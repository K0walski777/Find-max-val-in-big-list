package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
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
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	// Создаем слайс для хранения локальных максимумов и вычисляем размер одного чанка
	maxVal := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		// Вычисляем границы среза
		start := i * chunkSize
		end := start + chunkSize

		// Проверяем длину слайса при делении на 8
		// если длина не делится на 8 нацело, то приравниваем end = len(data) чтобы не потерять данные
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Увеличиваем счетчик горутин до их запуска
		wg.Add(1)

		// Запускаем горутину, передав ей анонимную функцию с индексом и срезом
		// В анонимной функции запускаем ранее написанную функцию maximum
		go func(i int, chunk []int) {
			defer wg.Done()
			maxVal[i] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(maxVal)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	randomSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(randomSlice)
	elapsed := time.Since(start)

	// По ТЗ надо сделать вывод в микросекундах, но для удобстава я также сделал в милисекундах
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us (или %d ms)\n", max, elapsed.Microseconds(), elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(randomSlice)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us (или %d ms)\n", max, elapsed.Microseconds(), elapsed.Milliseconds())
}
