package steps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"subscription-report/models"
	"subscription-report/services"
)

type _queryOrderDetailStep struct {
	url    string
	client *http.Client
}

func NewQueryOrderDetailStep(url string, client *http.Client) services.Step {
	return &_queryOrderDetailStep{url: url, client: client}
}

func (q _queryOrderDetailStep) Exec(input any, option *services.ReportOption) (any, *services.ReportOption) {
	orders, isOrders := input.(*[]models.Orders)
	if !isOrders {
		panic("input error type")
	}
	ids := make([]string, 0)
	for _, order := range *orders {
		ids = append(ids, strconv.Itoa(int(order.CustomerNumber)))
	}

	var response []map[string]any
	res, err := http.Get(q.url + "?ids=" + strings.Join(ids, ","))
	// TODO: batch process (redis)
	// handle batch fail

	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// report := make([]entities.Report, 0)
	// for _, item := range response {
	// 	report = append(report, entities.Report{
	// 		CustomerNumber: item["customerNumber"].(string),
	// 		CustomerName:   item["customerName"].(string),
	// 		Phone:          item["Phone"].(string),
	// 		City:           item["City"].(string),
	// 	})
	// }

	return &response, option
}
