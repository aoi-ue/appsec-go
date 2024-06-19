package http

import (
	"fmt"
	"net/http"
)

func HTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the HTTP Server!")
	})

	// go func() {
		fmt.Println("Starting HTTP server on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println(err)
		}
	// }()

	//select {}
}

