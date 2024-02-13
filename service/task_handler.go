package service


import (
	"github.com/google/uuid"
)


type TaskCase struct {
	task_id string
	file_id string
	err error
}


type TaskHandler struct {
	sales_data_handler SalesDataHandler
	task_channel chan TaskCase
	task_case_collection []*TaskCase
}


func (task_handler *TaskHandler) HandleNewTask(file_id string) string {
	task_id := uuid.NewString()
	go data_handler.HandleSalesFile(file_id, task_id)
	return task_id
}


func (task_handler *TaskHandler) CheckTaskStatus(task_id string) map[string]string {
	task_status_map = make(map[string]string)
	for _, task_case := range task_handler.task_case_collection {
		if task_case.task_id == task_id {
			if task_case.err == nil {
				task_status_map["status"] = "SUCCESS"
			} else {
				task_status_map["status"] = "ERROR"
			}
			return task_status_map
		}
	}
	task_status_map["status"] = "PENDING"
	return task_status_map
}


func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func (task_handler *TaskHandler) DeleteTaskCase(task_id string) {
	for index, task_case := range task_handler.task_case_collection {
		if task_case.task_id == task_id {
			data_handler.DeleteFiles(task_case.file_id)
			task_handler.task_case_collection[index] = task_handler.task_case_collection[len(task_handler.task_case_collection) - 1]
			task_handler.task_case_collection = task_handler.task_case_collection[:len(task_handler.task_case_collection) - 1]
		}
	}
}


func (task_handler *TaskHandler) TaskLoop() {
	for task_case := range task_handler.task_channel {
		task_handler.task_case_collection = append(task_handler.task_case_collection, task_case)
	}
}
