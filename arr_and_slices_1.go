package main
/* https://www.youtube.com/watch?v=-fhgodk2JLQ&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=16&ab_channel=ThisIsIT */
import "fmt"

func main(){
	arrays()
	slices()
}

func arrays(){
	fmt.Println("Arrays")

	var intArr [3]int
	fmt.Printf("TypeL %T, Value: %#v\n", intArr, intArr)
	
	intArr[0] = 5
	intArr[1] = 1

	fmt.Printf("TypeL %T, Value: %#v\n", intArr, intArr)

	// синтаксический сахар:
	
	people := [2]Person{
		{
			Age: 20,
			Name: "Katy",
		},
		{
			Age: 23,
			Name: "John",
		},
	}

	fmt.Printf("TypeL %T, Value: %#v\n", people, people)

	// неопределенное количество элементов
	stringArr := [...]string{"First", "Second", "Third", "Fourth"}
	fmt.Printf("TypeL %T, Value: %#v\n", stringArr, stringArr)

	fmt.Printf("TypeL %T, Value: %#v\n", len(stringArr), cap(stringArr))// len - длинна, cap - вместимость. len всегда равна cap

	// итерация по массиву:
	for i := 0; i < len(stringArr); i++ {
		fmt.Printf("TypeL %T, Value: %#v\n", i, stringArr[i])
	}

	for index, value := range stringArr{
		fmt.Printf("TypeL %T, Value: %#v\n", index, value)
	}

	for _, num := range intArr {
		fmt.Printf("Value: %#v\n", num)
	}

	newIntArr := changeArray(intArr)
	fmt.Printf("TypeL %T, Value: %#v\n", intArr, intArr)
	fmt.Printf("TypeL %T, Value: %#v\n", newIntArr, newIntArr)
}
// передача массива в функцию
func changeArray(arr [3]int) [3]int{
	arr[2] = 3
	return arr
}

// срезы массивов
func slices(){
	var defaultSlice []int // default value is nil
	fmt.Printf("TypeL %T, Value: %#v\n", defaultSlice, defaultSlice) 
	fmt.Printf("Len %T, Cap: %#v\n", len(defaultSlice), cap(defaultSlice))

	stringSliceLiteral := []string{"First", "Second"}
	fmt.Printf("TypeL %T, Value: %#v\n", stringSliceLiteral, stringSliceLiteral) 
	fmt.Printf("Len %T, Cap: %#v\n", len(stringSliceLiteral), cap(stringSliceLiteral))

	sliceByMake := make([]int, 3, 5) // помещаем 3 нуля дефолтных и задаем вместимось 5
	fmt.Printf("TypeL %T, Value: %#v\n", sliceByMake, sliceByMake) 
	fmt.Printf("Len %T, Cap: %#v\n", len(sliceByMake), cap(sliceByMake))

	// добавление элементов:
	sliceByMake = append(sliceByMake, 1,2,3,4)
	fmt.Println("Append:")
	fmt.Printf("TypeL %T, Value: %#v\n", sliceByMake, sliceByMake) 
	fmt.Printf("Len %T, Cap: %#v\n", len(sliceByMake), cap(sliceByMake))
	// цикл for
	for index, value := range sliceByMake {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}
