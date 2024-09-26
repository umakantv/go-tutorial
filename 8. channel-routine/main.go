package main

import (
	"fmt"
	"time"
)

func main() {

	basics()

	// bufferedChannels()

	// asynchronousTasks()

	// synchronousTasks()

	// rangingOverChannels()

	// selectStatement()
	selectStatementWithTwoWayCommunication()

	// checkSites()
	// go MultipleChannels()

}

// Define a task that takes some time to complete
// Ideal to be run in a goroutine
func task(delay int, c chan int64) {
	time.Sleep(time.Duration(delay) * time.Second)
	c <- time.Now().Unix()
}

// Basic usage of channels
// By default, sends and receives block until the other side is ready.
func basics() {

	c := make(chan int64)

	// Start the task, whenever done, it will send the result to the channel
	go task(1, c)

	// Wait til we receive from the channel
	fmt.Println(<-c)
}

// Buffered channels
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
func bufferedChannels() {
	c := make(chan int64, 2)

	c <- 1
	c <- 2

	// This would fail if the channel was not buffered
	// First we have to receive from the channel to make space for this
	// If we send more than the buffer length, it will block
	// c <- 3
	// Sending to a buffered channel is blocking only when the buffer is full.
	// Receiving is blocking when the buffer is empty.

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// Fetching data from multiple sources at once without blocking
// This is similar to await Promise.all(promise1, promise2) in JS
func asynchronousTasks() {
	// c := Fetch(4)
	c := GenericFetch(func(a chan int64) {
		task(4, a)
	})
	d := GenericFetch(func(a chan int64) {
		task(3, a)
	})
	// d := Fetch(3)
	fmt.Println(<-c, <-d)
}

// Fetching data from multiple sources at once with blocking
// Use this when there is a dependency between the tasks
// This is similar to await promise1, await promise2 in JS
func synchronousTasks() {
	c := <-Fetch(4)
	fmt.Println(c)
	d := <-Fetch(3)
	fmt.Println(d)
}

// We can range over channels to get results as they arrive and process them
// same as promise1.then(), promise2.then() in JS
func rangingOverChannels() {

	c := make(chan int64)

	// Start multiple tasks
	for i := 0; i < 5; i++ {
		go task(1, c)
	}

	go func() {
		// Wait for all tasks to complete
		time.Sleep(3 * time.Second)
		// Note: Channel should be closed by the sender, not the receiver
		close(c)
	}()

	// To terminal the range loop with channels, it must be closed
	for t := range c {
		fmt.Println(t)
	}
}

// A select blocks until one of its cases (both receiving and sending) can run, then it executes that case.
// It chooses one at random if multiple are ready.
// This will process the first result that arrives and ignore the rest
func selectStatement() {
	c := make(chan int64)
	d := make(chan int64)

	go task(4, c)
	go task(3, d)

	select {
	case res := <-c:
		fmt.Println("Select got results from first task", res)
	case res := <-d:
		fmt.Println("Select got results from second task", res)
	}

	e := make(chan int64)
	f := make(chan int64)
	// This quits after processing one result
	go task(4, e)
	go task(3, f)

	// To process all results, use a loop
	// we could also use empty for loop with additional case to exit the loop
	// NOTE: use return to exit the loop, break will only exit the select statement
	for i := 0; i < 2; i++ {
		select {
		case res := <-e:
			fmt.Println("Select got results from 3rd task", res)
		case res := <-f:
			fmt.Println("Select got results from 4th task", res)
		}
	}
}

func selectStatementWithTwoWayCommunication() {

	// select statement can also be used to block until a channel is ready to receive

	c := make(chan int64)
	quit := make(chan int)

	go func() {
		// fmt.Println("Waiting for 2 seconds before receiving")
		// time.Sleep(2 * time.Second)
		fmt.Println(<-c, "Received from channel")
		quit <- 0
		// Close the channel once sent
		close(quit)
	}()

	for {
		select {
		case c <- time.Now().Unix():
			fmt.Println("Sent to channel")
			// close(g)
		case <-quit:
			// This will allow us to wait until we receive from the channel
			fmt.Println("Receiver sent signal, done receiving")
			return
		}
	}
}

func Fetch(delay int) chan int64 {
	c := make(chan int64)
	go func() {
		task(delay, c)
	}()
	return c
}

func GenericFetch[T any](fn func(c chan T)) chan T {
	c := make(chan T)
	go func() {
		fn(c)
	}()
	return c
}
