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
func decodeToMinSumNumber(encoded string) string {
	n := len(encoded) + 1
	result := make([]int, n)
	// result[0] = 2
	// Forward pass: handle 'R' and '=' constraints
	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'R': // right value must be greater
			if result[i+1] <= result[i] {
				result[i+1] = result[i] + 1
			}
		case '=': // equal values
			result[i+1] = result[i]
			fmt.Println("case '='", result[i])
			fmt.Println("end--->", encoded[i])

		}
	}

	// Backward pass: handle 'L' constraints
	for i := len(encoded) - 1; i >= 0; i-- {
		if encoded[i] == 'L' { // left value must be greater
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}
		}
	}

	// Normalize: adjust the result to start from 0
	minValue := result[0]
	for _, val := range result {
		if val < minValue {
			minValue = val
		}
	}
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
	// รับข้อมูลจาก keyboard
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string: ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)

	// แปลงข้อมูลและแสดงผล
	decoded := decodeToMinSumNumber(encoded)
	fmt.Println("Decoded number sequence:", decoded)
}
