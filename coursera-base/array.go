package main

import "fmt"

func main() {
	// обявление массива
	deliveries := [...]int{1, 2, 3}
	fmt.Println("deliveries: ", deliveries)

	// слайс
	deliveriesDb := []int{1}
	deliveriesDb = append(deliveriesDb, 42)

	// сразу задаем с аллоцированием памяти из 5 элементов
	// len=5
	sliceWithSize := make([]int, 3)

	// обращение к элементу
	fmt.Println("val:", sliceWithSize[2])

	// добавление элементов len=10, cap=15
	sliceWithSize = append(sliceWithSize, 10, 15)
	// при увелечении размерности более чем есть
	// runtime делает x2 от размерности
	sliceWithSize = append(sliceWithSize, 19)

	// для добавления элементов другого слайса используют `...`
	sliceWithSize = append(sliceWithSize, deliveriesDb...)

	// просмотр иформации о слайсе
	// len - сколько элементов
	// cap - сколько памати аллоцировано
	fmt.Println("slice lin:", len(sliceWithSize), cap(sliceWithSize))

	fmt.Println("slice full:", sliceWithSize)

	// получение среза, 1,2,3
	sliceWithFilter := sliceWithSize[1:4] // >=1 && < 4
	fmt.Println("slice lin:", sliceWithFilter, len(sliceWithFilter), cap(sliceWithFilter))

	//
	sliceWithFilter = sliceWithSize[2:] // >=2 && до конца
	fmt.Println("slice from 3 ->:", sliceWithFilter, len(sliceWithFilter), cap(sliceWithFilter))

	//
	sliceWithFilter = sliceWithSize[3:5] // c начала и до 3 элемента
	fmt.Println("slice from < 4:", sliceWithFilter, len(sliceWithFilter), cap(sliceWithFilter))

	slice1 := []int{1, 2, 3}

	slice4 := slice1[:]
	slice4[0] = 10
	slice1[0] = 8
	// // то в slice1[0] будет 10
	fmt.Println("почти внезапно slice1[0]", slice1[0])
	fmt.Println("почти внезапно slice4[0]", slice4[0])

	// // но если
	slice5 := append(slice4[:], 22) //
	fmt.Println("почти внезапно slice1[0]", slice1[0], len(slice1))
	fmt.Println("почти внезапно slice5[last]", slice5[len(slice5)-1], len(slice5))

	// // и изменишь
	// slice1[0] = 42
	// // то в slice1[0] будет 1
	// fmt.Println("но без изменений slice1[0]: = 42", slice1[0])
	// fmt.Println("но без изменений slice2[0]: = 1", slice2[0])

	// // но если
	// slice3 := append(slice1[:], 33) //
	// // и изменишь
	// slice3[0] = 8
	// // то в slice1[0] будет 1
	// fmt.Println("но без изменений slice1[0]: = 8", slice1[0])
	// fmt.Println("но без изменений slice3[0]: = 42", slice3[0])

	// копирование
	numbers := []int{0, 1, 2, 3, 4}
	fmt.Println(numbers[1:4])
}
