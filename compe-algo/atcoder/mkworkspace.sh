#!/Users/masumi/local/bin/bash

url=$1
problem=$(basename $url)
dir=$(echo ${problem^^} | sed -E 's/_/\//g')

mkdir -p $dir
touch $dir/main.go
