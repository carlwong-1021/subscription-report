package steps

import (
	"fmt"
	"strconv"
	"subscription-report/interfaces/constants"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"

	"github.com/xuri/excelize/v2"
)

type _genrateReportStep struct {
}

func NewGenrateReportStep() steps_interfaces.GenerateReportStep {
	return &_genrateReportStep{}
}

func getColumnAlph(col int) string {
	name := make([]byte, 0, 3) // max 16,384 columns (2022)
	const aLen = 'Z' - 'A' + 1 // alphabet length
	for ; col > 0; col /= aLen + 1 {
		name = append(name, byte('A'+(col-1)%aLen))
	}
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}

func (s *_genrateReportStep) Exec(reports []*entities.SubscriptionReport, option *options_subscription_report.ReportOption) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	for i, definition := range constants.FieldDefinitions {
		f.SetCellValue("Sheet1", getColumnAlph(i+1)+"1", definition.EN)
		f.SetCellValue("Sheet1", getColumnAlph(i+1)+"2", definition.CH)
	}

	for i, report := range reports {
		for j, definition := range constants.FieldDefinitions {
			if err := f.SetCellValue("Sheet1", getColumnAlph(j+1)+strconv.Itoa(i+3), definition.GetValue(report)); err != nil {
				fmt.Println("Error when set cell value: ", err)
			}
		}
	}

	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println("Error when save excel file: ", err)
	}
}
