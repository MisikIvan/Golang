package main

import "fmt"

func main() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x

	fmt.Println("Битовые операции")

	fmt.Printf("^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)
	fmt.Printf("x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x<<2, x<<2)
	fmt.Printf("x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x>>2, x>>2)
	fmt.Printf("x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x&y, x&y)
	fmt.Printf("x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x|y, x|y)
	fmt.Printf("x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x^y, x^y)
	fmt.Printf("x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x&^y, x&^y)
	fmt.Printf("x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x%y, x%y)

	fmt.Println("\nБитовые операции с присваиванием")

	x = z
	x &= y
	fmt.Printf("x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x |= y
	fmt.Printf("x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x ^= y
	fmt.Printf("x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x &^= y
	fmt.Printf("x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x %= y
	fmt.Printf("x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)

	//Задание.
	//1. Пояснить результаты операций

	// ^x - Побітове НЕ (^) для числа x. Ця операція інвертує всі біти числа.
	// x << 2 - Зсув вліво (<<) для числа x на 2 позиції.
	// x >> 2 - Зсув вправо (>>) для числа x на 2 позиції.
	// x & y - Побітове І (&) для чисел x і y.
	// x | y - Побітове АБО (|) для чисел x і y.
	// x ^ y - Побітове ВИКЛЮЧНЕ АБО (^) для чисел x і y.
	// x &^ y - Побітове І НЕ (&^) для чисел x і y.
	// x % y - Остача від ділення числа x на y.

	//Операції з присвоєнням:

	// x &= y - Побітове І з присвоєнням.
	// x |= y - Побітове АБО з присвоєнням.
	// x ^= y - Побітове ВИКЛЮЧНЕ АБО з присвоєнням.
	// x &^= y - Побітове І НЕ з присвоєнням.
	// x %= y - Остача від ділення з присвоєнням.

}
