package main

import (
	"sales_data_handler/service"
	"github.com/gofiber/fiber/v2"
)


func main() {
	task_channel := make(chan service.TaskCase, 100)
	config := Config{
		FilePath: "./file_path/",
	}
	gateway_handler := BootstrapGatewayHandler(config.FilePath, task_channel)
	go gateway_handler.TaskHandler.TaskLoop()
	gateway_app := fiber.New()
	gateway_app.Post("/handle_file", gateway_handler.HandleFile)
	gateway_app.Get("/check_task_status/:task_id", gateway_handler.CheckTaskStatus)
	gateway_app.Get("/get_result/:task_id", gateway_handler.GetResultFile)
	gateway_app.Listen(":8080")
}
