package main

import (
	"os"
)

func main() {
	err := Run(os.Args[1:], os.Stdout)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
}
