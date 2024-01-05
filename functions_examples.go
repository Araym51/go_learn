package main

import "fmt"

// функция вернет переменную divide
func divid(x, y int) (divide int) {
	divide = x/y
	return //компилятор найдет указзаную переменную divide и вернет её без явного указания	
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





	fmt.Println("over")
}