package main

import "fmt"

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



func main() {
	p := People{name: "codelieche", age: 20}
	p.info()
	p.setAge(30)
	p.setName("codelieche.com")
	p.info()

	fmt.Println("p的名字：p.name: ", p.name)
}

