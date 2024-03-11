package main

import "fmt"

func main()  {
	array := []int{}
	var a, arr_elem int
	fmt.Scan(&a)
	for i:=0; i < a; i++{
		fmt.Scan(&arr_elem)
		array = append(array, arr_elem)
	}
	a = 0
	for i:=0; i < len(array); i++{
		if array[i] >= 0 {
			a ++
		}
	}
	fmt.Println(a)
}
