package model

// Log 表示日志记录的数据模型
type Log struct {
	Id           int    `xorm:"'id' not null pk autoincr comment('主键ID') INT(11)"`
	ClientIP     string `xorm:"'client_ip' default 'NULL' comment('客户端IP') VARCHAR(20)"`
	EdTime       int    `xorm:"'ed_time' default 0 comment('EdTime') INT(11)"`
	Method       string `xorm:"'method' default 'NULL' comment('请求方法') VARCHAR(50)"`
	RequestURI   string `xorm:"'request_uri' default 'NULL' comment('请求URI') VARCHAR(255)"`
	ResponseCode int    `xorm:"'response_code' default 0 comment('响应状态码') INT(11)"`
	ResponseData string `xorm:"'response_data' default 'NULL' comment('响应数据') VARCHAR(9999)"`
	ResponseMsg  string `xorm:"'response_msg' default 'NULL' comment('响应消息') VARCHAR(255)"`
	RoleName     string `xorm:"'role_name' default 'NULL' comment('角色名称') VARCHAR(50)"`
	StartTime    string `xorm:"'start_time' comment('开始时间') VARCHAR(50)"`
	Level        string `xorm:"'level' default 'info' comment('日志级别') VARCHAR(10)"`
	Msg          string `xorm:"'msg' default 'NULL' comment('日志消息')  VARCHAR(9999)"`
	Time         string `xorm:"'time' comment('记录时间') VARCHAR(50)"`
}
