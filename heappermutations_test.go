// Copyright (c) 2021 Nicolas Perraud <np bitbox io>

package heappermutations_test

import (
	"github.com/dbyio/heappermutations"
	"fmt"
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
	if len(perms) != 24 {		// factorial(len(s)) == 24
		t.Error("Expected 24")
	}
}
