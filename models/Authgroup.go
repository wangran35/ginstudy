package models

import (
	"time"
)

type Authgroup struct {
	Id      int64
	Pid     int `xorm:"varchar(200)"`
	Name    string
	Rules   string    `json:"rules" xorm:"LONGTEXT "`
	Status  int       `json:"status" xorm:"not null default 1 comment('账号状态 默认1 正常 0 未启用') TINYINT"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func (a *Authgroup) TableName() string {
	return "auth_group"
}
