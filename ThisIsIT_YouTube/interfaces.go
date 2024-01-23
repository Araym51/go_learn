package main

import "fmt"
//https://www.youtube.com/watch?v=Et3uJTXYNuI&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=13&ab_channel=ThisIsIT

// Вернуться к этому вопросу!

type Runner interface{
	Run() string
}

type Swimmer interface{
	Swim() string
}

type Flyer interface{
	Fly() string
}

type Ducker interface{
	Runner
	Swimmer
	Flyer
}

func main(){
	fmt.Println("Interfaces")
}