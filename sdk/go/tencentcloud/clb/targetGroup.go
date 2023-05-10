// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a CLB target group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Clb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Clb.NewTargetGroup(ctx, "test", &Clb.TargetGroupArgs{
// 			Port:            pulumi.Int(33),
// 			TargetGroupName: pulumi.String("test"),
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
// CLB target group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Clb/targetGroup:TargetGroup test lbtg-3k3io0i0
// ```
type TargetGroup struct {
	pulumi.CustomResourceState

	// The default port of target group, add server after can use it.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
	//
	// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
	TargetGroupInstances TargetGroupTargetGroupInstanceArrayOutput `pulumi:"targetGroupInstances"`
	// Target group name.
	TargetGroupName pulumi.StringPtrOutput `pulumi:"targetGroupName"`
	// VPC ID, default is based on the network.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewTargetGroup(ctx *pulumi.Context,
	name string, args *TargetGroupArgs, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	if args == nil {
		args = &TargetGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource TargetGroup
	err := ctx.RegisterResource("tencentcloud:Clb/targetGroup:TargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroup gets an existing TargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupState, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	var resource TargetGroup
	err := ctx.ReadResource("tencentcloud:Clb/targetGroup:TargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroup resources.
type targetGroupState struct {
	// The default port of target group, add server after can use it.
	Port *int `pulumi:"port"`
	// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
	//
	// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
	TargetGroupInstances []TargetGroupTargetGroupInstance `pulumi:"targetGroupInstances"`
	// Target group name.
	TargetGroupName *string `pulumi:"targetGroupName"`
	// VPC ID, default is based on the network.
	VpcId *string `pulumi:"vpcId"`
}

type TargetGroupState struct {
	// The default port of target group, add server after can use it.
	Port pulumi.IntPtrInput
	// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
	//
	// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
	TargetGroupInstances TargetGroupTargetGroupInstanceArrayInput
	// Target group name.
	TargetGroupName pulumi.StringPtrInput
	// VPC ID, default is based on the network.
	VpcId pulumi.StringPtrInput
}

func (TargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupState)(nil)).Elem()
}

type targetGroupArgs struct {
	// The default port of target group, add server after can use it.
	Port *int `pulumi:"port"`
	// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
	//
	// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
	TargetGroupInstances []TargetGroupTargetGroupInstance `pulumi:"targetGroupInstances"`
	// Target group name.
	TargetGroupName *string `pulumi:"targetGroupName"`
	// VPC ID, default is based on the network.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a TargetGroup resource.
type TargetGroupArgs struct {
	// The default port of target group, add server after can use it.
	Port pulumi.IntPtrInput
	// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
	//
	// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
	TargetGroupInstances TargetGroupTargetGroupInstanceArrayInput
	// Target group name.
	TargetGroupName pulumi.StringPtrInput
	// VPC ID, default is based on the network.
	VpcId pulumi.StringPtrInput
}

func (TargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupArgs)(nil)).Elem()
}

type TargetGroupInput interface {
	pulumi.Input

	ToTargetGroupOutput() TargetGroupOutput
	ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput
}

func (*TargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil)).Elem()
}

func (i *TargetGroup) ToTargetGroupOutput() TargetGroupOutput {
	return i.ToTargetGroupOutputWithContext(context.Background())
}

func (i *TargetGroup) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupOutput)
}

// TargetGroupArrayInput is an input type that accepts TargetGroupArray and TargetGroupArrayOutput values.
// You can construct a concrete instance of `TargetGroupArrayInput` via:
//
//          TargetGroupArray{ TargetGroupArgs{...} }
type TargetGroupArrayInput interface {
	pulumi.Input

	ToTargetGroupArrayOutput() TargetGroupArrayOutput
	ToTargetGroupArrayOutputWithContext(context.Context) TargetGroupArrayOutput
}

type TargetGroupArray []TargetGroupInput

func (TargetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetGroup)(nil)).Elem()
}

func (i TargetGroupArray) ToTargetGroupArrayOutput() TargetGroupArrayOutput {
	return i.ToTargetGroupArrayOutputWithContext(context.Background())
}

func (i TargetGroupArray) ToTargetGroupArrayOutputWithContext(ctx context.Context) TargetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupArrayOutput)
}

// TargetGroupMapInput is an input type that accepts TargetGroupMap and TargetGroupMapOutput values.
// You can construct a concrete instance of `TargetGroupMapInput` via:
//
//          TargetGroupMap{ "key": TargetGroupArgs{...} }
type TargetGroupMapInput interface {
	pulumi.Input

	ToTargetGroupMapOutput() TargetGroupMapOutput
	ToTargetGroupMapOutputWithContext(context.Context) TargetGroupMapOutput
}

type TargetGroupMap map[string]TargetGroupInput

func (TargetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetGroup)(nil)).Elem()
}

func (i TargetGroupMap) ToTargetGroupMapOutput() TargetGroupMapOutput {
	return i.ToTargetGroupMapOutputWithContext(context.Background())
}

func (i TargetGroupMap) ToTargetGroupMapOutputWithContext(ctx context.Context) TargetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupMapOutput)
}

type TargetGroupOutput struct{ *pulumi.OutputState }

func (TargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil)).Elem()
}

func (o TargetGroupOutput) ToTargetGroupOutput() TargetGroupOutput {
	return o
}

func (o TargetGroupOutput) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return o
}

// The default port of target group, add server after can use it.
func (o TargetGroupOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// It has been deprecated from version 1.77.3. please use `Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
//
// Deprecated: It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
func (o TargetGroupOutput) TargetGroupInstances() TargetGroupTargetGroupInstanceArrayOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupTargetGroupInstanceArrayOutput { return v.TargetGroupInstances }).(TargetGroupTargetGroupInstanceArrayOutput)
}

// Target group name.
func (o TargetGroupOutput) TargetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringPtrOutput { return v.TargetGroupName }).(pulumi.StringPtrOutput)
}

// VPC ID, default is based on the network.
func (o TargetGroupOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type TargetGroupArrayOutput struct{ *pulumi.OutputState }

func (TargetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetGroup)(nil)).Elem()
}

func (o TargetGroupArrayOutput) ToTargetGroupArrayOutput() TargetGroupArrayOutput {
	return o
}

func (o TargetGroupArrayOutput) ToTargetGroupArrayOutputWithContext(ctx context.Context) TargetGroupArrayOutput {
	return o
}

func (o TargetGroupArrayOutput) Index(i pulumi.IntInput) TargetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TargetGroup {
		return vs[0].([]*TargetGroup)[vs[1].(int)]
	}).(TargetGroupOutput)
}

type TargetGroupMapOutput struct{ *pulumi.OutputState }

func (TargetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetGroup)(nil)).Elem()
}

func (o TargetGroupMapOutput) ToTargetGroupMapOutput() TargetGroupMapOutput {
	return o
}

func (o TargetGroupMapOutput) ToTargetGroupMapOutputWithContext(ctx context.Context) TargetGroupMapOutput {
	return o
}

func (o TargetGroupMapOutput) MapIndex(k pulumi.StringInput) TargetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TargetGroup {
		return vs[0].(map[string]*TargetGroup)[vs[1].(string)]
	}).(TargetGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetGroupInput)(nil)).Elem(), &TargetGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetGroupArrayInput)(nil)).Elem(), TargetGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetGroupMapInput)(nil)).Elem(), TargetGroupMap{})
	pulumi.RegisterOutputType(TargetGroupOutput{})
	pulumi.RegisterOutputType(TargetGroupArrayOutput{})
	pulumi.RegisterOutputType(TargetGroupMapOutput{})
}
