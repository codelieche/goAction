package authserver

/**
这是个简单的ldap auth服务
一般就是：GET/POST，POST是用户登录、GET是校验用户是否登录
*/

import (
	"log"
	"net/http"
)

// 检查session的中间件
func checkSessionMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c, err := r.Cookie("usersession"); err != nil {
			// sessionid 不存在: 给请求添加一个

		} else {
			//log.Println(s.Name, s.Value)
			c.MaxAge = -100
		}

		// 1. 判断是POST还是其它方法
		if session, err := store.Get(r, "usersession"); err != nil {
			// 获取sessionid出错
			// log.Println(err)
			//http.Error(w, err.Error(), 400)
			//return
		} else {
			// log.Println(session.Values)
			session.Save(r, w)
			// 如果是GET访问，而且session是新的，就直接返回未登录吧
			if r.Method == "GET" && r.URL.Path == "/account/auth" {
				if session.IsNew {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(401)
					w.Write([]byte(`{"detail": "Authentication credentials were not provided."}`))
					//http.Error(w, `{"detail": "Authentication credentials were not provided."}`, 401)
					return
					// 跳转到登录页面
					//http.Redirect(w, r, "/account/login", 302)
					//return
				}
			}
		}
		// Next Handler
		next.ServeHTTP(w, r)
	})
}

// 检查session的中间件：
// 1. GET请求，没有sessionid的cookie，设置一个后直接返回401错误
// 2. POST请求，没有sessionid的cookie:
//    给响应设置个sessionid的cookie，同时加入到request中，方便后续调用
func checkSimpleSessionMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 判断是POST还是其它的方法
		if r.Method == "POST" {
			if _, err := r.Cookie("sessionid"); err != nil {
				// sessionid 不存在: 给请求添加一个
				sessionIdCookie := http.Cookie{
					Name:   "sessionid",
					Value:  GenerateSessionId(),
					MaxAge: 24 * 3600,
				}
				http.SetCookie(w, &sessionIdCookie)
				// 同时加入到请求之中
				r.AddCookie(&sessionIdCookie)
			} else {
				//log.Println(s.Name, s.Value)
			}
		}

		if r.Method == "GET" {
			if s, err := r.Cookie("sessionid"); err != nil {
				log.Print(err)
				sessionIdCookie := http.Cookie{
					Name:   "sessionid",
					Value:  GenerateSessionId(),
					MaxAge: 24 * 3600,
				}
				http.SetCookie(w, &sessionIdCookie)
				// 请求没有sessionid，那肯定是未登录，直接返回401
				w.Header().Add("Content-Type", "application/json")
				http.Error(w, `{"detail": "Authentication credentials were not provided."}`, 401)
				return
			} else {
				log.Println(s.Name, s.Value)
			}
		}
		// Next Handler
		next.ServeHTTP(w, r)
	})
}
