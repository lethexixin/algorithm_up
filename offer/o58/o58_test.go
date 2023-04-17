package o58

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))

}

func TestF2(t *testing.T) {
	s := "abcdefg"
	n := 2
	fmt.Println(reverseLeftWords(s, n))

	s = "lrloseumgh"
	n = 6
	fmt.Println(reverseLeftWords(s, n))
}
