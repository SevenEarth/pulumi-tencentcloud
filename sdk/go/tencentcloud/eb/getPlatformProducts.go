// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of eb platformProducts
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Eb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Eb.GetPlatformProducts(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPlatformProducts(ctx *pulumi.Context, args *GetPlatformProductsArgs, opts ...pulumi.InvokeOption) (*GetPlatformProductsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPlatformProductsResult
	err := ctx.Invoke("tencentcloud:Eb/getPlatformProducts:getPlatformProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlatformProducts.
type GetPlatformProductsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getPlatformProducts.
type GetPlatformProductsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Platform product list.
	PlatformProducts []GetPlatformProductsPlatformProduct `pulumi:"platformProducts"`
	ResultOutputFile *string                              `pulumi:"resultOutputFile"`
}

func GetPlatformProductsOutput(ctx *pulumi.Context, args GetPlatformProductsOutputArgs, opts ...pulumi.InvokeOption) GetPlatformProductsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPlatformProductsResult, error) {
			args := v.(GetPlatformProductsArgs)
			r, err := GetPlatformProducts(ctx, &args, opts...)
			var s GetPlatformProductsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPlatformProductsResultOutput)
}

// A collection of arguments for invoking getPlatformProducts.
type GetPlatformProductsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetPlatformProductsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPlatformProductsArgs)(nil)).Elem()
}

// A collection of values returned by getPlatformProducts.
type GetPlatformProductsResultOutput struct{ *pulumi.OutputState }

func (GetPlatformProductsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPlatformProductsResult)(nil)).Elem()
}

func (o GetPlatformProductsResultOutput) ToGetPlatformProductsResultOutput() GetPlatformProductsResultOutput {
	return o
}

func (o GetPlatformProductsResultOutput) ToGetPlatformProductsResultOutputWithContext(ctx context.Context) GetPlatformProductsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPlatformProductsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPlatformProductsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Platform product list.
func (o GetPlatformProductsResultOutput) PlatformProducts() GetPlatformProductsPlatformProductArrayOutput {
	return o.ApplyT(func(v GetPlatformProductsResult) []GetPlatformProductsPlatformProduct { return v.PlatformProducts }).(GetPlatformProductsPlatformProductArrayOutput)
}

func (o GetPlatformProductsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlatformProductsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPlatformProductsResultOutput{})
}
