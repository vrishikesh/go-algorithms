package quick_sort

type QuickSortRecursive[T int] struct{}

func NewRecursive[T int]() *QuickSortRecursive[T] {
	return &QuickSortRecursive[T]{}
}

func (s *QuickSortRecursive[T]) Sort(items []T) {
	s.dive(items, 0, len(items)-1)
}

func (s *QuickSortRecursive[T]) dive(slice []T, low, high int) {
	if low >= high {
		return
	}

	j := s.partition(slice, low, high)
	s.dive(slice, low, j-1)
	s.dive(slice, j+1, high)
}

func (s *QuickSortRecursive[T]) partition(slice []T, low, high int) int {
	var i, j = low, high + 1
	for {
		for i = i + 1; i < high; i++ {
			if !s.less(slice, i, low) {
				break
			}
		}

		for j = j - 1; j > low; j-- {
			if !s.less(slice, low, j) {
				break
			}
		}

		if i >= j {
			break
		}
		s.swap(slice, i, j)
	}

	s.swap(slice, low, j)
	return j
}

func (s *QuickSortRecursive[T]) less(slice []T, x, y int) bool {
	return slice[x] < slice[y]
}

func (s *QuickSortRecursive[T]) swap(slice []T, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}

func (s *QuickSortRecursive[T]) Name() string {
	return "Recursive Quick"
}
