package Translate

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

// Responseの構造体を作成
type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

var OutputText string

func TranslateFunc(req events.APIGatewayProxyRequest) (Response, error) {
	InputText := req.QueryStringParameters["InputText"]
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
}
