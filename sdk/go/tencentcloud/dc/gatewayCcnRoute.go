// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to creating direct connect gateway route entry.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ccn"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := Ccn.NewInstance(ctx, "main", &Ccn.InstanceArgs{
//				Description: pulumi.String("ci-temp-test-ccn-des"),
//				Qos:         pulumi.String("AG"),
//			})
//			if err != nil {
//				return err
//			}
//			ccnMain, err := Dc.NewGateway(ctx, "ccnMain", &Dc.GatewayArgs{
//				NetworkInstanceId: main.ID(),
//				NetworkType:       pulumi.String("CCN"),
//				GatewayType:       pulumi.String("NORMAL"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dc.NewGatewayCcnRoute(ctx, "route1", &Dc.GatewayCcnRouteArgs{
//				DcgId:     ccnMain.ID(),
//				CidrBlock: pulumi.String("10.1.1.0/32"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dc.NewGatewayCcnRoute(ctx, "route2", &Dc.GatewayCcnRouteArgs{
//				DcgId:     ccnMain.ID(),
//				CidrBlock: pulumi.String("192.1.1.0/32"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type GatewayCcnRoute struct {
	pulumi.CustomResourceState

	// As path list of the BGP.
	AsPaths pulumi.StringArrayOutput `pulumi:"asPaths"`
	// A network address segment of IDC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId pulumi.StringOutput `pulumi:"dcgId"`
}

// NewGatewayCcnRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayCcnRoute(ctx *pulumi.Context,
	name string, args *GatewayCcnRouteArgs, opts ...pulumi.ResourceOption) (*GatewayCcnRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.DcgId == nil {
		return nil, errors.New("invalid value for required argument 'DcgId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource GatewayCcnRoute
	err := ctx.RegisterResource("tencentcloud:Dc/gatewayCcnRoute:GatewayCcnRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayCcnRoute gets an existing GatewayCcnRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayCcnRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCcnRouteState, opts ...pulumi.ResourceOption) (*GatewayCcnRoute, error) {
	var resource GatewayCcnRoute
	err := ctx.ReadResource("tencentcloud:Dc/gatewayCcnRoute:GatewayCcnRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayCcnRoute resources.
type gatewayCcnRouteState struct {
	// As path list of the BGP.
	AsPaths []string `pulumi:"asPaths"`
	// A network address segment of IDC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId *string `pulumi:"dcgId"`
}

type GatewayCcnRouteState struct {
	// As path list of the BGP.
	AsPaths pulumi.StringArrayInput
	// A network address segment of IDC.
	CidrBlock pulumi.StringPtrInput
	// ID of the DCG.
	DcgId pulumi.StringPtrInput
}

func (GatewayCcnRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCcnRouteState)(nil)).Elem()
}

type gatewayCcnRouteArgs struct {
	// A network address segment of IDC.
	CidrBlock string `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId string `pulumi:"dcgId"`
}

// The set of arguments for constructing a GatewayCcnRoute resource.
type GatewayCcnRouteArgs struct {
	// A network address segment of IDC.
	CidrBlock pulumi.StringInput
	// ID of the DCG.
	DcgId pulumi.StringInput
}

func (GatewayCcnRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCcnRouteArgs)(nil)).Elem()
}

type GatewayCcnRouteInput interface {
	pulumi.Input

	ToGatewayCcnRouteOutput() GatewayCcnRouteOutput
	ToGatewayCcnRouteOutputWithContext(ctx context.Context) GatewayCcnRouteOutput
}

func (*GatewayCcnRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCcnRoute)(nil)).Elem()
}

func (i *GatewayCcnRoute) ToGatewayCcnRouteOutput() GatewayCcnRouteOutput {
	return i.ToGatewayCcnRouteOutputWithContext(context.Background())
}

func (i *GatewayCcnRoute) ToGatewayCcnRouteOutputWithContext(ctx context.Context) GatewayCcnRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteOutput)
}

// GatewayCcnRouteArrayInput is an input type that accepts GatewayCcnRouteArray and GatewayCcnRouteArrayOutput values.
// You can construct a concrete instance of `GatewayCcnRouteArrayInput` via:
//
//	GatewayCcnRouteArray{ GatewayCcnRouteArgs{...} }
type GatewayCcnRouteArrayInput interface {
	pulumi.Input

	ToGatewayCcnRouteArrayOutput() GatewayCcnRouteArrayOutput
	ToGatewayCcnRouteArrayOutputWithContext(context.Context) GatewayCcnRouteArrayOutput
}

type GatewayCcnRouteArray []GatewayCcnRouteInput

func (GatewayCcnRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCcnRoute)(nil)).Elem()
}

func (i GatewayCcnRouteArray) ToGatewayCcnRouteArrayOutput() GatewayCcnRouteArrayOutput {
	return i.ToGatewayCcnRouteArrayOutputWithContext(context.Background())
}

func (i GatewayCcnRouteArray) ToGatewayCcnRouteArrayOutputWithContext(ctx context.Context) GatewayCcnRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteArrayOutput)
}

// GatewayCcnRouteMapInput is an input type that accepts GatewayCcnRouteMap and GatewayCcnRouteMapOutput values.
// You can construct a concrete instance of `GatewayCcnRouteMapInput` via:
//
//	GatewayCcnRouteMap{ "key": GatewayCcnRouteArgs{...} }
type GatewayCcnRouteMapInput interface {
	pulumi.Input

	ToGatewayCcnRouteMapOutput() GatewayCcnRouteMapOutput
	ToGatewayCcnRouteMapOutputWithContext(context.Context) GatewayCcnRouteMapOutput
}

type GatewayCcnRouteMap map[string]GatewayCcnRouteInput

func (GatewayCcnRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCcnRoute)(nil)).Elem()
}

func (i GatewayCcnRouteMap) ToGatewayCcnRouteMapOutput() GatewayCcnRouteMapOutput {
	return i.ToGatewayCcnRouteMapOutputWithContext(context.Background())
}

func (i GatewayCcnRouteMap) ToGatewayCcnRouteMapOutputWithContext(ctx context.Context) GatewayCcnRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteMapOutput)
}

type GatewayCcnRouteOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCcnRoute)(nil)).Elem()
}

func (o GatewayCcnRouteOutput) ToGatewayCcnRouteOutput() GatewayCcnRouteOutput {
	return o
}

func (o GatewayCcnRouteOutput) ToGatewayCcnRouteOutputWithContext(ctx context.Context) GatewayCcnRouteOutput {
	return o
}

// As path list of the BGP.
func (o GatewayCcnRouteOutput) AsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCcnRoute) pulumi.StringArrayOutput { return v.AsPaths }).(pulumi.StringArrayOutput)
}

// A network address segment of IDC.
func (o GatewayCcnRouteOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCcnRoute) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// ID of the DCG.
func (o GatewayCcnRouteOutput) DcgId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCcnRoute) pulumi.StringOutput { return v.DcgId }).(pulumi.StringOutput)
}

type GatewayCcnRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCcnRoute)(nil)).Elem()
}

func (o GatewayCcnRouteArrayOutput) ToGatewayCcnRouteArrayOutput() GatewayCcnRouteArrayOutput {
	return o
}

func (o GatewayCcnRouteArrayOutput) ToGatewayCcnRouteArrayOutputWithContext(ctx context.Context) GatewayCcnRouteArrayOutput {
	return o
}

func (o GatewayCcnRouteArrayOutput) Index(i pulumi.IntInput) GatewayCcnRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayCcnRoute {
		return vs[0].([]*GatewayCcnRoute)[vs[1].(int)]
	}).(GatewayCcnRouteOutput)
}

type GatewayCcnRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCcnRoute)(nil)).Elem()
}

func (o GatewayCcnRouteMapOutput) ToGatewayCcnRouteMapOutput() GatewayCcnRouteMapOutput {
	return o
}

func (o GatewayCcnRouteMapOutput) ToGatewayCcnRouteMapOutputWithContext(ctx context.Context) GatewayCcnRouteMapOutput {
	return o
}

func (o GatewayCcnRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayCcnRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayCcnRoute {
		return vs[0].(map[string]*GatewayCcnRoute)[vs[1].(string)]
	}).(GatewayCcnRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteInput)(nil)).Elem(), &GatewayCcnRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteArrayInput)(nil)).Elem(), GatewayCcnRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteMapInput)(nil)).Elem(), GatewayCcnRouteMap{})
	pulumi.RegisterOutputType(GatewayCcnRouteOutput{})
	pulumi.RegisterOutputType(GatewayCcnRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayCcnRouteMapOutput{})
}
