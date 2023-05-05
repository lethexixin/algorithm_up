package o16

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(myPow1(3, 5))
	fmt.Println(myPow1(2.00000, 10))
	fmt.Println(myPow1(2.00000, -2))

	fmt.Println(myPow(3, 9))
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.00000, -2))
}
