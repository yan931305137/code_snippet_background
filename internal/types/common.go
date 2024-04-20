package types

type Token struct {
	Token string `form:"token" binding:"required"`
}

type Id struct {
	Id int `uri:"id" binding:"required"`
}
type Num struct {
	Num int `json:"num" binding:"min=0"`
}

type Page struct {
	Size int `uri:"size" binding:"min=0"`
	Num  int `uri:"num" binding:"min=0"`
}

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
	Token string      `json:"token"`
}

type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}
