package main

import (
	"fmt"
	"strings"
)

func main() {
	var response string
	var responded bool = false
	for !responded {
		fmt.Println("Would you like to host or connect? (h/c)")
		fmt.Scanf("%v", &response)
		response = strings.TrimRight(response, "\n")
		if response == "h" {
			responded = true
			createServer()
		} else if response == "c" {
			responded = true
			join()
		} else {
			fmt.Println("Did not specify connection!")
		}
	}
}

func createServer() {
	fmt.Println("created server")
}

func join() {
	fmt.Println("joined server")
}
