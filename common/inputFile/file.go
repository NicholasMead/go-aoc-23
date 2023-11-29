package inputFile

import "os"

type file interface {
	Read([]byte) (int, error)
	Close() error
}

var os_open = func(path string) (file, error) {
	return os.Open(path)
}
