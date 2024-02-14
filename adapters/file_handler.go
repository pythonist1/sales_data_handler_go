package adapters


import (
    "fmt"
    "github.com/xuri/excelize/v2"
	"os"
)


type FileHandler struct {
	FilePath string
}


func (file_handler FileHandler) ParseFile(file_id string) (map[string]string, error) {
	raw_data := make(map[string]string)

	path := file_handler.FilePath + file_id + ".xlsx"
    f, err := excelize.OpenFile(path)
    if err != nil {
        return raw_data, err
    }
    defer func() {
        if err := f.Close(); err != nil {
            return
        }
    }()

	rows, err := f.GetRows("Sheet1")
    if err != nil {
        return nil, err
    }

    for row_index, row := range rows {
		if row_index == 0 {
			continue
		}

		var date string
		var summ string

        for index, colCell := range row {
			if index > 1 {
				break
			} else if index == 0 {
				date = colCell
			} else if index == 1 {
				summ = colCell
			}
        }

		raw_data[date] = summ
    }
	return raw_data, err

}


func (file_handler FileHandler) CollectFile(raw_data map[string]string, file_id string) bool {
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            return
        }
    }()

	column_names := []interface{}{"Date", "Summ"}

	excel_list := [][]interface{}{}
	excel_list = append(excel_list, column_names)

	for date, summ := range raw_data {
		excel_list = append(excel_list, []interface{}{date, summ})
	}


	for index, row := range excel_list {
        cell, err := excelize.CoordinatesToCellName(1, index+1)
        if err != nil {
        }
		f.SetSheetRow("Sheet1", cell, &row)
	}
    if err := f.SaveAs(file_handler.FilePath + file_id + "_result.xlsx"); err != nil {
        fmt.Println(err)
    }
	return true
}


func (file_handler FileHandler) DeleteFiles(file_id string) bool {

	err := os.Remove(file_handler.FilePath + file_id + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}

	err_2 := os.Remove(file_handler.FilePath + file_id + "_result.xlsx")
	if err_2 != nil {
		fmt.Println(err_2)
	}

	return true

}
