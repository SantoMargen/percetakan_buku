package helpers

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "time"
	"siap_app/internal/app/entity/upload"
)

var uploadDir = "/Users/hisbikal/Documents/GitHub/sii_terbitan/uploads"

func UploadFileHandler(ctx context.Context, c *http.Request, path string) (resUpload []upload.RequestUpload, err error) {

	var (
		getFileUpload     upload.RequestUpload
		resultUploadsFile []upload.RequestUpload
	)

	pathDetail := uploadDir + "/" + path

	if err := ensureUploadDirExists(pathDetail); err != nil {
		return resultUploadsFile, errors.New("ERROR CREATING UPLOAD DIRECTORY")
	}

	files := c.MultipartForm.File["attachment"]

	for _, file := range files {
		src, err := file.Open()

		if err != nil {
			return resultUploadsFile, errors.New("ERROR RETRIEVE THE FILE")
		}
		defer src.Close()

		fileTypeSplit := strings.Split(file.Filename, ".")
		fileExtension := fileTypeSplit[len(fileTypeSplit)-1]

		if !isFileExtensionAllowed(fileExtension) {
			return resultUploadsFile, errors.New(file.Filename + " THE PROVIDED FILE FORMAT IS NOT ALLOWED. PLEASE UPLOAD AN ALLOWED FILE")
		}

		const MAX_UPLOAD_SIZE = 30 << 20

		errParse := c.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if errParse != nil {
			return resultUploadsFile, errors.New("FILE SIZE NOT ALLOWED")
		}

		randomTime := time.Now().Format("20060102150405")
		uniqueName := randomTime + "_" + file.Filename
		filePath := filepath.Join(pathDetail, uniqueName)

		dst, err := os.Create(filePath)

		if err != nil {
			return resultUploadsFile, errors.New("FAILED SAVE FILE")
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			return resultUploadsFile, errors.New("ERROR COPYING FILE CONTENT")
		}

		getFileUpload = upload.RequestUpload{
			IDFile:   randomTime,
			Filename: uniqueName,
			Filetype: fileExtension,
			Path:     pathDetail,
		}

		resultUploadsFile = append(resultUploadsFile, getFileUpload)
	}

	return resultUploadsFile, nil
}

func isFileExtensionAllowed(fileExtension string) bool {

	allowedExtensions := map[string]bool{
		"jpeg": true,
		"jpg":  true,
		"png":  true,
		"pdf":  true,
		"docx": true,
		"doc":  true,
		"xlsx": true,
		"csv":  true,
	}

	return allowedExtensions[fileExtension]
}

func ensureUploadDirExists(pathDetail string) error {

	if _, err := os.Stat(pathDetail); os.IsNotExist(err) {
		err := os.Mkdir(pathDetail, 0755)
		if err != nil {
			return errors.New("ERROR CREATING DIRECTORY FILE UPLOAD")
		} else {
			return nil
		}

	} else {
		return err
	}

}
