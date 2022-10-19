// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create a dayu CC self-define https policy
//
// > **NOTE:** creating CC self-define https policy need a valid resource `Dayu.L7Rule`; The resource only support Anti-DDoS of resource type `bgpip`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dayu.NewCcHttpsPolicy(ctx, "testPolicy", &Dayu.CcHttpsPolicyArgs{
//				ResourceType: pulumi.Any(tencentcloud_dayu_l7_rule.Test_rule.Resource_type),
//				ResourceId:   pulumi.Any(tencentcloud_dayu_l7_rule.Test_rule.Resource_id),
//				RuleId:       pulumi.Any(tencentcloud_dayu_l7_rule.Test_rule.Rule_id),
//				Domain:       pulumi.Any(tencentcloud_dayu_l7_rule.Test_rule.Domain),
//				Action:       pulumi.String("drop"),
//				Switch:       pulumi.Bool(true),
//				RuleLists: dayu.CcHttpsPolicyRuleListArray{
//					&dayu.CcHttpsPolicyRuleListArgs{
//						Skey:     pulumi.String("cgi"),
//						Operator: pulumi.String("include"),
//						Value:    pulumi.String("123"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CcHttpsPolicy struct {
	pulumi.CustomResourceState

	// Action mode. Valid values are `alg` and `drop`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Create time of the CC self-define https policy.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Ip of the CC self-define https policy.
	IpLists pulumi.StringArrayOutput `pulumi:"ipLists"`
	// Name of the CC self-define https policy. Length should between 1 and 20.
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the CC self-define https policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// ID of the resource that the CC self-define https policy works for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Rule list of the CC self-define https policy.
	RuleLists CcHttpsPolicyRuleListArrayOutput `pulumi:"ruleLists"`
	// Indicate the CC self-define https policy takes effect or not.
	Switch pulumi.BoolPtrOutput `pulumi:"switch"`
}

// NewCcHttpsPolicy registers a new resource with the given unique name, arguments, and options.
func NewCcHttpsPolicy(ctx *pulumi.Context,
	name string, args *CcHttpsPolicyArgs, opts ...pulumi.ResourceOption) (*CcHttpsPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.RuleLists == nil {
		return nil, errors.New("invalid value for required argument 'RuleLists'")
	}
	var resource CcHttpsPolicy
	err := ctx.RegisterResource("tencentcloud:Dayu/ccHttpsPolicy:CcHttpsPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCcHttpsPolicy gets an existing CcHttpsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCcHttpsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CcHttpsPolicyState, opts ...pulumi.ResourceOption) (*CcHttpsPolicy, error) {
	var resource CcHttpsPolicy
	err := ctx.ReadResource("tencentcloud:Dayu/ccHttpsPolicy:CcHttpsPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CcHttpsPolicy resources.
type ccHttpsPolicyState struct {
	// Action mode. Valid values are `alg` and `drop`.
	Action *string `pulumi:"action"`
	// Create time of the CC self-define https policy.
	CreateTime *string `pulumi:"createTime"`
	// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	Domain *string `pulumi:"domain"`
	// Ip of the CC self-define https policy.
	IpLists []string `pulumi:"ipLists"`
	// Name of the CC self-define https policy. Length should between 1 and 20.
	Name *string `pulumi:"name"`
	// Id of the CC self-define https policy.
	PolicyId *string `pulumi:"policyId"`
	// ID of the resource that the CC self-define https policy works for.
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
	ResourceType *string `pulumi:"resourceType"`
	// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	RuleId *string `pulumi:"ruleId"`
	// Rule list of the CC self-define https policy.
	RuleLists []CcHttpsPolicyRuleList `pulumi:"ruleLists"`
	// Indicate the CC self-define https policy takes effect or not.
	Switch *bool `pulumi:"switch"`
}

type CcHttpsPolicyState struct {
	// Action mode. Valid values are `alg` and `drop`.
	Action pulumi.StringPtrInput
	// Create time of the CC self-define https policy.
	CreateTime pulumi.StringPtrInput
	// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	Domain pulumi.StringPtrInput
	// Ip of the CC self-define https policy.
	IpLists pulumi.StringArrayInput
	// Name of the CC self-define https policy. Length should between 1 and 20.
	Name pulumi.StringPtrInput
	// Id of the CC self-define https policy.
	PolicyId pulumi.StringPtrInput
	// ID of the resource that the CC self-define https policy works for.
	ResourceId pulumi.StringPtrInput
	// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
	ResourceType pulumi.StringPtrInput
	// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	RuleId pulumi.StringPtrInput
	// Rule list of the CC self-define https policy.
	RuleLists CcHttpsPolicyRuleListArrayInput
	// Indicate the CC self-define https policy takes effect or not.
	Switch pulumi.BoolPtrInput
}

func (CcHttpsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ccHttpsPolicyState)(nil)).Elem()
}

type ccHttpsPolicyArgs struct {
	// Action mode. Valid values are `alg` and `drop`.
	Action *string `pulumi:"action"`
	// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	Domain string `pulumi:"domain"`
	// Name of the CC self-define https policy. Length should between 1 and 20.
	Name *string `pulumi:"name"`
	// ID of the resource that the CC self-define https policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
	ResourceType string `pulumi:"resourceType"`
	// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	RuleId string `pulumi:"ruleId"`
	// Rule list of the CC self-define https policy.
	RuleLists []CcHttpsPolicyRuleList `pulumi:"ruleLists"`
	// Indicate the CC self-define https policy takes effect or not.
	Switch *bool `pulumi:"switch"`
}

// The set of arguments for constructing a CcHttpsPolicy resource.
type CcHttpsPolicyArgs struct {
	// Action mode. Valid values are `alg` and `drop`.
	Action pulumi.StringPtrInput
	// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	Domain pulumi.StringInput
	// Name of the CC self-define https policy. Length should between 1 and 20.
	Name pulumi.StringPtrInput
	// ID of the resource that the CC self-define https policy works for.
	ResourceId pulumi.StringInput
	// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
	ResourceType pulumi.StringInput
	// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
	RuleId pulumi.StringInput
	// Rule list of the CC self-define https policy.
	RuleLists CcHttpsPolicyRuleListArrayInput
	// Indicate the CC self-define https policy takes effect or not.
	Switch pulumi.BoolPtrInput
}

func (CcHttpsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ccHttpsPolicyArgs)(nil)).Elem()
}

