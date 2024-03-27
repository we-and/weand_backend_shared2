package utils

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ResizeImage(originPath string, destpath string, width uint) (bool, error) {

	//	originFile, err := os.ReadFile(originPath)

	originFile, err := os.Open(originPath)
	if err != nil {
		return false, err
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(originFile)
	if err != nil {
		return false, err
	}
	originFile.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(width), 0, img, resize.Lanczos3)

	out, err := os.Create(destpath)
	if err != nil {
		return false, err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return true, nil
}
