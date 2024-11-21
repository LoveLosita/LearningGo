package main

import (
	"fmt"
	"net/http"
)

// ping 响应函数
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong!")
}

func echo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	message := query.Get("message")
	// 如果没有传入 message 参数，返回提示
	if message == "" {
		http.Error(w, "没有传入参数", http.StatusBadRequest)
		return
	}
	// 返回 message 参数内容
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/ping", ping)    // 创建ping路由
	http.HandleFunc("/echo", echo)    // 创建echo路由
	http.ListenAndServe(":8000", nil) // 监听端口及启动服务
}
