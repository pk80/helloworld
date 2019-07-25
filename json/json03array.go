package main

import (
	"text/template"
	"net/http"
)

var tpl3 *template.Template

func init()  {
	tpl3=template.Must(template.ParseGlob("json03.gohtml"))
}
type client3 int
func (c client3) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	tpl3.ExecuteTemplate(w,"json03.gohtml",r.Form)
}
func main()  {
	var c3 client3
	http.ListenAndServe(":8080",c3)
}