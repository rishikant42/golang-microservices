package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	c <- "1"
	fmt.Println("1")

	c <- "2"
	fmt.Println("2")

	c <- "3"
	fmt.Println("3")

	c <- "4"
	fmt.Println("4")

	// go func(input chan string) {
	// 	fmt.Println("1")
	// 	input <- "hello-1"

	// 	fmt.Println("2")
	// 	input <- "hello-2"

	// 	fmt.Println("3")
	// 	input <- "hello-3"
	// }(c)

	// for greeting := range c {
	// 	fmt.Println(greeting)
	// }
	// greeting := <-c
	// fmt.Println(greeting)
	// greeting2 := <-c
	// fmt.Println(greeting2)
	// greeting3 := <-c
	// fmt.Println(greeting3)

	// go helloWorld()
	// time.Sleep(1 * time.Millisecond)
}

func helloWorld() {
	fmt.Println("Hello world")
}
