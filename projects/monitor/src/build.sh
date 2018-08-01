# !/bin/bash

# 第1步：打包golang monitor
# 1-1：进入文件路径
cd ./codelieche.com/entry
echo "`date +"%F %T"`: 当前目录 ${PWD}"

# 1-2: 执行构建命令
echo "`date +"%F %T"`: 开始构建"

GOOS=linux GOARCH=amd64 go build ./monitor.go && echo "`date +"%F %T"`：构建成功"

# 1-3: 重命名打包后的文件
mv ./monitor ./webMonitor
echo ""
ls -al
echo ""

# 第2步：构建Docker镜像
# 2-1：移动文件到 Docker
mv ./webMonitor ../Docker

# 2-2：进入Docker目录
cd ../Docker

# 2-3：构建镜像
docker build -t ops-monitor:v1-alpine ./

# 2-4：删除golang打包的文件
rm ./webMonitor

# 最后输出结束信息
echo "`date +"%F %T"`: 执行完毕！"
