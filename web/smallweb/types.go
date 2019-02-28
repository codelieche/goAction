package smallweb

type Config struct {
	Host string // 监听的地址：默认127.0.0.1
	Port int    // 监听的端口：默认8080
}

type ApiResponse struct {
	Method    string      // 请求方法
	UserAgent string      // 请求头--UserAgent
	Data      interface{} // 请求内容
}
