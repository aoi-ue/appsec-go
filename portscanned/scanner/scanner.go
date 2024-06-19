package scanner

import (
	"flag"
	"fmt"
	"net"
	"time"
)

type Port struct {
	open bool
	closed bool 
	filtered bool 
}

func CLI() {
	localhost := flag.String("host", "localhost", "url")
	portPtr := flag.Int("port", 80, "port to use")

	flag.Parse()

	fmt.Printf("Scanning host: %s port: %d\n ", *localhost, *portPtr)
	fmt.Printf("Port: %d is open", 
	*portPtr)
}

func Scanning() {
	// target := "127.0.0.1"
	timeout := time.Second

	for port := 1; port <= 65535; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.DialTimeout("tcp", address, timeout)

		if err == nil {
			fmt.Printf("Port %d is open\n", port)
			conn.Close()
		}
	}
}
