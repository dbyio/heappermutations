The **heappermutations** package provides Golang primitives to generate all possible permutations of typed-datasets, following B. R. Heap's algorithm.

Based on [https://en.wikipedia.org/wiki/Heap's_algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm)


### Available functions


    func Float64s(a []float64) [][]float64

_Float64s_ returns all permutations of a slice of float64s.


    func Ints(a []int) [][]int

_Ints_ returns all permutations of a slice of ints.


    func Strings(a []string) [][]string

_Strings_ returns all permutations of a slice of strings.


### Usage

```go
import ( "github.com/dbyio/heappermutations" )

sInts := []int{1, 2, 3}
pInts := heappermutations.Ints(sInts)		// [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]

sStrings := []string{"a", "b", "c"}
pStrings := heappermutations.Strings(sStrings)	// [[a b c] [b a c] [c a b] [a c b] [b c a] [c b a]]
```
