#! /bin/bash

set -euo pipefail

base="$(basename $0)"
dir="$(dirname $0)"

input=./translations
output=./parquet
bin=./bin/pipeline

help() { cat <<EOF
usage: $base ARGS [FLAGS]

Take XMLs and turn it into parquet for accessibility
in the bible binary

help, -h, --help				Display help text
in DIR                          Input directory (default $input)
out DIR                         Output directory (default $output)
bin PATH                        Pipeline bin path (default $bin, will try make if not exist)
EOF
exit $1
}

while [[ "$#" -gt 0 && "${1:-}" != "--" ]]; do case $1 in
    help | -h | --help ) help 0;;
    bin ) bin=$1;;
    in ) in=$1 ;;
    out ) out=$1 ;;
    *) echo "Unknown arg $1"; help 1;;
esac; shift; done
if [[ "${1:-}" == '--' ]]; then shift; fi

if [[ -f $bin ]]; then
    echo "Making pipeline cli..."
    make pipeline
    bin=./bin/pipeline
fi

echo Making output dir
mkdir -p $output

for i in $(find $input -iname "*.xml" -type f -printf "%P "); do
    x="$output/$(dirname $i)"
    mkdir -p $x
    echo "Translating $input/$i"
    cat "$input/$i" | $bin > "$x/$(basename $i).parquet"
done