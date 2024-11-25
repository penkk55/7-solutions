package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Function to calculate the maximum path sum
func maxPathSum(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// Bottom-up DP approach
	dp := make([]int, len(triangle[n-1]))
	copy(dp, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + max(dp[j], dp[j+1])
		}
	}

	return dp[0]
}

// Utility function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to fetch and parse JSON input from URL
func fetchAndParseJSON(url string) ([][]int, error) {
	jsonFile, err := os.Open("hard.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	// [59], [73, 41], [52, 40, 53], [26, 53, 6, 34]
	// var result int
	var array [][]int
	json.Unmarshal([]byte(byteValue), &array)
	return array, nil
}

func main() {
	// Test case 1: Static input
	staticTriangle := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	fmt.Println("Test case 1 result:", maxPathSum(staticTriangle)) // Expected: 237

	// Test case 2: Fetch input from URL
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	triangle, err := fetchAndParseJSON(url)
	if err != nil {
		log.Fatalf("Error fetching or parsing JSON: %v", err)
	}
	fmt.Println("Test case 2 result:", maxPathSum(triangle)) // Expected: 7273
}
