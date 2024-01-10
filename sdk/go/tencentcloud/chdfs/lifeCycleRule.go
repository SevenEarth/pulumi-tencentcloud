// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chdfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a chdfs lifeCycleRule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Chdfs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Chdfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Chdfs.NewLifeCycleRule(ctx, "lifeCycleRule", &Chdfs.LifeCycleRuleArgs{
//				FileSystemId: pulumi.String("f14mpfy5lh4e"),
//				LifeCycleRule: &chdfs.LifeCycleRuleLifeCycleRuleArgs{
//					LifeCycleRuleName: pulumi.String("terraform-test"),
//					Path:              pulumi.String("/test"),
//					Status:            pulumi.Int(1),
//					Transitions: chdfs.LifeCycleRuleLifeCycleRuleTransitionArray{
//						&chdfs.LifeCycleRuleLifeCycleRuleTransitionArgs{
//							Days: pulumi.Int(30),
//							Type: pulumi.Int(1),
//						},
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
//
// ## Import
//
// chdfs life_cycle_rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule life_cycle_rule file_system_id#life_cycle_rule_id
//
// ```
type LifeCycleRule struct {
	pulumi.CustomResourceState

	// file system id.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// life cycle rule.
	LifeCycleRule LifeCycleRuleLifeCycleRuleOutput `pulumi:"lifeCycleRule"`
}

// NewLifeCycleRule registers a new resource with the given unique name, arguments, and options.
func NewLifeCycleRule(ctx *pulumi.Context,
	name string, args *LifeCycleRuleArgs, opts ...pulumi.ResourceOption) (*LifeCycleRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.LifeCycleRule == nil {
		return nil, errors.New("invalid value for required argument 'LifeCycleRule'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LifeCycleRule
	err := ctx.RegisterResource("tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifeCycleRule gets an existing LifeCycleRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifeCycleRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifeCycleRuleState, opts ...pulumi.ResourceOption) (*LifeCycleRule, error) {
	var resource LifeCycleRule
	err := ctx.ReadResource("tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifeCycleRule resources.
type lifeCycleRuleState struct {
	// file system id.
	FileSystemId *string `pulumi:"fileSystemId"`
	// life cycle rule.
	LifeCycleRule *LifeCycleRuleLifeCycleRule `pulumi:"lifeCycleRule"`
}

type LifeCycleRuleState struct {
	// file system id.
	FileSystemId pulumi.StringPtrInput
	// life cycle rule.
	LifeCycleRule LifeCycleRuleLifeCycleRulePtrInput
}

func (LifeCycleRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifeCycleRuleState)(nil)).Elem()
}

type lifeCycleRuleArgs struct {
	// file system id.
	FileSystemId string `pulumi:"fileSystemId"`
	// life cycle rule.
	LifeCycleRule LifeCycleRuleLifeCycleRule `pulumi:"lifeCycleRule"`
}

// The set of arguments for constructing a LifeCycleRule resource.
type LifeCycleRuleArgs struct {
	// file system id.
	FileSystemId pulumi.StringInput
	// life cycle rule.
	LifeCycleRule LifeCycleRuleLifeCycleRuleInput
}

func (LifeCycleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifeCycleRuleArgs)(nil)).Elem()
}

type LifeCycleRuleInput interface {
	pulumi.Input

	ToLifeCycleRuleOutput() LifeCycleRuleOutput
	ToLifeCycleRuleOutputWithContext(ctx context.Context) LifeCycleRuleOutput
}

func (*LifeCycleRule) ElementType() reflect.Type {
	return reflect.TypeOf((**LifeCycleRule)(nil)).Elem()
}

func (i *LifeCycleRule) ToLifeCycleRuleOutput() LifeCycleRuleOutput {
	return i.ToLifeCycleRuleOutputWithContext(context.Background())
}

func (i *LifeCycleRule) ToLifeCycleRuleOutputWithContext(ctx context.Context) LifeCycleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifeCycleRuleOutput)
}

// LifeCycleRuleArrayInput is an input type that accepts LifeCycleRuleArray and LifeCycleRuleArrayOutput values.
// You can construct a concrete instance of `LifeCycleRuleArrayInput` via:
//
//	LifeCycleRuleArray{ LifeCycleRuleArgs{...} }
type LifeCycleRuleArrayInput interface {
	pulumi.Input

	ToLifeCycleRuleArrayOutput() LifeCycleRuleArrayOutput
	ToLifeCycleRuleArrayOutputWithContext(context.Context) LifeCycleRuleArrayOutput
}

