package main

/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
Формат входных данных:
Вводится непустая последовательность натуральных чисел,
оканчивающаяся числом 0 (само число 0 в последовательность не входит,
а служит как признак ее окончания).
Формат выходных данных. Выведите ответ на задачу.
*/
import "fmt"

func main() {
	var n, max, i int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if max <= n {
			if max < n {
				max = n
				i = 1
			} else {
				i++
			}
		}
	}
	fmt.Println(i)
}
