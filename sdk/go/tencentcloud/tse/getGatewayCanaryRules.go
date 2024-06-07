// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of tse gatewayCanaryRules
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tse.GetGatewayCanaryRules(ctx, &tse.GetGatewayCanaryRulesArgs{
//				GatewayId: "gateway-xxxxxx",
//				ServiceId: "451a9920-e67a-4519-af41-fccac0e72005",
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
func LookupGatewayCanaryRules(ctx *pulumi.Context, args *LookupGatewayCanaryRulesArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCanaryRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayCanaryRulesResult
	err := ctx.Invoke("tencentcloud:Tse/getGatewayCanaryRules:getGatewayCanaryRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayCanaryRules.
type LookupGatewayCanaryRulesArgs struct {
	// gateway ID.
	GatewayId string `pulumi:"gatewayId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// service ID.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by getGatewayCanaryRules.
type LookupGatewayCanaryRulesResult struct {
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// canary rule configuration.
	Results []GetGatewayCanaryRulesResult `pulumi:"results"`
	// service ID.
	ServiceId string `pulumi:"serviceId"`
}

func LookupGatewayCanaryRulesOutput(ctx *pulumi.Context, args LookupGatewayCanaryRulesOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCanaryRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCanaryRulesResult, error) {
			args := v.(LookupGatewayCanaryRulesArgs)
			r, err := LookupGatewayCanaryRules(ctx, &args, opts...)
			var s LookupGatewayCanaryRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCanaryRulesResultOutput)
}

// A collection of arguments for invoking getGatewayCanaryRules.
type LookupGatewayCanaryRulesOutputArgs struct {
	// gateway ID.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// service ID.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupGatewayCanaryRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCanaryRulesArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayCanaryRules.
type LookupGatewayCanaryRulesResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCanaryRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCanaryRulesResult)(nil)).Elem()
}

func (o LookupGatewayCanaryRulesResultOutput) ToLookupGatewayCanaryRulesResultOutput() LookupGatewayCanaryRulesResultOutput {
	return o
}

func (o LookupGatewayCanaryRulesResultOutput) ToLookupGatewayCanaryRulesResultOutputWithContext(ctx context.Context) LookupGatewayCanaryRulesResultOutput {
	return o
}

func (o LookupGatewayCanaryRulesResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCanaryRulesResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayCanaryRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCanaryRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayCanaryRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayCanaryRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// canary rule configuration.
func (o LookupGatewayCanaryRulesResultOutput) Results() GetGatewayCanaryRulesResultArrayOutput {
	return o.ApplyT(func(v LookupGatewayCanaryRulesResult) []GetGatewayCanaryRulesResult { return v.Results }).(GetGatewayCanaryRulesResultArrayOutput)
}

// service ID.
func (o LookupGatewayCanaryRulesResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCanaryRulesResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCanaryRulesResultOutput{})
}
