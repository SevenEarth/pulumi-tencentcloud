// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ckafka aclRule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ckafka.NewAclRule(ctx, "aclRule", &Ckafka.AclRuleArgs{
// 			InstanceId:   pulumi.String("ckafka-xxx"),
// 			IsApplied:    pulumi.Int(1),
// 			Pattern:      pulumi.String("prefix"),
// 			PatternType:  pulumi.String("PREFIXED"),
// 			ResourceType: pulumi.String("Topic"),
// 			RuleLists: ckafka.AclRuleRuleListArray{
// 				&ckafka.AclRuleRuleListArgs{
// 					Host:           pulumi.String("*"),
// 					Operation:      pulumi.String("All"),
// 					PermissionType: pulumi.String("Deny"),
// 					Principal:      pulumi.String("User:*"),
// 				},
// 			},
// 			RuleName: pulumi.String("RuleName"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ckafka acl_rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ckafka/aclRule:AclRule acl_rule acl_rule_id
// ```
type AclRule struct {
	pulumi.CustomResourceState

	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whether the preset ACL rule is applied to the newly added topic.
	IsApplied pulumi.IntPtrOutput `pulumi:"isApplied"`
	// A value representing the prefix that the prefix matches.
	Pattern pulumi.StringPtrOutput `pulumi:"pattern"`
	// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
	PatternType pulumi.StringOutput `pulumi:"patternType"`
	// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// List of configured ACL rules.
	RuleLists AclRuleRuleListArrayOutput `pulumi:"ruleLists"`
	// rule name.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
}

// NewAclRule registers a new resource with the given unique name, arguments, and options.
func NewAclRule(ctx *pulumi.Context,
	name string, args *AclRuleArgs, opts ...pulumi.ResourceOption) (*AclRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PatternType == nil {
		return nil, errors.New("invalid value for required argument 'PatternType'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.RuleLists == nil {
		return nil, errors.New("invalid value for required argument 'RuleLists'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AclRule
	err := ctx.RegisterResource("tencentcloud:Ckafka/aclRule:AclRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclRule gets an existing AclRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclRuleState, opts ...pulumi.ResourceOption) (*AclRule, error) {
	var resource AclRule
	err := ctx.ReadResource("tencentcloud:Ckafka/aclRule:AclRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclRule resources.
type aclRuleState struct {
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Whether the preset ACL rule is applied to the newly added topic.
	IsApplied *int `pulumi:"isApplied"`
	// A value representing the prefix that the prefix matches.
	Pattern *string `pulumi:"pattern"`
	// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
	PatternType *string `pulumi:"patternType"`
	// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
	ResourceType *string `pulumi:"resourceType"`
	// List of configured ACL rules.
	RuleLists []AclRuleRuleList `pulumi:"ruleLists"`
	// rule name.
	RuleName *string `pulumi:"ruleName"`
}

type AclRuleState struct {
	// instance id.
	InstanceId pulumi.StringPtrInput
	// Whether the preset ACL rule is applied to the newly added topic.
	IsApplied pulumi.IntPtrInput
	// A value representing the prefix that the prefix matches.
	Pattern pulumi.StringPtrInput
	// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
	PatternType pulumi.StringPtrInput
	// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
	ResourceType pulumi.StringPtrInput
	// List of configured ACL rules.
	RuleLists AclRuleRuleListArrayInput
	// rule name.
	RuleName pulumi.StringPtrInput
}

func (AclRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleState)(nil)).Elem()
}

type aclRuleArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Whether the preset ACL rule is applied to the newly added topic.
	IsApplied *int `pulumi:"isApplied"`
	// A value representing the prefix that the prefix matches.
	Pattern *string `pulumi:"pattern"`
	// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
	PatternType string `pulumi:"patternType"`
	// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
	ResourceType string `pulumi:"resourceType"`
	// List of configured ACL rules.
	RuleLists []AclRuleRuleList `pulumi:"ruleLists"`
	// rule name.
	RuleName string `pulumi:"ruleName"`
}

// The set of arguments for constructing a AclRule resource.
type AclRuleArgs struct {
	// instance id.
	InstanceId pulumi.StringInput
	// Whether the preset ACL rule is applied to the newly added topic.
	IsApplied pulumi.IntPtrInput
	// A value representing the prefix that the prefix matches.
	Pattern pulumi.StringPtrInput
	// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
	PatternType pulumi.StringInput
	// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
	ResourceType pulumi.StringInput
	// List of configured ACL rules.
	RuleLists AclRuleRuleListArrayInput
	// rule name.
	RuleName pulumi.StringInput
}

func (AclRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleArgs)(nil)).Elem()
}

type AclRuleInput interface {
	pulumi.Input

	ToAclRuleOutput() AclRuleOutput
	ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput
}

func (*AclRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (i *AclRule) ToAclRuleOutput() AclRuleOutput {
	return i.ToAclRuleOutputWithContext(context.Background())
}

func (i *AclRule) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleOutput)
}

