package authserver

import (
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
			w.Write([]byte("登录成功"))
			return
		} else {
			// 用户登录失败的
			// 跳转去登录页
			http.Redirect(w, r, "/account/login", 301)
			return

		}
	}

}
