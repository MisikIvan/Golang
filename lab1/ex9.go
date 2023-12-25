package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	//Задание.
	//1. Пояснить результаты операций
	//!true - Логічне НЕ (!) для true дає false.
	//!false - Логічне НЕ для false дає true.
	//Логічне І (&&):

	//true && true - Логічне І для двох true дає true.
	//true && false - Логічне І для true та false дає false.
	//false && false - Логічне І для двох false дає false.
	//Логічне АБО (||):

	//true || true - Логічне АБО для двох true дає true.
	//true || false - Логічне АБО для true та false дає true.
	//false || false - Логічне АБО для двох false дає false.
	//Оператори порівняння:

	//2 < 3 - Перевірка, чи 2 менше 3 (true).
	//2 > 3 - Перевірка, чи 2 більше 3 (false).
	//3 < 3 - Перевірка, чи 3 менше 3 (false).
	//3 <= 3 - Перевірка, чи 3 менше або дорівнює 3 (true).
	//3 > 3 - Перевірка, чи 3 більше 3 (false).
	//3 >= 3 - Перевірка, чи 3 більше або дорівнює 3 (true).
	//2 == 3 - Перевірка, чи 2 дорівнює 3 (false).
	//3 == 3 - Перевірка, чи 3 дорівнює 3 (true).
	//2 != 3 - Перевірка, чи 2 не дорівнює 3 (true).
	//3 != 3 - Перевірка, чи 3 не дорівнює 3 (false).
}
