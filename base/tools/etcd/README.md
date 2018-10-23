## Golang 操作etcd

- 启动etcd：

```bash
etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379'
```

- 使用etcdctl 
> Set environment variable ETCDCTL_API=3 to use v3 API or ETCDCTL_API=2 to use v2 API.

```bash
 ETCDCTL_API=3 etcdctl set /study/t1 VALUE01

etcdctl get /study/t1
ETCDCTL_API=3 etcdctl put "/study/t1" "VALUE01-02"

ETCDCTL_API=3 etcdctl get /study --prefix
```
