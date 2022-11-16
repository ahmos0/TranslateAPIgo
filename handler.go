package main

import (
	"github.com/ahmos0/goLambdaFirst.git/Testgreeting"
	"github.com/ahmos0/goLambdaFirst.git/Translate"
	"github.com/aws/aws-lambda-go/lambda"
)

func exucuteFunc() {
	Testgreeting.Hello()

	Translate.TranslateFunc()

}

func main() {
	lambda.Start(exucuteFunc)
}
