package ddget

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Ddget struct {
	Ddb dynamodbiface.DynamoDBAPI
}

func New() (ddget *Ddget) {
	return &Ddget{
		Ddb: dynamodb.New(session.New()),
	}
}

func (ddg *Ddget) getKeyAttrName(table string) (keyAttrName string, err error) {
	svc := ddg.Ddb

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String(table),
	}

	resp, err := svc.DescribeTable(params)

	if err != nil {
		return
	}

	if len(resp.Table.KeySchema) > 1 {
		err = fmt.Errorf("Multiple key is not supported")
		return
	}

	keyAttrName = *resp.Table.KeySchema[0].AttributeName

	return
}

func (ddg *Ddget) getItemByKey(table string, keyAttrName string, valueAttrName string, key string) (item string, err error) {
	svc := ddg.Ddb

	params := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			keyAttrName: {S: aws.String(key)},
		},
	}

	resp, err := svc.GetItem(params)

	if err != nil {
		return
	}

	if resp.Item == nil {
		err = fmt.Errorf("Item not found")
		return
	}

	var attrVal *dynamodb.AttributeValue

	for k, v := range resp.Item {
		if valueAttrName != "" {
			if k == valueAttrName {
				attrVal = v
			}
		} else {
			if k != keyAttrName {
				attrVal = v
			}
		}

		if attrVal != nil {
			if attrVal.S != nil {
				item = *attrVal.S
			} else if attrVal.N != nil {
				item = *attrVal.N
			} else {
				err = fmt.Errorf("Unsupported item value type")
			}

			return
		}
	}

	err = fmt.Errorf("Item value not found")

	return
}

func (ddg *Ddget) GetItem(table string, valueAttrName string, key string) (item string, err error) {
	keyAttrName, err := ddg.getKeyAttrName(table)

	if err != nil {
		return
	}

	item, err = ddg.getItemByKey(table, keyAttrName, valueAttrName, key)

	return
}
