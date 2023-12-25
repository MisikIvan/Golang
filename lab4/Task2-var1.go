package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть країну: ")
	country, _ := reader.ReadString('\n')
	country = strings.TrimSpace(country)

	fmt.Print("Введіть пору року: ")
	season, _ := reader.ReadString('\n')
	season = strings.TrimSpace(season)

	fmt.Print("Введіть кількість днів: ")
	daysStr, _ := reader.ReadString('\n')
	days, _ := strconv.Atoi(strings.TrimSpace(daysStr))

	fmt.Print("Чи потрібен гід (так/ні): ")
	guideStr, _ := reader.ReadString('\n')
	guide := strings.TrimSpace(guideStr) == "так"

	fmt.Print("Чи вибрано номер люкс (так/ні): ")
	luxuryStr, _ := reader.ReadString('\n')
	luxury := strings.TrimSpace(luxuryStr) == "так"

	var price float64
	switch country {
	case "Болгарія":
		if season == "літо" {
			price = 100
		} else {
			price = 150
		}
	case "Німеччина":
		if season == "літо" {
			price = 160
		} else {
			price = 200
		}
	case "Польща":
		if season == "літо" {
			price = 120
		} else {
			price = 180
		}
	}

	cost := float64(days) * price
	if guide {
		cost += 50 * float64(days)
	}
	if luxury {
		cost *= 1.2
	}

	fmt.Printf("Вартість туру складає $%.2f\n", cost)
}
