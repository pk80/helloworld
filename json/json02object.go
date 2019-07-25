package main

import (
	"text/template"
	"net/http"
)

var tpl *template.Template

func init()  {
	tpl=template.Must(template.ParseGlob("json02.gohtml"))
}
type client int
func (c client) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	tpl.ExecuteTemplate(w,"json02.gohtml",r.Form)
}
func main()  {
	var c1 client
	http.ListenAndServe(":8080",c1)
}