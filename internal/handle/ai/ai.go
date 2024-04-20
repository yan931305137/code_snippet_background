package handle

import (
	"github.com/gin-gonic/gin"
)

// Ai控制器管理对象
var Ai = new(aiHandler)

type aiHandler struct{}

func (c *aiHandler) AiKnow(ctx *gin.Context) {

}
