package main

import (
	"fmt"
	"sync"
	"time"

	"elementary_sorts/bubble_sort"
	"elementary_sorts/heap_sort"
	"elementary_sorts/insertion_sort"
	"elementary_sorts/selection_sort"
	"elementary_sorts/shell_sort"
	"elementary_sorts/shuffling"
	"elementary_sorts/types"
	"merge_sort/merge_sort"
)

func main() {
	var slice = shuffling.Generate(10000)
	var wg sync.WaitGroup

	fmt.Printf("Random Numbers: %v\n\n", len(slice))

	sorts := []types.Sorter[int]{
		bubble_sort.New(),
		insertion_sort.New(),
		selection_sort.New(),
		heap_sort.New(slice...),
		shell_sort.New(),
		merge_sort.NewRecursive(),
	}

	wg.Add(len(sorts))

	for _, s := range sorts {
		s := s
		go func() {
			defer wg.Done()
			DoSort(s.Name(), s, slice)
		}()
	}

	wg.Wait()
}

func DoSort(sortName string, sorter types.Sorter[int], initialSlice []int) {
	var slice = make([]int, len(initialSlice))
	copy(slice, initialSlice)

	// Before(sortName, slice)
	start := time.Now()
	sorter.Sort(slice)
	elapsed := time.Since(start)
	After(sortName, slice, elapsed)
}

func Before(sortName string, slice []int) {
	fmt.Printf("%6s %9s Sort: %v\n", "Before", sortName, len(slice))
}

func After(sortName string, slice []int, elapsed time.Duration) {
	fmt.Printf("%6s %9s Sort: %v took %s\n\n", "After", sortName, len(slice), elapsed)
}
