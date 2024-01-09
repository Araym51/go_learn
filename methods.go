package main

import (
	"fmt"
	"time"
)

type OurType string
//func (reciever T or *T) methodName(){}
func (t OurType) Hello(){
	fmt.Println("Hello")
}

type Square struct{
	Side int
}

func (s Square) Perimeter(){
	fmt.Printf("Периметр квадрата: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int){ //указатель на созданный объект! pointReciever
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func definition(){
	square := Square{Side: 4}
	pSquare := &square

	square.Perimeter()
	pSquare.Perimeter()
	fmt.Println("------------")
	pSquare.Scale(2)
	fmt.Println("------------")
	// но можно и так:
	square.Scale(2) // == (&square).Scale
	pSquare.Perimeter() // == (*pSquare).Perimeter() Squa{Side:8}
	
}