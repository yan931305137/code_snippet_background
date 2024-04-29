package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

func Model(XormDb *xorm.Engine) {
	//直接通过结构体，在数据库中创建对应的表【同步结构体与数据表】
	if err := XormDb.Sync(new(User)); err != nil {
		fmt.Println("User表结构同步失败")
	}
	if err := XormDb.Sync(new(Log)); err != nil {
		fmt.Println("Log表结构同步失败")
	}
	if err := XormDb.Sync(new(Ai)); err != nil {
		fmt.Println("Ai表结构同步失败")
	}
	if err := XormDb.Sync(new(Code)); err != nil {
		fmt.Println("Code表结构同步失败")
	}
	if err := XormDb.Sync(new(Top)); err != nil {
		fmt.Println("Top表结构同步失败")
	}
	if err := XormDb.Sync(new(Like)); err != nil {
		fmt.Println("Like表结构同步失败")
	}
	if err := XormDb.Sync(new(Collect)); err != nil {
		fmt.Println("Collect表结构同步失败")
	}

	//给ai表的user_id外键联系到user表的id
	XormDb.Exec("alter table ai add constraint ai_user_id_fk foreign key (user_id) references user (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给code表的user_id外键联系到user表的id
	XormDb.Exec("alter table code add constraint code_user_id_fk foreign key (user_id) references user (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给top表的code_id外键联系到code表的id
	XormDb.Exec("alter table top add constraint top_code_id_fk foreign key (code_id) references code (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给like表的code_id外键联系到code表的id
	XormDb.Exec("alter table like add constraint like_code_id_fk foreign key (code_id) references code (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给like表的user_id外键联系到user表的id
	XormDb.Exec("alter table like add constraint like_user_id_fk foreign key (user_id) references user (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给collect表的code_id外键联系到code表的id
	XormDb.Exec("alter table focus add constraint collect_code_id_fk foreign key (code_id) references code (id) ON UPDATE CASCADE ON DELETE CASCADE")
	//给collect表的user_id外键联系到user表的id
	XormDb.Exec("alter table focus add constraint collect_user_id_fk foreign key (user_id) references user (id) ON UPDATE CASCADE ON DELETE CASCADE")
}
