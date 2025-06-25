package lib

import (
	"io/fs"
	"path/filepath"
)

type ListFileNameEventHandler interface {
	OnError(err error, fileName string)
}

type ListFileNameErrorFunc func(err error, fileName string)

var _ ListFileNameEventHandler = (ListFileNameErrorFunc)(nil)

func (f ListFileNameErrorFunc) OnError(err error, fileName string) {
	f(err, fileName)
}

func ListFileName(accountHome string, eventHandler ListFileNameEventHandler) []string {
	fileNames := make([]string, 0)

	filepath.Walk(accountHome, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			eventHandler.OnError(err, path)

			return nil
		}

		if info.IsDir() {
			return nil
		}

		fileNames = append(fileNames, path)

		return nil
	})

	return fileNames
}
