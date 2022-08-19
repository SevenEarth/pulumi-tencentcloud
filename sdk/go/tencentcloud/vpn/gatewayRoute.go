// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a VPN gateway route.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpn.NewGatewayRoute(ctx, "route", &Vpn.GatewayRouteArgs{
// 			DestinationCidrBlock: pulumi.String("10.0.0.0/16"),
// 			InstanceId:           pulumi.String("vpnx-5b5dmao3"),
// 			InstanceType:         pulumi.String("VPNCONN"),
// 			Priority:             pulumi.Int(100),
// 			Status:               pulumi.String("DISABLE"),
// 			VpnGatewayId:         pulumi.String("vpngw-ak9sjem2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// VPN gateway route can be imported using the id, the id format must be '{vpn_gateway_id}#{route_id}', e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Vpn/gatewayRoute:GatewayRoute route1 vpngw-ak9sjem2#vpngw-8ccsnclt
// ```
type GatewayRoute struct {
	pulumi.CustomResourceState

	// Create time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Destination IDC IP range.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// Instance ID of the next hop.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Priority. Valid values: 0 and 100.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Route ID.
	RouteId pulumi.StringOutput `pulumi:"routeId"`
	// Status. Valid values: ENABLE and DISABLE.
	Status pulumi.StringOutput `pulumi:"status"`
	// Route type. Default value: Static.
	Type pulumi.StringOutput `pulumi:"type"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// VPN gateway ID.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayRoute(ctx *pulumi.Context,
	name string, args *GatewayRouteArgs, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.VpnGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VpnGatewayId'")
	}
	var resource GatewayRoute
	err := ctx.RegisterResource("tencentcloud:Vpn/gatewayRoute:GatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRoute gets an existing GatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteState, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	var resource GatewayRoute
	err := ctx.ReadResource("tencentcloud:Vpn/gatewayRoute:GatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRoute resources.
type gatewayRouteState struct {
	// Create time.
	CreateTime *string `pulumi:"createTime"`
	// Destination IDC IP range.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Instance ID of the next hop.
	InstanceId *string `pulumi:"instanceId"`
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType *string `pulumi:"instanceType"`
	// Priority. Valid values: 0 and 100.
	Priority *int `pulumi:"priority"`
	// Route ID.
	RouteId *string `pulumi:"routeId"`
	// Status. Valid values: ENABLE and DISABLE.
	Status *string `pulumi:"status"`
	// Route type. Default value: Static.
	Type *string `pulumi:"type"`
	// Update time.
	UpdateTime *string `pulumi:"updateTime"`
	// VPN gateway ID.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type GatewayRouteState struct {
	// Create time.
	CreateTime pulumi.StringPtrInput
	// Destination IDC IP range.
	DestinationCidrBlock pulumi.StringPtrInput
	// Instance ID of the next hop.
	InstanceId pulumi.StringPtrInput
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType pulumi.StringPtrInput
	// Priority. Valid values: 0 and 100.
	Priority pulumi.IntPtrInput
	// Route ID.
	RouteId pulumi.StringPtrInput
	// Status. Valid values: ENABLE and DISABLE.
	Status pulumi.StringPtrInput
	// Route type. Default value: Static.
	Type pulumi.StringPtrInput
	// Update time.
	UpdateTime pulumi.StringPtrInput
	// VPN gateway ID.
	VpnGatewayId pulumi.StringPtrInput
}

func (GatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteState)(nil)).Elem()
}

type gatewayRouteArgs struct {
	// Destination IDC IP range.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// Instance ID of the next hop.
	InstanceId string `pulumi:"instanceId"`
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType string `pulumi:"instanceType"`
	// Priority. Valid values: 0 and 100.
	Priority int `pulumi:"priority"`
	// Status. Valid values: ENABLE and DISABLE.
	Status string `pulumi:"status"`
	// VPN gateway ID.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a GatewayRoute resource.
type GatewayRouteArgs struct {
	// Destination IDC IP range.
	DestinationCidrBlock pulumi.StringInput
	// Instance ID of the next hop.
	InstanceId pulumi.StringInput
	// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
	InstanceType pulumi.StringInput
	// Priority. Valid values: 0 and 100.
	Priority pulumi.IntInput
	// Status. Valid values: ENABLE and DISABLE.
	Status pulumi.StringInput
	// VPN gateway ID.
	VpnGatewayId pulumi.StringInput
}

func (GatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteArgs)(nil)).Elem()
}

type GatewayRouteInput interface {
	pulumi.Input

	ToGatewayRouteOutput() GatewayRouteOutput
	ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput
}

func (*GatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (i *GatewayRoute) ToGatewayRouteOutput() GatewayRouteOutput {
	return i.ToGatewayRouteOutputWithContext(context.Background())
}

func (i *GatewayRoute) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteOutput)
}

// GatewayRouteArrayInput is an input type that accepts GatewayRouteArray and GatewayRouteArrayOutput values.
// You can construct a concrete instance of `GatewayRouteArrayInput` via:
//
//          GatewayRouteArray{ GatewayRouteArgs{...} }
type GatewayRouteArrayInput interface {
	pulumi.Input

	ToGatewayRouteArrayOutput() GatewayRouteArrayOutput
	ToGatewayRouteArrayOutputWithContext(context.Context) GatewayRouteArrayOutput
}

type GatewayRouteArray []GatewayRouteInput

func (GatewayRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return i.ToGatewayRouteArrayOutputWithContext(context.Background())
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteArrayOutput)
}

// GatewayRouteMapInput is an input type that accepts GatewayRouteMap and GatewayRouteMapOutput values.
// You can construct a concrete instance of `GatewayRouteMapInput` via:
//
//          GatewayRouteMap{ "key": GatewayRouteArgs{...} }
type GatewayRouteMapInput interface {
	pulumi.Input

	ToGatewayRouteMapOutput() GatewayRouteMapOutput
	ToGatewayRouteMapOutputWithContext(context.Context) GatewayRouteMapOutput
}

type GatewayRouteMap map[string]GatewayRouteInput

func (GatewayRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteMap) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return i.ToGatewayRouteMapOutputWithContext(context.Background())
}

func (i GatewayRouteMap) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteMapOutput)
}

type GatewayRouteOutput struct{ *pulumi.OutputState }

func (GatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteOutput) ToGatewayRouteOutput() GatewayRouteOutput {
	return o
}

func (o GatewayRouteOutput) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return o
}

// Create time.
func (o GatewayRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Destination IDC IP range.
func (o GatewayRouteOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// Instance ID of the next hop.
func (o GatewayRouteOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
func (o GatewayRouteOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// Priority. Valid values: 0 and 100.
func (o GatewayRouteOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Route ID.
func (o GatewayRouteOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

// Status. Valid values: ENABLE and DISABLE.
func (o GatewayRouteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Route type. Default value: Static.
func (o GatewayRouteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Update time.
func (o GatewayRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// VPN gateway ID.
func (o GatewayRouteOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.VpnGatewayId }).(pulumi.StringOutput)
}

type GatewayRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) Index(i pulumi.IntInput) GatewayRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].([]*GatewayRoute)[vs[1].(int)]
	}).(GatewayRouteOutput)
}

type GatewayRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].(map[string]*GatewayRoute)[vs[1].(string)]
	}).(GatewayRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteInput)(nil)).Elem(), &GatewayRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteArrayInput)(nil)).Elem(), GatewayRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteMapInput)(nil)).Elem(), GatewayRouteMap{})
	pulumi.RegisterOutputType(GatewayRouteOutput{})
	pulumi.RegisterOutputType(GatewayRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteMapOutput{})
}
