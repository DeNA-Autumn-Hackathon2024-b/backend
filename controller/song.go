package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ct *Controller) UploadSong(c echo.Context) error {
	// リクエストから音楽ファイルを取得
	file, err := c.FormFile("song")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer src.Close()

	// フォームデータから追加情報を取得
	var cassetteID pgtype.UUID
	var userID pgtype.UUID
	err = cassetteID.Scan(c.FormValue("cassette_id"))
	err = userID.Scan(c.FormValue("user_id"))
	songNumber, _ := strconv.Atoi(c.FormValue("song_number"))
	songTime, _ := strconv.Atoi(c.FormValue("song_time"))
	name := c.FormValue("name")
	// uploadUser := c.FormValue("upload_user")

	// S3にアップロード
	songID := uuid.New().String()

	res, err := ct.db.PostSong(c.Request().Context(), sqlc.PostSongParams{
		CassetteID: cassetteID,
		UserID:     userID,
		SongNumber: int32(songNumber),
		SongTime:   pgtype.Int4{Int32: int32(songTime), Valid: true},
		Name:       name,
		Url:        os.Getenv("S3_URL") + "/" + songID + "/" + songID + ".m3u8",
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create cassette")
	}

	err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+"/original.mp3", src)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// S3のURL生成
	s3URL := os.Getenv("S3_URL")
	key := s3URL + "/" + songID + "/original.mp3"
	// key = "https://cassette-songs.s3.ap-southeast-2.amazonaws.com/34ccd82a-6020-4bd5-b197-098f62f6bcc7.mp3"
	err = ct.infra.ConvertVideoHLS(c.Request().Context(), songID, key)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	// HLSファイルのアップロード
	dirPath := "output"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), songID) {
			fileD, err := os.Open(filepath.Join("output", file.Name()))
			if err != nil {
				c.Logger().Error(fmt.Errorf("1%v", err))
				return fmt.Errorf("1%v", err)
			}
			defer fileD.Close()

			err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", songID+"/"+file.Name(), fileD)
			if err != nil {
				c.Logger().Error(fmt.Errorf("1%v", err))
				return fmt.Errorf("1%v", err)
			}

			fmt.Println(res)
		}
	}

	// 不要な一時ファイルを削除
	err = os.Remove(filepath.Join(dirPath, songID+".m3u8"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	for i := 0; ; i++ {
		tsFile := filepath.Join(dirPath, fmt.Sprintf("%s%05d.ts", songID, i))
		if _, err := os.Stat(tsFile); os.IsNotExist(err) {
			break
		}
		err = os.Remove(tsFile)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
	}

	fmt.Println(cassetteID, userID, songNumber, songTime, name)

	m3u8URL := os.Getenv("S3_URL") + "/" + songID + "/" + songID + ".m3u8"
	// レスポンスのJSONを構築
	response := map[string]string{
		"url": m3u8URL,
	}

	return c.JSON(http.StatusOK, response)
}
