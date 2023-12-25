package main

import "fmt"

func main() {
	//Инициализация переменных
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 //Такой вариант инициализации также возможен

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	//Краткая запись объявления переменной
	//только для новых переменных
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Задание.
	//1. Вывести типы всех переменных
	fmt.Printf("Type of userinit8: %T\n", userinit8)
	fmt.Printf("Type of userinit16: %T\n", userinit16)
	fmt.Printf("Type of userinit64: %T\n", userinit64)
	fmt.Printf("Type of userautoinit: %T\n", userautoinit)
	fmt.Printf("Type of intVar: %T\n", intVar)

	//2. Присвоить переменной intVar переменные userinit16 и userautoinit. Результат вывести.
	intVar = int(userinit16)
	fmt.Printf("Updated intVar after assigning userinit16: %d\n", intVar)

	intVar = int(userautoinit)
	fmt.Printf("Updated intVar after assigning userautoinit: %d\n", intVar)
}
