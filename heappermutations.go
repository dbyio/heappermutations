// Copyright (c) 2024 Nicolas Perraud <np bitbox io>

// The heappermutations package a generic  primitive to generate all possible
// permutations following Heap's algorithm on typed collection.
package heappermutations

// Permute returns all permutations of a slice
func Permute[T any](arr []T) [][]T {
	return heapPermutations(arr)
}

// An implementation of Heap's algorithm
func heapPermutations[T any](arr []T) [][]T {
	var permutations [][]T

	var generatePermutations func(int, []T)
	generatePermutations = func(n int, a []T) {
		if n == 1 {
			// Create a copy of the array to avoid modifying the original slice
			temp := make([]T, len(a))
			copy(temp, a)
			permutations = append(permutations, temp)
			return
		}

		for i := 0; i < n; i++ {
			generatePermutations(n-1, a)

			// Swap elements if n is even, otherwise swap the first and last element
			if n%2 == 0 {
				a[i], a[n-1] = a[n-1], a[i]
			} else {
				a[0], a[n-1] = a[n-1], a[0]
			}
		}
	}

	generatePermutations(len(arr), arr)
	return permutations
}
