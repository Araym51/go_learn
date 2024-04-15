/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
Входные данные:
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
Выходные данные:
Выведите максимальную цифру, которая встречается во введенной строке.
*/

package main

import "fmt"

func main(){
	inpt_str := ""
	max_num := 0

	fmt.Scanln(&inpt_str)
	for _, value := range(inpt_str){
		if int(value) > max_num {
			max_num = int(value)
		}
	}
	fmt.Println(string(max_num))
}