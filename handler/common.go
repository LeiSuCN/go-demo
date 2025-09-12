package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

// 定义学生结构体
type Student struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

// 定义响应结构体
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// 定义查询响应结构体
type QueryResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Student `json:"data,omitempty"`
}

// 存储管理器，用于管理学生数据
var (
	students map[string]Student
	mu       sync.Mutex // 用于保护students map的并发访问
)

// 初始化存储管理器
func InitStorage() {
	students = make(map[string]Student)
}

// 添加测试数据
func AddTestData() {
	AddStudent(Student{StudentID: "2024001", Name: "张三", Gender: "男", BirthDate: "2000-01-01"})
	AddStudent(Student{StudentID: "2024002", Name: "李四", Gender: "女", BirthDate: "2000-02-02"})
}

// 添加学生
func AddStudent(student Student) {
	mu.Lock()
	defer mu.Unlock()
	students[student.StudentID] = student
}

// 获取所有学生
func GetAllStudents() []Student {
	mu.Lock()
	defer mu.Unlock()

	var result []Student
	for _, student := range students {
		result = append(result, student)
	}
	return result
}

// 根据ID获取学生
func GetStudentByID(studentID string) (Student, bool) {
	mu.Lock()
	defer mu.Unlock()

	student, exists := students[studentID]
	return student, exists
}

// 检查学生是否存在
func StudentExists(studentID string) bool {
	mu.Lock()
	defer mu.Unlock()

	_, exists := students[studentID]
	return exists
}

// 返回成功响应
func RespondWithSuccess(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Success: true,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// 返回错误响应
func RespondWithError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Success: false,
		Message: message,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}