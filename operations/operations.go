package operations

import (
    "context"
    "fmt"
    "strings"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

type Operations struct {
    minioClient *minio.Client
}

func NewOperations(endpoint, accessKey, secretKey string) (*Operations, error) {
    // Initialize MinIO client
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false, // Set to true if MinIO server is secure (HTTPS)
    })

    if err != nil {
        return nil, err
    }

    return &Operations{
        minioClient: minioClient,
    }, nil
}

func (o *Operations) ListFiles(bucketName string) error {
    ctx := context.Background()
    objectsCh := o.minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{})
    for object := range objectsCh {
        if object.Err != nil {
            return object.Err
        }
        fmt.Println(object.Key)
    }
    return nil
}

func (o *Operations) ModifyFile(bucketName, objectName, content string) error {
    ctx := context.Background()
    _, err := o.minioClient.PutObject(ctx, bucketName, objectName, strings.NewReader(content), int64(len(content)), minio.PutObjectOptions{})
    return err
}
func (o *Operations) DeleteFile(bucketName, objectName string) error {
	ctx := context.Background()
	err := o.minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	return err
}