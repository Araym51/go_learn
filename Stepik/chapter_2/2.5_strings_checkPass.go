package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/
func main(){
	pass := ""
	fmt.Scanln(&pass)
	fmt.Println(only_latin(pass))
}

func str_len(input_string string) bool{
	if len(input_string) >= 5 {
		return true
	} else {
		return false
	}
}

func only_latin(input_string string) bool{
	if utf8.RuneCountInString(input_string)==len(input_string) {
		return true
	} else {
		return false
	}

}