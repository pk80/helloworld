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
	p1:=people{"Praveen","Kuvvarapu"}
	p2:=people{"Vasu","Muly"}
	fmt.Println(p1.fname +" and "+p2.fname+ " are friends." )
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	fmt.Println(p1.fname + " take lessons from "+p2.fname)
>>>>>>> brhello
=======
	fmt.Println(p1.fname + " take lessons from "+p2.fname)
>>>>>>> brhello
=======
	fmt.Println(p2.fname + " trains "+p2.fname+" learn go.")
>>>>>>> brhello
}
