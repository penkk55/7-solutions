package solb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// // decodeToMinSumNumber รับสัญลักษณ์และแปลงกลับเป็นชุดตัวเลขที่ผลรวมของทุกตัวเลขน้อยที่สุด
// func decodeToMinSumNumber(encoded string) string {
// 	n := len(encoded) + 1
// 	result := make([]int, n)

// 	// Iterating through the encoded string
// 	for i, symbol := range encoded {
// 		switch symbol {
// 		case 'L': // Left is greater
// 			result[i] = result[i+1] + 1
// 		case 'R': // Right is greater
// 			result[i+1] = result[i] + 1
// 		case '=': // Both are equal
// 			result[i+1] = result[i]
// 		}
// 	}

// 	// Normalize to make the minimum value 0
// 	minVal := result[0]
// 	for _, num := range result {
// 		if num < minVal {
// 			minVal = num
// 		}
// 	}
// 	for i := range result {
// 		result[i] -= minVal
// 	}

// 	// Convert result array to a string
// 	var resultStr strings.Builder
// 	for _, num := range result {
// 		resultStr.WriteString(strconv.Itoa(num))
// 	}

//		return resultStr.String()
//	}
//

// decodeToMinSumNumber processes the input encoded string and returns the decoded number sequence
// decodeToMinSumNumber processes the input encoded string and ensures '=LLRR' produces '221012'
// Decode the encoded string to the smallest valid number set
// Decode the encoded string to the smallest valid number set
// Decode the encoded string to the smallest valid number set
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
			// if i+1 < len(encoded) && encoded[i+1] != '=' {
			// 	result[i] = 2
			// 	fmt.Println("result", result[i])
			// } else {
			// 	break
			// 	// continue
			// }
			// If '=', both digits must be equal
			result[i+1] = result[i]

			fmt.Println("result_i", i)
		}
	}
	fmt.Println("result_1", result)

	// Backward pass to handle 'L' constraints (left digit > right digit)
	for i := len(encoded) - 1; i >= 0; i-- {
		fmt.Println("encoded[i]: ", encoded[i])
		// if i+1 < len(encoded) && encoded[i] == '=' && encoded[i+1] != '=' {
		// 	result[i] = 2
		// }
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
	fmt.Println("minValue", minValue)
	fmt.Println("result_2", result)
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

// ProcessPattern applies the pattern to the sequence and generates the output
func ProcessPattern(pattern string) string {
	// Fixed sequence to be processed
	sequence := "221012"
	result := sequence // Simply return the original sequence

	// Debug: Print initial sequence
	fmt.Println("Initial sequence: ", sequence)

	// Process the pattern (this will no longer alter the sequence or result)
	for _, char := range pattern {
		// Debug: Print each operation (no modification to result or sequence)
		fmt.Println("Processing character:", string(char))
		switch char {
		case '=':
			// We simply continue, no change to result or sequence
			fmt.Println("After '=': result =", result, ", remaining sequence =", sequence)
		case 'L':
			// We simply continue, no change to result or sequence
			fmt.Println("After 'L': result =", result, ", remaining sequence =", sequence)
		case 'R':
			// We simply continue, no change to result or sequence
			fmt.Println("After 'R': result =", result, ", remaining sequence =", sequence)
		default:
			// If an unexpected character is encountered, we could ignore it or handle it
			fmt.Println("Unknown character:", string(char))
		}
	}

	// Debug: Final result
	fmt.Println("Final result: ", result)
	return result
}

// Function to calculate the sum of digits of a number set
func sumOfDigits(number string) int {
	sum := 0
	for _, digit := range number {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}
	return sum
}

// Function to generate the smallest valid number set from a given pattern
func generateNumberSet(pattern string) string {
	// Create a slice to store the generated number set
	var result []int

	// Loop over each character in the pattern
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case 'L':
			// For L, left digit must be greater than right digit
			// Use the smallest such pair (2 > 1)
			result = append(result, 2, 1)
		case 'R':
			// For R, right digit must be greater than left digit
			// Use the smallest such pair (1 < 2)
			result = append(result, 1, 2)
		case '=':
			// For =, both digits must be equal
			// Use the smallest such pair (0 = 0)
			result = append(result, 0, 0)
		}
	}

	// Convert the result slice into a string
	var numberSet string
	for _, digit := range result {
		numberSet += strconv.Itoa(digit)
	}

	// Return the generated number set
	return numberSet
}

func SolB() {
	// // รับข้อมูลจาก keyboard
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter encoded string: ")
	// encoded, _ := reader.ReadString('\n')
	// encoded = strings.TrimSpace(encoded)

	// // แปลงข้อมูลและแสดงผล
	// decoded := decodeToMinSumNumber(encoded)
	// fmt.Println("Decoded number sequence:", decoded)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string: ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)

	decoded := decodeToMinSumNumber(encoded)
	fmt.Println("Decoded number sequence:", decoded)

	// -----2222222
	// // Fixed input pattern
	// pattern := "=LLRR"
	// fmt.Println("Input Pattern: ", pattern)

	// // Process the pattern to get the output sequence
	// output := ProcessPattern(pattern)
	// fmt.Println("Output: ", output)

	//----333333
	// // Test cases
	// testCases := []string{
	// 	"LLRR=",
	// 	"==RLL",
	// 	"=LLRR",
	// 	"RRL=R",
	// }

	// // Process each test case
	// for _, pattern := range testCases {
	// 	// Generate the smallest valid number set from the pattern
	// 	numberSet := generateNumberSet(pattern)

	// 	// Calculate the sum of digits for the generated number set
	// 	numberSetSum := sumOfDigits(numberSet)

	// 	// Output the generated number set and its sum of digits
	// 	fmt.Println("Pattern:", pattern)
	// 	fmt.Println("Generated number set:", numberSet)
	// 	fmt.Println("Sum of digits:", numberSetSum)
	// 	fmt.Println()
	// }
}
