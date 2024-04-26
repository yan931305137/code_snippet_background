package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/utils"
)

// TopService 表示日志管理服务。
type TopService struct{}

// NewTopService 创建一个新的 TopService 实例。
func NewTopService() *TopService {
	return &TopService{}
}

func (s TopService) GetTopSlider() (*[]model.Code, error) {
	var code []model.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 0).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopHop() (*[]model.Code, error) {
	var code []model.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 1).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopNew() (*[]model.Code, error) {
	var code []model.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 2).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopFocus() (*[]model.Code, error) {
	var code []model.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 3).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}
