package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}
