package fitbit

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	region                   = "ap-northeast-1"
	fitbitTokenParameterName = "AccessToken"
)

var ssmSVC = ssm.New(session.New(), &aws.Config{Region: aws.String(region)})

// FetchToken from ssm parameter
func FetchToken() *string {
	res, err := ssmSVC.GetParameter(
		&ssm.GetParameterInput{
			Name:           aws.String(fitbitTokenParameterName),
			WithDecryption: aws.Bool(true),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return res.Parameter.Value
}
