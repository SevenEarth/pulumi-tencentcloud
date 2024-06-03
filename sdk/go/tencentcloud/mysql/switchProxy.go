// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

type SwitchProxy struct {
	pulumi.CustomResourceState

	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Proxy group id.
	ProxyGroupId pulumi.StringOutput `pulumi:"proxyGroupId"`
}

// NewSwitchProxy registers a new resource with the given unique name, arguments, and options.
func NewSwitchProxy(ctx *pulumi.Context,
	name string, args *SwitchProxyArgs, opts ...pulumi.ResourceOption) (*SwitchProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ProxyGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchProxy
	err := ctx.RegisterResource("tencentcloud:Mysql/switchProxy:SwitchProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchProxy gets an existing SwitchProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchProxyState, opts ...pulumi.ResourceOption) (*SwitchProxy, error) {
	var resource SwitchProxy
	err := ctx.ReadResource("tencentcloud:Mysql/switchProxy:SwitchProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchProxy resources.
type switchProxyState struct {
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Proxy group id.
	ProxyGroupId *string `pulumi:"proxyGroupId"`
}

type SwitchProxyState struct {
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Proxy group id.
	ProxyGroupId pulumi.StringPtrInput
}

func (SwitchProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchProxyState)(nil)).Elem()
}

type switchProxyArgs struct {
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Proxy group id.
	ProxyGroupId string `pulumi:"proxyGroupId"`
}

// The set of arguments for constructing a SwitchProxy resource.
type SwitchProxyArgs struct {
	// Instance id.
	InstanceId pulumi.StringInput
	// Proxy group id.
	ProxyGroupId pulumi.StringInput
}

func (SwitchProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchProxyArgs)(nil)).Elem()
}

type SwitchProxyInput interface {
	pulumi.Input

	ToSwitchProxyOutput() SwitchProxyOutput
	ToSwitchProxyOutputWithContext(ctx context.Context) SwitchProxyOutput
}

func (*SwitchProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchProxy)(nil)).Elem()
}

func (i *SwitchProxy) ToSwitchProxyOutput() SwitchProxyOutput {
	return i.ToSwitchProxyOutputWithContext(context.Background())
}

func (i *SwitchProxy) ToSwitchProxyOutputWithContext(ctx context.Context) SwitchProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchProxyOutput)
}

// SwitchProxyArrayInput is an input type that accepts SwitchProxyArray and SwitchProxyArrayOutput values.
// You can construct a concrete instance of `SwitchProxyArrayInput` via:
//
//	SwitchProxyArray{ SwitchProxyArgs{...} }
type SwitchProxyArrayInput interface {
	pulumi.Input

	ToSwitchProxyArrayOutput() SwitchProxyArrayOutput
	ToSwitchProxyArrayOutputWithContext(context.Context) SwitchProxyArrayOutput
}

type SwitchProxyArray []SwitchProxyInput

func (SwitchProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchProxy)(nil)).Elem()
}

func (i SwitchProxyArray) ToSwitchProxyArrayOutput() SwitchProxyArrayOutput {
	return i.ToSwitchProxyArrayOutputWithContext(context.Background())
}

func (i SwitchProxyArray) ToSwitchProxyArrayOutputWithContext(ctx context.Context) SwitchProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchProxyArrayOutput)
}

// SwitchProxyMapInput is an input type that accepts SwitchProxyMap and SwitchProxyMapOutput values.
// You can construct a concrete instance of `SwitchProxyMapInput` via:
//
//	SwitchProxyMap{ "key": SwitchProxyArgs{...} }
type SwitchProxyMapInput interface {
	pulumi.Input

	ToSwitchProxyMapOutput() SwitchProxyMapOutput
	ToSwitchProxyMapOutputWithContext(context.Context) SwitchProxyMapOutput
}

type SwitchProxyMap map[string]SwitchProxyInput

func (SwitchProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchProxy)(nil)).Elem()
}

func (i SwitchProxyMap) ToSwitchProxyMapOutput() SwitchProxyMapOutput {
	return i.ToSwitchProxyMapOutputWithContext(context.Background())
}

func (i SwitchProxyMap) ToSwitchProxyMapOutputWithContext(ctx context.Context) SwitchProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchProxyMapOutput)
}

type SwitchProxyOutput struct{ *pulumi.OutputState }

func (SwitchProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchProxy)(nil)).Elem()
}

func (o SwitchProxyOutput) ToSwitchProxyOutput() SwitchProxyOutput {
	return o
}

func (o SwitchProxyOutput) ToSwitchProxyOutputWithContext(ctx context.Context) SwitchProxyOutput {
	return o
}

// Instance id.
func (o SwitchProxyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchProxy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Proxy group id.
func (o SwitchProxyOutput) ProxyGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchProxy) pulumi.StringOutput { return v.ProxyGroupId }).(pulumi.StringOutput)
}

type SwitchProxyArrayOutput struct{ *pulumi.OutputState }

func (SwitchProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchProxy)(nil)).Elem()
}

func (o SwitchProxyArrayOutput) ToSwitchProxyArrayOutput() SwitchProxyArrayOutput {
	return o
}

func (o SwitchProxyArrayOutput) ToSwitchProxyArrayOutputWithContext(ctx context.Context) SwitchProxyArrayOutput {
	return o
}

func (o SwitchProxyArrayOutput) Index(i pulumi.IntInput) SwitchProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchProxy {
		return vs[0].([]*SwitchProxy)[vs[1].(int)]
	}).(SwitchProxyOutput)
}

type SwitchProxyMapOutput struct{ *pulumi.OutputState }

func (SwitchProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchProxy)(nil)).Elem()
}

func (o SwitchProxyMapOutput) ToSwitchProxyMapOutput() SwitchProxyMapOutput {
	return o
}

func (o SwitchProxyMapOutput) ToSwitchProxyMapOutputWithContext(ctx context.Context) SwitchProxyMapOutput {
	return o
}

func (o SwitchProxyMapOutput) MapIndex(k pulumi.StringInput) SwitchProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchProxy {
		return vs[0].(map[string]*SwitchProxy)[vs[1].(string)]
	}).(SwitchProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchProxyInput)(nil)).Elem(), &SwitchProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchProxyArrayInput)(nil)).Elem(), SwitchProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchProxyMapInput)(nil)).Elem(), SwitchProxyMap{})
	pulumi.RegisterOutputType(SwitchProxyOutput{})
	pulumi.RegisterOutputType(SwitchProxyArrayOutput{})
	pulumi.RegisterOutputType(SwitchProxyMapOutput{})
}
