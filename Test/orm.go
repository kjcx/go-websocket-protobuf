package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Member struct {
	Member_id int `orm:"pk;column(member_id);"`
	Member_name string
}
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/chuangke?charset=utf8")
	//orm.RegisterModel(new (Member))
	orm.RegisterModelWithPrefix("chuangke_", new(Member))
}


func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	s := &Member{}
	s.Member_id = 1
	s.Member_name ="admin"

	err := o.Read(s,"Member_name")

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(s)
	}

}
