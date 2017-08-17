package main

import (
	"strings"
	"fmt"
	"time"
	"net/http"
	"io"
)
//create a simple function to return 3 characters
func firstThree(s string)string  {//strings package
	s=strings.TrimSpace(s)
	if len(s)>=3{
		s=s[:3]//slice of 3 characters from start
	}
	return s
}
//create a function to display date as string
func monthDayYear(t time.Time)string  {//time package
	// 01/02 03:04:05PM '06 -0700
	return t.Format("02-01-2006")//02-Jan-2006
}
//Handle function
func d(response http.ResponseWriter,request *http.Request)  {//net/http package
	io.WriteString(response,"dog dog dog")

}
func main()  {
	//string function
	s1:="Buddha"
	fmt.Println(firstThree(s1))
	//date time function
	fmt.Println(time.Now())//2017-08-11 08:44:31.8226254 +0530 IST
	fmt.Println(monthDayYear(time.Now()))
	/*
	//Handle function
	http.HandleFunc("/dog",d)
	http.ListenAndServe(":8080",nil)//killing this process tree ???
	*/
}