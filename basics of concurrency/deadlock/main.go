package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 1)
	c <- "hello" //send will block until something is ready to recieve, therefore after this line
	// no more statements will be executed and program will end in a deadlock, to make this work
	// we need to recieve this channel in a seperate go routine or make a buffered channel(won't block util the
	// channel is full)
	var msg string
	fmt.Println(msg)
}
