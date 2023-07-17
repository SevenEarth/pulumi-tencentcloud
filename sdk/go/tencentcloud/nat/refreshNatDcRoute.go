// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc refreshNatDcRoute
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Nat"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Nat.NewRefreshNatDcRoute(ctx, "refreshNatDcRoute", &Nat.RefreshNatDcRouteArgs{
// 			DryRun:       pulumi.Bool(true),
// 			NatGatewayId: pulumi.String("nat-gnxkey2e"),
// 			VpcId:        pulumi.String("vpc-pyyv5k3v"),
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
// vpc refresh_nat_dc_route can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Nat/refreshNatDcRoute:RefreshNatDcRoute refresh_nat_dc_route vpc_id#nat_gateway_id
// ```
type RefreshNatDcRoute struct {
	pulumi.CustomResourceState

	// Whether to pre-refresh, valid values: True:yes, False:no.
	DryRun pulumi.BoolOutput `pulumi:"dryRun"`
	// Unique identifier of Nat Gateway.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// Unique identifier of Vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRefreshNatDcRoute registers a new resource with the given unique name, arguments, and options.
func NewRefreshNatDcRoute(ctx *pulumi.Context,
	name string, args *RefreshNatDcRouteArgs, opts ...pulumi.ResourceOption) (*RefreshNatDcRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DryRun == nil {
		return nil, errors.New("invalid value for required argument 'DryRun'")
	}
	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RefreshNatDcRoute
	err := ctx.RegisterResource("tencentcloud:Nat/refreshNatDcRoute:RefreshNatDcRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRefreshNatDcRoute gets an existing RefreshNatDcRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRefreshNatDcRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RefreshNatDcRouteState, opts ...pulumi.ResourceOption) (*RefreshNatDcRoute, error) {
	var resource RefreshNatDcRoute
	err := ctx.ReadResource("tencentcloud:Nat/refreshNatDcRoute:RefreshNatDcRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RefreshNatDcRoute resources.
type refreshNatDcRouteState struct {
	// Whether to pre-refresh, valid values: True:yes, False:no.
	DryRun *bool `pulumi:"dryRun"`
	// Unique identifier of Nat Gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Unique identifier of Vpc.
	VpcId *string `pulumi:"vpcId"`
}

type RefreshNatDcRouteState struct {
	// Whether to pre-refresh, valid values: True:yes, False:no.
	DryRun pulumi.BoolPtrInput
	// Unique identifier of Nat Gateway.
	NatGatewayId pulumi.StringPtrInput
	// Unique identifier of Vpc.
	VpcId pulumi.StringPtrInput
}

func (RefreshNatDcRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshNatDcRouteState)(nil)).Elem()
}

type refreshNatDcRouteArgs struct {
	// Whether to pre-refresh, valid values: True:yes, False:no.
	DryRun bool `pulumi:"dryRun"`
	// Unique identifier of Nat Gateway.
	NatGatewayId string `pulumi:"natGatewayId"`
	// Unique identifier of Vpc.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RefreshNatDcRoute resource.
type RefreshNatDcRouteArgs struct {
	// Whether to pre-refresh, valid values: True:yes, False:no.
	DryRun pulumi.BoolInput
	// Unique identifier of Nat Gateway.
	NatGatewayId pulumi.StringInput
	// Unique identifier of Vpc.
	VpcId pulumi.StringInput
}

func (RefreshNatDcRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshNatDcRouteArgs)(nil)).Elem()
}

type RefreshNatDcRouteInput interface {
	pulumi.Input

	ToRefreshNatDcRouteOutput() RefreshNatDcRouteOutput
	ToRefreshNatDcRouteOutputWithContext(ctx context.Context) RefreshNatDcRouteOutput
}

func (*RefreshNatDcRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshNatDcRoute)(nil)).Elem()
}

func (i *RefreshNatDcRoute) ToRefreshNatDcRouteOutput() RefreshNatDcRouteOutput {
	return i.ToRefreshNatDcRouteOutputWithContext(context.Background())
}

func (i *RefreshNatDcRoute) ToRefreshNatDcRouteOutputWithContext(ctx context.Context) RefreshNatDcRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshNatDcRouteOutput)
}

// RefreshNatDcRouteArrayInput is an input type that accepts RefreshNatDcRouteArray and RefreshNatDcRouteArrayOutput values.
// You can construct a concrete instance of `RefreshNatDcRouteArrayInput` via:
//
//          RefreshNatDcRouteArray{ RefreshNatDcRouteArgs{...} }
type RefreshNatDcRouteArrayInput interface {
	pulumi.Input

	ToRefreshNatDcRouteArrayOutput() RefreshNatDcRouteArrayOutput
	ToRefreshNatDcRouteArrayOutputWithContext(context.Context) RefreshNatDcRouteArrayOutput
}

type RefreshNatDcRouteArray []RefreshNatDcRouteInput

func (RefreshNatDcRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RefreshNatDcRoute)(nil)).Elem()
}

