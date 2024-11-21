/*Lv2. 使用web框架编写一个学生管理系统，学生属性可自定义(例如学号，籍贯，生日，性别等)，需实现接口：
添加学生，通过id查询学生，修改学生属性；可以使用json持久化数据，有能力可使用数据库进行持久化。*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name      string `json:"name"`
	StudentID string `json:"studentID"`
	Birthday  string `json:"birthday"`
	Sex       string `json:"sex"`
}

type studentList struct {
	Students []student
}

func search(w http.ResponseWriter, r *http.Request, list studentList) { //用id查询学生
	query := r.URL.Query()
	id := query.Get("id")
	// 如果没有传入 学生id 参数，返回提示
	if id == "" {
		http.Error(w, "没有传入id，查询失败！", http.StatusBadRequest)
		return
	}
	for _, singleStudent := range list.Students {
		if singleStudent.StudentID == id {
			jsoned, err := json.Marshal(singleStudent)
			if err != nil {
				http.Error(w, "Unable to encode student data", http.StatusInternalServerError)
				return
			}
			w.Write(jsoned)
			return
		}
	}
	fmt.Fprintf(w, "没找到")
	return
}

func addStudent(w http.ResponseWriter, r *http.Request, list *studentList) {
	query := r.URL.Query()
	list.Students = append(list.Students, student{})
	list.Students[len(list.Students)-1].Name = query.Get("name")
	list.Students[len(list.Students)-1].StudentID = query.Get("id")
	list.Students[len(list.Students)-1].Birthday = query.Get("birthday")
	list.Students[len(list.Students)-1].Sex = query.Get("sex")
	fmt.Fprintf(w, "该添加的学生信息如下：\n姓名：%s\n学号：%s\n生日：%s\n性别：%s\n", list.Students[len(list.Students)-1].Name, list.Students[len(list.Students)-1].StudentID, list.Students[len(list.Students)-1].Birthday, list.Students[len(list.Students)-1].Sex)
}

func change(w http.ResponseWriter, r *http.Request, list *studentList) { //修改学生属性
	query := r.URL.Query()
	id := query.Get("id")
	newID := query.Get("newid")
	newName := query.Get("newname")
	newBirth := query.Get("newbirth")
	newSex := query.Get("newsex")
	// 如果没有传入 学生id 参数，返回提示
	if id == "" {
		http.Error(w, "没有传入id，修改失败！", http.StatusBadRequest)
		return
	}
	for index, singleStudent := range list.Students {
		if singleStudent.StudentID == id {
			fmt.Fprintf(w, "修改前该学生信息如下：\n姓名：%s\n学号：%s\n生日：%s\n性别：%s\n", singleStudent.Name, singleStudent.StudentID, singleStudent.Birthday, singleStudent.Sex)
			if newName != "" {
				list.Students[index].Name = newName
			}
			if newID != "" {
				list.Students[index].StudentID = newID
			}
			if newBirth != "" {
				list.Students[index].Birthday = newBirth
			}
			if newSex != "" {
				list.Students[index].Sex = newSex
			}
			fmt.Fprintf(w, "修改后该学生信息如下：\n姓名：%s\n学号：%s\n生日：%s\n性别：%s\n", list.Students[index].Name, list.Students[index].StudentID, list.Students[index].Birthday, list.Students[index].Sex)
			return
		}
	}
	fmt.Fprintf(w, "没找到该学生")
}

func main() {
	list := studentList{Students: []student{{Name: "雪豹", StudentID: "001", Sex: "男", Birthday: "2000-01-01"}}}
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		search(w, r, list)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		change(w, r, &list)
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		addStudent(w, r, &list)
	})
	http.ListenAndServe(":8000", nil) // 监听端口及启动服务
}
