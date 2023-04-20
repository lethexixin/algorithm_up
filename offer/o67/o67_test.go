package o67

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(strToInt("15869yhk"))
	fmt.Println(strToInt("  -42"))
	fmt.Println(strToInt("  +42"))
	fmt.Println(strToInt("  -"))
	fmt.Println(strToInt("  +"))
	fmt.Println(strToInt("   "))
}
