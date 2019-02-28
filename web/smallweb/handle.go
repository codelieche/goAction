package smallweb

import (
	"encoding/json"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Index"))
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
