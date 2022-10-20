// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create a dayu CC self-define http policy
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
//			_, err := Dayu.NewCcHttpPolicy(ctx, "testBgpip", &Dayu.CcHttpPolicyArgs{
//				Action:       pulumi.String("drop"),
//				ResourceId:   pulumi.String("bgpip-00000294"),
//				ResourceType: pulumi.String("bgpip"),
//				RuleLists: dayu.CcHttpPolicyRuleListArray{
//					&dayu.CcHttpPolicyRuleListArgs{
//						Operator: pulumi.String("include"),
//						Skey:     pulumi.String("host"),
//						Value:    pulumi.String("123"),
//					},
//				},
//				Smode:  pulumi.String("matching"),
//				Switch: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dayu.NewCcHttpPolicy(ctx, "testNet", &Dayu.CcHttpPolicyArgs{
//				Action:       pulumi.String("drop"),
//				ResourceId:   pulumi.String("net-0000007e"),
//				ResourceType: pulumi.String("net"),
//				RuleLists: dayu.CcHttpPolicyRuleListArray{
//					&dayu.CcHttpPolicyRuleListArgs{
//						Operator: pulumi.String("equal"),
//						Skey:     pulumi.String("cgi"),
//						Value:    pulumi.String("123"),
//					},
//				},
//				Smode:  pulumi.String("matching"),
//				Switch: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dayu.NewCcHttpPolicy(ctx, "testBgpmultip", &Dayu.CcHttpPolicyArgs{
//				Action:       pulumi.String("alg"),
//				Ip:           pulumi.String("111.230.178.25"),
//				ResourceId:   pulumi.String("bgp-0000008o"),
//				ResourceType: pulumi.String("bgp-multip"),
//				RuleLists: dayu.CcHttpPolicyRuleListArray{
//					&dayu.CcHttpPolicyRuleListArgs{
//						Operator: pulumi.String("not_include"),
//						Skey:     pulumi.String("referer"),
//						Value:    pulumi.String("123"),
//					},
//				},
//				Smode:  pulumi.String("matching"),
//				Switch: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dayu.NewCcHttpPolicy(ctx, "testBgp", &Dayu.CcHttpPolicyArgs{
//				Action:       pulumi.String("alg"),
//				ResourceId:   pulumi.String("bgp-000006mq"),
//				ResourceType: pulumi.String("bgp"),
//				RuleLists: dayu.CcHttpPolicyRuleListArray{
//					&dayu.CcHttpPolicyRuleListArgs{
//						Operator: pulumi.String("not_include"),
//						Skey:     pulumi.String("ua"),
//						Value:    pulumi.String("123"),
//					},
//				},
//				Smode:  pulumi.String("matching"),
//				Switch: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CcHttpPolicy struct {
	pulumi.CustomResourceState

	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Create time of the CC self-define http policy.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	Frequency pulumi.IntOutput `pulumi:"frequency"`
	// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Name of the CC self-define http policy. Length should between 1 and 20.
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the CC self-define http policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// ID of the resource that the CC self-define http policy works for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	RuleLists CcHttpPolicyRuleListArrayOutput `pulumi:"ruleLists"`
	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	Smode pulumi.StringPtrOutput `pulumi:"smode"`
	// Indicate the CC self-define http policy takes effect or not.
	Switch pulumi.BoolPtrOutput `pulumi:"switch"`
}

