package main

import (
	"fmt"
	"strings"
)

func main(){
	var x, y string
	fmt.Scan(&x, &y)
	final := strings.ReplaceAll(string(x),string(y),"")
	fmt.Println(final)
}