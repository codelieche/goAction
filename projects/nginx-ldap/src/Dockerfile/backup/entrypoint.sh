#!/bin/sh

# 1. 检查日志目录是否为空
# 1-1: 检查目录
if [ ! -d /data/logs/nginx ];then
    mkdir -p /data/logs/nginx;
fi

# 1-2：检查文件
if [ `ls /data/logs/nginx/ | wc -l` -eq 0];
then
    touch /data/logs/nginx/access.log /data/logs/nginx/error.log;
    chown nginx:nginx -R /data/logs/nginx;
fi;

# 2. 检查nginx的ldap app配置文件
# 2-1: 检查目录
if [ ! -d /data/app ];then
    mkdir -p /data/app;
fi

# 2-2: 检查文件
if [ `ls /data/app/ | wc -l` -eq 0 ];
then
    cp -rf /var/backup/app/* /data/app/
fi;

# 3. 检查etc配置文件
# 3-1: 检查目录
if [ ! -d /data/etc/nginx ];then
    mkdir -p /data/etc/nginx;
fi

# 3-2: 检查文件
if [ `ls /data/etc/nginx/ | wc -l` -eq 0 ];
then
    cp -rf /var/backup/conf/* /data/etc/nginx/
fi;

# 检查web服务是否启动
if [ `netstat -lnput | grep 9000 | wc -l` -eq 0 ];
then
    cd /data/app && nohup ./app --config=./config.json & ls
fi;

exec "$@"