package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL 			string
	Height,Width 	int
}
type img struct {
	Width,Height int
	Title 		 string
	Thumbnail	 thumbnail
	Animated	 bool
	IDs			 []int
}

func main()  {
	fmt.Println("encoding/json package, unmarshalling")
	var data img
	//refers to json04.gohtml template
	rcd:=`{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://7xlx3k.com1.z0.glb.clouddn.com/golang.jpg","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`
	//unmarshalling
	err:=json.Unmarshal([]byte(rcd),&data)
	if err!=nil{
		log.Fatalln("error unmarshalling",err)
	}
	fmt.Println(data)

	for i, v:=range data.IDs{
		fmt.Println(i,v)
	}

	fmt.Println(data.Thumbnail.URL)
}