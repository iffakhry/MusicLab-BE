package helper

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"strings"

	"musiclab-be/app/config"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	DEFAULT_GCS_LINK string = "https://storage.googleapis.com/altabucket/"
)

type ClientUploader struct {
	storageClient *storage.Client
	projectID     string
	bucketName    string
	uploadPath    string
}

var clientUploader *ClientUploader

func GetStorageClient() *ClientUploader {
	if clientUploader == nil {
		client, err := storage.NewClient(context.Background(), option.WithoutAuthentication())
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		clientUploader = &ClientUploader{
			storageClient: client,
			bucketName:    config.GCP_BUCKET_NAME,
			projectID:     config.GCP_PROJECT_ID,
			uploadPath:    "musiclab/images/",
		}

		return clientUploader
	}
	return clientUploader
}

// ======================================================================
// UPLOAD IMAGE PROGRESS
// ======================================================================
func GetUrlImagesFromGCS(fileData multipart.FileHeader) (string, error) {

	if fileData.Filename != "" && fileData.Size != 0 {
		if fileData.Size > 500000 {
			return "", errors.New("file size max 500kb")
		}
		file, err := fileData.Open()
		if err != nil {
			return "", errors.New("error open fileData")
		}
		// Validasi Type
		tipeNameFile, err := TypeFile(file)
		if err != nil {
			return "", errors.New("file type error only jpg or png file can be upload")
		}
		defer file.Close()

		log.Println("size:", fileData.Filename, file)
		namaFile := GenerateRandomString()
		namaFile = namaFile + tipeNameFile
		fileData.Filename = namaFile
		log.Println(namaFile)
		file2, _ := fileData.Open()
		defer file2.Close()
		// uploadURL, err := UploadToS3(fileData.Filename, file2)
		uploadURL, err := GetStorageClient().UploadFile(file2, fileData.Filename)
		if err != nil {
			return "", errors.New("cannot upload to GCS server error")
		}
		return uploadURL, nil
	}
	return "", nil
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, objectName string) (fileLocation string, err error) {
	ctx := context.Background()

	// Upload an object with storage.Writer.
	wc := c.storageClient.Bucket(c.bucketName).Object(c.uploadPath + objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return DEFAULT_GCS_LINK + wc.Name, nil
}

func (c *ClientUploader) DeleteFile(objectName string) error {
	ctx := context.Background()

	wc := c.storageClient.Bucket(c.bucketName).Object(strings.Replace(objectName, DEFAULT_GCS_LINK, "", 1))
	if err := wc.Delete(ctx); err != nil {
		return err
	}

	return nil
}
