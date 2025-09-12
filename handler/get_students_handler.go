package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 处理获取学生列表请求
func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "只允许GET请求", http.StatusMethodNotAllowed)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 获取学生ID参数
	studentID := r.URL.Query().Get("studentId")

	// 创建响应数据切片
	var responseData []Student

	if studentID != "" {
		// 查询单个学生
		if student, exists := GetStudentByID(studentID); exists {
			responseData = append(responseData, student)
		}
	} else {
		// 查询所有学生
		responseData = GetAllStudents()
	}

	// 返回响应
	response := QueryResponse{
		Success: true,
		Message: fmt.Sprintf("共查询到 %d 条学生信息", len(responseData)),
		Data:    responseData,
	}

	json.NewEncoder(w).Encode(response)
}