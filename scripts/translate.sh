#! /bin/sh

if [ $1 == "" ]; then
    echo "Pass a file argument"
    exit 1
fi

if [ ! -e $1 ]; then
    echo "File passed does not exist: $1"
    exit 1
fi

set -x
PIPELINE_BINARY="${PIPELINE_BINARY:-./bin/pipeline}"
set +x

if [ -n $PIPELINE_BINARY ]; then
    echo "Pipeline binary doesn't exist, you may need to compile it with 'make pipeline'"
    exit 1
fi

cat $1 | $(PIPELINE_BINARY) > 