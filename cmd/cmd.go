package cmd

import (
	"lambdafunction/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func Cmd() {
	lambda.Start(handler.Handler)	
}
