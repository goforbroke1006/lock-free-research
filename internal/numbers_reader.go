package internal

import (
	"bufio"
	"os"
	"strconv"
)

func ReadSamples(filename string) ([]int64, error) {
	numbers := make([]int64, 0, 1024)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	return numbers, nil
}
