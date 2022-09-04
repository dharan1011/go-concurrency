package patterns

import "fmt"

type Message[Type interface{}] struct {
	Data Type
	Wait chan bool
}

func DataGenerators() <-chan Message[int] {
	inputChannel := make(chan Message[int])
	wait := make(chan bool)
	// genrator pattern to generate some data
	go func() {
		for i := 0; ; i++ {
			inputChannel <- Message[int]{i, wait}
			<-wait // hold till the receiver signals
		}
	}()
	return inputChannel
}

func SequenceGenerators() {
	a := DataGenerators()
	b := DataGenerators()
	for i := 0; i < 5; i++ {
		msgA := <-a
		fmt.Println("Value from a", msgA.Data)
		msgB := <-b
		fmt.Println("Value from b", msgB.Data)
		msgA.Wait <- true
		msgB.Wait <- true
	}
}
