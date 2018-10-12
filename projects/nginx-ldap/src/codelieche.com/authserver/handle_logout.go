package authserver

import "net/http"

// 用户退出登录
func logout(w http.ResponseWriter, r *http.Request) {
	// 1. 获取session
	if session, err := store.Get(r, "usersession"); err != nil {
		// 获取sessionid出错
		// log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	} else {
		session.Values["authenticated"] = false
		session.Values["username"] = ""
		session.Save(r, w)

		http.Redirect(w, r, "/account/login", 302)
	}

	// 2. 设置authenticated为False
}
