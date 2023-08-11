// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcr tag retention rule.
//
// ## Example Usage
// ### Create a tcr tag retention rule instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleInstance, err := Tcr.NewInstance(ctx, "exampleInstance", &Tcr.InstanceArgs{
// 			InstanceType: pulumi.String("basic"),
// 			DeleteBucket: pulumi.Bool(true),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNamespace, err := Tcr.NewNamespace(ctx, "exampleNamespace", &Tcr.NamespaceArgs{
// 			InstanceId:   exampleInstance.ID(),
// 			IsPublic:     pulumi.Bool(true),
// 			IsAutoScan:   pulumi.Bool(true),
// 			IsPreventVul: pulumi.Bool(true),
// 			Severity:     pulumi.String("medium"),
// 			CveWhitelistItems: tcr.NamespaceCveWhitelistItemArray{
// 				&tcr.NamespaceCveWhitelistItemArgs{
// 					CveId: pulumi.String("cve-xxxxx"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tcr.NewTagRetentionRule(ctx, "myRule", &Tcr.TagRetentionRuleArgs{
// 			RegistryId:    exampleInstance.ID(),
// 			NamespaceName: exampleNamespace.Name,
// 			RetentionRule: &tcr.TagRetentionRuleRetentionRuleArgs{
// 				Key:   pulumi.String("nDaysSinceLastPush"),
// 				Value: pulumi.Int(2),
// 			},
// 			CronSetting: pulumi.String("daily"),
// 			Disabled:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TagRetentionRule struct {
	pulumi.CustomResourceState

	// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
	CronSetting pulumi.StringOutput `pulumi:"cronSetting"`
	// Whether to disable the rule, with the default value of false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The Name of the namespace.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The main instance ID.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The ID of the retention task.
	RetentionId pulumi.IntOutput `pulumi:"retentionId"`
	// Retention Policy.
	RetentionRule TagRetentionRuleRetentionRuleOutput `pulumi:"retentionRule"`
}

// NewTagRetentionRule registers a new resource with the given unique name, arguments, and options.
func NewTagRetentionRule(ctx *pulumi.Context,
	name string, args *TagRetentionRuleArgs, opts ...pulumi.ResourceOption) (*TagRetentionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronSetting == nil {
		return nil, errors.New("invalid value for required argument 'CronSetting'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.RetentionRule == nil {
		return nil, errors.New("invalid value for required argument 'RetentionRule'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TagRetentionRule
	err := ctx.RegisterResource("tencentcloud:Tcr/tagRetentionRule:TagRetentionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagRetentionRule gets an existing TagRetentionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagRetentionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagRetentionRuleState, opts ...pulumi.ResourceOption) (*TagRetentionRule, error) {
	var resource TagRetentionRule
	err := ctx.ReadResource("tencentcloud:Tcr/tagRetentionRule:TagRetentionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagRetentionRule resources.
type tagRetentionRuleState struct {
	// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
	CronSetting *string `pulumi:"cronSetting"`
	// Whether to disable the rule, with the default value of false.
	Disabled *bool `pulumi:"disabled"`
	// The Name of the namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	// The main instance ID.
	RegistryId *string `pulumi:"registryId"`
	// The ID of the retention task.
	RetentionId *int `pulumi:"retentionId"`
	// Retention Policy.
	RetentionRule *TagRetentionRuleRetentionRule `pulumi:"retentionRule"`
}

type TagRetentionRuleState struct {
	// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
	CronSetting pulumi.StringPtrInput
	// Whether to disable the rule, with the default value of false.
	Disabled pulumi.BoolPtrInput
	// The Name of the namespace.
	NamespaceName pulumi.StringPtrInput
	// The main instance ID.
	RegistryId pulumi.StringPtrInput
	// The ID of the retention task.
	RetentionId pulumi.IntPtrInput
	// Retention Policy.
	RetentionRule TagRetentionRuleRetentionRulePtrInput
}

func (TagRetentionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRetentionRuleState)(nil)).Elem()
}

type tagRetentionRuleArgs struct {
	// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
	CronSetting string `pulumi:"cronSetting"`
	// Whether to disable the rule, with the default value of false.
	Disabled *bool `pulumi:"disabled"`
	// The Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The main instance ID.
	RegistryId string `pulumi:"registryId"`
	// Retention Policy.
	RetentionRule TagRetentionRuleRetentionRule `pulumi:"retentionRule"`
}

// The set of arguments for constructing a TagRetentionRule resource.
type TagRetentionRuleArgs struct {
	// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
	CronSetting pulumi.StringInput
	// Whether to disable the rule, with the default value of false.
	Disabled pulumi.BoolPtrInput
	// The Name of the namespace.
	NamespaceName pulumi.StringInput
	// The main instance ID.
	RegistryId pulumi.StringInput
	// Retention Policy.
	RetentionRule TagRetentionRuleRetentionRuleInput
}

func (TagRetentionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagRetentionRuleArgs)(nil)).Elem()
}

type TagRetentionRuleInput interface {
	pulumi.Input

	ToTagRetentionRuleOutput() TagRetentionRuleOutput
	ToTagRetentionRuleOutputWithContext(ctx context.Context) TagRetentionRuleOutput
}

func (*TagRetentionRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TagRetentionRule)(nil)).Elem()
}

func (i *TagRetentionRule) ToTagRetentionRuleOutput() TagRetentionRuleOutput {
	return i.ToTagRetentionRuleOutputWithContext(context.Background())
}

func (i *TagRetentionRule) ToTagRetentionRuleOutputWithContext(ctx context.Context) TagRetentionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagRetentionRuleOutput)
}

