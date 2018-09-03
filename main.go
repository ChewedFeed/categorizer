package categorizer

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/bugfixes/go-bugfixes"
	"gopkg.in/jdkato/prose.v2"

	"github.com/chewedfeed/categorizer/categories"
)

var AllTitles = AllItems{}
var AllTags = categories.TagsStruct{}

func GetTitles() {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		bugfixes.Error("Session", err)
	}

	svc := dynamodb.New(session)
	filter := expression.Name("itemId").Equal(expression.Value(os.Getenv("FEEDID")))
	projection := expression.NamesList(expression.Name("title"), expression.Name("itemId"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	if err != nil {
		bugfixes.Error("Expression", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("items"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		bugfixes.Error("Scan", err)
	}

	items := []Item{}

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			bugfixes.Error("Item Unmarshal", err)
		}

		bugfixes.Info("I", i)

		items = append(items, item)
	}

	AllTitles.Items = items
}

func ParseTitles() {
	if len(AllTitles.Items) >= 1 {
		for i := 0; i < len(AllTitles.Items); i++ {
		    tags := categories.TagStruct{}
		    item := AllTitles.Items[i]

			doc, err := prose.NewDocument(item.Title)
			if err != nil {
				bugfixes.Fatal("Parse Title", err)
			}

			fmt.Println("--- Tokens ---")
			for _, tok := range doc.Tokens() {
				fmt.Println(tok.Text, tok.Tag, tok.Label)
			}

			fmt.Println("--- Entities ---")
			for _, ent := range doc.Entities() {
				fmt.Println(ent.Text, ent.Label)

				if ent.Label == "GPE" {
                    tags.ItemId = item.ItemId
                    tags.Tag = strings.Trim(ent.Text, " ")

                    AllTags.Tags = append(AllTags.Tags, tags)
				}
			}

			fmt.Println("--- Sentences ---")
			for _, sent := range doc.Sentences() {
				fmt.Println(sent.Text)
			}
		}
	}

	bugfixes.Info("Tags", AllTags)
}

func RunApp() {
	GetTitles()
	ParseTitles()
	categories.PutTags(AllTags)
}

func main() {
	RunApp()
}
