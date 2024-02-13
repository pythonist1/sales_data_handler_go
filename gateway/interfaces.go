package api


type TaskHandlerInterface interface {
	HandleNewTask(file_id string) string
	CheckTaskStatus(task_id string) map[string]string
	DeleteTaskCase(task_id string)
}
