package main

import (
	"goAction/base/demo/errorhandling/filelistingserver/filelisting"
	"net/http"
	"os"

	"fmt"

	"github.com/gpmgo/gopm/modules/log"
)

type appHander func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHander) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// recover 错误
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		code := http.StatusOK
		if err != nil {
			// 加入log
			log.Warn("Error handling request: %s", err.Error())
			// 如果是自定义的userError
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// 处理err
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				// 404
			case os.IsPermission(err):
				// 没有权限：403
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	// 这边定义的接口，只管要实现个error, 还要个Message()
	// 用的人自己管自己的实现，两者不需要互相能看见
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

	// http://127.0.0.1:8888/list/tmp/fib.txt
}
