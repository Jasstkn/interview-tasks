package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	arraySize, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	line, _ = reader.ReadString('\n')
	array := readArray(line, arraySize)

	line, _ = reader.ReadString('\n')
	desiredNum, _ := strconv.Atoi(strings.TrimSpace(line))

	sort.Ints(array)

	left := 0
	right := int(arraySize) - 1

	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == desiredNum {
			fmt.Fprintf(writer, "%d %d", array[left], array[right])
			writer.Flush()
			return
		}
		if currentSum < desiredNum {
			left++
		} else {
			right--
		}
	}

	fmt.Fprintf(writer, "None")

	writer.Flush()
}

func readArray(input string, len int64) []int {
	splitted := strings.Split(input, " ")
	array := make([]int, len)

	for i, elem := range splitted {
		x, _ := strconv.ParseInt(strings.TrimSpace(elem), 10, 0)
		array[i] = int(x)
	}
	return array
}
