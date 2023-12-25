package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Задание.
	// 1. Создайте 2 переменные разных типов. Выполните арифметические операции. Результат вывести

	// Створення двох змінних різних типів
	var numberInt int = 42
	var numberFloat float64 = 3.14

	// Арифметичні операції
	sum := numberInt + int(numberFloat)
	difference := numberInt - int(numberFloat)
	product := numberInt * int(numberFloat)
	quotient := numberInt / int(numberFloat)

	// Вивід результатів
	fmt.Printf("\nАрифметичні операції:\n")
	fmt.Printf("Сума: %d\n", sum)
	fmt.Printf("Різниця: %d\n", difference)
	fmt.Printf("Добуток: %d\n", product)
	fmt.Printf("Частка: %d\n", quotient)
}
