package ldaplib

import "testing"

func TestAuth(t *testing.T) {
	if Auth("codelieche", "password") {
		t.Log("登录成功")
	} else {
		t.Error("登录失败")
	}
}
