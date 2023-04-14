package o50

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(string(firstUniqChar("abaccdeff")))
	fmt.Println(string(firstUniqChar("")))
}
