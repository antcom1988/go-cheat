package sampletest

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

//Before
//func GetItems(client *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error){
//	return client.GetItem(&dynamodb.GetItemInput{
//		AttributesToGet:         []*string{aws.String("bar")},
//		Key: map[string]*dynamodb.AttributeValue{
//			"foo":{
//				BOOL: aws.Bool(true),
//			},
//		},
//	})
//}

func GetItems(client dynamodbiface.DynamoDBAPI) (*dynamodb.GetItemOutput, error){
	return client.GetItem(&dynamodb.GetItemInput{
		AttributesToGet:         []*string{aws.String("bar")},
		Key: map[string]*dynamodb.AttributeValue{
			"foo":{
				BOOL: aws.Bool(true),
			},
		},
	})
}