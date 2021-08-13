// Copyright (c) 2021 Nicolas Perraud <np bitbox io>
// Thanks: @yberman

package heappermutations_test

import (
	"fmt"
	"github.com/dbyio/heappermutations"
	"testing"
)

func ExampleInts() {
	s := []int{1, 2, 3}
	perms := heappermutations.Ints(s)
	fmt.Println(perms)
	//Output: [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]
}

func TestStrings(t *testing.T) {
	s := []string{"abc", "def", "ghi", "jkl"}
	perms := heappermutations.Strings(s)
	if len(perms) != 24 { // factorial(len(s)) == 24
		t.Error("Expected 24")
	}
}

func BenchmarkInts(b *testing.B) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	b.ResetTimer()
	heappermutations.Ints(s)
}
