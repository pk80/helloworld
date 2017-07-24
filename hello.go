package main

import "fmt"

type people struct {
	fname string
	lname string
}

func main()  {
	//fmt.Println("Hello world !")
	p1:=people{"Praveen","Kuvvarapu"}
	p2:=people{"Vasu","Muly"}
	fmt.Println(p1.fname +" and "+p2.fname+ " are friends." )
	fmt.Println(p1.fname + " take lessons from "+p2.fname)
}
