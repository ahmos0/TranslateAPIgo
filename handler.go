package main

import (
	"fmt"
	"log"

	"github.com/ahmos0/goLambdaFirst.git/Testgreeting"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func ExucuteFunc(request events.APIGatewayProxyRequest) (Response, error) {
	Testgreeting.Hello()
	InputText := request.QueryStringParameters["InputText"]
	InputLang := "ja"
	OutputLang := "en"
	sess := session.Must(session.NewSession())
	trans := translate.New(sess)

	log.Println(InputText + "これです")

	result, err := trans.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(InputLang),
		TargetLanguageCode: aws.String(OutputLang),
		Text:               aws.String(InputText),
	})
	if err != nil {
		log.Print("おかしいぜ")
	}

	OutputText := *result.TranslatedText

	fmt.Println(OutputText + "わーい")

	return Response{
		StatusCode: 200,
		Body:       OutputText,
	}, nil
	//Translate.TranslateFunc(events.APIGatewayProxyRequese{})<-引数がおそらくうまくいっていない
	//Translate.TranslateFunc(events.APIGatewayProxyRequest{})

}

func main() {
	lambda.Start(ExucuteFunc)
}
