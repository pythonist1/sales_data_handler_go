package service


import (
	"sales_data_handler_go/domain"
)


type SalesDataHandler struct {
	sales_service SalesServiceAdapter
	file_handler FileHandlerAdapter
	task_channel chan TaskCase
}


func (data_handler *SalesDataHandler) HandleSalesFile(file_id string, task_id string) {
	raw_data := data_handler.file_handler.ParseFile(file_id)
	sales_data := domain.NewSalesData(raw_data)

	additional_data := data_handler.sales_service.GetSalesSummByDates(sales_data.empty_dates)
	sales_data.CompleteEmptyDates(additional_data)
	raw_data = sales_data.GetRawData()
	data_handler.file_handler.CollectFile(raw_data, file_id)
	data_handler.task_channel <- TaskCase {
		task_id: task_id
		file_id: file_id
	}
}


func (data_handler *SalesDataHandler) DeleteRecords(file_id string) {
	data_handler.file_handler.DeleteFiles(file_id)
}
