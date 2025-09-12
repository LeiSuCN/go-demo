package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 处理学生信息提交请求
func SubmitStudentHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许POST请求
	if r.Method != http.MethodPost {
		http.Error(w, "只允许POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 解析请求体
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		RespondWithError(w, "请求数据格式错误")
		return
	}

	// 验证必填字段
	if student.StudentID == "" || student.Name == "" || student.Gender == "" || student.BirthDate == "" {
		RespondWithError(w, "请填写所有必填字段")
		return
	}

	// 检查学号是否已存在（确保唯一性）
	if StudentExists(student.StudentID) {
		RespondWithError(w, fmt.Sprintf("学号 %s 已存在，请使用不同的学号", student.StudentID))
		return
	}

	// 保存学生信息
	AddStudent(student)

	// 打印保存的学生信息（用于调试）
	fmt.Printf("保存学生信息: %+v\n", student)

	// 返回成功响应
	RespondWithSuccess(w, "学生信息提交成功")
}