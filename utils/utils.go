package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pipeline-prototype/models"
)

var (
	apiTestKey = "887f3f93c4e372163118de9351b382b9"
	url        = fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s", apiTestKey)
)

// GetLatestRates comments
func GetLatestRates() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseBody models.ResponseBody

	err = json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody.Rates["MXN"])
}
