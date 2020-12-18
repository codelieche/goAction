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


在Go语言中，只要两个接口拥有相同的方法列表，那么它们就是相同的，可以相互赋值。

```go
package one

type ReadWrite interface {
	Read(buf [] byte) (n int, err error)
	Writer(buf [] byte)(n int, err error)
}
```

第二个接口：
```go
package two
type IStream interface {
	Write(buf [] byte) (n int, err error)
	Read(buf [] byte) (n int, err error)
}
```
在这里我们定义了两个接口，一个叫one.ReadWriter, 一个叫two.IStream,两者都定义了Read()和Write()方法，  
只是定义的次序相反。  
在Go语言中，这两个接口实际上并无区别。

