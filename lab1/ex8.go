package main

// Імпорт нескольких пакетів
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultFloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Задание.
	// 1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести

	// Створення змінних різних типів
	shortFloat := float32(3.14)
	shortDouble := 42.0

	// Вивід результатів
	fmt.Printf("shortFloat (%T) = %f\n", shortFloat, shortFloat)
	fmt.Printf("shortDouble (%T) = %f\n", shortDouble, shortDouble)
}
