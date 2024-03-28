package main

import "fmt"

func main(){
	input_string := ""
	new_string := ""
	fmt.Scanln(&input_string)
	for k, v := range input_string{
		if k % 2 != 0{
		new_string += string(v)
		}
	}
	fmt.Println(new_string)
} 