package minio

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	miniosdk "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/quikzens/rest-api-boilerplate/config"
	"github.com/quikzens/rest-api-boilerplate/helper"
)

// Minio Configurtion
var (
	endpoint      = config.MinioEndpoint
	accessKeyID   = config.MinioAccessKeyId
	accessKeyPass = config.MinioAccessKeyPass
	bucketName    = config.MinioBucketName
	useSSL        = false
)

// Minio Folder Names
var (
	UserFolder = "user"
)

// Keep Minio global instance for usage across the app
var MinioClient = connectMinio()

// Initialize minio client object instance
func connectMinio() *miniosdk.Client {
	minioClient, err := miniosdk.New(endpoint, &miniosdk.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, accessKeyPass, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	return minioClient
}

// Create new Minio Object
func CreateObject(object *multipart.FileHeader, folder string) (string, error) {
	file, err := object.Open()
	if err != nil {
		return "", err
	}

	objectName := fmt.Sprintf("%v%v", helper.RandomString(20), filepath.Ext(object.Filename))
	objectFullPath := fmt.Sprintf("%v/%v", folder, objectName)

	// upload object to minio
	_, err = MinioClient.PutObject(context.TODO(), bucketName, objectFullPath, file, object.Size, miniosdk.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return objectName, nil
}

// Remove existing Minio Object
func RemoveObject(folder string, deletedObject string) error {
	deletedObjectFullPath := fmt.Sprintf("%v/%v", folder, deletedObject)

	if deletedObject != "" {
		err := MinioClient.RemoveObject(context.TODO(), bucketName, deletedObjectFullPath, miniosdk.RemoveObjectOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// Override existing Minio Object
func UpdateObject(folder string, newObject *multipart.FileHeader, replacedObject string) (string, error) {
	newObjectName, err := CreateObject(newObject, folder)
	if err != nil {
		return "", err
	}

	err = RemoveObject(folder, replacedObject)
	if err != nil {
		return "", err
	}

	return newObjectName, nil
}
