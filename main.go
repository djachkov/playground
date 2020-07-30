package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
)

func hello(w http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/index.html"))
	tmpl.Execute(w, "")
}
func world(w http.ResponseWriter, request *http.Request) {
	dumpedRequest, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(dumpedRequest))
}
func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/admin", world)
	http.ListenAndServe(":8090", nil)
}
