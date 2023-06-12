package main

import "github.com/wontw/algo/structs/stack"

func isValid(s string) bool {
	var stack stack.Stack

	mapping1 := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	mapping2 := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, symbol := range s {
		var ok bool

		// if symbol is opening - add it to stack
		_, ok = mapping1[symbol]
		if ok {
			stack.Push(string(symbol))
			continue
		}

		// if symbol is closing - pop from stach and check if it closes right parentheses
		targetOpening, ok := mapping2[symbol]
		if !ok {
			return false
		}

		currentTop, _ := stack.Pop()

		if currentTop != string(targetOpening) {
			return false
		}
	}

	return stack.IsEmpty()
}
