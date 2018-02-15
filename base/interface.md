## 接口（interface）

### 接口的定义

```go
type Retriever interface {
	Get(source string) string
}

func download(retriever Retriever) string {
	return retriever.Get("http://www.codelieche.com")
}
```

### 接口的实现

- 接口的实现是隐式的
- 只要实现接口里面的方法

