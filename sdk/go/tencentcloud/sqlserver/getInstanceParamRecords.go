// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of sqlserver instanceParamRecords
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sqlserver.GetInstanceParamRecords(ctx, &sqlserver.GetInstanceParamRecordsArgs{
//				InstanceId: "mssql-qelbzgwf",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceParamRecords(ctx *pulumi.Context, args *GetInstanceParamRecordsArgs, opts ...pulumi.InvokeOption) (*GetInstanceParamRecordsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceParamRecordsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getInstanceParamRecords:getInstanceParamRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceParamRecords.
type GetInstanceParamRecordsArgs struct {
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter InstanceId of the DescribeDBInstances API.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceParamRecords.
type GetInstanceParamRecordsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Parameter modification records.
	Items            []GetInstanceParamRecordsItem `pulumi:"items"`
	ResultOutputFile *string                       `pulumi:"resultOutputFile"`
}

func GetInstanceParamRecordsOutput(ctx *pulumi.Context, args GetInstanceParamRecordsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceParamRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceParamRecordsResult, error) {
			args := v.(GetInstanceParamRecordsArgs)
			r, err := GetInstanceParamRecords(ctx, &args, opts...)
			var s GetInstanceParamRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceParamRecordsResultOutput)
}

// A collection of arguments for invoking getInstanceParamRecords.
type GetInstanceParamRecordsOutputArgs struct {
	// Instance ID in the format of mssql-dj5i29c5n. It is the same as the instance ID displayed in the TencentDB console and the response parameter InstanceId of the DescribeDBInstances API.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceParamRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceParamRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceParamRecords.
type GetInstanceParamRecordsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceParamRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceParamRecordsResult)(nil)).Elem()
}

func (o GetInstanceParamRecordsResultOutput) ToGetInstanceParamRecordsResultOutput() GetInstanceParamRecordsResultOutput {
	return o
}

func (o GetInstanceParamRecordsResultOutput) ToGetInstanceParamRecordsResultOutputWithContext(ctx context.Context) GetInstanceParamRecordsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceParamRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceParamRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance ID.
func (o GetInstanceParamRecordsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceParamRecordsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Parameter modification records.
func (o GetInstanceParamRecordsResultOutput) Items() GetInstanceParamRecordsItemArrayOutput {
	return o.ApplyT(func(v GetInstanceParamRecordsResult) []GetInstanceParamRecordsItem { return v.Items }).(GetInstanceParamRecordsItemArrayOutput)
}

func (o GetInstanceParamRecordsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceParamRecordsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceParamRecordsResultOutput{})
}
