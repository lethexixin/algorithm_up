package o58

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	s := "abcdefg"
	n := 2
	fmt.Println(reverseLeftWords(s, n))

	s = "lrloseumgh"
	n = 6
	fmt.Println(reverseLeftWords(s, n))
}
