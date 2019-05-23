package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		io.WriteString(w, "udddnknow")
	} else {
		io.WriteString(w, host)
	}
}

func main() {
	http.HandleFunc("/who", handler)
	http.ListenAndServe(":80", nil)
}

