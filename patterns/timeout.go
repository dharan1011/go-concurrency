package patterns

import (
	"fmt"
	"time"
)

func timeout(input <-chan string) {
	/*
		Time out set for each message.
		wait for 1 second to receive message from input channel
		before the timeout case is called

	*/
	for {
		select { // the whole select statement is evaluated before
		// moving to a blocked state
		case val := <-input:
			fmt.Println("Message", val)
		case <-time.After(time.Second * 1):
			fmt.Println("operation timeout")
			break
		}
	}
}

func OverallTimeout(input <-chan string) {
	/*
		The overall for loop is timed out after the set time out
	*/
	timeout := time.After(time.Minute * 1)
	for {
		select {
		case msg := <-input:
			fmt.Println("Message", msg)
		case <-timeout:
			fmt.Println("Loop timed out")
		}
	}
}
