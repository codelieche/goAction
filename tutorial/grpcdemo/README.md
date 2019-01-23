## Grpc的基本使用

### 安装
```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```
安装protoc:  `https://github.com/protocolbuffers/protobuf/releases`

查看protoc版本：
```bash
protoc --version
```

### 示例

#### proto

- permission.proto

```
syntax = "proto3";

package proto;

service PermissionCheck {
    rpc CheckPermission(CheckRequest) returns (CheckResponse){}
}

message CheckRequest {
    string username = 1;
    string permission = 2;
}

message CheckResponse {
    bool status = 1;
    string message = 2;
}
```

- 生成相关文件
```bash
protoc --go_out=plugins=grpc:. *.proto
```

#### Service


启动服务报错，说没2个包，安装：

```bash
go get github.com/sirupsen/logrus
go get gopkg.in/alecthomas/kingpin.v2
```

#### Client



### 参考文档
- [grpc](https://grpc.io/)
- [proto releases](https://github.com/protocolbuffers/protobuf/releases)
- [protocol buffers proto3](https://developers.google.com/protocol-buffers/docs/proto3)
