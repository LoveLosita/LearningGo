package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type UserStudent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordStudent struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func Login(ctx context.Context, c *app.RequestContext) {
	student := UserStudent{}
	err := c.BindJSON(&student)
	name := ""
	getPassword := ""
	query := "SELECT name,password FROM students WHERE username=?"
	rows, err := Db.Query(query, student.Username)
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
		if getPassword != student.Password {
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
	student := ChangePasswordStudent{}
	err := c.BindJSON(&student)
	oldPassword := ""
	username := ""
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	query := "SELECT name,password FROM students WHERE username=?"
	rows, err := Db.Query(query, student.Username)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if !rows.Next() {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "用户名不存在！",
		})
		return
	} else {
		err = rows.Scan(&username, &oldPassword)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if student.OldPassword != oldPassword {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": "旧密码错误，修改失败！",
			})
			return
		} else {
			query = "UPDATE students SET password=? WHERE username=?"
			result, err := Db.Exec(query, student.NewPassword, student.Username)
			if err != nil {
				c.JSON(consts.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
				return
			}
			// 获取修改的行数
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				c.JSON(consts.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
				return
			}
			// 返回成功更新的消息
			c.JSON(consts.StatusOK, map[string]string{
				"成功": "成功更新了" + strconv.FormatInt(rowsAffected, 10) + "行数据",
			})
		}
	}
}
