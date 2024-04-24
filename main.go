package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
		os.Exit(1)
	}

	path := os.Args[1]

	data, err := ReadFile(path)
	if err != nil {
		log.Fatal("Error reading data:", err)
	}

	average, median, variance, stdDeviation := CalcStats(data)

	fmt.Printf("Average: %d\n", round(average))
	fmt.Printf("Median: %d\n", round(median))
	fmt.Printf("Variance: %d\n", round(variance))
	fmt.Printf("Standard Deviation: %d\n", round(stdDeviation))
}

// ReadFile reads data from a file and returns a slice of integers.
func ReadFile(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// CalcStats calculates the average, median, variance, and standard deviation of a slice of integers.
func CalcStats(data []float64) (float64, float64, float64, float64) {
	// Calculate average
	var sum float64
	for _, value := range data {
		sum += value
	}
	average := float64(sum) / float64(len(data))

	// Calculate median
	sort.Float64s(data)
	var median float64
	if len(data)%2 == 1 {
		median = data[len(data)/2]
	} else {
		median = (data[len(data)/2-1] + data[len(data)/2]) / 2
	}

	// Calculate variance
	var varianceSum float64
	for _, value := range data {
		diff := float64(value) - average
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(len(data))

	// Calculate standard deviation
	stdDeviation := math.Sqrt(variance)

	return average, median, variance, stdDeviation
}

// function round rounds a float64 to the nearest integer.
func round(f float64) int {
	return int(math.Round(f))
}
