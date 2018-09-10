package authserver

import (
	"html/template"
	"net/http"

	"fmt"

	"log"

	"codelieche.com/ldaplib"
)

// 账号登录
// URL是：/account/login
// 支持防范：POST/GET
func login(w http.ResponseWriter, r *http.Request) {
	// 1. 获取到session：
	session, err := store.Get(r, "usersession")
	if err != nil {
		// session 有误
		http.Error(w, err.Error(), 400)
		return
	}
	// 1. POST登录：

	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username == "" || password == "" {
			lr := LoginResponse{false, "用户名/密码为空"}
			w.WriteHeader(400)
			w.Write(lr.Marshal())
			return
		}

		// ldap verify
		if result := ldaplib.Auth(username, password); result {
			// 设置session
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
			// 校验成功
			lr := LoginResponse{true, "登录成功:" + username}
			//w.Write(lr.Marshal())
			if t, err := template.ParseFiles("templates/login.html"); err != nil {
				log.Println(err)
				msg := fmt.Sprintf("加载模板出错: %s", err.Error())
				http.Error(w, msg, 500)
				return
			} else {
				t.Execute(w, lr)
				return
			}
			return
		} else {
			// 登录失败

			// 设置session
			session.Values["authenticated"] = false
			session.Save(r, w)
			lr := LoginResponse{false, "登录失败:用户名/密码错误"}
			w.WriteHeader(400)
			// 模板渲染登录结果
			if t, err := template.ParseFiles("templates/login.html"); err != nil {
				log.Println(err)
				msg := fmt.Sprintf("加载模板出错: %s", err.Error())
				http.Error(w, msg, 500)
				return
			} else {
				t.Execute(w, lr)
				return
			}
			//w.Write(lr.Marshal())
			return
		}
	} else {
		// 获取cookie
		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			// 用户登录成功
			msg := fmt.Sprintf("登录用户:%s", session.Values["username"])
			lr := LoginResponse{true, msg}
			//w.Write(lr.Marshal())
			// 模板渲染登录结果
			if t, err := template.ParseFiles("templates/login.html"); err != nil {
				log.Println(err)
				msg := fmt.Sprintf("加载模板出错: %s", err.Error())
				http.Error(w, msg, 500)
				return
			} else {
				t.Execute(w, lr)
				return
			}
			return
		} else {
			// 用户登录失败的：渲染登录页面
			if t, err := template.ParseFiles("templates/login.html"); err != nil {
				log.Println(err)
				msg := fmt.Sprintf("加载模板出错: %s", err.Error())
				http.Error(w, msg, 500)
				return
			} else {
				t.Execute(w, nil)
				return
			}
		}
	}

}
