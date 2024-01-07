package main

import "fmt"

// функция вернет переменную divide
func divid(x, y int) (divide int) {
	divide = x/y
	return //компилятор найдет указзаную переменную divide и вернет её без явного указания	
}



func calc(x, y int, action(x, y int) int)int{
	return action(x, y)
}

// возврат функции из функции
func createDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}


func main() {
	// объявление функции и вередача в переменную
	var mult func(x, y int) int
	mult = func(x, y int) int {return x * y}
	res_mult := mult(2,3)
	fmt.Println(res_mult)


	// передача функции в переменную без объявления
	summa := func (x, y int) int {return x + y}
	res_sum := summa(2,3)
	fmt.Println(res_sum)


	fmt.Println("Вложенная функция:")
	sumFunction := func(x, y int) int {
		return x + y
	}
	fmt.Println(calc(2,3, sumFunc)) // функцию передаём без вызова!
	fmt.Println("over")
}