package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"Welcom to Index page!")
}
func ToDoIndex(w http.ResponseWriter,r *http.Request)  {
	/*todos:=ToDos{
		ToDo{Name: "Praveen"},
		ToDo{Name: "Vasu"},
	}*///added in repo.go as var
	//adding responsibilities
	w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err:=json.NewEncoder(w).Encode(todos); err!=nil{
		panic(err)
	}
}
func ToDoShow(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	todoId:=vars["todoId"]
	fmt.Fprintln(w,"Todo Show : ",todoId)
}
func ToDoCreate(w http.ResponseWriter,r *http.Request)  {
	var todo ToDo
	body, err:=ioutil.ReadAll(io.LimitReader(r.Body,1048576))
	//Notice that we use io.LimitReader. This is a good way to protect against malicious attacks on your server.
	if err!=nil{
		panic(err)
	}
	if err:=r.Body.Close();err!=nil{
		panic(err)
	}
	if err:=json.Unmarshal(body,&todo);err!=nil{
		/*After we have read the body, we then Unmarshal it to our Todo struct. If that fails,
		we will do the right thing and not only respond with the appropriate status code, 422,
		but we will also send back the error in a json string. This will allow the client to
		understand not only that something went wrong, but we have the ability to communicate
		specifically what went wrong.*/
		w.Header().Set("Content-Type","application/json;charset=UTF-8")
		w.WriteHeader(422)//unprocessable entity
		if err:=json.NewEncoder(w).Encode(err);err!=nil{
			panic(err)
		}
	}
	t:=RepoCreateToDo(todo)
	/*Finally, if all has gone well, we send back the status code of 201,
	which means that the entity was successfully created. We also send back the
	json representation of entity we created, as it contains an id that the
	client will likely need for their next step.*/
	w.Header().Set("Content-Type","applicatoin/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)//201
	if err:=json.NewEncoder(w).Encode(t);err!=nil{
		panic(err)
	}
}