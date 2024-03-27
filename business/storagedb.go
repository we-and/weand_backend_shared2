package business

import (
	"fmt"
	config "stretches-common-api/config"
)

func GetBucketname(conf config.AppConfig, dbKey string) string {
	switch dbKey {
	case "live":
		return conf.Google.CLOUD_STORAGE_BUCKET_UPLOADS.Prod
	case "test":
		return conf.Google.CLOUD_STORAGE_BUCKET_UPLOADS.Stag
	case "local":
		return conf.Google.CLOUD_STORAGE_BUCKET_UPLOADS.Dev
	default:
		return "unimplemented-bucketname"
	}
}
func GetFinalUrl(conf *config.AppConfig, dbKey string, targetname string) string {
	if conf == nil {
		return ""
	}
	return fmt.Sprintf("https://storage.googleapis.com/%v/%v", GetBucketname(*conf, dbKey), targetname)
}
