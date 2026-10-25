package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-3, 0, 1, 10, 100} {
		data := generateRandomElements(size)
		wantSize := size
		if wantSize < 0 {
			wantSize = 0
		}
		if len(data) != wantSize {
			t.Fatalf("generateRandomElements(%d) returned %d elements, want %d", size, len(data), wantSize)
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("generateRandomElements(%d) returned non-positive value %d", size, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "single", data: []int{12}, want: 12},
		{name: "positive", data: []int{2, 9, 1, 8}, want: 9},
		{name: "negative", data: []int{-10, -7, -20}, want: -7},
		{name: "mixed", data: []int{-5, 0, 10, -1}, want: 10},
		{name: "duplicates", data: []int{4, 4, 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "one", data: []int{5}, want: 5},
		{name: "fewer than eight", data: []int{1, 3, 2, 9, 4}, want: 9},
		{name: "exactly eight", data: []int{8, 7, 6, 5, 4, 3, 2, 1}, want: 8},
		{name: "nine, remainder at end", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{name: "seventeen, remainder at end", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 99}, want: 99},
		{name: "negative", data: []int{-20, -30, -10, -45, -5, -100, -40, -55, -60}, want: -5},
		{name: "zero", data: []int{-10, -5, 0, -1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestParallelMatchesSequential(t *testing.T) {
	generator := rand.New(rand.NewSource(42))
	for size := 0; size <= 1000; size++ {
		data := make([]int, size)
		for i := range data {
			data[i] = generator.Intn(2001) - 1000
		}
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Fatalf("length %d: maxChunks = %d; maximum = %d", size, got, want)
		}
	}
}
