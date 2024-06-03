// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tdmq rabbitmqVipInstance
//
// ## Import
//
// tdmq rabbitmq_vip_instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance example amqp-mok52gmn
// ```
type RabbitmqVipInstance struct {
	pulumi.CustomResourceState

	// Automatic renewal, the default is true.
	AutoRenewFlag pulumi.BoolPtrOutput `pulumi:"autoRenewFlag"`
	// cluster name.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Mirrored queue, the default is false.
	EnableCreateDefaultHaMirrorQueue pulumi.BoolPtrOutput `pulumi:"enableCreateDefaultHaMirrorQueue"`
	// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
	NodeNum pulumi.IntPtrOutput `pulumi:"nodeNum"`
	// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
	NodeSpec pulumi.StringPtrOutput `pulumi:"nodeSpec"`
	// Single node storage specification, the default is 200G.
	StorageSize pulumi.IntPtrOutput `pulumi:"storageSize"`
	// Private network SubnetId.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Purchase duration, the default is 1 (month).
	TimeSpan pulumi.IntPtrOutput `pulumi:"timeSpan"`
	// Private network VpcId.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// availability zone.
	ZoneIds pulumi.IntArrayOutput `pulumi:"zoneIds"`
}

// NewRabbitmqVipInstance registers a new resource with the given unique name, arguments, and options.
func NewRabbitmqVipInstance(ctx *pulumi.Context,
	name string, args *RabbitmqVipInstanceArgs, opts ...pulumi.ResourceOption) (*RabbitmqVipInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.ZoneIds == nil {
		return nil, errors.New("invalid value for required argument 'ZoneIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RabbitmqVipInstance
	err := ctx.RegisterResource("tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRabbitmqVipInstance gets an existing RabbitmqVipInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRabbitmqVipInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RabbitmqVipInstanceState, opts ...pulumi.ResourceOption) (*RabbitmqVipInstance, error) {
	var resource RabbitmqVipInstance
	err := ctx.ReadResource("tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RabbitmqVipInstance resources.
type rabbitmqVipInstanceState struct {
	// Automatic renewal, the default is true.
	AutoRenewFlag *bool `pulumi:"autoRenewFlag"`
	// cluster name.
	ClusterName *string `pulumi:"clusterName"`
	// Mirrored queue, the default is false.
	EnableCreateDefaultHaMirrorQueue *bool `pulumi:"enableCreateDefaultHaMirrorQueue"`
	// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
	NodeNum *int `pulumi:"nodeNum"`
	// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
	NodeSpec *string `pulumi:"nodeSpec"`
	// Single node storage specification, the default is 200G.
	StorageSize *int `pulumi:"storageSize"`
	// Private network SubnetId.
	SubnetId *string `pulumi:"subnetId"`
	// Purchase duration, the default is 1 (month).
	TimeSpan *int `pulumi:"timeSpan"`
	// Private network VpcId.
	VpcId *string `pulumi:"vpcId"`
	// availability zone.
	ZoneIds []int `pulumi:"zoneIds"`
}

type RabbitmqVipInstanceState struct {
	// Automatic renewal, the default is true.
	AutoRenewFlag pulumi.BoolPtrInput
	// cluster name.
	ClusterName pulumi.StringPtrInput
	// Mirrored queue, the default is false.
	EnableCreateDefaultHaMirrorQueue pulumi.BoolPtrInput
	// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
	NodeNum pulumi.IntPtrInput
	// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
	NodeSpec pulumi.StringPtrInput
	// Single node storage specification, the default is 200G.
	StorageSize pulumi.IntPtrInput
	// Private network SubnetId.
	SubnetId pulumi.StringPtrInput
	// Purchase duration, the default is 1 (month).
	TimeSpan pulumi.IntPtrInput
	// Private network VpcId.
	VpcId pulumi.StringPtrInput
	// availability zone.
	ZoneIds pulumi.IntArrayInput
}

func (RabbitmqVipInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitmqVipInstanceState)(nil)).Elem()
}

type rabbitmqVipInstanceArgs struct {
	// Automatic renewal, the default is true.
	AutoRenewFlag *bool `pulumi:"autoRenewFlag"`
	// cluster name.
	ClusterName string `pulumi:"clusterName"`
	// Mirrored queue, the default is false.
	EnableCreateDefaultHaMirrorQueue *bool `pulumi:"enableCreateDefaultHaMirrorQueue"`
	// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
	NodeNum *int `pulumi:"nodeNum"`
	// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
	NodeSpec *string `pulumi:"nodeSpec"`
	// Single node storage specification, the default is 200G.
	StorageSize *int `pulumi:"storageSize"`
	// Private network SubnetId.
	SubnetId string `pulumi:"subnetId"`
	// Purchase duration, the default is 1 (month).
	TimeSpan *int `pulumi:"timeSpan"`
	// Private network VpcId.
	VpcId string `pulumi:"vpcId"`
	// availability zone.
	ZoneIds []int `pulumi:"zoneIds"`
}

// The set of arguments for constructing a RabbitmqVipInstance resource.
type RabbitmqVipInstanceArgs struct {
	// Automatic renewal, the default is true.
	AutoRenewFlag pulumi.BoolPtrInput
	// cluster name.
	ClusterName pulumi.StringInput
	// Mirrored queue, the default is false.
	EnableCreateDefaultHaMirrorQueue pulumi.BoolPtrInput
	// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
	NodeNum pulumi.IntPtrInput
	// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
	NodeSpec pulumi.StringPtrInput
	// Single node storage specification, the default is 200G.
	StorageSize pulumi.IntPtrInput
	// Private network SubnetId.
	SubnetId pulumi.StringInput
	// Purchase duration, the default is 1 (month).
	TimeSpan pulumi.IntPtrInput
	// Private network VpcId.
	VpcId pulumi.StringInput
	// availability zone.
	ZoneIds pulumi.IntArrayInput
}

func (RabbitmqVipInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitmqVipInstanceArgs)(nil)).Elem()
}

type RabbitmqVipInstanceInput interface {
	pulumi.Input

	ToRabbitmqVipInstanceOutput() RabbitmqVipInstanceOutput
	ToRabbitmqVipInstanceOutputWithContext(ctx context.Context) RabbitmqVipInstanceOutput
}

func (*RabbitmqVipInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitmqVipInstance)(nil)).Elem()
}

