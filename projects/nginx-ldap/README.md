## Nginx LDAP


### 背景
> 有些服务是没有认证功能的，比如：`elasticsearch`, `kibana`.  
这个时候我们如果想给它们加入个校验功能。就可以用到nginx的`http_auth_request`模块。

### 参考文档
- [ngx_http_auth_request_module](http://nginx.org/en/docs/http/ngx_http_auth_request_module.html)
- [nginx-ldap-auth](https://github.com/nginxinc/nginx-ldap-auth)


