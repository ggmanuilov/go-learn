package main

func main() {
	str := "threetwoone!"
	var arr [256]uint8
	for _, s := range str {
		arr[s]++
	}
	for i, s := range arr {
		if s != 0 {
			println(string(i), ":", s)
		}
	}
}
