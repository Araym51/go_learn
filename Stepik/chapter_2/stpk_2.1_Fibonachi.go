package main

import "fmt"

func main(){
	a := fibonacci(5)
	fmt.Print(a)
}

func fibonacci(n int) int {
    a, b := 0, 1
	for i:=0; n > i; i++{
		a, b = b, b + a
	}
	return a
}