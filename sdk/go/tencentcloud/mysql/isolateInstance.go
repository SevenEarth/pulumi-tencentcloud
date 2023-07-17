// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IsolateInstance struct {
	pulumi.CustomResourceState

	// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
	// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
	// value of the field InstanceId in the output parameter.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Instance status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIsolateInstance registers a new resource with the given unique name, arguments, and options.
func NewIsolateInstance(ctx *pulumi.Context,
	name string, args *IsolateInstanceArgs, opts ...pulumi.ResourceOption) (*IsolateInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IsolateInstance
	err := ctx.RegisterResource("tencentcloud:Mysql/isolateInstance:IsolateInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIsolateInstance gets an existing IsolateInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIsolateInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IsolateInstanceState, opts ...pulumi.ResourceOption) (*IsolateInstance, error) {
	var resource IsolateInstance
	err := ctx.ReadResource("tencentcloud:Mysql/isolateInstance:IsolateInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IsolateInstance resources.
type isolateInstanceState struct {
	// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
	// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
	// value of the field InstanceId in the output parameter.
	InstanceId *string `pulumi:"instanceId"`
	// Instance status.
	Status *string `pulumi:"status"`
}

type IsolateInstanceState struct {
	// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
	// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
	// value of the field InstanceId in the output parameter.
	InstanceId pulumi.StringPtrInput
	// Instance status.
	Status pulumi.StringPtrInput
}

func (IsolateInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*isolateInstanceState)(nil)).Elem()
}

type isolateInstanceArgs struct {
	// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
	// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
	// value of the field InstanceId in the output parameter.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a IsolateInstance resource.
type IsolateInstanceArgs struct {
	// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
	// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
	// value of the field InstanceId in the output parameter.
	InstanceId pulumi.StringInput
}

func (IsolateInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*isolateInstanceArgs)(nil)).Elem()
}

type IsolateInstanceInput interface {
	pulumi.Input

	ToIsolateInstanceOutput() IsolateInstanceOutput
	ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput
}

func (*IsolateInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**IsolateInstance)(nil)).Elem()
}

func (i *IsolateInstance) ToIsolateInstanceOutput() IsolateInstanceOutput {
	return i.ToIsolateInstanceOutputWithContext(context.Background())
}

func (i *IsolateInstance) ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceOutput)
}

// IsolateInstanceArrayInput is an input type that accepts IsolateInstanceArray and IsolateInstanceArrayOutput values.
// You can construct a concrete instance of `IsolateInstanceArrayInput` via:
//
//          IsolateInstanceArray{ IsolateInstanceArgs{...} }
type IsolateInstanceArrayInput interface {
	pulumi.Input

	ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput
	ToIsolateInstanceArrayOutputWithContext(context.Context) IsolateInstanceArrayOutput
}

type IsolateInstanceArray []IsolateInstanceInput

func (IsolateInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsolateInstance)(nil)).Elem()
}

func (i IsolateInstanceArray) ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput {
	return i.ToIsolateInstanceArrayOutputWithContext(context.Background())
}

func (i IsolateInstanceArray) ToIsolateInstanceArrayOutputWithContext(ctx context.Context) IsolateInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceArrayOutput)
}

// IsolateInstanceMapInput is an input type that accepts IsolateInstanceMap and IsolateInstanceMapOutput values.
// You can construct a concrete instance of `IsolateInstanceMapInput` via:
//
//          IsolateInstanceMap{ "key": IsolateInstanceArgs{...} }
type IsolateInstanceMapInput interface {
	pulumi.Input

	ToIsolateInstanceMapOutput() IsolateInstanceMapOutput
	ToIsolateInstanceMapOutputWithContext(context.Context) IsolateInstanceMapOutput
}

type IsolateInstanceMap map[string]IsolateInstanceInput

func (IsolateInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsolateInstance)(nil)).Elem()
}

func (i IsolateInstanceMap) ToIsolateInstanceMapOutput() IsolateInstanceMapOutput {
	return i.ToIsolateInstanceMapOutputWithContext(context.Background())
}

func (i IsolateInstanceMap) ToIsolateInstanceMapOutputWithContext(ctx context.Context) IsolateInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceMapOutput)
}

type IsolateInstanceOutput struct{ *pulumi.OutputState }

func (IsolateInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceOutput) ToIsolateInstanceOutput() IsolateInstanceOutput {
	return o
}

func (o IsolateInstanceOutput) ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput {
	return o
}

// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console
// page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the
// value of the field InstanceId in the output parameter.
func (o IsolateInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IsolateInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance status.
func (o IsolateInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IsolateInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type IsolateInstanceArrayOutput struct{ *pulumi.OutputState }

func (IsolateInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceArrayOutput) ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput {
	return o
}

func (o IsolateInstanceArrayOutput) ToIsolateInstanceArrayOutputWithContext(ctx context.Context) IsolateInstanceArrayOutput {
	return o
}

func (o IsolateInstanceArrayOutput) Index(i pulumi.IntInput) IsolateInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IsolateInstance {
		return vs[0].([]*IsolateInstance)[vs[1].(int)]
	}).(IsolateInstanceOutput)
}

type IsolateInstanceMapOutput struct{ *pulumi.OutputState }

func (IsolateInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceMapOutput) ToIsolateInstanceMapOutput() IsolateInstanceMapOutput {
	return o
}

func (o IsolateInstanceMapOutput) ToIsolateInstanceMapOutputWithContext(ctx context.Context) IsolateInstanceMapOutput {
	return o
}

func (o IsolateInstanceMapOutput) MapIndex(k pulumi.StringInput) IsolateInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IsolateInstance {
		return vs[0].(map[string]*IsolateInstance)[vs[1].(string)]
	}).(IsolateInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceInput)(nil)).Elem(), &IsolateInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceArrayInput)(nil)).Elem(), IsolateInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceMapInput)(nil)).Elem(), IsolateInstanceMap{})
	pulumi.RegisterOutputType(IsolateInstanceOutput{})
	pulumi.RegisterOutputType(IsolateInstanceArrayOutput{})
	pulumi.RegisterOutputType(IsolateInstanceMapOutput{})
}
