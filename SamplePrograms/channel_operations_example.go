package main

import (
	"fmt"
	"time"
)

func channel_op() {
	req := make(chan struct{}, 1)
	data := make(chan int, 1)
	block := make(chan struct{}, 1)
	go consumeData(data, block)
	go processData(req, data, block)
	time.Sleep(360 * time.Second)

}

func processData(req chan struct{}, data chan int, block chan struct{}) {
	fmt.Println("process data")
	fmt.Println("Length data", len(data))
	fmt.Println("process data next")
	for {
		if len(data) > 0 {
			select {
			case req <- struct{}{}:
				sendData(block, data, req)
			default:
			}
		}
		fmt.Println("End of IF")
	}

}

func sendData(block chan struct{}, data chan int, req chan struct{}) {
	fmt.Println("sendata")
	block <- struct{}{}
	for i := 0; i < 100; i++ {
		fmt.Println("consuming: ", <-data)
	}
	<-block
	<-req

}

func consumeData(data chan int, block chan struct{}) {
	fmt.Println("consume data")
	block <- struct{}{}
	fmt.Println("write to block48")
	for i := 0; i < 1; i++ {

		fmt.Println("write to data 50", i)
		fmt.Println("Length data  53", len(data))
		data <- i
		fmt.Println("Length data  55", len(data))
	}
	<-block
	fmt.Println("consumed data completely")

}
