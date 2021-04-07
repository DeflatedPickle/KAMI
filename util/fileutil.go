package util

import (
	"io/ioutil"
)

func SReadFile(path string) string {
	return string(ReadFile(path))
}

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	FCheckErr(err, "unable to read file: %v")
	return data
}

func CheckReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
