#!/bin/bash

# 第1步：打包golang程序
# 1-1: 进入文件路径
cd ../entry
echo "`date +%"F %T"`: 当前目录 ${PWD}"

# 1-2: 执行go build命令
echo "`date +%"F %T"`: 开始打包nginxldap"
GOOS=linux GOARCH=amd64 go build ./nginxldap.go && echo "`date +%"F %T"`: 构建成功" || echo "`date +%"F %T"`: 构建失败！！！"

# 1-3: 重命名打包后的文件(而且重命名为app)
mv ./nginxldap ../src/Dockerfile/backup/app/app || exit 1
echo ""

# 第2步：构建镜像
# 2-1：进入Dockerfile的目录
cd ../src/Dockerfile
echo "`date +%"F %T"`: 当前目录 ${PWD}"

# 2-2：构建镜像
docker build -t nginxldap:v1 ./

# 2-3: 构建完后删除文件
# rm ./backup/app/app

# 第3步：最后输出结束信息
echo "`date +"%F %T"`: 执行完毕！"

