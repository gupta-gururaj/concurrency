package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true //go routine is finished so notify on channel(bool type),  response sent on channel
}

//An example of using a blocking receive to wait for a goroutine to finish.
//By default sends and receives block until both the sender and receiver are ready.

func main() {
	done := make(chan bool /*, 0 (automatically unbuffered)*/)
	go worker(done)
	<-done //Response recieved from channel, if no recieve then worker will be blocked from line 11
	fmt.Println("main function ends here . . .")
}

/*A send operation on an unbuffered channel blocks the sending goroutine until
another goroutine executes a corresponding receive on the same channel, at which
point the value is transmitted and both goroutines may continue.
Conversely, if the receive operation was attempted first, the receiving
goroutine is blocked until another goroutine performs a send on the same channel.*/
