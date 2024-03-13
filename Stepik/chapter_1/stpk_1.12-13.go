package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	maximum := array[0]
	for idx, _ := range array{
		if maximum < array[idx]{
			maximum = array[idx]
		}
	}
	fmt.Println(maximum)
}