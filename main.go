package main

import (
	"net/http"
	"text/template"
)
var tpl1 *template.Template

func init()  {
	tpl1=template.Must(template.ParseGlob("templates/*.gohtml"))
}
type client1 int

func (c client1) ServeHTTP(w http.ResponseWriter, r *http.Request) {//need to workout this !!!
	s:=r.URL.Path
	tpl1.ExecuteTemplate(w,"templates"+s+".gohtml",r.URL.Path)
	//fmt.Fprintln(w,s)
}

func main() {
	var c1 client1
	http.ListenAndServe(":8080", c1)
}
