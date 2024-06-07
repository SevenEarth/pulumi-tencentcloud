// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of lighthouse instanceVncUrl
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lighthouse.GetInstanceVncUrl(ctx, &lighthouse.GetInstanceVncUrlArgs{
//				InstanceId: "lhins-123456",
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
func GetInstanceVncUrl(ctx *pulumi.Context, args *GetInstanceVncUrlArgs, opts ...pulumi.InvokeOption) (*GetInstanceVncUrlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceVncUrlResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getInstanceVncUrl:getInstanceVncUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceVncUrl.
type GetInstanceVncUrlArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceVncUrl.
type GetInstanceVncUrlResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Instance VNC URL.
	InstanceVncUrl   string  `pulumi:"instanceVncUrl"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetInstanceVncUrlOutput(ctx *pulumi.Context, args GetInstanceVncUrlOutputArgs, opts ...pulumi.InvokeOption) GetInstanceVncUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceVncUrlResult, error) {
			args := v.(GetInstanceVncUrlArgs)
			r, err := GetInstanceVncUrl(ctx, &args, opts...)
			var s GetInstanceVncUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceVncUrlResultOutput)
}

// A collection of arguments for invoking getInstanceVncUrl.
type GetInstanceVncUrlOutputArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceVncUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceVncUrlArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceVncUrl.
type GetInstanceVncUrlResultOutput struct{ *pulumi.OutputState }

func (GetInstanceVncUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceVncUrlResult)(nil)).Elem()
}

func (o GetInstanceVncUrlResultOutput) ToGetInstanceVncUrlResultOutput() GetInstanceVncUrlResultOutput {
	return o
}

func (o GetInstanceVncUrlResultOutput) ToGetInstanceVncUrlResultOutputWithContext(ctx context.Context) GetInstanceVncUrlResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceVncUrlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceVncUrlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceVncUrlResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceVncUrlResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance VNC URL.
func (o GetInstanceVncUrlResultOutput) InstanceVncUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceVncUrlResult) string { return v.InstanceVncUrl }).(pulumi.StringOutput)
}

func (o GetInstanceVncUrlResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceVncUrlResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceVncUrlResultOutput{})
}
