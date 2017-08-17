package main

import (
	"time"
	"net/http"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"log"
)

//A basic web server :
/*
A RESTful service starts with fundamentally being a web service first.
Here is a really basic web server that responds to any requests by simply outputting the request url
*/
/*func main()  {//net/http,fmt,html & log packages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))
	})
	//http.HandleFunc("/about",About)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatalln(err)
	}
}*/
//Adding a router : third party mux routers from Gorilla Web toolkit, Julien Schmidt
/*
The below example creates a basic router, adds the route / and assigns the Index handler to run when that endpoint is called.
You will also notice now that before we could ask for http://localhost:8080/... and that worked.
That will no longer work now as there is no route defined. Only http://localhost:8080 will be a valid response.
*/
/*func main()  {//net/http,fmt,html,log packages & github.com/gorilla/mux is retrieved by go get statement from gitbash/cmd
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)//defining route is here
	router.HandleFunc("/hi",Index)//2nd route
	router.HandleFunc("/about",About)//3rd route need About func to handle
	log.Fatalln(http.ListenAndServe(":8080",router))
}
func Index(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome to Index page, %q",html.EscapeString(r.URL.Path))
}
func About(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome to About page, %q",r.URL.Path)
}*/
//Creating basic routes :
/*
We have now added two more endpoints (or routes)- todo, todo/{todoId}
This is the beginning of a RESTful design.
Pay close attention to the last route where we added a variable in the route, called todoId: http://localhost:8080/todos/{todoId}
This will allow us to pass in id’s to the route and respond with the proper records.
 */
/*func main()  {
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/about",About)
	router.HandleFunc("/todo",ToDoIndex)
	router.HandleFunc("/todo/{todoId}",ToDoShow)
	log.Fatalln(http.ListenAndServe(":8080",router))
}
func ToDoIndex(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome to ToDo Index page, %q",r.URL.Path)
}
func ToDoShow(w http.ResponseWriter,r *http.Request)  {
	vars:=mux.Vars(r)
	todoId:=vars["todoId"]
	fmt.Fprintln(w,"Welcome to ToDo Show page.\nTodo Show: ",todoId)
}*/
//A Basic Model :
/*
Now that we have routes in place, it’s time to create a basic Todo model that we can send and retrieve data with.
In Go, a struct will typically serve as your model. Many other languages use classes for this purpose.
 */
/*type ToDo struct {
	Name string
	Completed bool
	Due time.Time
}
type ToDos []ToDo //a slice (an ordered collection) of Todo
//Send Back some JSON : Now that we have a basic model, we can simulate a real response and mock out the TodoIndex with static data.
func ToDoIndex(w http.ResponseWriter,r *http.Request)  {
	//w.Header().Set("Content-Type","application/json")//this gives values of json application
	todos:=ToDos{
		ToDo{Name:"Praveen"},
		ToDo{Name:"Vasu"},
	}
	json.NewEncoder(w).Encode(todos)

}//For now, we are just creating a static slice of Todos to send back to the client. Now if you requesthttp://localhost:8080/todos
func main()  {
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos",ToDoIndex)
	log.Fatalln(http.ListenAndServe(":8080",router))
}*/
//A Better Model :
type ToDo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}
//By adding struct tags you can control exactly what and how your struct will be marshalled to JSON.
type ToDos []ToDo