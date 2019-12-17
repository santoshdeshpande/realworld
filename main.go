package main

import "github.com/santoshdeshpande/realworld/pkg/api"

func main() {

	srv, _ := api.NewServer(3000)
	srv.Start()

	// fmt.Println("Hello World")
}
