package main

import "fmt"
/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. 
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

func main(){
	fmt.Println(sumInt(1, 0))
}

func sumInt(args ...int)(i int, sum int) {
	for i = range args {
		sum += args[i]
	}
	return len(args), sum
}