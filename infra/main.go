package main

import (
	"os"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	_jsii_ "github.com/aws/jsii-runtime-go"
	"github.com/samber/lo"
)

type BatchHistoryStackProps struct {
	awscdk.StackProps
}

func main() {
	defer _jsii_.Close()

	app := awscdk.NewApp(nil)

	NewBatchHistoryStack(app, "BatchHistoryStack", &BatchHistoryStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: lo.ToPtr(os.Getenv("CDK_DEFAULT_REGION")),
			},
		},
	})

	app.Synth(nil)
}

func NewBatchHistoryStack(scope constructs.Construct, id string, props *BatchHistoryStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	table := awsdynamodb.NewTable(stack, lo.ToPtr("BatchHistoryTable"), &awsdynamodb.TableProps{
		TableName: lo.ToPtr("batch-history"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: lo.ToPtr("batch_name"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: lo.ToPtr("executed_at"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: lo.ToPtr("expires_at"),
		RemovalPolicy:       removalPolicyForApp(os.Getenv("APP")),
	})

	table.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: lo.ToPtr("status-index"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: lo.ToPtr("status"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: lo.ToPtr("executed_at"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	awscdk.NewCfnOutput(stack, lo.ToPtr("BatchHistoryTableName"), &awscdk.CfnOutputProps{
		Value: table.TableName(),
	})

	return stack
}

func removalPolicyForApp(app string) awscdk.RemovalPolicy {
	if strings.EqualFold(app, "local") {
		return awscdk.RemovalPolicy_DESTROY
	}

	return awscdk.RemovalPolicy_RETAIN
}
