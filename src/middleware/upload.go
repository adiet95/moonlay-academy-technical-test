package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/adiet95/moonlay-academy-technical-test/src/libs"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		const max = 1024 * 1024 // 1MB
		e.Request().Body = http.MaxBytesReader(e.Response().Writer, e.Request().Body, max)
		if err := e.Request().ParseMultipartForm(max); err != nil {
			libs.New("The uploaded file is less than 1MB in size", 400, true)
			e.Error(err)
		}
		// check attribut file yang di upload
		file, fileHeader, err := e.Request().FormFile("file")
		if err != nil {
			libs.New("invalid attribute", 401, true).Send(e)
			e.Error(err)
		}
		defer file.Close()

		//Checking extension
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			e.Error(err)
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			libs.New("Extension file not allowed", 401, true).Send(e)
			e.Error(err)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			e.Error(err)
		}

		// Membuat folder file upload
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			libs.New("error build file location", 401, true).Send(e)
			e.Error(err)
		}

		// Membuat file baru di direktori file upload
		fileName := e.FormValue("file_name")
		temp := fmt.Sprintf("%d-%s-%s", time.Now().UnixNano(), fileName, filepath.Ext(fileHeader.Filename))
		res := fmt.Sprintf("./uploads/%d-%s-%s", time.Now().UnixNano(), fileName, filepath.Ext(fileHeader.Filename))
		dst, err := os.Create(res)
		if err != nil {
			libs.New("error while upload file", 401, true).Send(e)
			e.Error(err)
		}
		var name interface{} = temp

		defer dst.Close()

		// Mengcopy file ke filesistem sesuai direktorinya
		_, err = io.Copy(dst, file)
		if err != nil {
			libs.New("error copy filesystem", 401, true).Send(e)
			e.Error(err)
		}

		e.Set("dir", name)
		return next(e)
	}
}
