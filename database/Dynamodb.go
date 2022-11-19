package database

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	TimeStamp   string `json:"timestamp" dynamodbav:"timestamp"`
	InputTextj  string `json:"inputtext" dynamodbav:"InputText"`
	OutputTextj string `json:"outputtext" dynamodbav:"OutputText"`
}

func OperateDB(InputText string, OutputText string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	fmt.Println(InputText)
	fmt.Println(OutputText)
	dynamo := dynamodb.New(sess)

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

}
