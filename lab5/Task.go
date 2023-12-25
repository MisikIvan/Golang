package main

import (
	"fmt"
	"time"
)

// Worker структура, що представляє інформацію про працівника
type Worker struct {
	Name      string
	Year      int
	Month     int
	WorkPlace string
}

// Company структура, що містить інформацію про компанію
type Company struct {
	Name     string
	Position string
	Salary   float64
}

// NewWorker конструктор для створення нового об'єкту Worker
func NewWorker(name string, year, month int, workPlace string) *Worker {
	return &Worker{
		Name:      name,
		Year:      year,
		Month:     month,
		WorkPlace: workPlace,
	}
}

// NewCompany конструктор для створення нового об'єкту Company
func NewCompany(name, position string, salary float64) *Company {
	return &Company{
		Name:     name,
		Position: position,
		Salary:   salary,
	}
}

// GetWorkerPosition повертає посаду робочого
func (w *Worker) GetWorkerPosition() string {
	return w.WorkPlace
}

// GetWorkExperience обчислює та повертає стаж роботи у місяцях
func (w *Worker) GetWorkExperience() int {
	currentTime := time.Now()
	startDate := time.Date(w.Year, time.Month(w.Month), 1, 0, 0, 0, 0, time.UTC)
	months := int(currentTime.Sub(startDate).Hours() / 24 / 30)
	return months
}

// GetTotalMoney повертає загальну суму зароблених коштів за всі місяці роботи
func (w *Worker) GetTotalMoney() float64 {
	months := w.GetWorkExperience()
	return float64(months) * w.Salary
}

// ReadWorkersArray читає з клавіатури дані і повертає множину об’єктів типу Worker (n штук)
func ReadWorkersArray() []*Worker {
	var n int
	fmt.Print("Введіть кількість працівників: ")
	fmt.Scan(&n)

	workers := make([]*Worker, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nВведіть дані для працівника #%d:\n", i+1)
		fmt.Print("Прізвище та ініціали: ")
		var name string
		fmt.Scan(&name)

		fmt.Print("Рік початку роботи: ")
		var year int
		fmt.Scan(&year)

		fmt.Print("Місяць початку роботи: ")
		var month int
		fmt.Scan(&month)

		fmt.Print("Місце роботи: ")
		var workPlace string
		fmt.Scan(&workPlace)

		workers[i] = NewWorker(name, year, month, workPlace)
	}

	return workers
}

// PrintWorker приймає тип Worker і виводить його на екран
func PrintWorker(worker *Worker) {
	fmt.Printf("\nПрізвище та ініціали: %s\n", worker.Name)
	fmt.Printf("Рік початку роботи: %d\n", worker.Year)
	fmt.Printf("Місяць початку роботи: %d\n", worker.Month)
	fmt.Printf("Місце роботи: %s\n", worker.WorkPlace)
}

// PrintWorkers приймає множину типу Worker і виводить їх на екран
func PrintWorkers(workers []*Worker) {
	for i, worker := range workers {
		fmt.Printf("\nІнформація про працівника #%d:\n", i+1)
		PrintWorker(worker)
	}
}

// GetWorkersInfo приймає множину типу Worker і повертає через вихідні параметри найбільшу та найменшу зарплату серед усіх працівників
func GetWorkersInfo(workers []*Worker) (minSalary, maxSalary float64) {
	if len(workers) == 0 {
		return 0, 0
	}

	minSalary, maxSalary = workers[0].Salary, workers[0].Salary

	for _, worker := range workers {
		if worker.Salary < minSalary {
			minSalary = worker.Salary
		}

		if worker.Salary > maxSalary {
			maxSalary = worker.Salary
		}
	}

	return minSalary, maxSalary
}

func main() {
	// Читання даних про працівників
	workers := ReadWorkersArray()

	// Виведення інформації про працівників на екран
	fmt.Println("\nІнформація про працівників:")
	PrintWorkers(workers)

	// Отримання найбільшої та найменшої зарплати
	minSalary, maxSalary := GetWorkersInfo(workers)
	fmt.Printf("\nНайменша зарплата: %.2f\n", minSalary)
	fmt.Printf("Найбільша зарплата: %.2f\n", maxSalary)
}
