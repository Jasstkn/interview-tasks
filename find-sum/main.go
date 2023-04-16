package main

import (
	"bufio"
	"fmt"
	"os"
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
	desiredNum, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	var i, j int
	for i = 0; i < int(arraySize); i++ {
		for j = i + 1; j < int(arraySize); j++ {
			if array[i]+array[j] == desiredNum {
				fmt.Fprintf(writer, "%d %d", array[i], array[j])
				writer.Flush()
				return
			}
		}
	}
	fmt.Fprintf(writer, "None")

	writer.Flush()
}

func readArray(input string, len int64) []int64 {
	splitted := strings.Split(input, " ")
	array := make([]int64, len)

	for i, elem := range splitted {
		x, _ := strconv.ParseInt(strings.TrimSpace(elem), 10, 64)
		array[i] = x
	}

	return array
}
