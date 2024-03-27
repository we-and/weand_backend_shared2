package path

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func GetAssetPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("ERROR cannot read current directory")
	}
	if runtime.GOOS == "windows" {
		path = fmt.Sprintf("%v/assets", path)
	} else {
		path = "/go/assets"
	}
	return path
}

