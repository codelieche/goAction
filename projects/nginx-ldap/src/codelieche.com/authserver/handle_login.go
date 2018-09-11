package authserver

import (
	"html/template"
	"net/http"

	"fmt"

	"log"

	"strings"

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
		nextUrl := r.PostFormValue("next")

		if username == "" || password == "" {
			lr := LoginResponse{false, "用户名/密码为空", nextUrl}
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
			if nextUrl != "" {
				// 跳转链接
				http.Redirect(w, r, nextUrl, 302)
				return
			} else {
				// 登录成功
				lr := LoginResponse{true, "登录成功:" + username, nextUrl}
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
			}
		} else {
			// 登录失败

			// 设置session
			session.Values["authenticated"] = false
			session.Save(r, w)
			lr := LoginResponse{false, "登录失败:用户名/密码错误", nextUrl}
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
		r.ParseForm()
		nextUrl := r.Form.Get("next")
		if nextUrl == "" {
			nextUrl = r.Header.Get("X-Next")
		}
		// 获取cookie
		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			if nextUrl != "" && !strings.HasPrefix(nextUrl, "/account/login") {
				// 跳转链接
				http.Redirect(w, r, nextUrl, 302)
				return
			} else {
				// 用户登录成功
				username := session.Values["username"].(string)
				lr := LoginResponse{true, "登录用户:" + username, nextUrl}
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
			}

		} else {
			// 用户未登录：渲染登录页面
			if t, err := template.ParseFiles("templates/login.html"); err != nil {
				log.Println(err)
				msg := fmt.Sprintf("加载模板出错: %s", err.Error())
				http.Error(w, msg, 500)
				return
			} else {
				var lr LoginResponse
				if nextUrl != "" && !strings.HasPrefix(nextUrl, "/account/login") {
					lr = LoginResponse{false, "", nextUrl}
				}
				t.Execute(w, lr)
				return
			}
		}
	}

}
