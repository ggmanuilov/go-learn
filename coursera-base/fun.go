package main

import "fmt"

//	именованныё результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multipleReturn(amount float32) (float32, error) {
	if amount > 10 {
		return 0, fmt.Errorf("invalida amount")
	}
	return amount, nil
}

// возврат нескольких именованных результатов
func multiNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error here!")
		// если заполнить err то оно будет возвражено
		// по умолчанию, даже без указания
		// return rez, err
		return
	}

	rez = 4
	return
}


// не фиксированное количество параметров одного типа
func sum(in ...int) (result int) {
	for _, val := range in {
		result += val
	}
	return result
}

func sumNoNamed(in ...int) int {
	result := 0
	for _, val := range in {
		result += val
	}
	return result
}

func main() {
	val, err := multiNamedReturn(true)

	if err != nil {
		fmt.Println("multiNamedReturn: ", err)
	} else {
		fmt.Println("multiNamedReturn val: ", val)
	}

	nums := []int{1, 2, 3, 4, 5, 6}

	// sum
	fmt.Println("sum: ", sum(1, 2, 3, 4, 5))

	// sum
	fmt.Println("sum nums: ", sum(nums...))

	// sumNoNamed
	fmt.Println("sumNoNamed: ", sumNoNamed(1, 2, 3, 4, 5))
}
