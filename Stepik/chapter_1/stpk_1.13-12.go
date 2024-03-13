package main

import "fmt"
/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. 
Между числом и словом должен стоять ровно один пробел.
*/
func main(){
	var cow int
	fmt.Scan(&cow)
	if cow != 11 && cow % 10 == 1{
		fmt.Printf("%d korova", cow)
	}else if (cow % 10 == 2 || cow % 10 == 3 || cow % 10 == 4) && cow != 14 && cow != 13 && cow != 12 {
		fmt.Printf("%d korovy", cow)
	}else {
		fmt.Printf("%d korov", cow)
	}
}