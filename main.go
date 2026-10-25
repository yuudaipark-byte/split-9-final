package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns a slice of positive random integers.
// For a non-positive size it returns an empty slice.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum returns the largest element of data, or 0 for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks finds the maximum using up to CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkCount := CHUNKS
	if len(data) < chunkCount {
		chunkCount = len(data)
	}
	chunkSize := len(data) / chunkCount
	chunkMaxima := make([]int, chunkCount)

	var wg sync.WaitGroup
	wg.Add(chunkCount)
	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == chunkCount-1 {
			end = len(data) // Include any remainder in the final chunk.
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	sequentialMax := maximum(data)
	sequentialElapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", sequentialMax, sequentialElapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	parallelMax := maxChunks(data)
	parallelElapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", parallelMax, parallelElapsed)
}
