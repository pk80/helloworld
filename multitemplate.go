package main

import (
	"net/http"
	"text/template"
)
var tpl *template.Template

func init()  {
	tpl=template.Must(template.ParseGlob("templates/*.gohtml"))
}
type client int

func (c client) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/index":
		tpl.ExecuteTemplate(w,"index.gohtml",req.URL.Path)
	case "/login":
		tpl.ExecuteTemplate(w,"login.gohtml",req.URL.Path)
	case "/contact":
		tpl.ExecuteTemplate(w,"contact.gohtml",req.URL.Path)
	case "/about":
		tpl.ExecuteTemplate(w,"about.gohtml",req.URL.Path)
	}
}

func main() {
	var c1 client
	http.ListenAndServe(":8080", c1)
}