package authserver

import "encoding/json"

type LoginResponse struct {
	Status  bool   // 登录成功状态
	Message string // 相关消息
}

func (lr *LoginResponse) Marshal() []byte {
	if data, err := json.Marshal(lr); err != nil {
		return []byte("")
	} else {
		return data
	}
}
