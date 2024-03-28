package main

import (
	"fmt"
	"strings"
)

func main(){
	input_str := ""
	fmt.Scanln(&input_str)
	for _, v := range input_str {
		if strings.Count(input_str, string(v)) > 1 {
			input_str = strings.Replace(input_str, string(v), "", -1)
		}
	}
	fmt.Println(input_str)
}