package main

import (
	"fmt"
	"sync"
)

func main() {
	wg:=new (sync.WaitGroup)
	for i:=0; i< 10; i++ {
		// Need add to group before run goroutine
		wg.Add(1)
		name := fmt.Sprintf("Hi, %d", i)
		go customLog(name, wg)
	}
	wg.Wait()
}

func customLog(name string, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println(name)
}