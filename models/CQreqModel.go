package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"
)


/**
 * Struct of CQreq
 */
type CQreq struct {
	Id int
	CreateTime time.Time
	Message string `orm:"size(1024)"`	// length is 1024
	Delete int 							// 0 or 1

}


/**
 * CQreq query param
 */
type CQreqQueryParam struct {
	BaseQueryParam
	NameLike string // name_isstartwith
}

/**
 * [init orm of CQreq]
 */
func init() {
	orm.RegisterModel(new(CQreq))
}

/**
 * [SaveCQreq into db]
 * @param {[type]} m *CQreq [description]
 */
func SaveCQreq(m *CQreq){
	o := orm.NewOrm()

	m.CreateTime = time.Now()
	m.Delete = 0
	// output result of exec
	logs.Info(o.Insert(m))
}

/**
 * [PageListCQreg :Query with Pagers]
 * @param {[type]} )([]*CQreq, int64 [description]
 */
func PageListCQreg(params *BaseQueryParam)([]*CQreq, int64){
	
	query := orm.NewOrm().QueryTable("c_qreq")
	data := make([]*CQreq, 0)

	total, _ := query.Count()

	query.Limit(params.Limit, (params.Page-1)*params.Limit).All(&data)
	return data, total
}


