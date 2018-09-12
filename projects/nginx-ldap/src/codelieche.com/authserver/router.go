package authserver

import (
	"net/http"

	"codelieche.com/webhealth"
)

// web服务的路由入口
func webRooter(w http.ResponseWriter, r *http.Request) {

	// 1. 处理
	switch {
	case r.URL.Path == "/account/auth" || r.URL.Path == "/account/auth/":
		auth(w, r)
	case r.URL.Path == "/account/login" || r.URL.Path == "/account/login/":
		login(w, r)
	case r.URL.Path == "/account/logout" || r.URL.Path == "/account/logout/":
		logout(w, r)
	case r.URL.Path == "/account/users" || r.URL.Path == "/account/users/":
		handlerUsers(w, r)
	case r.URL.Path == "/health" || r.URL.Path == "/health/":
		webhealth.CheckHealth(w, r)

	default:
		http.Error(w, "Page Not Found", 404)
		return
	}
}
