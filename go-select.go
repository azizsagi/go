package main

import (
	"fmt"
	"time"
)

func main() {

	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		for {
			c1 <- "Message 1 500 MS"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {
		for {
			c2 <- "Message 2  2 Seconds"
			time.Sleep(time.Second * 2)
		}

	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}

	}

}
