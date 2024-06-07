// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a organization orgManagePolicy
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Organization"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Organization.NewOrgManagePolicy(ctx, "orgManagePolicy", &Organization.OrgManagePolicyArgs{
//				Content:     pulumi.String("{\"version\":\"2.0\",\"statement\":[{\"effect\":\"allow\",\"action\":\"*\",\"resource\":\"*\"}]}"),
//				Description: pulumi.String("Full access policy"),
//				Type:        pulumi.String("SERVICE_CONTROL_POLICY"),
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
//
// ## Import
//
// organization org_manage_policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Organization/orgManagePolicy:OrgManagePolicy org_manage_policy policy_id#type
// ```
type OrgManagePolicy struct {
	pulumi.CustomResourceState

	// Policy content. Refer to the CAM policy syntax.
	Content pulumi.StringOutput `pulumi:"content"`
	// Policy description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Policy name.
	// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy Id.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Policy type. Default value is SERVICE_CONTROL_POLICY.
	// Valid values:
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewOrgManagePolicy registers a new resource with the given unique name, arguments, and options.
func NewOrgManagePolicy(ctx *pulumi.Context,
	name string, args *OrgManagePolicyArgs, opts ...pulumi.ResourceOption) (*OrgManagePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgManagePolicy
	err := ctx.RegisterResource("tencentcloud:Organization/orgManagePolicy:OrgManagePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgManagePolicy gets an existing OrgManagePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgManagePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgManagePolicyState, opts ...pulumi.ResourceOption) (*OrgManagePolicy, error) {
	var resource OrgManagePolicy
	err := ctx.ReadResource("tencentcloud:Organization/orgManagePolicy:OrgManagePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgManagePolicy resources.
type orgManagePolicyState struct {
	// Policy content. Refer to the CAM policy syntax.
	Content *string `pulumi:"content"`
	// Policy description.
	Description *string `pulumi:"description"`
	// Policy name.
	// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
	Name *string `pulumi:"name"`
	// Policy Id.
	PolicyId *string `pulumi:"policyId"`
	// Policy type. Default value is SERVICE_CONTROL_POLICY.
	// Valid values:
	Type *string `pulumi:"type"`
}

type OrgManagePolicyState struct {
	// Policy content. Refer to the CAM policy syntax.
	Content pulumi.StringPtrInput
	// Policy description.
	Description pulumi.StringPtrInput
	// Policy name.
	// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
	Name pulumi.StringPtrInput
	// Policy Id.
	PolicyId pulumi.StringPtrInput
	// Policy type. Default value is SERVICE_CONTROL_POLICY.
	// Valid values:
	Type pulumi.StringPtrInput
}

func (OrgManagePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgManagePolicyState)(nil)).Elem()
}

type orgManagePolicyArgs struct {
	// Policy content. Refer to the CAM policy syntax.
	Content string `pulumi:"content"`
	// Policy description.
	Description *string `pulumi:"description"`
	// Policy name.
	// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
	Name *string `pulumi:"name"`
	// Policy type. Default value is SERVICE_CONTROL_POLICY.
	// Valid values:
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OrgManagePolicy resource.
type OrgManagePolicyArgs struct {
	// Policy content. Refer to the CAM policy syntax.
	Content pulumi.StringInput
	// Policy description.
	Description pulumi.StringPtrInput
	// Policy name.
	// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
	Name pulumi.StringPtrInput
	// Policy type. Default value is SERVICE_CONTROL_POLICY.
	// Valid values:
	Type pulumi.StringPtrInput
}

func (OrgManagePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgManagePolicyArgs)(nil)).Elem()
}

type OrgManagePolicyInput interface {
	pulumi.Input

	ToOrgManagePolicyOutput() OrgManagePolicyOutput
	ToOrgManagePolicyOutputWithContext(ctx context.Context) OrgManagePolicyOutput
}

func (*OrgManagePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgManagePolicy)(nil)).Elem()
}

func (i *OrgManagePolicy) ToOrgManagePolicyOutput() OrgManagePolicyOutput {
	return i.ToOrgManagePolicyOutputWithContext(context.Background())
}

func (i *OrgManagePolicy) ToOrgManagePolicyOutputWithContext(ctx context.Context) OrgManagePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgManagePolicyOutput)
}

// OrgManagePolicyArrayInput is an input type that accepts OrgManagePolicyArray and OrgManagePolicyArrayOutput values.
// You can construct a concrete instance of `OrgManagePolicyArrayInput` via:
//
//	OrgManagePolicyArray{ OrgManagePolicyArgs{...} }
type OrgManagePolicyArrayInput interface {
	pulumi.Input

	ToOrgManagePolicyArrayOutput() OrgManagePolicyArrayOutput
	ToOrgManagePolicyArrayOutputWithContext(context.Context) OrgManagePolicyArrayOutput
}

