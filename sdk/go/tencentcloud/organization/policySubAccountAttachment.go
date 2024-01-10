// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a organization policySubAccountAttachment
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Organization"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Organization.NewPolicySubAccountAttachment(ctx, "policySubAccountAttachment", &Organization.PolicySubAccountAttachmentArgs{
//				MemberUin:        pulumi.Int(100028582828),
//				OrgSubAccountUin: pulumi.Int(100028223737),
//				PolicyId:         pulumi.Int(144256499),
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
// organization policy_sub_account_attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment policy_sub_account_attachment policyId#memberUin#orgSubAccountUin
//
// ```
type PolicySubAccountAttachment struct {
	pulumi.CustomResourceState

	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Manage Identity ID.
	IdentityId pulumi.IntOutput `pulumi:"identityId"`
	// Identity role alias name.
	IdentityRoleAliasName pulumi.StringOutput `pulumi:"identityRoleAliasName"`
	// Identity role name.
	IdentityRoleName pulumi.StringOutput `pulumi:"identityRoleName"`
	// Organization member uin.
	MemberUin pulumi.IntOutput `pulumi:"memberUin"`
	// Organization administrator sub account name.
	OrgSubAccountName pulumi.StringOutput `pulumi:"orgSubAccountName"`
	// Organization administrator sub account uin list.
	OrgSubAccountUin pulumi.IntOutput `pulumi:"orgSubAccountUin"`
	// Policy ID.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// Policy name.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPolicySubAccountAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicySubAccountAttachment(ctx *pulumi.Context,
	name string, args *PolicySubAccountAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicySubAccountAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemberUin == nil {
		return nil, errors.New("invalid value for required argument 'MemberUin'")
	}
	if args.OrgSubAccountUin == nil {
		return nil, errors.New("invalid value for required argument 'OrgSubAccountUin'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PolicySubAccountAttachment
	err := ctx.RegisterResource("tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySubAccountAttachment gets an existing PolicySubAccountAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySubAccountAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySubAccountAttachmentState, opts ...pulumi.ResourceOption) (*PolicySubAccountAttachment, error) {
	var resource PolicySubAccountAttachment
	err := ctx.ReadResource("tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySubAccountAttachment resources.
type policySubAccountAttachmentState struct {
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// Manage Identity ID.
	IdentityId *int `pulumi:"identityId"`
	// Identity role alias name.
	IdentityRoleAliasName *string `pulumi:"identityRoleAliasName"`
	// Identity role name.
	IdentityRoleName *string `pulumi:"identityRoleName"`
	// Organization member uin.
	MemberUin *int `pulumi:"memberUin"`
	// Organization administrator sub account name.
	OrgSubAccountName *string `pulumi:"orgSubAccountName"`
	// Organization administrator sub account uin list.
	OrgSubAccountUin *int `pulumi:"orgSubAccountUin"`
	// Policy ID.
	PolicyId *int `pulumi:"policyId"`
	// Policy name.
	PolicyName *string `pulumi:"policyName"`
	// Update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type PolicySubAccountAttachmentState struct {
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// Manage Identity ID.
	IdentityId pulumi.IntPtrInput
	// Identity role alias name.
	IdentityRoleAliasName pulumi.StringPtrInput
	// Identity role name.
	IdentityRoleName pulumi.StringPtrInput
	// Organization member uin.
	MemberUin pulumi.IntPtrInput
	// Organization administrator sub account name.
	OrgSubAccountName pulumi.StringPtrInput
	// Organization administrator sub account uin list.
	OrgSubAccountUin pulumi.IntPtrInput
	// Policy ID.
	PolicyId pulumi.IntPtrInput
	// Policy name.
	PolicyName pulumi.StringPtrInput
	// Update time.
	UpdateTime pulumi.StringPtrInput
}

func (PolicySubAccountAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySubAccountAttachmentState)(nil)).Elem()
}

type policySubAccountAttachmentArgs struct {
	// Organization member uin.
	MemberUin int `pulumi:"memberUin"`
	// Organization administrator sub account uin list.
	OrgSubAccountUin int `pulumi:"orgSubAccountUin"`
	// Policy ID.
	PolicyId int `pulumi:"policyId"`
}

// The set of arguments for constructing a PolicySubAccountAttachment resource.
type PolicySubAccountAttachmentArgs struct {
	// Organization member uin.
	MemberUin pulumi.IntInput
	// Organization administrator sub account uin list.
	OrgSubAccountUin pulumi.IntInput
	// Policy ID.
	PolicyId pulumi.IntInput
}

func (PolicySubAccountAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySubAccountAttachmentArgs)(nil)).Elem()
}

type PolicySubAccountAttachmentInput interface {
	pulumi.Input

	ToPolicySubAccountAttachmentOutput() PolicySubAccountAttachmentOutput
	ToPolicySubAccountAttachmentOutputWithContext(ctx context.Context) PolicySubAccountAttachmentOutput
}

func (*PolicySubAccountAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySubAccountAttachment)(nil)).Elem()
}

