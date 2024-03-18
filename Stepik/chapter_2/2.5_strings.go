package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strings"
)

func main(){
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strings.TrimSuffix(text, "\n")
	rs := []rune(text)
	if unicode.IsUpper(rs[0]) && (rs[len(rs)-1] == 10){
		fmt.Println("Right")
	}else{
		fmt.Println("Wrong")
	}
}