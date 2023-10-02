package util

import (
	"io"
	"os"
)

func ReadFile(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()
	dat, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
