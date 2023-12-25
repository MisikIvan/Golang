package main

/*
#include <stdio.h>
#include <stdlib.h>

double calculateCost(double price, double width, double height) {
    return price * width * height;
}

double addWindowSill(double cost, int hasWindowSill) {
    if (hasWindowSill) {
        return cost + 350;
    }
    return cost;
}
*/
import "C"
import "fmt"

func main() {
	// Змінні для вхідних даних
	var glassType, width, height int
	var hasWindowSill bool

	// Зчитування даних від користувача
	fmt.Println("Введіть тип склопакета (1 - однокамерний, дерев'яний, 2 - двокамерний, дерев'яний, 3 - однокамерний, металевий, 4 - двокамерний, металевий, 5 - однокамерний, металопластиковий, 6 - двокамерний, металопластиковий):")
	fmt.Scan(&glassType)
	fmt.Println("Введіть ширину вікна в см:")
	fmt.Scan(&width)
	fmt.Println("Введіть висоту вікна в см:")
	fmt.Scan(&height)
	fmt.Println("Введіть наявність підвіконня (true - так, false - ні):")
	fmt.Scan(&hasWindowSill)

	// Змінна для ціни за 1 см2 склопакета
	var price float64

	// Вибір ціни за типом склопакета
	switch glassType {
	case 1:
		price = 2.5
	case 2:
		price = 3
	case 3:
		price = 0.5
	case 4:
		price = 1
	case 5:
		price = 1.5
	case 6:
		price = 2
	default:
		fmt.Println("Невірний тип склопакета")
		return
	}

	// Обчислення вартості склопакета
	cost := C.calculateCost(C.double(price), C.double(width), C.double(height))

	// Додавання вартості підвіконня
	cost = C.addWindowSill(cost, C.int(hasWindowSill))

	// Виведення результату
	fmt.Printf("Вартість склопакета складає %.2f грн\n", cost)
}
