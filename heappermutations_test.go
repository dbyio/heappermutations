// Copyright (c) 2024 Nicolas Perraud <np bitbox io>
// Thanks: @yberman

package heappermutations_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/dbyio/heappermutations"
)

func ExamplePermute() {
	s := []int{1, 2, 3}
	perms := heappermutations.Permute(s)
	fmt.Println(perms)
	//Output: [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]
}

func TestPermute(t *testing.T) {
	s := []string{"abc", "def", "ghi", "jkl"}
	perms := heappermutations.Permute(s)
	if len(perms) != 24 { // factorial(len(s)) == 24
		t.Error("Expected 24")
	}
}

var result [][]int

func BenchmarkInts10(b *testing.B) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(100))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = heappermutations.Permute(s)
	}
}
