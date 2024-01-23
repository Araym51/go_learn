package main

import (
	"fmt"
	"time"
)

// https://youtu.be/3SYe5TepPwI?si=wBZNlJt_1y4Lvbss
//при создании кастомного типа необходимо указать для него примитивный тип
type OurString string

// объявление структуры:
type Person struct{
	Name string // аналогично def __init__(self):
	Age int
}

func main(){
	var customString OurString = "Structures and types"

	fmt.Println(customString)
	// приведение к базовому типу:
	fmt.Println(string(customString))

	var Billy Person // иницилизация структуры. Можно ещё так: Billy = Person{}
	fmt.Printf("%T, %#v \n", Billy, Billy)
	// присвоение данных полям:
	Billy.Name = "Billy"
	Billy.Age = 18
	fmt.Println(Billy)

	Anna := Person{
		Name: "Ann",
		Age: 18,
	}
	fmt.Println(Anna)
	
	Pedro := Person{"Pedro", 18}

	fmt.Println(Pedro)

	// доступ к полям через указатель на структуру
	pPedro := &Pedro
	fmt.Println((*pPedro).Age)
	fmt.Println(pPedro.Age) // синтаксический сахар, результат записи аналогичен записи выше

	//создание указателя при создании переменной
	pIvan := &Person{"Ivan", 20}
	fmt.Println(pIvan)

	// анонимные структуры
	unnamedStruct := struct{
		Name, LastName, BirthDate string
	}{
		Name: "NoName",
		LastName: "NoLastName",
		BirthDate: fmt.Sprintf("%s", time.Now()),
		//BirthDate: time.Now().String() // или так
	}

	fmt.Println(unnamedStruct)
}