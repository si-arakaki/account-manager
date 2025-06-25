package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/si-arakaki/account-manager/lib"
)

func main() {
	accountHome := os.Getenv("ACCOUNT_HOME")
	if accountHome == "" {
		accountHome = fmt.Sprintf("%s/.account", os.Getenv("HOME"))

		log.Warnf("no environment found ACCOUNT_HOME")
		log.Warnf("using %s instead", accountHome)
	}

	mode, err := getMode()
	if err != nil {
		log.Error("getMode", "err", err)
		return
	}

	if mode == ReadAccountModeKeyword {
		keyword, err := getKeyword()
		if err != nil {
			log.Error("getKeyword", "err", err)
			return
		}

		keywordPattern := regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, keyword))

		fileNames := lib.ListFileName(accountHome, lib.ListFileNameErrorFunc(func(err error, fileName string) {
			log.Warn("lib.ListFileName", "err", err, "file", fileName)
		}))

		blockLines := make([]string, 0)
		hasKeywordFoundInBlock := false

		for _, fileName := range fileNames {
			cleanup := func() {
				if len(blockLines) > 0 && hasKeywordFoundInBlock {
					fmt.Printf("\u001b[38;5;240m%s\u001b[0m\n", fileName)

					for _, blockLine := range blockLines {
						if keywordPattern != nil {
							fmt.Println(keywordPattern.ReplaceAllString(blockLine, "\u001b[38;5;220m$1\u001b[0m"))
						} else {
							fmt.Println(blockLine)
						}
					}

					fmt.Println()
				}

				blockLines = make([]string, 0)
				hasKeywordFoundInBlock = false
			}

			lib.ReadLine(fileName, lib.ReadLineFunc(func(line string) {
				if len(line) > 0 && line[0:1] == "#" {
					cleanup()
				}

				if keywordPattern.MatchString(line) {
					hasKeywordFoundInBlock = true
				}

				blockLines = append(blockLines, line)
			}))

			cleanup()
		}
	}

	if mode == ReadAccountModeFileNameLike {
		fileNameLike, err := getFileNameLike()
		if err != nil {
			log.Error("getFileNameLike", "err", err)
			return
		}

		fileNamePattern := regexp.MustCompile(fmt.Sprintf(`(?i)(%s)`, fileNameLike))

		fileNames := lib.ListFileName(accountHome, lib.ListFileNameErrorFunc(func(err error, fileName string) {
			log.Warn("lib.ListFileName", "err", err, "file", fileName)
		}))

		for _, fileName := range fileNames {
			if !fileNamePattern.MatchString(fileName) {
				continue
			}

			fmt.Printf("\u001b[38;5;240m%s\u001b[0m\n", fileName)

			lib.ReadLine(fileName, lib.ReadLineFunc(func(line string) {
				fmt.Println(line)
			}))
		}
	}
}

type ReadAccountMode string

const ReadAccountModeKeyword ReadAccountMode = "keyword"

const ReadAccountModeFileNameLike ReadAccountMode = "FileNameLike"

func getMode() (ReadAccountMode, error) {
	mode := ReadAccountModeKeyword

	if err := huh.NewSelect[ReadAccountMode]().
		Title("mode?").
		Options(
			huh.NewOption("keyword", ReadAccountModeKeyword),
			huh.NewOption("filename", ReadAccountModeFileNameLike),
		).
		Value(&mode).
		Run(); err != nil {
		return "", err
	}

	return mode, nil
}

func getKeyword() (string, error) {
	keyword := ""

	if err := huh.NewInput().
		Title("keyword?").
		Value(&keyword).
		Run(); err != nil {
		return "", err
	}

	return keyword, nil
}

func getFileNameLike() (string, error) {
	keyword := ""

	if err := huh.NewInput().
		Title("filename?").
		Value(&keyword).
		Run(); err != nil {
		return "", err
	}

	return keyword, nil
}
