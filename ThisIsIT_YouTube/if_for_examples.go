package main

import "fmt"

// if
func isChildren(age int) bool {
	return age < 18
}

func main() {
	age := 19

	//короткий синтаксис объявления функции
	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are kid")
	}

	// for
	for i := 0; i < 20; i++ {
		if i%2==1{
			continue // выход из текущей итерации
		}
		if i == 10 {
			break // принудительное завершение цикла
		}
		fmt.Println(i)
	}

	j := 1
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// бесконечный цикл
	for {
		fmt.Println("STOP ME PLEASE!")
	}

}
