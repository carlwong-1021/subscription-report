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

func (s *_genrateReportStep) Exec(reports []*entities.SubscriptionReport, option *options_subscription_report.ReportOption) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	for char, i := 'A', 0; i < len(constants.EnFieldName); char, i = char+1, i+1 {
		f.SetCellValue("Sheet1", string(char)+"1", constants.EnFieldName[0])
		f.SetCellValue("Sheet1", string(char)+"2", constants.ChFieldName[0])
	}

	for i, report := range reports {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+3), report.SubscriptionDate)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+3), report.SubscriptionSource)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+3), report.AppName)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+3), report.AppPrice)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+3), report.AppDeveloper)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+3), report.MerchantID)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+3), report.Handle)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+3), report.ShopName)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+3), report.ShopURL)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+3), report.SubscriptionCycle)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+3), report.SubscriptionFee)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+3), report.OneOffActivationFee)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+3), report.SubscriptionCurrency)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+3), report.BillNumber)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+3), report.OrderId)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+3), report.InvoiceType)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+3), report.CompanyName)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+3), report.TaxID)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+3), report.CompanyAddress)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+3), report.InvoiceReceiptEmail)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+3), report.DeveloperRevenue)
		f.SetCellValue("Sheet1", "V"+strconv.Itoa(i+3), report.RebateAccountDate)
		f.SetCellValue("Sheet1", "W"+strconv.Itoa(i+3), report.RebateRatio)
		f.SetCellValue("Sheet1", "X"+strconv.Itoa(i+3), report.RebateAmount)
		f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i+3), report.RebatePaymentDate)
		f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i+3), report.InvoiceIssueDate)
	}

	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println(err)
	}
}
