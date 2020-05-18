package settings

import (
	"encoding/json"
	"log"

	"fmt"

	"github.com/levigross/grequests"
)

/**
登录web系统
*/

type LoginResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// 登录web系统，并返回session指针
func Login() (*grequests.Session, error) {
	// 第1步：先实例化Session
	sRo := &grequests.RequestOptions{
		//RequestTimeout: time.Duration(5000 * time.Millisecond),
	}
	session := grequests.NewSession(sRo)

	// 第2步：登录系统
	// 2-1: 请求选项
	loginRo := &grequests.RequestOptions{
		JSON:    map[string]string{"username": Config.Web.UserName, "password": Config.Web.PassWord},
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	// 2-2: 登录请求
	if resp, err := session.Post(Config.Web.GetLoginUrl(), loginRo); err != nil {
		log.Println("登录失败")
		return nil, err
	} else {
		// body := resp.String()
		data := LoginResponse{}
		if err := json.Unmarshal(resp.Bytes(), &data); err != nil {
			log.Println(err.Error())
			return nil, err
		} else {
			if data.Status == "success" {
				// 登录成功
				return session, nil
			} else {
				// 登录失败
				message := fmt.Sprintf("登录失败:%s", data.Message)
				log.Println(message)
				return nil, fmt.Errorf(message)
			}
		}
	}
}
