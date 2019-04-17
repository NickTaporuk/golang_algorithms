package stack

import "fmt"

func StackExample() {
	s := New()

	s.Push("test")
	s.Push(1)

	fmt.Println("Len:=>", s.Len())


	for s.Len() != 0 {
		val := s.Pop()
		fmt.Println(val)
	}
}
