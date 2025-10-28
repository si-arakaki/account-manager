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
	accountHome := getAccountHome()

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
	hasFilenameMatched := false
	hasKeywordFoundInBlock := false

	for _, fileName := range fileNames {
		if keywordPattern.MatchString(fileName) {
			hasFilenameMatched = true
		}

		cleanup := func() {
			if len(blockLines) > 0 && (hasFilenameMatched || hasKeywordFoundInBlock) {
				fmt.Printf("\u001b[38;5;240m%s\u001b[0m\n", keywordPattern.ReplaceAllString(fileName, "\u001b[38;5;220m$1\u001b[0m\u001b[38;5;240m"))

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
			hasFilenameMatched = false
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

	fmt.Printf("accountHome: %s\nkeyword: %s\n", accountHome, keyword)
}

func getAccountHome() string {
	accountHome := os.Getenv("ACCOUNT_HOME")
	if accountHome == "" {
		accountHome = fmt.Sprintf("%s/.account", os.Getenv("HOME"))

		log.Warnf("no environment found ACCOUNT_HOME")
		log.Warnf("using %s instead", accountHome)
	}

	return accountHome
}

func getKeyword() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	keyword := ""

	if err := huh.NewInput().
		Title("keyword?").
		Value(&keyword).
		Run(); err != nil {
		return "", err
	}

	return keyword, nil
}
