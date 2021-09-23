package main

import "fmt"

func main() {
	strVar := "name2"
	switch strVar {
	case "name":
		fallthrough
	case "test", "name2":
		// some code
	default:
		//
	}

	// выход из цикла, находясь внутри switch
	mapVal := map[string]string{"name": "Table of fish"}
Loop:
	for key, val := range mapVal {
		println("switch: ", key, val)
		switch {
		case key == "lastName":
			break                 // выйдет из кейса, заверша его на середине
			println("somethings") // не будет выведена
		case key == "name" && val == "Table of fish":
			println("end loop", val)
			break Loop
		}
	}

	// циклы - while(true)
	isRun := true
	for isRun {
		println("is runned")
		isRun = false
	}

	// цикл с условием
	for i := 0; i < 3; i++ {
		if i == 0 {
			continue
		}
		println("for i", i)
	}

	for idx, val := range mapVal {
		fmt.Println("range of slice is pretty:", idx, val)
	}

	// в мапе порядок ключей не опеределен
	// в разных запусках программы, порядок ключей бет разным!

	str := "Привет мир, Ёпта!"
	// mapKeys := []uint8{255}
	var arr [2048]uint
	for _, char := range str[:len(str)-1] {
		arr[char]++
	}

	for key, count := range arr {
		fmt.Println("key:", key, count)
	}
}
