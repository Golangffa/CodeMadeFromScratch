package calculator

import (
	"errors"
)

// Add performs addition of two numbers
func Add(x, y int) int {
	return x + y
}

// Subtract performs subtraction of two numbers
func Subtract(x, y int) int {
	return x - y
}

// Multiply performs multiplication of two numbers
func Multiply(x, y int) int {
	return x * y
}

// Divide performs division of two numbers
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}
	return x / y, nil
}

func main() {
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			res, err := processStack(expressions)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			fmt.Print("gocalc>")
		}
	}