package main

import "fmt"

const pi = 123

type UserId = int

func main() {
	idx := 1
	var uid UserId = 42

	//	даже при одинаковом базовом типе,
	// разные типы не совместимы
	myID := UserId(idx)

	fmt.Println("Result: ", myID)
	fmt.Println("Result: ", uid)

	a := 10
	b := &a // b указатель на а
	*b = 20 // по ссылке b = 20, значит a = не 10, а 20

	fmt.Println("Result *b 1: ", *b)

	a = 30 // присвоение а = 30, *b = 30

	fmt.Println("Result *b 2: ", a, *b)
}
