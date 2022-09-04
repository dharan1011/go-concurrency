package patterns

import (
	"fmt"
)

/*
Daisy chain patter

Chinese Whispers simulation
*/
func whisper(left <-chan int, right chan<- int) {
	for {
		val := <-left + 1
		fmt.Println(val)
		right <- val
	}
}

/*
start
  |
left -> right -> left -> right -> left
*/

func ChineseWhispers() {
	start := make(chan int)
	left := start
	for i := 0; i < 5; i++ {
		right := make(chan int)
		go whisper(left, right)
		left = right
	}
	start <- 0
	fmt.Println(<-left)
}
