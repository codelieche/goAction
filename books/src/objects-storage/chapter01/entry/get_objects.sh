#!/bin/bash

function get_object(){
   url="http://127.0.0.1:9000/objects/$1"
   echo $url
   curl $url
}

echo "开始下载图片"

for i in {1..10}
do
  file="file${i}"
  get_object $file
  sleep $i
done;

echo "执行完毕"
