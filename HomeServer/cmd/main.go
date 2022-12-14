package main

import (
	server "HomeServer/pkg"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello world\n")
	host := os.Args[1]

	server.RunServer(host)
}
