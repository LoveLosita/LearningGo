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

func DeleteStudents(ctx context.Context, c *app.RequestContext) { //删除学生模块
	//1.从GET请求中获取数据
	id := c.Query("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	query := "DELETE FROM students WHERE id = ?"
	// 2.执行删除操作
	result, err := Db.Exec(query, int(intID))
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	// 3.检查删除的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	//4.通过删除行数倒推是否删除，如果为0就是没删除，相当于没找到
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

func AddStudents(ctx context.Context, c *app.RequestContext) { //添加学生模块
	//1.从POST请求中获取信息
	studentInfo := Student{}
	err := c.BindJSON(&studentInfo) //从Json请求中获取信息
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": "未知的请求数据",
		})
		return
	}
	//2.执行数据库操作，准备插入新学生
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

func ChangeStudentInfo(ctx context.Context, c *app.RequestContext) { //改变学生信息模块
	//1.从POST请求中获取信息
	newStudentInfo := Student{}
	err := c.BindJSON(&newStudentInfo)
	targetID := c.Query("id")
	intTargetID, err := strconv.ParseInt(targetID, 10, 0)
	student := Student{}
	rows, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE id = ?", int(intTargetID))
	if newStudentInfo.Id != 0 { //此处为了防止更改id，导致自增炸
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
	if newStudentInfo.Username != "" { //判断是否在新信息中填写了名字，即是否想改名字
		ifSameUsername, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE username = ?", newStudentInfo.Username)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if ifSameUsername.Next() { //如果通过要改的名字能找到学生，说明用户名已存在
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": "用户名已存在！",
			})
			return
		}
	}
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
	//下面是一个选择性填写的设计，只需要填写想更改的信息就行了，不想改可以直接不用填写
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
		//这里是判断性别是否正确，把错误扼杀在写入数据库之前
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
	//选择性填写设计模块至此结束
	//下面开始更新数据库信息
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

func FindStudent(ctx context.Context, c *app.RequestContext) { //查询学生的模块
	//1.从url中获取想要查找学生的id
	id := c.Query("id")
	rows, err := Db.Query("SELECT id ,name, age, gender, class,username,password FROM students WHERE id = ?", id)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	student := Student{} //准备结构体用于写入信息并呈现
	if rows.Next() {     //找到学生了
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Class, &student.Username, &student.Password)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		// 如果没有找到学生
		c.JSON(consts.StatusNotFound, map[string]string{
			"error": "没有找到指定的学生",
		})
		return
	}
	c.JSON(consts.StatusOK, student)
}
