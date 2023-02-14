package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"))
	fmt.Println("Start MongoDB connection: ", uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	database := client.Database(os.Getenv("MONGO_DATABASE"))

	httpClient := &http.Client{}

	reportSteps := []services.Step{}
	queryLatestChangeStep := steps.NewQueryLatestChangeStep(database)
	queryOrderDetailStep := steps.NewQueryOrderDetailStep(os.Getenv("API_URL"), httpClient)
	generateReportStep := steps.NewGenrateReportStep()
	reportSteps = append(reportSteps, queryLatestChangeStep, queryOrderDetailStep, generateReportStep)
	reportService := services.NewReportService(reportSteps)
	reportService.Exec(fromRange, toRange)
}
