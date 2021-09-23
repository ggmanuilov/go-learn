/**
Паника - это исключительная ситуация. 
Приводит к краху программы.
Её нельзя использовать как try catch в других языках.
*/
package main

import "fmt"

func deferTest() {
	// defer func - будет выполнен всеровно после падения, что интересно
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
			// таким образом, при обработке паники
			// программ не завершится с паникой
		}
	}()
	fmt.Println("Work job...")
	panic("somthing bad happend")
	return
}

func main() {
	deferTest()
	return
}
