package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())

	for i := range arr {
		arr[i] = rand.Intn(100) + 1
	}

	slice := make([]int, len(arr))
	copy(slice, arr[:])

	sort.Ints(slice)

	fmt.Println("Исходный массив:    ", arr)
	fmt.Println("Отсортированный слайс:", slice)
}
