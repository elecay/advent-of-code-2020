package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []int {
	localPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(localPath + path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []int
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, value)
	}
	return result
}

func RemoveByIndex(elements []int, index int) []int {
	elements[index] = elements[len(elements) - 1]
	return elements[:len(elements) - 1]
}
