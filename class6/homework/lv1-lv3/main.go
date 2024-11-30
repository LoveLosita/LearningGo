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
	h.GET("/change", ChangeStudentInfo)
	h.GET("/register", AddStudents)
	h.GET("/delete", DeleteStudents)
	h.GET("login", Login)
	h.Spin()
}
