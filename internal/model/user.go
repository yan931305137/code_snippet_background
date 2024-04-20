package model

import "time"

type User struct {
	Id           int       `xorm:"id not null pk autoincr comment('主键ID') INT(11)"`
	UserName     string    `xorm:"user_name default 'NULL' comment('登录用户名') VARCHAR(50)"`
	Password     string    `xorm:"password default 'NULL' comment('登录密码') VARCHAR(150)"`
	Gender       int       `xorm:"gender default 3 comment('性别:1男 2女 3保密') TINYINT(1)"`
	Avatar       string    `xorm:"avatar default 'NULL' comment('头像') VARCHAR(150)"`
	Mobile       string    `xorm:"mobile default 'NULL' comment('手机号码') CHAR(11)"`
	Email        string    `xorm:"email default 'NULL' comment('邮箱地址') VARCHAR(30)"`
	Birthday     time.Time `xorm:"birthday comment('出生日期') DATE"`
	ProvinceCode string    `xorm:"province_code default 'NULL' comment('省份编号') VARCHAR(50)"`
	CityCode     string    `xorm:"city_code default 'NULL' comment('市区编号') VARCHAR(50)"`
	DistrictCode string    `xorm:"district_code default 'NULL' comment('区县编号') VARCHAR(50)"`
	Address      string    `xorm:"address default 'NULL' comment('详细地址') VARCHAR(255)"`
	CityName     string    `xorm:"city_name default 'NULL' comment('所属城市') VARCHAR(150)"`
	Status       int       `xorm:"status default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	LoginNum     int       `xorm:"login_num default 0 comment('登录次数') INT(11)"`
	LoginIp      string    `xorm:"login_ip default 'NULL' comment('最近登录IP') VARCHAR(20)"`
	LoginTime    time.Time `xorm:"login_time comment('最近登录时间') DATETIME"`
	CreateTime   time.Time `xorm:"create_time comment('创建时间') DATETIME"`
	Mark         int       `xorm:"mark default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}
