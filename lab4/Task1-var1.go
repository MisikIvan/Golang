package main

import "fmt"

// Константи для цін за 1 см2 склопакета
const (
	OneChamberWooden  = 2.5
	TwoChamberWooden  = 3
	OneChamberMetal   = 0.5
	TwoChamberMetal   = 1
	OneChamberPlastic = 1.5
	TwoChamberPlastic = 2
)

// Константа для вартості підвіконня
const WindowSill = 350

// Функція для обчислення вартості склопакета
func CalculateCost(price, width, height float64) float64 {
	return price * width * height
}

// Функція для додавання вартості підвіконня
func AddWindowSill(cost float64, hasWindowSill bool) float64 {
	if hasWindowSill {
		return cost + WindowSill
	}
	return cost
}

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
		price = OneChamberWooden
	case 2:
		price = TwoChamberWooden
	case 3:
		price = OneChamberMetal
	case 4:
		price = TwoChamberMetal
	case 5:
		price = OneChamberPlastic
	case 6:
		price = TwoChamberPlastic
	default:
		fmt.Println("Невірний тип склопакета")
		return
	}

	// Обчислення вартості склопакета
	cost := CalculateCost(price, float64(width), float64(height))

	// Додавання вартості підвіконня
	cost = AddWindowSill(cost, hasWindowSill)

	// Виведення результату
	fmt.Printf("Вартість склопакета складає %.2f грн\n", cost)
}
