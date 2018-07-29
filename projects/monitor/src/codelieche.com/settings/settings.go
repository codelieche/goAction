package settings

import (
	"flag"
	"fmt"
	"strings"
)

/**
监控项目配置相关数据:
settings感觉搞复杂了，但是为了可维护，好传参，先这样
*/

/**
后端api要用到的数据
- LoginUrl: 登录地址
- UserName：用户名
- UserPassword: 用户密码
- MonitorListUrl：监控列表地址
- FreshListInteraval: 刷新列表的频率
*/

type WebConfig struct {
	// 环境, Domain
	Env, Domain, UserName, PassWord string
	Interval                        int
	LoginUrl                        string
	MonitorListUrl                  string
	EventCreateUrl                  string
	EventAutoFixUrl                 string
}

func (c *WebConfig) GetLoginUrl() string {
	c.checkDomain()
	loginUrl := fmt.Sprintf("%sapi/1.0/account/login", c.Domain)
	return loginUrl
}

func (c *WebConfig) GetMonitorListUrl() string {
	// 检查下域名前缀与后缀
	c.checkDomain()
	listUrl := fmt.Sprintf("%sapi/1.0/monitor/all", c.Domain)
	return listUrl
}

func (c *WebConfig) checkDomain() {
	// 检查域名
	if !strings.HasPrefix(c.Domain, "http") {
		c.Domain = fmt.Sprintf("http://%s", c.Domain)
	}
	if !strings.HasSuffix(c.Domain, "/") {
		c.Domain = fmt.Sprintf("%s/", c.Domain)
	}
}

type DatabaseConfig struct {
	Host           string
	Port           int
	Database       string
	User           string
	Password       string
	WriteLogLength int
}

type ProjectConfig struct {
	Web      WebConfig
	DataBase DatabaseConfig
}

func ParseParams() (WebConfig, DatabaseConfig) {
	// web相关的配置
	env := flag.String("env", "local", "执行环境：Develop(dev)、Test(test)、Product(prod)")
	username := flag.String("username", "admin", "用户名")
	password := flag.String("password", "admin", "用户密码")
	interval := flag.Int("interval", 30, "刷新监控列表频率(秒)")
	var domain string
	flag.StringVar(&domain, "domain", "http://127.0.0.1:8080/", "服务器地址")

	// 数据库相关的配置
	dbuser := flag.String("dbuser", "admin", "数据库用户")
	dbpwd := flag.String("dbpwd", "admin", "数据库用户密码")
	dbhost := flag.String("dbhost", "127.0.0.1", "数据库Host")
	port := flag.Int("port", 8086, "数据库端口号")
	db := flag.String("db", "monitor", "数据库名")
	logLength := flag.Int("loglength", 50, "一次写入多少条日志")

	flag.Parse()
	webConfig := WebConfig{
		Env:      *env,
		Domain:   domain,
		UserName: *username,
		PassWord: *password,
		Interval: *interval,
	}
	webConfig.checkDomain()
	webConfig.LoginUrl = webConfig.GetLoginUrl()
	webConfig.MonitorListUrl = webConfig.GetMonitorListUrl()
	webConfig.EventCreateUrl = fmt.Sprintf("%sapi/1.0/monitor/event/create", webConfig.Domain)
	webConfig.EventAutoFixUrl = fmt.Sprintf("%sapi/1.0/monitor/autofix", webConfig.Domain)

	dbConfig := DatabaseConfig{*dbhost, *port, *db, *dbuser, *dbpwd, *logLength}
	return webConfig, dbConfig
}

// 解析项目相关配置：Web和DataBase
var webConfig, dbConfig = ParseParams()
var Config = &ProjectConfig{
	Web:      webConfig,
	DataBase: dbConfig,
}

// 所以项目在启动的时候，需要传递参数：
