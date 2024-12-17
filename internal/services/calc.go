package services

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	var ops []byte
	var nums []float64
	var i int

	precedence := map[byte]int{
		'+': 1, '-': 1,
		'*': 2, '/': 2,
	}

	applyOp := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return errors.New("invalid expression")
		}
		b, a := nums[len(nums)-1], nums[len(nums)-2]
		op := ops[len(ops)-1]
		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]

		var result float64
		switch op {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return errors.New("division by zero")
			}
			result = a / b
		}
		nums = append(nums, result)
		return nil
	}

	for i < len(expression) {
		switch {
		case expression[i] == ' ':
			i++
		case expression[i] >= '0' && expression[i] <= '9':
			start := i
			for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		case expression[i] == '(':
			ops = append(ops, expression[i])
			i++
		case expression[i] == ')':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			ops = ops[:len(ops)-1]
			i++
		case strings.ContainsRune("+-*/", rune(expression[i])):
			for len(ops) > 0 && ops[len(ops)-1] != '(' && precedence[ops[len(ops)-1]] >= precedence[expression[i]] {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, expression[i])
			i++
		default:
			return 0, errors.New("invalid character in expression")
		}
	}

	for len(ops) > 0 {
		if err := applyOp(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, errors.New("invalid expression")
	}
	return nums[0], nil
}
