package utils

import (
	"bufio"
	"log"
	"os"
)

func loadDataFromFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed to read data from file: %s", err)
		return ""
	}
	content := string(bytes)
	return content
}

func WriteToFile(fileName string, input []string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range input {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed writing to file: %s", err)
		}
	}
	writer.Flush()
}
