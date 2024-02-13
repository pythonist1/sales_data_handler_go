package gateway


import (
	"github.com/google/uuid"
)


type gateway_handler struct {
	task_handler TaskHandlerInterface
}


func (h *handler) HandleFile(ctx *fiber.Ctx) error {

	file, err := c.FormFile("document")
	if err != nil {
		return err
	}

	file_id := uuid.NewString()

	c.SaveFile(file, fmt.Sprintf("./%s", file_id))
	task_id := task_handler.HandleNewTask(file_id)

	return ctx.JSON("123")
}


func (h *handler) CheckTaskStatus(ctx *fiber.Ctx) error {
}


func (h *handler) GetFile(ctx *fiber.Ctx) error {
}
