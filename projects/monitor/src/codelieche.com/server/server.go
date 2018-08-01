package server

import (
	"codelieche.com/settings"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (server *MonitorServer) start() {
	http.HandleFunc("/info", func(write http.ResponseWriter, request *http.Request) {
		server.process.Info.RunTime = time.Now().Sub(server.startTime).String()
		server.process.Info.TaskChanLen = len(server.process.TaskChan)
		server.process.Info.LogChanLen = len(server.process.LogChan)

		if result, err := json.MarshalIndent(server.process.Info, "", "\t"); err != nil {
			io.WriteString(write, err.Error())

		} else {
			io.WriteString(write, string(result))
		}
	})

	// 启动web服务
	addr := fmt.Sprintf(":%d", settings.Config.Web.ServerPort)
	http.ListenAndServe(addr, nil)
}
