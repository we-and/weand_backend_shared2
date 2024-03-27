package instance

import (
	"context"
	"fmt"
	"log"

	fp "path/filepath"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func (c *CommonInstance) GenerateGCPClient() (*storage.Client, error) {
	fmt.Printf("[%v][GCP] GCP client\n", c.ServerId)
	//create gcp client
	ctx := context.Background()
	path := c.GetPathInDevFolder(fp.Join("gcpconfig", c.AppConfig.Google.CLOUD_SERVICEACCOUNT_JSONPATH))
	fmt.Printf("[%v][GCP] GCP config : %v\n", c.ServerId, path)
	gcpClient, err := storage.NewClient(ctx, option.WithCredentialsFile(path))
	if err != nil {
		fmt.Printf("[%v][GCP] Error creating GCP client %vn", c.ServerId, err)
		log.Fatalf("Error creating GCP client %vn", err)
	} else {
		fmt.Printf("[%v][GCP] Creating GCP client OK\n", c.ServerId)
	}
	return gcpClient, err
}
func (c *CommonInstance) GenerateFirebaseClient() (*firebase.App, error) {
	fmt.Printf("[%v][FIREBASE] Firebase client\n", c.ServerId)

	// Initialize firebase
	path2 := c.GetPathInDevFolder(fp.Join("firebase", fmt.Sprintf("%v%v", c.AppConfig.Firebase.SERVICEACCOUNT_JSONPATH, ".json")))
	fmt.Printf("[%v][FIREBASE] Firebase config : %v\n", c.ServerId, path2)
	opt := option.WithCredentialsFile(path2)
	conf := firebase.Config{
		ProjectID: "acto-24edc",
	}
	firebaseClient, err := firebase.NewApp(context.Background(), &conf, opt)
	if err != nil {
		//        log.Fatalf("error initializing app: %v\n", err)
		log.Fatalf("[%v][FIREBASE] Error creating firebase client %vn", c.ServerId, err)
	} else {
		fmt.Printf("[%v][FIREBASE] Creating firebase client OK\n", c.ServerId)
	}
	return firebaseClient, err
}
