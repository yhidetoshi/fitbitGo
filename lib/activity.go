package fitbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//var accessToken = os.Getenv("AccessToken")

const (
	urlActive = "https://api.fitbit.com/1/user/-/activities/date/"
)

func Activity(accessToken string) {
	client := &http.Client{}
	now := time.Now().Format("2006-01-02")

	req, _ := http.NewRequest("GET", urlActive+now+".json", nil)
	req.Header.Set("Authorization", " Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	//fmt.Println(bodyStr)

	jsonBytes := ([]byte)(bodyStr)

	s := &Activities{}
	if err = json.Unmarshal(jsonBytes, s); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.ActiveSummary.CaloriesOut)
	fmt.Println(s.ActiveSummary.Steps)
}

type Activities struct {
	ActiveSummary ActiveSummary `json:"summary"`
}

type ActiveSummary struct {
	CaloriesOut          int     `json:"caloriesOut"`
	Steps                int     `json:"steps"`
}
