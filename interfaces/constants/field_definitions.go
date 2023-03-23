package constants

import (
	"subscription-report/internal/entities"
	"time"
)

func formatDate(target *time.Time) string {
	if target == nil {
		return ""
	}
	return target.Format("20060102")
}

var FieldDefinitions []entities.FieldDefinition = []entities.FieldDefinition{
	{func(report *entities.SubscriptionReport) any { return formatDate(report.SubscriptionDate) }, "Subscription Date", "訂閱日期"},
	{func(report *entities.SubscriptionReport) any { return report.SubscriptionSource }, "Subscription Source", "訂閱來源"},
	{func(report *entities.SubscriptionReport) any { return report.AppName }, "App Name", "App 名稱"},
	{func(report *entities.SubscriptionReport) any { return report.AppPrice }, "App Price", "App 價格設定"},
	{func(report *entities.SubscriptionReport) any { return report.AppDeveloper }, "App Developer Name", "App 開發者"},
	{func(report *entities.SubscriptionReport) any { return report.MerchantID }, "Merchant id", "Merchant id"},
	{func(report *entities.SubscriptionReport) any { return report.Handle }, "Merchant handle", "handle"},
	{func(report *entities.SubscriptionReport) any { return report.ShopName }, "Shop Name", "商店名稱"},
	{func(report *entities.SubscriptionReport) any { return report.ShopURL }, "Shop Address", "商店網址"},
	{func(report *entities.SubscriptionReport) any { return report.SubscriptionCycle }, "Subscription Cycle", "訂閱週期"},
	{func(report *entities.SubscriptionReport) any { return report.SubscriptionFee }, "Subscription Fee", "訂閱費"},
	{func(report *entities.SubscriptionReport) any { return report.OneOffActivationFee }, "One-off activation fee", "一次性開通費"},
	{func(report *entities.SubscriptionReport) any { return report.SubscriptionCurrency }, "Merchant Subscription Currency", "幣別"},
	{func(report *entities.SubscriptionReport) any { return report.BillNumber }, "Billing Number", "帳單號碼"},
	{func(report *entities.SubscriptionReport) any { return report.OrderId }, "Order ID", "帳單參考編號"},
	{func(report *entities.SubscriptionReport) any { return report.InvoiceType }, "Invoice Type", "發票類型"},
	{func(report *entities.SubscriptionReport) any { return report.CompanyName }, "Company Name", "公司名稱"},
	{func(report *entities.SubscriptionReport) any { return report.TaxID }, "Tax ID Number", "統一編號"},
	{func(report *entities.SubscriptionReport) any { return report.CompanyAddress }, "Company Contact Address", "公司聯絡地址"},
	{func(report *entities.SubscriptionReport) any { return report.InvoiceReceiptEmail }, "Invoice Receipt Email", "發票收件電郵"},
	{func(report *entities.SubscriptionReport) any { return report.DeveloperRevenue }, "", "開發者收入"},
	{func(report *entities.SubscriptionReport) any { return formatDate(report.RebateAccountDate) }, "", "回饋金結算日期"},
	{func(report *entities.SubscriptionReport) any { return report.RebateRatio }, "", "回饋金分潤比例"},
	{func(report *entities.SubscriptionReport) any { return report.RebateAmount }, "", "回饋金應付金額"},
	{func(report *entities.SubscriptionReport) any { return formatDate(report.RebatePaymentDate) }, "", "回饋金匯款日期"},
	{func(report *entities.SubscriptionReport) any { return formatDate(report.InvoiceIssueDate) }, "", "開立發票日期"},
}
