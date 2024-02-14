package service


type TaskCase struct {
	task_id string
	err error
}


type TaskHandler struct {
	sales_data_handler SalesDataHandler
	task_channel chan TaskCase
	task_case_collection []*TaskCase
}


// Enums
const (
	Success string = "SUCCESS"
	Pending = "PENDING"
	Error = "ERROR"
)


func (t *TaskHandler) HandleNewTask(task_id string) string {
	t.sales_data_handler.HandleSalesFile(task_id)
}


func (t *TaskHandler) CheckTaskStatus(task_id string) task_status string {
	for _, task_case := range t.task_case_collection {
		if task_case.task_id == task_id {
			if task_case.err == nil {
				task_status = Success
			} else {
				task_status= Error
			}
			return task_status
		}
	}
	task_status = Pending
	return task_status_map
}


func (t *TaskHandler) DeleteTaskCase(task_id string) {
	for index, task_case := range t.task_case_collection {
		if task_case.task_id == task_id {
			t.sales_data_handler.DeleteFiles(task_id)
			t.task_case_collection[index] = t.task_case_collection[len(t.task_case_collection) - 1]
			t.task_case_collection = t.task_case_collection[:len(t.task_case_collection) - 1]
		}
	}
}


func (t *TaskHandler) TaskLoop() {
	for task_case := range t.task_channel {
		t.task_case_collection = append(task_handler.task_case_collection, task_case)
	}
}
