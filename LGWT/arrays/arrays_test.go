package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("want %d got %d given %v", want, got, numbers)
	}
}
func TestSumRange(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sumrange(numbers)
	want := 15
	if got != want {
		t.Errorf("want %d got %d given %v", want, got, numbers)
	}
}
func ExampleSum() {
	a := []int{1, 2, 3, 4, 5}
	sum := Sumrange(a)
	fmt.Println(sum)
	//Output:15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
