package main

import (
	"fmt"
)

type client struct {
	uname string
	pwd   string
}
type clients struct {
	aut []client
	cmp string
}

type comps struct {
	cmps clients
	area string
}

func pt(a comps) {
	a1 := ""
	fmt.Println("**", a.cmps.cmp, "is located in", a.area, "**")
	for i := 0; i < 2; i++ {
		if a1 != "" {
			a1 = a1 + "," + a.cmps.aut[i].uname
		} else {
			a1 = a.cmps.aut[i].uname
		}
	}
	fmt.Println(a1, "are in", a.cmps.cmp)
	for i := 0; i < 2; i++ {
		fmt.Println(a.cmps.aut[i].uname, "uses", a.cmps.aut[i].pwd, "as password")
	}
}

func main() {
	c1 := client{"Praveen", "p2345"}
	c2 := client{"Vasu", "v7890"}
	c3 := client{"Pramod", "pv135"}
	cabc := clients{[]client{c1, c2}, "ACB Company"}
	c123 := clients{[]client{c1, c3}, "123 Company"}
	comp1 := comps{cabc, "US"}
	comp2 := comps{c123, "India"}

	pt(comp1)
	pt(comp2)
}