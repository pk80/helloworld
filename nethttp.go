package main

import (
	"fmt"
	"log"
	"net/http"
	//"reflect"
	//"crypto/aes"
)

/*//Handler
type helloHandler struct {

}
func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello World 'Handler'! You have queried for : %s\n",r.URL.Path)
}

func main()  {
	err:=http.ListenAndServe(":8080",helloHandler{})
	if err!=nil{
		log.Fatalln(err)
	}
}*/
/*//HandlerFunc
func main()  {
	h:=http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Hi 'World HandlerFunc'! You have queried for : %s\n",r.URL.Path)
	})
	err:=http.ListenAndServe(":8080",h)
	if err!=nil{
		log.Fatalln(err)
	}
}*/
/*//ServeMux
func main()  {
	h:=http.NewServeMux()
	h.HandleFunc("/hi world sm1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Hi World 'ServeMux'! You have queried for sm1.")
	})
	h.HandleFunc("/hi world sm2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Hi World 'ServeMux'! You have queried for sm2.")
	})

	err:=http.ListenAndServe(":8080",h)
	if err!=nil {
		log.Fatalln(err)
	}
}*/
/*//Composability
type numberDumper int

func (n numberDumper)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Here is what you entered : %d\n",n)
}
func main()  {
	h:=http.NewServeMux()
	h.Handle("/one",numberDumper(1))
	h.Handle("/nine",numberDumper(9))

	err:=http.ListenAndServe(":8080",h)
	if err!=nil{
		log.Fatalln(err)
	}
}*/
//Testing package ."testing" ****TODO****
/*//Middleware
type numberDumper int

func (n numberDumper)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Here's your number: %d\n",n)
}
func logger(h http.Handler)http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		log.Printf("%s requested %s",r.RemoteAddr,r.URL)
		h.ServeHTTP(w,r)
	})
}
func main()  {
	h:=http.NewServeMux()
	h.Handle("/one",numberDumper(1))
	h.Handle("/nine",numberDumper(9))
	h1:=logger(h)
	err:=http.ListenAndServe(":8080",h1)
	if err!=nil{
		log.Fatalln(err)
	}
}*/
/*//Middleware Chaining
type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Here's your number : %d\n",n)
}
func logger(h http.Handler)http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		log.Printf("%s requested %s",r.RemoteAddr,r.URL)
		h.ServeHTTP(w,r)
	})
}
type headerSetter struct {
		key, val string
		handler  http.Handler
	}

func (hs headerSetter)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set(hs.key,hs.val)
	hs.handler.ServeHTTP(w,r)
}
func newHeaderSetter(key,val string) func(http.Handler) http.Handler  {
	//func newHeaderSetter(key, val string, h http.Handler) http.Handler
	return func(h http.Handler) http.Handler{
		return headerSetter{key,val,h}
	}
}
func main(){
	h:=http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/nine",numberDumper(9))

	h.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(404)
		fmt.Fprintln(w,"\nThat's not a supported number!")
	})
	h1:=logger(h)
	hhs:=newHeaderSetter("X-FOO","BAR")(h1)

	err:=http.ListenAndServe(":8080",hhs)
	if err!=nil{
		log.Fatalln(err)
	}
}*/
//Composing middleware with alice
type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

type headerSetter struct {
	key, val string
	handler  http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}
func newHeaderSetter(key, val string) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return headerSetter{key, val, h}
	}
}
func main() {
	h := http.NewServeMux()
	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))
	h.Handle("/six", numberDumper(6))
	h.Handle("/seven", numberDumper(7))
	h.Handle("/eight", numberDumper(8))
	h.Handle("/nine", numberDumper(9))

	tenHS := newHeaderSetter("X-Five", "First two digit number")
	h.Handle("/ten", tenHS(numberDumper(10)))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	//	chain:=alice.New(//alice is in new pakage in middleware
	//		newHeaderSetter("X-Foo","Bar"),
	//		newHeaderSetter("X-Baz","Buz"),
	//		logger,
	//	).Then (h)
	err := http.ListenAndServe(":8080", h)
	log.Fatal(err)
}
