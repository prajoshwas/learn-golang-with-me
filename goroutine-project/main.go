package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)

	doneChan <- true
	close(doneChan)
}

func main() {

	/* dones := make([]chan bool, 4) */

	/* dones[0] = make(chan bool)
	dones[1] = make(chan bool)
	dones[2] = make(chan bool)
	dones[3] = make(chan bool) */
	done := make(chan bool)
	go greet("Greetings!", done)
	go greet("How are You?", done)
	go slowGreet("Slow Greetings!", done)
	go greet("Wow Concurrent greetings", done)
	/* <-dones[0]
	<-dones[1]
	<-dones[2]
	<-dones[3] */
	for _ = range done {
	}
}
