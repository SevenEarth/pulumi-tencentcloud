// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWafInfos(ctx *pulumi.Context, args *GetWafInfosArgs, opts ...pulumi.InvokeOption) (*GetWafInfosResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetWafInfosResult
	err := ctx.Invoke("tencentcloud:Waf/getWafInfos:getWafInfos", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafInfos.
type GetWafInfosArgs struct {
	Params           []GetWafInfosParam `pulumi:"params"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

// A collection of values returned by getWafInfos.
type GetWafInfosResult struct {
	HostLists []GetWafInfosHostList `pulumi:"hostLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	Params           []GetWafInfosParam `pulumi:"params"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

func GetWafInfosOutput(ctx *pulumi.Context, args GetWafInfosOutputArgs, opts ...pulumi.InvokeOption) GetWafInfosResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWafInfosResult, error) {
			args := v.(GetWafInfosArgs)
			r, err := GetWafInfos(ctx, &args, opts...)
			var s GetWafInfosResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWafInfosResultOutput)
}

// A collection of arguments for invoking getWafInfos.
type GetWafInfosOutputArgs struct {
	Params           GetWafInfosParamArrayInput `pulumi:"params"`
	ResultOutputFile pulumi.StringPtrInput      `pulumi:"resultOutputFile"`
}

func (GetWafInfosOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafInfosArgs)(nil)).Elem()
}

// A collection of values returned by getWafInfos.
type GetWafInfosResultOutput struct{ *pulumi.OutputState }

func (GetWafInfosResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafInfosResult)(nil)).Elem()
}

func (o GetWafInfosResultOutput) ToGetWafInfosResultOutput() GetWafInfosResultOutput {
	return o
}

func (o GetWafInfosResultOutput) ToGetWafInfosResultOutputWithContext(ctx context.Context) GetWafInfosResultOutput {
	return o
}

func (o GetWafInfosResultOutput) HostLists() GetWafInfosHostListArrayOutput {
	return o.ApplyT(func(v GetWafInfosResult) []GetWafInfosHostList { return v.HostLists }).(GetWafInfosHostListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafInfosResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafInfosResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWafInfosResultOutput) Params() GetWafInfosParamArrayOutput {
	return o.ApplyT(func(v GetWafInfosResult) []GetWafInfosParam { return v.Params }).(GetWafInfosParamArrayOutput)
}

func (o GetWafInfosResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafInfosResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafInfosResultOutput{})
}
