package main

import (
	"github.com/yhidetoshi/fitbitGo/lib"
	"os"
)

var accessToken = os.Getenv("AccessToken")

func main() {
	//fitbit.Activity(accessToken)
	fitbit.FetchSleep(accessToken)
}