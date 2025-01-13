package utils

func Range(start, stop, step int) []int {
	list := make([]int, 0)
	for i := start; i < stop; i += step {
		list = append(list, i)
	}
	return list
}

func CombinationsWithReplacement[T any](input []T, length int) [][]T {
	combinations := make([][]T, 0)
	n := len(input)
	if length > n || length < 0 {
		return nil
	}
	indices := make([]int, length)
	addCombination := func() {
		combination := make([]T, length)
		for i, idx := range indices {
			combination[i] = input[idx]
		}
		combinations = append(combinations, combination)
	}
	addCombination()
	for {
		i := length - 1
		b := false
		for ; i >= 0; i-- {
			if indices[i] != n-1 {
				b = true
				break
			}
		}
		if !b {
			break
		}
		indices[i]++
		for j := i + 1; j < length; j++ {
			indices[j] = indices[i]
		}
		addCombination()
	}
	return combinations
}
