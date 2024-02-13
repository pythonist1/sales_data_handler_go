package service


type SalesServiceAdapter interface {
	GetSalesSummByDates(dates []string) map[string]string
}


type FileHandlerAdapter interface {
	ParseFile(file_id string) map[string]string
	CollectFile(map[string]string, file_id string) bool
	DeleteFiles(file_id string) bool
}

