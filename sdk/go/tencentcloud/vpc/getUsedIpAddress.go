// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpc usedIpAddress
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpc.GetUsedIpAddress(ctx, &vpc.GetUsedIpAddressArgs{
// 			VpcId: "vpc-4owdpnwr",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetUsedIpAddress(ctx *pulumi.Context, args *GetUsedIpAddressArgs, opts ...pulumi.InvokeOption) (*GetUsedIpAddressResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetUsedIpAddressResult
	err := ctx.Invoke("tencentcloud:Vpc/getUsedIpAddress:getUsedIpAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsedIpAddress.
type GetUsedIpAddressArgs struct {
	// IPs to query.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Subnet instance ID.
	SubnetId *string `pulumi:"subnetId"`
	// VPC instance ID.
	VpcId string `pulumi:"vpcId"`
}

// A collection of values returned by getUsedIpAddress.
type GetUsedIpAddressResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Information of resources bound with the queried IPs Note: This parameter may return null, indicating that no valid values can be obtained.
	IpAddressStates  []GetUsedIpAddressIpAddressState `pulumi:"ipAddressStates"`
	IpAddresses      []string                         `pulumi:"ipAddresses"`
	ResultOutputFile *string                          `pulumi:"resultOutputFile"`
	// Subnet instance ID.
	SubnetId *string `pulumi:"subnetId"`
	// VPC instance ID.
	VpcId string `pulumi:"vpcId"`
}

func GetUsedIpAddressOutput(ctx *pulumi.Context, args GetUsedIpAddressOutputArgs, opts ...pulumi.InvokeOption) GetUsedIpAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsedIpAddressResult, error) {
			args := v.(GetUsedIpAddressArgs)
			r, err := GetUsedIpAddress(ctx, &args, opts...)
			var s GetUsedIpAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsedIpAddressResultOutput)
}

// A collection of arguments for invoking getUsedIpAddress.
type GetUsedIpAddressOutputArgs struct {
	// IPs to query.
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Subnet instance ID.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// VPC instance ID.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetUsedIpAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsedIpAddressArgs)(nil)).Elem()
}

// A collection of values returned by getUsedIpAddress.
type GetUsedIpAddressResultOutput struct{ *pulumi.OutputState }

func (GetUsedIpAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsedIpAddressResult)(nil)).Elem()
}

func (o GetUsedIpAddressResultOutput) ToGetUsedIpAddressResultOutput() GetUsedIpAddressResultOutput {
	return o
}

func (o GetUsedIpAddressResultOutput) ToGetUsedIpAddressResultOutputWithContext(ctx context.Context) GetUsedIpAddressResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsedIpAddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) string { return v.Id }).(pulumi.StringOutput)
}

// Information of resources bound with the queried IPs Note: This parameter may return null, indicating that no valid values can be obtained.
func (o GetUsedIpAddressResultOutput) IpAddressStates() GetUsedIpAddressIpAddressStateArrayOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) []GetUsedIpAddressIpAddressState { return v.IpAddressStates }).(GetUsedIpAddressIpAddressStateArrayOutput)
}

func (o GetUsedIpAddressResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o GetUsedIpAddressResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Subnet instance ID.
func (o GetUsedIpAddressResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// VPC instance ID.
func (o GetUsedIpAddressResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsedIpAddressResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsedIpAddressResultOutput{})
}
