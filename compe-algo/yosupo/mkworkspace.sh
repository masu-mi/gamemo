#!/Users/masumi/local/bin/bash

url=$1
problem=$(basename $url)
dir=$(echo ${problem})

mkdir -p $dir
touch $dir/main.go
