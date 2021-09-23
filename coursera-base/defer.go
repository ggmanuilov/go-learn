/**
* defer - ключевое слово
* выполняеются в обратном порядке объявления
* FILO - first in last out
*
* Конструкции с defer полезны для восстановления из паники.
* Смотри файл recovery.go
 */
package main

import "fmt"

func getSomeVars() string {
	fmt.Println("get some vars")
	return "[GET SOME VARS] result"
}

func deferOne() {
	defer fmt.Println("After work 3")
	defer fmt.Println("After work 2")
	defer fmt.Println(getSomeVars())
	defer fmt.Println("After work 1.1")
	defer fmt.Println("After work 1")
	fmt.Println("General work")

	/** output
	get some vars
	General work
	After work 1
	After work 1.1
	[GET SOME VARS] result
	After work 2
	After work 3
	*/
}

func deferLine() {
	fmt.Println("Для удобства выполнения мзавернем отложенную функцию в колбэк")
	defer fmt.Println("After work 3")
	defer fmt.Println("After work 2")
	defer func() {
		fmt.Println(getSomeVars())
	}()
	defer fmt.Println("After work 1.1")
	defer fmt.Println("After work 1")
	fmt.Println("General work")
	/** output
	  General work
	  After work 1
	  After work 1.1
	  get some vars
	  [GET SOME VARS] result
	  After work 2
	  After work 3
	*/
}

func main() {
	deferOne()
	deferLine()
}
