package main

import "fmt"

const age = 25

func main() {
	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(host, port)
}