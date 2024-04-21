package model

import "time"

type Ai struct {
	Id         int       `xorm:"id not null pk autoincr comment('主键ID') INT(11)"`
	UserId     int       `xorm:"user_id not null comment('用户id外键关联') INT(11)"`
	Content    string    `xorm:"content default 'NULL' comment('对话记录') VARCHAR(2000)"`
	Status     int       `xorm:"status default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	UpdateTime time.Time `xorm:"update_time comment('创建时间') DATETIME"`
	Mark       int       `xorm:"mark default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}
