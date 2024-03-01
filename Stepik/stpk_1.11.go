package main

import "fmt"

func main(){
	var workArray [10]uint8
	var a, b, i int
	for fmt.Scan(&workArray[i]); i < 9; fmt.Scan(&workArray[i]){
		i++
	}
	i = 0
	for ;i < 3;{
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
		i ++
	}
	i = 0
	for _, elem := range workArray {
		fmt.Printf("%d ",elem)
	}
}