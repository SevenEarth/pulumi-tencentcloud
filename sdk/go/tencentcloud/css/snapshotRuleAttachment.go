// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a css snapshotRule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Css.NewSnapshotRuleAttachment(ctx, "snapshotRule", &Css.SnapshotRuleAttachmentArgs{
//				AppName:    pulumi.String("qqq"),
//				DomainName: pulumi.String("177154.push.tlivecloud.com"),
//				StreamName: pulumi.String("ppp"),
//				TemplateId: pulumi.Int(12838073),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// css snapshot_rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Css/snapshotRuleAttachment:SnapshotRuleAttachment snapshot_rule templateId#domainName
//
// ```
type SnapshotRuleAttachment struct {
	pulumi.CustomResourceState

	// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
	AppName pulumi.StringPtrOutput `pulumi:"appName"`
	// Streaming domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
	StreamName pulumi.StringPtrOutput `pulumi:"streamName"`
	// Template ID.
	TemplateId pulumi.IntOutput `pulumi:"templateId"`
}

// NewSnapshotRuleAttachment registers a new resource with the given unique name, arguments, and options.
func NewSnapshotRuleAttachment(ctx *pulumi.Context,
	name string, args *SnapshotRuleAttachmentArgs, opts ...pulumi.ResourceOption) (*SnapshotRuleAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SnapshotRuleAttachment
	err := ctx.RegisterResource("tencentcloud:Css/snapshotRuleAttachment:SnapshotRuleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotRuleAttachment gets an existing SnapshotRuleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotRuleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotRuleAttachmentState, opts ...pulumi.ResourceOption) (*SnapshotRuleAttachment, error) {
	var resource SnapshotRuleAttachment
	err := ctx.ReadResource("tencentcloud:Css/snapshotRuleAttachment:SnapshotRuleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotRuleAttachment resources.
type snapshotRuleAttachmentState struct {
	// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
	AppName *string `pulumi:"appName"`
	// Streaming domain name.
	DomainName *string `pulumi:"domainName"`
	// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
	StreamName *string `pulumi:"streamName"`
	// Template ID.
	TemplateId *int `pulumi:"templateId"`
}

type SnapshotRuleAttachmentState struct {
	// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
	AppName pulumi.StringPtrInput
	// Streaming domain name.
	DomainName pulumi.StringPtrInput
	// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
	StreamName pulumi.StringPtrInput
	// Template ID.
	TemplateId pulumi.IntPtrInput
}

func (SnapshotRuleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotRuleAttachmentState)(nil)).Elem()
}

type snapshotRuleAttachmentArgs struct {
	// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
	AppName *string `pulumi:"appName"`
	// Streaming domain name.
	DomainName string `pulumi:"domainName"`
	// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
	StreamName *string `pulumi:"streamName"`
	// Template ID.
	TemplateId int `pulumi:"templateId"`
}

// The set of arguments for constructing a SnapshotRuleAttachment resource.
type SnapshotRuleAttachmentArgs struct {
	// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
	AppName pulumi.StringPtrInput
	// Streaming domain name.
	DomainName pulumi.StringInput
	// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
	StreamName pulumi.StringPtrInput
	// Template ID.
	TemplateId pulumi.IntInput
}

func (SnapshotRuleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotRuleAttachmentArgs)(nil)).Elem()
}

type SnapshotRuleAttachmentInput interface {
	pulumi.Input

	ToSnapshotRuleAttachmentOutput() SnapshotRuleAttachmentOutput
	ToSnapshotRuleAttachmentOutputWithContext(ctx context.Context) SnapshotRuleAttachmentOutput
}

func (*SnapshotRuleAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotRuleAttachment)(nil)).Elem()
}

func (i *SnapshotRuleAttachment) ToSnapshotRuleAttachmentOutput() SnapshotRuleAttachmentOutput {
	return i.ToSnapshotRuleAttachmentOutputWithContext(context.Background())
}

func (i *SnapshotRuleAttachment) ToSnapshotRuleAttachmentOutputWithContext(ctx context.Context) SnapshotRuleAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRuleAttachmentOutput)
}

// SnapshotRuleAttachmentArrayInput is an input type that accepts SnapshotRuleAttachmentArray and SnapshotRuleAttachmentArrayOutput values.
// You can construct a concrete instance of `SnapshotRuleAttachmentArrayInput` via:
//
//	SnapshotRuleAttachmentArray{ SnapshotRuleAttachmentArgs{...} }
type SnapshotRuleAttachmentArrayInput interface {
	pulumi.Input

	ToSnapshotRuleAttachmentArrayOutput() SnapshotRuleAttachmentArrayOutput
	ToSnapshotRuleAttachmentArrayOutputWithContext(context.Context) SnapshotRuleAttachmentArrayOutput
}

type SnapshotRuleAttachmentArray []SnapshotRuleAttachmentInput

func (SnapshotRuleAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotRuleAttachment)(nil)).Elem()
}

func (i SnapshotRuleAttachmentArray) ToSnapshotRuleAttachmentArrayOutput() SnapshotRuleAttachmentArrayOutput {
	return i.ToSnapshotRuleAttachmentArrayOutputWithContext(context.Background())
}

