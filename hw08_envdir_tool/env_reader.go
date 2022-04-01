package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func getFileData(filePath string) (string, fs.FileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}

	defer func() {
		err = file.Close()
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var fileContent string

	for scanner.Scan() {
		fileContent = scanner.Text()
		break
	}
	return fileContent, fileInfo, err
}

func FillEnv(myEnvs Environment) Environment {
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]

		myEnvs[key] = EnvValue{
			Value:      splits[1],
			NeedRemove: false,
		}
	}
	return myEnvs
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirList, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	environment := FillEnv(make(Environment))
	for _, d := range dirList {
		if !d.IsDir() {
			filepath := fmt.Sprintf("%s/%s", dir, d.Name())
			fileContent, fileInfo, err := getFileData(filepath)
			if err != nil {
				return nil, err
			}

			trimsSlice := []string{" ", "\n", "\t", "\x00"}
			for _, trimSign := range trimsSlice {
				fileContent = strings.ReplaceAll(fileContent, "\x00", "\n")
				fileContent = strings.TrimRight(fileContent, trimSign)
			}

			var needRemove bool
			if fileInfo.Size() == 0 {
				needRemove = true
			} else {
				needRemove = false
			}
			environment[d.Name()] = EnvValue{
				Value:      fileContent,
				NeedRemove: needRemove,
			}
		}
	}
	return environment, nil
}
