package userlist

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type mockError struct {
	msg  string
	code string
}

func (t *mockError) Error() string {
	return t.msg
}

// Returns the short phrase depicting the classification of the error.
func (t *mockError) Code() string {
	return t.code
}

// Returns the error details message.
func (t *mockError) Message() string {
	panic("not implemented") // TODO: Implement
}

// Returns the original error if one was set.  Nil is returned if not set.
func (t *mockError) OrigErr() error {
	panic("not implemented") // TODO: Implement
}

type DynamoDBMock struct {
	GetItemFn func(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	PutItemFn func(_ *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func (d *DynamoDBMock) BatchExecuteStatement(_ *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchExecuteStatementWithContext(_ aws.Context, _ *dynamodb.BatchExecuteStatementInput, _ ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchExecuteStatementRequest(_ *dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchGetItem(_ *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchGetItemWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchGetItemRequest(_ *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchGetItemPages(_ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchGetItemPagesWithContext(_ aws.Context, _ *dynamodb.BatchGetItemInput, _ func(*dynamodb.BatchGetItemOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchWriteItem(_ *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchWriteItemWithContext(_ aws.Context, _ *dynamodb.BatchWriteItemInput, _ ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) BatchWriteItemRequest(_ *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateBackup(_ *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateBackupWithContext(_ aws.Context, _ *dynamodb.CreateBackupInput, _ ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateBackupRequest(_ *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateGlobalTable(_ *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateGlobalTableWithContext(_ aws.Context, _ *dynamodb.CreateGlobalTableInput, _ ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateGlobalTableRequest(_ *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateTable(_ *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateTableWithContext(_ aws.Context, _ *dynamodb.CreateTableInput, _ ...request.Option) (*dynamodb.CreateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) CreateTableRequest(_ *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteBackup(_ *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteBackupWithContext(_ aws.Context, _ *dynamodb.DeleteBackupInput, _ ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteBackupRequest(_ *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteItem(_ *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteItemWithContext(_ aws.Context, _ *dynamodb.DeleteItemInput, _ ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteItemRequest(_ *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteTable(_ *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteTableWithContext(_ aws.Context, _ *dynamodb.DeleteTableInput, _ ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DeleteTableRequest(_ *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeBackup(_ *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeBackupWithContext(_ aws.Context, _ *dynamodb.DescribeBackupInput, _ ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeBackupRequest(_ *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContinuousBackups(_ *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.DescribeContinuousBackupsInput, _ ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContinuousBackupsRequest(_ *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContributorInsights(_ *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContributorInsightsWithContext(_ aws.Context, _ *dynamodb.DescribeContributorInsightsInput, _ ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeContributorInsightsRequest(_ *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeEndpoints(_ *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeEndpointsWithContext(_ aws.Context, _ *dynamodb.DescribeEndpointsInput, _ ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeEndpointsRequest(_ *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeExport(_ *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeExportWithContext(_ aws.Context, _ *dynamodb.DescribeExportInput, _ ...request.Option) (*dynamodb.DescribeExportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeExportRequest(_ *dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTable(_ *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTableWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTableRequest(_ *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTableSettings(_ *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.DescribeGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeGlobalTableSettingsRequest(_ *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeImport(_ *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeImportWithContext(_ aws.Context, _ *dynamodb.DescribeImportInput, _ ...request.Option) (*dynamodb.DescribeImportOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeImportRequest(_ *dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeKinesisStreamingDestination(_ *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.DescribeKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeKinesisStreamingDestinationRequest(_ *dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeLimits(_ *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeLimitsWithContext(_ aws.Context, _ *dynamodb.DescribeLimitsInput, _ ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeLimitsRequest(_ *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTable(_ *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTableWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTableRequest(_ *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTableReplicaAutoScaling(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.DescribeTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTableReplicaAutoScalingRequest(_ *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTimeToLive(_ *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTimeToLiveWithContext(_ aws.Context, _ *dynamodb.DescribeTimeToLiveInput, _ ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DescribeTimeToLiveRequest(_ *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DisableKinesisStreamingDestination(_ *dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DisableKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.DisableKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) DisableKinesisStreamingDestinationRequest(_ *dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) EnableKinesisStreamingDestination(_ *dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) EnableKinesisStreamingDestinationWithContext(_ aws.Context, _ *dynamodb.EnableKinesisStreamingDestinationInput, _ ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) EnableKinesisStreamingDestinationRequest(_ *dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteStatement(_ *dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteStatementWithContext(_ aws.Context, _ *dynamodb.ExecuteStatementInput, _ ...request.Option) (*dynamodb.ExecuteStatementOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteStatementRequest(_ *dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteTransaction(_ *dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteTransactionWithContext(_ aws.Context, _ *dynamodb.ExecuteTransactionInput, _ ...request.Option) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExecuteTransactionRequest(_ *dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExportTableToPointInTime(_ *dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExportTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.ExportTableToPointInTimeInput, _ ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ExportTableToPointInTimeRequest(_ *dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) GetItemWithContext(ctx aws.Context, in *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	if f := d.GetItemFn; f != nil {
		return f(in)
	}

	panic("not assigned")
}

func (d *DynamoDBMock) GetItemRequest(_ *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ImportTable(_ *dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ImportTableWithContext(_ aws.Context, _ *dynamodb.ImportTableInput, _ ...request.Option) (*dynamodb.ImportTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ImportTableRequest(_ *dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListBackups(_ *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListBackupsWithContext(_ aws.Context, _ *dynamodb.ListBackupsInput, _ ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListBackupsRequest(_ *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListContributorInsights(_ *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListContributorInsightsWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListContributorInsightsRequest(_ *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListContributorInsightsPages(_ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListContributorInsightsPagesWithContext(_ aws.Context, _ *dynamodb.ListContributorInsightsInput, _ func(*dynamodb.ListContributorInsightsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListExports(_ *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListExportsWithContext(_ aws.Context, _ *dynamodb.ListExportsInput, _ ...request.Option) (*dynamodb.ListExportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListExportsRequest(_ *dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListExportsPages(_ *dynamodb.ListExportsInput, _ func(*dynamodb.ListExportsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListExportsPagesWithContext(_ aws.Context, _ *dynamodb.ListExportsInput, _ func(*dynamodb.ListExportsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListGlobalTables(_ *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListGlobalTablesWithContext(_ aws.Context, _ *dynamodb.ListGlobalTablesInput, _ ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListGlobalTablesRequest(_ *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListImports(_ *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListImportsWithContext(_ aws.Context, _ *dynamodb.ListImportsInput, _ ...request.Option) (*dynamodb.ListImportsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListImportsRequest(_ *dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListImportsPages(_ *dynamodb.ListImportsInput, _ func(*dynamodb.ListImportsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListImportsPagesWithContext(_ aws.Context, _ *dynamodb.ListImportsInput, _ func(*dynamodb.ListImportsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTables(_ *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTablesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ ...request.Option) (*dynamodb.ListTablesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTablesRequest(_ *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTablesPages(_ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTablesPagesWithContext(_ aws.Context, _ *dynamodb.ListTablesInput, _ func(*dynamodb.ListTablesOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTagsOfResource(_ *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTagsOfResourceWithContext(_ aws.Context, _ *dynamodb.ListTagsOfResourceInput, _ ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ListTagsOfResourceRequest(_ *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) PutItem(_ *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) PutItemWithContext(_ aws.Context, in *dynamodb.PutItemInput, _ ...request.Option) (*dynamodb.PutItemOutput, error) {
	if f := d.PutItemFn; f != nil {
		return f(in)
	}
	panic("not assigned")
}

func (d *DynamoDBMock) PutItemRequest(_ *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) Query(_ *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) QueryWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) QueryRequest(_ *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) QueryPages(_ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) QueryPagesWithContext(_ aws.Context, _ *dynamodb.QueryInput, _ func(*dynamodb.QueryOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableFromBackup(_ *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableFromBackupWithContext(_ aws.Context, _ *dynamodb.RestoreTableFromBackupInput, _ ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableFromBackupRequest(_ *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableToPointInTime(_ *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableToPointInTimeWithContext(_ aws.Context, _ *dynamodb.RestoreTableToPointInTimeInput, _ ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) RestoreTableToPointInTimeRequest(_ *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) Scan(_ *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ScanWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ ...request.Option) (*dynamodb.ScanOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ScanRequest(_ *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ScanPages(_ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) ScanPagesWithContext(_ aws.Context, _ *dynamodb.ScanInput, _ func(*dynamodb.ScanOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TagResource(_ *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TagResourceWithContext(_ aws.Context, _ *dynamodb.TagResourceInput, _ ...request.Option) (*dynamodb.TagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TagResourceRequest(_ *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactGetItems(_ *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactGetItemsWithContext(_ aws.Context, _ *dynamodb.TransactGetItemsInput, _ ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactGetItemsRequest(_ *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactWriteItems(_ *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactWriteItemsWithContext(_ aws.Context, _ *dynamodb.TransactWriteItemsInput, _ ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) TransactWriteItemsRequest(_ *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UntagResource(_ *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UntagResourceWithContext(_ aws.Context, _ *dynamodb.UntagResourceInput, _ ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UntagResourceRequest(_ *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContinuousBackups(_ *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContinuousBackupsWithContext(_ aws.Context, _ *dynamodb.UpdateContinuousBackupsInput, _ ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContinuousBackupsRequest(_ *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContributorInsights(_ *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContributorInsightsWithContext(_ aws.Context, _ *dynamodb.UpdateContributorInsightsInput, _ ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateContributorInsightsRequest(_ *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTable(_ *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTableWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTableRequest(_ *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTableSettings(_ *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTableSettingsWithContext(_ aws.Context, _ *dynamodb.UpdateGlobalTableSettingsInput, _ ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateGlobalTableSettingsRequest(_ *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateItem(_ *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateItemWithContext(_ aws.Context, _ *dynamodb.UpdateItemInput, _ ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateItemRequest(_ *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTable(_ *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTableWithContext(_ aws.Context, _ *dynamodb.UpdateTableInput, _ ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTableRequest(_ *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTableReplicaAutoScaling(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTableReplicaAutoScalingWithContext(_ aws.Context, _ *dynamodb.UpdateTableReplicaAutoScalingInput, _ ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTableReplicaAutoScalingRequest(_ *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTimeToLive(_ *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTimeToLiveWithContext(_ aws.Context, _ *dynamodb.UpdateTimeToLiveInput, _ ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) UpdateTimeToLiveRequest(_ *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) WaitUntilTableExists(_ *dynamodb.DescribeTableInput) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) WaitUntilTableExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) WaitUntilTableNotExists(_ *dynamodb.DescribeTableInput) error {
	panic("not implemented") // TODO: Implement
}

func (d *DynamoDBMock) WaitUntilTableNotExistsWithContext(_ aws.Context, _ *dynamodb.DescribeTableInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}