// NewCcHttpPolicy registers a new resource with the given unique name, arguments, and options.
func NewCcHttpPolicy(ctx *pulumi.Context,
	name string, args *CcHttpPolicyArgs, opts ...pulumi.ResourceOption) (*CcHttpPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CcHttpPolicy
	err := ctx.RegisterResource("tencentcloud:Dayu/ccHttpPolicy:CcHttpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCcHttpPolicy gets an existing CcHttpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCcHttpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CcHttpPolicyState, opts ...pulumi.ResourceOption) (*CcHttpPolicy, error) {
	var resource CcHttpPolicy
	err := ctx.ReadResource("tencentcloud:Dayu/ccHttpPolicy:CcHttpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CcHttpPolicy resources.
type ccHttpPolicyState struct {
	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	Action *string `pulumi:"action"`
	// Create time of the CC self-define http policy.
	CreateTime *string `pulumi:"createTime"`
	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	Frequency *int `pulumi:"frequency"`
	// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
	Ip *string `pulumi:"ip"`
	// Name of the CC self-define http policy. Length should between 1 and 20.
	Name *string `pulumi:"name"`
	// Id of the CC self-define http policy.
	PolicyId *string `pulumi:"policyId"`
	// ID of the resource that the CC self-define http policy works for.
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType *string `pulumi:"resourceType"`
	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	RuleLists []CcHttpPolicyRuleList `pulumi:"ruleLists"`
	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	Smode *string `pulumi:"smode"`
	// Indicate the CC self-define http policy takes effect or not.
	Switch *bool `pulumi:"switch"`
}

type CcHttpPolicyState struct {
	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	Action pulumi.StringPtrInput
	// Create time of the CC self-define http policy.
	CreateTime pulumi.StringPtrInput
	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	Frequency pulumi.IntPtrInput
	// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
	Ip pulumi.StringPtrInput
	// Name of the CC self-define http policy. Length should between 1 and 20.
	Name pulumi.StringPtrInput
	// Id of the CC self-define http policy.
	PolicyId pulumi.StringPtrInput
	// ID of the resource that the CC self-define http policy works for.
	ResourceId pulumi.StringPtrInput
	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringPtrInput
	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	RuleLists CcHttpPolicyRuleListArrayInput
	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	Smode pulumi.StringPtrInput
	// Indicate the CC self-define http policy takes effect or not.
	Switch pulumi.BoolPtrInput
}

func (CcHttpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ccHttpPolicyState)(nil)).Elem()
}

type ccHttpPolicyArgs struct {
	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	Action *string `pulumi:"action"`
	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	Frequency *int `pulumi:"frequency"`
	// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
	Ip *string `pulumi:"ip"`
	// Name of the CC self-define http policy. Length should between 1 and 20.
	Name *string `pulumi:"name"`
	// ID of the resource that the CC self-define http policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType string `pulumi:"resourceType"`
	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	RuleLists []CcHttpPolicyRuleList `pulumi:"ruleLists"`
	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	Smode *string `pulumi:"smode"`
	// Indicate the CC self-define http policy takes effect or not.
	Switch *bool `pulumi:"switch"`
}

// The set of arguments for constructing a CcHttpPolicy resource.
type CcHttpPolicyArgs struct {
	// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
	Action pulumi.StringPtrInput
	// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
	Frequency pulumi.IntPtrInput
	// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
	Ip pulumi.StringPtrInput
	// Name of the CC self-define http policy. Length should between 1 and 20.
	Name pulumi.StringPtrInput
	// ID of the resource that the CC self-define http policy works for.
	ResourceId pulumi.StringInput
	// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringInput
	// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
	RuleLists CcHttpPolicyRuleListArrayInput
	// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
	Smode pulumi.StringPtrInput
	// Indicate the CC self-define http policy takes effect or not.
	Switch pulumi.BoolPtrInput
}

func (CcHttpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ccHttpPolicyArgs)(nil)).Elem()
}

type CcHttpPolicyInput interface {
	pulumi.Input

	ToCcHttpPolicyOutput() CcHttpPolicyOutput
	ToCcHttpPolicyOutputWithContext(ctx context.Context) CcHttpPolicyOutput
}

func (*CcHttpPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CcHttpPolicy)(nil)).Elem()
}

func (i *CcHttpPolicy) ToCcHttpPolicyOutput() CcHttpPolicyOutput {
	return i.ToCcHttpPolicyOutputWithContext(context.Background())
}

func (i *CcHttpPolicy) ToCcHttpPolicyOutputWithContext(ctx context.Context) CcHttpPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpPolicyOutput)
}

// CcHttpPolicyArrayInput is an input type that accepts CcHttpPolicyArray and CcHttpPolicyArrayOutput values.
// You can construct a concrete instance of `CcHttpPolicyArrayInput` via:
//
//	CcHttpPolicyArray{ CcHttpPolicyArgs{...} }
type CcHttpPolicyArrayInput interface {
	pulumi.Input

	ToCcHttpPolicyArrayOutput() CcHttpPolicyArrayOutput
	ToCcHttpPolicyArrayOutputWithContext(context.Context) CcHttpPolicyArrayOutput
}

type CcHttpPolicyArray []CcHttpPolicyInput

func (CcHttpPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcHttpPolicy)(nil)).Elem()
}

func (i CcHttpPolicyArray) ToCcHttpPolicyArrayOutput() CcHttpPolicyArrayOutput {
	return i.ToCcHttpPolicyArrayOutputWithContext(context.Background())
}

func (i CcHttpPolicyArray) ToCcHttpPolicyArrayOutputWithContext(ctx context.Context) CcHttpPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpPolicyArrayOutput)
}

