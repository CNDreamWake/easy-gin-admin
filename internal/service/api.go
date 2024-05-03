package service

import (
	"gorm.io/gorm"
	"server/internal/vo"
	"server/model"
	"server/pkg/db"
)

type IApi interface {
	Test() error
}

var Api IApi

func NewImplApi() IApi {
	return &implApi{
		db: db.GetDb(),
	}
}

type implApi struct {
	db *gorm.DB
}

func (i *implApi) Test() error {
	if err := i.db.Create(&model.T{}).Error; err != nil {
		return vo.NewErr().Msg("新增测试数据失败")
	}
	return nil
}
