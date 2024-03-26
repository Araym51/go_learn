package main

import ("fmt"
		"strings"
		)
/*
check with "cordoba oba"
*/
func main(){
	word_1, word_2 := "", ""
	fmt.Scanln(&word_1)
	fmt.Scanln(&word_2)
	if strings.Contains(word_1, word_2){
		for _, i := range(word_2){
			fmt.Println(strings.Index(word_1, string(i)))
			break
		}
	} else {
		fmt.Println(-1)
	}
}