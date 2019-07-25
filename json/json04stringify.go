package main

import (
	"text/template"
	"net/http"
)

var tpl4 *template.Template

func init()  {
	tpl4=template.Must(template.ParseGlob("json04.gohtml"))
}

type client4 int

func (c client4) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	tpl4.ExecuteTemplate(w,"json04.gohtml",r.Form)
}
func main()  {
	var c4 client4
	http.ListenAndServe(":8080",c4)
}
