package smallweb

import "flag"

var config Config

func parseConfig() {
	var host = flag.String("host", "127.0.0.1", "监听的地址")
	var port = flag.Int("port", 8080, "端口号")
	flag.Parse()
	config.Host = *host
	config.Port = *port
}
