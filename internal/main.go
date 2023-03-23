package internal

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	developer_api "subscription-report/internal/developer-api"
	"subscription-report/internal/steps"
	"subscription-report/internal/subscription_report"
	options_subscription_report "subscription-report/internal/subscription_report/options"
)

func Exec() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var fromRange, toRange string
	flag.StringVar(&fromRange, "s", "", "from range")
	flag.StringVar(&toRange, "e", "", "to range")
	flag.Parse()
	layout := "2006-01-02"
	from, _ := time.Parse(layout, fromRange)
	to, _ := time.Parse(layout, toRange)

	httpClient := &http.Client{}
	developerApiClient := developer_api.NewDeveloperApiClient()

	options := &options_subscription_report.ReportOption{
		From: from,
		To:   to,
	}
	reportService := subscription_report.NewReportService(options)

	fetchAppSubscriptionStep := steps.NewFetchAppSubscriptionStep(developerApiClient)
	fetchMerchantInfoStep := steps.NewFetchMerchantInfoStep(httpClient, os.Getenv("MBS_API_ENDPOINT"))
	generateReportStep := steps.NewGenrateReportStep()

	reportService.Exec(fetchAppSubscriptionStep, fetchMerchantInfoStep, generateReportStep)
}
