package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model/common/response"
)

type UpgradeApi struct {
}

// UploadUpgradeFile 上传更新包
// @Tags FileBait
// @Summary 上传的更新包
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Success 200 {object}
// @Router /upgrade/upload/img [post]
func (upgradeApi *UpgradeApi) UploadUpgradeFile(c *gin.Context) {
	imgType, _ := c.GetPostForm("imgType")
	screenType, _ := c.GetPostForm("screenType")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithDetailed(err.Error(), "接收文件失败", c)
		return
	}
	id := "xxxxx"
	userId, b := c.Get("userId")
	if b {
		id = userId.(string)
	}
	err, info := upgradeService.UploadUpgradeFile(header, id, imgType, screenType)
	if err != nil {
		response.FailWithDetailed("fail", err.Error(), c)
		return
	}
	response.OkWithDetailed(info, "上传成功", c)
}
