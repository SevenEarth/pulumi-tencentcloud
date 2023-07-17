// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tsf usableUnitNamespaces
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tsf.GetUsableUnitNamespaces(ctx, &tsf.GetUsableUnitNamespacesArgs{
// 			SearchWord: pulumi.StringRef(""),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupUsableUnitNamespaces(ctx *pulumi.Context, args *LookupUsableUnitNamespacesArgs, opts ...pulumi.InvokeOption) (*LookupUsableUnitNamespacesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupUsableUnitNamespacesResult
	err := ctx.Invoke("tencentcloud:Tsf/getUsableUnitNamespaces:getUsableUnitNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsableUnitNamespaces.
type LookupUsableUnitNamespacesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// search by namespace id or namespace Name.
	SearchWord *string `pulumi:"searchWord"`
}

// A collection of values returned by getUsableUnitNamespaces.
type LookupUsableUnitNamespacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// namespace object list.
	Results    []GetUsableUnitNamespacesResult `pulumi:"results"`
	SearchWord *string                         `pulumi:"searchWord"`
}

func LookupUsableUnitNamespacesOutput(ctx *pulumi.Context, args LookupUsableUnitNamespacesOutputArgs, opts ...pulumi.InvokeOption) LookupUsableUnitNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUsableUnitNamespacesResult, error) {
			args := v.(LookupUsableUnitNamespacesArgs)
			r, err := LookupUsableUnitNamespaces(ctx, &args, opts...)
			var s LookupUsableUnitNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUsableUnitNamespacesResultOutput)
}

// A collection of arguments for invoking getUsableUnitNamespaces.
type LookupUsableUnitNamespacesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// search by namespace id or namespace Name.
	SearchWord pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (LookupUsableUnitNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsableUnitNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getUsableUnitNamespaces.
type LookupUsableUnitNamespacesResultOutput struct{ *pulumi.OutputState }

func (LookupUsableUnitNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsableUnitNamespacesResult)(nil)).Elem()
}

func (o LookupUsableUnitNamespacesResultOutput) ToLookupUsableUnitNamespacesResultOutput() LookupUsableUnitNamespacesResultOutput {
	return o
}

func (o LookupUsableUnitNamespacesResultOutput) ToLookupUsableUnitNamespacesResultOutputWithContext(ctx context.Context) LookupUsableUnitNamespacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUsableUnitNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUsableUnitNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUsableUnitNamespacesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsableUnitNamespacesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// namespace object list.
func (o LookupUsableUnitNamespacesResultOutput) Results() GetUsableUnitNamespacesResultArrayOutput {
	return o.ApplyT(func(v LookupUsableUnitNamespacesResult) []GetUsableUnitNamespacesResult { return v.Results }).(GetUsableUnitNamespacesResultArrayOutput)
}

func (o LookupUsableUnitNamespacesResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsableUnitNamespacesResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUsableUnitNamespacesResultOutput{})
}