type CcHttpsPolicyInput interface {
	pulumi.Input

	ToCcHttpsPolicyOutput() CcHttpsPolicyOutput
	ToCcHttpsPolicyOutputWithContext(ctx context.Context) CcHttpsPolicyOutput
}

func (*CcHttpsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CcHttpsPolicy)(nil)).Elem()
}

func (i *CcHttpsPolicy) ToCcHttpsPolicyOutput() CcHttpsPolicyOutput {
	return i.ToCcHttpsPolicyOutputWithContext(context.Background())
}

func (i *CcHttpsPolicy) ToCcHttpsPolicyOutputWithContext(ctx context.Context) CcHttpsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpsPolicyOutput)
}

// CcHttpsPolicyArrayInput is an input type that accepts CcHttpsPolicyArray and CcHttpsPolicyArrayOutput values.
// You can construct a concrete instance of `CcHttpsPolicyArrayInput` via:
//
//	CcHttpsPolicyArray{ CcHttpsPolicyArgs{...} }
type CcHttpsPolicyArrayInput interface {
	pulumi.Input

	ToCcHttpsPolicyArrayOutput() CcHttpsPolicyArrayOutput
	ToCcHttpsPolicyArrayOutputWithContext(context.Context) CcHttpsPolicyArrayOutput
}

type CcHttpsPolicyArray []CcHttpsPolicyInput

func (CcHttpsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcHttpsPolicy)(nil)).Elem()
}

func (i CcHttpsPolicyArray) ToCcHttpsPolicyArrayOutput() CcHttpsPolicyArrayOutput {
	return i.ToCcHttpsPolicyArrayOutputWithContext(context.Background())
}

func (i CcHttpsPolicyArray) ToCcHttpsPolicyArrayOutputWithContext(ctx context.Context) CcHttpsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpsPolicyArrayOutput)
}

// CcHttpsPolicyMapInput is an input type that accepts CcHttpsPolicyMap and CcHttpsPolicyMapOutput values.
// You can construct a concrete instance of `CcHttpsPolicyMapInput` via:
//
//	CcHttpsPolicyMap{ "key": CcHttpsPolicyArgs{...} }
type CcHttpsPolicyMapInput interface {
	pulumi.Input

	ToCcHttpsPolicyMapOutput() CcHttpsPolicyMapOutput
	ToCcHttpsPolicyMapOutputWithContext(context.Context) CcHttpsPolicyMapOutput
}

