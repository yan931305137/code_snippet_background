package service

import (
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
)

// TopService 表示日志管理服务。
type TopService struct{}

// NewTopService 创建一个新的 TopService 实例。
func NewTopService() *TopService {
	return &TopService{}
}

func (s TopService) GetTopSlider() (*[]types.Code, error) {
	var code []types.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 0).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopHop() (*[]types.Code, error) {
	var code []types.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 1).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopNew() (*[]types.Code, error) {
	var code []types.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 2).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}

func (s TopService) GetTopFocus() (*[]types.Code, error) {
	var code []types.Code
	if err := utils.XormDb.Table("top").Where("top.type = ?", 3).Join("INNER", "code", "code.id = top.code_id").Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}
