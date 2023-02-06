package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum of Array of length 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Sum of Slice of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Sum of multiple Slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9})
		want := []int{6, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum of tails of multiple Slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 9})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Safe sum of tails of multiple Slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 9}, []int{})
		want := []int{5, 9, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
