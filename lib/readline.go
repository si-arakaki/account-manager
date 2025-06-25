package lib

import (
	"bufio"
	"os"
)

type ReadLineEventHandler interface {
	OnReadLine(line string)
}

type ReadLineFunc func(line string)

var _ ReadLineEventHandler = (ReadLineFunc)(nil)

func (f ReadLineFunc) OnReadLine(line string) {
	f(line)
}

func ReadLine(fileName string, eventHandler ReadLineEventHandler) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		eventHandler.OnReadLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
