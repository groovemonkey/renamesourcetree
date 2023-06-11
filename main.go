package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func replaceInFile(path string, oldString string, newString string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)
	content = strings.ReplaceAll(content, oldString, newString)
	data = []byte(content)

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func renameIfNeeded(path string, oldString string, newString string) error {
	pathSegments := strings.Split(path, "/")
	// replace the string if it occurs in the last path segment
	// (all other occurrences higher up the filesystem tree will be separate entries so no need to do anything)
	lastIdx := len(pathSegments) - 1

	if strings.Contains(pathSegments[lastIdx], oldString) {

		// destructive modification of the string
		replacedSegment := strings.ReplaceAll(pathSegments[lastIdx], oldString, newString)
		pathSegments[lastIdx] = replacedSegment

		newPath := strings.Join(pathSegments, "/")
		if newPath != path {
			err := os.Rename(path, newPath)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func processDirectory(root string, oldString string, newString string) error {
	filePaths := make([]string, 0)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		filePaths = append(filePaths, path)
		return nil
	})

	if err != nil {
		return err
	}

	for i := len(filePaths) - 1; i >= 0; i-- {
		if info, err := os.Lstat(filePaths[i]); err == nil {
			if info.Mode()&os.ModeSymlink == 0 {
				// Regular file?
				if !info.IsDir() {
					err := replaceInFile(filePaths[i], oldString, newString)
					if err != nil {
						return err
					}
				}
				err := renameIfNeeded(filePaths[i], oldString, newString)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func main() {
	var startTime, endTime time.Time
	startTime = time.Now()

	// go run main.go --old=oldString --new=newString --targetdir=targetDirectory

	var oldString string
	var newString string
	var targetDir string
	var helpText string

	// TODO helptext
	flag.StringVar(&helpText, "help", "", "prints usage information")

	flag.StringVar(&oldString, "old", "", "The old string you want to replace. Enter a CamelCase string if any occurrences will be in CamelCase")
	flag.StringVar(&newString, "new", "", "The new string you want to replace the old one with. This must be CamelCase if you want to replace any CamelCase occurrences of the 'old' string.")
	flag.StringVar(&targetDir, "targetdir", "", "The target directory you want to process.")

	flag.Parse()

	// Set lowercase values
	oldStringLowercase := strings.ToLower(oldString)
	newStringLowercase := strings.ToLower(newString)

	// TODO logic based on if oldlower == old; newlower == new
	err := processDirectory(targetDir, oldString, newString)
	if err != nil {
		panic(fmt.Sprintf("error on first pass: %v", err))
	}

	err = processDirectory(targetDir, oldStringLowercase, newStringLowercase)
	if err != nil {
		panic(fmt.Sprintf("error on second pass: %v", err))
	}

	endTime = time.Now()
	fmt.Printf("Time elapsed: %s", endTime.Sub(startTime))
}
