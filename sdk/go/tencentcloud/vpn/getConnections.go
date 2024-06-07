// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of VPN connections.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpn.GetConnections(ctx, &vpn.GetConnectionsArgs{
//				CustomerGatewayId: pulumi.StringRef(""),
//				Id:                pulumi.StringRef("vpnx-xfqag"),
//				Name:              pulumi.StringRef("main"),
//				Tags: map[string]interface{}{
//					"test": "tf",
//				},
//				VpcId:        pulumi.StringRef("cgw-xfqag"),
//				VpnGatewayId: pulumi.StringRef("vpngw-8ccsnclt"),
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
func GetConnections(ctx *pulumi.Context, args *GetConnectionsArgs, opts ...pulumi.InvokeOption) (*GetConnectionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetConnectionsResult
	err := ctx.Invoke("tencentcloud:Vpn/getConnections:getConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConnections.
type GetConnectionsArgs struct {
	// Customer gateway ID of the VPN connection.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// ID of the VPN connection.
	Id *string `pulumi:"id"`
	// Name of the VPN connection. The length of character is limited to 1-60.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tags of the VPN connection to be queried.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// VPN gateway ID of the VPN connection.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// A collection of values returned by getConnections.
type GetConnectionsResult struct {
	// Information list of the dedicated connections.
	ConnectionLists []GetConnectionsConnectionList `pulumi:"connectionLists"`
	// ID of the customer gateway.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// ID of the VPN connection.
	Id *string `pulumi:"id"`
	// Name of the VPN connection.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// A list of tags used to associate different resources.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// ID of the VPN gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

func GetConnectionsOutput(ctx *pulumi.Context, args GetConnectionsOutputArgs, opts ...pulumi.InvokeOption) GetConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConnectionsResult, error) {
			args := v.(GetConnectionsArgs)
			r, err := GetConnections(ctx, &args, opts...)
			var s GetConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetConnectionsResultOutput)
}

// A collection of arguments for invoking getConnections.
type GetConnectionsOutputArgs struct {
	// Customer gateway ID of the VPN connection.
	CustomerGatewayId pulumi.StringPtrInput `pulumi:"customerGatewayId"`
	// ID of the VPN connection.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the VPN connection. The length of character is limited to 1-60.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Tags of the VPN connection to be queried.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the VPC.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// VPN gateway ID of the VPN connection.
	VpnGatewayId pulumi.StringPtrInput `pulumi:"vpnGatewayId"`
}

func (GetConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConnectionsArgs)(nil)).Elem()
}

// A collection of values returned by getConnections.
type GetConnectionsResultOutput struct{ *pulumi.OutputState }

func (GetConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConnectionsResult)(nil)).Elem()
}

func (o GetConnectionsResultOutput) ToGetConnectionsResultOutput() GetConnectionsResultOutput {
	return o
}

func (o GetConnectionsResultOutput) ToGetConnectionsResultOutputWithContext(ctx context.Context) GetConnectionsResultOutput {
	return o
}

// Information list of the dedicated connections.
func (o GetConnectionsResultOutput) ConnectionLists() GetConnectionsConnectionListArrayOutput {
	return o.ApplyT(func(v GetConnectionsResult) []GetConnectionsConnectionList { return v.ConnectionLists }).(GetConnectionsConnectionListArrayOutput)
}

// ID of the customer gateway.
func (o GetConnectionsResultOutput) CustomerGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.CustomerGatewayId }).(pulumi.StringPtrOutput)
}

// ID of the VPN connection.
func (o GetConnectionsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the VPN connection.
func (o GetConnectionsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetConnectionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// A list of tags used to associate different resources.
func (o GetConnectionsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetConnectionsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// ID of the VPC.
func (o GetConnectionsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// ID of the VPN gateway.
func (o GetConnectionsResultOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectionsResult) *string { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConnectionsResultOutput{})
}
