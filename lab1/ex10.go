package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Задание.
	// 1. Вывести украинскую букву 'Ї'
	ukrainianChar := 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrainianChar, ukrainianChar)

	// 2. Пояснить назначение типа "rune"
	var runeVar rune = 'Я'
	fmt.Printf("Code '%c' - %d (Type: %T)\n", runeVar, runeVar, runeVar)
}
