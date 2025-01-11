package utils

import (
	"golang.org/x/exp/constraints"
)

func Map2[T, U any](input []T, f func(T) U) []U {
	mapped := make([]U, len(input))
	for i, element := range input {
		mapped[i] = f(element)
	}
	return mapped
}

func Reduce[T, U any](input []T, f func(U, T) U, initial U) U {
	acc := initial
	for _, element := range input {
		acc = f(acc, element)
	}
	return acc
}

func Filter[T any](input []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, element := range input {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func Any[T any](input []T, f func(T) bool) bool {
	for _, element := range input {
		if f(element) {
			return true
		}
	}
	return false
}

func All[T any](input []T, f func(T) bool) bool {
	for _, element := range input {
		if !f(element) {
			return false
		}
	}
	return true
}

func Find[T any](input []T, f func(T) bool, defaultValue T) T {
	for _, element := range input {
		if f(element) {
			return element
		}
	}
	return defaultValue
}

func Index[T any](input []T, f func(T) bool) int {
	for i, element := range input {
		if f(element) {
			return i
		}
	}
	return -1
}

func Unique[T comparable](input []T) []T {
	keys := make(map[T]bool)
	var unique []T
	for _, element := range input {
		if _, exists := keys[element]; !exists {
			keys[element] = true
			unique = append(unique, element)
		}
	}
	return unique
}

func Zip[T any, U any](first []T, second []U) []struct {
	First  T
	Second U
} {
	n := len(first)
	m := len(second)
	var length int
	if m < n {
		length = m
	} else {
		length = n
	}
	var result []struct {
		First  T
		Second U
	}
	for i := 0; i < length; i++ {
		result = append(result, struct {
			First  T
			Second U
		}{first[i], second[i]})
	}
	return result
}

func Contains[T comparable](input []T, element T) bool {
	for _, e := range input {
		if e == element {
			return true
		}
	}
	return false
}

func Min[T constraints.Ordered](input []T, key func(T) T) T {
	if len(input) == 0 {
		panic("Slice cannot be empty.")
	}
	if key == nil {
		key = func(x T) T {
			return x
		}
	}
	minValue := input[0]
	minKey := key(minValue)
	for _, x := range input[1:] {
		if key(x) < minKey {
			minValue = x
			minKey = key(minValue)
		}
	}
	return minValue
}

func Max[T constraints.Ordered](input []T, key func(T) T) T {
	if len(input) == 0 {
		panic("Slice cannot be empty.")
	}
	if key == nil {
		key = func(x T) T {
			return x
		}
	}
	maxValue := input[0]
	maxKey := key(maxValue)
	for _, x := range input[1:] {
		if key(x) < maxKey {
			maxValue = x
			maxKey = key(maxValue)
		}
	}
	return maxValue
}
