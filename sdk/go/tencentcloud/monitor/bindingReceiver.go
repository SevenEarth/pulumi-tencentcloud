// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for bind receivers to a policy group resource.
type BindingReceiver struct {
	pulumi.CustomResourceState

	// Policy group ID for binding receivers.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	Receivers BindingReceiverReceiversPtrOutput `pulumi:"receivers"`
}

// NewBindingReceiver registers a new resource with the given unique name, arguments, and options.
func NewBindingReceiver(ctx *pulumi.Context,
	name string, args *BindingReceiverArgs, opts ...pulumi.ResourceOption) (*BindingReceiver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BindingReceiver
	err := ctx.RegisterResource("tencentcloud:Monitor/bindingReceiver:BindingReceiver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBindingReceiver gets an existing BindingReceiver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBindingReceiver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingReceiverState, opts ...pulumi.ResourceOption) (*BindingReceiver, error) {
	var resource BindingReceiver
	err := ctx.ReadResource("tencentcloud:Monitor/bindingReceiver:BindingReceiver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BindingReceiver resources.
type bindingReceiverState struct {
	// Policy group ID for binding receivers.
	GroupId *int `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	Receivers *BindingReceiverReceivers `pulumi:"receivers"`
}

type BindingReceiverState struct {
	// Policy group ID for binding receivers.
	GroupId pulumi.IntPtrInput
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	Receivers BindingReceiverReceiversPtrInput
}

func (BindingReceiverState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingReceiverState)(nil)).Elem()
}

type bindingReceiverArgs struct {
	// Policy group ID for binding receivers.
	GroupId int `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	Receivers *BindingReceiverReceivers `pulumi:"receivers"`
}

// The set of arguments for constructing a BindingReceiver resource.
type BindingReceiverArgs struct {
	// Policy group ID for binding receivers.
	GroupId pulumi.IntInput
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	Receivers BindingReceiverReceiversPtrInput
}

func (BindingReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingReceiverArgs)(nil)).Elem()
}

type BindingReceiverInput interface {
	pulumi.Input

	ToBindingReceiverOutput() BindingReceiverOutput
	ToBindingReceiverOutputWithContext(ctx context.Context) BindingReceiverOutput
}

func (*BindingReceiver) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingReceiver)(nil)).Elem()
}

func (i *BindingReceiver) ToBindingReceiverOutput() BindingReceiverOutput {
	return i.ToBindingReceiverOutputWithContext(context.Background())
}

func (i *BindingReceiver) ToBindingReceiverOutputWithContext(ctx context.Context) BindingReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingReceiverOutput)
}

// BindingReceiverArrayInput is an input type that accepts BindingReceiverArray and BindingReceiverArrayOutput values.
// You can construct a concrete instance of `BindingReceiverArrayInput` via:
//
//          BindingReceiverArray{ BindingReceiverArgs{...} }
type BindingReceiverArrayInput interface {
	pulumi.Input

	ToBindingReceiverArrayOutput() BindingReceiverArrayOutput
	ToBindingReceiverArrayOutputWithContext(context.Context) BindingReceiverArrayOutput
}

type BindingReceiverArray []BindingReceiverInput

func (BindingReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindingReceiver)(nil)).Elem()
}

func (i BindingReceiverArray) ToBindingReceiverArrayOutput() BindingReceiverArrayOutput {
	return i.ToBindingReceiverArrayOutputWithContext(context.Background())
}

func (i BindingReceiverArray) ToBindingReceiverArrayOutputWithContext(ctx context.Context) BindingReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingReceiverArrayOutput)
}

// BindingReceiverMapInput is an input type that accepts BindingReceiverMap and BindingReceiverMapOutput values.
// You can construct a concrete instance of `BindingReceiverMapInput` via:
//
//          BindingReceiverMap{ "key": BindingReceiverArgs{...} }
type BindingReceiverMapInput interface {
	pulumi.Input

	ToBindingReceiverMapOutput() BindingReceiverMapOutput
	ToBindingReceiverMapOutputWithContext(context.Context) BindingReceiverMapOutput
}

type BindingReceiverMap map[string]BindingReceiverInput

func (BindingReceiverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindingReceiver)(nil)).Elem()
}

func (i BindingReceiverMap) ToBindingReceiverMapOutput() BindingReceiverMapOutput {
	return i.ToBindingReceiverMapOutputWithContext(context.Background())
}

func (i BindingReceiverMap) ToBindingReceiverMapOutputWithContext(ctx context.Context) BindingReceiverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingReceiverMapOutput)
}

type BindingReceiverOutput struct{ *pulumi.OutputState }

func (BindingReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingReceiver)(nil)).Elem()
}

func (o BindingReceiverOutput) ToBindingReceiverOutput() BindingReceiverOutput {
	return o
}

func (o BindingReceiverOutput) ToBindingReceiverOutputWithContext(ctx context.Context) BindingReceiverOutput {
	return o
}

// Policy group ID for binding receivers.
func (o BindingReceiverOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *BindingReceiver) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
func (o BindingReceiverOutput) Receivers() BindingReceiverReceiversPtrOutput {
	return o.ApplyT(func(v *BindingReceiver) BindingReceiverReceiversPtrOutput { return v.Receivers }).(BindingReceiverReceiversPtrOutput)
}

type BindingReceiverArrayOutput struct{ *pulumi.OutputState }

func (BindingReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindingReceiver)(nil)).Elem()
}

func (o BindingReceiverArrayOutput) ToBindingReceiverArrayOutput() BindingReceiverArrayOutput {
	return o
}

func (o BindingReceiverArrayOutput) ToBindingReceiverArrayOutputWithContext(ctx context.Context) BindingReceiverArrayOutput {
	return o
}

func (o BindingReceiverArrayOutput) Index(i pulumi.IntInput) BindingReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BindingReceiver {
		return vs[0].([]*BindingReceiver)[vs[1].(int)]
	}).(BindingReceiverOutput)
}

type BindingReceiverMapOutput struct{ *pulumi.OutputState }

func (BindingReceiverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindingReceiver)(nil)).Elem()
}

func (o BindingReceiverMapOutput) ToBindingReceiverMapOutput() BindingReceiverMapOutput {
	return o
}

func (o BindingReceiverMapOutput) ToBindingReceiverMapOutputWithContext(ctx context.Context) BindingReceiverMapOutput {
	return o
}

func (o BindingReceiverMapOutput) MapIndex(k pulumi.StringInput) BindingReceiverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BindingReceiver {
		return vs[0].(map[string]*BindingReceiver)[vs[1].(string)]
	}).(BindingReceiverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingReceiverInput)(nil)).Elem(), &BindingReceiver{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingReceiverArrayInput)(nil)).Elem(), BindingReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingReceiverMapInput)(nil)).Elem(), BindingReceiverMap{})
	pulumi.RegisterOutputType(BindingReceiverOutput{})
	pulumi.RegisterOutputType(BindingReceiverArrayOutput{})
	pulumi.RegisterOutputType(BindingReceiverMapOutput{})
}
