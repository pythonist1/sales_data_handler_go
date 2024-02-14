package main

import (
	"sales_data_handler/adapters"
	"sales_data_handler/service"
	"sales_data_handler/gateway"
)


func BootstrapSalesServiceAdapter() adapters.FakeSalesServiceAdapter {

	sales_service_adapter := adapters.FakeSalesServiceAdapter{}
	return sales_service_adapter

}


func BootstrapFileHandler(file_path string) adapters.FileHandler {

	file_handler := adapters.FileHandler{file_path}
	return file_handler
}


func BootstrapSalesDataHandler(
	sales_service adapters.FakeSalesServiceAdapter,
	file_handler adapters.FileHandler,
	task_channel chan service.TaskCase) service.SalesDataHandler {

		sales_data_handler := service.SalesDataHandler{sales_service, file_handler, task_channel}
		return sales_data_handler

}


func BootstrapTaskHandler(sales_data_handler service.SalesDataHandler, task_channel chan service.TaskCase) service.TaskHandler {

	task_handler := service.TaskHandler{
		SalesDataHandler: sales_data_handler,
		TaskChannel: task_channel,
	}
	return task_handler

}


func BootstrapGatewayHandler(file_path string, task_channel chan service.TaskCase) gateway.GatewayHandler {

	sales_service_adapter := BootstrapSalesServiceAdapter()
	file_handler := BootstrapFileHandler(file_path)
	sales_data_handler := BootstrapSalesDataHandler(sales_service_adapter, file_handler, task_channel)
	task_handler := BootstrapTaskHandler(sales_data_handler, task_channel)

	gateway_handler := gateway.GatewayHandler{
		TaskHandler: task_handler,
		FilePath: file_path,
	}
	return gateway_handler

}
