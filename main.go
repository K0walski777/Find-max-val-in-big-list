package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.Intn(math.MaxInt) + 1
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// инициализируем переменную с первым элементом в списке для последующего сравнения
	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	return 0
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randomSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(randomSlice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)

	//fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
