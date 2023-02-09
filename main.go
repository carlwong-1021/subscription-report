package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"subscription-report/services"
	"subscription-report/steps"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	httpClient := &http.Client{}

	reportSteps := []services.Step{}
	queryLatestChangeStep := steps.NewQueryLatestChangeStep(db)
	queryOrderDetailStep := steps.NewQueryOrderDetailStep(os.Getenv("API_URL"), httpClient)
	reportSteps = append(reportSteps, queryLatestChangeStep, queryOrderDetailStep)
	reportService := services.NewReportService(reportSteps)
	report := reportService.Exec(fromRange, toRange)

	fmt.Println(report)
}
