package providers

import (
	"context"
	"fmt"
	"hulk/go-webservice/common"
	"hulk/go-webservice/infrastructure/config"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

func UploadAzBlob(file *multipart.FileHeader) (string, error) {
	config := config.AppConfig()
	url := fmt.Sprintf("https://%s.blob.core.windows.net", config.AzureStorageAccount)
	containerName := config.AzureContainerName
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return "", err
	}

	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		return "", err
	}

	upFile, err := os.Open(file.Filename)
	if err != nil {
		return "", err
	}

	defer upFile.Close()

	var fileSize int64 = file.Size
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	fileKey := strconv.Itoa(int(common.NowMinisecond())) + path.Ext(file.Filename)

	contentType := "image/jpg"
	option := azblob.UploadBufferOptions{
		HTTPHeaders: &blob.HTTPHeaders{
			BlobContentType: &contentType, //  Add any needed headers here
		},
	}

	_, err = client.UploadBuffer(ctx, containerName, fileKey, fileBuffer, &option)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", url, containerName, fileKey), nil
}
