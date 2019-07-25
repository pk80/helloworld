package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {
	/*
	h:=http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type","text/html;charset=utf-8")
		fmt.Fprintf(w,"Hi 'World HandlerFunc'! You have queried for : %s\n",`<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	})
	err:=http.ListenAndServe(":8080",h)
	if err!=nil{
		log.Fatalln(err)
	}
	*/

	h1:=http.NewServeMux()
	h1.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/html;charset=utf-8")
		fmt.Fprintf(w,"Hi 'World HandlerFunc'! You have queried for : %s\n",`<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	})

	err1:=http.ListenAndServe(":8080",h1)
	if err1!=nil {
		log.Fatalln(err1)
	}
}
