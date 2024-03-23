package sliceop

import (
	"fmt"
	"testing"
)

func Test_filter(t *testing.T) {
	{
		s := []int{1, 2, 3}
		s = Filter(s, func(i int) bool { return i == 2 })
		if fmt.Sprint(s) != "[2]" {
			t.Error("assert 1 failed")
		}
	}
	{
		s := []int{1, 2, 3}
		s = Filter(s, func(i int) bool { return i != 2 })
		if fmt.Sprint(s) != "[1 3]" {
			t.Error("assert 2 failed")
		}
	}
	{
		s := []int{}
		s = Filter(s, func(i int) bool { return i == 2 })
		if fmt.Sprint(s) != "[]" {
			t.Error("assert 3 failed")
		}
	}
}

func Test_group(t *testing.T) {
	{
		s := []int{}
		s = Group(s, func(i int) int { return i }, func(i, j int) int { return i + j })
		if fmt.Sprint(s) != "[]" {
			t.Error("assert 0 failed")
		}
	}
	{
		s := []int{1, 2, 3}
		s = Group(s, func(i int) int { return i }, func(i, j int) int { return i + j })
		if fmt.Sprint(s) != "[1 2 3]" {
			t.Error("assert 1 failed")
		}
	}
	{
		s := []int{3, 2, 1}
		s = Group(s, func(i int) int { return i }, func(i, j int) int { return i + j })
		if fmt.Sprint(s) != "[1 2 3]" {
			t.Error("assert 2 failed")
		}
	}
	{
		s := []int{3, 2, 23, 3, 2, 1}
		s = Group(s, func(i int) int { return i }, func(i, j int) int { return i + j })
		if fmt.Sprint(s) != "[1 4 6 23]" {
			t.Error("assert 3 failed")
		}
	}

}

func Test_sort(t *testing.T) {
	{
		s := []int{3, 2, 23, 3, 2, 1}
		SortDesc(s, func(i int) int { return i })
		if fmt.Sprint(s) != "[23 3 3 2 2 1]" {
			t.Error("assert 0 failed")
		}
	}
}
