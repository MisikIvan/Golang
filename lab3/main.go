package main

import (
	"fmt"
	"math"
	"math/rand"
)

func GeneratePseudoRandomIntegers(a, c, m, n int) []int {
	// Створюємо послідовність псевдовипадкових цілочислових значень.
	sequence := make([]int, n)
	x := rand.Intn(m)

	for i := 0; i < n; i++ {
		sequence[i] = (a*x + c) % m
	}

	return sequence
}

func CalculateFrequency(sequence []int, interval int) []int {
	// Створюємо масив частоти інтервалів появи випадкових величин.
	frequency := make([]int, interval)

	for _, value := range sequence {
		frequency[value%interval]++
	}

	return frequency
}

func CalculateProbability(frequency []int, n int) []float64 {
	// Створюємо масив статистичної імовірності появи випадкових величин.
	probability := make([]float64, len(frequency))

	for i := 0; i < len(probability); i++ {
		probability[i] = float64(frequency[i]) / float64(n)
	}

	return probability
}

func CalculateMean(sequence []int, n int) float64 {
	// Створюємо змінну для зберігання математичного сподівання.
	mean := 0.0

	for _, value := range sequence {
		mean += float64(value)
	}

	mean /= float64(n)

	return mean
}

func CalculateVariance(sequence []int, mean float64, n int) float64 {
	// Створюємо змінну для зберігання дисперсії.
	variance := 0.0

	for _, value := range sequence {
		variance += float64((value - mean) * (value - mean))
	}

	variance /= float64(n)

	return variance
}

func CalculateStandardDeviation(variance float64) float64 {
	// Створюємо змінну для зберігання середньоквадратичного відхилення.
	standardDeviation := math.Sqrt(variance)

	return standardDeviation
}

func main() {
	// Створюємо послідовність псевдовипадкових цілочислових значень.
	sequence := GeneratePseudoRandomIntegers(65539, 0, 4294967296, 30000)

	// Розраховуємо частоту інтервалів появи випадкових величин.
	frequency := CalculateFrequency(sequence, 1)

	// Розраховуємо статистичну імовірність появи випадкових величин.
	probability := CalculateProbability(frequency, 30000)

	// Розраховуємо математичне сподівання випадкових величин.
	mean := CalculateMean(sequence, 30000)

	// Розраховуємо дисперсію випадкових величин.
	variance := CalculateVariance(sequence, mean, 30000)

	// Розраховуємо середньоквадратичне відхилення випадкових величин.
	standardDeviation := CalculateStandardDeviation(variance)

	// Виводимо результати обробки.
	fmt.Println("Frequency:")
	for i, value := range frequency {
		fmt.Println(i, value)
	}

	fmt.Println("Probability:")
	for i, value := range probability {
		fmt.Println(i, value)
	}

	fmt.Println("Mean:", mean)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard deviation:", standardDeviation)
}
