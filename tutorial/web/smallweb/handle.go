package smallweb

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	var ips = []string{}
	var hostName string
	var err error

	if hostName, err = os.Hostname(); err != nil {
		http.Error(w, "获取主机名出错", 500)
		return
	}

	if addrs, err := net.InterfaceAddrs(); err != nil {
		log.Println(err.Error())
		http.Error(w, "获取网卡的地址出错", 500)
		return
	} else {
		//log.Println(addrs)
		//log.Println(len(addrs))

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

		// 返回结果
		response := IndexResponse{
			HostName: hostName,
			IP:       ips,
			Datetime: time.Now(),
		}
		if data, err := json.Marshal(response); err != nil {

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}

	//w.Write([]byte("Hello Index"))
}

func api(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	r.ParseForm()
	response := ApiResponse{
		Method:    method,
		UserAgent: r.Header.Get("User-Agent"),
		Data:      r.Form,
	}

	if data, err := json.Marshal(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
