package ominio

import (
	"DIMSMonitorPlat/config"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"time"
)

var MinioClient *minio.Client
var Err error

func NewMinioClient() error {
	MinioClient, Err = minio.New(config.Conf.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.Minio.AccessKeyID, config.Conf.Minio.SecretAccessKey, ""),
		Secure: config.Conf.Minio.UseSSL,
	})
	if Err != nil {
		return Err
	}
	return nil
}

func Download(ctx context.Context, bucketName string, objectName string, filePath string) error {
	err := MinioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	return err
}
func DownloadBinary(ctx context.Context, bucketName string, objectName string) ([]byte, error) {
	object, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	stat, err := object.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size == int64(len(bs)) {
		return bs, nil
	} else {
		return nil, errors.New("download file error")
	}
}
func UploadBinary(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64) error {
	_, err := MinioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{})
	return err
}
func UploadFile(ctx context.Context, bucketName string, objectName string, filePath string) error {
	_, err := MinioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	return err
}
func GetObjectUrl(ctx context.Context, bucketName string, objectName string, expires time.Duration) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment")
	url, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, expires, reqParams)

	return url.String(), err
}
