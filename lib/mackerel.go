package fitbit

import (
	"os"

	"github.com/mackerelio/mackerel-client-go"
)

var (
	mkrKey = os.Getenv("MKRKEY")
	client = mackerel.NewClient(mkrKey)
)

const (
	serviceName = "Fitbit"
	timezone    = "Asia/Tokyo"
	offset      = 9 * 60 * 60
)
