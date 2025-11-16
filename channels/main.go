package main

import (
	"fmt"
	"sync"
	"time"
)

func simpleSendReceive() {
	fmt.Println("=== Simple Send / Receive ===")

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)
}

func bufferedChannel() {
	fmt.Println("\n=== Buffered Channel ===")

	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	fmt.Println("Buffered values sent:", 10, 20)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Reading:", <-ch)
	}()

	// This will block until goroutine reads from the channel
	ch <- 30
	fmt.Println("Sent:", 30)
}

func producerConsumer() {
	fmt.Println("\n=== Producer / Consumer ===")

	ch := make(chan int)
	var wg sync.WaitGroup

	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // consumer can finish
	}()

	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Consumed:", v)
		}
	}()

	wg.Wait()
}

func multiChannelSelect() {
	fmt.Println("\n=== Multi-Channel Select ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)

		case msg := <-ch2:
			fmt.Println("Received:", msg)
		}
	}
}

func timeoutWithSelect() {
	fmt.Println("\n=== Timeout with Select ===")

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Operation completed"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)

	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! No response.")
	}
}

func main() {
	// simpleSendReceive()
	bufferedChannel()
	// producerConsumer()
	// multiChannelSelect()
	// timeoutWithSelect()
}
