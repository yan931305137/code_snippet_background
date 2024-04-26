package handle

import (
	"code_snippet/internal/model"
	service "code_snippet/internal/services/ai"
	"code_snippet/internal/types"
	"code_snippet/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Ai控制器管理对象
var Ai = new(aiHandler)

type aiHandler struct{}

func (c *aiHandler) GetAiKnow(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	if userid, err := utils.GetTokenToRedis(token); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "请先登录再进行操作!",
		})
		return
	} else {
		if content, err := service.NewAiService().GetAiKnow(userid); err != nil {
			// 获取对话失败
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code: -1,
				Msg:  "获取对话失败",
			})
			return
		} else {
			// 获取对话成功
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code:  0,
				Msg:   "获取对话成功",
				Data:  *content,
				Count: int64(len(*content)),
			})
		}
	}
}

// PostAiKnow 更新对话
func (c *aiHandler) PutAiKnow(ctx *gin.Context) {
	var req *types.AiReq
	// 获取参数并验证
	if err := ctx.ShouldBind(&req); err != nil {
		// 返回错误信息
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	token := utils.GetToken(ctx)
	if userid, err := utils.GetTokenToRedis(token); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "请先登录再进行操作!",
		})
		return
	} else {
		if aitalk, err := service.NewAiService().GetAiTalk(req.Id, userid); err != nil {
			fmt.Println(aitalk, err.Error())
			// 更新对话失败
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code: -1,
				Msg:  "获取历史对话失败",
			})
			return
		} else {
			// 添加新的内容
			req.Content = strings.Replace(req.Content, "\\", "\\\\", -1)
			req.Content = strings.Replace(req.Content, "\"", "'", -1)
			req.Content = strings.Replace(req.Content, "\n", "\\n", -1)
			var newMessage string
			if req.Id == -1 {
				newMessage = fmt.Sprintf(aitalk + `{"role":"user","content":"` + req.Content + `"}`)
			} else {
				newMessage = fmt.Sprintf(aitalk + `,{"role":"user","content":"` + req.Content + `"}`)
			}
			content := utils.AiMessage(newMessage)
			content = strings.Replace(content, "\\", "\\\\", -1)
			content = strings.Replace(content, "\"", "'", -1)
			newcontent := strings.Replace(content, "\n", "\\n", -1)
			newJson := fmt.Sprintf(newMessage + `,{"role": "assistant", "content": "` + newcontent + `"}`)
			var aiInfo model.Ai
			aiInfo.Content = newJson
			aiInfo.Id = req.Id
			aiInfo.UserId = userid
			if content != "" {
				if err := service.NewAiService().PostAiKnow(aiInfo); err != nil {
					// 更新对话失败
					ctx.JSON(http.StatusOK, types.JsonResult{
						Code: -1,
						Msg:  "更新对话失败",
					})
					return
				} else {
					// 更新对话成功
					ctx.JSON(http.StatusOK, types.JsonResult{
						Code: 0,
						Msg:  "更新对话成功",
						Data: content,
					})
				}
			} else {
				ctx.JSON(http.StatusOK, types.JsonResult{
					Code: -1,
					Msg:  "更新对话失败",
				})
			}
		}

	}
}

func (c *aiHandler) DeleteListAiKnow(ctx *gin.Context) {
	token := utils.GetToken(ctx)
	if userid, err := utils.GetTokenToRedis(token); err != nil {
		ctx.JSON(http.StatusOK, types.JsonResult{
			Code: -1,
			Msg:  "请先登录再进行操作!",
		})
		return
	} else {
		if err := service.NewAiService().DeleteListAiKnow(userid); err != nil {
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code: -1,
				Msg:  "删除历史对话记录失败",
			})
			return
		} else {
			ctx.JSON(http.StatusOK, types.JsonResult{
				Code: 0,
				Msg:  "删除历史对话记录成功",
			})
		}
	}
}
