package mymath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Тестуємо додачу трьох чисел
	sum := Add(1.3, 2.33, -3.2)
	if sum != 0.93 {
		t.Errorf("Невірний результат: %f", sum)
	}

	// Тестуємо додачу нуля
	sum = Add(0, 0, 0)
	if sum != 0 {
		t.Errorf("Невірний результат: %f", sum)
	}
}
