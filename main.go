package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"subscription-report/services"
	"subscription-report/steps"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var fromRange, toRange string
	flag.StringVar(&fromRange, "s", "", "from range")
	flag.StringVar(&toRange, "e", "", "to range")
	flag.Parse()

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_PROTO"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	httpClient := &http.Client{}

	reportSteps := []services.Step{}
	queryLatestChangeStep := steps.NewQueryLatestChangeStep(db)
	queryOrderDetailStep := steps.NewQueryOrderDetailStep(os.Getenv("API_URL"), httpClient)
	generateReportStep := steps.NewGenrateReportStep()
	reportSteps = append(reportSteps, queryLatestChangeStep, queryOrderDetailStep, generateReportStep)
	reportService := services.NewReportService(reportSteps)
	reportService.Exec(fromRange, toRange)
}
