package controller

import (
	"mime/multipart"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	UserID string                `form:"user_id" validate:"required"`
	Name   string                `form:"name" validate:"required"`
	icon   *multipart.FileHeader `form:"icon"`
}

// func (ct *Controller) PostUser(c echo.Context) error {
// 	var req CreateUserRequest
// 	if err := c.Bind(&req); err != nil {
// 		return err
// 	}
// 	var uuid pgtype.UUID
// 	err := uuid.Scan(req.UserID)

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return echo.ErrBadRequest
// 	}
// 	var image *multipart.FileHeader

// 	imageFiles, ok := form.File["image"]
// 	if !ok {
// 		image = nil
// 	} else {
// 		image = imageFiles[0]
// 	}
// 	if image != nil {
// 		src, err := image.Open()
// 		if err != nil {
// 			c.Logger().Error(err)
// 			return err
// 		}
// 		defer src.Close()
// 		data, err := io.ReadAll(src)
// 		err = ct.infra.UploadFile(c.Request().Context(), "cassette-songs", fmt.Sprintf("%s.%s", uuid, image.Filename), data)
// 		if err != nil {
// 			c.Logger().Error(err)
// 			return err
// 		}
// 	}

// 	return c.JSON(http.StatusOK, "Post User")
// }

// e.GET("/users/:id", getUser)
func (ct *Controller) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid UUID format")
	}
	res, err := ct.db.GetUser(c.Request().Context(), uuid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user")
	}

	return c.JSON(http.StatusOK, res)
}
