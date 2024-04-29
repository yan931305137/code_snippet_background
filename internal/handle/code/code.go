package code

import (
	service "code_snippet/internal/services/code"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Code = new(codeHandler)

type codeHandler struct{}

func (c *codeHandler) PostCode(ctx *gin.Context) {
	var req types.CodeReq
	// 获取参数并验证
	if err := ctx.ShouldBind(&req); err != nil {
		// 返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			//Msg:  "参数错误",
			Msg: err.Error(),
		})
		return
	} else {
		token := utils.GetToken(ctx)
		if userid, err := utils.GetTokenToRedis(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, types.JsonResult{
				Code: -1,
				Msg:  "请先登录再进行操作!",
			})
			return
		} else {
			if err := service.NewCodeService().PostCode(req, userid); err != nil {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: -1,
					Msg:  err.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: 0,
					Msg:  "创建代码段成功",
				})
			}
		}
	}
}

func (c *codeHandler) GetMyCode(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	if userid, err := utils.GetTokenToRedis(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, types.JsonResult{
			Code: -1,
			Msg:  "请先登录再进行操作!",
		})
		return
	} else {
		if codeInfo, err := service.NewCodeService().GetMyCode(userid); err != nil {
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
}

func (c *codeHandler) GetCodes(ctx *gin.Context) {
	if codeInfo, err := service.NewCodeService().GetCodes(); err != nil {
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

func (c *codeHandler) SearchGetCodes(ctx *gin.Context) {
	var str types.Str
	if err := ctx.ShouldBind(&str); err != nil {
		// 返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			//Msg:  "参数错误",
			Msg: err.Error(),
		})
		return
	} else {
		if codeInfo, err := service.NewCodeService().SearchGetCodes(str.Value); err != nil {
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
}

func (c *codeHandler) PostLike(ctx *gin.Context) {
	var req types.CodeId
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		token := utils.GetToken(ctx)
		if userid, err := utils.GetTokenToRedis(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, types.JsonResult{
				Code: -1,
				Msg:  "请先登录再进行操作!",
			})
			return
		} else {
			if err := service.NewCodeService().PostLike(req.CodeId, userid); err != nil {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: -1,
					Msg:  err.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: 0,
					Msg:  "代码段点赞成功",
				})
			}
		}
	}
}

func (c *codeHandler) PostCollect(ctx *gin.Context) {
	var req types.CodeId
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	} else {
		token := utils.GetToken(ctx)
		if userid, err := utils.GetTokenToRedis(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, types.JsonResult{
				Code: -1,
				Msg:  "请先登录再进行操作!",
			})
			return
		} else {
			if err := service.NewCodeService().PostCollect(req.CodeId, userid); err != nil {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: -1,
					Msg:  err.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: 0,
					Msg:  "代码段收藏成功",
				})
			}
		}
	}
}
