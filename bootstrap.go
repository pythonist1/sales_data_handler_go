package main

import (
	"sales_data_handler_go/adapters"
	"sales_data_handler_go/service"
)

func BootstrapSalesServiceAdapter() {

}


func BootstrapFileHandler(file_path string) {

}


func BootstrapSalesDataHandler(
	sales_service adapter.FakeSalesServiceAdapter,
	file_handler adapter.FileHandlerAdapter,
	task_channel chan) {

}


func BootstrapTaskHandler(SalesDataHandler service.SalesDataHandler, task_channel chan) {

}


func BootstrapGatewayHandler(task_handler service.TaskHandler) {

}
