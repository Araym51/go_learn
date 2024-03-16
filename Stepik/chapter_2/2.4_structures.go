package main

import "fmt"

/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни.
Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем
testStruct в функции main, в дальнейшем программа проверит результат.
*/
type Character struct {
	On bool
	Ammo int
	Power int
}

func (field *Character) Shoot() bool {
	if (field.Ammo > 0) && (field.On == true) {
		field.Ammo -= 1
		return true 
	} else {
		return false
	}
}

func (field *Character) RideBike() bool{
	if (field.Power > 0) && (field.On == true){
		field.Power -= 1
		return true
	} else {
		return false
	}
}

func main() {
    sample := Character{true, 1, 2}
    testStruct := &sample
	fmt.Println(testStruct.On)
	fmt.Println(testStruct.Ammo)
	fmt.Println(testStruct.Power)
	fmt.Println(testStruct.Shoot(), testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), testStruct.Power)	
}

