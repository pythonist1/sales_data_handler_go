package gateway


import (
	"github.com/google/uuid"
)


type GatewayHandler struct {
	task_handler TaskHandlerInterface
}


func (handler *GatewayHandler) HandleFile(ctx *fiber.Ctx) error {

	file, err := c.FormFile("document")
	if err != nil {
		return err
	}

	task_id := uuid.NewString()

	c.SaveFile(file, fmt.Sprintf("./%s", task_id))

	go handler.task_handler.HandleNewTask(task_id)

	response_map := fiber.Map{
		"status": "PENDING",
	  }
	
	return ctx.JSON(response_map)
}


func (handler *GatewayHandler) CheckTaskStatus(ctx *fiber.Ctx) error {
	task_id := ctx.Query("task_id")
	task_status := handler.task_handler.CheckTaskStatus(task_id)

	response_map := fiber.Map{
		"status": "PENDING",
	  }

	return ctx.JSON(response_map)
}


func (h *handler) GetFile(ctx *fiber.Ctx) error {
}
