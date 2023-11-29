package inputFile

import (
	"bufio"
)

func ReadInputFile(filePath string) []string {
	file, err := os_open(filePath)
	if err != nil {
		panic(err)
	} else {
		defer file.Close()
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
