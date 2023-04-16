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
	arraySize, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)

	line, _ = reader.ReadString('\n')
	arrayOne := readArray(line, arraySize)

	line, _ = reader.ReadString('\n')
	arrayTwo := readArray(line, arraySize)

	arrayZipped := make([]int64, 0, arraySize*2)

	for i := 0; i < int(arraySize); i++ {
		arrayZipped = append(arrayZipped, arrayOne[i])
		arrayZipped = append(arrayZipped, arrayTwo[i])
	}

	for _, elem := range arrayZipped {
		fmt.Fprintf(writer, "%d ", elem)
	}

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
