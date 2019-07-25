package main

import "fmt"

type people struct {
	fname string
	lname string
}

func main() {
	//fmt.Println("Hello world !")
	p1 := people{"Praveen", "Kuvvarapu"}
	p2 := people{"Vasu", "Muly"}
	fmt.Println(p1.fname + " and " + p2.fname + " are friends.")
	fmt.Println(p2.fname + " trains " + p2.fname + " learn go.")
	fmt.Println(p2.fname + " says hi to Mr. " + p1.fname)
}
