// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdcpg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tdcpg instances.
//
// > **NOTE:** This data source is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tdcpg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdcpg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tdcpg.GetInstances(ctx, &tdcpg.GetInstancesArgs{
//				ClusterId:    "",
//				InstanceId:   pulumi.StringRef(""),
//				InstanceName: pulumi.StringRef(""),
//				InstanceType: pulumi.StringRef(""),
//				Status:       pulumi.StringRef(""),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("tencentcloud:Tdcpg/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// instance id.
	ClusterId string `pulumi:"clusterId"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// instance name.
	InstanceName *string `pulumi:"instanceName"`
	// instance type.
	InstanceType *string `pulumi:"instanceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// instance status.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// cluster id.
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// instance name.
	InstanceName *string `pulumi:"instanceName"`
	// instance type.
	InstanceType *string `pulumi:"instanceType"`
	// instance list.
	Lists            []GetInstancesList `pulumi:"lists"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	// status.
	Status *string `pulumi:"status"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// instance id.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// instance id.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// instance name.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// instance type.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// instance status.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// cluster id.
func (o GetInstancesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// instance id.
func (o GetInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// instance name.
func (o GetInstancesResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// instance type.
func (o GetInstancesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// instance list.
func (o GetInstancesResultOutput) Lists() GetInstancesListArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesList { return v.Lists }).(GetInstancesListArrayOutput)
}

func (o GetInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// status.
func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
