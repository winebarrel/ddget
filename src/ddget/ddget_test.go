package ddget

import (
	. "."
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mockaws"
	"testing"
)

func TestGetString(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mddb := mockaws.NewMockDynamoDBAPI(ctrl)

	mddb.EXPECT().DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String("table")},
	).Return(
		&dynamodb.DescribeTableOutput{
			Table: &dynamodb.TableDescription{
				KeySchema: []*dynamodb.KeySchemaElement{
					&dynamodb.KeySchemaElement{
						AttributeName: aws.String("key"),
						KeyType:       aws.String("HASH")}}}},
		nil,
	)

	mddb.EXPECT().GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String("table"),
			Key: map[string]*dynamodb.AttributeValue{
				"key": {S: aws.String("foo")}}},
	).Return(
		&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"key":   {S: aws.String("foo")},
				"value": {S: aws.String("bar")}}},
		nil,
	)

	ddg := &Ddget{Ddb: mddb}
	item, _ := ddg.GetItem("table", "", "foo")

	assert.Equal("bar", item)
}

func TestGetNumber(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mddb := mockaws.NewMockDynamoDBAPI(ctrl)

	mddb.EXPECT().DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String("table")},
	).Return(
		&dynamodb.DescribeTableOutput{
			Table: &dynamodb.TableDescription{
				KeySchema: []*dynamodb.KeySchemaElement{
					&dynamodb.KeySchemaElement{
						AttributeName: aws.String("key"),
						KeyType:       aws.String("HASH")}}}},
		nil,
	)

	mddb.EXPECT().GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String("table"),
			Key: map[string]*dynamodb.AttributeValue{
				"key": {S: aws.String("foo")}}},
	).Return(
		&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"key":   {S: aws.String("foo")},
				"value": {N: aws.String("1")}}},
		nil,
	)

	ddg := &Ddget{Ddb: mddb}
	item, _ := ddg.GetItem("table", "", "foo")

	assert.Equal("1", item)
}

func TestGetStringWithValueSchema(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mddb := mockaws.NewMockDynamoDBAPI(ctrl)

	mddb.EXPECT().DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String("table")},
	).Return(
		&dynamodb.DescribeTableOutput{
			Table: &dynamodb.TableDescription{
				KeySchema: []*dynamodb.KeySchemaElement{
					&dynamodb.KeySchemaElement{
						AttributeName: aws.String("key"),
						KeyType:       aws.String("HASH")}}}},
		nil,
	)

	mddb.EXPECT().GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String("table"),
			Key: map[string]*dynamodb.AttributeValue{
				"key": {S: aws.String("foo")}}},
	).Return(
		&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"key":    {S: aws.String("foo")},
				"value":  {S: aws.String("bar")},
				"value1": {S: aws.String("zoo")}}},
		nil,
	)

	ddg := &Ddget{Ddb: mddb}
	item, _ := ddg.GetItem("table", "value1", "foo")

	assert.Equal("zoo", item)
}

func TestGetNumberWithValueSchema(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mddb := mockaws.NewMockDynamoDBAPI(ctrl)

	mddb.EXPECT().DescribeTable(
		&dynamodb.DescribeTableInput{
			TableName: aws.String("table")},
	).Return(
		&dynamodb.DescribeTableOutput{
			Table: &dynamodb.TableDescription{
				KeySchema: []*dynamodb.KeySchemaElement{
					&dynamodb.KeySchemaElement{
						AttributeName: aws.String("key"),
						KeyType:       aws.String("HASH")}}}},
		nil,
	)

	mddb.EXPECT().GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String("table"),
			Key: map[string]*dynamodb.AttributeValue{
				"key": {S: aws.String("foo")}}},
	).Return(
		&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"key":    {S: aws.String("foo")},
				"value":  {N: aws.String("1")},
				"value1": {N: aws.String("2")}}},
		nil,
	)

	ddg := &Ddget{Ddb: mddb}
	item, _ := ddg.GetItem("table", "value1", "foo")

	assert.Equal("2", item)
}
