package app

import (
	"fmt"
	"io/ioutil"
	"os"
	ut "stretches-common-api/utils"
)

func WriteFormImageToFile(r RouteContext, formkey string, tmpfilename string) (bool, string, []byte, *os.File) {
	c := r.FiberCtx
	//GET FILE HEADER
	fileHeader, err := c.FormFile(formkey)
	if err != nil {
		SetAndSaveBadRequest(r, "Cannot read file header.", err, "ME00429")
		return false, "", []byte{}, nil
	}
	fmt.Printf("NAME %v\n", fileHeader.Filename)
	fmt.Printf("SIZE %v\n", fileHeader.Size)

	//OPEN FILE
	fileIOReader, err := fileHeader.Open()
	if err != nil {
		SetAndSaveBadRequest(r, "Cannot read file.", err, "ME00429")
		return false, "", []byte{}, nil
	}

	//CONTENT TO STRING
	byteContainer, err := ioutil.ReadAll(fileIOReader) // why the long names though?
	fmt.Printf("UploadFromFormToGCP READ SIZE:%d", len(byteContainer))
	contents := string(byteContainer)

	fmt.Println(" * UploadFromFormToGCP WRITE FILE")
	absolutetmpfilepath, tmpfile, err5 := ut.WriteFile(contents, tmpfilename)
	if err5 != nil {
		SetAndSaveBadRequest(r, "Cannot create temp file ", err5, "ME00429")
		return false, "", []byte{}, nil
	}
	fmt.Println("UploadFromFormToGCPWRITE FILE OK", absolutetmpfilepath)
	return true, absolutetmpfilepath, byteContainer, tmpfile
}
