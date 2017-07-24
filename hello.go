package main

import (
	"fmt"
	//"bufio"
	//"os"
)

func main()  {
	fmt.Println("Hello world !")
	//reader:=bufio.NewReader(os.Stdin)
	u1:=""
	fmt.Println("Enter your name:")
	//u1, _:=reader.ReadString('\n')
	fmt.Scanln(&u1)
	//fmt.Println(u1)
	p1:=""
	fmt.Println("Enter password:")
	fmt.Scanln(&p1)
	if u1=="Praveen" && p1=="abcdef"{
		fmt.Println("login success...")
	}else {
		fmt.Println("login failed...")
	}
}
