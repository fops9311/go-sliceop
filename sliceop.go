package sliceop

import (
	"sort"
)

func Filter[T any](s []T, filterFunc func(T) bool) []T {
	if len(s) == 0 {
		return s
	}
	start := 0
	end := 0
	for i := range s {
		allowed := filterFunc(s[i])
		if i == start && !allowed {
			start++
			end++
			continue
		}
		if allowed {
			s[i], s[end] = s[end], s[i]
			end++
			continue
		}
	}
	return s[start:end]
}

func SortAsc[T any, T1 Ordered](s []T, sortBy func(T) T1) {
	sort.Slice(s, func(i, j int) bool { return sortBy(s[i]) < sortBy(s[j]) })
}
func SortDesc[T any, T1 Ordered](s []T, sortBy func(T) T1) {
	sort.Slice(s, func(i, j int) bool { return sortBy(s[i]) > sortBy(s[j]) })
}

func Group[T any, T1 Ordered](s []T, groupBy func(T) T1, accum func(T, T) T) []T {
	if len(s) < 2 {
		return s
	}
	SortAsc(s, groupBy)
	ptr := 0
	totalLen := len(s)
	mem := groupBy(s[0])
	isSame := func(s T1) bool {
		result := mem == s
		mem = s
		return result
	}
	count := 0
	isSame(groupBy(s[0]))
	for i := 1; i < totalLen; i++ {
		if isSame(groupBy(s[i])) {
			s[ptr] = accum(s[ptr], s[i])
			continue
		}
		ptr++
		if i != ptr {
			copy(s[ptr:totalLen], s[i:totalLen])
		}
		lost := i - ptr
		totalLen -= lost
		i -= lost
		count += lost
	}
	lost := (totalLen - 1) - ptr
	count += lost
	return s[:len(s)-count]
}

type Ordered interface {
	Integer | Float | ~string
}
type Integer interface {
	Signed | Unsigned
}
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Float interface {
	~float32 | ~float64
}
