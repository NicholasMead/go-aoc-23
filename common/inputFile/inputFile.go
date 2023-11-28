package inputFile

import (
	"bufio"
	"os"
)

type file interface {
	Read([]byte) (int, error)
	Close() error
}

var os_open = func(path string) (file, error) {
	return os.Open(path)
}

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
