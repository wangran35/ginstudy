package models

import (
	"errors"
	"time"
)

type Authrule struct {
	Id         int64
	Type       int `json:"type" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"`
	Pid        int
	Name       string `json:"name"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Permission string `json:"permission"`
	Remark     string
	Ismenu     int    `json:"ismenu" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"`
	Menutype   string `xorm:"varchar(59)"`
	Weigh      int
	Status     int        `json:"status" xorm:"not null default 1 comment('账号状态 默认1 正常 0 未启用') TINYINT"`
	Created    time.Time  `xorm:"created"`
	Updated    time.Time  `xorm:"updated"`
	Children   []Authrule `json:"children,omitempty" gorm:"-"`
}

func (a *Authrule) TableName() string {
	return "auth_rule"
}

//根据用户名密码查询用户
func Selectrules() (*Authrule, error) {
	a := new(Authrule)
	has, err := orm.Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("没有数据！")
	}
	return a, nil

}
