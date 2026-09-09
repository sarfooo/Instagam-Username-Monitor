package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func openFile(filename string) (slice []string) {
	file, _ := os.Open(filename)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	return slice
}

func splitSliceAtDelimiter(slice []string, delimiter string, index int) (splitSlice []string) {
	for i := range slice {
		parts := strings.Split(slice[i], delimiter)
		splitSlice = append(splitSlice, parts[index])
	}
	return splitSlice
}

func formatNumber(number int64) string {
	var in string = strconv.FormatInt(number, 10)
	var out []byte = make([]byte, len(in)+(len(in)-1)/3)
	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
