package client

import (
	"web2app/global"
	"web2app/model"
	"web2app/model/request"
)

type AppUserLoginLogService struct {
}

func NewAppUserLoginLogService() *AppUserLoginLogService {

	return &AppUserLoginLogService{}
}

func (s *AppUserLoginLogService) GetAppUserLoginLogList(info request.AppUserLoginLogSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.AppUserLoginLog{})
	var appList []model.AppUserLoginLog

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&appList).Error
	return err, appList, total
}
