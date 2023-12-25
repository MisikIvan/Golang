package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateCost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	days, _ := strconv.Atoi(r.FormValue("days"))
	country := r.FormValue("country")
	season := r.FormValue("season")
	guide := r.FormValue("guide") == "on"
	luxury := r.FormValue("luxury") == "on"

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

	fmt.Fprintf(w, "Вартість туру складає $%.2f\n", cost)
}

func main() {
	http.HandleFunc("/calculate", calculateCost)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
