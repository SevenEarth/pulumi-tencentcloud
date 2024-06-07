// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of vpc accountAttributes
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.GetAccountAttributes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAccountAttributes(ctx *pulumi.Context, args *GetAccountAttributesArgs, opts ...pulumi.InvokeOption) (*GetAccountAttributesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountAttributesResult
	err := ctx.Invoke("tencentcloud:Vpc/getAccountAttributes:getAccountAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountAttributes.
type GetAccountAttributesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAccountAttributes.
type GetAccountAttributesResult struct {
	// User account attribute object.
	AccountAttributeSets []GetAccountAttributesAccountAttributeSet `pulumi:"accountAttributeSets"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetAccountAttributesOutput(ctx *pulumi.Context, args GetAccountAttributesOutputArgs, opts ...pulumi.InvokeOption) GetAccountAttributesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountAttributesResult, error) {
			args := v.(GetAccountAttributesArgs)
			r, err := GetAccountAttributes(ctx, &args, opts...)
			var s GetAccountAttributesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccountAttributesResultOutput)
}

// A collection of arguments for invoking getAccountAttributes.
type GetAccountAttributesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAccountAttributesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAttributesArgs)(nil)).Elem()
}

// A collection of values returned by getAccountAttributes.
type GetAccountAttributesResultOutput struct{ *pulumi.OutputState }

func (GetAccountAttributesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAttributesResult)(nil)).Elem()
}

func (o GetAccountAttributesResultOutput) ToGetAccountAttributesResultOutput() GetAccountAttributesResultOutput {
	return o
}

func (o GetAccountAttributesResultOutput) ToGetAccountAttributesResultOutputWithContext(ctx context.Context) GetAccountAttributesResultOutput {
	return o
}

// User account attribute object.
func (o GetAccountAttributesResultOutput) AccountAttributeSets() GetAccountAttributesAccountAttributeSetArrayOutput {
	return o.ApplyT(func(v GetAccountAttributesResult) []GetAccountAttributesAccountAttributeSet {
		return v.AccountAttributeSets
	}).(GetAccountAttributesAccountAttributeSetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountAttributesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAttributesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountAttributesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountAttributesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountAttributesResultOutput{})
}
