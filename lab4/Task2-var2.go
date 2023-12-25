package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

double cost(char* country, char* season, int days, int guide, int luxury) {
    double price;
    if (strcmp(country, "Болгарія") == 0) {
        if (strcmp(season, "літо") == 0) {
            price = 100;
        } else {
            price = 150;
        }
    } else if (strcmp(country, "Німеччина") == 0) {
        if (strcmp(season, "літо") == 0) {
            price = 160;
        } else {
            price = 200;
        }
    } else if (strcmp(country, "Польща") == 0) {
        if (strcmp(season, "літо") == 0) {
            price = 120;
        } else {
            price = 180;
        }
    }

    double cost = days * price;
    if (guide) {
        cost += 50 * days;
    }
    if (luxury) {
        cost *= 1.2;
    }

    return cost;
}
*/
import "C"
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

	countryC := C.CString(country)
	seasonC := C.CString(season)
	daysC := C.int(days)
	guideC := C.int(0)
	if guide {
		guideC = C.int(1)
	}
	luxuryC := C.int(0)
	if luxury {
		luxuryC = C.int(1)
	}

	cost := C.cost(countryC, seasonC, daysC, guideC, luxuryC)

	fmt.Printf("Вартість туру складає $%.2f\n", float64(cost))
}
