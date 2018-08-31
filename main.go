package categorizer

import (
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/bugfixes/go-bugfixes"
  "github.com/aws/aws-sdk-go/service/dynamodb/expression"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "fmt"
)

func GetTitles() {
  session, err := session.NewSession(&aws.Config{
    Region: aws.String(os.Getenv("AWS_REGION")),
  })
  if err != nil {
    bugfixes.Error("Session", err)
  }

  svc := dynamodb.New(session)
  filter := expression.Name("itemId").Equal(expression.Value("00a5a300-0456-5f27-8c0b-4d37d8e05b45"))
  projection := expression.NamesList(expression.Name("title"))
  expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
  if err != nil {
    bugfixes.Error("Expression", err)
  }

  params := &dynamodb.ScanInput{
    ExpressionAttributeNames: expr.Names(),
    ExpressionAttributeValues: expr.Values(),
    FilterExpression: expr.Filter(),
    ProjectionExpression: expr.Projection(),
    TableName: aws.String("items"),
  }
  result, err := svc.Scan(params)
  if err != nil {
    bugfixes.Error("Scan", err)
  }

  for _, i := range result.Items {
    item := Item{}

    err = dynamodbattribute.UnmarshalMap(i, &item)
    if err != nil {
      bugfixes.Error("Item Unmarshal", err)
    }

    fmt.Println("Item", item)
  }
}

func main() {

}
