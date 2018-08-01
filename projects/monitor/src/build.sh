# !/bin/bash
# 进入文件路径
cd ./codelieche.com/entry
echo "当前目录：${PWD}"

# 执行构建命令
echo "`date +"%F %T"`: 开始构建"

GOOS=linux GOARCH=amd64 go build ./monitor.go && echo "`date +"%F %T"`：构建成功"

# 重命名打包后的文件
mv ./monitor ./webMonitor
echo ""

ls -al

# 移动文件到 Docker
mv ./webMonitor ../Docker

# 进入Docker目录
cd ../Docker

# 构建镜像
docker build -t ops-monitor:v1-alpine ./

