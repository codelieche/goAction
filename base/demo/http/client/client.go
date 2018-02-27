package main

import (
	"fmt"
	"net/http"
)

func httpClientDemo01() {
	response, err := http.Get("http://www.codelieche.com")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println(response)

}

func main() {
	httpClientDemo01()

}
