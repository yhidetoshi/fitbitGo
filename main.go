package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yhidetoshi/fitbitGo/lib"
)

func main() {
	lambda.Start(Handler)
}

// Handler lambda
func Handler() {
	// func main() {

	accessToken := fitbit.FetchToken()
	fitbit.DoActivity(*accessToken)
	fitbit.DoSleep(*accessToken)
}
