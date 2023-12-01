The **heappermutations** package provides a primitive to generate all possible permutations of typed-datasets, following B. R. Heap's algorithm.

Based on [https://en.wikipedia.org/wiki/Heap's_algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm)


### Available functions

    func Permute[T any]([]T) [][]T
    
_Permute_ returns all permutations of a typed slice.


### Usage

```go
import ( "github.com/dbyio/heappermutations" )

sInts := []int{1, 2, 3}
pInts := heappermutations.Permute(sInts)		// [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]

sStrings := []string{"a", "b", "c"}
pStrings := heappermutations.Permute(sStrings)	// [[a b c] [b a c] [c a b] [a c b] [b c a] [c b a]]
```
