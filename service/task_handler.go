package service


import (
	"time"
	"fmt"
)


type TaskCase struct {
	task_id string
	err error
}


type TaskHandler struct {
	SalesDataHandler *SalesDataHandler
	TaskChannel chan TaskCase
	task_case_collection []*TaskCase
	delete_case_collection []string
}


// Enums
const (
	Success string = "SUCCESS"
	Pending = "PENDING"
	Error = "ERROR"
)


func (t *TaskHandler) HandleNewTask(task_id string) {
	t.SalesDataHandler.HandleSalesFile(task_id)
}


func (t *TaskHandler) CheckTaskStatus(task_id string) string {
	for _, task_case := range t.task_case_collection {
		if task_case.task_id == task_id {
			if task_case.err == nil {
				return Success
			} else {
				return Error
			}
		}
	}
	return Pending
}


func (t *TaskHandler) DeleteTaskCase(task_id string) {
	for _, delete_id := range t.delete_case_collection {
		if delete_id == task_id {
			return
		}
	}

	t.delete_case_collection = append(t.delete_case_collection, task_id)

	time.Sleep(time.Minute * 5)

	for index, task_case := range t.task_case_collection {
		if task_case.task_id == task_id {
			t.SalesDataHandler.DeleteRecords(task_id)
			t.task_case_collection[index] = t.task_case_collection[len(t.task_case_collection) - 1]
			t.task_case_collection = t.task_case_collection[:len(t.task_case_collection) - 1]
		}
	}
}


func (t *TaskHandler) TaskLoop() {
	for task_case := range t.TaskChannel {
		t.task_case_collection = append(t.task_case_collection, &task_case)
	}
}
