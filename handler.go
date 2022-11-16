package main

import (
	"github.com/ahmos0/goLambdaFirst.git/Testgreeting"
	"github.com/aws/aws-lambda-go/lambda"
)

func exucuteFunc() {
	Testgreeting.Hello()
}

func main() {
	lambda.Start(exucuteFunc)
}
