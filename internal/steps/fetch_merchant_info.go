package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	steps_interfaces "subscription-report/interfaces/steps"
	"subscription-report/internal/entities"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

type _queryOrderDetailStep struct {
	client      *http.Client
	mbsEndpoint string
}

type MerchantInfoRes struct {
	Data MerchantInfo `json:"data"`
}

type MerchantInfo struct {
	Handle  string `json:"handle"`
	Name    string `json:"name"`
	ShopURL string `json:"brand_home_url"`
	// InvoiceType         string `json:"handle"`
	// CompanyName         string `json:"handle"`
	// TaxID               string `json:"handle"`
	// CompanyAddress      string `json:"handle"`
	// InvoiceReceiptEmail string `json:"handle"`
}

func NewFetchMerchantInfoStep(client *http.Client, mbsEndpoint string) steps_interfaces.FetchMerchantInfoStep {
	return &_queryOrderDetailStep{client: client, mbsEndpoint: mbsEndpoint}
}

func (s *_queryOrderDetailStep) Exec(ctx context.Context, reports []*entities.SubscriptionReport, opts *options_subscription_report.ReportOption) ([]*entities.SubscriptionReport, error) {
	for _, report := range reports {
		info, err := s.getMerchantInfo(report.MerchantID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		report.Handle = info.Handle
		report.ShopName = info.Name
		report.ShopURL = info.ShopURL
	}
	return reports, nil
}

func (s *_queryOrderDetailStep) getMerchantInfo(id string) (MerchantInfo, error) {
	info := MerchantInfo{}
	response, err := http.Get(s.mbsEndpoint + "/v1/" + id)
	if err != nil {
		return info, err
	}

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}

	fmt.Println(info)
	return info, nil
}