func (i *PolicySubAccountAttachment) ToPolicySubAccountAttachmentOutput() PolicySubAccountAttachmentOutput {
	return i.ToPolicySubAccountAttachmentOutputWithContext(context.Background())
}

func (i *PolicySubAccountAttachment) ToPolicySubAccountAttachmentOutputWithContext(ctx context.Context) PolicySubAccountAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySubAccountAttachmentOutput)
}

// PolicySubAccountAttachmentArrayInput is an input type that accepts PolicySubAccountAttachmentArray and PolicySubAccountAttachmentArrayOutput values.
// You can construct a concrete instance of `PolicySubAccountAttachmentArrayInput` via:
//
//	PolicySubAccountAttachmentArray{ PolicySubAccountAttachmentArgs{...} }
type PolicySubAccountAttachmentArrayInput interface {
	pulumi.Input

	ToPolicySubAccountAttachmentArrayOutput() PolicySubAccountAttachmentArrayOutput
	ToPolicySubAccountAttachmentArrayOutputWithContext(context.Context) PolicySubAccountAttachmentArrayOutput
}

type PolicySubAccountAttachmentArray []PolicySubAccountAttachmentInput

func (PolicySubAccountAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicySubAccountAttachment)(nil)).Elem()
}

func (i PolicySubAccountAttachmentArray) ToPolicySubAccountAttachmentArrayOutput() PolicySubAccountAttachmentArrayOutput {
	return i.ToPolicySubAccountAttachmentArrayOutputWithContext(context.Background())
}

func (i PolicySubAccountAttachmentArray) ToPolicySubAccountAttachmentArrayOutputWithContext(ctx context.Context) PolicySubAccountAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySubAccountAttachmentArrayOutput)
}

// PolicySubAccountAttachmentMapInput is an input type that accepts PolicySubAccountAttachmentMap and PolicySubAccountAttachmentMapOutput values.
// You can construct a concrete instance of `PolicySubAccountAttachmentMapInput` via:
//
//	PolicySubAccountAttachmentMap{ "key": PolicySubAccountAttachmentArgs{...} }
type PolicySubAccountAttachmentMapInput interface {
	pulumi.Input

	ToPolicySubAccountAttachmentMapOutput() PolicySubAccountAttachmentMapOutput
	ToPolicySubAccountAttachmentMapOutputWithContext(context.Context) PolicySubAccountAttachmentMapOutput
}

type PolicySubAccountAttachmentMap map[string]PolicySubAccountAttachmentInput

func (PolicySubAccountAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicySubAccountAttachment)(nil)).Elem()
}

func (i PolicySubAccountAttachmentMap) ToPolicySubAccountAttachmentMapOutput() PolicySubAccountAttachmentMapOutput {
	return i.ToPolicySubAccountAttachmentMapOutputWithContext(context.Background())
}

func (i PolicySubAccountAttachmentMap) ToPolicySubAccountAttachmentMapOutputWithContext(ctx context.Context) PolicySubAccountAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySubAccountAttachmentMapOutput)
}

