package utils

import (
	"context"
	"mime/multipart"
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

	cld, _ := cloudinary.NewFromParams("sandydev99", "917665392796572", "J6m152XVs7TyyfxJYn9oIHjPiGc")
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: fileName})

	if err != nil {
		ErrChan <- "Failed to upload file."
	}

	UrlChan <- resp.SecureURL
}
