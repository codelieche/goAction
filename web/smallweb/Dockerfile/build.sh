#!/bin/bash
# 构建smallweb镜像

# 第1步：定义变量
NAME=smallweb
TAG=v1

# 第2步：进入entry打包程序
# 2-1: 打包
cd ../entry
GOOS=linux GOARCH=amd64 go build ./smallweb.go && echo "`date +%"F %T"`: 构建成功" || (echo "`date +%"F %T"`: 构建失败！！！" && exit 1)
# 2-2： 把打包后的文件移动到Dockerfile目录
mv ./smallweb ../Dockerfile

# 第3步：进入Dockerfile目录
# 3-1: 进入目录
cd ../Dockerfile
# 3-2：执行构建镜像
docker build . -t "${NAME}:${TAG}"
# 3-3: 删除打包的文件
rm ./smallweb

# 第4步：执行查看镜像命令
docker images | grep $NAME


