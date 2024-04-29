package model

import "time"

// Like
type Like struct {
	Id         int       `xorm:"'id' not null pk autoincr comment('主键ID') INT(11)"`
	CodeID     int       `xorm:"code_id default 0 comment('代码ID') INT(11)"`
	UserID     int       `xorm:"user_id default 0 comment('代码ID') INT(11)"`
	Status     int       `xorm:"status default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	UpdateTime time.Time `xorm:"update_time comment('创建时间') DATETIME"`
	Mark       int       `xorm:"mark default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}
