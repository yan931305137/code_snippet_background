package types

type AiReq struct {
	Id      int    `form:"id"  binding:"required"`
	Content string `form:"content" binding:"required"`
}

type BdAiReq struct {
	Result string `form:"result" binding:"required"`
}
