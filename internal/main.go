package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"dynamodb_practice/internal/library/dynamodb"
	"dynamodb_practice/internal/library/env"
)

type batchHistoryRecord struct {
	BatchName  string `dynamodbav:"batch_name"`
	ExecutedAt string `dynamodbav:"executed_at"`
	Status     string `dynamodbav:"status"`
	ExpiresAt  int64  `dynamodbav:"expires_at"`
}

func main() {
	ctx := context.Background()

	envReader := env.New()

	client, err := dynamodb.NewClient(ctx, envReader)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().UTC()
	record := batchHistoryRecord{
		BatchName:  "sample-batch",
		ExecutedAt: now.Format(time.RFC3339),
		Status:     "SUCCESS",
		ExpiresAt:  now.Add(24 * time.Hour).Unix(),
	}

	if err := client.CreateRecord(ctx, record); err != nil {
		log.Fatal(err)
	}

	log.Println("record created")

	records, err := client.ListRecords(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("records:")
	for _, record := range records {
		fmt.Printf(
			"- batch_name=%v executed_at=%v status=%v \n",
			record["batch_name"],
			record["executed_at"],
			record["status"],
		)
	}
}
