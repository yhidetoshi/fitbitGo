package main

import (
	"os"

	"github.com/yhidetoshi/fitbitGo/lib"
)

var accessToken = os.Getenv("AccessToken")

func main() {
	fitbit.DoActivity(accessToken)
	fitbit.DoSleep(accessToken)
}