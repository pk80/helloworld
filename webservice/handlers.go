package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"Welcom to Index page!")
}
func ToDoIndex(w http.ResponseWriter,r *http.Request)  {
	todos:=ToDos{
		ToDo{Name: "Praveen"},
		ToDo{Name: "Vasu"},
	}
	if err:=json.NewEncoder(w).Encode(todos); err!=nil{
		panic(err)
	}
}
func ToDoShow(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	todoId:=vars["todoId"]
	fmt.Fprintln(w,"Todo Show : ",todoId)
}