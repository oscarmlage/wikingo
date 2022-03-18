package main

import (
	"fmt"
	"github.com/oscarmlage/wikingo/server"
)

func main() {
	fmt.Println("Starting wikingo...")

	server.Serve()
}
