package handler

import (
	service "code_snippet/internal/services/top"
	"code_snippet/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TopHandler top处理器
var Top = new(topHandler)

type topHandler struct{}

// TopList 获取日志列表
func (c *topHandler) GetTopSlider(ctx *gin.Context) {
	if codeInfo, err := service.NewTopService().GetTopSlider(); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "获取代码段成功",
			Data: &codeInfo,
		})
	}
}

func (c *topHandler) GetTopHop(ctx *gin.Context) {
	if codeInfo, err := service.NewTopService().GetTopHop(); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "获取代码段成功",
			Data: &codeInfo,
		})
	}
}

func (c *topHandler) GetTopNew(ctx *gin.Context) {
	if codeInfo, err := service.NewTopService().GetTopNew(); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "获取代码段成功",
			Data: &codeInfo,
		})
	}
}

func (c *topHandler) GetTopFocus(ctx *gin.Context) {
	if codeInfo, err := service.NewTopService().GetTopFocus(); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "获取代码段成功",
			Data: &codeInfo,
		})
	}
}
