package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
)

// CodeService 表示日志管理服务。
type CodeService struct{}

// NewCodeService 创建一个新的 CodeService 实例。
func NewCodeService() *CodeService {
	return &CodeService{}
}

func (s *CodeService) PostCode(req types.CodeReq, userid int) error {
	var codeInfo model.Code
	codeInfo.UserID = userid
	codeInfo.Content = req.Content
	codeInfo.Category = req.Category
	codeInfo.Description = req.Description
	codeInfo.Title = req.Title
	codeInfo.Tags = req.Tags
	//codeInfo.ExpireTime= req.ExpireTime
	codeInfo.Authority = req.Authority
	codeInfo.CodePassword = req.CodePassword
	if _, err := utils.XormDb.Table("code").Insert(&codeInfo); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *CodeService) GetMyCode(userid int) (*[]model.Code, error) {
	var code []model.Code
	if err := utils.XormDb.Table("code").Where("user_id = ?", userid).Find(&code); err != nil {
		return nil, err
	} else {
		return &code, nil
	}
}
