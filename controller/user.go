package controller

import (
	"be/model"
	"be/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

//	用户组的handler

func LoginByUP(c *gin.Context) {
	type Login struct {
		UserName string `json:"userName" form:"userName"`
		Password string `json:"password" form:"password"`
		Remember bool   `json:"remember" form:"remember"`
	}
	l := Login{}
	err := c.ShouldBind(&l)
	if err != nil {
		//	解析数据失败
		c.JSON(200, gin.H{
			"code": 0,
			"data": gin.H{
				"msg": "解析数据失败" + fmt.Sprintf("%s", err),
			},
		})
		c.Abort()
		return
	}
	//	到数据库查询，用户是否存在
	i, u := model.IsExist(l.UserName)
	if i != 1 {
		c.JSON(200, gin.H{
			"code": 0,
			"data": gin.H{
				"msg": "用户不存在,请注册",
			},
		})
		c.Abort()
		return
	}
	if !utils.CheckHash(l.Password, u.PasswordHash) {
		c.JSON(200, gin.H{
			"code": 0,
			"data": gin.H{
				"msg": "密码错误,请重新输入或者找回",
			},
		})
		c.Abort()
		return
	}
	myClaims := utils.MyClaims{
		UId:              int(u.ID),
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	tokenString, err := utils.GeneToken(myClaims)
	if err != nil {
		fmt.Println("token生成出错" + fmt.Sprintf("%s", err))
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": gin.H{
			"msg":   "登录成功",
			"token": tokenString,
		},
	})
}

func Register(c *gin.Context) {
	type Register struct {
		UserName string `json:"userName" form:"userName"`
		Password string `json:"password" form:"password"`
	}
	r := Register{}
	c.ShouldBind(&r)
	u := model.User{
		UserName:     r.UserName,
		PasswordHash: utils.GenerateHash(r.Password),
	}
	model.Add(&u)
}
