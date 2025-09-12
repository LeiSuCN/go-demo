package main

import (
	"fmt"
	"net/http"
	"demo/handler"
)

// 主函数
func main() {
	// 初始化学生数据存储
	handler.InitStorage()

	// 添加测试数据
	handler.AddTestData()

	// 设置静态文件服务，提供public目录下的文件
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// 设置API路由，使用handler包中的函数
	http.HandleFunc("/api/submit-student", handler.SubmitStudentHandler)
	http.HandleFunc("/api/get-students", handler.GetStudentsHandler)

	// 启动HTTP服务器
	fmt.Println("服务器启动成功！访问 http://localhost:8080")
	var err error
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("服务器启动失败: %s\n", err)
	}
}
