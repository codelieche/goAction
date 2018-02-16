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

**接口变量里面有什么**
接口变量：实现者的类型，实现者的值/指针。  

- 接口变量自带指针
- 接口变量同样采用值传递，几乎不需要实现接口的指针
- 指针接收者实现只能以指针方式使用；值接收者都可以

**查看接口变量**
- 表示任何类型：interface{}
- Type Assertion
- Type Switch



