package steps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
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
	rt := reflect.TypeOf(input)
	if rt.String() != "*[]models.Orders" {
		fmt.Println("error")
	}
	orders, _ := input.(*[]models.Orders)
	ids := make([]string, 0)
	for _, order := range *orders {
		ids = append(ids, strconv.Itoa(int(order.CustomerNumber)))
	}

	var response []map[string]any
	res, err := http.Get(q.url + "?ids=" + strings.Join(ids, ","))

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

	return response, option
}