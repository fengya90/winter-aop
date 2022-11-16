package util

import (
	"io/ioutil"
	"path"
	"strings"
)

func GetFilePathsInTheDirWithSuffix(dir, suffix string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, file := range files {
		if file.IsDir() {
			tmpFiles, err := GetFilePathsInTheDirWithSuffix(path.Join(dir, file.Name()), suffix)
			if err != nil {
				return nil, err
			} else {
				result = append(result, tmpFiles...)
			}
		}
		if suffix == "" || strings.HasSuffix(file.Name(), suffix) {
			result = append(result, path.Join(dir, file.Name()))
		}
	}
	return result, nil
}
