package steps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	comments, isComments := input.([]models.Comment)
	if !isComments {
		panic("wrong input type")
	}
	ids := make([]string, 0)
	for _, comment := range comments {
		ids = append(ids, comment.MovieId.Hex())
	}

	var result []map[string]any

	for start, end := 0, 0; start <= len(ids)-1; start = end {
		end = start + option.BatchSize
		if end > len(ids) {
			end = len(ids)
		}
		batch := ids[start:end]
		response, err := q.fetchData(batch)
		if err != nil {
			fmt.Println("fetch api error: ", err)
			break
		}
		result = append(result, response...)
	}

	return result, option
}

func (q _queryOrderDetailStep) fetchData(ids []string) ([]map[string]any, error) {
	var response []map[string]any
	res, err := http.Get(q.url + "?ids=" + strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return response, nil
}
