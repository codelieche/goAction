package authserver

import (
	"fmt"
	"net/http"
)

/**
Handle Auth
*/

// 校验auth的处理器
func auth(w http.ResponseWriter, r *http.Request) {
	// 1. 获取到session：
	session, err := store.Get(r, "usersession")
	if err != nil {
		// session 有误
		http.Error(w, err.Error(), 400)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "方法错误", 405)
		return
	} else {
		// 获取cookie

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			// 用户登录成功
			// 判断是否是可以登录本系统的用户
			username := session.Values["username"].(string)
			var canLogginSystem = userCanLoginSystem(username)

			if canLogginSystem {
				w.Write([]byte("校验成功"))
			} else {
				// 用户已没有权限访问本系统
				w.WriteHeader(403)
				msg := fmt.Sprintf("%s无权限访问本系统", username)
				w.Write([]byte(msg))
			}
			return
		} else {
			// 用户登录失败的
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(401)
			// 如果再w.Write之前未调用过WriteHeader，默认会是200
			w.Write([]byte(`{"detail": "Authentication credentials were not provided."}`))

			// 跳转去登录页
			//http.Redirect(w, r, "/account/login", 302)
			return

		}
	}

}
