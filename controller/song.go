package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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
	err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+".mp3", src)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	fmt.Println("うまくいった")
	// mp3をHLSに変換
	err = ct.infra.ConvertVideoHLS(c.Request().Context(), songID, "tmep/input.mp3")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// TODO:output をアップロード

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
