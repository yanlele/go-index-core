package api

import (
	"gin-example/pkg/e"
	"gin-example/pkg/logging"
	"gin-example/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(context *gin.Context) {
	code := e.SUCCESS
	resData := make(map[string]interface{})
	file, image, err := context.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"data":    resData,
		})
	}
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := context.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				resData["image_url"] = upload.GetImageFullUrl(imageName)
				resData["image_save_url"] = savePath + imageName
			}
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    resData,
	})
}
