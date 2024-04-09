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

func (s *UpgradeService) UploadUpgradeFile(header *multipart.FileHeader, id, imgType, screenType string) (err error, imgInfos []*model.ImgInfo) {
	upload := &upload.Uploads{}
	filePath, _fileName, uploadErr := upload.UploadFile(header, global.GVA_CONFIG.Local.UploadPath+"/"+id)
	fmt.Println(_fileName)
	if uploadErr != nil {
		return uploadErr, nil
	}
	if !utils.IsPNGImage(filePath) && !utils.IsJpgImage(filePath) {
		//if !utils.IsJpgImage(filePath) {
		return errors.New("必须是Png文件或者jpg"), nil
		//}
	}
	imgInfos = make([]*model.ImgInfo, 0)
	//imgInfo = new(model.ImgInfo)
	dstName := utils.RandCreator(12)
	imgPathFile := global.GVA_CONFIG.Local.UploadPath + "/" + id + "/"
	if utils.IsPNGImage(filePath) {
		if imgType == global.ICON_IMG {
			scaleImgFile := imgPathFile + dstName + global.BIG_ICON + ".png"
			if err = utils.ImgScale(filePath, scaleImgFile, 512, 512); err != nil {
				return err, nil
			}
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + dstName + global.BIG_ICON + ".png",
				Width:   512,
				Height:  512,
				ImgType: "png",
			})
			//utils.CopyFile(scaleImgFile, imgPathFile+dstName+".png")
			scaleImgFile = imgPathFile + dstName + global.SMALL_ICON + ".png"
			if err = utils.ImgScale(filePath, scaleImgFile, 192, 192); err != nil {
				return err, nil
			}
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + dstName + global.SMALL_ICON + ".png",
				Width:   192,
				Height:  192,
				ImgType: "png",
			})
		} else if imgType == global.APP_IMG {
			x, y, format := utils.GetImgWH(filePath)
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + _fileName,
				Width:   x,
				Height:  y,
				ImgType: format,
			})
			//if screenType == global.Vertical_screen {
			//	scaleImgFile := imgPathFile + dstName + "526x296" + ".png"
			//	if err = utils.ImgScale(filePath, scaleImgFile, 526, 296); err != nil {
			//		return err, nil
			//	}
			//	imgInfos = append(imgInfos, &model.ImgInfo{
			//		Url:     "/img/" + id + "/" + dstName + "526x296" + ".png",
			//		Width:   526,
			//		Height:  296,
			//		ImgType: "png",
			//	})
			//} else if screenType == global.Landscape_screen {
			//	scaleImgFile := imgPathFile + dstName + "144x266" + ".png"
			//	if err = utils.ImgScale(filePath, scaleImgFile, 144, 266); err != nil {
			//		return err, nil
			//	}
			//	imgInfos = append(imgInfos, &model.ImgInfo{
			//		Url:     "/img/" + id + "/" + dstName + "144x266" + ".png",
			//		Width:   144,
			//		Height:  266,
			//		ImgType: "png",
			//	})
			//}
		}
	} else if utils.IsJpgImage(filePath) {
		if imgType == global.ICON_IMG {
			scaleImgFile := imgPathFile + dstName + global.BIG_ICON + ".jpg"
			if err = utils.ImgJpgScale(filePath, scaleImgFile, 512, 512); err != nil {
				return err, nil
			}
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + dstName + global.BIG_ICON + ".jpg",
				Width:   512,
				Height:  512,
				ImgType: "jpg",
			})

			scaleImgFile = imgPathFile + dstName + global.SMALL_ICON + ".jpg"
			if err = utils.ImgJpgScale(filePath, scaleImgFile, 192, 192); err != nil {
				return err, nil
			}
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + dstName + global.SMALL_ICON + ".jpg",
				Width:   192,
				Height:  192,
				ImgType: "jpg",
			})
		} else if imgType == global.APP_IMG {
			x, y, format := utils.GetImgWH(filePath)
			imgInfos = append(imgInfos, &model.ImgInfo{
				Url:     "/img/" + id + "/" + _fileName,
				Width:   x,
				Height:  y,
				ImgType: format,
			})
		}
	}
	//fmt.Println("----------", global.GVA_CONFIG.Local.ImgUrl)
	for _, imgInfo := range imgInfos {
		imgInfo.Url = global.GVA_CONFIG.Local.ImgUrl + imgInfo.Url
	}

	return nil, imgInfos
}
