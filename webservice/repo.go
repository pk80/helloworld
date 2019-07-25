package main

import "fmt"

var currentID int
var todos ToDos
//Giving some seed data
func init(){
	RepoCreateToDo(ToDo{Name:"Praveen"})
	RepoCreateToDo(ToDo{Name:"Vasu"})
}

func RepoFindToDo(id int) ToDo  {
	for _,t:=range todos{
		if t.Id==id{
			return t
		}
	}
	return ToDo{}//if not found empty todo is reuturend.
}

func RepoCreateToDo(t ToDo)ToDo  {
	currentID+=1
	t.Id=currentID
	todos=append(todos,t)
	return t
}
func RepoDestroyToDo(id int)error  {
	for i,t:=range todos{
		if t.Id==id{
			todos=append(todos[:i],todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find ToDo with id of %d to delete.",id)
}