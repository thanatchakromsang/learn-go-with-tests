package arraysandslices

import "testing"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func TestSum(t *testing.T) {
	t.Run("collection of any sizes", func(t *testing.T) {
		numbers := []int{1, 3, 5}

		got := Sum(numbers)
		want := 9

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
