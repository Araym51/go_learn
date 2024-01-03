package main

import "fmt"

func main(){
	var a, b int
	var operation string
	fmt.Scan(&a, &operation, &b)
	switch{
	case operation == "+":
		fmt.Println(a+b)
	case operation == "-":
		fmt.Println(a-b)
	case operation == "*":
		fmt.Println(a*b)
	case operation == "/":
		fmt.Println(a/b)
	}
	
}