package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(h http.ResponseWriter, r *http.Request) {
	http.ServeFile(h, r, "./webserver/index.html")
}