func (i *RabbitmqVipInstance) ToRabbitmqVipInstanceOutput() RabbitmqVipInstanceOutput {
	return i.ToRabbitmqVipInstanceOutputWithContext(context.Background())
}

func (i *RabbitmqVipInstance) ToRabbitmqVipInstanceOutputWithContext(ctx context.Context) RabbitmqVipInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVipInstanceOutput)
}

// RabbitmqVipInstanceArrayInput is an input type that accepts RabbitmqVipInstanceArray and RabbitmqVipInstanceArrayOutput values.
// You can construct a concrete instance of `RabbitmqVipInstanceArrayInput` via:
//
//	RabbitmqVipInstanceArray{ RabbitmqVipInstanceArgs{...} }
type RabbitmqVipInstanceArrayInput interface {
	pulumi.Input

	ToRabbitmqVipInstanceArrayOutput() RabbitmqVipInstanceArrayOutput
	ToRabbitmqVipInstanceArrayOutputWithContext(context.Context) RabbitmqVipInstanceArrayOutput
}

type RabbitmqVipInstanceArray []RabbitmqVipInstanceInput

func (RabbitmqVipInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitmqVipInstance)(nil)).Elem()
}

func (i RabbitmqVipInstanceArray) ToRabbitmqVipInstanceArrayOutput() RabbitmqVipInstanceArrayOutput {
	return i.ToRabbitmqVipInstanceArrayOutputWithContext(context.Background())
}

func (i RabbitmqVipInstanceArray) ToRabbitmqVipInstanceArrayOutputWithContext(ctx context.Context) RabbitmqVipInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVipInstanceArrayOutput)
}

// RabbitmqVipInstanceMapInput is an input type that accepts RabbitmqVipInstanceMap and RabbitmqVipInstanceMapOutput values.
// You can construct a concrete instance of `RabbitmqVipInstanceMapInput` via:
//
//	RabbitmqVipInstanceMap{ "key": RabbitmqVipInstanceArgs{...} }
type RabbitmqVipInstanceMapInput interface {
	pulumi.Input

	ToRabbitmqVipInstanceMapOutput() RabbitmqVipInstanceMapOutput
	ToRabbitmqVipInstanceMapOutputWithContext(context.Context) RabbitmqVipInstanceMapOutput
}

type RabbitmqVipInstanceMap map[string]RabbitmqVipInstanceInput

func (RabbitmqVipInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitmqVipInstance)(nil)).Elem()
}

func (i RabbitmqVipInstanceMap) ToRabbitmqVipInstanceMapOutput() RabbitmqVipInstanceMapOutput {
	return i.ToRabbitmqVipInstanceMapOutputWithContext(context.Background())
}

