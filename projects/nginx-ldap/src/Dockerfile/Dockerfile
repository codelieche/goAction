FROM nginx:alpine

# 设置环境变量
ENV LANG='C.UTF-8' PYTHONIOENCODING='UTF-8' TZ='Asia/Shanghai'

# 设置工作目录
WORKDIR /var/backup

RUN apk add curl && rm -rf /var/cache/apk/*

# 复制相关文件
# COPY ["./app/", "./conf/", "./entrypoint.sh", "/var/backup/"]
COPY ./backup /var/backup

RUN mv /var/backup/conf/nginx.conf /etc/nginx/nginx.conf && \
    chmod -R 755 /var/backup && \
    mkdir -p /data/etc/nginx && mkdir -p /data/logs/nginx/ && \
    mkdir -p /data/app/

EXPOSE 80
VOLUME ["/data/"]
ENTRYPOINT ["/var/backup/entrypoint.sh"]

# LDAP校验的程序放在：/var/backup/app/app
# app运行命令：/var/backup/app/app -config /data/conf/nginx-ldap.conf

# 启动nginx
CMD ["nginx", "-g", "daemon off;"]
# CMD ["/var/backup/app/app"]