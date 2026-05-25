package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"dynamodb_practice/internal/library/dynamodb"
	"dynamodb_practice/internal/library/env"
)

type BatchHistoryRecord struct {
	BatchName  string `dynamodbav:"batch_name"`
	ExecutedAt string `dynamodbav:"executed_at"`
	Status     string `dynamodbav:"status"`
	ExpiresAt  int64  `dynamodbav:"expires_at"`
}

func main() {
	ctx := context.Background()
	envReader := env.New()

	app, err := envReader.GetEnv("APP")
	if err != nil {
		fmt.Println("error: APP environment variable is not set")
		return
	}
	if app != "local" {
		fmt.Println("error: seeder can only run when APP=local")
		return
	}

	client, err := dynamodb.NewClient(ctx, envReader)
	if err != nil {
		fmt.Printf("error: DynamoDB client の初期化に失敗しました: %s\n", err)
		os.Exit(1)
	}

	if err := seed(ctx, client, seedRecords()); err != nil {
		fmt.Printf("error: データ投入中にエラーが発生しました: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("seed completed for APP=%s\n", app)
}

func seed(ctx context.Context, client dynamodb.ClientInterface, records []BatchHistoryRecord) error {
	for _, record := range records {
		if err := client.CreateRecord(ctx, record); err != nil {
			return fmt.Errorf("batch_name=%s executed_at=%s: %w", record.BatchName, record.ExecutedAt, err)
		}
	}

	return nil
}

func seedRecords() []BatchHistoryRecord {
	expiresAt := time.Date(2099, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()

	return []BatchHistoryRecord{
		{
			BatchName:  "daily-aggregation",
			ExecutedAt: "2026-01-01T00:00:00Z",
			Status:     "SUCCESS",
			ExpiresAt:  expiresAt,
		},
		{
			BatchName:  "daily-aggregation",
			ExecutedAt: "2026-01-02T00:00:00Z",
			Status:     "FAILED",
			ExpiresAt:  expiresAt,
		},
		{
			BatchName:  "billing-export",
			ExecutedAt: "2026-01-01T03:00:00Z",
			Status:     "SUCCESS",
			ExpiresAt:  expiresAt,
		},
	}
}