// TagRetentionRuleArrayInput is an input type that accepts TagRetentionRuleArray and TagRetentionRuleArrayOutput values.
// You can construct a concrete instance of `TagRetentionRuleArrayInput` via:
//
//          TagRetentionRuleArray{ TagRetentionRuleArgs{...} }
type TagRetentionRuleArrayInput interface {
	pulumi.Input

	ToTagRetentionRuleArrayOutput() TagRetentionRuleArrayOutput
	ToTagRetentionRuleArrayOutputWithContext(context.Context) TagRetentionRuleArrayOutput
}

type TagRetentionRuleArray []TagRetentionRuleInput

func (TagRetentionRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagRetentionRule)(nil)).Elem()
}

func (i TagRetentionRuleArray) ToTagRetentionRuleArrayOutput() TagRetentionRuleArrayOutput {
	return i.ToTagRetentionRuleArrayOutputWithContext(context.Background())
}

func (i TagRetentionRuleArray) ToTagRetentionRuleArrayOutputWithContext(ctx context.Context) TagRetentionRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagRetentionRuleArrayOutput)
}

// TagRetentionRuleMapInput is an input type that accepts TagRetentionRuleMap and TagRetentionRuleMapOutput values.
// You can construct a concrete instance of `TagRetentionRuleMapInput` via:
//
//          TagRetentionRuleMap{ "key": TagRetentionRuleArgs{...} }
type TagRetentionRuleMapInput interface {
	pulumi.Input

	ToTagRetentionRuleMapOutput() TagRetentionRuleMapOutput
	ToTagRetentionRuleMapOutputWithContext(context.Context) TagRetentionRuleMapOutput
}

type TagRetentionRuleMap map[string]TagRetentionRuleInput

func (TagRetentionRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagRetentionRule)(nil)).Elem()
}

func (i TagRetentionRuleMap) ToTagRetentionRuleMapOutput() TagRetentionRuleMapOutput {
	return i.ToTagRetentionRuleMapOutputWithContext(context.Background())
}

func (i TagRetentionRuleMap) ToTagRetentionRuleMapOutputWithContext(ctx context.Context) TagRetentionRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagRetentionRuleMapOutput)
}

type TagRetentionRuleOutput struct{ *pulumi.OutputState }

func (TagRetentionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagRetentionRule)(nil)).Elem()
}

func (o TagRetentionRuleOutput) ToTagRetentionRuleOutput() TagRetentionRuleOutput {
	return o
}

func (o TagRetentionRuleOutput) ToTagRetentionRuleOutputWithContext(ctx context.Context) TagRetentionRuleOutput {
	return o
}

// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
func (o TagRetentionRuleOutput) CronSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *TagRetentionRule) pulumi.StringOutput { return v.CronSetting }).(pulumi.StringOutput)
}

// Whether to disable the rule, with the default value of false.
func (o TagRetentionRuleOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TagRetentionRule) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The Name of the namespace.
func (o TagRetentionRuleOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagRetentionRule) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// The main instance ID.
func (o TagRetentionRuleOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagRetentionRule) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The ID of the retention task.
func (o TagRetentionRuleOutput) RetentionId() pulumi.IntOutput {
	return o.ApplyT(func(v *TagRetentionRule) pulumi.IntOutput { return v.RetentionId }).(pulumi.IntOutput)
}

// Retention Policy.
func (o TagRetentionRuleOutput) RetentionRule() TagRetentionRuleRetentionRuleOutput {
	return o.ApplyT(func(v *TagRetentionRule) TagRetentionRuleRetentionRuleOutput { return v.RetentionRule }).(TagRetentionRuleRetentionRuleOutput)
}

type TagRetentionRuleArrayOutput struct{ *pulumi.OutputState }

func (TagRetentionRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagRetentionRule)(nil)).Elem()
}

func (o TagRetentionRuleArrayOutput) ToTagRetentionRuleArrayOutput() TagRetentionRuleArrayOutput {
	return o
}

func (o TagRetentionRuleArrayOutput) ToTagRetentionRuleArrayOutputWithContext(ctx context.Context) TagRetentionRuleArrayOutput {
	return o
}

func (o TagRetentionRuleArrayOutput) Index(i pulumi.IntInput) TagRetentionRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagRetentionRule {
		return vs[0].([]*TagRetentionRule)[vs[1].(int)]
	}).(TagRetentionRuleOutput)
}

type TagRetentionRuleMapOutput struct{ *pulumi.OutputState }

func (TagRetentionRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagRetentionRule)(nil)).Elem()
}

func (o TagRetentionRuleMapOutput) ToTagRetentionRuleMapOutput() TagRetentionRuleMapOutput {
	return o
}

func (o TagRetentionRuleMapOutput) ToTagRetentionRuleMapOutputWithContext(ctx context.Context) TagRetentionRuleMapOutput {
	return o
}

func (o TagRetentionRuleMapOutput) MapIndex(k pulumi.StringInput) TagRetentionRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagRetentionRule {
		return vs[0].(map[string]*TagRetentionRule)[vs[1].(string)]
	}).(TagRetentionRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagRetentionRuleInput)(nil)).Elem(), &TagRetentionRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagRetentionRuleArrayInput)(nil)).Elem(), TagRetentionRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagRetentionRuleMapInput)(nil)).Elem(), TagRetentionRuleMap{})
	pulumi.RegisterOutputType(TagRetentionRuleOutput{})
	pulumi.RegisterOutputType(TagRetentionRuleArrayOutput{})
	pulumi.RegisterOutputType(TagRetentionRuleMapOutput{})
}
