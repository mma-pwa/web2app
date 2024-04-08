package client

import (
	"errors"
	"fmt"
	"mime/multipart"
	"web2app/global"
	"web2app/model"
	"web2app/utils"
	"web2app/utils/upload"
)

type UpgradeService struct {
}

func NewUpgradeService() *UpgradeService {

	return &UpgradeService{}
}

func (s *UpgradeService) UploadUpgradeFile(header *multipart.FileHeader, id, imgType, screenType string) (err error, imgInfo *model.ImgInfo) {
	upload := &upload.Uploads{}
	filePath, _, uploadErr := upload.UploadFile(header, global.GVA_CONFIG.Local.UploadPath+"/"+id)
	if uploadErr != nil {
		return uploadErr, nil
	}
	if !utils.IsPNGImage(filePath) && !utils.IsJpgImage(filePath) {
		//if !utils.IsJpgImage(filePath) {
		return errors.New("必须是Png文件或者jpg"), nil
		//}
	}

	imgInfo = new(model.ImgInfo)
	dstName := utils.RandCreator(12)
	imgPathFile := global.GVA_CONFIG.Local.UploadPath + "/" + id + "/"
	if utils.IsPNGImage(filePath) {
		if imgType == global.ICON_IMG {
			scaleImgFile := imgPathFile + dstName + global.BIG_ICON + ".png"
			if err = utils.ImgScale(filePath, scaleImgFile, 512, 512); err != nil {
				return err, nil
			}
			//imgInfo.Url = "/img/" + id + "/" + dstName + global.BIG_ICON + ".png"

			scaleImgFile = imgPathFile + dstName + global.SMALL_ICON + ".png"
			if err = utils.ImgScale(filePath, scaleImgFile, 192, 192); err != nil {
				return err, nil
			}
			imgInfo.Url = "/img/" + id + "/" + dstName + ".png"
			utils.CopyFile(scaleImgFile, imgPathFile+dstName+".png")
		} else if imgType == global.APP_IMG {
			if screenType == global.Vertical_screen {
				scaleImgFile := imgPathFile + dstName + "192" + ".png"
				if err = utils.ImgScale(filePath, scaleImgFile, 526, 296); err != nil {
					return err, nil
				}
				imgInfo.Url = "/img/" + id + "/" + dstName + "526x296" + ".png"
			} else if screenType == global.Landscape_screen {
				scaleImgFile := imgPathFile + dstName + "192" + ".png"
				if err = utils.ImgScale(filePath, scaleImgFile, 144, 266); err != nil {
					return err, nil
				}
				imgInfo.Url = "/img/" + id + "/" + dstName + "144x266" + ".png"
			}
		}
	} else if utils.IsJpgImage(filePath) {
		if imgType == global.ICON_IMG {
			scaleImgFile := imgPathFile + dstName + global.BIG_ICON + ".jpg"
			if err = utils.ImgJpgScale(filePath, scaleImgFile, 512, 512); err != nil {
				return err, nil
			}
			//imgInfo.Url = "/img/" + id + "/" + dstName + global.BIG_ICON + ".jpg"

			scaleImgFile = imgPathFile + dstName + global.SMALL_ICON + ".jpg"
			if err = utils.ImgJpgScale(filePath, scaleImgFile, 192, 192); err != nil {
				return err, nil
			}
			imgInfo.Url = "/img/" + id + "/" + dstName + ".jpg"
			utils.CopyFile(filePath, imgPathFile+dstName+".jpg")
		} else if imgType == global.APP_IMG {
			if screenType == global.Vertical_screen {
				scaleImgFile := imgPathFile + dstName + "192" + ".jpg"
				if err = utils.ImgJpgScale(filePath, scaleImgFile, 526, 296); err != nil {
					return err, nil
				}
				imgInfo.Url = "/img/" + id + "/" + dstName + "526x296" + ".jpg"
			} else if screenType == global.Landscape_screen {
				scaleImgFile := imgPathFile + dstName + "192" + ".jpg"
				if err = utils.ImgJpgScale(filePath, scaleImgFile, 144, 266); err != nil {
					return err, nil
				}
				imgInfo.Url = "/img/" + id + "/" + dstName + "144x266" + ".jpg"
			}
		}
	}
	fmt.Println("----------", global.GVA_CONFIG.Local.ImgUrl)
	imgInfo.Url = global.GVA_CONFIG.Local.ImgUrl + imgInfo.Url
	return nil, imgInfo
}
