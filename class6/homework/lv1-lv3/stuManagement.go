package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Class  string `json:"class"`
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
	id := c.Query("id")
	rows, err := Db.Query("SELECT id ,name, age, gender, class FROM students WHERE id = ?", id)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if rows.Next() {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "该id已存在，请换一个",
		})
		return
	}
	name := c.Query("name")
	age := c.Query("age")
	gender := c.Query("gender")
	class := c.Query("class")
	username := c.Query("username")
	password := c.Query("password")
	query := ""
	if name == "" || age == "" || gender == "" || class == "" || password == "" || username == "" {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "参数不全",
		})
		return
	}
	intAge, err := strconv.ParseInt(age, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	if id != "" {
		query = "INSERT INTO students (id,name, age, gender, class,username,password) VALUES (?, ?, ?, ?,?,?)"
		// 执行插入操作
		_, err := Db.Exec(query, id, name, int(intAge), gender, class, username, password)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		query = "INSERT INTO students (name, age, gender, class,username,password) VALUES (?, ?, ?, ?,?,?)"
		_, err := Db.Exec(query, name, int(intAge), gender, class, username, password)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	}
	c.JSON(consts.StatusOK, map[string]string{
		"成功": "已经存入数据库",
	})
}

func ChangeStudentInfo(ctx context.Context, c *app.RequestContext) {
	targetID := c.Query("id")
	intTargetID, err := strconv.ParseInt(targetID, 10, 0)
	student := Student{}
	rows, err := Db.Query("SELECT id ,name, age, gender, class FROM students WHERE id = ?", int(intTargetID))
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	// 这里加入 rows.Next()
	if rows.Next() { // 确保至少有一行返回
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Class)
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
	name := c.Query("name")
	id := c.Query("newid")
	intID, err := strconv.ParseInt(id, 10, 0)
	gender := c.Query("gender")
	class := c.Query("class")
	age := c.Query("age")
	intAge, err := strconv.ParseInt(age, 10, 0)
	if name != "" {
		student.Name = name
	}
	if id != "" {
		student.Id = int(intID)
	}
	if gender != "" && (gender == "男" || gender == "女") {
		student.Gender = gender
	} else if gender != "" && !(gender == "男" || gender == "女") {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "性别错误！",
		})
	}
	if class != "" {
		student.Class = class
	}
	if age != "" {
		student.Age = int(intAge)
	}
	query := "UPDATE students SET id=?,name=?, age=?,gender=?,class=? WHERE id=?"
	result, err := Db.Exec(query, student.Id, student.Name, student.Age, student.Gender, student.Class, intTargetID)
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
	rows, err := Db.Query("SELECT id ,name, age, gender, class FROM students WHERE id = ?", id)
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
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Class)
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
