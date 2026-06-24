package network

import "fmt"

func Ping(host string) bool {
	fmt.Println("pinging", host)
	return true
}
