// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package drsiface provides an interface to enable mocking the Elastic Disaster Recovery Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package drsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/drs"
)

// DrsAPI provides an interface to enable mocking the
// drs.Drs service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Elastic Disaster Recovery Service.
//    func myFunc(svc drsiface.DrsAPI) bool {
//        // Make svc.CreateReplicationConfigurationTemplate request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := drs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDrsClient struct {
//        drsiface.DrsAPI
//    }
//    func (m *mockDrsClient) CreateReplicationConfigurationTemplate(input *drs.CreateReplicationConfigurationTemplateInput) (*drs.CreateReplicationConfigurationTemplateOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDrsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DrsAPI interface {
	CreateReplicationConfigurationTemplate(*drs.CreateReplicationConfigurationTemplateInput) (*drs.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateWithContext(aws.Context, *drs.CreateReplicationConfigurationTemplateInput, ...request.Option) (*drs.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateRequest(*drs.CreateReplicationConfigurationTemplateInput) (*request.Request, *drs.CreateReplicationConfigurationTemplateOutput)

	DeleteJob(*drs.DeleteJobInput) (*drs.DeleteJobOutput, error)
	DeleteJobWithContext(aws.Context, *drs.DeleteJobInput, ...request.Option) (*drs.DeleteJobOutput, error)
	DeleteJobRequest(*drs.DeleteJobInput) (*request.Request, *drs.DeleteJobOutput)

	DeleteRecoveryInstance(*drs.DeleteRecoveryInstanceInput) (*drs.DeleteRecoveryInstanceOutput, error)
	DeleteRecoveryInstanceWithContext(aws.Context, *drs.DeleteRecoveryInstanceInput, ...request.Option) (*drs.DeleteRecoveryInstanceOutput, error)
	DeleteRecoveryInstanceRequest(*drs.DeleteRecoveryInstanceInput) (*request.Request, *drs.DeleteRecoveryInstanceOutput)

	DeleteReplicationConfigurationTemplate(*drs.DeleteReplicationConfigurationTemplateInput) (*drs.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateWithContext(aws.Context, *drs.DeleteReplicationConfigurationTemplateInput, ...request.Option) (*drs.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateRequest(*drs.DeleteReplicationConfigurationTemplateInput) (*request.Request, *drs.DeleteReplicationConfigurationTemplateOutput)

	DeleteSourceServer(*drs.DeleteSourceServerInput) (*drs.DeleteSourceServerOutput, error)
	DeleteSourceServerWithContext(aws.Context, *drs.DeleteSourceServerInput, ...request.Option) (*drs.DeleteSourceServerOutput, error)
	DeleteSourceServerRequest(*drs.DeleteSourceServerInput) (*request.Request, *drs.DeleteSourceServerOutput)

	DescribeJobLogItems(*drs.DescribeJobLogItemsInput) (*drs.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsWithContext(aws.Context, *drs.DescribeJobLogItemsInput, ...request.Option) (*drs.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsRequest(*drs.DescribeJobLogItemsInput) (*request.Request, *drs.DescribeJobLogItemsOutput)

	DescribeJobLogItemsPages(*drs.DescribeJobLogItemsInput, func(*drs.DescribeJobLogItemsOutput, bool) bool) error
	DescribeJobLogItemsPagesWithContext(aws.Context, *drs.DescribeJobLogItemsInput, func(*drs.DescribeJobLogItemsOutput, bool) bool, ...request.Option) error

	DescribeJobs(*drs.DescribeJobsInput) (*drs.DescribeJobsOutput, error)
	DescribeJobsWithContext(aws.Context, *drs.DescribeJobsInput, ...request.Option) (*drs.DescribeJobsOutput, error)
	DescribeJobsRequest(*drs.DescribeJobsInput) (*request.Request, *drs.DescribeJobsOutput)

	DescribeJobsPages(*drs.DescribeJobsInput, func(*drs.DescribeJobsOutput, bool) bool) error
	DescribeJobsPagesWithContext(aws.Context, *drs.DescribeJobsInput, func(*drs.DescribeJobsOutput, bool) bool, ...request.Option) error

	DescribeRecoveryInstances(*drs.DescribeRecoveryInstancesInput) (*drs.DescribeRecoveryInstancesOutput, error)
	DescribeRecoveryInstancesWithContext(aws.Context, *drs.DescribeRecoveryInstancesInput, ...request.Option) (*drs.DescribeRecoveryInstancesOutput, error)
	DescribeRecoveryInstancesRequest(*drs.DescribeRecoveryInstancesInput) (*request.Request, *drs.DescribeRecoveryInstancesOutput)

	DescribeRecoveryInstancesPages(*drs.DescribeRecoveryInstancesInput, func(*drs.DescribeRecoveryInstancesOutput, bool) bool) error
	DescribeRecoveryInstancesPagesWithContext(aws.Context, *drs.DescribeRecoveryInstancesInput, func(*drs.DescribeRecoveryInstancesOutput, bool) bool, ...request.Option) error

	DescribeRecoverySnapshots(*drs.DescribeRecoverySnapshotsInput) (*drs.DescribeRecoverySnapshotsOutput, error)
	DescribeRecoverySnapshotsWithContext(aws.Context, *drs.DescribeRecoverySnapshotsInput, ...request.Option) (*drs.DescribeRecoverySnapshotsOutput, error)
	DescribeRecoverySnapshotsRequest(*drs.DescribeRecoverySnapshotsInput) (*request.Request, *drs.DescribeRecoverySnapshotsOutput)

	DescribeRecoverySnapshotsPages(*drs.DescribeRecoverySnapshotsInput, func(*drs.DescribeRecoverySnapshotsOutput, bool) bool) error
	DescribeRecoverySnapshotsPagesWithContext(aws.Context, *drs.DescribeRecoverySnapshotsInput, func(*drs.DescribeRecoverySnapshotsOutput, bool) bool, ...request.Option) error

	DescribeReplicationConfigurationTemplates(*drs.DescribeReplicationConfigurationTemplatesInput) (*drs.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesWithContext(aws.Context, *drs.DescribeReplicationConfigurationTemplatesInput, ...request.Option) (*drs.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesRequest(*drs.DescribeReplicationConfigurationTemplatesInput) (*request.Request, *drs.DescribeReplicationConfigurationTemplatesOutput)

	DescribeReplicationConfigurationTemplatesPages(*drs.DescribeReplicationConfigurationTemplatesInput, func(*drs.DescribeReplicationConfigurationTemplatesOutput, bool) bool) error
	DescribeReplicationConfigurationTemplatesPagesWithContext(aws.Context, *drs.DescribeReplicationConfigurationTemplatesInput, func(*drs.DescribeReplicationConfigurationTemplatesOutput, bool) bool, ...request.Option) error

	DescribeSourceServers(*drs.DescribeSourceServersInput) (*drs.DescribeSourceServersOutput, error)
	DescribeSourceServersWithContext(aws.Context, *drs.DescribeSourceServersInput, ...request.Option) (*drs.DescribeSourceServersOutput, error)
	DescribeSourceServersRequest(*drs.DescribeSourceServersInput) (*request.Request, *drs.DescribeSourceServersOutput)

	DescribeSourceServersPages(*drs.DescribeSourceServersInput, func(*drs.DescribeSourceServersOutput, bool) bool) error
	DescribeSourceServersPagesWithContext(aws.Context, *drs.DescribeSourceServersInput, func(*drs.DescribeSourceServersOutput, bool) bool, ...request.Option) error

	DisconnectRecoveryInstance(*drs.DisconnectRecoveryInstanceInput) (*drs.DisconnectRecoveryInstanceOutput, error)
	DisconnectRecoveryInstanceWithContext(aws.Context, *drs.DisconnectRecoveryInstanceInput, ...request.Option) (*drs.DisconnectRecoveryInstanceOutput, error)
	DisconnectRecoveryInstanceRequest(*drs.DisconnectRecoveryInstanceInput) (*request.Request, *drs.DisconnectRecoveryInstanceOutput)

	DisconnectSourceServer(*drs.DisconnectSourceServerInput) (*drs.DisconnectSourceServerOutput, error)
	DisconnectSourceServerWithContext(aws.Context, *drs.DisconnectSourceServerInput, ...request.Option) (*drs.DisconnectSourceServerOutput, error)
	DisconnectSourceServerRequest(*drs.DisconnectSourceServerInput) (*request.Request, *drs.DisconnectSourceServerOutput)

	GetFailbackReplicationConfiguration(*drs.GetFailbackReplicationConfigurationInput) (*drs.GetFailbackReplicationConfigurationOutput, error)
	GetFailbackReplicationConfigurationWithContext(aws.Context, *drs.GetFailbackReplicationConfigurationInput, ...request.Option) (*drs.GetFailbackReplicationConfigurationOutput, error)
	GetFailbackReplicationConfigurationRequest(*drs.GetFailbackReplicationConfigurationInput) (*request.Request, *drs.GetFailbackReplicationConfigurationOutput)

	GetLaunchConfiguration(*drs.GetLaunchConfigurationInput) (*drs.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationWithContext(aws.Context, *drs.GetLaunchConfigurationInput, ...request.Option) (*drs.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationRequest(*drs.GetLaunchConfigurationInput) (*request.Request, *drs.GetLaunchConfigurationOutput)

	GetReplicationConfiguration(*drs.GetReplicationConfigurationInput) (*drs.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationWithContext(aws.Context, *drs.GetReplicationConfigurationInput, ...request.Option) (*drs.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationRequest(*drs.GetReplicationConfigurationInput) (*request.Request, *drs.GetReplicationConfigurationOutput)

	InitializeService(*drs.InitializeServiceInput) (*drs.InitializeServiceOutput, error)
	InitializeServiceWithContext(aws.Context, *drs.InitializeServiceInput, ...request.Option) (*drs.InitializeServiceOutput, error)
	InitializeServiceRequest(*drs.InitializeServiceInput) (*request.Request, *drs.InitializeServiceOutput)

	ListTagsForResource(*drs.ListTagsForResourceInput) (*drs.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *drs.ListTagsForResourceInput, ...request.Option) (*drs.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*drs.ListTagsForResourceInput) (*request.Request, *drs.ListTagsForResourceOutput)

	RetryDataReplication(*drs.RetryDataReplicationInput) (*drs.RetryDataReplicationOutput, error)
	RetryDataReplicationWithContext(aws.Context, *drs.RetryDataReplicationInput, ...request.Option) (*drs.RetryDataReplicationOutput, error)
	RetryDataReplicationRequest(*drs.RetryDataReplicationInput) (*request.Request, *drs.RetryDataReplicationOutput)

	StartFailbackLaunch(*drs.StartFailbackLaunchInput) (*drs.StartFailbackLaunchOutput, error)
	StartFailbackLaunchWithContext(aws.Context, *drs.StartFailbackLaunchInput, ...request.Option) (*drs.StartFailbackLaunchOutput, error)
	StartFailbackLaunchRequest(*drs.StartFailbackLaunchInput) (*request.Request, *drs.StartFailbackLaunchOutput)

	StartRecovery(*drs.StartRecoveryInput) (*drs.StartRecoveryOutput, error)
	StartRecoveryWithContext(aws.Context, *drs.StartRecoveryInput, ...request.Option) (*drs.StartRecoveryOutput, error)
	StartRecoveryRequest(*drs.StartRecoveryInput) (*request.Request, *drs.StartRecoveryOutput)

	StopFailback(*drs.StopFailbackInput) (*drs.StopFailbackOutput, error)
	StopFailbackWithContext(aws.Context, *drs.StopFailbackInput, ...request.Option) (*drs.StopFailbackOutput, error)
	StopFailbackRequest(*drs.StopFailbackInput) (*request.Request, *drs.StopFailbackOutput)

	TagResource(*drs.TagResourceInput) (*drs.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *drs.TagResourceInput, ...request.Option) (*drs.TagResourceOutput, error)
	TagResourceRequest(*drs.TagResourceInput) (*request.Request, *drs.TagResourceOutput)

	TerminateRecoveryInstances(*drs.TerminateRecoveryInstancesInput) (*drs.TerminateRecoveryInstancesOutput, error)
	TerminateRecoveryInstancesWithContext(aws.Context, *drs.TerminateRecoveryInstancesInput, ...request.Option) (*drs.TerminateRecoveryInstancesOutput, error)
	TerminateRecoveryInstancesRequest(*drs.TerminateRecoveryInstancesInput) (*request.Request, *drs.TerminateRecoveryInstancesOutput)

	UntagResource(*drs.UntagResourceInput) (*drs.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *drs.UntagResourceInput, ...request.Option) (*drs.UntagResourceOutput, error)
	UntagResourceRequest(*drs.UntagResourceInput) (*request.Request, *drs.UntagResourceOutput)

	UpdateFailbackReplicationConfiguration(*drs.UpdateFailbackReplicationConfigurationInput) (*drs.UpdateFailbackReplicationConfigurationOutput, error)
	UpdateFailbackReplicationConfigurationWithContext(aws.Context, *drs.UpdateFailbackReplicationConfigurationInput, ...request.Option) (*drs.UpdateFailbackReplicationConfigurationOutput, error)
	UpdateFailbackReplicationConfigurationRequest(*drs.UpdateFailbackReplicationConfigurationInput) (*request.Request, *drs.UpdateFailbackReplicationConfigurationOutput)

	UpdateLaunchConfiguration(*drs.UpdateLaunchConfigurationInput) (*drs.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationWithContext(aws.Context, *drs.UpdateLaunchConfigurationInput, ...request.Option) (*drs.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationRequest(*drs.UpdateLaunchConfigurationInput) (*request.Request, *drs.UpdateLaunchConfigurationOutput)

	UpdateReplicationConfiguration(*drs.UpdateReplicationConfigurationInput) (*drs.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationWithContext(aws.Context, *drs.UpdateReplicationConfigurationInput, ...request.Option) (*drs.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationRequest(*drs.UpdateReplicationConfigurationInput) (*request.Request, *drs.UpdateReplicationConfigurationOutput)

	UpdateReplicationConfigurationTemplate(*drs.UpdateReplicationConfigurationTemplateInput) (*drs.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateWithContext(aws.Context, *drs.UpdateReplicationConfigurationTemplateInput, ...request.Option) (*drs.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateRequest(*drs.UpdateReplicationConfigurationTemplateInput) (*request.Request, *drs.UpdateReplicationConfigurationTemplateOutput)
}

var _ DrsAPI = (*drs.Drs)(nil)
