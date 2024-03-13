package main

import "fmt"
/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, 
что φn=A. Если А не является числом Фибоначчи, выведите число -1.
*/
func main(){
	var x, i int
	a, b := 0, 1
	fmt.Scan(&x)
	for x != a  {
		a, b = b, a + b
		i++
		if a > x {
			break
		}
	}
	if a > x{
		fmt.Println(-1)
	}else{
	fmt.Println(i)
	}
}