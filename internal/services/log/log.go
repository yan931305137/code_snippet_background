package service

import (
	"code_snippet/internal/model"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
	"fmt"
)

// LogService 表示日志管理服务。
type LogService struct{}

// NewLogService 创建一个新的 LogService 实例。
func NewLogService() *LogService {
	return &LogService{}
}

// GetLog 根据ID检索日志。
func (s *LogService) GetLog(id int) (*model.Log, error) {
	var logInfo model.Log
	_, err := utils.XormDb.Table("log").ID(id).Get(&logInfo)
	if err != nil {
		return nil, fmt.Errorf("获取ID为 %d 的日志失败: %w", id, err)
	}
	return &logInfo, nil
}

// GetLogList 检索具有分页的日志列表。
func (s *LogService) GetLogList(pageSize, pageNum int) ([]*model.Log, int64, error) {
	var logInfoList []*model.Log
	err := utils.XormDb.Table("log").Limit(pageSize, pageSize*(pageNum-1)).Find(&logInfoList)
	if err != nil {
		return nil, 0, fmt.Errorf("获取日志列表失败: %w", err)
	}
	total, err := utils.XormDb.Table("log").Count()
	if err != nil {
		return nil, 0, fmt.Errorf("获取日志总数失败: %w", err)
	}
	return logInfoList, total, nil
}

// CreateLog 创建新的日志。
func (s *LogService) CreateLog(logInfo types.LogInfo) error {
	_, err := utils.XormDb.Table("log").Insert(&logInfo)
	if err != nil {
		return fmt.Errorf("创建日志失败: %w", err)
	}
	return nil
}

// DeleteLog 删除日志。
func (s *LogService) DeleteLog(id int) error {
	_, err := utils.XormDb.Table("log").ID(id).Delete(&model.Log{})
	if err != nil {
		return fmt.Errorf("删除ID为 %d 的日志失败: %w", id, err)
	}
	return nil
}

// UpdateLog 更新日志。
func (s *LogService) UpdateLog(logInfo types.LogInfo) error {
	_, err := utils.XormDb.Table("log").ID(logInfo.Id).Update(&logInfo)
	if err != nil {
		return fmt.Errorf("更新ID为 %d 的日志失败: %w", logInfo.Id, err)
	}
	return nil
}
