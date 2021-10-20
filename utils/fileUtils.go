package utils

import (
	"context"
	"mime/multipart"
	cc "slot/config"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func CreateFileName(filename string) string {
	newFile := strings.Split(filename, ".")
	return newFile[0]
}

var UrlChan = make(chan string)
var ErrChan = make(chan string)

func UploadToCloud(file multipart.File, Filename string) {
	fileName := CreateFileName(Filename)
	var ctx = context.Background()

	cld, _ := cloudinary.NewFromParams(cc.CloudConfig.CloudName, cc.CloudConfig.ApiKey, cc.CloudConfig.ApiSecret)
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: fileName})

	if err != nil {
		ErrChan <- "Failed to upload file."
	}

	UrlChan <- resp.SecureURL
}
