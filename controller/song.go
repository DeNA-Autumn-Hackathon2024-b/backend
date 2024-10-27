package controller

import (
	"fmt"
	"log"
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
	key := s3URL + "/" + songID + "/original.mp3"
	// key = "https://cassette-songs.s3.ap-southeast-2.amazonaws.com/34ccd82a-6020-4bd5-b197-098f62f6bcc7.mp3"
	err = ct.infra.ConvertVideoHLS(c.Request().Context(), songID, key)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	dirPath := "output"

	// ディレクトリの内容を取得
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// ファイルを一つずつ表示（prefix が "song" のもののみ）
	for _, file := range files {
		if strings.HasPrefix(file.Name(), songID) {
			fileD, err := os.Open("output/" + file.Name())
			if err != nil {
				c.Logger().Error(fmt.Errorf("1%v", err))
				return fmt.Errorf("1%v", err)
			}
			err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+"/"+songID+file.Name(), fileD)
			if err != nil {
				c.Logger().Error(fmt.Errorf("1%v", err))
				return fmt.Errorf("1%v", err)
			}
		}
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

	// ct.db.PostCassette()

	defer src.Close()
	return c.String(http.StatusOK, "id")
}