// CcHttpPolicyMapInput is an input type that accepts CcHttpPolicyMap and CcHttpPolicyMapOutput values.
// You can construct a concrete instance of `CcHttpPolicyMapInput` via:
//
//	CcHttpPolicyMap{ "key": CcHttpPolicyArgs{...} }
type CcHttpPolicyMapInput interface {
	pulumi.Input

	ToCcHttpPolicyMapOutput() CcHttpPolicyMapOutput
	ToCcHttpPolicyMapOutputWithContext(context.Context) CcHttpPolicyMapOutput
}

type CcHttpPolicyMap map[string]CcHttpPolicyInput

func (CcHttpPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcHttpPolicy)(nil)).Elem()
}

func (i CcHttpPolicyMap) ToCcHttpPolicyMapOutput() CcHttpPolicyMapOutput {
	return i.ToCcHttpPolicyMapOutputWithContext(context.Background())
}

func (i CcHttpPolicyMap) ToCcHttpPolicyMapOutputWithContext(ctx context.Context) CcHttpPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CcHttpPolicyMapOutput)
}

type CcHttpPolicyOutput struct{ *pulumi.OutputState }

func (CcHttpPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CcHttpPolicy)(nil)).Elem()
}

func (o CcHttpPolicyOutput) ToCcHttpPolicyOutput() CcHttpPolicyOutput {
	return o
}

func (o CcHttpPolicyOutput) ToCcHttpPolicyOutputWithContext(ctx context.Context) CcHttpPolicyOutput {
	return o
}

// Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
func (o CcHttpPolicyOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Create time of the CC self-define http policy.
func (o CcHttpPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
func (o CcHttpPolicyOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// Ip of the CC self-define http policy, only valid when `resourceType` is `bgp-multip`. The num of list items can only be set one.
func (o CcHttpPolicyOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Name of the CC self-define http policy. Length should between 1 and 20.
func (o CcHttpPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Id of the CC self-define http policy.
func (o CcHttpPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// ID of the resource that the CC self-define http policy works for.
func (o CcHttpPolicyOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
func (o CcHttpPolicyOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Rule list of the CC self-define http policy,  only valid when `smode` is `matching`.
func (o CcHttpPolicyOutput) RuleLists() CcHttpPolicyRuleListArrayOutput {
	return o.ApplyT(func(v *CcHttpPolicy) CcHttpPolicyRuleListArrayOutput { return v.RuleLists }).(CcHttpPolicyRuleListArrayOutput)
}

// Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set one.
func (o CcHttpPolicyOutput) Smode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.StringPtrOutput { return v.Smode }).(pulumi.StringPtrOutput)
}

// Indicate the CC self-define http policy takes effect or not.
func (o CcHttpPolicyOutput) Switch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CcHttpPolicy) pulumi.BoolPtrOutput { return v.Switch }).(pulumi.BoolPtrOutput)
}

type CcHttpPolicyArrayOutput struct{ *pulumi.OutputState }

func (CcHttpPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CcHttpPolicy)(nil)).Elem()
}

func (o CcHttpPolicyArrayOutput) ToCcHttpPolicyArrayOutput() CcHttpPolicyArrayOutput {
	return o
}

func (o CcHttpPolicyArrayOutput) ToCcHttpPolicyArrayOutputWithContext(ctx context.Context) CcHttpPolicyArrayOutput {
	return o
}

func (o CcHttpPolicyArrayOutput) Index(i pulumi.IntInput) CcHttpPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CcHttpPolicy {
		return vs[0].([]*CcHttpPolicy)[vs[1].(int)]
	}).(CcHttpPolicyOutput)
}

type CcHttpPolicyMapOutput struct{ *pulumi.OutputState }

func (CcHttpPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CcHttpPolicy)(nil)).Elem()
}

func (o CcHttpPolicyMapOutput) ToCcHttpPolicyMapOutput() CcHttpPolicyMapOutput {
	return o
}

func (o CcHttpPolicyMapOutput) ToCcHttpPolicyMapOutputWithContext(ctx context.Context) CcHttpPolicyMapOutput {
	return o
}

func (o CcHttpPolicyMapOutput) MapIndex(k pulumi.StringInput) CcHttpPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CcHttpPolicy {
		return vs[0].(map[string]*CcHttpPolicy)[vs[1].(string)]
	}).(CcHttpPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpPolicyInput)(nil)).Elem(), &CcHttpPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpPolicyArrayInput)(nil)).Elem(), CcHttpPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CcHttpPolicyMapInput)(nil)).Elem(), CcHttpPolicyMap{})
	pulumi.RegisterOutputType(CcHttpPolicyOutput{})
	pulumi.RegisterOutputType(CcHttpPolicyArrayOutput{})
	pulumi.RegisterOutputType(CcHttpPolicyMapOutput{})
}
