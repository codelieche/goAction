package webhealth

import (
	"encoding/json"
	"net/http"
	"time"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	content := CheckResponse{
		Status:  true,
		Message: "系统正在运行",
		Time:    time.Now(),
	}
	if contentData, err := json.Marshal(content); err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(contentData)
		return
	}
}
