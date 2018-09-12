## Nginx LDAP


### 背景
> 有些服务是没有认证功能的，比如：`elasticsearch`, `kibana`.  
这个时候我们如果想给它们加入个校验功能。就可以用到nginx的`http_auth_request`模块。

### 依赖
- ldap.v2

安装：`go get gopkg.in/ldap.v2`

- gorilla/sessions

安装：`go get github.com/gorilla/sessions`

### nginx.conf

```bash
upstream kibana-server {
    server 127.0.0.1:8000;
}

server {
    linsten 80;
    server_name codelieche.com;
    
    access_log /data/logs/access.log;
    
    location / {
        auth_request /account/auth;
        
        error_page 401 403=200 /account/login;
        
        proxy_pass kibana-server;
    }
    
    location /account/login {
        proxy_pass http://127.0.0.1:9000/account/login;
        proxy_set_header X-Next $request_uri;
    }
    
    location /account/auth {
        proxy_pass http://127.0.0.1:9000;
        # proxy_cache auth_cache;
        
        proxy_pass_request_body off;
        proxy_set_header Content-Length "";
        proxy_set_header X-CookieName "usersession";
        proxy_set_header Cookie usersession=$cookie_usersession;
    }
}
```

/etc/nginx/conf.d/default.conf

### 参考文档
- [ngx_http_auth_request_module](http://nginx.org/en/docs/http/ngx_http_auth_request_module.html)
- [ldap.V2](https://godoc.org/gopkg.in/ldap.v2)
- [nginx-ldap-auth](https://github.com/nginxinc/nginx-ldap-auth)
- [gorialla/sessions](https://github.com/gorilla/sessions)


### 使用

```bash
docker run -itd --name nginxldap -p 9090:80 nginxldap:v1
docker exec -it nginxldap /bin/sh

docker stop nginxldap
docker rm nginxldap
docker rmi nginxldap:v1
```