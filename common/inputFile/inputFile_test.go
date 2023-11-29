package inputFile

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var (
	text = "" +
		"line 1\n" +
		"line 2\n" +
		"line 3\n"
	path      = "textFile"
	mf        = mockFile{0, text, false}
	mock_open = func(p string) (file, error) {
		if p != path {
			return nil, os.ErrNotExist
		}
		return &mf, nil
	}
)

func TestReadInputFile(t *testing.T) {
	os_open = mock_open

	t.Run("FileExists", func(t *testing.T) {
		input := ReadInputFile(path)

		if len(input) != 3 {
			t.Fatalf("Expected %v lines, got %v", 3, len(input))
		}
		for index, line := range input {
			if line != fmt.Sprintf("line %v", index+1) {
				t.Fatalf("Expected 'line %v', got %v", index+1, line)
			}
		}
		if !mf.didClose {
			t.Error("File did not close")
		}
	})

	t.Run("FileNotExists_Panic", func(t *testing.T) {
		defer func() {
			err := recover()

			if err == nil {
				t.Fatal("Did not panic")
			} else {
				switch err := err.(type) {
				case error:
					if !errors.Is(os.ErrNotExist, err) {
						t.Fatalf("Wrong panic error %v", err)
					}
				default:
					t.Fatalf("Panic not an error! %v", err)
				}
			}
		}()

		ReadInputFile("")
	})

}
