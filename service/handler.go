package service


import (
	"sales_data_handler/domain"
)


type SalesDataHandler struct {
	SalesService SalesServiceAdapter
	FileHandler FileHandlerAdapter
	TaskChannel chan TaskCase
}


func (data_handler *SalesDataHandler) HandleSalesFile(task_id string) {
	raw_data, _ := data_handler.FileHandler.ParseFile(task_id)
	sales_data := domain.NewSalesData(raw_data)

	additional_data := data_handler.SalesService.GetSalesSummByDates(sales_data.EmptyDates)
	sales_data.CompleteEmptyDates(additional_data)
	raw_data = sales_data.GetRawData()
	data_handler.FileHandler.CollectFile(raw_data, task_id)
	data_handler.TaskChannel <- TaskCase {
		task_id: task_id,
	}
}


func (data_handler *SalesDataHandler) DeleteRecords(file_id string) {
	data_handler.FileHandler.DeleteFiles(file_id)
}