type OrgManagePolicyArray []OrgManagePolicyInput

func (OrgManagePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgManagePolicy)(nil)).Elem()
}

func (i OrgManagePolicyArray) ToOrgManagePolicyArrayOutput() OrgManagePolicyArrayOutput {
	return i.ToOrgManagePolicyArrayOutputWithContext(context.Background())
}

func (i OrgManagePolicyArray) ToOrgManagePolicyArrayOutputWithContext(ctx context.Context) OrgManagePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgManagePolicyArrayOutput)
}

// OrgManagePolicyMapInput is an input type that accepts OrgManagePolicyMap and OrgManagePolicyMapOutput values.
// You can construct a concrete instance of `OrgManagePolicyMapInput` via:
//
//	OrgManagePolicyMap{ "key": OrgManagePolicyArgs{...} }
type OrgManagePolicyMapInput interface {
	pulumi.Input

	ToOrgManagePolicyMapOutput() OrgManagePolicyMapOutput
	ToOrgManagePolicyMapOutputWithContext(context.Context) OrgManagePolicyMapOutput
}

type OrgManagePolicyMap map[string]OrgManagePolicyInput

func (OrgManagePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgManagePolicy)(nil)).Elem()
}

func (i OrgManagePolicyMap) ToOrgManagePolicyMapOutput() OrgManagePolicyMapOutput {
	return i.ToOrgManagePolicyMapOutputWithContext(context.Background())
}

func (i OrgManagePolicyMap) ToOrgManagePolicyMapOutputWithContext(ctx context.Context) OrgManagePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgManagePolicyMapOutput)
}

type OrgManagePolicyOutput struct{ *pulumi.OutputState }

func (OrgManagePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgManagePolicy)(nil)).Elem()
}

func (o OrgManagePolicyOutput) ToOrgManagePolicyOutput() OrgManagePolicyOutput {
	return o
}

func (o OrgManagePolicyOutput) ToOrgManagePolicyOutputWithContext(ctx context.Context) OrgManagePolicyOutput {
	return o
}

// Policy content. Refer to the CAM policy syntax.
func (o OrgManagePolicyOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgManagePolicy) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Policy description.
func (o OrgManagePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgManagePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Policy name.
// The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
func (o OrgManagePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgManagePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy Id.
func (o OrgManagePolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgManagePolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// Policy type. Default value is SERVICE_CONTROL_POLICY.
// Valid values:
func (o OrgManagePolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgManagePolicy) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type OrgManagePolicyArrayOutput struct{ *pulumi.OutputState }

func (OrgManagePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgManagePolicy)(nil)).Elem()
}

func (o OrgManagePolicyArrayOutput) ToOrgManagePolicyArrayOutput() OrgManagePolicyArrayOutput {
	return o
}

func (o OrgManagePolicyArrayOutput) ToOrgManagePolicyArrayOutputWithContext(ctx context.Context) OrgManagePolicyArrayOutput {
	return o
}

func (o OrgManagePolicyArrayOutput) Index(i pulumi.IntInput) OrgManagePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgManagePolicy {
		return vs[0].([]*OrgManagePolicy)[vs[1].(int)]
	}).(OrgManagePolicyOutput)
}

type OrgManagePolicyMapOutput struct{ *pulumi.OutputState }

func (OrgManagePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgManagePolicy)(nil)).Elem()
}

func (o OrgManagePolicyMapOutput) ToOrgManagePolicyMapOutput() OrgManagePolicyMapOutput {
	return o
}

func (o OrgManagePolicyMapOutput) ToOrgManagePolicyMapOutputWithContext(ctx context.Context) OrgManagePolicyMapOutput {
	return o
}

func (o OrgManagePolicyMapOutput) MapIndex(k pulumi.StringInput) OrgManagePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgManagePolicy {
		return vs[0].(map[string]*OrgManagePolicy)[vs[1].(string)]
	}).(OrgManagePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgManagePolicyInput)(nil)).Elem(), &OrgManagePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgManagePolicyArrayInput)(nil)).Elem(), OrgManagePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgManagePolicyMapInput)(nil)).Elem(), OrgManagePolicyMap{})
	pulumi.RegisterOutputType(OrgManagePolicyOutput{})
	pulumi.RegisterOutputType(OrgManagePolicyArrayOutput{})
	pulumi.RegisterOutputType(OrgManagePolicyMapOutput{})
}