func (i RabbitmqVipInstanceMap) ToRabbitmqVipInstanceMapOutputWithContext(ctx context.Context) RabbitmqVipInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVipInstanceMapOutput)
}

type RabbitmqVipInstanceOutput struct{ *pulumi.OutputState }

func (RabbitmqVipInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitmqVipInstance)(nil)).Elem()
}

func (o RabbitmqVipInstanceOutput) ToRabbitmqVipInstanceOutput() RabbitmqVipInstanceOutput {
	return o
}

func (o RabbitmqVipInstanceOutput) ToRabbitmqVipInstanceOutputWithContext(ctx context.Context) RabbitmqVipInstanceOutput {
	return o
}

// Automatic renewal, the default is true.
func (o RabbitmqVipInstanceOutput) AutoRenewFlag() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.BoolPtrOutput { return v.AutoRenewFlag }).(pulumi.BoolPtrOutput)
}

// cluster name.
func (o RabbitmqVipInstanceOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Mirrored queue, the default is false.
func (o RabbitmqVipInstanceOutput) EnableCreateDefaultHaMirrorQueue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.BoolPtrOutput { return v.EnableCreateDefaultHaMirrorQueue }).(pulumi.BoolPtrOutput)
}

// The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
func (o RabbitmqVipInstanceOutput) NodeNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.IntPtrOutput { return v.NodeNum }).(pulumi.IntPtrOutput)
}

// Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
func (o RabbitmqVipInstanceOutput) NodeSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.StringPtrOutput { return v.NodeSpec }).(pulumi.StringPtrOutput)
}

// Single node storage specification, the default is 200G.
func (o RabbitmqVipInstanceOutput) StorageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.IntPtrOutput { return v.StorageSize }).(pulumi.IntPtrOutput)
}

// Private network SubnetId.
func (o RabbitmqVipInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Purchase duration, the default is 1 (month).
func (o RabbitmqVipInstanceOutput) TimeSpan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.IntPtrOutput { return v.TimeSpan }).(pulumi.IntPtrOutput)
}

// Private network VpcId.
func (o RabbitmqVipInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// availability zone.
func (o RabbitmqVipInstanceOutput) ZoneIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RabbitmqVipInstance) pulumi.IntArrayOutput { return v.ZoneIds }).(pulumi.IntArrayOutput)
}

type RabbitmqVipInstanceArrayOutput struct{ *pulumi.OutputState }

func (RabbitmqVipInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitmqVipInstance)(nil)).Elem()
}

func (o RabbitmqVipInstanceArrayOutput) ToRabbitmqVipInstanceArrayOutput() RabbitmqVipInstanceArrayOutput {
	return o
}

func (o RabbitmqVipInstanceArrayOutput) ToRabbitmqVipInstanceArrayOutputWithContext(ctx context.Context) RabbitmqVipInstanceArrayOutput {
	return o
}

func (o RabbitmqVipInstanceArrayOutput) Index(i pulumi.IntInput) RabbitmqVipInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RabbitmqVipInstance {
		return vs[0].([]*RabbitmqVipInstance)[vs[1].(int)]
	}).(RabbitmqVipInstanceOutput)
}

type RabbitmqVipInstanceMapOutput struct{ *pulumi.OutputState }

func (RabbitmqVipInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitmqVipInstance)(nil)).Elem()
}

func (o RabbitmqVipInstanceMapOutput) ToRabbitmqVipInstanceMapOutput() RabbitmqVipInstanceMapOutput {
	return o
}

func (o RabbitmqVipInstanceMapOutput) ToRabbitmqVipInstanceMapOutputWithContext(ctx context.Context) RabbitmqVipInstanceMapOutput {
	return o
}

func (o RabbitmqVipInstanceMapOutput) MapIndex(k pulumi.StringInput) RabbitmqVipInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RabbitmqVipInstance {
		return vs[0].(map[string]*RabbitmqVipInstance)[vs[1].(string)]
	}).(RabbitmqVipInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVipInstanceInput)(nil)).Elem(), &RabbitmqVipInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVipInstanceArrayInput)(nil)).Elem(), RabbitmqVipInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVipInstanceMapInput)(nil)).Elem(), RabbitmqVipInstanceMap{})
	pulumi.RegisterOutputType(RabbitmqVipInstanceOutput{})
	pulumi.RegisterOutputType(RabbitmqVipInstanceArrayOutput{})
	pulumi.RegisterOutputType(RabbitmqVipInstanceMapOutput{})
}
