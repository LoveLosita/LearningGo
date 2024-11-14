package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Author struct {
	Name string
	Bio  string
}

type Post struct {
	Title   string
	Content string
	Author  Author
	Tags    []string
}

func main() {
	author := Author{Name: "Losita", Bio: "Man!"}
	post := Post{Title: "Today is crazy Thursday!",
		Content: "V me 50 Yuan.", Author: author, Tags: []string{"Beg", "KFC", "Shit"}}
	postJson, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(postJson))
	postJsonSlice := []string{} //此处输出Json字符串
	postJsonString := string(postJson)
	postJsonString = strings.Trim(postJsonString, "{}") //掐头去尾
	tempStr := ""
	going := false
	//选择自己处理分隔，保证Author部分的复合字符串能被保留
	for i := 0; i < len(postJsonString); i++ {
		if string(postJsonString[i]) == "{" || string(postJsonString[i]) == "[" { //遇到左大括号或者中括号，开启不跳过逗号模式
			tempStr += string(postJsonString[i])
			going = true
		} else if string(postJsonString[i]) == "}" || string(postJsonString[i]) == "]" { //遇到右大括号或者中括号，结束不跳过逗号模式
			going = false
			tempStr += string(postJsonString[i])
			postJsonSlice = append(postJsonSlice, tempStr)
			tempStr = ""
		} else if string(postJsonString[i]) == "," && !going { //不是不跳过逗号模式时如果遇到逗号，那就结算
			postJsonSlice = append(postJsonSlice, tempStr)
			tempStr = ""
		} else { //其他情况就正常加入
			tempStr += string(postJsonString[i])
		}
	}
	postJsonSlice = append(postJsonSlice, tempStr) //最后退出时也要结算
	for _, value := range postJsonSlice {          //此处输出键值对
		fmt.Println(value)
	}
	authorString := postJsonSlice[2]
	authorString = authorString[9:]
	authorStruct := Author{}
	err = json.Unmarshal([]byte(authorString), &authorStruct)
	fmt.Printf("Name:%s\nBio:%s", authorStruct.Name, authorStruct.Bio)
}
