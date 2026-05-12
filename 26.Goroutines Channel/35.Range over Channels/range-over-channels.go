package main

import "fmt"

func main() {
	queue := make(chan int, 2)
	queue <- 1
	queue <- 2
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
