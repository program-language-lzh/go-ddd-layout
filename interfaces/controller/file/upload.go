package file

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/domain/file/entity"
	"server/infrastructure/common/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctl *EndpointCtl) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	tempZipPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempZipPath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	defer os.Remove(tempZipPath)

	extractDir := filepath.Join(os.TempDir(), "extracted")
	if err := os.MkdirAll(extractDir, 0755); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	defer os.RemoveAll(extractDir)
	if err := unzip(tempZipPath, extractDir); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}
	// Traverse the extracted files and save file info to database
	err = filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// hash := md5.Sum(fileData)
			// fileMD5 := hex.EncodeToString(hash[:])
			fileType := http.DetectContentType(fileData)

			fileInfo := &entity.File{
				Name: info.Name(),
				// FileMD5:     fileMD5,
				Type: fileType,
				// StoragePath: path,
			}

			_, err = ctl.Srv.CreateFile(fileInfo)
			if err != nil {
				log.Printf("Failed to save file info: %v", err)
			}
		}
		return nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusOK, response.Ok())

}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
