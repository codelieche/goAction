package ldaplib

import (
	"crypto/tls"
	"fmt"
	"log"

	"codelieche.com/settings"
	"gopkg.in/ldap.v2"
)

// 认证用户
// 传入用户名和密码
func Auth(username string, password string) bool {

	addr := fmt.Sprintf("%s:%d", settings.Config.LdapServer, settings.Config.LdapPort)
	l, err := ldap.Dial("tcp", addr)
	if err != nil {
		log.Print("连接Ldap Server出错：", err.Error())
		return false
	}
	defer l.Close()

	// Reconnect with TLS
	// 建立StartTLS连接，这个是建立纯文本的TLS协议，
	// 允许将非加密的通讯升级为TLS加密，而不需要另外使用一个新的端口
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	// First bind with a read only user
	// 首先我们bind账号：禁止匿名查询，就需要先bind个
	if err = l.Bind(settings.Config.BindDN, settings.Config.BindPW); err != nil {
		// Bind Error
		log.Println("Bind Failed:", err.Error())
		return false
	}

	filterString := fmt.Sprintf(settings.Config.FilterStr, username)

	searchRequest := ldap.NewSearchRequest(
		// 这里是baseDN,我们将从这个节点开始搜索：eg: "DC=example,DC=com"
		settings.Config.BaseDN,
		// 参数：scope, defrefAliases, sizeLimit, timeLimit, typesOnly
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// 我们需要认证的用户名filter字符
		filterString,
		// Attributes: 这里是查询返回的属性，以数组的形式提供，如果为空就会返回所有的属性：eg: []string{"dn"} / nil
		[]string{"dn"},
		nil,
	)

	// 开始搜索
	if searchResult, err := l.Search(searchRequest); err != nil {
		// 搜索出错
		log.Printf("搜索失败：%s\n", err.Error())
		return false
	} else {
		if len(searchResult.Entries) != 1 {
			// 搜索的用户不存在或者是多个值
			log.Println("搜索的结果长度不等于1")
			return false
		}

		// 得到要获取用户的DN
		userDN := searchResult.Entries[0].DN
		//log.Println(userDN)

		// Bind as the user to verify their password
		// 拿这个 dn 和它的密码去做 bind 验证
		if err := l.Bind(userDN, password); err != nil {
			log.Println("密码不正确:", err.Error())
			return false
		} else {
			// 只有到这里才是正确的
			return true

			// Rebind as ther read only user for any further queries
			// 后续还要操作就用最初bind的账号重新bind回来，恢复出事的权限
		}
	}
}
