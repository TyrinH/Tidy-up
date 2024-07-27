package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dstDir := fmt.Sprintf("%s/Downloads/", homeDir)
	dstDirContents, err := os.ReadDir(dstDir)
	if err != nil {
		log.Fatal(err)
	}
	extTypes := make(map[string]string)
	for i := 0; i < len(dstDirContents); i++ {
		ext := filepath.Ext(dstDirContents[i].Name())
		if ext == "" {
			continue
		}
		_, ok := extTypes[ext]
		if !ok {
			_,extDirString,_ := strings.Cut(ext, ".")
			extDirString = strings.ToUpper(extDirString)
			extTypes[ext] = extDirString
		}
	}
	fmt.Println(extTypes)
	if err != nil {
		log.Print(err)
	}
	for ext, extDir := range extTypes {
		wg.Add(1)
		go func(ext, extDir string) {
			defer wg.Done()
			err := MoveFiles(ext, extDir)
			if err != nil {
				log.Println(err)
			}
		}(ext, extDir)

	}
	wg.Wait()

}

func MoveFiles(ext string, extDir string) error {
	changesCount := 0
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dstDir := fmt.Sprintf("%s/Downloads/", homeDir)
	dirContent, err := os.ReadDir(dstDir)
	if err != nil {
		return fmt.Errorf("error reading directory content: %w", err)
	}
	extDirPath := fmt.Sprintf("%s%s", dstDir, extDir)
	fmt.Println(extDirPath)
	_, err = os.ReadDir(extDirPath)
	if os.IsNotExist(err) {
		fmt.Println("Dir does not exist, creating it now.")
		err := os.Mkdir(extDirPath, 0750)
			if err != nil {
				return err
			}
	}

	for i := 0; i < len(dirContent); i ++ {
		fileExt := filepath.Ext(dirContent[i].Name())
		if fileExt == ext {
			oldFilePath := fmt.Sprintf("%s%s", dstDir, dirContent[i].Name())
			newFilePath := fmt.Sprintf("%s%s/%s", dstDir, extDir, dirContent[i].Name())
			err := os.Rename(oldFilePath, newFilePath)
			if err != nil {
				log.Print(err)
			}
			changesCount ++
		}

	}
	log.Printf("%d %s files moved.", changesCount, ext)
	return nil
}