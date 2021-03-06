package _9_Amortized_Time_Complexity

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArrayNP()
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	for i := 0; i < 4; i++ {
		arr.RemoveFirst()
		fmt.Println(arr)
	}
}
