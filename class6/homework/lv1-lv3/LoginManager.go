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

func Login(ctx context.Context, c *app.RequestContext) { //登录模块
	//1.处理POST请求
	student := UserStudent{}
	err := c.BindJSON(&student) //从POST请求中获取学生数据，存入student结构体中
	//2.准备开始对比用户名密码
	name := ""
	getPassword := ""
	query := "SELECT name,password FROM students WHERE username=?" //从数据库中提取用户名和密码，用于下面的比较
	rows, err := Db.Query(query, student.Username)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	//3.对比密码
	if !rows.Next() { //上面使用POST请求中获取的用户名扫描未获得结果，即用户名错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"登录失败": "错误的用户名或密码",
		})
	} else { //找到了用户，进行下一步密码的判断
		err = rows.Scan(&name, &getPassword)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if getPassword != student.Password { //POST请求中的密码和数据库中的密码不相等，即密码错误
			c.JSON(consts.StatusBadRequest, map[string]string{
				"登录失败": "错误的用户名或密码",
			})
		} else { //相等，即成功登录
			c.JSON(consts.StatusBadRequest, map[string]string{
				"登录成功": "欢迎" + name + "同学",
			})
		}
	}
}

func ChangePassword(ctx context.Context, c *app.RequestContext) {
	//1.从POST请求获取数据
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
	//2.从数据库中提取名字和密码，用于下一步比较
	query := "SELECT name,password FROM students WHERE username=?"
	rows, err := Db.Query(query, student.Username)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	//3.开始比较
	if !rows.Next() { //找不到用户名，说明用户名输入错误
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "用户名不存在！",
		})
		return
	} else { //找到了用户名，需要进一步对比旧密码是否正确
		err = rows.Scan(&username, &oldPassword) //从数据库中提取用户名和旧密码
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if student.OldPassword != oldPassword { //如果POST请求中的旧密码和数据库中的密码不一样，即旧密码错误
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": "旧密码错误，修改失败！",
			})
			return
		} else { //密码正确，开始执行修改的流程
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
