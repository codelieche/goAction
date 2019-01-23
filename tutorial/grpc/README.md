```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```
安装protoc:  `https://github.com/protocolbuffers/protobuf/releases`

查看protoc版本：
```bash
protoc --version
```

- `protoc --go_out=plugins=grpc:. helloworld.proto`