// AclRuleArrayInput is an input type that accepts AclRuleArray and AclRuleArrayOutput values.
// You can construct a concrete instance of `AclRuleArrayInput` via:
//
//          AclRuleArray{ AclRuleArgs{...} }
type AclRuleArrayInput interface {
	pulumi.Input

	ToAclRuleArrayOutput() AclRuleArrayOutput
	ToAclRuleArrayOutputWithContext(context.Context) AclRuleArrayOutput
}

type AclRuleArray []AclRuleInput

func (AclRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (i AclRuleArray) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return i.ToAclRuleArrayOutputWithContext(context.Background())
}

func (i AclRuleArray) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleArrayOutput)
}

// AclRuleMapInput is an input type that accepts AclRuleMap and AclRuleMapOutput values.
// You can construct a concrete instance of `AclRuleMapInput` via:
//
//          AclRuleMap{ "key": AclRuleArgs{...} }
type AclRuleMapInput interface {
	pulumi.Input

	ToAclRuleMapOutput() AclRuleMapOutput
	ToAclRuleMapOutputWithContext(context.Context) AclRuleMapOutput
}

type AclRuleMap map[string]AclRuleInput

func (AclRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (i AclRuleMap) ToAclRuleMapOutput() AclRuleMapOutput {
	return i.ToAclRuleMapOutputWithContext(context.Background())
}

func (i AclRuleMap) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleMapOutput)
}

type AclRuleOutput struct{ *pulumi.OutputState }

func (AclRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (o AclRuleOutput) ToAclRuleOutput() AclRuleOutput {
	return o
}

func (o AclRuleOutput) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return o
}

// instance id.
func (o AclRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the preset ACL rule is applied to the newly added topic.
func (o AclRuleOutput) IsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.IntPtrOutput { return v.IsApplied }).(pulumi.IntPtrOutput)
}

// A value representing the prefix that the prefix matches.
func (o AclRuleOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

// Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
func (o AclRuleOutput) PatternType() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.PatternType }).(pulumi.StringOutput)
}

// Acl resource type, currently only supports Topic, enumeration value list{Topic}.
func (o AclRuleOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// List of configured ACL rules.
func (o AclRuleOutput) RuleLists() AclRuleRuleListArrayOutput {
	return o.ApplyT(func(v *AclRule) AclRuleRuleListArrayOutput { return v.RuleLists }).(AclRuleRuleListArrayOutput)
}

// rule name.
func (o AclRuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

type AclRuleArrayOutput struct{ *pulumi.OutputState }

func (AclRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) Index(i pulumi.IntInput) AclRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].([]*AclRule)[vs[1].(int)]
	}).(AclRuleOutput)
}

type AclRuleMapOutput struct{ *pulumi.OutputState }

func (AclRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (o AclRuleMapOutput) ToAclRuleMapOutput() AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) MapIndex(k pulumi.StringInput) AclRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].(map[string]*AclRule)[vs[1].(string)]
	}).(AclRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleInput)(nil)).Elem(), &AclRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleArrayInput)(nil)).Elem(), AclRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleMapInput)(nil)).Elem(), AclRuleMap{})
	pulumi.RegisterOutputType(AclRuleOutput{})
	pulumi.RegisterOutputType(AclRuleArrayOutput{})
	pulumi.RegisterOutputType(AclRuleMapOutput{})
}
