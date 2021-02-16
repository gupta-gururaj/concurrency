package main

import "fmt"

//import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		fmt.Println("Hello")
	}()
	//<-messages
	/* msg := <-messages //recieve
	fmt.Println(msg) */

	/* messages := make(chan string, 2)

	   messages <- "buffered"
	   messages <- "channel"

	   fmt.Println(<-messages)
	   fmt.Println(<-messages) */
}
