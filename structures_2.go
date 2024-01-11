package main

import "fmt"

// https://www.youtube.com/watch?v=Mi2mqnGm1h8&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=14&ab_channel=ThisIsIT

// встраивание типов в структуру

type Person struct{
	Name string
	Age int
}


type WorkExpirience struct{
	Name string
	Age int
}

type WoodBuilder struct{
	Person
	//Name string  // если создавать дубль поля от существующей стректуры - можно огрести проблем от Shadowing
	//WorkExpirience  // Colliding когда те же методы на том же уровне вложенности
}

func main(){
	fmt.Println("Interfaces 2")
	exp()
}

func exp(){
	builder := WoodBuilder{Person{Name: "Vasyan", Age: 20}}

	fmt.Printf("Type: %T Value: %#v\n", builder, builder)
	fmt.Println(builder.Person.Age) // ==
	fmt.Println(builder.Age)
	fmt.Println(builder.Name)

}