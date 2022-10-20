// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of VPN gateways.
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
//			_, err := Vpn.GetGateways(ctx, &vpn.GetGatewaysArgs{
//				DestinationCidrBlock: "vpngw-8ccsnclt",
//				InstanceId:           "ap-guangzhou-3",
//				InstanceType:         "1.1.1.1",
//				Tags: map[string]interface{}{
//					"test": "tf",
//				},
//				VpnGatewayId: "main",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetGatewayRoutes(ctx *pulumi.Context, args *GetGatewayRoutesArgs, opts ...pulumi.InvokeOption) (*GetGatewayRoutesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetGatewayRoutesResult
	err := ctx.Invoke("tencentcloud:Vpn/getGatewayRoutes:getGatewayRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayRoutes.
type GetGatewayRoutesArgs struct {
	// Destination IDC IP range.
	DestinationCidr *string `pulumi:"destinationCidr"`
	// Instance ID of the next hop.
	InstanceId *string `pulumi:"instanceId"`
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType *string `pulumi:"instanceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// VPN gateway ID.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// A collection of values returned by getGatewayRoutes.
type GetGatewayRoutesResult struct {
	DestinationCidr *string `pulumi:"destinationCidr"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       *string `pulumi:"instanceId"`
	InstanceType     *string `pulumi:"instanceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	VpnGatewayId     string  `pulumi:"vpnGatewayId"`
	// Information list of the vpn gateway routes.
	VpnGatewayRouteLists []GetGatewayRoutesVpnGatewayRouteList `pulumi:"vpnGatewayRouteLists"`
}

func GetGatewayRoutesOutput(ctx *pulumi.Context, args GetGatewayRoutesOutputArgs, opts ...pulumi.InvokeOption) GetGatewayRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewayRoutesResult, error) {
			args := v.(GetGatewayRoutesArgs)
			r, err := GetGatewayRoutes(ctx, &args, opts...)
			var s GetGatewayRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewayRoutesResultOutput)
}

// A collection of arguments for invoking getGatewayRoutes.
type GetGatewayRoutesOutputArgs struct {
	// Destination IDC IP range.
	DestinationCidr pulumi.StringPtrInput `pulumi:"destinationCidr"`
	// Instance ID of the next hop.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// VPN gateway ID.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

func (GetGatewayRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayRoutes.
type GetGatewayRoutesResultOutput struct{ *pulumi.OutputState }

func (GetGatewayRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayRoutesResult)(nil)).Elem()
}

func (o GetGatewayRoutesResultOutput) ToGetGatewayRoutesResultOutput() GetGatewayRoutesResultOutput {
	return o
}

func (o GetGatewayRoutesResultOutput) ToGetGatewayRoutesResultOutputWithContext(ctx context.Context) GetGatewayRoutesResultOutput {
	return o
}

func (o GetGatewayRoutesResultOutput) DestinationCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) *string { return v.DestinationCidr }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewayRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewayRoutesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetGatewayRoutesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GetGatewayRoutesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGatewayRoutesResultOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) string { return v.VpnGatewayId }).(pulumi.StringOutput)
}

// Information list of the vpn gateway routes.
func (o GetGatewayRoutesResultOutput) VpnGatewayRouteLists() GetGatewayRoutesVpnGatewayRouteListArrayOutput {
	return o.ApplyT(func(v GetGatewayRoutesResult) []GetGatewayRoutesVpnGatewayRouteList { return v.VpnGatewayRouteLists }).(GetGatewayRoutesVpnGatewayRouteListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewayRoutesResultOutput{})
}
