// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cynosdb wan
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.NewWan(ctx, "wan", &Cynosdb.WanArgs{
// 			ClusterId:     pulumi.String("cynosdbmysql-bws8h88b"),
// 			InstanceGrpId: pulumi.String("cynosdbmysql-grp-lxav0p9z"),
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
// cynosdb wan can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cynosdb/wan:Wan wan cynosdbmysql-bws8h88b#cynosdbmysql-grp-lxav0p9z
// ```
type Wan struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Instance Group ID.
	InstanceGrpId pulumi.StringOutput `pulumi:"instanceGrpId"`
	// Domain name.
	WanDomain pulumi.StringOutput `pulumi:"wanDomain"`
	// Network ip.
	WanIp pulumi.StringOutput `pulumi:"wanIp"`
	// Internet port.
	WanPort pulumi.IntOutput `pulumi:"wanPort"`
	// Internet status.
	WanStatus pulumi.StringOutput `pulumi:"wanStatus"`
}

// NewWan registers a new resource with the given unique name, arguments, and options.
func NewWan(ctx *pulumi.Context,
	name string, args *WanArgs, opts ...pulumi.ResourceOption) (*Wan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceGrpId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceGrpId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Wan
	err := ctx.RegisterResource("tencentcloud:Cynosdb/wan:Wan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWan gets an existing Wan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanState, opts ...pulumi.ResourceOption) (*Wan, error) {
	var resource Wan
	err := ctx.ReadResource("tencentcloud:Cynosdb/wan:Wan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wan resources.
type wanState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Instance Group ID.
	InstanceGrpId *string `pulumi:"instanceGrpId"`
	// Domain name.
	WanDomain *string `pulumi:"wanDomain"`
	// Network ip.
	WanIp *string `pulumi:"wanIp"`
	// Internet port.
	WanPort *int `pulumi:"wanPort"`
	// Internet status.
	WanStatus *string `pulumi:"wanStatus"`
}

type WanState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Instance Group ID.
	InstanceGrpId pulumi.StringPtrInput
	// Domain name.
	WanDomain pulumi.StringPtrInput
	// Network ip.
	WanIp pulumi.StringPtrInput
	// Internet port.
	WanPort pulumi.IntPtrInput
	// Internet status.
	WanStatus pulumi.StringPtrInput
}

func (WanState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanState)(nil)).Elem()
}

type wanArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Instance Group ID.
	InstanceGrpId string `pulumi:"instanceGrpId"`
}

// The set of arguments for constructing a Wan resource.
type WanArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Instance Group ID.
	InstanceGrpId pulumi.StringInput
}

func (WanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanArgs)(nil)).Elem()
}

type WanInput interface {
	pulumi.Input

	ToWanOutput() WanOutput
	ToWanOutputWithContext(ctx context.Context) WanOutput
}

func (*Wan) ElementType() reflect.Type {
	return reflect.TypeOf((**Wan)(nil)).Elem()
}

func (i *Wan) ToWanOutput() WanOutput {
	return i.ToWanOutputWithContext(context.Background())
}

func (i *Wan) ToWanOutputWithContext(ctx context.Context) WanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanOutput)
}

// WanArrayInput is an input type that accepts WanArray and WanArrayOutput values.
// You can construct a concrete instance of `WanArrayInput` via:
//
//          WanArray{ WanArgs{...} }
type WanArrayInput interface {
	pulumi.Input

	ToWanArrayOutput() WanArrayOutput
	ToWanArrayOutputWithContext(context.Context) WanArrayOutput
}

type WanArray []WanInput

func (WanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wan)(nil)).Elem()
}

func (i WanArray) ToWanArrayOutput() WanArrayOutput {
	return i.ToWanArrayOutputWithContext(context.Background())
}

func (i WanArray) ToWanArrayOutputWithContext(ctx context.Context) WanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanArrayOutput)
}

// WanMapInput is an input type that accepts WanMap and WanMapOutput values.
// You can construct a concrete instance of `WanMapInput` via:
//
//          WanMap{ "key": WanArgs{...} }
type WanMapInput interface {
	pulumi.Input

	ToWanMapOutput() WanMapOutput
	ToWanMapOutputWithContext(context.Context) WanMapOutput
}

type WanMap map[string]WanInput

func (WanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wan)(nil)).Elem()
}

func (i WanMap) ToWanMapOutput() WanMapOutput {
	return i.ToWanMapOutputWithContext(context.Background())
}

func (i WanMap) ToWanMapOutputWithContext(ctx context.Context) WanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanMapOutput)
}

type WanOutput struct{ *pulumi.OutputState }

func (WanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wan)(nil)).Elem()
}

func (o WanOutput) ToWanOutput() WanOutput {
	return o
}

func (o WanOutput) ToWanOutputWithContext(ctx context.Context) WanOutput {
	return o
}

// Cluster ID.
func (o WanOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Wan) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Instance Group ID.
func (o WanOutput) InstanceGrpId() pulumi.StringOutput {
	return o.ApplyT(func(v *Wan) pulumi.StringOutput { return v.InstanceGrpId }).(pulumi.StringOutput)
}

// Domain name.
func (o WanOutput) WanDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Wan) pulumi.StringOutput { return v.WanDomain }).(pulumi.StringOutput)
}

// Network ip.
func (o WanOutput) WanIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Wan) pulumi.StringOutput { return v.WanIp }).(pulumi.StringOutput)
}

// Internet port.
func (o WanOutput) WanPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Wan) pulumi.IntOutput { return v.WanPort }).(pulumi.IntOutput)
}

// Internet status.
func (o WanOutput) WanStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Wan) pulumi.StringOutput { return v.WanStatus }).(pulumi.StringOutput)
}

type WanArrayOutput struct{ *pulumi.OutputState }

func (WanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wan)(nil)).Elem()
}

func (o WanArrayOutput) ToWanArrayOutput() WanArrayOutput {
	return o
}

func (o WanArrayOutput) ToWanArrayOutputWithContext(ctx context.Context) WanArrayOutput {
	return o
}

func (o WanArrayOutput) Index(i pulumi.IntInput) WanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wan {
		return vs[0].([]*Wan)[vs[1].(int)]
	}).(WanOutput)
}

type WanMapOutput struct{ *pulumi.OutputState }

func (WanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wan)(nil)).Elem()
}

func (o WanMapOutput) ToWanMapOutput() WanMapOutput {
	return o
}

func (o WanMapOutput) ToWanMapOutputWithContext(ctx context.Context) WanMapOutput {
	return o
}

func (o WanMapOutput) MapIndex(k pulumi.StringInput) WanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wan {
		return vs[0].(map[string]*Wan)[vs[1].(string)]
	}).(WanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanInput)(nil)).Elem(), &Wan{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanArrayInput)(nil)).Elem(), WanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanMapInput)(nil)).Elem(), WanMap{})
	pulumi.RegisterOutputType(WanOutput{})
	pulumi.RegisterOutputType(WanArrayOutput{})
	pulumi.RegisterOutputType(WanMapOutput{})
}
