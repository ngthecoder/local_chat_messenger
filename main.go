package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "server" {
		server()
	} else if len(os.Args) == 2 && os.Args[1] == "client" {
		client()
	} else {
		fmt.Println("Incorrect usage")
	}
}
