package userlist

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Mock struct {
	GetObjectFn func(in *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	PutObjectFn func(_ *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func (s *S3Mock) AbortMultipartUpload(_ *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) AbortMultipartUploadWithContext(_ aws.Context, _ *s3.AbortMultipartUploadInput, _ ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) AbortMultipartUploadRequest(_ *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CompleteMultipartUpload(_ *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CompleteMultipartUploadWithContext(_ aws.Context, _ *s3.CompleteMultipartUploadInput, _ ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CompleteMultipartUploadRequest(_ *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CopyObject(_ *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CopyObjectWithContext(_ aws.Context, _ *s3.CopyObjectInput, _ ...request.Option) (*s3.CopyObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CopyObjectRequest(_ *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateBucket(_ *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateBucketWithContext(_ aws.Context, _ *s3.CreateBucketInput, _ ...request.Option) (*s3.CreateBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateBucketRequest(_ *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateMultipartUpload(_ *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateMultipartUploadWithContext(_ aws.Context, _ *s3.CreateMultipartUploadInput, _ ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) CreateMultipartUploadRequest(_ *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucket(_ *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketWithContext(_ aws.Context, _ *s3.DeleteBucketInput, _ ...request.Option) (*s3.DeleteBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketRequest(_ *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketAnalyticsConfiguration(_ *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketAnalyticsConfigurationWithContext(_ aws.Context, _ *s3.DeleteBucketAnalyticsConfigurationInput, _ ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketAnalyticsConfigurationRequest(_ *s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketCors(_ *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketCorsWithContext(_ aws.Context, _ *s3.DeleteBucketCorsInput, _ ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketCorsRequest(_ *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketEncryption(_ *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketEncryptionWithContext(_ aws.Context, _ *s3.DeleteBucketEncryptionInput, _ ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketEncryptionRequest(_ *s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketIntelligentTieringConfiguration(_ *s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketIntelligentTieringConfigurationWithContext(_ aws.Context, _ *s3.DeleteBucketIntelligentTieringConfigurationInput, _ ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketIntelligentTieringConfigurationRequest(_ *s3.DeleteBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.DeleteBucketIntelligentTieringConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketInventoryConfiguration(_ *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketInventoryConfigurationWithContext(_ aws.Context, _ *s3.DeleteBucketInventoryConfigurationInput, _ ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketInventoryConfigurationRequest(_ *s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketLifecycle(_ *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketLifecycleWithContext(_ aws.Context, _ *s3.DeleteBucketLifecycleInput, _ ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketLifecycleRequest(_ *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketMetricsConfiguration(_ *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketMetricsConfigurationWithContext(_ aws.Context, _ *s3.DeleteBucketMetricsConfigurationInput, _ ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketMetricsConfigurationRequest(_ *s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketOwnershipControls(_ *s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketOwnershipControlsWithContext(_ aws.Context, _ *s3.DeleteBucketOwnershipControlsInput, _ ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketOwnershipControlsRequest(_ *s3.DeleteBucketOwnershipControlsInput) (*request.Request, *s3.DeleteBucketOwnershipControlsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketPolicy(_ *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketPolicyWithContext(_ aws.Context, _ *s3.DeleteBucketPolicyInput, _ ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketPolicyRequest(_ *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketReplication(_ *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketReplicationWithContext(_ aws.Context, _ *s3.DeleteBucketReplicationInput, _ ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketReplicationRequest(_ *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketTagging(_ *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketTaggingWithContext(_ aws.Context, _ *s3.DeleteBucketTaggingInput, _ ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketTaggingRequest(_ *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketWebsite(_ *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketWebsiteWithContext(_ aws.Context, _ *s3.DeleteBucketWebsiteInput, _ ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteBucketWebsiteRequest(_ *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObject(_ *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectWithContext(_ aws.Context, _ *s3.DeleteObjectInput, _ ...request.Option) (*s3.DeleteObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectRequest(_ *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectTagging(_ *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectTaggingWithContext(_ aws.Context, _ *s3.DeleteObjectTaggingInput, _ ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectTaggingRequest(_ *s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjects(_ *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectsWithContext(_ aws.Context, _ *s3.DeleteObjectsInput, _ ...request.Option) (*s3.DeleteObjectsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeleteObjectsRequest(_ *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeletePublicAccessBlock(_ *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeletePublicAccessBlockWithContext(_ aws.Context, _ *s3.DeletePublicAccessBlockInput, _ ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) DeletePublicAccessBlockRequest(_ *s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAccelerateConfiguration(_ *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAccelerateConfigurationWithContext(_ aws.Context, _ *s3.GetBucketAccelerateConfigurationInput, _ ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAccelerateConfigurationRequest(_ *s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAcl(_ *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAclWithContext(_ aws.Context, _ *s3.GetBucketAclInput, _ ...request.Option) (*s3.GetBucketAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAclRequest(_ *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAnalyticsConfiguration(_ *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAnalyticsConfigurationWithContext(_ aws.Context, _ *s3.GetBucketAnalyticsConfigurationInput, _ ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketAnalyticsConfigurationRequest(_ *s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketCors(_ *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketCorsWithContext(_ aws.Context, _ *s3.GetBucketCorsInput, _ ...request.Option) (*s3.GetBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketCorsRequest(_ *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketEncryption(_ *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketEncryptionWithContext(_ aws.Context, _ *s3.GetBucketEncryptionInput, _ ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketEncryptionRequest(_ *s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketIntelligentTieringConfiguration(_ *s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketIntelligentTieringConfigurationWithContext(_ aws.Context, _ *s3.GetBucketIntelligentTieringConfigurationInput, _ ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketIntelligentTieringConfigurationRequest(_ *s3.GetBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.GetBucketIntelligentTieringConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketInventoryConfiguration(_ *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketInventoryConfigurationWithContext(_ aws.Context, _ *s3.GetBucketInventoryConfigurationInput, _ ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketInventoryConfigurationRequest(_ *s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycle(_ *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycleWithContext(_ aws.Context, _ *s3.GetBucketLifecycleInput, _ ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycleRequest(_ *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycleConfiguration(_ *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycleConfigurationWithContext(_ aws.Context, _ *s3.GetBucketLifecycleConfigurationInput, _ ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLifecycleConfigurationRequest(_ *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLocation(_ *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLocationWithContext(_ aws.Context, _ *s3.GetBucketLocationInput, _ ...request.Option) (*s3.GetBucketLocationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLocationRequest(_ *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLogging(_ *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLoggingWithContext(_ aws.Context, _ *s3.GetBucketLoggingInput, _ ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketLoggingRequest(_ *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketMetricsConfiguration(_ *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketMetricsConfigurationWithContext(_ aws.Context, _ *s3.GetBucketMetricsConfigurationInput, _ ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketMetricsConfigurationRequest(_ *s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotification(_ *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotificationWithContext(_ aws.Context, _ *s3.GetBucketNotificationConfigurationRequest, _ ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotificationRequest(_ *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotificationConfiguration(_ *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotificationConfigurationWithContext(_ aws.Context, _ *s3.GetBucketNotificationConfigurationRequest, _ ...request.Option) (*s3.NotificationConfiguration, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketNotificationConfigurationRequest(_ *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketOwnershipControls(_ *s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketOwnershipControlsWithContext(_ aws.Context, _ *s3.GetBucketOwnershipControlsInput, _ ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketOwnershipControlsRequest(_ *s3.GetBucketOwnershipControlsInput) (*request.Request, *s3.GetBucketOwnershipControlsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicy(_ *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicyWithContext(_ aws.Context, _ *s3.GetBucketPolicyInput, _ ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicyRequest(_ *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicyStatus(_ *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicyStatusWithContext(_ aws.Context, _ *s3.GetBucketPolicyStatusInput, _ ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketPolicyStatusRequest(_ *s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketReplication(_ *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketReplicationWithContext(_ aws.Context, _ *s3.GetBucketReplicationInput, _ ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketReplicationRequest(_ *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketRequestPayment(_ *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketRequestPaymentWithContext(_ aws.Context, _ *s3.GetBucketRequestPaymentInput, _ ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketRequestPaymentRequest(_ *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketTagging(_ *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketTaggingWithContext(_ aws.Context, _ *s3.GetBucketTaggingInput, _ ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketTaggingRequest(_ *s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketVersioning(_ *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketVersioningWithContext(_ aws.Context, _ *s3.GetBucketVersioningInput, _ ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketVersioningRequest(_ *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketWebsite(_ *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketWebsiteWithContext(_ aws.Context, _ *s3.GetBucketWebsiteInput, _ ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetBucketWebsiteRequest(_ *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObject(_ *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectWithContext(_ aws.Context, in *s3.GetObjectInput, _ ...request.Option) (*s3.GetObjectOutput, error) {
	if f := s.GetObjectFn; f != nil {
		return f(in)
	}
	panic("not assigned")
}

func (s *S3Mock) GetObjectRequest(_ *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAcl(_ *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAclWithContext(_ aws.Context, _ *s3.GetObjectAclInput, _ ...request.Option) (*s3.GetObjectAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAclRequest(_ *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAttributes(_ *s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAttributesWithContext(_ aws.Context, _ *s3.GetObjectAttributesInput, _ ...request.Option) (*s3.GetObjectAttributesOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectAttributesRequest(_ *s3.GetObjectAttributesInput) (*request.Request, *s3.GetObjectAttributesOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLegalHold(_ *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLegalHoldWithContext(_ aws.Context, _ *s3.GetObjectLegalHoldInput, _ ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLegalHoldRequest(_ *s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLockConfiguration(_ *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLockConfigurationWithContext(_ aws.Context, _ *s3.GetObjectLockConfigurationInput, _ ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectLockConfigurationRequest(_ *s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectRetention(_ *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectRetentionWithContext(_ aws.Context, _ *s3.GetObjectRetentionInput, _ ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectRetentionRequest(_ *s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTagging(_ *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTaggingWithContext(_ aws.Context, _ *s3.GetObjectTaggingInput, _ ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTaggingRequest(_ *s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTorrent(_ *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTorrentWithContext(_ aws.Context, _ *s3.GetObjectTorrentInput, _ ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetObjectTorrentRequest(_ *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetPublicAccessBlock(_ *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetPublicAccessBlockWithContext(_ aws.Context, _ *s3.GetPublicAccessBlockInput, _ ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) GetPublicAccessBlockRequest(_ *s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadBucket(_ *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadBucketWithContext(_ aws.Context, _ *s3.HeadBucketInput, _ ...request.Option) (*s3.HeadBucketOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadBucketRequest(_ *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadObject(_ *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadObjectWithContext(_ aws.Context, _ *s3.HeadObjectInput, _ ...request.Option) (*s3.HeadObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) HeadObjectRequest(_ *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketAnalyticsConfigurations(_ *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketAnalyticsConfigurationsWithContext(_ aws.Context, _ *s3.ListBucketAnalyticsConfigurationsInput, _ ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketAnalyticsConfigurationsRequest(_ *s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketIntelligentTieringConfigurations(_ *s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketIntelligentTieringConfigurationsWithContext(_ aws.Context, _ *s3.ListBucketIntelligentTieringConfigurationsInput, _ ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketIntelligentTieringConfigurationsRequest(_ *s3.ListBucketIntelligentTieringConfigurationsInput) (*request.Request, *s3.ListBucketIntelligentTieringConfigurationsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketInventoryConfigurations(_ *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketInventoryConfigurationsWithContext(_ aws.Context, _ *s3.ListBucketInventoryConfigurationsInput, _ ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketInventoryConfigurationsRequest(_ *s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketMetricsConfigurations(_ *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketMetricsConfigurationsWithContext(_ aws.Context, _ *s3.ListBucketMetricsConfigurationsInput, _ ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketMetricsConfigurationsRequest(_ *s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBuckets(_ *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketsWithContext(_ aws.Context, _ *s3.ListBucketsInput, _ ...request.Option) (*s3.ListBucketsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListBucketsRequest(_ *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListMultipartUploads(_ *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListMultipartUploadsWithContext(_ aws.Context, _ *s3.ListMultipartUploadsInput, _ ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListMultipartUploadsRequest(_ *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListMultipartUploadsPages(_ *s3.ListMultipartUploadsInput, _ func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListMultipartUploadsPagesWithContext(_ aws.Context, _ *s3.ListMultipartUploadsInput, _ func(*s3.ListMultipartUploadsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectVersions(_ *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectVersionsWithContext(_ aws.Context, _ *s3.ListObjectVersionsInput, _ ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectVersionsRequest(_ *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectVersionsPages(_ *s3.ListObjectVersionsInput, _ func(*s3.ListObjectVersionsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectVersionsPagesWithContext(_ aws.Context, _ *s3.ListObjectVersionsInput, _ func(*s3.ListObjectVersionsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjects(_ *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsWithContext(_ aws.Context, _ *s3.ListObjectsInput, _ ...request.Option) (*s3.ListObjectsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsRequest(_ *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsPages(_ *s3.ListObjectsInput, _ func(*s3.ListObjectsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsPagesWithContext(_ aws.Context, _ *s3.ListObjectsInput, _ func(*s3.ListObjectsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsV2(_ *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsV2WithContext(_ aws.Context, _ *s3.ListObjectsV2Input, _ ...request.Option) (*s3.ListObjectsV2Output, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsV2Request(_ *s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsV2Pages(_ *s3.ListObjectsV2Input, _ func(*s3.ListObjectsV2Output, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListObjectsV2PagesWithContext(_ aws.Context, _ *s3.ListObjectsV2Input, _ func(*s3.ListObjectsV2Output, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListParts(_ *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListPartsWithContext(_ aws.Context, _ *s3.ListPartsInput, _ ...request.Option) (*s3.ListPartsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListPartsRequest(_ *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListPartsPages(_ *s3.ListPartsInput, _ func(*s3.ListPartsOutput, bool) bool) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) ListPartsPagesWithContext(_ aws.Context, _ *s3.ListPartsInput, _ func(*s3.ListPartsOutput, bool) bool, _ ...request.Option) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAccelerateConfiguration(_ *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAccelerateConfigurationWithContext(_ aws.Context, _ *s3.PutBucketAccelerateConfigurationInput, _ ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAccelerateConfigurationRequest(_ *s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAcl(_ *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAclWithContext(_ aws.Context, _ *s3.PutBucketAclInput, _ ...request.Option) (*s3.PutBucketAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAclRequest(_ *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAnalyticsConfiguration(_ *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAnalyticsConfigurationWithContext(_ aws.Context, _ *s3.PutBucketAnalyticsConfigurationInput, _ ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketAnalyticsConfigurationRequest(_ *s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketCors(_ *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketCorsWithContext(_ aws.Context, _ *s3.PutBucketCorsInput, _ ...request.Option) (*s3.PutBucketCorsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketCorsRequest(_ *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketEncryption(_ *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketEncryptionWithContext(_ aws.Context, _ *s3.PutBucketEncryptionInput, _ ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketEncryptionRequest(_ *s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketIntelligentTieringConfiguration(_ *s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketIntelligentTieringConfigurationWithContext(_ aws.Context, _ *s3.PutBucketIntelligentTieringConfigurationInput, _ ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketIntelligentTieringConfigurationRequest(_ *s3.PutBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.PutBucketIntelligentTieringConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketInventoryConfiguration(_ *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketInventoryConfigurationWithContext(_ aws.Context, _ *s3.PutBucketInventoryConfigurationInput, _ ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketInventoryConfigurationRequest(_ *s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycle(_ *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycleWithContext(_ aws.Context, _ *s3.PutBucketLifecycleInput, _ ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycleRequest(_ *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycleConfiguration(_ *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycleConfigurationWithContext(_ aws.Context, _ *s3.PutBucketLifecycleConfigurationInput, _ ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLifecycleConfigurationRequest(_ *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLogging(_ *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLoggingWithContext(_ aws.Context, _ *s3.PutBucketLoggingInput, _ ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketLoggingRequest(_ *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketMetricsConfiguration(_ *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketMetricsConfigurationWithContext(_ aws.Context, _ *s3.PutBucketMetricsConfigurationInput, _ ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketMetricsConfigurationRequest(_ *s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotification(_ *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotificationWithContext(_ aws.Context, _ *s3.PutBucketNotificationInput, _ ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotificationRequest(_ *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotificationConfiguration(_ *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotificationConfigurationWithContext(_ aws.Context, _ *s3.PutBucketNotificationConfigurationInput, _ ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketNotificationConfigurationRequest(_ *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketOwnershipControls(_ *s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketOwnershipControlsWithContext(_ aws.Context, _ *s3.PutBucketOwnershipControlsInput, _ ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketOwnershipControlsRequest(_ *s3.PutBucketOwnershipControlsInput) (*request.Request, *s3.PutBucketOwnershipControlsOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketPolicy(_ *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketPolicyWithContext(_ aws.Context, _ *s3.PutBucketPolicyInput, _ ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketPolicyRequest(_ *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketReplication(_ *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketReplicationWithContext(_ aws.Context, _ *s3.PutBucketReplicationInput, _ ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketReplicationRequest(_ *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketRequestPayment(_ *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketRequestPaymentWithContext(_ aws.Context, _ *s3.PutBucketRequestPaymentInput, _ ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketRequestPaymentRequest(_ *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketTagging(_ *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketTaggingWithContext(_ aws.Context, _ *s3.PutBucketTaggingInput, _ ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketTaggingRequest(_ *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketVersioning(_ *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketVersioningWithContext(_ aws.Context, _ *s3.PutBucketVersioningInput, _ ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketVersioningRequest(_ *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketWebsite(_ *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketWebsiteWithContext(_ aws.Context, _ *s3.PutBucketWebsiteInput, _ ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutBucketWebsiteRequest(_ *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObject(_ *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectWithContext(_ aws.Context, in *s3.PutObjectInput, _ ...request.Option) (*s3.PutObjectOutput, error) {
	if f := s.PutObjectFn; f != nil {
		return f(in)
	}
	panic("not assigned")
}

func (s *S3Mock) PutObjectRequest(_ *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectAcl(_ *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectAclWithContext(_ aws.Context, _ *s3.PutObjectAclInput, _ ...request.Option) (*s3.PutObjectAclOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectAclRequest(_ *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLegalHold(_ *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLegalHoldWithContext(_ aws.Context, _ *s3.PutObjectLegalHoldInput, _ ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLegalHoldRequest(_ *s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLockConfiguration(_ *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLockConfigurationWithContext(_ aws.Context, _ *s3.PutObjectLockConfigurationInput, _ ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectLockConfigurationRequest(_ *s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectRetention(_ *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectRetentionWithContext(_ aws.Context, _ *s3.PutObjectRetentionInput, _ ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectRetentionRequest(_ *s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectTagging(_ *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectTaggingWithContext(_ aws.Context, _ *s3.PutObjectTaggingInput, _ ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutObjectTaggingRequest(_ *s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutPublicAccessBlock(_ *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutPublicAccessBlockWithContext(_ aws.Context, _ *s3.PutPublicAccessBlockInput, _ ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) PutPublicAccessBlockRequest(_ *s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) RestoreObject(_ *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) RestoreObjectWithContext(_ aws.Context, _ *s3.RestoreObjectInput, _ ...request.Option) (*s3.RestoreObjectOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) RestoreObjectRequest(_ *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) SelectObjectContent(_ *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) SelectObjectContentWithContext(_ aws.Context, _ *s3.SelectObjectContentInput, _ ...request.Option) (*s3.SelectObjectContentOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) SelectObjectContentRequest(_ *s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPart(_ *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPartWithContext(_ aws.Context, _ *s3.UploadPartInput, _ ...request.Option) (*s3.UploadPartOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPartRequest(_ *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPartCopy(_ *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPartCopyWithContext(_ aws.Context, _ *s3.UploadPartCopyInput, _ ...request.Option) (*s3.UploadPartCopyOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) UploadPartCopyRequest(_ *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WriteGetObjectResponse(_ *s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WriteGetObjectResponseWithContext(_ aws.Context, _ *s3.WriteGetObjectResponseInput, _ ...request.Option) (*s3.WriteGetObjectResponseOutput, error) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WriteGetObjectResponseRequest(_ *s3.WriteGetObjectResponseInput) (*request.Request, *s3.WriteGetObjectResponseOutput) {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilBucketExists(_ *s3.HeadBucketInput) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilBucketExistsWithContext(_ aws.Context, _ *s3.HeadBucketInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilBucketNotExists(_ *s3.HeadBucketInput) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilBucketNotExistsWithContext(_ aws.Context, _ *s3.HeadBucketInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilObjectExists(_ *s3.HeadObjectInput) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilObjectExistsWithContext(_ aws.Context, _ *s3.HeadObjectInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilObjectNotExists(_ *s3.HeadObjectInput) error {
	panic("not implemented") // TODO: Implement
}

func (s *S3Mock) WaitUntilObjectNotExistsWithContext(_ aws.Context, _ *s3.HeadObjectInput, _ ...request.WaiterOption) error {
	panic("not implemented") // TODO: Implement
}

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
