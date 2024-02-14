package gateway


import (
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
	"sales_data_handler/service"
)


type GatewayHandler struct {
	TaskHandler service.TaskHandler
	FilePath string
}


func (handler *GatewayHandler) HandleFile(ctx *fiber.Ctx) error {

	file, err := ctx.FormFile("document")
	if err != nil {
		return err
	}

	task_id := uuid.NewString()

	ctx.SaveFile(file, handler.FilePath + task_id + ".xlsx")

	go handler.TaskHandler.HandleNewTask(task_id)

	response_map := fiber.Map{
		"status": "PENDING",
		"task_id": task_id,
	  }
	
	return ctx.JSON(response_map)
}


func (handler *GatewayHandler) CheckTaskStatus(ctx *fiber.Ctx) error {
	task_id := ctx.Params("task_id")
	task_status := handler.TaskHandler.CheckTaskStatus(task_id)

	response_map := fiber.Map{
		"status": task_status,
	  }

	return ctx.JSON(response_map)
}


func (h *GatewayHandler) GetResultFile(ctx *fiber.Ctx) error {
	task_id := ctx.Params("task_id")

	go h.TaskHandler.DeleteTaskCase(task_id)

	return ctx.Download(h.FilePath + task_id + "_result.xlsx")
}
