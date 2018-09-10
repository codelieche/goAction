package authserver

import (
	"net/http"

	"log"

	"codelieche.com/ldaplib"
)

/**
Handle Auth
*/

// 校验auth的处理器
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	// 1. POST登录：
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username == "" || password == "" {
			lr := LoginResponse{false, "用户名/密码为空"}
			w.Write(lr.Marshal())
			w.WriteHeader(400)
			return
		}

		// 开登录
		if loginSuccess := ldaplib.Auth(username, password); loginSuccess {
			// 登录成功
			lr := LoginResponse{true, "登录成功:" + username}
			w.Write(lr.Marshal())
			return
		} else {
			// 登录成功
			lr := LoginResponse{true, "登录失败:用户名/密码错误"}
			w.Write(lr.Marshal())
			w.WriteHeader(400)
			return
		}
	} else {
		// 获取cookie
		if sessionIdCookie, err := r.Cookie("sessionid"); err != nil {
			log.Println(err)
		} else {
			// 返回sessionid
			w.Write([]byte("Session Id Is:" + sessionIdCookie.Value))
		}
	}

}
