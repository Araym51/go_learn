package main
import "fmt"

func main()  {
	array := [5]int{}
	var a, maximum int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	for idx, _ := range array{
		if maximum < array[idx]{
			maximum = array[idx]
		}
	}
	fmt.Println(maximum)
}