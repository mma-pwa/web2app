package client

import (
	"time"
	"web2app/global"
	"web2app/model"
	"web2app/model/request"
)

type AppCustomURLService struct {
}

func NewAppCustomURLService() *AppCustomURLService {

	return &AppCustomURLService{}
}

func (s *AppCustomURLService) CreateAppCustomURL(appCustom model.AppCustomURL) error {
	appCustom.CreatedAt = time.Now()
	appCustom.UpdatedAt = time.Now()
	return global.GVA_DB.Create(&appCustom).Error
}
func (s *AppCustomURLService) DeleteAppCustomURL(appCustom model.AppCustomURL) error {

	return global.GVA_DB.Delete(&appCustom, appCustom.ID).Error
}

func (s *AppCustomURLService) UpdateAppCustomURL(appCustom model.AppCustomURL) error {

	return global.GVA_DB.Save(&appCustom).Error
}

func (s *AppCustomURLService) GetAppCustomList(info request.AppCustomURLSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.AppCustomURL{})
	var appList []model.AppCustomURLRes

	err = db.Count(&total).Error
	if info.UserId != "" {
		db = db.Where("a.user_id = ?", info.UserId)
	}
	err = db.Table(model.TableNameAppCustomURL).Select("app_custom_url.id,app_custom_url.custom_url,app_custom_url.status,app_custom_url.created_at,"+
		"app_custom_url.updated_at, app_custom_url.app_id,a.name").
		Joins("left join app as a on a.id=app_custom_url.app_id ", info.UserId).Scan(&appList).Limit(limit).Offset(offset).Error
	//err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&appList).Error
	return err, appList, total
}
