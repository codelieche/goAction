package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
制作镜像
simpleweb:go-v1
simpleweb:v1
*/

// 定义全局变量
// 系统开始的时间
var startTime = time.Now() // 系统启动的时间
var host string            // web监听的地址
var port int               // web启动的端口号
var duration int           // 监控检查通过的时间
var version int = 1        // 本程序的版本v1  改成用程序控制
//var version int = 2        // 本程序的版本v2
//var version int = 3        // 本程序的版本v3

// 首页处理器
func handleIndex(w http.ResponseWriter, r *http.Request) {
	// 获取本机的主机名和第一个ip地址
	// 获取主机名
	var hostName string
	var err error
	if hostName, err = os.Hostname(); err != nil {
		log.Print("获取主机名出错：", err)
		hostName = "Error Host Name"
	}

	// 获取网卡地址
	var ips []string
	if addrs, err := net.InterfaceAddrs(); err != nil {
		log.Println(err.Error())
		http.Error(w, "获取网卡的地址出错", 500)
		return
	} else {
		//log.Println(addrs)
		for _, address := range addrs {
			//log.Println(address)
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip := ipnet.IP.String()
					//log.Println(ip)
					ips = append(ips, ip)
				}
			}
		}
	}

	content := fmt.Sprintf("Host:%s | IP:%s | Version:%d\n", hostName, strings.Join(ips, ","), version)
	w.Write([]byte(content))
	return
}

// api页面处理器
func handleApi(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf("Api Page:%s | Version:%d\n", r.URL, version)
	w.Write([]byte(content))
	return
}

// 监控检查处理器
func handleHealth(w http.ResponseWriter, r *http.Request) {
	var now time.Time
	var durationTime time.Duration

	// 如果是Delete方法就重置开始时间，这里是为了测试后续修改系统的health状态
	if r.Method == "DELETE" {
		startTime = time.Now()
		w.WriteHeader(204)
		return
	}

	now = time.Now()
	durationTime = now.Sub(startTime)
	//log.Println(duration)

	// 默认是30秒才准备好
	if durationTime <= time.Duration(duration)*time.Second {
		content := fmt.Sprintf("Not Reading!(%s) | Version:%d", durationTime, version)
		http.Error(w, content, 500)
	} else {
		content := fmt.Sprintf("Is OK!(%s) | Version:%d\n", durationTime, version)
		w.Write([]byte(content))
		return
	}

}

// 显示请求的Headers
func handleSHowHeaders(w http.ResponseWriter, r *http.Request) {
	if data, err := json.Marshal(r.Header); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}
}

func parseConfig() (string, int, int, int) {
	var host = flag.String("host", "", "监听的地址")
	var port = flag.Int("port", 8080, "端口号")
	var duration = flag.Int("duration", 30, "监控检查通过的时间")
	var version = flag.Int("version", 1, "程序版本")
	flag.Parse()
	return *host, *port, *duration, *version
}

// request: 请求页面
func handleRequestApi(w http.ResponseWriter, r *http.Request) {
	// 判断是否有next
	query := r.URL.Query()
	next := query.Get("next")

	nextRequestBody := ""
	statusCode := 200
	// 判断next的前缀
	if strings.HasPrefix(next, "/") {
		next = fmt.Sprintf("http://%s:%d%s", host, port, next)
	}

	// 需要是http开头
	if !strings.HasPrefix(next, "http") {
		if next != "" {
			nextRequestBody = fmt.Sprintf("Next: %s", next)
		}
	}

	// 发起http/https请求
	if next != "" && strings.HasPrefix(next, "http") {
		// log.Printf("有next的值：%s\n", next)
		// 发起请求
		if req, err := http.NewRequest("GET", next, nil); err != nil {
			nextRequestBody = err.Error()
			statusCode = 500
		} else {
			// 发起请求
			client := &http.Client{}
			if response, err := client.Do(req); err != nil {
				nextRequestBody = err.Error()
				statusCode = 500
			} else {
				defer response.Body.Close()

				// 响应码: 以大的为准，比如开始是200，现在是404，那么设置为404
				// 如果开始是404，后面出现了500，那么设置为500
				if response.StatusCode > statusCode {
					statusCode = response.StatusCode
				}

				// 读取body的内容
				if body, err := ioutil.ReadAll(response.Body); err != nil {
					nextRequestBody = err.Error()
					statusCode = 500
				} else {
					// 设置nextRequestBody
					nextRequestBody = fmt.Sprintf("\nNext: %s\nStatusCode：%d\nBody: %s\n", next, response.StatusCode, body)
				}
			}
		}

	}

	// 兼容：/api /request的请求，设置个区别
	prefix := "Api"
	if strings.HasPrefix(r.URL.Path, "/request") {
		prefix = "Request"
	}
	// 响应的主体内容
	content := fmt.Sprintf("%s Page:%s | Version:%d\n%s", prefix, r.URL, version, nextRequestBody)

	// 写入响应状态码和响应内容
	w.WriteHeader(statusCode)
	w.Write([]byte(content))
	return
}

func init() {
	host, port, duration, version = parseConfig()
}

func webRoute(w http.ResponseWriter, r *http.Request) {
	// 打印出客户端访问的的URL地址
	var remoteAddr = r.RemoteAddr
	msg := fmt.Sprintf("%s\t %s\t %s", r.URL, remoteAddr, r.Header.Get("User-Agent"))
	log.Println(msg)
	switch {
	case r.URL.Path == "/":
		handleIndex(w, r)
	// 监控检查
	case r.URL.Path == "/health" || r.URL.Path == "/health/":
		handleHealth(w, r)
	//	以/api开头的地址都用handleApi来处理
	case strings.HasPrefix(r.URL.Path, "/api"):
		// handleApi(w, r)
		handleRequestApi(w, r)
	case r.URL.Path == "/header" || r.URL.Path == "/header/" || r.URL.Path == "/headers" || r.URL.Path == "/headers/":
		handleSHowHeaders(w, r)
	//	以/rrequest开头的地址
	case strings.HasPrefix(r.URL.Path, "/request"):
		handleRequestApi(w, r)
	default:
		http.Error(w, "Page Not Fount", 404)
		return
	}
}

func main() {
	// 启动服务示例
	// go run main.go --port 9000 --host 0.0.0.0 --duration 10
	log.Println("启动Web Server")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
	// 获取传递的参数：host, port, duration
	log.Printf("Host: %s\tPort:%d\tDuration:%d\n", host, port, duration)
	addr := fmt.Sprintf("%s:%d", host, port)

	http.HandleFunc("/", webRoute)

	//	静态文件
	fs := http.FileServer(http.Dir("/data"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(addr, err.Error())
	} else {
		msg := "Web Server执行完毕"
		log.Println(msg)
	}
}
