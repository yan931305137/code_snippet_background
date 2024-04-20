package handle

import (
	service "code_snippet/internal/services/user"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 用户控制器管理对象
var User = new(userHandler)

type userHandler struct{}

func (c *userHandler) Login(ctx *gin.Context) {
	var req *types.LoginReq
	// 获取参数并验证
	if err := ctx.ShouldBind(&req); err != nil {
		// 返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 校验验证码
	verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
	if !verifyRes {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "验证码不正确",
		})
		return
	}
	// 管理员登录
	if id, token, err := service.NewUserService().PostLogin(req.UserName, req.Password); err != nil {
		// 登录错误
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "账号不正确",
		})
		return
	} else {
		// 将 token 存储到 Redis 中
		if err := utils.SaveTokenToRedis(token, id); err != nil {
			// 存储失败
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code: -1,
				Msg:  "登录失败，请稍后重试",
			})
			return
		}
		// 登录成功
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "登录成功",
		})
	}
}

// 验证码
func (c *userHandler) Captcha(ctx *gin.Context) {
	idKeyC, base64stringC := utils.Captcha()
	// 返回结果集
	ctx.JSON(http.StatusOK, types.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

func (c *userHandler) Exit(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	// 将 token 存储到 Redis 中
	if err := utils.RD.Del(ctx, token).Err(); err != nil {
		// 存储失败
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "退出失败，请稍后重试",
		})
		return
	}
	// 登录成功
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "退出成功",
	})
}

func (c *userHandler) GetUsername(context *gin.Context) {

}

func (c *userHandler) Register(context *gin.Context) {

}
