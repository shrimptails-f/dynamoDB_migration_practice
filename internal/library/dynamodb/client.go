package dynamodb

import (
	"context"

	"dynamodb_practice/internal/library/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const defaultTableName = "batch-history"

var _ ClientInterface = (*Client)(nil)

type Client struct {
	dynamoDB  *dynamodb.Client
	tableName string
}

func NewClient(ctx context.Context, env env.EnvInterface) (*Client, error) {
	region, err := env.GetEnv("AWS_DEFAULT_REGION")
	if err != nil {
		return nil, err
	}

	loadOptions := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}

	app, err := env.GetEnv("APP")
	if err == nil && app == "local" {
		endpointURL, err := env.GetEnv("AWS_ENDPOINT_URL")
		if err != nil {
			return nil, err
		}
		loadOptions = append(loadOptions, config.WithBaseEndpoint(endpointURL))
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		loadOptions...,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		dynamoDB:  dynamodb.NewFromConfig(cfg),
		tableName: defaultTableName, // pocなので定数にする
	}, nil
}

// CreateRecord はレコードを作成します。
func (c *Client) CreateRecord(ctx context.Context, item any) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = c.dynamoDB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(c.tableName),
		Item:      av,
	})

	return err
}

// ListRecords Pocで作成したレコードを表示するための関数
func (c *Client) ListRecords(ctx context.Context) ([]map[string]any, error) {
	output, err := c.dynamoDB.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(c.tableName),
	})
	if err != nil {
		return nil, err
	}

	var records []map[string]any
	if err := attributevalue.UnmarshalListOfMaps(output.Items, &records); err != nil {
		return nil, err
	}

	return records, nil
}
