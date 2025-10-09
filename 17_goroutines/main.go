package main

import "fmt"

func task(id int) {
	fmt.Println("Doing Task", id)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i)
	}
}
