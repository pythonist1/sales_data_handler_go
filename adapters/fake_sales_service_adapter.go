package adapters

import (
	"math/rand"
	"strconv"
)

type FakeSalesServiceAdapter struct {
}


func (adapter FakeSalesServiceAdapter) GetSalesSummByDates(dates []string) map[string]string {
	additional_data = make(map[string]string)

	for _, date := range dates {
		rand_summ := strconv.FormatFloat(1 + rand.Float64() * (10 - 1), 'f', 2, 64)
		additional_data[date] = *rand_summ
	}
	return additional_data
}
