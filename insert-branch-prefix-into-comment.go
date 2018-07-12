package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func loadCurrentBranchName() (string, error) {
	branchName, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return string(branchName), err
}

func loadFileContent(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}

func extractPrefixName(name string) string {
	index := strings.Index(name, "/")
	if index >= 0 {
		return string(name[:index])
	} else {
		return ""
	}
}

func main() {
	filename := os.Args[1]
	content, err := loadFileContent(filename)
	if err != nil {
		log.Fatal(err)
	}

	branchName, err := loadCurrentBranchName()
	if err != nil {
		log.Fatal(err)
	}

	prefixName := extractPrefixName(branchName)
	if prefixName != "" && !strings.Contains(content, prefixName) {
		ioutil.WriteFile(filename, []byte(prefixName+" "+content), 0666)
	}
}
