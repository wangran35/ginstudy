package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id           int64
	Username     string    `xorm:"varchar(200)" json:"username"`
	Nickname     string    `xorm:"varchar(200)" json:"nickname"`
	Salt         string    `xorm:"varchar(200)" json:"salt"`
	Age          int       `xorm:"int(2)" json:"age"`
	Avatar       string    `xorm:"TEXT" json:"avatar"`
	Loginfailure int       `xorm:"int(10)" json:"loginfailure"`
	Logintime    int       `xorm:"int(10)" json:"logintime"`
	Loginip      string    `xorm:"varchar(200)" json:"loginip"`
	Token        string    `xorm:"varchar(59)"`
	Password     string    `xorm:"varchar(200)"`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}

func (a *Admin) TableName() string {
	return "admin"
}

//根据用户名密码查询用户
func SelectUserByUserName(userName string) (*Admin, error) {
	a := new(Admin)
	has, err := orm.Where("username = ?", userName).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户未找到！")
	}
	return a, nil

}

//add
func AddAdmin(a *Admin) error {
	_, err := orm.Insert(a)
	return err
}
