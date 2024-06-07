// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tsf operateGroup
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.NewOperateGroup(ctx, "operateGroup", &Tsf.OperateGroupArgs{
//				GroupId: pulumi.String("group-ynd95rea"),
//				Operate: pulumi.String("start"),
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
type OperateGroup struct {
	pulumi.CustomResourceState

	// group id.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Operation, `start`- start the group, `stop`- stop the group.
	Operate pulumi.StringOutput `pulumi:"operate"`
}

// NewOperateGroup registers a new resource with the given unique name, arguments, and options.
func NewOperateGroup(ctx *pulumi.Context,
	name string, args *OperateGroupArgs, opts ...pulumi.ResourceOption) (*OperateGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Operate == nil {
		return nil, errors.New("invalid value for required argument 'Operate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OperateGroup
	err := ctx.RegisterResource("tencentcloud:Tsf/operateGroup:OperateGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOperateGroup gets an existing OperateGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOperateGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperateGroupState, opts ...pulumi.ResourceOption) (*OperateGroup, error) {
	var resource OperateGroup
	err := ctx.ReadResource("tencentcloud:Tsf/operateGroup:OperateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OperateGroup resources.
type operateGroupState struct {
	// group id.
	GroupId *string `pulumi:"groupId"`
	// Operation, `start`- start the group, `stop`- stop the group.
	Operate *string `pulumi:"operate"`
}

type OperateGroupState struct {
	// group id.
	GroupId pulumi.StringPtrInput
	// Operation, `start`- start the group, `stop`- stop the group.
	Operate pulumi.StringPtrInput
}

func (OperateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*operateGroupState)(nil)).Elem()
}

type operateGroupArgs struct {
	// group id.
	GroupId string `pulumi:"groupId"`
	// Operation, `start`- start the group, `stop`- stop the group.
	Operate string `pulumi:"operate"`
}

// The set of arguments for constructing a OperateGroup resource.
type OperateGroupArgs struct {
	// group id.
	GroupId pulumi.StringInput
	// Operation, `start`- start the group, `stop`- stop the group.
	Operate pulumi.StringInput
}

func (OperateGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*operateGroupArgs)(nil)).Elem()
}

type OperateGroupInput interface {
	pulumi.Input

	ToOperateGroupOutput() OperateGroupOutput
	ToOperateGroupOutputWithContext(ctx context.Context) OperateGroupOutput
}

func (*OperateGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**OperateGroup)(nil)).Elem()
}

func (i *OperateGroup) ToOperateGroupOutput() OperateGroupOutput {
	return i.ToOperateGroupOutputWithContext(context.Background())
}

func (i *OperateGroup) ToOperateGroupOutputWithContext(ctx context.Context) OperateGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperateGroupOutput)
}

// OperateGroupArrayInput is an input type that accepts OperateGroupArray and OperateGroupArrayOutput values.
// You can construct a concrete instance of `OperateGroupArrayInput` via:
//
//	OperateGroupArray{ OperateGroupArgs{...} }
type OperateGroupArrayInput interface {
	pulumi.Input

	ToOperateGroupArrayOutput() OperateGroupArrayOutput
	ToOperateGroupArrayOutputWithContext(context.Context) OperateGroupArrayOutput
}

type OperateGroupArray []OperateGroupInput

func (OperateGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OperateGroup)(nil)).Elem()
}

func (i OperateGroupArray) ToOperateGroupArrayOutput() OperateGroupArrayOutput {
	return i.ToOperateGroupArrayOutputWithContext(context.Background())
}

func (i OperateGroupArray) ToOperateGroupArrayOutputWithContext(ctx context.Context) OperateGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperateGroupArrayOutput)
}

// OperateGroupMapInput is an input type that accepts OperateGroupMap and OperateGroupMapOutput values.
// You can construct a concrete instance of `OperateGroupMapInput` via:
//
//	OperateGroupMap{ "key": OperateGroupArgs{...} }
type OperateGroupMapInput interface {
	pulumi.Input

	ToOperateGroupMapOutput() OperateGroupMapOutput
	ToOperateGroupMapOutputWithContext(context.Context) OperateGroupMapOutput
}

type OperateGroupMap map[string]OperateGroupInput

func (OperateGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OperateGroup)(nil)).Elem()
}

func (i OperateGroupMap) ToOperateGroupMapOutput() OperateGroupMapOutput {
	return i.ToOperateGroupMapOutputWithContext(context.Background())
}

func (i OperateGroupMap) ToOperateGroupMapOutputWithContext(ctx context.Context) OperateGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperateGroupMapOutput)
}

type OperateGroupOutput struct{ *pulumi.OutputState }

func (OperateGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperateGroup)(nil)).Elem()
}

func (o OperateGroupOutput) ToOperateGroupOutput() OperateGroupOutput {
	return o
}

func (o OperateGroupOutput) ToOperateGroupOutputWithContext(ctx context.Context) OperateGroupOutput {
	return o
}

// group id.
func (o OperateGroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *OperateGroup) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Operation, `start`- start the group, `stop`- stop the group.
func (o OperateGroupOutput) Operate() pulumi.StringOutput {
	return o.ApplyT(func(v *OperateGroup) pulumi.StringOutput { return v.Operate }).(pulumi.StringOutput)
}

type OperateGroupArrayOutput struct{ *pulumi.OutputState }

func (OperateGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OperateGroup)(nil)).Elem()
}

func (o OperateGroupArrayOutput) ToOperateGroupArrayOutput() OperateGroupArrayOutput {
	return o
}

func (o OperateGroupArrayOutput) ToOperateGroupArrayOutputWithContext(ctx context.Context) OperateGroupArrayOutput {
	return o
}

func (o OperateGroupArrayOutput) Index(i pulumi.IntInput) OperateGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OperateGroup {
		return vs[0].([]*OperateGroup)[vs[1].(int)]
	}).(OperateGroupOutput)
}

type OperateGroupMapOutput struct{ *pulumi.OutputState }

func (OperateGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OperateGroup)(nil)).Elem()
}

func (o OperateGroupMapOutput) ToOperateGroupMapOutput() OperateGroupMapOutput {
	return o
}

func (o OperateGroupMapOutput) ToOperateGroupMapOutputWithContext(ctx context.Context) OperateGroupMapOutput {
	return o
}

func (o OperateGroupMapOutput) MapIndex(k pulumi.StringInput) OperateGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OperateGroup {
		return vs[0].(map[string]*OperateGroup)[vs[1].(string)]
	}).(OperateGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OperateGroupInput)(nil)).Elem(), &OperateGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperateGroupArrayInput)(nil)).Elem(), OperateGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperateGroupMapInput)(nil)).Elem(), OperateGroupMap{})
	pulumi.RegisterOutputType(OperateGroupOutput{})
	pulumi.RegisterOutputType(OperateGroupArrayOutput{})
	pulumi.RegisterOutputType(OperateGroupMapOutput{})
}
