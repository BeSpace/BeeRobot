package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type FAQ struct {
	Id         int
	Question   string `orm:"size(512)"`
	Answer     string `orm:"size(512)"`
	Tag        string `orm:"size(512)"`
	Actived    int
	CreateTime time.Time
	EditTime   time.Time
	Deleted    int
}

func (m *FAQ) TableName() string {
	return "faq"

}

func init() {
	orm.RegisterModel(new(FAQ))
}

func SaveFAQ(m *FAQ) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
