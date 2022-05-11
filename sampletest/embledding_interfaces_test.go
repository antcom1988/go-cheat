package sampletest

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"reflect"
	"testing"
)

type mockDynamoDBCLient struct {
	dynamodbiface.DynamoDBAPI
	output *dynamodb.GetItemOutput
	err error
}

func (m *mockDynamoDBCLient) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.output, m.err
}

func TestGetItems(t *testing.T) {
	exampleErr := errors.New("error")
	exampleResponse := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"foo":{
				BOOL: aws.Bool(false),
			},
		},
	}

	tests := []struct {
		name    string
		client *mockDynamoDBCLient
		response *dynamodb.GetItemOutput
		err error
	}{
		{
			name:     "success",
			client:   &mockDynamoDBCLient{
				output: exampleResponse,
				err: nil,
			},
			response: exampleResponse,
			err:      nil,
		},
		{
			name:     "failed",
			client:   &mockDynamoDBCLient{
				err:         exampleErr,
			},
			err:      exampleErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetItems(tt.client)
			if !reflect.DeepEqual(actual, tt.response){
				t.Errorf("expected (%v), got (%v)", tt.response, actual)
			}
			if !errors.Is(err, tt.err){
				t.Errorf("expected error (%v), got error (%v)", tt.err, err)
			}
		})
	}
}
