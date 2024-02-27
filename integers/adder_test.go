package integers

import (
	"fmt"
	"slices"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
		numbers := [...]int{2, 3, 13, 3, 2}
		got := Sum(numbers[:])
		want := 23
		if got != want {
			t.Errorf("Want %d, but got %d, %v", want, got, numbers)
		}
	})

	t.Run("sum of collection of any size", func(t *testing.T) {
		numbers := []int{3, 3, 2}
		got := Sum(numbers)
		want := 8
		if got != want {
			t.Errorf("Want %d, but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of legit collection of slices", func(t *testing.T) {
		got := SumAll([]int{2, 3, 4}, []int{5, 6, 7})
		want := []int{9, 18}
		// want := "wanted"
		// if got != want {
		// if !reflect.DeepEqual(got, want){
		// ⬆️ Doesn't respect types
		if !slices.Equal(got, want) {
			t.Errorf("wanted %v, but got %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T){
	got := SumAllTails([]int{2,3,4}, []int{4,5,6})
	want := []int{7,11}
	if(!slices.Equal(got, want)){
		t.Errorf("wanted %v, but got %v", want, got)
	}
}