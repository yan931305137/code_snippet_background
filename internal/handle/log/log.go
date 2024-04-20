package handler

import (
	"net/http"

	service "code_snippet/internal/services/log"
	"code_snippet/internal/types"
	"github.com/gin-gonic/gin"
)

// LogHandler 日志处理器
var Log = new(logHandler)

type logHandler struct{}

// LogList 获取日志列表
func (c *logHandler) LogList(ctx *gin.Context) {
	var page types.Page
	// 绑定 URI 中的分页参数
	if err := ctx.ShouldBindUri(&page); err != nil {
		// 如果绑定失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 创建日志服务实例
	logService := service.NewLogService()
	// 获取日志列表
	logList, total, err := logService.GetLogList(page.Size, page.Num)
	if err != nil {
		// 如果获取失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回成功的结果，包括日志列表数据
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "获取日志列表信息成功",
		Data: gin.H{
			"logList": logList,
			"total":   total,
		},
	})
}

// LogAdd 添加日志信息
func (c *logHandler) LogAdd(ctx *gin.Context) {
	var logInfo types.LogInfo
	// 绑定请求中的日志信息
	if err := ctx.ShouldBind(&logInfo); err != nil {
		// 如果绑定失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 创建日志服务实例
	logService := service.NewLogService()
	// 添加日志信息
	if err := logService.CreateLog(logInfo); err != nil {
		// 如果添加失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回成功的结果
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "添加日志信息成功",
	})
}

// LogDelete 删除日志信息
func (c *logHandler) LogDelete(ctx *gin.Context) {
	var logInfo types.LogInfo
	// 绑定请求中的日志信息
	if err := ctx.ShouldBind(&logInfo); err != nil {
		// 如果绑定失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 创建日志服务实例
	logService := service.NewLogService()
	// 删除日志信息
	if err := logService.DeleteLog(logInfo.Id); err != nil {
		// 如果删除失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回成功的结果
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "删除日志信息成功",
	})
}

// LogAlter 修改日志信息
func (c *logHandler) LogAlter(ctx *gin.Context) {
	var logInfo types.LogInfo
	// 绑定请求中的日志信息
	if err := ctx.ShouldBind(&logInfo); err != nil {
		// 如果绑定失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 创建日志服务实例
	logService := service.NewLogService()
	// 修改日志信息
	if err := logService.UpdateLog(logInfo); err != nil {
		// 如果修改失败，返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回成功的结果
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "修改日志信息成功",
	})
}
