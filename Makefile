.DEFAULT: all
.PHONY: fmt test gen clean run help sql docker

# command aliases
test := CONFIG_ENV=test go test ./...

targets := bible pipeline

VERSION ?= v?.?.?
COMMIT ?= $(shell git rev-list -1 HEAD)
ifneq (,$(wildcard ./vendor))
	$(info Found vendor directory; setting "-mod vendor" to any "go build" commands)
	BUILD_FLAGS += -mod vendor
endif

#=================================
# Assets
#=================================
filename := master.zip
download: ## Download all open bible translations, unzip
	wget https://github.com/gratis-bible/bible/archive/refs/heads/master.zip
	unzip -tq $(filename)
	unzip $(filename)
	rm $(filename)
	mv bible-master translations

LANG ?= en
outdir=internal/codex
translate-lang: pipeline ## translate all files in $LANG (defaults to en)
	mkdir -p protos/$(LANG)
	for i in $(shell find translations/$(LANG) -type f -iname "*.xml"); do \
	 	file=$$(basename $$i); \
		out=$(outdir)/$(LANG)/$${file%.xml}.pbf.gz; \
		cat $$i | ./bin/pipeline | gzip > $$out; \
		echo "Completed $$i -> $$out"; \
	done

#======================================
# Builds
#======================================
$(targets): ## Build a target server binary
	go build $(BUILD_FLAGS) -ldflags "-X main.version=$(COMMIT)" -o bin/$@ ./cmd/$@

all: $(targets) ## Build all targets

#======================================
# Running
#======================================
run-%: ## Run the server using .env variables
	export $$(cat .env | xargs) && ./bin/$(patsubst run-%,%,$@)

up-compose: ## Run a binary with docker compose
	docker-compose -f ./docker/compose.yaml up

proto: ## Generate proto files
	buf generate

#======================================
# App hygiene
#======================================
clean: ## gofmt, go generate, then go mod tidy, and finally rm -rf bin/
	find . -iname *.go -type f -exec gofmt -w -s {} \;
	go generate ./...
	go mod tidy
	rm -rf ./bin

test: ## Run go vet, then test all files
	go vet ./...
	$(test)

help: ## Print help
	@printf "\033[36m%-30s\033[0m %s\n" "(target)" "Build a target binary in current arch for running locally: $(targets)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
