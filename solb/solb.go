package solb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Decode the encoded string to the smallest valid number set
func decodeToMinSumNumber(encoded string) string {

	n := len(encoded) + 1
	result := make([]int, n)
	// Initialize result with a baseline of 0
	for i := range result {
		result[i] = 0
	}
	// Forward pass to handle 'R' and '=' constraints
	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'R':
			// If 'R', right digit must be greater than the left digit
			if result[i+1] <= result[i] {
				result[i+1] = result[i] + 1
			}
		case '=':

			result[i+1] = result[i]

			fmt.Println("result_i", i)
		}
	}
	fmt.Println("result_1", result)

	// Backward pass to handle 'L' constraints (left digit > right digit)
	for i := len(encoded) - 1; i >= 0; i-- {
		fmt.Println("encoded[i]: ", encoded[i])

		if encoded[i] == 'L' && result[i] <= result[i+1] {
			result[i] = result[i+1] + 1
		}
	}

	// Normalize the result to make the smallest number start from 0

	minValue := result[0]
	for _, val := range result {
		if val < minValue {
			minValue = val
		}
	}
	// Subtract the minimum value from each result element
	for i := range result {
		result[i] -= minValue
	}

	// Convert result to string
	var sb strings.Builder
	for _, num := range result {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()

}

func SolB() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string: ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)

	decoded := decodeToMinSumNumber(encoded)
	fmt.Println("Decoded number sequence:", decoded)

}
