package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Class    string `json:"class"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func DeleteStudents(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	query := "DELETE FROM students WHERE id = ?"
	// 执行删除操作
	result, err := Db.Exec(query, int(intID))
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	// 检查删除的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "没找到学生",
		})
		return
	}
	c.JSON(consts.StatusBadRequest, map[string]string{
		"成功": "已经删除此学生",
	})
}

func AddStudents(ctx context.Context, c *app.RequestContext) {
	studentInfo := Student{}
	err := c.BindJSON(&studentInfo) //从Json请求中获取信息
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "未知的请求数据",
		})
		return
	}
	query := ""
	query = "INSERT INTO students (name, age, gender, class,username,password) VALUES (?, ?, ?, ?,?,?)"
	_, err = Db.Exec(query, studentInfo.Name, studentInfo.Age, studentInfo.Gender, studentInfo.Class, studentInfo.Username, studentInfo.Password)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, map[string]string{
		"成功": "已经存入数据库",
	})
}

func ChangeStudentInfo(ctx context.Context, c *app.RequestContext) {
	newStudentInfo := Student{}
	err := c.BindJSON(&newStudentInfo)
	targetID := c.Query("id")
	intTargetID, err := strconv.ParseInt(targetID, 10, 0)
	student := Student{}
	rows, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE id = ?", int(intTargetID))
	if newStudentInfo.Id != 0 {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "不允许更改id！",
		})
		return
	}
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if newStudentInfo.Username != "" {
		ifSameUsername, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE username = ?", newStudentInfo.Username)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if ifSameUsername.Next() {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": "用户名已存在！",
			})
			return
		}
	}
	// 这里加入 rows.Next()
	if rows.Next() { // 确保至少有一行返回
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Class, &student.Username, &student.Password)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		// 如果没有找到对应的学生，返回错误信息
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "未找到该学生信息！",
		})
		return
	}
	if newStudentInfo.Name != "" {
		student.Name = newStudentInfo.Name
	}
	if newStudentInfo.Username != "" {
		student.Username = newStudentInfo.Username
	}
	if newStudentInfo.Password != "" {
		student.Password = newStudentInfo.Password
	}
	if newStudentInfo.Gender != "" && (newStudentInfo.Gender == "男" || newStudentInfo.Gender == "女") {
		student.Gender = newStudentInfo.Gender
	} else if newStudentInfo.Gender != "" && !(newStudentInfo.Gender == "男" || newStudentInfo.Gender == "女") {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "性别错误！",
		})
	}
	if newStudentInfo.Class != "" {
		student.Class = newStudentInfo.Class
	}
	if newStudentInfo.Age != 0 {
		student.Age = newStudentInfo.Age
	}
	query := "UPDATE students SET name=?, age=?,gender=?,class=?,username=?,password=? WHERE id=?"
	result, err := Db.Exec(query, student.Name, student.Age, student.Gender, student.Class, student.Username, student.Password, intTargetID)
	if err != nil {
		// 如果发生错误，返回错误信息
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

func FindStudent(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	rows, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE id = ?", id)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	student := Student{}
	// 使用 rows.Next() 遍历查询结果
	if rows.Next() {
		// 只有在 rows.Next() 返回 true 时才调用 rows.Scan
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Class, &student.Username, &student.Password)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		// 如果没有找到数据，返回相应的错误信息
		c.JSON(consts.StatusNotFound, map[string]string{
			"error": "没有找到指定的学生",
		})
		return
	}
	c.JSON(consts.StatusOK, student)
}