type CcHttpsPolicyMap map[string]CcHttpsPolicyInput

func (CcHttpsPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcHttpsPolicy)(nil)).Elem()
}

func (i CcHttpsPolicyMap) ToCcHttpsPolicyMapOutput() CcHttpsPolicyMapOutput {
	return i.ToCcHttpsPolicyMapOutputWithContext(context.Background())
}

func (i CcHttpsPolicyMap) ToCcHttpsPolicyMapOutputWithContext(ctx context.Context) CcHttpsPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpsPolicyMapOutput)
}

type CcHttpsPolicyOutput struct{ *pulumi.OutputState }

func (CcHttpsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CcHttpsPolicy)(nil)).Elem()
}

func (o CcHttpsPolicyOutput) ToCcHttpsPolicyOutput() CcHttpsPolicyOutput {
	return o
}

func (o CcHttpsPolicyOutput) ToCcHttpsPolicyOutputWithContext(ctx context.Context) CcHttpsPolicyOutput {
	return o
}

// Action mode. Valid values are `alg` and `drop`.
func (o CcHttpsPolicyOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Create time of the CC self-define https policy.
func (o CcHttpsPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
func (o CcHttpsPolicyOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Ip of the CC self-define https policy.
func (o CcHttpsPolicyOutput) IpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringArrayOutput { return v.IpLists }).(pulumi.StringArrayOutput)
}

// Name of the CC self-define https policy. Length should between 1 and 20.
func (o CcHttpsPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Id of the CC self-define https policy.
func (o CcHttpsPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// ID of the resource that the CC self-define https policy works for.
func (o CcHttpsPolicyOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
func (o CcHttpsPolicyOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
func (o CcHttpsPolicyOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Rule list of the CC self-define https policy.
func (o CcHttpsPolicyOutput) RuleLists() CcHttpsPolicyRuleListArrayOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) CcHttpsPolicyRuleListArrayOutput { return v.RuleLists }).(CcHttpsPolicyRuleListArrayOutput)
}

// Indicate the CC self-define https policy takes effect or not.
func (o CcHttpsPolicyOutput) Switch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CcHttpsPolicy) pulumi.BoolPtrOutput { return v.Switch }).(pulumi.BoolPtrOutput)
}

type CcHttpsPolicyArrayOutput struct{ *pulumi.OutputState }

func (CcHttpsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcHttpsPolicy)(nil)).Elem()
}

func (o CcHttpsPolicyArrayOutput) ToCcHttpsPolicyArrayOutput() CcHttpsPolicyArrayOutput {
	return o
}

func (o CcHttpsPolicyArrayOutput) ToCcHttpsPolicyArrayOutputWithContext(ctx context.Context) CcHttpsPolicyArrayOutput {
	return o
}

func (o CcHttpsPolicyArrayOutput) Index(i pulumi.IntInput) CcHttpsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CcHttpsPolicy {
		return vs[0].([]*CcHttpsPolicy)[vs[1].(int)]
	}).(CcHttpsPolicyOutput)
}

type CcHttpsPolicyMapOutput struct{ *pulumi.OutputState }

func (CcHttpsPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcHttpsPolicy)(nil)).Elem()
}

func (o CcHttpsPolicyMapOutput) ToCcHttpsPolicyMapOutput() CcHttpsPolicyMapOutput {
	return o
}

func (o CcHttpsPolicyMapOutput) ToCcHttpsPolicyMapOutputWithContext(ctx context.Context) CcHttpsPolicyMapOutput {
	return o
}

func (o CcHttpsPolicyMapOutput) MapIndex(k pulumi.StringInput) CcHttpsPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CcHttpsPolicy {
		return vs[0].(map[string]*CcHttpsPolicy)[vs[1].(string)]
	}).(CcHttpsPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpsPolicyInput)(nil)).Elem(), &CcHttpsPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpsPolicyArrayInput)(nil)).Elem(), CcHttpsPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpsPolicyMapInput)(nil)).Elem(), CcHttpsPolicyMap{})
	pulumi.RegisterOutputType(CcHttpsPolicyOutput{})
	pulumi.RegisterOutputType(CcHttpsPolicyArrayOutput{})
	pulumi.RegisterOutputType(CcHttpsPolicyMapOutput{})
}
