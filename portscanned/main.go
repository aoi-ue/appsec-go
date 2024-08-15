package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sending request and scanning!!!")
	go func() {
	 	HTTPServer()
	}()
	Scanning()
}
