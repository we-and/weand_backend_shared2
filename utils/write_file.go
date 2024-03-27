package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFile(content string, tempfilename string) (string, *os.File, error) {
	tmpfilepath := os.TempDir() // fp.Join(os.TempDir(), tempfilename)
	tmpFile, err := ioutil.TempFile(os.TempDir(), tempfilename)
	if err != nil {
		return "", tmpFile, errors.New("Cannot create temporary file")
	}

	// Remember to clean up the file afterwards
	//	defer os.Remove(tmpFile.Name())
	fmt.Println("Tempfilepath: " + tmpfilepath)

	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	text := []byte(content)
	if _, err = tmpFile.Write(text); err != nil {
		return "", tmpFile, errors.New("Failed to write to temporary file")
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		return "", tmpFile, err
	}

	fi, err := os.Stat(tmpFile.Name())
	if err != nil {
		return "", tmpFile, err
	}
	// get the size
	fmt.Printf("SIZE %v\n", fi.Size())
	fmt.Printf("SIZE INPUT %v\n", len(content))

	return tmpFile.Name(), tmpFile, nil
}