func (i SnapshotRuleAttachmentArray) ToSnapshotRuleAttachmentArrayOutputWithContext(ctx context.Context) SnapshotRuleAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRuleAttachmentArrayOutput)
}

// SnapshotRuleAttachmentMapInput is an input type that accepts SnapshotRuleAttachmentMap and SnapshotRuleAttachmentMapOutput values.
// You can construct a concrete instance of `SnapshotRuleAttachmentMapInput` via:
//
//	SnapshotRuleAttachmentMap{ "key": SnapshotRuleAttachmentArgs{...} }
type SnapshotRuleAttachmentMapInput interface {
	pulumi.Input

	ToSnapshotRuleAttachmentMapOutput() SnapshotRuleAttachmentMapOutput
	ToSnapshotRuleAttachmentMapOutputWithContext(context.Context) SnapshotRuleAttachmentMapOutput
}

type SnapshotRuleAttachmentMap map[string]SnapshotRuleAttachmentInput

func (SnapshotRuleAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotRuleAttachment)(nil)).Elem()
}

func (i SnapshotRuleAttachmentMap) ToSnapshotRuleAttachmentMapOutput() SnapshotRuleAttachmentMapOutput {
	return i.ToSnapshotRuleAttachmentMapOutputWithContext(context.Background())
}

func (i SnapshotRuleAttachmentMap) ToSnapshotRuleAttachmentMapOutputWithContext(ctx context.Context) SnapshotRuleAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRuleAttachmentMapOutput)
}

type SnapshotRuleAttachmentOutput struct{ *pulumi.OutputState }

func (SnapshotRuleAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotRuleAttachment)(nil)).Elem()
}

func (o SnapshotRuleAttachmentOutput) ToSnapshotRuleAttachmentOutput() SnapshotRuleAttachmentOutput {
	return o
}

func (o SnapshotRuleAttachmentOutput) ToSnapshotRuleAttachmentOutputWithContext(ctx context.Context) SnapshotRuleAttachmentOutput {
	return o
}

// The streaming path is consistent with the AppName in the streaming and playback addresses. The default is live.
func (o SnapshotRuleAttachmentOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotRuleAttachment) pulumi.StringPtrOutput { return v.AppName }).(pulumi.StringPtrOutput)
}

// Streaming domain name.
func (o SnapshotRuleAttachmentOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotRuleAttachment) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Stream name. Note: If this parameter is set to a non empty string, the rule will only work on this streaming.
func (o SnapshotRuleAttachmentOutput) StreamName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotRuleAttachment) pulumi.StringPtrOutput { return v.StreamName }).(pulumi.StringPtrOutput)
}

// Template ID.
func (o SnapshotRuleAttachmentOutput) TemplateId() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotRuleAttachment) pulumi.IntOutput { return v.TemplateId }).(pulumi.IntOutput)
}

type SnapshotRuleAttachmentArrayOutput struct{ *pulumi.OutputState }

func (SnapshotRuleAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotRuleAttachment)(nil)).Elem()
}

func (o SnapshotRuleAttachmentArrayOutput) ToSnapshotRuleAttachmentArrayOutput() SnapshotRuleAttachmentArrayOutput {
	return o
}

func (o SnapshotRuleAttachmentArrayOutput) ToSnapshotRuleAttachmentArrayOutputWithContext(ctx context.Context) SnapshotRuleAttachmentArrayOutput {
	return o
}

func (o SnapshotRuleAttachmentArrayOutput) Index(i pulumi.IntInput) SnapshotRuleAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotRuleAttachment {
		return vs[0].([]*SnapshotRuleAttachment)[vs[1].(int)]
	}).(SnapshotRuleAttachmentOutput)
}

type SnapshotRuleAttachmentMapOutput struct{ *pulumi.OutputState }

func (SnapshotRuleAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotRuleAttachment)(nil)).Elem()
}

func (o SnapshotRuleAttachmentMapOutput) ToSnapshotRuleAttachmentMapOutput() SnapshotRuleAttachmentMapOutput {
	return o
}

func (o SnapshotRuleAttachmentMapOutput) ToSnapshotRuleAttachmentMapOutputWithContext(ctx context.Context) SnapshotRuleAttachmentMapOutput {
	return o
}

func (o SnapshotRuleAttachmentMapOutput) MapIndex(k pulumi.StringInput) SnapshotRuleAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotRuleAttachment {
		return vs[0].(map[string]*SnapshotRuleAttachment)[vs[1].(string)]
	}).(SnapshotRuleAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRuleAttachmentInput)(nil)).Elem(), &SnapshotRuleAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRuleAttachmentArrayInput)(nil)).Elem(), SnapshotRuleAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRuleAttachmentMapInput)(nil)).Elem(), SnapshotRuleAttachmentMap{})
	pulumi.RegisterOutputType(SnapshotRuleAttachmentOutput{})
	pulumi.RegisterOutputType(SnapshotRuleAttachmentArrayOutput{})
	pulumi.RegisterOutputType(SnapshotRuleAttachmentMapOutput{})
}
