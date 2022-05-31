package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func getRPN(input string) (rpn []string, err error) {
	operators := make([]rune, 0)
	operand, operandsCounter := strings.Builder{}, 0
	for _, character := range input {
		if character == '+' || character == '-' {
			if operand.Len() == 0 {
				rpn = append(rpn, "0")
			} else {
				rpn = append(rpn, operand.String())
				operand.Reset()
				operandsCounter += 1
			}
			if n := len(operators); n > 0 && (operators[n-1] == '+' || operators[n-1] == '-') {
				rpn = append(rpn, string(operators[n-1]))
				operators = operators[:n-1]
			}
			operators = append(operators, character)
		} else if !unicode.IsSpace(character) {
			operand.WriteRune(character)
		}
	}
	if operand.Len() != 0 {
		rpn = append(rpn, operand.String())
		operandsCounter += 1
	}
	for _, operator := range operators {
		rpn = append(rpn, string(operator))
	}
	if len(rpn) == 0 || (len(rpn) == 1 && len(rpn[0]) == 0) {
		return nil, errorEmptyInput
	}
	if operandsCounter != 2 {
		return nil, errorNotTwoOperands
	}
	return rpn, nil
}

func calculate(rpn []string) (output string, err error) {
	stack := make([]int64, 0)
	for _, str := range rpn {
		switch str {
		case "+":
			if n := len(stack); n >= 2 {
				sum := stack[n-2] + stack[n-1]
				stack = stack[:n-2]
				stack = append(stack, sum)
			}
		case "-":
			if n := len(stack); n >= 2 {
				sum := stack[n-2] - stack[n-1]
				stack = stack[:n-2]
				stack = append(stack, sum)
			}
		default:
			number, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return "", fmt.Errorf("an error occurred: %w", err)
			}
			stack = append(stack, number)
		}
	}
	output = strconv.FormatInt(stack[0], 10)
	return
}

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
func StringSum(input string) (output string, err error) {
	rpn, err := getRPN(input)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	fmt.Println(rpn)
	output, err = calculate(rpn)
	return
}
