package authserver

import "net/http"

// web服务的路由入口
func webRooter(w http.ResponseWriter, r *http.Request) {

	// 1. 处理
	HandleAuth(w, r)
}
