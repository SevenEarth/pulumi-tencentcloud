// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of vpc addressQuota
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eip"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Eip.GetAddressQuota(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAddressQuota(ctx *pulumi.Context, args *GetAddressQuotaArgs, opts ...pulumi.InvokeOption) (*GetAddressQuotaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAddressQuotaResult
	err := ctx.Invoke("tencentcloud:Eip/getAddressQuota:getAddressQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddressQuota.
type GetAddressQuotaArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAddressQuota.
type GetAddressQuotaResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The specified account EIP quota information.
	QuotaSets        []GetAddressQuotaQuotaSet `pulumi:"quotaSets"`
	ResultOutputFile *string                   `pulumi:"resultOutputFile"`
}

func GetAddressQuotaOutput(ctx *pulumi.Context, args GetAddressQuotaOutputArgs, opts ...pulumi.InvokeOption) GetAddressQuotaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddressQuotaResult, error) {
			args := v.(GetAddressQuotaArgs)
			r, err := GetAddressQuota(ctx, &args, opts...)
			var s GetAddressQuotaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddressQuotaResultOutput)
}

// A collection of arguments for invoking getAddressQuota.
type GetAddressQuotaOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAddressQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressQuotaArgs)(nil)).Elem()
}

// A collection of values returned by getAddressQuota.
type GetAddressQuotaResultOutput struct{ *pulumi.OutputState }

func (GetAddressQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressQuotaResult)(nil)).Elem()
}

func (o GetAddressQuotaResultOutput) ToGetAddressQuotaResultOutput() GetAddressQuotaResultOutput {
	return o
}

func (o GetAddressQuotaResultOutput) ToGetAddressQuotaResultOutputWithContext(ctx context.Context) GetAddressQuotaResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddressQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddressQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The specified account EIP quota information.
func (o GetAddressQuotaResultOutput) QuotaSets() GetAddressQuotaQuotaSetArrayOutput {
	return o.ApplyT(func(v GetAddressQuotaResult) []GetAddressQuotaQuotaSet { return v.QuotaSets }).(GetAddressQuotaQuotaSetArrayOutput)
}

func (o GetAddressQuotaResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressQuotaResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddressQuotaResultOutput{})
}
