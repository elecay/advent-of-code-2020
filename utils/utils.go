package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) []string {
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

	var result []string
	for scanner.Scan() {
		value := scanner.Text()
		result = append(result, value)
	}
	return result
}

func RemoveByIndex(elements []string, index int) []string {
	elements[index] = elements[len(elements)-1]
	return elements[:len(elements)-1]
}
