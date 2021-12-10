package controllers

import (
	"ginstudy/lib"
	"ginstudy/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Adminform struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// type AdminController struct {
// 	BaseController //继承统一判断是否登录或者是否有权限类
// 	// beego.Controller
// }
func AddAdmin(c *gin.Context) {
	var admindata Adminform
	if err := c.ShouldBind(&admindata); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"msg":     "表单未完整！",
			"msgcode": "-1",
			// "data":    reqIP,
		})
		return
	}
	// //读取数据库
	Admin := new(models.Admin)
	Admin.Username = admindata.Username
	info, _ := models.SelectUserByUserName(Admin.Username) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该用户已经存在！",
		})
		return
	}
	pwd, salt := lib.Password(4, admindata.Password) //截取四位随机盐+上这个做原始密码
	Admin.Password = pwd
	Admin.Salt = salt
	Admin.Created = time.Now()
	err := models.AddAdmin(Admin) //判断账号是否存在！
	if err != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		result := make(map[string]interface{})
		result["id"] = Admin.Id //返回当前总数
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "数据添加成功！",
			"data": result,
		})

	}
}
