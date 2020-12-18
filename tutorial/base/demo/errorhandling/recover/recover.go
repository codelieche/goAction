package main

import (
	"errors"
	"fmt"
)

func throwError() {
	panic(errors.New("This is a error"))

	//a := 0
	//b := 100 / a
	//fmt.Println(b)

	//panic("abcdefg")
}

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("r is nil")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do %v\n", r))
		}
	}()
	throwError()
}

func main() {
	tryRecover()
}
