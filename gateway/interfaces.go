package gateway


type TaskHandlerInterface interface {
	HandleNewTask(file_id string) string
	CheckTaskStatus(task_id string) string
	DeleteTaskCase(task_id string)
}