type PolicySubAccountAttachmentOutput struct{ *pulumi.OutputState }

func (PolicySubAccountAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySubAccountAttachment)(nil)).Elem()
}

func (o PolicySubAccountAttachmentOutput) ToPolicySubAccountAttachmentOutput() PolicySubAccountAttachmentOutput {
	return o
}

func (o PolicySubAccountAttachmentOutput) ToPolicySubAccountAttachmentOutputWithContext(ctx context.Context) PolicySubAccountAttachmentOutput {
	return o
}

// Creation time.
func (o PolicySubAccountAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Manage Identity ID.
func (o PolicySubAccountAttachmentOutput) IdentityId() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.IntOutput { return v.IdentityId }).(pulumi.IntOutput)
}

// Identity role alias name.
func (o PolicySubAccountAttachmentOutput) IdentityRoleAliasName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.IdentityRoleAliasName }).(pulumi.StringOutput)
}

// Identity role name.
func (o PolicySubAccountAttachmentOutput) IdentityRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.IdentityRoleName }).(pulumi.StringOutput)
}

// Organization member uin.
func (o PolicySubAccountAttachmentOutput) MemberUin() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.IntOutput { return v.MemberUin }).(pulumi.IntOutput)
}

// Organization administrator sub account name.
func (o PolicySubAccountAttachmentOutput) OrgSubAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.OrgSubAccountName }).(pulumi.StringOutput)
}

// Organization administrator sub account uin list.
func (o PolicySubAccountAttachmentOutput) OrgSubAccountUin() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.IntOutput { return v.OrgSubAccountUin }).(pulumi.IntOutput)
}

// Policy ID.
func (o PolicySubAccountAttachmentOutput) PolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.IntOutput { return v.PolicyId }).(pulumi.IntOutput)
}

// Policy name.
func (o PolicySubAccountAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// Update time.
func (o PolicySubAccountAttachmentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySubAccountAttachment) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type PolicySubAccountAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PolicySubAccountAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicySubAccountAttachment)(nil)).Elem()
}

func (o PolicySubAccountAttachmentArrayOutput) ToPolicySubAccountAttachmentArrayOutput() PolicySubAccountAttachmentArrayOutput {
	return o
}

func (o PolicySubAccountAttachmentArrayOutput) ToPolicySubAccountAttachmentArrayOutputWithContext(ctx context.Context) PolicySubAccountAttachmentArrayOutput {
	return o
}

func (o PolicySubAccountAttachmentArrayOutput) Index(i pulumi.IntInput) PolicySubAccountAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicySubAccountAttachment {
		return vs[0].([]*PolicySubAccountAttachment)[vs[1].(int)]
	}).(PolicySubAccountAttachmentOutput)
}

type PolicySubAccountAttachmentMapOutput struct{ *pulumi.OutputState }

func (PolicySubAccountAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicySubAccountAttachment)(nil)).Elem()
}

func (o PolicySubAccountAttachmentMapOutput) ToPolicySubAccountAttachmentMapOutput() PolicySubAccountAttachmentMapOutput {
	return o
}

func (o PolicySubAccountAttachmentMapOutput) ToPolicySubAccountAttachmentMapOutputWithContext(ctx context.Context) PolicySubAccountAttachmentMapOutput {
	return o
}

func (o PolicySubAccountAttachmentMapOutput) MapIndex(k pulumi.StringInput) PolicySubAccountAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicySubAccountAttachment {
		return vs[0].(map[string]*PolicySubAccountAttachment)[vs[1].(string)]
	}).(PolicySubAccountAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySubAccountAttachmentInput)(nil)).Elem(), &PolicySubAccountAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySubAccountAttachmentArrayInput)(nil)).Elem(), PolicySubAccountAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySubAccountAttachmentMapInput)(nil)).Elem(), PolicySubAccountAttachmentMap{})
	pulumi.RegisterOutputType(PolicySubAccountAttachmentOutput{})
	pulumi.RegisterOutputType(PolicySubAccountAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PolicySubAccountAttachmentMapOutput{})
}
