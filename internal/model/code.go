package model

import "time"

type Code struct {
	Id           int       `xorm:"code.id not null pk autoincr comment('主键ID') INT(11)"`
	UserID       int       `xorm:"user_id default 0 comment('用户ID') INT(11)"`
	Content      string    `xorm:"content default 'NULL' comment('内容') VARCHAR(2000)"`
	Category     string    `xorm:"category default 'NULL' comment('类别') VARCHAR(50)"`
	Description  string    `xorm:"description default 'NULL' comment('描述') VARCHAR(2000)"`
	Title        string    `xorm:"title default 'NULL' comment('标题') VARCHAR(150)"`
	Tags         string    `xorm:"tags default 'NULL' comment('标签') VARCHAR(255)"`
	ExpireTime   time.Time `xorm:"expire_time comment('过期时间') DATETIME"`
	Authority    int       `xorm:"authority default 1 comment('权限(1公开 2私人 3加密)') TINYINT(1)"`
	Look         int       `xorm:"look default 0 comment('查看人数') INT(11)"`
	Like         int       `xorm:"like default 0 comment('点赞人数') INT(11)"`
	Collect      int       `xorm:"collect default 0 comment('收藏人数') INT(11)"`
	CodePassword string    `xorm:"code_password default 'NULL' comment('密码') VARCHAR(150)"`
	CreateTime   time.Time `xorm:"create_time comment('创建时间') DATETIME"`
	Mark         int       `xorm:"mark default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}
