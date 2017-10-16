// Copyright (c) 2017 Nicolas Perraud <np bitbox io>

// The heappermutations package implements primitives to generate all
// possible permutations following Heap's algorithm on typed collection.
package heappermutations

// An internal type that defines the generic structure of permutable collections
type heapInterface interface {
	// Len returns the number of elements in the collection
	Len() int
	// Swap swaps the elements with indexes i and j
	Swap(i, j int)
	// Copy copies the existing slice in a new slice
	Copy() heapInterface
}

// intSlice attaches the methods of heapInterface to []int.
type intSlice []int

func (p intSlice) Len() int      { return len(p) }
func (p intSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p intSlice) Copy() heapInterface {
	A := make(intSlice, p.Len())
	copy(A, p)
	return A
}

// float64Slice attaches the methods of heapInterface to []int.
type float64Slice []float64

func (p float64Slice) Len() int      { return len(p) }
func (p float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p float64Slice) Copy() heapInterface {
	A := make(float64Slice, p.Len())
	copy(A, p)
	return A
}

// stringSlice attaches the methods of heapInterface to []int.
type stringSlice []string

func (p stringSlice) Len() int      { return len(p) }
func (p stringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p stringSlice) Copy() heapInterface {
	A := make(stringSlice, p.Len())
	copy(A, p)
	return A
}

// Ints returns all permutations of a slice of ints.
func Ints(a []int) [][]int {
	interfaceSlice := heapPermutations(intSlice(a))
	permutations := make([][]int, len(interfaceSlice))

	// Assert and transfer each permuted item to our target-typed structure
	for i, d := range interfaceSlice {
		permutations[i] = d.(intSlice)
	}
	return permutations
}

// Strings returns all permutations of a slice of strings.
func Strings(a []string) [][]string {
	interfaceSlice := heapPermutations(stringSlice(a))
	permutations := make([][]string, len(interfaceSlice))

	for i, d := range interfaceSlice {
		permutations[i] = d.(stringSlice)
	}
	return permutations
}

// Float64s returns all permutations of a slice of float64s.
func Float64s(a []float64) [][]float64 {
	interfaceSlice := heapPermutations(float64Slice(a))
	permutations := make([][]float64, len(interfaceSlice))

	for i, d := range interfaceSlice {
		permutations[i] = d.(float64Slice)
	}
	return permutations
}

// An implementation of Heap's algorithm
func heapPermutations(data heapInterface) []heapInterface {
	permutations := []heapInterface{}
	var generate func(int, heapInterface)

	generate = func(n int, arr heapInterface) {
		if n == 1 {
			A := arr.Copy()
			permutations = append(permutations, A)
		} else {
			for i := 0; i < n; i++ {
				generate(n-1, arr)
				if n%2 == 0 {
					arr.Swap(i, n-1)
				} else {
					arr.Swap(0, n-1)
				}
			}
		}
	}
	generate(data.Len(), data)
	return permutations
}
