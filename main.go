package main

import (
	"fmt"
	"net/http"
)

// 主函数
func main() {
	// 设置静态文件服务，提供public目录下的文件
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	
	// 启动HTTP服务器
	fmt.Println("服务器启动成功！访问 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("服务器启动失败: %s\n", err)
	}
}