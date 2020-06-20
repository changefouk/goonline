package main

import "github.com/changefouk/todo"

func main() {
	r := todo.SetupRouter()
	r.Run(":1234")
}
