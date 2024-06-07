// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tmp tke template attachment
type TmpTkeTemplateAttachment struct {
	pulumi.CustomResourceState

	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsOutput `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewTmpTkeTemplateAttachment registers a new resource with the given unique name, arguments, and options.
func NewTmpTkeTemplateAttachment(ctx *pulumi.Context,
	name string, args *TmpTkeTemplateAttachmentArgs, opts ...pulumi.ResourceOption) (*TmpTkeTemplateAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TmpTkeTemplateAttachment
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpTkeTemplateAttachment:TmpTkeTemplateAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpTkeTemplateAttachment gets an existing TmpTkeTemplateAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpTkeTemplateAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpTkeTemplateAttachmentState, opts ...pulumi.ResourceOption) (*TmpTkeTemplateAttachment, error) {
	var resource TmpTkeTemplateAttachment
	err := ctx.ReadResource("tencentcloud:Monitor/tmpTkeTemplateAttachment:TmpTkeTemplateAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpTkeTemplateAttachment resources.
type tmpTkeTemplateAttachmentState struct {
	// Sync target details.
	Targets *TmpTkeTemplateAttachmentTargets `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId *string `pulumi:"templateId"`
}

type TmpTkeTemplateAttachmentState struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsPtrInput
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringPtrInput
}

func (TmpTkeTemplateAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeTemplateAttachmentState)(nil)).Elem()
}

type tmpTkeTemplateAttachmentArgs struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargets `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId string `pulumi:"templateId"`
}

// The set of arguments for constructing a TmpTkeTemplateAttachment resource.
type TmpTkeTemplateAttachmentArgs struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsInput
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringInput
}

func (TmpTkeTemplateAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeTemplateAttachmentArgs)(nil)).Elem()
}

type TmpTkeTemplateAttachmentInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput
	ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput
}

func (*TmpTkeTemplateAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i *TmpTkeTemplateAttachment) ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput {
	return i.ToTmpTkeTemplateAttachmentOutputWithContext(context.Background())
}

func (i *TmpTkeTemplateAttachment) ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentOutput)
}

// TmpTkeTemplateAttachmentArrayInput is an input type that accepts TmpTkeTemplateAttachmentArray and TmpTkeTemplateAttachmentArrayOutput values.
// You can construct a concrete instance of `TmpTkeTemplateAttachmentArrayInput` via:
//
//	TmpTkeTemplateAttachmentArray{ TmpTkeTemplateAttachmentArgs{...} }
type TmpTkeTemplateAttachmentArrayInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput
	ToTmpTkeTemplateAttachmentArrayOutputWithContext(context.Context) TmpTkeTemplateAttachmentArrayOutput
}

type TmpTkeTemplateAttachmentArray []TmpTkeTemplateAttachmentInput

func (TmpTkeTemplateAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i TmpTkeTemplateAttachmentArray) ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput {
	return i.ToTmpTkeTemplateAttachmentArrayOutputWithContext(context.Background())
}

func (i TmpTkeTemplateAttachmentArray) ToTmpTkeTemplateAttachmentArrayOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentArrayOutput)
}

// TmpTkeTemplateAttachmentMapInput is an input type that accepts TmpTkeTemplateAttachmentMap and TmpTkeTemplateAttachmentMapOutput values.
// You can construct a concrete instance of `TmpTkeTemplateAttachmentMapInput` via:
//
//	TmpTkeTemplateAttachmentMap{ "key": TmpTkeTemplateAttachmentArgs{...} }
type TmpTkeTemplateAttachmentMapInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput
	ToTmpTkeTemplateAttachmentMapOutputWithContext(context.Context) TmpTkeTemplateAttachmentMapOutput
}

type TmpTkeTemplateAttachmentMap map[string]TmpTkeTemplateAttachmentInput

func (TmpTkeTemplateAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i TmpTkeTemplateAttachmentMap) ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput {
	return i.ToTmpTkeTemplateAttachmentMapOutputWithContext(context.Background())
}

func (i TmpTkeTemplateAttachmentMap) ToTmpTkeTemplateAttachmentMapOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentMapOutput)
}

type TmpTkeTemplateAttachmentOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentOutput) ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput {
	return o
}

func (o TmpTkeTemplateAttachmentOutput) ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput {
	return o
}

// Sync target details.
func (o TmpTkeTemplateAttachmentOutput) Targets() TmpTkeTemplateAttachmentTargetsOutput {
	return o.ApplyT(func(v *TmpTkeTemplateAttachment) TmpTkeTemplateAttachmentTargetsOutput { return v.Targets }).(TmpTkeTemplateAttachmentTargetsOutput)
}

// The ID of the template, which is used for the outgoing reference.
func (o TmpTkeTemplateAttachmentOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeTemplateAttachment) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

type TmpTkeTemplateAttachmentArrayOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentArrayOutput) ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput {
	return o
}

func (o TmpTkeTemplateAttachmentArrayOutput) ToTmpTkeTemplateAttachmentArrayOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentArrayOutput {
	return o
}

func (o TmpTkeTemplateAttachmentArrayOutput) Index(i pulumi.IntInput) TmpTkeTemplateAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpTkeTemplateAttachment {
		return vs[0].([]*TmpTkeTemplateAttachment)[vs[1].(int)]
	}).(TmpTkeTemplateAttachmentOutput)
}

type TmpTkeTemplateAttachmentMapOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentMapOutput) ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput {
	return o
}

func (o TmpTkeTemplateAttachmentMapOutput) ToTmpTkeTemplateAttachmentMapOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentMapOutput {
	return o
}

func (o TmpTkeTemplateAttachmentMapOutput) MapIndex(k pulumi.StringInput) TmpTkeTemplateAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpTkeTemplateAttachment {
		return vs[0].(map[string]*TmpTkeTemplateAttachment)[vs[1].(string)]
	}).(TmpTkeTemplateAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentInput)(nil)).Elem(), &TmpTkeTemplateAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentArrayInput)(nil)).Elem(), TmpTkeTemplateAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentMapInput)(nil)).Elem(), TmpTkeTemplateAttachmentMap{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentOutput{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentArrayOutput{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentMapOutput{})
}
