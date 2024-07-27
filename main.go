package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main () {

	count, err := MoveFiles(".pdf", "/PDFs")
	if err != nil {
		log.Print(err)
	}
	countStr := fmt.Sprintf("%d files moved.", count)
	fmt.Println(countStr)

}

func MoveFiles (ext string, extDir string) (int, error) {
	changesCount := 0
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dstDir := fmt.Sprintf("%s/Downloads", homeDir)
	fmt.Println(dstDir)
	dirContent, err := os.ReadDir(dstDir)
	if err != nil {
		log.Print("Error reading directory content, ",err)
	}
	_, err = os.ReadDir(dstDir)
	if os.IsNotExist(err) {
		fmt.Println("Dir does not exist, creating it now.")
		err := os.Mkdir(dstDir, 0750)
			if err != nil {
				log.Fatal(err)
			}
	}

	for i := 0; i < len(dirContent); i ++ {
		fileExt := filepath.Ext(dirContent[i].Name())
		if fileExt == ext {
			oldFilePath := fmt.Sprintf("%s/%s", dstDir, dirContent[i].Name())
			newFilePath := fmt.Sprintf("%s%s/%s", dstDir, extDir, dirContent[i].Name())
			err := os.Rename(oldFilePath, newFilePath)
			if err != nil {
				log.Fatal(err)
			}
			changesCount ++
		}

	}
	return changesCount, err
}