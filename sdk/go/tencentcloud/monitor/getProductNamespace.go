// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query product namespace in monitor)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.GetProductNamespace(ctx, &monitor.GetProductNamespaceArgs{
//				Name: pulumi.StringRef("Redis"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProductNamespace(ctx *pulumi.Context, args *GetProductNamespaceArgs, opts ...pulumi.InvokeOption) (*GetProductNamespaceResult, error) {
	var rv GetProductNamespaceResult
	err := ctx.Invoke("tencentcloud:Monitor/getProductNamespace:getProductNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProductNamespace.
type GetProductNamespaceArgs struct {
	// Name for filter, eg:`Load Banlancer`.
	Name *string `pulumi:"name"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getProductNamespace.
type GetProductNamespaceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list product namespaces. Each element contains the following attributes:
	Lists            []GetProductNamespaceList `pulumi:"lists"`
	Name             *string                   `pulumi:"name"`
	ResultOutputFile *string                   `pulumi:"resultOutputFile"`
}

func GetProductNamespaceOutput(ctx *pulumi.Context, args GetProductNamespaceOutputArgs, opts ...pulumi.InvokeOption) GetProductNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductNamespaceResult, error) {
			args := v.(GetProductNamespaceArgs)
			r, err := GetProductNamespace(ctx, &args, opts...)
			var s GetProductNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductNamespaceResultOutput)
}

// A collection of arguments for invoking getProductNamespace.
type GetProductNamespaceOutputArgs struct {
	// Name for filter, eg:`Load Banlancer`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetProductNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getProductNamespace.
type GetProductNamespaceResultOutput struct{ *pulumi.OutputState }

func (GetProductNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductNamespaceResult)(nil)).Elem()
}

func (o GetProductNamespaceResultOutput) ToGetProductNamespaceResultOutput() GetProductNamespaceResultOutput {
	return o
}

func (o GetProductNamespaceResultOutput) ToGetProductNamespaceResultOutputWithContext(ctx context.Context) GetProductNamespaceResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProductNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list product namespaces. Each element contains the following attributes:
func (o GetProductNamespaceResultOutput) Lists() GetProductNamespaceListArrayOutput {
	return o.ApplyT(func(v GetProductNamespaceResult) []GetProductNamespaceList { return v.Lists }).(GetProductNamespaceListArrayOutput)
}

func (o GetProductNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetProductNamespaceResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductNamespaceResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductNamespaceResultOutput{})
}
