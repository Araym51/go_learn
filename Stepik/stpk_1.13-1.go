package main

import "fmt"
/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные 
Вводится два целых числа a и b (a≤b).

Выходные данные 
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.
*/
func main()  {
	var a, b int
	fmt.Scan(&a, &b)
	num := b/7*7
	if num > a && b > 0 {
		fmt.Println(num)
	} else if b < 0{
		num = (b/7-1)*7
		fmt.Println(num)
	} else if num == 0 {
		fmt.Println(num)
    }else {
    fmt.Println("NO")
    }
}