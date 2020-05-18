package settings

import (
	"flag"
	"log"
)

/**
从配置文件中获取LDAP的相关信息
*/

var Config = LdapConfig{}

// 获取配置内容
func getConfigFromFile() {
	var filePath = flag.String("config", "./config.json", "请传入配置文件的路径")
	flag.Parse()

	if err := Config.LoadFromFile(*filePath); err != nil {
		log.Print("获取配置内容出错:", err.Error())
		panic(err)
	} else {
		// 获取配置成功
		log.Println("读取配置成功！")
	}
}

func init() {
	getConfigFromFile()
}
