package categories

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/bugfixes/go-bugfixes"
    "github.com/satori/go.uuid"
)

func generateUUID(tag string) (string) {
    id := uuid.NewV5(uuid.NamespaceURL, tag)

    return id.String()
}

func PutTags(tags TagsStruct) {
    bugfixes.Info("Tags", tags)

    svc := dynamodb.New(session.New())

    for i := 0; i < len(tags.Tags); i++ {
        tag := tags.Tags[i]

        input := &dynamodb.PutItemInput{
            Item: map[string]*dynamodb.AttributeValue{
                "Tag": {
                    S: aws.String(tag.Tag),
                },
                "itemId": {
                    S: aws.String(tag.ItemId),
                },
                "tagId": {
                    S: aws.String(generateUUID(tag.Tag)),
                },
            },
            TableName: aws.String("tags"),
        }

        bugfixes.Info("Input", input)

        result, err := svc.PutItem(input)
        if err != nil {
            bugfixes.Error("Add Tag", err)
        }

        bugfixes.Info("Add Item", result)
    }
}