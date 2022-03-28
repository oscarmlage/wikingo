package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oscarmlage/wikingo/server"
)

var (
	app_name    = "Wikingo"
	app_version = "0.0.1"
)

func main() {
	flag_debug := flag.Bool("d", false, "Enable debug mode")
	flag_version := flag.Bool("v", false, "Show version")
	flag.Parse()

	// Version
	if *flag_version {
		fmt.Printf("%s - v%s\n", app_name, app_version)
		os.Exit(0)
	}

	// Logger
	server.InitLogger(*flag_debug)
	server.Debug.Printf("Yo there")

	fmt.Println("Starting wikingo...")
	fmt.Printf("Debug mode: %t\n", *flag_debug)
	fmt.Printf("Version: %t\n", *flag_version)
	server.Serve()
}
