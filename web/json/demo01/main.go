package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Admin struct {
	Users []*User `json:"users"`
	Name  string  `json:"name"`
}

func main() {
	// 有2个用户 Lisi，WangWu
	u1 := User{"Lisi", 18}
	u2 := User{"WangeWu", 28}

	a := Admin{[]*User{&u1, &u2}, "管理员"}

	if data, err := json.Marshal(a); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(string(data))
	}
}
