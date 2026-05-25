package dynamodb

import "context"

type ClientInterface interface {
	CreateRecord(ctx context.Context, item any) error
	ListRecords(ctx context.Context) ([]map[string]any, error)
}
