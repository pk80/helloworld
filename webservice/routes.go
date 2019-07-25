package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes =Routes{
	Route{"Index",
		"GET",
		"/",
		Index,
	},
	Route{"ToDoIndex",
		"GET",
		"/todos",
		ToDoIndex,
	},
	Route{"ToDoShow",
		"GET",
		"/todos/{todoId}",
		ToDoShow,
	},
	Route{"ToDoCreate",
		"POST",
		"/todos",
		ToDoCreate,
	},
}