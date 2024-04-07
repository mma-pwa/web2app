package client

import (
	"web2app/global"
	"web2app/model"
	"web2app/model/request"
)

type AppUserReportDayService struct {
}

func NewAppUserReportDayService() *AppUserReportDayService {

	return &AppUserReportDayService{}
}
func (s *AppUserReportDayService) GetAppUserReportDayList(info request.AppUserReportDaySearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.AppUserReportDay{})
	var appList []model.AppUserReportDay

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&appList).Error
	return err, appList, total
}
