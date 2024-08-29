package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Функция для вычисления медианы
func median(numbers []int) float64 {
	sort.Ints(numbers)
	result := float64(0)
	if lenNumbers := len(numbers); lenNumbers%2 == 0 {
		result = float64(numbers[lenNumbers/2-1]+numbers[lenNumbers/2]) / 2
	} else {
		result = float64(numbers[lenNumbers/2-1])
	}
	return result
}

// Функция для вычисления моды
func mode(numbers []int) int {
	counts := make(map[int]int)
	maxCount := 0
	for _, number := range numbers {
		counts[number]++
		if maxCount > counts[number] {
			maxCount = counts[number]
		}
	}
	mode := 0
	for numberKey, countValue := range counts {
		if countValue >= maxCount {
			mode = numberKey
		}
	}
	return mode
}

// Функция для вычисления среднего значения
func mean(numbers []int) float64 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}

// Функция для вычисления стандартного отклонения
func standardDeviation(numbers []int) float64 {
	m := mean(numbers)
	var sumSquares float64
	for _, number := range numbers {
		sumSquares += (float64(number) - m) * (float64(number) - m)
	}
	return math.Sqrt(sumSquares / float64(len(numbers)))
}

// Функция для чтения входных данных и проверки корректности
func scanInput() ([]int, error) {
	numbers := make([]int, 0, 8)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		number, errAtoi := strconv.Atoi(line)
		if errAtoi != nil || number < -100000 || number > 100000 {
			return nil, errors.New("invalid input: " + line)
		}
		numbers = append(numbers, number)

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func main() {
	meanFlag := flag.Bool("mean", false, "Display mean")
	medianFlag := flag.Bool("median", false, "Display median")
	modeFlag := flag.Bool("mode", false, "Display mode")
	SDFlag := flag.Bool("sd", false, "Display standard deviation")

	flag.Parse()

	// Чтение входных данных
	numbers, err := scanInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Если не указаны конкретные флаги, показываем все метрики
	if !*meanFlag && !*medianFlag && !*modeFlag && !*SDFlag {
		*meanFlag = true
		*medianFlag = true
		*modeFlag = true
		*SDFlag = true
	}

	// Вычисление и вывод выбранных метрик
	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", mean(numbers))
	}
	if *medianFlag {
		fmt.Printf("Median: %.2f\n", median(numbers))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", mode(numbers))
	}
	if *SDFlag {
		fmt.Printf("SD: %.2f\n", standardDeviation(numbers))
	}
}
