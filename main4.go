package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"time"
)

func main() {
	MinioClient, Err := minio.New("47.108.20.64:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("g5fMoJm1xGdBSbgMrDpw", "XdIi8sUQSqMR0k1XD1VDtpxOc0ECmWpkT6vnXkFI", ""),
		Secure: false,
	})
	fmt.Println(Err)
	url, err := MinioClient.PresignedGetObject(context.Background(), "monitordownloadurl", "MonitorZip.zip", time.Second*60*60*24*4, make(url.Values))
	fmt.Println(url, err)
}
