package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default() //启动服务
	err := ConnectDB()    //链接数据库
	if err != nil {
		fmt.Println(err)
		return
	}
	defer DisconnectDB()                 //在退出时关闭数据库连接
	h.GET("/search", FindStudent)        //搜索模块
	h.POST("/change", ChangeStudentInfo) //改变学生信息模块
	h.POST("/register", AddStudents)     //新增学生模块
	h.GET("/delete", DeleteStudents)     //删除学生模块
	h.POST("/login", Login)              //学生用户名密码登录模块
	h.POST("/changePwd", ChangePassword) //学生改密码模块
	h.Spin()                             //使服务保持运行
}
