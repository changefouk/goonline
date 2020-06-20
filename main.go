package main

func main() {
	r := todo.SetupRouter()
	r.Run(":1234")
}
