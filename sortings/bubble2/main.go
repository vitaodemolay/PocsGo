package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Hacker Rank chalenge
 * https://www.hackerrank.com/challenges/ctci-bubble-sort/problem
 *
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(array []int32) {
	// Write your code here
	swaps := 0
	for i := 0; i < len(array)-1; i++ {
		hasSwapped := false
		for j := 0; j < len(array)-1-i; j++ {
			hasSwapped = swap(&array[j], &array[j+1])
			if hasSwapped {
				swaps++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", array[0])
	fmt.Printf("Last Element: %d\n", array[len(array)-1])
}

func swap(i, j *int32) bool {
	if *i > *j {
		*i, *j = *j, *i
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
