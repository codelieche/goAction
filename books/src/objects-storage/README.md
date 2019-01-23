## 《分布式对象存储》
> 分布式对象存储--原理、架构及Go语言实现。

### 第1章：对象存储简介
实现一个简单的对象获取和上传接口：
1. `GET`: 下载对象： `GET /objects/:name`
2. `PUT`: 上传对象: `PUT /objects/:name`

入口文件：
- 启动服务端端：`chapter01/entry/main.go`
- 测试PUT功能：`chapter01/entry/put_objects.sh`
- 测试GET功能：`chapter01/entry/get_objects.sh`

启动服务端：
```bash
go run chapter01/entry/main.go
```



