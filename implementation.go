package lab2

import (
	"fmt"
	"strings"
)

func Operators(sign string) bool {
	return sign == "+" || sign == "-" || sign == "*" || sign == "/" || sign == "^"
}

func PostfixToInfix(postfix string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(postfix)
	for _, token := range tokens {
		if !Operators(token) {
			stack = append(stack, token)
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf("Incorrect postfix example")
			}
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			infix := fmt.Sprintf("(%s %s %s)", op1, token, op2)
			stack = append(stack, infix)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("Incorrect postfix example")
	}
	return stack[0], nil
}
