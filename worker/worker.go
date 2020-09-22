package worker

import (
	"fmt"
	"net"
)

// Worker does the heavy lifting for port scanner
func Worker(ports chan int, results chan int, IP string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", IP, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
