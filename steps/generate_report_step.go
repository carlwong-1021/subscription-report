package steps

import (
	"fmt"
	"strconv"
	"subscription-report/services"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type _genrateReportStep struct {
	db *gorm.DB
}

func NewGenrateReportStep() services.Step {
	return &_genrateReportStep{}
}

func (q _genrateReportStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	report, isReport := input.(*[]map[string]any)
	if !isReport {
		panic("input error type")
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	f.SetCellValue("Sheet1", "A1", "Customer Number")
	f.SetCellValue("Sheet1", "B1", "Customer Name")
	for i, item := range *report {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), item["customerNumber"])
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), item["customerName"])
	}

	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println(err)
	}

	return input, option
}
