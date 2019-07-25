package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type person struct {
	Fname string
	Lname string
	Items []string//slice of strings
}

func main()  {
	http.HandleFunc("/",Jason1)
	http.HandleFunc("/mshl",mshl)
	http.HandleFunc("/encd",encd)
	http.ListenAndServe(":8080",nil)
}
func Jason1(w http.ResponseWriter,r *http.Request)  {
	s:=`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Foo</title>
		</head>
		<body>
			<h4>You are at Foo page</h4>
		</body>
		</html>`
	w.Write([]byte(s))
}
func mshl(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	p:=person{
		Fname: "Praveen",
		Lname: "Kuvvarapu",
		Items:[]string{"Research","Testing","Development"},
	}
	j,err:=json.Marshal(p)
	if err!=nil{
		log.Fatalln(err)
	}
	w.Write(j)
}
func encd(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	p:=person{
		Fname: "Praveen",
		Lname: "Kuvvarapu",
		Items:[]string{"Research","Testing","Development"},
	}
	err:=json.NewEncoder(w).Encode(p)
	if err!=nil{
		log.Fatalln(err)
	}
}
