package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var filename string
	var cmd string

	quitted := false

	CreateMainUI()

	for !quitted {
		cmd = ShowCmdPrompt()
		switch cmd {
		case "select_file":
			filename = ShowFilePrompt(getFileList())
			if isDir(filename) {
				err := os.Chdir(filename)
				if err != nil {
					log.Fatalln(err)
				}
				filename = ""
			}
			fmt.Println(filename)
		case "quit":
			quitted = true
		}
		CallClear()
	}
}

func getFileList() []string {
	filelist := []string{
		"..",
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, aFile := range files {
		if isDir(aFile.Name()) {
			filelist = append(filelist, aFile.Name())
		} else {
			filelist = append(filelist, aFile.Name())
		}
	}

	return filelist
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	default:
		return false
	}
}
