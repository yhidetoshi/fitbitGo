package fitbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

const (
	urlSleep = "https://api.fitbit.com/1.2/user/-/sleep/date/"
	hour = 60.0
)

// Sleep set value
type Sleep struct {
	SleepSummary SleepSummary `json:"summary"`
}

// SleepSummary set value
type SleepSummary struct {
	TotalMinutesAsleep float64 `json:"totalMinutesAsleep"`
	TotalSleepRecords  int     `json:"totalSleepRecords"`
	TotalTimeInBed     int     `json:"totalTimeInBed"`
}

// DoSleep Fetch values by fitbit api
func DoSleep(accessToken string) {
	client := &http.Client{}
	now := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	jst := time.FixedZone(timezone, offset)
	nowPostTime := time.Now().In(jst)

	req, _ := http.NewRequest("GET", urlSleep+now+".json", nil)
	req.Header.Set("Authorization", " Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	jsonBytes := ([]byte)(bodyStr)

	s := &Sleep{}
	if err = json.Unmarshal(jsonBytes, s); err != nil {
		fmt.Println(err)
	}

	// fmt.Println(s.SleepSummary.TotalTimeInBed)
	hourTotalTimeInBed := float32(s.SleepSummary.TotalTimeInBed) / hour
	TotalTimeInBed := fmt.Sprintf("%.2f", hourTotalTimeInBed)
	float64TotalTimeInBed, _ := strconv.ParseFloat(TotalTimeInBed, 64)

	err = PostSleepValuesToMackerel(float64TotalTimeInBed, nowPostTime)
	if err != nil {
		fmt.Println(err)
	}
}

// PostSleepValuesToMackerel Post Metrics to Mackerel
func PostSleepValuesToMackerel(TotalTimeInBed float64, nowTime time.Time) error {
	// Post Steps metrics
	errSteps := client.PostServiceMetricValues(serviceName, []*mackerel.MetricValue{
		&mackerel.MetricValue{
			Name:  "TotalTimeInBed.totalTimeInBed",
			Time:  nowTime.Unix(),
			Value: TotalTimeInBed,
		},
	})
	if errSteps != nil {
		fmt.Println(errSteps)
	}

	return nil
}
