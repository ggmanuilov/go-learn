package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Producer struct {
	OutChan chan int
}

func (p *Producer) produce()  {
	for{
		time.Sleep(3 * time.Second)
		// write to channel
		p.OutChan <- rand.Int()
	}
}

func (p *Producer) getOutChan() <-chan int  {
	return p.OutChan
}

func main() {
	p := Producer{
		OutChan: make(chan int, 10),
	}
	go p.produce()
	// read channel only
	//for i:= range p.getOutChan(){
	//	fmt.Println("Message from channel ", i)
	//}

	ticker := time.NewTicker(2*time.Second)
	for{
		// non block read from multi channels
		select {
		case prodMessage:= <- p.getOutChan():
			fmt.Println("Producer message: ", prodMessage)
		case tickerMessage := <- ticker.C:
			fmt.Println("Ticker message: ", tickerMessage)
		}
	}
}