type LifeCycleRuleArray []LifeCycleRuleInput

func (LifeCycleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LifeCycleRule)(nil)).Elem()
}

func (i LifeCycleRuleArray) ToLifeCycleRuleArrayOutput() LifeCycleRuleArrayOutput {
	return i.ToLifeCycleRuleArrayOutputWithContext(context.Background())
}

func (i LifeCycleRuleArray) ToLifeCycleRuleArrayOutputWithContext(ctx context.Context) LifeCycleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifeCycleRuleArrayOutput)
}

// LifeCycleRuleMapInput is an input type that accepts LifeCycleRuleMap and LifeCycleRuleMapOutput values.
// You can construct a concrete instance of `LifeCycleRuleMapInput` via:
//
//	LifeCycleRuleMap{ "key": LifeCycleRuleArgs{...} }
type LifeCycleRuleMapInput interface {
	pulumi.Input

	ToLifeCycleRuleMapOutput() LifeCycleRuleMapOutput
	ToLifeCycleRuleMapOutputWithContext(context.Context) LifeCycleRuleMapOutput
}

type LifeCycleRuleMap map[string]LifeCycleRuleInput

func (LifeCycleRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LifeCycleRule)(nil)).Elem()
}

func (i LifeCycleRuleMap) ToLifeCycleRuleMapOutput() LifeCycleRuleMapOutput {
	return i.ToLifeCycleRuleMapOutputWithContext(context.Background())
}

func (i LifeCycleRuleMap) ToLifeCycleRuleMapOutputWithContext(ctx context.Context) LifeCycleRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifeCycleRuleMapOutput)
}

type LifeCycleRuleOutput struct{ *pulumi.OutputState }

func (LifeCycleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifeCycleRule)(nil)).Elem()
}

func (o LifeCycleRuleOutput) ToLifeCycleRuleOutput() LifeCycleRuleOutput {
	return o
}

func (o LifeCycleRuleOutput) ToLifeCycleRuleOutputWithContext(ctx context.Context) LifeCycleRuleOutput {
	return o
}

// file system id.
func (o LifeCycleRuleOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *LifeCycleRule) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// life cycle rule.
func (o LifeCycleRuleOutput) LifeCycleRule() LifeCycleRuleLifeCycleRuleOutput {
	return o.ApplyT(func(v *LifeCycleRule) LifeCycleRuleLifeCycleRuleOutput { return v.LifeCycleRule }).(LifeCycleRuleLifeCycleRuleOutput)
}

type LifeCycleRuleArrayOutput struct{ *pulumi.OutputState }

func (LifeCycleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LifeCycleRule)(nil)).Elem()
}

func (o LifeCycleRuleArrayOutput) ToLifeCycleRuleArrayOutput() LifeCycleRuleArrayOutput {
	return o
}

func (o LifeCycleRuleArrayOutput) ToLifeCycleRuleArrayOutputWithContext(ctx context.Context) LifeCycleRuleArrayOutput {
	return o
}

func (o LifeCycleRuleArrayOutput) Index(i pulumi.IntInput) LifeCycleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LifeCycleRule {
		return vs[0].([]*LifeCycleRule)[vs[1].(int)]
	}).(LifeCycleRuleOutput)
}

type LifeCycleRuleMapOutput struct{ *pulumi.OutputState }

func (LifeCycleRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LifeCycleRule)(nil)).Elem()
}

func (o LifeCycleRuleMapOutput) ToLifeCycleRuleMapOutput() LifeCycleRuleMapOutput {
	return o
}

func (o LifeCycleRuleMapOutput) ToLifeCycleRuleMapOutputWithContext(ctx context.Context) LifeCycleRuleMapOutput {
	return o
}

func (o LifeCycleRuleMapOutput) MapIndex(k pulumi.StringInput) LifeCycleRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LifeCycleRule {
		return vs[0].(map[string]*LifeCycleRule)[vs[1].(string)]
	}).(LifeCycleRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifeCycleRuleInput)(nil)).Elem(), &LifeCycleRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifeCycleRuleArrayInput)(nil)).Elem(), LifeCycleRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifeCycleRuleMapInput)(nil)).Elem(), LifeCycleRuleMap{})
	pulumi.RegisterOutputType(LifeCycleRuleOutput{})
	pulumi.RegisterOutputType(LifeCycleRuleArrayOutput{})
	pulumi.RegisterOutputType(LifeCycleRuleMapOutput{})
}
