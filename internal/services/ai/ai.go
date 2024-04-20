package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/utils"
	"fmt"
)

// AiService 表示日志管理服务。
type AiService struct{}

// NewAiService 创建一个新的 AiService 实例。
func NewAiService() *AiService {
	return &AiService{}
}

// PostAiKnow 根据ID检索日志。
func (s *AiService) PostAiKnow(id int) (*model.Ai, error) {
	var aiInfo model.Ai
	_, err := utils.XormDb.Table("ai").ID(id).Get(&aiInfo)
	if err != nil {
		return nil, fmt.Errorf("获取ID为 %d 的日志失败: %w", id, err)
	}
	return &aiInfo, nil
}
