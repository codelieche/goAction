package objects

import (
	"net/http"
	"strings"
)

// 服务端的入口函数
func handlerRouter(w http.ResponseWriter, r *http.Request) {
	// 1. 根据路径处理
	switch {
	// URL是/objects/开头的函数
	case strings.HasPrefix(r.URL.Path, "/objects/"):
		if r.Method == "PUT" {
			handlerPutObject(w, r)
		} else if r.Method == "GET" {
			handlerGetObject(w, r)
		} else {
			http.Error(w, "方法不支持", 405)
		}
	default:
		http.Error(w, "Not Found", 404)
	}
}
