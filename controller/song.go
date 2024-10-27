package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ct *Controller) UploadSong(c echo.Context) error {
	file, err := c.FormFile("song")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	err = os.MkdirAll("output", 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// TODO:mp3をs3にアップロード
	songID := uuid.New().String()
	err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+"/original.mp3", src)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	s3URL := os.Getenv("S3_URL")
	// mp3をHLSに変換
	err = ct.infra.ConvertVideoHLS(c.Request().Context(), songID, s3URL+"/"+songID+"/original.mp3")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// TODO:output をアップロード
	err = filepath.Walk("output/"+songID, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		// 対象のファイルかどうかを確認
		if strings.HasPrefix(filepath.Base(path), songID) && (strings.HasSuffix(path, ".m3u8") || strings.HasSuffix(path, ".ts")) {
			// TODO: 失敗した時にtsファイルを削除できるように修正する
			file, err := os.Open(path)
			if err != nil {
				c.Logger().Error(err)
				return err
			}

			defer func() error {
				err = os.Remove(path)
				if err != nil {
					c.Logger().Error(err)
					return err
				}
				return nil
			}()

			err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+file.Name(), file)
			if err != nil {
				c.Logger().Error(err)
				return err
			}

		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to remove output files: %w", err)
	}

	// outputのやつ消す
	outputDir := "output"

	// songIDに関連するファイルを削除
	err = os.Remove(filepath.Join(outputDir, songID+".m3u8"))
	if err != nil {
		fmt.Println("Error deleting m3u8 file:", err)
		c.Logger().Error(err)
		return err
	}

	// .tsファイルを削除
	for i := 0; ; i++ {
		tsFile := filepath.Join(outputDir, fmt.Sprintf("%s%05d.ts", songID, i))
		if _, err := os.Stat(tsFile); os.IsNotExist(err) {
			break // ファイルが存在しない場合はループを終了
		}
		err = os.Remove(tsFile)
		if err != nil {
			fmt.Println("Error deleting ts file:", err)
		}
	}

	defer src.Close()
	return c.String(http.StatusOK, "id")
}
