package handler

import (
	"fmt"
	"github.com/drakenchef/RIP/MyMinio"
	"github.com/drakenchef/RIP/internal/app/utils"
	"github.com/minio/minio-go"
	"mime/multipart"
)

func (h *Handler) createImageInMinio(file *multipart.File, header *multipart.FileHeader) (string, error) {
	objectName := header.Filename
	if errName := utils.GenerateUniqueName(&objectName); errName != nil {
		return "", errName
	}

	if _, err := h.Minio.PutObject("amsflights", objectName, *file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	}); err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", MyMinio.MinioHost, MyMinio.BucketName, objectName), nil
}
