package main

import (
	"mime"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

var status bool = false;

func toggle(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		status = !status
	}
	fmt.Fprintf(w, "status: %t", status)

}
func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("src/webserver" + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(r.URL.Path)))
	w.Write(body)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/state", toggle)
	http.ListenAndServe(":8090", nil)
}
