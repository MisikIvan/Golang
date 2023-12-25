package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateCost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	width, _ := strconv.ParseFloat(r.FormValue("width"), 64)
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	material := r.FormValue("material")
	sill := r.FormValue("sill") == "on"

	var price float64
	switch material {
	case "Однокамерний, дерев'яний":
		price = 2.5
	case "Двокамерний, дерев'яний":
		price = 3
	case "Однокамерний, металевий":
		price = 0.5
	case "Двокамерний, металевий":
		price = 1
	case "Однокамерний, металопластиковий":
		price = 1.5
	case "Двокамерний, металопластиковий":
		price = 2
	}

	cost := price * width * height
	if sill {
		cost += 350
	}

	fmt.Fprintf(w, "Вартість склопакета складає %.2f грн", cost)
}

func main() {
	http.HandleFunc("/calculate", calculateCost)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
