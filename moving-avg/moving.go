package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	arraySize, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	line, _ = reader.ReadString('\n')
	array := readArray(line, arraySize)

	line, _ = reader.ReadString('\n')
	k, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	arrayAvg := make([]float64, 0, arraySize-(k-1))

	// 1st sum
	var sum float64
	for i := 0; i < int(k); i++ {
		sum += float64(array[i])
	}
	arrayAvg = append(arrayAvg, sum/float64(k))

	for i := 0; i < int(arraySize-k); i++ {
		sum -= float64(array[i])
		sum += float64(array[i+int(k)])
		arrayAvg = append(arrayAvg, sum/float64(k))
	}

	for _, elem := range arrayAvg {
		fmt.Fprintf(writer, "%f ", elem)
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
