package main

import (
	// "appsec-go/http"
	"appsec-go/scanner"
	"fmt"
)

func main() {
	fmt.Println("Sending request and scanning!!!")
	// go func() {
	//  	http.HTTPServer()
	// }()
	scanner.Scanning()
}
