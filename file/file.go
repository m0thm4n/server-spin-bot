package file

import (
	"fmt"
	"io"
	"os"
)

// WriteToFile writes to a file
func WriteToFile(filename string, games []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, game := range games {
		_, err := io.WriteString(file, game)
		_, err = io.WriteString(file, "\n")
		if err != nil {
			return err
		}
	}

	return file.Sync()
}

// FileExists check for the file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func MakeFile(filename string) {
	if FileExists("games.txt") {
		fmt.Println("This is a random game picker")
	} else {
		os.Create("games.txt")
	}
}
