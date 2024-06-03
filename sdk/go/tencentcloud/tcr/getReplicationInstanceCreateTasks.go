// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of tcr replicationInstanceCreateTasks
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tcr.GetReplicationInstanceCreateTasks(ctx, &tcr.GetReplicationInstanceCreateTasksArgs{
//				ReplicationRegistryId: local.Dst_registry_id,
//				ReplicationRegionId:   local.Dst_region_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetReplicationInstanceCreateTasks(ctx *pulumi.Context, args *GetReplicationInstanceCreateTasksArgs, opts ...pulumi.InvokeOption) (*GetReplicationInstanceCreateTasksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReplicationInstanceCreateTasksResult
	err := ctx.Invoke("tencentcloud:Tcr/getReplicationInstanceCreateTasks:getReplicationInstanceCreateTasks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplicationInstanceCreateTasks.
type GetReplicationInstanceCreateTasksArgs struct {
	// synchronization instance region Id, see ReplicationRegionId in DescribeReplicationInstances.
	ReplicationRegionId int `pulumi:"replicationRegionId"`
	// synchronization instance Id, see RegistryId in DescribeReplicationInstances.
	ReplicationRegistryId string `pulumi:"replicationRegistryId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getReplicationInstanceCreateTasks.
type GetReplicationInstanceCreateTasksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                    string  `pulumi:"id"`
	ReplicationRegionId   int     `pulumi:"replicationRegionId"`
	ReplicationRegistryId string  `pulumi:"replicationRegistryId"`
	ResultOutputFile      *string `pulumi:"resultOutputFile"`
	// overall task status.
	Status string `pulumi:"status"`
	// task details.
	TaskDetails []GetReplicationInstanceCreateTasksTaskDetail `pulumi:"taskDetails"`
}

func GetReplicationInstanceCreateTasksOutput(ctx *pulumi.Context, args GetReplicationInstanceCreateTasksOutputArgs, opts ...pulumi.InvokeOption) GetReplicationInstanceCreateTasksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReplicationInstanceCreateTasksResult, error) {
			args := v.(GetReplicationInstanceCreateTasksArgs)
			r, err := GetReplicationInstanceCreateTasks(ctx, &args, opts...)
			var s GetReplicationInstanceCreateTasksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReplicationInstanceCreateTasksResultOutput)
}

// A collection of arguments for invoking getReplicationInstanceCreateTasks.
type GetReplicationInstanceCreateTasksOutputArgs struct {
	// synchronization instance region Id, see ReplicationRegionId in DescribeReplicationInstances.
	ReplicationRegionId pulumi.IntInput `pulumi:"replicationRegionId"`
	// synchronization instance Id, see RegistryId in DescribeReplicationInstances.
	ReplicationRegistryId pulumi.StringInput `pulumi:"replicationRegistryId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetReplicationInstanceCreateTasksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReplicationInstanceCreateTasksArgs)(nil)).Elem()
}

// A collection of values returned by getReplicationInstanceCreateTasks.
type GetReplicationInstanceCreateTasksResultOutput struct{ *pulumi.OutputState }

func (GetReplicationInstanceCreateTasksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReplicationInstanceCreateTasksResult)(nil)).Elem()
}

func (o GetReplicationInstanceCreateTasksResultOutput) ToGetReplicationInstanceCreateTasksResultOutput() GetReplicationInstanceCreateTasksResultOutput {
	return o
}

func (o GetReplicationInstanceCreateTasksResultOutput) ToGetReplicationInstanceCreateTasksResultOutputWithContext(ctx context.Context) GetReplicationInstanceCreateTasksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetReplicationInstanceCreateTasksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetReplicationInstanceCreateTasksResultOutput) ReplicationRegionId() pulumi.IntOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) int { return v.ReplicationRegionId }).(pulumi.IntOutput)
}

func (o GetReplicationInstanceCreateTasksResultOutput) ReplicationRegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) string { return v.ReplicationRegistryId }).(pulumi.StringOutput)
}

func (o GetReplicationInstanceCreateTasksResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// overall task status.
func (o GetReplicationInstanceCreateTasksResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) string { return v.Status }).(pulumi.StringOutput)
}

// task details.
func (o GetReplicationInstanceCreateTasksResultOutput) TaskDetails() GetReplicationInstanceCreateTasksTaskDetailArrayOutput {
	return o.ApplyT(func(v GetReplicationInstanceCreateTasksResult) []GetReplicationInstanceCreateTasksTaskDetail {
		return v.TaskDetails
	}).(GetReplicationInstanceCreateTasksTaskDetailArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReplicationInstanceCreateTasksResultOutput{})
}
