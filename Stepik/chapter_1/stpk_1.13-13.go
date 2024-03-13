package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var outputStr string

	_, _ = fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		x := int(math.Exp2(float64(i)))
		if x > n {
			break
		}
		outputStr += fmt.Sprintf("%d ", x)
	}
	fmt.Println(outputStr)
}