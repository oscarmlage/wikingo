package main

import (
	"flag"
	"fmt"

	"github.com/oscarmlage/wikingo/server"
)

func main() {
	fmt.Println("Starting wikingo...")
	flag_debug := flag.Bool("d", false, "Enable debug mode")
	flag.Parse()

	// Logger
	server.InitLogger(*flag_debug)
	server.Debug.Printf("Yo there")

	fmt.Println("Starting wikingo...")
	fmt.Printf("Debug mode: %t\n", *flag_debug)
	server.Serve()
}
