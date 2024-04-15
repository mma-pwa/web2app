package client

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
	"web2app/global"
	"web2app/model"
	"web2app/model/request"
	"web2app/utils"
)

type AppService struct {
}

func NewAppService() *AppService {

	return &AppService{}
}

func (s *AppService) CreateAPP(app model.App) error {
	icon, _ := json.Marshal(app.IconImg)
	appScreen, _ := json.Marshal(app.ScreenImg)
	app.Icons = string(icon)
	app.AppScreen = string(appScreen)
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	app.AppURL = "https://" + app.AppURL
	app.ID = utils.MD5V([]byte(uuid.New().String()))
	return global.GVA_DB.Create(app).Error
}

func (s *AppService) DeleteApp(app model.App) error {

	return global.GVA_DB.Delete(app).Error
}

func (s *AppService) UpdateApp(app model.App) error {
	icon, _ := json.Marshal(app.IconImg)
	appScreen, _ := json.Marshal(app.ScreenImg)
	app.Icons = string(icon)
	app.AppScreen = string(appScreen)
	app.UpdatedAt = time.Now()
	return global.GVA_DB.Save(app).Error
}

func (s *AppService) GetAppInfo(id string) (err error, app *model.App) {

	err = global.GVA_DB.Where("id = ?", id).First(&app).Error
	return
}
func (s *AppService) GetAppInfoList(info request.AppSearch) (err error, list interface{}, total int64) {
	if info.UserID == "" {
		return errors.New("unauthorized user"), nil, 0
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.App{})
	if info.UserID != "" {
		db = db.Where("user_id = ?", info.UserID)
	}
	if info.Status != nil {
		db = db.Where("status= ?", info.Status)
	}
	appList := make([]*model.App, 0)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&appList).Error
	for _, v := range appList {
		imgInfo := make([]model.ImgInfo, 0)
		json.Unmarshal([]byte(v.Icons), &imgInfo)
		v.IconImg = imgInfo
		screenImg := make([]model.ImgInfo, 0)
		json.Unmarshal([]byte(v.AppScreen), &screenImg)
		v.ScreenImg = screenImg
	}
	return err, appList, total
}

func (s *AppService) GetAppList(userId string) (err error, list interface{}) {
	if userId == "" {
		return errors.New("unauthorized user"), nil
	}
	db := global.GVA_DB.Model(&model.App{})
	if userId != "" {
		db = db.Where("user_id = ?", userId)
	}
	var appList []model.AppRes
	err = db.Limit(1000).Offset(0).Order("updated_at desc").Find(&appList).Error
	return err, appList
}
