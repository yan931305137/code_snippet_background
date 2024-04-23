package types

type CodeReq struct {
	Content      string `form:"content" binding:"required"`
	Category     string `form:"category"`
	Description  string `form:"description" `
	Title        string `form:"title" `
	Tags         string `form:"tags" `
	ExpireTime   string `form:"expire_time"`
	Authority    int    `form:"authority" binding:"required"`
	CodePassword string `form:"code_password" `
}
