// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpc vpnCustomerGatewayVendors
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpn.GetCustomerGatewayVendors(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCustomerGatewayVendors(ctx *pulumi.Context, args *GetCustomerGatewayVendorsArgs, opts ...pulumi.InvokeOption) (*GetCustomerGatewayVendorsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCustomerGatewayVendorsResult
	err := ctx.Invoke("tencentcloud:Vpn/getCustomerGatewayVendors:getCustomerGatewayVendors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerGatewayVendors.
type GetCustomerGatewayVendorsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCustomerGatewayVendors.
type GetCustomerGatewayVendorsResult struct {
	// Customer Gateway Vendor Set.
	CustomerGatewayVendorSets []GetCustomerGatewayVendorsCustomerGatewayVendorSet `pulumi:"customerGatewayVendorSets"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCustomerGatewayVendorsOutput(ctx *pulumi.Context, args GetCustomerGatewayVendorsOutputArgs, opts ...pulumi.InvokeOption) GetCustomerGatewayVendorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomerGatewayVendorsResult, error) {
			args := v.(GetCustomerGatewayVendorsArgs)
			r, err := GetCustomerGatewayVendors(ctx, &args, opts...)
			var s GetCustomerGatewayVendorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomerGatewayVendorsResultOutput)
}

// A collection of arguments for invoking getCustomerGatewayVendors.
type GetCustomerGatewayVendorsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCustomerGatewayVendorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerGatewayVendorsArgs)(nil)).Elem()
}

// A collection of values returned by getCustomerGatewayVendors.
type GetCustomerGatewayVendorsResultOutput struct{ *pulumi.OutputState }

func (GetCustomerGatewayVendorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerGatewayVendorsResult)(nil)).Elem()
}

func (o GetCustomerGatewayVendorsResultOutput) ToGetCustomerGatewayVendorsResultOutput() GetCustomerGatewayVendorsResultOutput {
	return o
}

func (o GetCustomerGatewayVendorsResultOutput) ToGetCustomerGatewayVendorsResultOutputWithContext(ctx context.Context) GetCustomerGatewayVendorsResultOutput {
	return o
}

// Customer Gateway Vendor Set.
func (o GetCustomerGatewayVendorsResultOutput) CustomerGatewayVendorSets() GetCustomerGatewayVendorsCustomerGatewayVendorSetArrayOutput {
	return o.ApplyT(func(v GetCustomerGatewayVendorsResult) []GetCustomerGatewayVendorsCustomerGatewayVendorSet {
		return v.CustomerGatewayVendorSets
	}).(GetCustomerGatewayVendorsCustomerGatewayVendorSetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomerGatewayVendorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerGatewayVendorsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCustomerGatewayVendorsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomerGatewayVendorsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomerGatewayVendorsResultOutput{})
}
