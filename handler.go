package main

import (
	"github.com/ahmos0/goLambdaFirst.git/Testgreeting"
	"github.com/ahmos0/goLambdaFirst.git/Translate"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func ExucuteFunc(request events.APIGatewayProxyRequest) (Response, error) {
	Testgreeting.Hello()
	OutputText := Translate.TranslateFunc(request)
	return Response{
		StatusCode: 200,
		Body:       OutputText,
	}, nil
}

func main() {
	lambda.Start(ExucuteFunc)
}
