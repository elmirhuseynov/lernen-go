package main

import "fmt"

func main() {
	var s = "()[]{}" // {()}
	var isValid = isValid(s)

	fmt.Println(isValid)
}

func isValid(s string) bool {
	stack := Stack{}

	for _, ch := range s {
		if ch == '(' {
			stack.Push(')')
		} else if ch == '{' {
			stack.Push('}')
		} else if ch == '[' {
			stack.Push(']')
		} else if stack.IsEmpty() || stack.Pop() != ch {
			return false
		}
	}

	return stack.IsEmpty()
}

type Stack struct {
	data []interface{}
}

func (stack *Stack) Push(d interface{}) {
	stack.data = append(stack.data, d)
}

func (stack *Stack) Pop() interface{} {
	if len(stack.data) == 0 {
		fmt.Println("Stack is empty.")
		return nil
	}
	top := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return top
}

func (stack *Stack) Peek() interface{} {
	if len(stack.data) == 0 {
		fmt.Println("Stack is empty.")
		return nil
	}
	return stack.data[len(stack.data)-1]
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.data) == 0
}
