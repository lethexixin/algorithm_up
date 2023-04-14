package o10

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(43))
	fmt.Println(fib(44))
	fmt.Println(fib(45))

	fmt.Println(fib2(2))
	fmt.Println(fib2(5))
	fmt.Println(fib2(6))
	fmt.Println(fib2(43))
	fmt.Println(fib2(44))
	fmt.Println(fib2(45))
}

func TestF2(t *testing.T) {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
	fmt.Println(numWays(0))
}
