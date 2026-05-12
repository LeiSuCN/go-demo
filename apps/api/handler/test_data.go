package handler

import "go-demo/pkgs/models"

func AddTestData() {
	mu.Lock()
	defer mu.Unlock()
	store["1"] = models.Student{ID: "1", Name: "张三", Age: 20, Grade: "大一"}
	store["2"] = models.Student{ID: "2", Name: "李四", Age: 21, Grade: "大二"}
	store["3"] = models.Student{ID: "3", Name: "王五", Age: 22, Grade: "大三"}
}
