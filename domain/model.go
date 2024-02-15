package domain

import (
	"fmt"
)


type Sale struct {
	Date string
	Summ string
}


type SalesGroup []*Sale


type SalesData struct {
	sales_group SalesGroup
	EmptyDates []string
}


func (sales_data *SalesData) CompleteEmptyDates(additional_data map[string]string) {
	fmt.Println("Sales group len ", len(sales_data.sales_group))
	for date, summ := range additional_data {
		sale := Sale{date, summ}
		sales_data.sales_group = append(sales_data.sales_group, &sale)
	}
	fmt.Println("Sales group len after", len(sales_data.sales_group))
}


func (sales_data *SalesData) GetRawData() map[string]string {
	raw_data := make(map[string]string)
	for _, sale := range sales_data.sales_group {
		raw_data[sale.Date] = sale.Summ
	}
	return raw_data
}


func NewSalesData(raw_data map[string]string) *SalesData {
	sales_group := SalesGroup{}
	empty_dates := []string{}
	for date, summ := range raw_data {
		if summ == "" {
			empty_dates = append(empty_dates, date)
		} else {
			sale := Sale{date, summ}
			sales_group = append(sales_group, &sale)
		}
	}
	return &SalesData{sales_group, empty_dates}
}
