package Translate

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

func TranslateFunc() {
	InputText := "日本語だよ!こっちは英語"
	InputLang := "ja"
	OutputLang := "en"
	sess := session.Must(session.NewSession())
	trans := translate.New(sess)

	result, err := trans.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(InputLang),
		TargetLanguageCode: aws.String(OutputLang),
		Text:               aws.String(InputText),
	})
	if err != nil {
		log.Print("おかしいぜ")
	}

	fmt.Println(*result.TranslatedText)
}
