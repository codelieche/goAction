## Go语言的面向对象

- go语言仅支持封装，不支持继承和多态
- 面向接口编程
- go语言没有class，只有struct


### 值接收者 VS 指针接收者

- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者
- 一致性：如果有指针接收者，最好都是指针接收者【建议】

**值接收者**是go的一个独特之处
指针接收者，其它语言都有，比如py的self，java中的变量引用等

```go
type People struct {
	age int
	name string
}

func (p People) info (){
	// 值传递
	fmt.Printf("Name is %s, age is %d\n", p.name, p.age)
}

func (p *People)setAge(age int) {
	// 指针传递
	p.age = age
}

func (p *People)setName(name string){
	// 指针传递
	p.name = name
}
```

### 封装
- Go语言名字一般使用CamelCase，python常使用caml_case
- 首字母大写：public
- 首字母小写：private
- 注意public/private是针对包来说的

#### 包
- 每个目录只能一个包【比如这个目录有了main包了，就不可加其它包了】
- main包包含可执行入口
- 为结构体定义的方法必须放在同一个包内
- 可以是不同文件

