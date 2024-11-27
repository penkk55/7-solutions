package main

import (
	// "7-solutions/sola"
	"7-solutions/solb"
	"fmt"
	"sync"
)

// Declare a waitgroup
var wg sync.WaitGroup

// // Function to calculate the maximum path sum
// func maxPathSum(triangle [][]int) int {
// 	n := len(triangle)
// 	if n == 0 {
// 		return 0
// 	}

// 	// Bottom-up DP approach
// 	dp := make([]int, len(triangle[n-1]))
// 	copy(dp, triangle[n-1])

// 	for i := n - 2; i >= 0; i-- {
// 		for j := 0; j < len(triangle[i]); j++ {
// 			dp[j] = triangle[i][j] + max(dp[j], dp[j+1])
// 		}
// 	}

// 	return dp[0]
// }

// // Utility function to get the maximum of two integers
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // Function to fetch and parse JSON input from URL
// func fetchAndParseJSON() ([][]int, error) {
// 	jsonFile, err := os.Open("hard.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer jsonFile.Close()
// 	byteValue, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// [59], [73, 41], [52, 40, 53], [26, 53, 6, 34]
// 	// var result int
// 	var array [][]int
// 	json.Unmarshal([]byte(byteValue), &array)
// 	return array, nil
// }

func main() {

	// parallel

	// // Add(3) because of 3 functions of concurrent
	// wg.Add(3)

	// // Concurrent function list
	// go solA()
	// go solB()
	// go solC()
	// // Waiting to finish
	// wg.Wait()

	// async
	// sola.solA()
	// sola.SolA()
	solb.SolB()
	// solB()
	// solC()
}

// func solA() {
// 	// defer wg.Done()
// 	fmt.Println("s----solA----")
// 	// Test case 1: Static input
// 	staticTriangle := [][]int{
// 		{59},
// 		{73, 41},
// 		{52, 40, 53},
// 		{26, 53, 6, 34},
// 	}
// 	fmt.Println("Test case 1 result:", maxPathSum(staticTriangle)) // Expected: 237

// 	// Test case 2: Fetch input from URL
// 	triangle, err := fetchAndParseJSON()
// 	if err != nil {
// 		log.Fatalf("Error fetching or parsing JSON: %v", err)
// 	}
// 	fmt.Println("Test case 2 result:", maxPathSum(triangle)) // Expected: 7273
// 	fmt.Println("e----solA----")
// }

func solB() {
	// defer wg.Done()
	var i int
	fmt.Println("s----solB----")
	fmt.Scan(&i)
	fmt.Println("Your number is:", i)
	fmt.Println("e----solB----")
}

func solC() {
	// defer wg.Done()

	fmt.Println("s----solC----")
	fmt.Println("e----solC----")
}
