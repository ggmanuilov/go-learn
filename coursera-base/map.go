package main

import "fmt"

func main() {
	var user map[string]string = map[string]string{
		"name":     "Vasia",
		"lastname": "Petrov",
	}
	// создаем сразу нужной ёмкостью
	profile := make(map[string]string, 10)
	// количество элементов
	fmt.Println("len:", len(user))
	fmt.Println("len:", len(profile))
	// если ключа нет, вернет значение по умолчанию для типа
	mName, mNameExist := user["middleName"]
	// но его нет в мэре! go вернет занчене по умолчанию
	fmt.Println("mName:", mName, len(mName))
	// проверим что ключ был инициирован, а не вернулось значение по умолчанию
	fmt.Println("key mName exists:", mNameExist)
}
