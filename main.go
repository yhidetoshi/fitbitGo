package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yhidetoshi/fitbitGo/lib"
)

var accessToken = os.Getenv("ACCESSTOKEN")

func main() {
	lambda.Start(Handler)
}

func Handler() {

//func main() {
	fitbit.DoActivity(accessToken)
	fitbit.DoSleep(accessToken)
}