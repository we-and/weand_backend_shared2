package uploads

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	fp "path/filepath"
	"stretches-common-api/app"
	business "stretches-common-api/business"
	config "stretches-common-api/config"
	ut "stretches-common-api/utils"
	"strings"

	"cloud.google.com/go/storage"
)

func StoreGCS1(appConfig config.AppConfig, gcpClient *storage.Client, absolutefilename string, targetname string, dbKey string) error {
	ctx := context.Background()
	path := fp.Join("/gcpconfig", appConfig.Google.CLOUD_SERVICEACCOUNT_JSONPATH)
	//	path := appConfig.GetPath( appConfig.Google.CLOUD_SERVICEACCOUNT_JSONPATH)
	fmt.Printf("------------------------\n")
	fmt.Printf("StoreGCS1 %v\n", path)
	fmt.Printf("GOOGLE CLIENT %v\n", path)

	filepath := absolutefilename
	fmt.Printf("StoreGCS1 write file %v\n", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Printf("StoreGCS1 writen file %v\n", filepath)

	bucketname := business.GetBucketname(appConfig, dbKey)
	fmt.Printf("StoreGCS1 BUCKET %v\n", bucketname)
	fmt.Printf("StoreGCS1 TARGET %v\n", targetname)
	wc := gcpClient.Bucket(bucketname).Object(targetname).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	fmt.Printf("StoreGCS1 uploaded %v\n", filepath)

	return nil
}
func UploadFromFormToGCP(r app.RouteContext, formkey string, extension string, mime string, tmpfilename string, targetname string, dbKey string) bool {

	success, absolutetmpfilepath, _, _ := app.WriteFormImageToFile(r, formkey, tmpfilename)
	if !success {
		app.SetAndSaveBadRequest(r, "Cannot create temp file ", errors.New(""), "ME00429")
		return false
	}
	//h := r.AppCtx
	//GET FILE HEADER
	/*	fileHeader, err := c.FormFile(formkey)
		if err != nil {
			app.SetAndSaveBadRequest(r, "Cannot read file header.", err, "ME00429")
			return false
		}
		fmt.Printf("NAME %v\n", fileHeader.Filename)
		fmt.Printf("SIZE %v\n", fileHeader.Size)

		//OPEN FILE
		file, err := fileHeader.Open()
		if err != nil {
			app.SetAndSaveBadRequest(r, "Cannot read file.", err, "ME00429")
			return false
		}

		//CONTENT TO STRING
		byteContainer, err := ioutil.ReadAll(file) // why the long names though?
		fmt.Printf("UploadFromFormToGCP READ SIZE:%d", len(byteContainer))
		contents := string(byteContainer)
	*/
	//FILTER
	ext, _ := ut.FilterStringKeepAlphaAndDot(extension)
	if strings.Contains(mime, "image") {
		//		absolutetmpfilepath=
		successResize, errResize := ut.ResizeImage(absolutetmpfilepath, absolutetmpfilepath, 256)
		if !successResize {
			app.SetAndSaveBadRequest(r, "Error resizing image file", errResize, "ME00429")
			return false

		}
	}
	//WRITE TEMP FILE
	fmt.Printf("UploadFromFormToGCP EXT %v\n", ext)
	fmt.Printf("UploadFromFormToGCP MIME %v\n", mime)
	//	fmt.Printf("Image %v\n", req.Image)
	fmt.Println(" * UploadFromFormToGCP WRITE FILE")
	fmt.Println("UploadFromFormToGCPWRITE FILE OK", absolutetmpfilepath)

	//UPLOAD FILE
	config := r.GetConfigR()
	if config != nil {
		err6 := StoreGCS1(*config, (r.AppCtx).GetGCP(), absolutetmpfilepath, targetname, dbKey)
		if err6 != nil {
			app.SetAndSaveBadRequest(r, "Error uploading file", err6, "ME00429")
			return false
		}
	}
	defer os.Remove(absolutetmpfilepath)
	return true
}
