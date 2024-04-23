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
			Code:  0,
			Msg:   "登录成功",
			Token: token,
		})
	}
}

func (c *userHandler) Register(ctx *gin.Context) {
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
	// 用户注册
	if err := service.NewUserService().PostRegister(req.UserName, req.Password); err != nil {
		// 注册错误
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "账号不正确",
		})
		return
	} else {
		// 注册成功
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "注册成功",
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

func (c *userHandler) GetUsername(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	// 将 username 从 Token 中解析出来
	userName, err := utils.ParseToken(token)
	if err != nil {
		// 存储失败
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "获取用户名失败，请稍后重试",
		})
		return
	}
	// 登录成功
	ctx.JSON(http.StatusOK, types.JsonResult{
		Code: 0,
		Msg:  "获取用户名成功!",
		Data: userName,
	})
}

func (c *userHandler) GetInformation(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	UserName, err := utils.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	if inform, err := service.NewUserService().GetInformation(UserName); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "获取个人信息失败",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: 0,
			Msg:  "获取个人信息成功",
			Data: inform,
		})
	}
}
