// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to attach an existing subnet to Network ACL.
//
// ## Import
//
// Acl attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vpc/aclAttachment:AclAttachment attachment acl-eotx5qsg#subnet-91x0geu6
//
// ```
type AclAttachment struct {
	pulumi.CustomResourceState

	// ID of the attached ACL.
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// The Subnet instance ID.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewAclAttachment registers a new resource with the given unique name, arguments, and options.
func NewAclAttachment(ctx *pulumi.Context,
	name string, args *AclAttachmentArgs, opts ...pulumi.ResourceOption) (*AclAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclId == nil {
		return nil, errors.New("invalid value for required argument 'AclId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AclAttachment
	err := ctx.RegisterResource("tencentcloud:Vpc/aclAttachment:AclAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclAttachment gets an existing AclAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclAttachmentState, opts ...pulumi.ResourceOption) (*AclAttachment, error) {
	var resource AclAttachment
	err := ctx.ReadResource("tencentcloud:Vpc/aclAttachment:AclAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclAttachment resources.
type aclAttachmentState struct {
	// ID of the attached ACL.
	AclId *string `pulumi:"aclId"`
	// The Subnet instance ID.
	SubnetId *string `pulumi:"subnetId"`
}

type AclAttachmentState struct {
	// ID of the attached ACL.
	AclId pulumi.StringPtrInput
	// The Subnet instance ID.
	SubnetId pulumi.StringPtrInput
}

func (AclAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclAttachmentState)(nil)).Elem()
}

type aclAttachmentArgs struct {
	// ID of the attached ACL.
	AclId string `pulumi:"aclId"`
	// The Subnet instance ID.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a AclAttachment resource.
type AclAttachmentArgs struct {
	// ID of the attached ACL.
	AclId pulumi.StringInput
	// The Subnet instance ID.
	SubnetId pulumi.StringInput
}

func (AclAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclAttachmentArgs)(nil)).Elem()
}

type AclAttachmentInput interface {
	pulumi.Input

	ToAclAttachmentOutput() AclAttachmentOutput
	ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput
}

func (*AclAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AclAttachment)(nil)).Elem()
}

func (i *AclAttachment) ToAclAttachmentOutput() AclAttachmentOutput {
	return i.ToAclAttachmentOutputWithContext(context.Background())
}

func (i *AclAttachment) ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentOutput)
}

// AclAttachmentArrayInput is an input type that accepts AclAttachmentArray and AclAttachmentArrayOutput values.
// You can construct a concrete instance of `AclAttachmentArrayInput` via:
//
//	AclAttachmentArray{ AclAttachmentArgs{...} }
type AclAttachmentArrayInput interface {
	pulumi.Input

	ToAclAttachmentArrayOutput() AclAttachmentArrayOutput
	ToAclAttachmentArrayOutputWithContext(context.Context) AclAttachmentArrayOutput
}

type AclAttachmentArray []AclAttachmentInput

func (AclAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclAttachment)(nil)).Elem()
}

func (i AclAttachmentArray) ToAclAttachmentArrayOutput() AclAttachmentArrayOutput {
	return i.ToAclAttachmentArrayOutputWithContext(context.Background())
}

func (i AclAttachmentArray) ToAclAttachmentArrayOutputWithContext(ctx context.Context) AclAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentArrayOutput)
}

// AclAttachmentMapInput is an input type that accepts AclAttachmentMap and AclAttachmentMapOutput values.
// You can construct a concrete instance of `AclAttachmentMapInput` via:
//
//	AclAttachmentMap{ "key": AclAttachmentArgs{...} }
type AclAttachmentMapInput interface {
	pulumi.Input

	ToAclAttachmentMapOutput() AclAttachmentMapOutput
	ToAclAttachmentMapOutputWithContext(context.Context) AclAttachmentMapOutput
}

type AclAttachmentMap map[string]AclAttachmentInput

func (AclAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclAttachment)(nil)).Elem()
}

func (i AclAttachmentMap) ToAclAttachmentMapOutput() AclAttachmentMapOutput {
	return i.ToAclAttachmentMapOutputWithContext(context.Background())
}

func (i AclAttachmentMap) ToAclAttachmentMapOutputWithContext(ctx context.Context) AclAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclAttachmentMapOutput)
}

type AclAttachmentOutput struct{ *pulumi.OutputState }

func (AclAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclAttachment)(nil)).Elem()
}

func (o AclAttachmentOutput) ToAclAttachmentOutput() AclAttachmentOutput {
	return o
}

func (o AclAttachmentOutput) ToAclAttachmentOutputWithContext(ctx context.Context) AclAttachmentOutput {
	return o
}

// ID of the attached ACL.
func (o AclAttachmentOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// The Subnet instance ID.
func (o AclAttachmentOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclAttachment) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

type AclAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AclAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclAttachment)(nil)).Elem()
}

func (o AclAttachmentArrayOutput) ToAclAttachmentArrayOutput() AclAttachmentArrayOutput {
	return o
}

func (o AclAttachmentArrayOutput) ToAclAttachmentArrayOutputWithContext(ctx context.Context) AclAttachmentArrayOutput {
	return o
}

func (o AclAttachmentArrayOutput) Index(i pulumi.IntInput) AclAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclAttachment {
		return vs[0].([]*AclAttachment)[vs[1].(int)]
	}).(AclAttachmentOutput)
}

type AclAttachmentMapOutput struct{ *pulumi.OutputState }

func (AclAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclAttachment)(nil)).Elem()
}

func (o AclAttachmentMapOutput) ToAclAttachmentMapOutput() AclAttachmentMapOutput {
	return o
}

func (o AclAttachmentMapOutput) ToAclAttachmentMapOutputWithContext(ctx context.Context) AclAttachmentMapOutput {
	return o
}

func (o AclAttachmentMapOutput) MapIndex(k pulumi.StringInput) AclAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclAttachment {
		return vs[0].(map[string]*AclAttachment)[vs[1].(string)]
	}).(AclAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentInput)(nil)).Elem(), &AclAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentArrayInput)(nil)).Elem(), AclAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclAttachmentMapInput)(nil)).Elem(), AclAttachmentMap{})
	pulumi.RegisterOutputType(AclAttachmentOutput{})
	pulumi.RegisterOutputType(AclAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AclAttachmentMapOutput{})
}
