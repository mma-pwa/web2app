package client

import (
	"web2app/global"
	"web2app/model"
)

type DomainService struct {
}

func NewDomainService() *DomainService {

	return &DomainService{}
}

func (s *DomainService) CreateDomain(domain model.Domain) error {

	return global.GVA_DB.Create(&domain).Error
}

func (s *DomainService) GetDomainList() (err error, list interface{}) {

	db := global.GVA_DB.Model(&model.Domain{})

	var domainList []model.Domain
	err = db.Limit(1000).Offset(0).Order("updated_at desc").Find(&domainList).Error
	return err, domainList
}
