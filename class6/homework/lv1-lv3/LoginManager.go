package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	name := ""
	getPassword := ""
	query := "SELECT name,password FROM students WHERE username=?"
	rows, err := Db.Query(query, username)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if !rows.Next() {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"登录失败": "错误的用户名或密码",
		})
	} else {
		err = rows.Scan(&name, &getPassword)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if getPassword != password {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"登录失败": "错误的用户名或密码",
			})
		} else {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"登录成功": "欢迎" + name + "同学",
			})
		}
	}
}

func ChangePassword(ctx context.Context, c *app.RequestContext) {

}
