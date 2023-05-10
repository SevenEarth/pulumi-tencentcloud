// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EncryptAttributes struct {
	pulumi.CustomResourceState

	// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
	// 0-disable, 1-enable.
	EncryptEnabled pulumi.IntOutput `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewEncryptAttributes registers a new resource with the given unique name, arguments, and options.
func NewEncryptAttributes(ctx *pulumi.Context,
	name string, args *EncryptAttributesArgs, opts ...pulumi.ResourceOption) (*EncryptAttributes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncryptEnabled == nil {
		return nil, errors.New("invalid value for required argument 'EncryptEnabled'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EncryptAttributes
	err := ctx.RegisterResource("tencentcloud:Mariadb/encryptAttributes:EncryptAttributes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptAttributes gets an existing EncryptAttributes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptAttributes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptAttributesState, opts ...pulumi.ResourceOption) (*EncryptAttributes, error) {
	var resource EncryptAttributes
	err := ctx.ReadResource("tencentcloud:Mariadb/encryptAttributes:EncryptAttributes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptAttributes resources.
type encryptAttributesState struct {
	// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
	// 0-disable, 1-enable.
	EncryptEnabled *int `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
}

type EncryptAttributesState struct {
	// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
	// 0-disable, 1-enable.
	EncryptEnabled pulumi.IntPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
}

func (EncryptAttributesState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptAttributesState)(nil)).Elem()
}

type encryptAttributesArgs struct {
	// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
	// 0-disable, 1-enable.
	EncryptEnabled int `pulumi:"encryptEnabled"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a EncryptAttributes resource.
type EncryptAttributesArgs struct {
	// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
	// 0-disable, 1-enable.
	EncryptEnabled pulumi.IntInput
	// instance id.
	InstanceId pulumi.StringInput
}

func (EncryptAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptAttributesArgs)(nil)).Elem()
}

type EncryptAttributesInput interface {
	pulumi.Input

	ToEncryptAttributesOutput() EncryptAttributesOutput
	ToEncryptAttributesOutputWithContext(ctx context.Context) EncryptAttributesOutput
}

func (*EncryptAttributes) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptAttributes)(nil)).Elem()
}

func (i *EncryptAttributes) ToEncryptAttributesOutput() EncryptAttributesOutput {
	return i.ToEncryptAttributesOutputWithContext(context.Background())
}

func (i *EncryptAttributes) ToEncryptAttributesOutputWithContext(ctx context.Context) EncryptAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesOutput)
}

// EncryptAttributesArrayInput is an input type that accepts EncryptAttributesArray and EncryptAttributesArrayOutput values.
// You can construct a concrete instance of `EncryptAttributesArrayInput` via:
//
//          EncryptAttributesArray{ EncryptAttributesArgs{...} }
type EncryptAttributesArrayInput interface {
	pulumi.Input

	ToEncryptAttributesArrayOutput() EncryptAttributesArrayOutput
	ToEncryptAttributesArrayOutputWithContext(context.Context) EncryptAttributesArrayOutput
}

type EncryptAttributesArray []EncryptAttributesInput

func (EncryptAttributesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptAttributes)(nil)).Elem()
}

func (i EncryptAttributesArray) ToEncryptAttributesArrayOutput() EncryptAttributesArrayOutput {
	return i.ToEncryptAttributesArrayOutputWithContext(context.Background())
}

func (i EncryptAttributesArray) ToEncryptAttributesArrayOutputWithContext(ctx context.Context) EncryptAttributesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesArrayOutput)
}

// EncryptAttributesMapInput is an input type that accepts EncryptAttributesMap and EncryptAttributesMapOutput values.
// You can construct a concrete instance of `EncryptAttributesMapInput` via:
//
//          EncryptAttributesMap{ "key": EncryptAttributesArgs{...} }
type EncryptAttributesMapInput interface {
	pulumi.Input

	ToEncryptAttributesMapOutput() EncryptAttributesMapOutput
	ToEncryptAttributesMapOutputWithContext(context.Context) EncryptAttributesMapOutput
}

type EncryptAttributesMap map[string]EncryptAttributesInput

func (EncryptAttributesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptAttributes)(nil)).Elem()
}

func (i EncryptAttributesMap) ToEncryptAttributesMapOutput() EncryptAttributesMapOutput {
	return i.ToEncryptAttributesMapOutputWithContext(context.Background())
}

func (i EncryptAttributesMap) ToEncryptAttributesMapOutputWithContext(ctx context.Context) EncryptAttributesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptAttributesMapOutput)
}

type EncryptAttributesOutput struct{ *pulumi.OutputState }

func (EncryptAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptAttributes)(nil)).Elem()
}

func (o EncryptAttributesOutput) ToEncryptAttributesOutput() EncryptAttributesOutput {
	return o
}

func (o EncryptAttributesOutput) ToEncryptAttributesOutputWithContext(ctx context.Context) EncryptAttributesOutput {
	return o
}

// whether to enable data encryption, it is not supported to turn it off after it is turned on. The optional values:
// 0-disable, 1-enable.
func (o EncryptAttributesOutput) EncryptEnabled() pulumi.IntOutput {
	return o.ApplyT(func(v *EncryptAttributes) pulumi.IntOutput { return v.EncryptEnabled }).(pulumi.IntOutput)
}

// instance id.
func (o EncryptAttributesOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EncryptAttributes) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type EncryptAttributesArrayOutput struct{ *pulumi.OutputState }

func (EncryptAttributesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EncryptAttributes)(nil)).Elem()
}

func (o EncryptAttributesArrayOutput) ToEncryptAttributesArrayOutput() EncryptAttributesArrayOutput {
	return o
}

func (o EncryptAttributesArrayOutput) ToEncryptAttributesArrayOutputWithContext(ctx context.Context) EncryptAttributesArrayOutput {
	return o
}

func (o EncryptAttributesArrayOutput) Index(i pulumi.IntInput) EncryptAttributesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EncryptAttributes {
		return vs[0].([]*EncryptAttributes)[vs[1].(int)]
	}).(EncryptAttributesOutput)
}

type EncryptAttributesMapOutput struct{ *pulumi.OutputState }

func (EncryptAttributesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EncryptAttributes)(nil)).Elem()
}

func (o EncryptAttributesMapOutput) ToEncryptAttributesMapOutput() EncryptAttributesMapOutput {
	return o
}

func (o EncryptAttributesMapOutput) ToEncryptAttributesMapOutputWithContext(ctx context.Context) EncryptAttributesMapOutput {
	return o
}

func (o EncryptAttributesMapOutput) MapIndex(k pulumi.StringInput) EncryptAttributesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EncryptAttributes {
		return vs[0].(map[string]*EncryptAttributes)[vs[1].(string)]
	}).(EncryptAttributesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesInput)(nil)).Elem(), &EncryptAttributes{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesArrayInput)(nil)).Elem(), EncryptAttributesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptAttributesMapInput)(nil)).Elem(), EncryptAttributesMap{})
	pulumi.RegisterOutputType(EncryptAttributesOutput{})
	pulumi.RegisterOutputType(EncryptAttributesArrayOutput{})
	pulumi.RegisterOutputType(EncryptAttributesMapOutput{})
}
