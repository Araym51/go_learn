package main

import "fmt"

/*
palindrome
*/
func main(){
	var str, reverse_str string
	fmt.Scan(&str)
	
	for _, v := range(str){
		reverse_str = string(v) + reverse_str
	}

	if str == reverse_str {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}