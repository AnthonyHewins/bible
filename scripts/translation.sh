mkdir -p parquet
cat $1 | ./bin/pipeline > parquet/$2.parquet