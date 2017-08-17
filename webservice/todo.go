package main

import "time"

type ToDo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}
//By adding struct tags you can control exactly what and how your struct will be marshalled to JSON.
type ToDos []ToDo
