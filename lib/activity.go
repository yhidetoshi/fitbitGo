package fitbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

const urlActive = "https://api.fitbit.com/1/user/-/activities/date/"

// Activities set value
type Activities struct {
	ActiveSummary ActiveSummary `json:"summary"`
}

// ActiveSummary set value
type ActiveSummary struct {
	CaloriesOut int `json:"caloriesOut"`
	Steps       int `json:"steps"`
}

// DoActivity Fetch values by fitbit api
func DoActivity(accessToken string) {
	client := &http.Client{}

	now := time.Now().Format("2006-01-02")
	jst := time.FixedZone(timezone, offset)
	nowPostTime := time.Now().In(jst)

	req, _ := http.NewRequest("GET", urlActive+now+".json", nil)
	req.Header.Set("Authorization", " Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	jsonBytes := ([]byte)(bodyStr)

	s := &Activities{}
	if err = json.Unmarshal(jsonBytes, s); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.ActiveSummary.CaloriesOut, s.ActiveSummary.Steps)

	PostActiveValuesToMackerel(s.ActiveSummary.Steps, s.ActiveSummary.CaloriesOut, nowPostTime)
}

// PostActiveValuesToMackerel Post Metrics to Mackerel
func PostActiveValuesToMackerel(steps int, caloriesOut int, nowTime time.Time) error {
	// Post Steps metrics
	errSteps := client.PostServiceMetricValues(serviceName, []*mackerel.MetricValue{
		&mackerel.MetricValue{
			Name:  "Steps.steps",
			Time:  nowTime.Unix(),
			Value: steps,
		},
	})
	if errSteps != nil {
		fmt.Println(errSteps)
	}

	// Post CaloriesOut metrics
	errCaloriesOut := client.PostServiceMetricValues(serviceName, []*mackerel.MetricValue{
		&mackerel.MetricValue{
			Name:  "CaloriesOut.caloriesOut",
			Time:  nowTime.Unix(),
			Value: caloriesOut,
		},
	})
	if errCaloriesOut != nil {
		fmt.Println(errCaloriesOut)
	}

	return nil
}
