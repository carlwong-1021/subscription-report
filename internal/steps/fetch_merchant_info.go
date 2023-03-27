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
	client       *http.Client
	msEndpoint   string
	mainEndpoint string
}

type InvoiceSettingRes struct {
	Data InvoiceSetting `json:"data"`
}

type InvoiceSetting struct {
	InvoiceAttribute InvoiceAttribute `json:"attributes"`
}

type InvoiceAttribute struct {
	InvoiceType         string `json:"drawee"`
	CompanyName         string `json:"company_name"`
	TaxID               string `json:"company_unified_id"`
	CompanyAddress      string `json:"company_registered_address"`
	InvoiceReceiptEmail string `json:"email"`
}

type MerchantInfoRes struct {
	Data MerchantInfo `json:"data"`
}

type MerchantInfo struct {
	Handle  string `json:"handle"`
	Name    string `json:"name"`
	ShopURL string `json:"brand_home_url"`
}

func NewFetchMerchantInfoStep(client *http.Client, msEndpoint string, mainEndpoint string) steps_interfaces.FetchMerchantInfoStep {
	return &_queryOrderDetailStep{client: client, msEndpoint: msEndpoint, mainEndpoint: mainEndpoint}
}

func (s *_queryOrderDetailStep) Exec(ctx context.Context, reports []*entities.SubscriptionReport, opts *options_subscription_report.ReportOption) ([]*entities.SubscriptionReport, error) {
	for _, report := range reports {
		info, err := s.getMerchantInfo(report.MerchantID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		invoice, err := s.getInvoiceSetting(report.MerchantID)
		report.Handle = info.Handle
		report.ShopName = info.Name
		report.ShopURL = info.ShopURL
		report.InvoiceType = invoice.InvoiceType
		report.CompanyName = invoice.CompanyName
		report.TaxID = invoice.TaxID
		report.CompanyAddress = invoice.CompanyAddress
		report.InvoiceReceiptEmail = invoice.InvoiceReceiptEmail
	}
	return reports, nil
}

func (s *_queryOrderDetailStep) getInvoiceSetting(id string) (*InvoiceAttribute, error) {
	result := &InvoiceSettingRes{}
	response, err := http.Get(s.msEndpoint + "/api/merchants/" + id + "/invoice")
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return &result.Data.InvoiceAttribute, nil
}

func (s *_queryOrderDetailStep) getMerchantInfo(id string) (*MerchantInfo, error) {
	result := &MerchantInfoRes{}
	response, err := http.Get(s.mainEndpoint + "/v1/" + id)
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}
