#!/bin/bash

# 上传文件

function put_object(){
   echo "上传的文件名是： $1"
   echo "上传文件的内容是：$*"

   # 通过curl上传对象到服务器
   url="http://127.0.0.1:9000/objects/${1}"
   content="上传文件的内容是：$*"
   curl -v $url -XPUT -d "$content"
}

for i in {1..10}
do
    name="file$i"
    content="File Content $i"
    put_object $name $content
    sleep $i
done

echo "执行完毕"