package main

import (
	"testing"
)

// Тесты для проверки generateRandomElements
func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	got := generateRandomElements(0)
	if len(got) != 0 {
		t.Errorf("ожидался пустой слайс при size = 0, получена длина %d", len(got))
	}
}

func TestGenerateRandomElements_NegativeSize(t *testing.T) {
	got := generateRandomElements(-5)
	if len(got) != 0 {
		t.Errorf("ожидался пустой слайс при size = -5, получена длина %d", len(got))
	}
}

func TestGenerateRandomElements_SingleElement(t *testing.T) {
	got := generateRandomElements(1)
	if len(got) != 1 {
		t.Fatalf("ожидалась длина 1, получено %d", len(got))
	}
}

func TestGenerateRandomElements_StandSize(t *testing.T) {
	size := 100
	got := generateRandomElements(size)
	if len(got) != size {
		t.Fatalf("ожидалась длина %d, получено %d", size, len(got))
	}

	for i, v := range got {
		if v <= 0 {
			t.Errorf("элемент по индексу %d не положительный: %d", i, v)
			break
		}
	}
}
