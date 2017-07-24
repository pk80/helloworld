package main

import (
	"fmt"
)

type people struct {
	fname string
	lname string
}

func main()  {
	//fmt.Println("Hello world !")
	p1:=people{"Praveen","Kumar"}
	fmt.Println(p1)
}
