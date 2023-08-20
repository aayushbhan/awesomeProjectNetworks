package HelperLibrary

import (
	"log"
	"os"
	"path/filepath"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetFilePath(fileName string) string {
	wd, err := os.Getwd()
	HandleError(err)
	return filepath.Join(wd, "Assets", fileName)
}
