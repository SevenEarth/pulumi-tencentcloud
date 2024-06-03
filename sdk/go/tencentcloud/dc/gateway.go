// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to creating direct connect gateway instance.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := Vpc.NewInstance(ctx, "main", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dc.NewGateway(ctx, "vpcMain", &Dc.GatewayArgs{
//				NetworkInstanceId: main.ID(),
//				NetworkType:       pulumi.String("VPC"),
//				GatewayType:       pulumi.String("NAT"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Direct connect gateway instance can be imported, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Dc/gateway:Gateway instance dcg-id
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
	CnnRouteType pulumi.StringOutput `pulumi:"cnnRouteType"`
	// Creation time of resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether the BGP is enabled.
	EnableBgp pulumi.BoolOutput `pulumi:"enableBgp"`
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType pulumi.StringPtrOutput `pulumi:"gatewayType"`
	// Name of the DCG.
	Name pulumi.StringOutput `pulumi:"name"`
	// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceId pulumi.StringOutput `pulumi:"networkInstanceId"`
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInstanceId'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("tencentcloud:Dc/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("tencentcloud:Dc/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
	CnnRouteType *string `pulumi:"cnnRouteType"`
	// Creation time of resource.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether the BGP is enabled.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType *string `pulumi:"gatewayType"`
	// Name of the DCG.
	Name *string `pulumi:"name"`
	// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceId *string `pulumi:"networkInstanceId"`
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType *string `pulumi:"networkType"`
}

type GatewayState struct {
	// Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
	CnnRouteType pulumi.StringPtrInput
	// Creation time of resource.
	CreateTime pulumi.StringPtrInput
	// Indicates whether the BGP is enabled.
	EnableBgp pulumi.BoolPtrInput
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType pulumi.StringPtrInput
	// Name of the DCG.
	Name pulumi.StringPtrInput
	// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceId pulumi.StringPtrInput
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType *string `pulumi:"gatewayType"`
	// Name of the DCG.
	Name *string `pulumi:"name"`
	// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceId string `pulumi:"networkInstanceId"`
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType string `pulumi:"networkType"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	GatewayType pulumi.StringPtrInput
	// Name of the DCG.
	Name pulumi.StringPtrInput
	// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
	NetworkInstanceId pulumi.StringInput
	// Type of associated network. Valid value: `VPC` and `CCN`.
	NetworkType pulumi.StringInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//	GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//	GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
func (o GatewayOutput) CnnRouteType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.CnnRouteType }).(pulumi.StringOutput)
}

// Creation time of resource.
func (o GatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether the BGP is enabled.
func (o GatewayOutput) EnableBgp() pulumi.BoolOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolOutput { return v.EnableBgp }).(pulumi.BoolOutput)
}

// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
func (o GatewayOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.GatewayType }).(pulumi.StringPtrOutput)
}

// Name of the DCG.
func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If the `networkType` value is `VPC`, the available value is VPC ID. But when the `networkType` value is `CCN`, the available value is CCN instance ID.
func (o GatewayOutput) NetworkInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.NetworkInstanceId }).(pulumi.StringOutput)
}

// Type of associated network. Valid value: `VPC` and `CCN`.
func (o GatewayOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
