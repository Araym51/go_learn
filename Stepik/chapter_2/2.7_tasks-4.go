/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число. 

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
*/

package main

import ("fmt"
		"math"
		"strconv"
		)

func main(){
	var inpt_str, fnl_str string
	fmt.Scanln(&inpt_str)

	for _, value := range(inpt_str){
		y, err := strconv.Atoi(string(value))
		if err == nil{
			x := math.Pow(float64(y), 2)
			fnl_str += strconv.Itoa(int(x))
		}
	}
	fmt.Println(fnl_str)
}