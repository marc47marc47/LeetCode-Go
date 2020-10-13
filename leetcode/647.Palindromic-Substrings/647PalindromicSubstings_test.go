package leetcode

import (
	"fmt"
	"testing"
)

func TestProblem647(t *testing.T) {
	s := "aaaa"
	fmt.Printf("%v: %v\n", s, countSubstrings(s))
	if countSubstrings(s) != 6 {
		panic(fmt.Errorf("result should by 6"))
	}
}
