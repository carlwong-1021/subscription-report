package internal

import (
	"flag"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"subscription-report/internal/services"

	_ "github.com/go-sql-driver/mysql"
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

	httpClient := &http.Client{}
	options := &services.ReportOption{
		From: fromRange,
		To:   toRange,
	}
	reportService := services.NewReportService(options)

	reportService.Exec()
}
