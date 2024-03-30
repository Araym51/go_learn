package main

import (
	"fmt"
	"regexp"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/
func main() {
	pass := ""
	fmt.Scanln(&pass)
	if reg_exp(pass) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func reg_exp(input_string string) bool {
	re, _ := regexp.MatchString(`^[a-zA-Z0-9]{5,}$`, input_string)
	return re
}
