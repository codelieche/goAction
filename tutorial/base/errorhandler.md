## 错误处理

### error
大多数函数的最后一个返回值会是error类型，以便在错误情况下返回详细信息。  
error类型只是一个系统内置的`interface`。

```go
type error interface{
	Error() string
}
```

### defer
> defer的语句含义是不管程序是否出现异常，均在函数退出时自动执行相关代码。  
注意是：先进后出(执行)的。

### panic
> 和其它语言throw相通，抛出异常，没人接就抛出异常。

- 停止当前函数执行
- 一直向上返回，执行每一层的defer
- 如果没有遇到recover，程序退出

### recover
- 尽在defer调用中使用
- 获取panic的值
- 如果无法处理，可以重新panic


### error vs panic
- 意料之中的：使用error。如：文件打不开，不存在等
- 意料之外的：使用panic。如：数组越界
