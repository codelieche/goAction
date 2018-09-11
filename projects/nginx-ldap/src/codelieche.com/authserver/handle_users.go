package authserver

import (
	"net/http"

	"fmt"
	"strings"

	"codelieche.com/settings"
)

/**
可以登录的User 列表
*/

var usersListStr = settings.Config.Admin

func handlerUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(usersListStr))
		return
	} else if r.Method == "POST" {
		// 简单方式：就取出users然后修改新的usersListStr, 根据逗号分隔多个用户
		users := r.PostFormValue("users")
		if strings.Index(users, settings.Config.Admin) < 0 {
			users = fmt.Sprintf("%s,%s", settings.Config.Admin, users)
		}
		// 修改usersListStr
		usersListStr = users
		// 返回结果
		w.Write([]byte(usersListStr))
		return
	} else {
		http.Error(w, "Method Not Allow", 405)
		return
	}
}
