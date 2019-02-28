package smallweb

import (
	"log"
	"net/http"
)

func webRoute(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	switch {
	case r.URL.Path == "/":
		index(w, r)
	case r.URL.Path == "/api" || r.URL.Path == "/api/":
		api(w, r)
	default:
		http.Error(w, "Page Not Fount", 404)
		return
	}
}
