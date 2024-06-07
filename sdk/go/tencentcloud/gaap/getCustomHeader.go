// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of gaap custom header
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Gaap.GetCustomHeader(ctx, &gaap.GetCustomHeaderArgs{
//				RuleId: "rule-hddrxgpd",
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
func LookupCustomHeader(ctx *pulumi.Context, args *LookupCustomHeaderArgs, opts ...pulumi.InvokeOption) (*LookupCustomHeaderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomHeaderResult
	err := ctx.Invoke("tencentcloud:Gaap/getCustomHeader:getCustomHeader", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomHeader.
type LookupCustomHeaderArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Rule IdNote: This field may return null, indicating that a valid value cannot be obtained.
	RuleId string `pulumi:"ruleId"`
}

// A collection of values returned by getCustomHeader.
type LookupCustomHeaderResult struct {
	// HeadersNote: This field may return null, indicating that a valid value cannot be obtained.
	Headers []GetCustomHeaderHeader `pulumi:"headers"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RuleId           string  `pulumi:"ruleId"`
}

func LookupCustomHeaderOutput(ctx *pulumi.Context, args LookupCustomHeaderOutputArgs, opts ...pulumi.InvokeOption) LookupCustomHeaderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomHeaderResult, error) {
			args := v.(LookupCustomHeaderArgs)
			r, err := LookupCustomHeader(ctx, &args, opts...)
			var s LookupCustomHeaderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomHeaderResultOutput)
}

// A collection of arguments for invoking getCustomHeader.
type LookupCustomHeaderOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Rule IdNote: This field may return null, indicating that a valid value cannot be obtained.
	RuleId pulumi.StringInput `pulumi:"ruleId"`
}

func (LookupCustomHeaderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomHeaderArgs)(nil)).Elem()
}

// A collection of values returned by getCustomHeader.
type LookupCustomHeaderResultOutput struct{ *pulumi.OutputState }

func (LookupCustomHeaderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomHeaderResult)(nil)).Elem()
}

func (o LookupCustomHeaderResultOutput) ToLookupCustomHeaderResultOutput() LookupCustomHeaderResultOutput {
	return o
}

func (o LookupCustomHeaderResultOutput) ToLookupCustomHeaderResultOutputWithContext(ctx context.Context) LookupCustomHeaderResultOutput {
	return o
}

// HeadersNote: This field may return null, indicating that a valid value cannot be obtained.
func (o LookupCustomHeaderResultOutput) Headers() GetCustomHeaderHeaderArrayOutput {
	return o.ApplyT(func(v LookupCustomHeaderResult) []GetCustomHeaderHeader { return v.Headers }).(GetCustomHeaderHeaderArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCustomHeaderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomHeaderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomHeaderResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomHeaderResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupCustomHeaderResultOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomHeaderResult) string { return v.RuleId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomHeaderResultOutput{})
}
