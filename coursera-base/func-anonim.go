package main

import "fmt"

func main() {
	// ананимная функция
	func(in string) {
		fmt.Println(in)
	}("nobody")

	// присвоение анонимной функции в переменную
	doAny := func(in string) {
		fmt.Println("do any:", in)
	}

	doAny("do anythink")

	// определяем тип функции
	// чтобы мочь передовать ее как аргумент в другие функции
	type strFuncType func(string, string) string
	type strFuncTypeInt func(string, string) (int, error)

	/**
	 * пример 2 передача анонимной функции
	 */
	fmt.Println("пример 2 передача анонимной функции")

	type strType func(string)

	logWrite := func(in string) {
		fmt.Println("do any:", in)
	}

	worker := func(callback strType) {
		callback("call as callback")
	}
	worker(logWrite)

	/**
	 * Пример 3 - замыкания
	 */
	 fmt.Println("Пример 3 - замыкания")
	 logPrfixer := func(prefix string) strType{
		 return func(in string) {
			 fmt.Printf("[%s] %s \n", prefix, in)
		 }
	 }
	 successLogger := logPrfixer("CRON")
	 successLogger("Exchange begin.")
	 successLogger("Exchange end.")
}
