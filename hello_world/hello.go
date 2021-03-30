package main

import "fmt"

const prefixHello = "Hello, "

func Hello(name string) string {
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("world"))
}
