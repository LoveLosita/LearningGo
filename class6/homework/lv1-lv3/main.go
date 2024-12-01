package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer DisconnectDB()
	h.GET("/search", FindStudent)
	h.POST("/change", ChangeStudentInfo)
	h.POST("/register", AddStudents)
	h.GET("/delete", DeleteStudents)
	h.POST("/login", Login)
	h.POST("/changePwd", ChangePassword)
	h.Spin()
}
