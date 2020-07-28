#!/usr/bin/env bash

url=$1
problem=$(basename $url)
# dir=$(echo ${problem} | sed -E 's/_/\//g')
dir=$(echo ${problem^^} | sed -E 's/_/\//g')
echo $dir

mkdir -p $dir
touch $dir/main.go
