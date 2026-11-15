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

// Тесты для проверки maximum
func TestMaximum_ZeroSize(t *testing.T) {
	if got := maximum([]int{}); got != 0 {
		t.Errorf("ожидаемое значение 0 для пустого слайса, получено %d", got)
	}

	if got := maximum(nil); got != 0 {
		t.Errorf("ожидаемое значение 0 для nil-слайса, получено %d", got)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	data := []int{10}
	if got := maximum(data); got != 10 {
		t.Errorf("ожидаемое значение 10, получено %d", got)
	}
}

func TestMaximum_MaxAtBegin(t *testing.T) {
	data := []int{15000, 24, 100, 1, 160}
	if got := maximum(data); got != 15000 {
		t.Errorf("ожидаемое значение 15000, получено %d", got)
	}
}

func TestMaximum_MaxAtMiddle(t *testing.T) {
	data := []int{15000, 24, 20000, 1, 160}
	if got := maximum(data); got != 20000 {
		t.Errorf("ожидаемое значение 20000, получено %d", got)
	}
}

func TestMaximum_MaxAtEnd(t *testing.T) {
	data := []int{15000, 24, 100, 1, 31099}
	if got := maximum(data); got != 31099 {
		t.Errorf("ожидаемое значение 31099, получено %d", got)
	}
}

func TestMaximum_EqualElements(t *testing.T) {
	data := []int{1, 1, 1, 1}
	if got := maximum(data); got != 1 {
		t.Errorf("ожидаемое значение 7, получено %d", got)
	}
}

func TestMaximum_NegativeNumbers(t *testing.T) {
	data := []int{-1, -3, -100, -100500, -111}
	if got := maximum(data); got != -1 {
		t.Errorf("ожидаемое значение -100500, получено %d", got)
	}
}

// Тесты для проверки maxChunks
func TestMaxChunks_ZeroSize(t *testing.T) {
	if got := maxChunks([]int{}); got != 0 {
		t.Errorf("ожидаемое значение 0 для пустого слайса, получено %d", got)
	}

	if got := maxChunks(nil); got != 0 {
		t.Errorf("ожидаемое значение 0 для nil-слайса, получено %d", got)
	}
}

func TestMaxChunks_LessThanChunks(t *testing.T) {
	data := []int{1524124, 15352, 1, 53, 93993}
	if got := maxChunks(data); got != 1524124 {
		t.Errorf("ожидаемое значение 1524124, получено %d", got)
	}
}

func TestMaxChunks_NotDivByChunks(t *testing.T) {
	data := []int{1, 2, 3, 33, 44, 55, 66, 100, 101, 102, 103}
	if got := maxChunks(data); got != 103 {
		t.Errorf("ожидаемое значение 103, получено %d", got)
	}
}

func TestMaxChunks_CompareWithSingleThread(t *testing.T) {
	data := generateRandomElements(10_000)
	expected := maximum(data)
	got := maxChunks(data)

	if got != expected {
		t.Errorf("не совпадает результат сравнения: maximum вернул %d, maxChunks возвращает %d", expected, got)
	}
}
