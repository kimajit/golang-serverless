package ccitrigger

import (
	"lambdafunction/models"
	"net/http"
	"strings"
)

func PostRequestToCircleCI(url string, params *strings.Reader, circleCiToken string, client models.Client) (string, error) {
	req, err := http.NewRequest("POST", url, params)
	if err != nil {
		return "", err
	}
	req.Header.Add("Circle-Token", circleCiToken)
	req.Header.Add("Content-Type", "application/json")
	stat := "200 OK"
	// // res, err := client.Do(req)
	// if err != nil {
	// 	return "", err
	// }
	// stat := res.Status
	return stat, nil
}
