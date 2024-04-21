package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/utils"
)

// AiService 表示日志管理服务。
type AiService struct{}

// NewAiService 创建一个新的 AiService 实例。
func NewAiService() *AiService {
	return &AiService{}
}

// GetAiKnow  根据ID获取对话。
func (s *AiService) GetAiKnow(userid int) (*[]model.Ai, error) {
	var aiInfo []model.Ai
	if err := utils.XormDb.Table("ai").Where("user_id = ?", userid).Find(&aiInfo); err != nil {
		return nil, err
	} else {
		return &aiInfo, nil
	}
}

// PostAiKnow  根据ID更新对话。
func (s *AiService) PostAiKnow(aiInfo model.Ai) error {
	if aiInfo.Id != -1 {
		if _, err := utils.XormDb.Table(`ai`).ID(aiInfo.Id).Where("user_id = ? ", aiInfo.UserId).Update(&aiInfo); err != nil {
			return err
		}
		return nil
	} else {
		var p int
		aiInfo.Id = p
		if _, err := utils.XormDb.Table(`ai`).Insert(&aiInfo); err != nil {
			return err
		}
		return nil
	}
}

func (s *AiService) GetAiTalk(id int, userid int) (string, error) {
	var aiInfo model.Ai
	if _, err := utils.XormDb.Table(`ai`).ID(id).Where("user_id = ? ", userid).Get(&aiInfo); err != nil {
		return "", err
	}
	return aiInfo.Content, nil
}

func (s *AiService) DeleteListAiKnow(userid int) error {
	var aiInfo model.Ai
	if _, err := utils.XormDb.Table(`ai`).Where("user_id = ? ", userid).Delete(&aiInfo); err != nil {
		return err
	}
	return nil
}