func (i RefreshNatDcRouteArray) ToRefreshNatDcRouteArrayOutput() RefreshNatDcRouteArrayOutput {
	return i.ToRefreshNatDcRouteArrayOutputWithContext(context.Background())
}

func (i RefreshNatDcRouteArray) ToRefreshNatDcRouteArrayOutputWithContext(ctx context.Context) RefreshNatDcRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshNatDcRouteArrayOutput)
}

// RefreshNatDcRouteMapInput is an input type that accepts RefreshNatDcRouteMap and RefreshNatDcRouteMapOutput values.
// You can construct a concrete instance of `RefreshNatDcRouteMapInput` via:
//
//          RefreshNatDcRouteMap{ "key": RefreshNatDcRouteArgs{...} }
type RefreshNatDcRouteMapInput interface {
	pulumi.Input

	ToRefreshNatDcRouteMapOutput() RefreshNatDcRouteMapOutput
	ToRefreshNatDcRouteMapOutputWithContext(context.Context) RefreshNatDcRouteMapOutput
}

type RefreshNatDcRouteMap map[string]RefreshNatDcRouteInput

func (RefreshNatDcRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RefreshNatDcRoute)(nil)).Elem()
}

func (i RefreshNatDcRouteMap) ToRefreshNatDcRouteMapOutput() RefreshNatDcRouteMapOutput {
	return i.ToRefreshNatDcRouteMapOutputWithContext(context.Background())
}

func (i RefreshNatDcRouteMap) ToRefreshNatDcRouteMapOutputWithContext(ctx context.Context) RefreshNatDcRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshNatDcRouteMapOutput)
}

type RefreshNatDcRouteOutput struct{ *pulumi.OutputState }

func (RefreshNatDcRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshNatDcRoute)(nil)).Elem()
}

func (o RefreshNatDcRouteOutput) ToRefreshNatDcRouteOutput() RefreshNatDcRouteOutput {
	return o
}

func (o RefreshNatDcRouteOutput) ToRefreshNatDcRouteOutputWithContext(ctx context.Context) RefreshNatDcRouteOutput {
	return o
}

// Whether to pre-refresh, valid values: True:yes, False:no.
func (o RefreshNatDcRouteOutput) DryRun() pulumi.BoolOutput {
	return o.ApplyT(func(v *RefreshNatDcRoute) pulumi.BoolOutput { return v.DryRun }).(pulumi.BoolOutput)
}

// Unique identifier of Nat Gateway.
func (o RefreshNatDcRouteOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *RefreshNatDcRoute) pulumi.StringOutput { return v.NatGatewayId }).(pulumi.StringOutput)
}

// Unique identifier of Vpc.
func (o RefreshNatDcRouteOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RefreshNatDcRoute) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RefreshNatDcRouteArrayOutput struct{ *pulumi.OutputState }

func (RefreshNatDcRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RefreshNatDcRoute)(nil)).Elem()
}

func (o RefreshNatDcRouteArrayOutput) ToRefreshNatDcRouteArrayOutput() RefreshNatDcRouteArrayOutput {
	return o
}

func (o RefreshNatDcRouteArrayOutput) ToRefreshNatDcRouteArrayOutputWithContext(ctx context.Context) RefreshNatDcRouteArrayOutput {
	return o
}

func (o RefreshNatDcRouteArrayOutput) Index(i pulumi.IntInput) RefreshNatDcRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RefreshNatDcRoute {
		return vs[0].([]*RefreshNatDcRoute)[vs[1].(int)]
	}).(RefreshNatDcRouteOutput)
}

type RefreshNatDcRouteMapOutput struct{ *pulumi.OutputState }

func (RefreshNatDcRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RefreshNatDcRoute)(nil)).Elem()
}

func (o RefreshNatDcRouteMapOutput) ToRefreshNatDcRouteMapOutput() RefreshNatDcRouteMapOutput {
	return o
}

func (o RefreshNatDcRouteMapOutput) ToRefreshNatDcRouteMapOutputWithContext(ctx context.Context) RefreshNatDcRouteMapOutput {
	return o
}

func (o RefreshNatDcRouteMapOutput) MapIndex(k pulumi.StringInput) RefreshNatDcRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RefreshNatDcRoute {
		return vs[0].(map[string]*RefreshNatDcRoute)[vs[1].(string)]
	}).(RefreshNatDcRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RefreshNatDcRouteInput)(nil)).Elem(), &RefreshNatDcRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*RefreshNatDcRouteArrayInput)(nil)).Elem(), RefreshNatDcRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RefreshNatDcRouteMapInput)(nil)).Elem(), RefreshNatDcRouteMap{})
	pulumi.RegisterOutputType(RefreshNatDcRouteOutput{})
	pulumi.RegisterOutputType(RefreshNatDcRouteArrayOutput{})
	pulumi.RegisterOutputType(RefreshNatDcRouteMapOutput{})
}
