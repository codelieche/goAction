## 构建镜像


## 使用

```bash
docker run -e "DOMAIN=http://192.168.1.101:8080" \
 -e "USERNAME=admin" -e "PASSWORD=pwd"  -p 9000:9000 \
  -e "DB=devops" -e "DBHOST=192.168.1.101" -e "DBUSER=admin" -e "DBPWD=***"\
 -itd  --name monitor ops-monitor:v1-alpine
```

## 执行镜像遇到的问题
> failed to load system roots and no roots provided

**解决方式：**

```
RUN apk add --no-cache ca-certificates 
```