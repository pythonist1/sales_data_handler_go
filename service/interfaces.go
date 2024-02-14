package service


type SalesServiceAdapter interface {
	GetSalesSummByDates(dates []string) map[string]string
}


type FileHandlerAdapter interface {
	ParseFile(file_id string) (map[string]string, error)
	CollectFile(map[string]string, string) bool
	DeleteFiles(file_id string) bool
}

