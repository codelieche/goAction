package authserver

import (
	"net/http"

	"fmt"
	"strings"

	"goAction/projects/nginx-ldap/web/settings"
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

func userCanLoginSystem(username string) bool {
	// 判断是否是可以登录本系统的用户
	usersList := strings.Split(usersListStr, ",")
	// 先使用O(N)算法来执行
	for _, i := range usersList {
		if i == username {
			return true
		}
	}
	// 跳出循环了，表示未搜寻到用户
	return false
}
