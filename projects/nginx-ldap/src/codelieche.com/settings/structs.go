package settings

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// LDAP配置: 默认先不使用SSL
type LdapConfig struct {
	// Ldap Server
	LdapServer string // LDAP的服务器地址
	LdapPort   int    // LDAP的端口号，默认一般是389
	BaseDN     string // 我们将从这个节点开始搜索：eg: "DC=example,DC=com"
	// 用来获取查询权限的 bind 用户。如果 ldap禁止了匿名查询，那我们就需要先bind个用户才可以查询
	BindDN    string // bind 的账号通常要使用完整的DN信息：比如：cn=Admin,dc=example,dc=org
	BindPW    string
	FilterStr string // (sAMAccountName=%s) 一般是这样的
	UserSSL   bool   // 是否使用加密传输

	WebSecretKey string // web生成session的密匙
}

// 从文件读取配置内容
func (config *LdapConfig) LoadFromFile(path string) error {
	// 从文件读取配置文件
	if file, err := os.Open(path); err != nil {
		msg := fmt.Sprintf("读取配置文件出错：%s", path)
		log.Println(msg)
		log.Println(err.Error())
		return err
	} else {
		// 关闭文件
		defer file.Close()
		// 解析文件
		if err := json.NewDecoder(file).Decode(config); err != nil {
			log.Println("解析配置文件出错！")
			log.Println(err.Error())
			return err
		}

		// 没出错的话返回nil
		return nil
	}
}
