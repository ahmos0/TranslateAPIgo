package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmos0/goLambdaFirst.git/Testgreeting"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/translate"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type Item struct {
	TimeStamp   string `json:"timestamp" dynamodbav:"timestamp"`
	InputTextj  string `json:"inputtext" dynamodbav:"InputText"`
	OutputTextj string `json:"outputtext" dynamodbav:"OutputText"`
}

func ExucuteFunc(request events.APIGatewayProxyRequest) (Response, error) {
	Testgreeting.Hello()
	InputText := request.QueryStringParameters["InputText"]
	InputLang := "ja"
	OutputLang := "en"

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sessi := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamo := dynamodb.New(sessi)
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

	t := time.Now()
	item := Item{
		TimeStamp:   *aws.String(t.String()),
		InputTextj:  *aws.String(InputText),
		OutputTextj: *aws.String(OutputText),
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	tableName := "translate-history"
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = dynamo.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
	}

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
