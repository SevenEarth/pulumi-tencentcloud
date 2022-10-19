// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor tmp recordingRule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.NewTmpRecordingRule(ctx, "recordingRule", &Monitor.TmpRecordingRuleArgs{
//				Group:      pulumi.String("LS0tDQpuYW1lOiBleGFtcGxlDQpydWxlczoNCiAgLSByZWNvcmQ6IGpvYjpodHRwX2lucHJvZ3Jlc3NfcmVxdWVzdHM6c3VtDQogICAgZXhwcjogc3VtIGJ5IChqb2IpIChodHRwX2lucHJvZ3Jlc3NfcmVxdWVzdHMp"),
//				InstanceId: pulumi.String("prom-c89b3b3u"),
//				RuleState:  pulumi.Int(2),
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
// monitor recordingRule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule recordingRule instanceId#recordingRule_id
//
// ```
type TmpRecordingRule struct {
	pulumi.CustomResourceState

	// Recording rule group.
	Group pulumi.StringOutput `pulumi:"group"`
	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Recording rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule state.
	RuleState pulumi.IntPtrOutput `pulumi:"ruleState"`
}

// NewTmpRecordingRule registers a new resource with the given unique name, arguments, and options.
func NewTmpRecordingRule(ctx *pulumi.Context,
	name string, args *TmpRecordingRuleArgs, opts ...pulumi.ResourceOption) (*TmpRecordingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource TmpRecordingRule
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpRecordingRule gets an existing TmpRecordingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpRecordingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpRecordingRuleState, opts ...pulumi.ResourceOption) (*TmpRecordingRule, error) {
	var resource TmpRecordingRule
	err := ctx.ReadResource("tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpRecordingRule resources.
type tmpRecordingRuleState struct {
	// Recording rule group.
	Group *string `pulumi:"group"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Recording rule name.
	Name *string `pulumi:"name"`
	// Rule state.
	RuleState *int `pulumi:"ruleState"`
}

type TmpRecordingRuleState struct {
	// Recording rule group.
	Group pulumi.StringPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Recording rule name.
	Name pulumi.StringPtrInput
	// Rule state.
	RuleState pulumi.IntPtrInput
}

func (TmpRecordingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpRecordingRuleState)(nil)).Elem()
}

type tmpRecordingRuleArgs struct {
	// Recording rule group.
	Group string `pulumi:"group"`
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Recording rule name.
	Name *string `pulumi:"name"`
	// Rule state.
	RuleState *int `pulumi:"ruleState"`
}

// The set of arguments for constructing a TmpRecordingRule resource.
type TmpRecordingRuleArgs struct {
	// Recording rule group.
	Group pulumi.StringInput
	// Instance id.
	InstanceId pulumi.StringInput
	// Recording rule name.
	Name pulumi.StringPtrInput
	// Rule state.
	RuleState pulumi.IntPtrInput
}

func (TmpRecordingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpRecordingRuleArgs)(nil)).Elem()
}

type TmpRecordingRuleInput interface {
	pulumi.Input

	ToTmpRecordingRuleOutput() TmpRecordingRuleOutput
	ToTmpRecordingRuleOutputWithContext(ctx context.Context) TmpRecordingRuleOutput
}

func (*TmpRecordingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpRecordingRule)(nil)).Elem()
}

func (i *TmpRecordingRule) ToTmpRecordingRuleOutput() TmpRecordingRuleOutput {
	return i.ToTmpRecordingRuleOutputWithContext(context.Background())
}

func (i *TmpRecordingRule) ToTmpRecordingRuleOutputWithContext(ctx context.Context) TmpRecordingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpRecordingRuleOutput)
}

// TmpRecordingRuleArrayInput is an input type that accepts TmpRecordingRuleArray and TmpRecordingRuleArrayOutput values.
// You can construct a concrete instance of `TmpRecordingRuleArrayInput` via:
//
//	TmpRecordingRuleArray{ TmpRecordingRuleArgs{...} }
type TmpRecordingRuleArrayInput interface {
	pulumi.Input

	ToTmpRecordingRuleArrayOutput() TmpRecordingRuleArrayOutput
	ToTmpRecordingRuleArrayOutputWithContext(context.Context) TmpRecordingRuleArrayOutput
}

type TmpRecordingRuleArray []TmpRecordingRuleInput

func (TmpRecordingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpRecordingRule)(nil)).Elem()
}

func (i TmpRecordingRuleArray) ToTmpRecordingRuleArrayOutput() TmpRecordingRuleArrayOutput {
	return i.ToTmpRecordingRuleArrayOutputWithContext(context.Background())
}

func (i TmpRecordingRuleArray) ToTmpRecordingRuleArrayOutputWithContext(ctx context.Context) TmpRecordingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpRecordingRuleArrayOutput)
}

// TmpRecordingRuleMapInput is an input type that accepts TmpRecordingRuleMap and TmpRecordingRuleMapOutput values.
// You can construct a concrete instance of `TmpRecordingRuleMapInput` via:
//
//	TmpRecordingRuleMap{ "key": TmpRecordingRuleArgs{...} }
type TmpRecordingRuleMapInput interface {
	pulumi.Input

	ToTmpRecordingRuleMapOutput() TmpRecordingRuleMapOutput
	ToTmpRecordingRuleMapOutputWithContext(context.Context) TmpRecordingRuleMapOutput
}

type TmpRecordingRuleMap map[string]TmpRecordingRuleInput

func (TmpRecordingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpRecordingRule)(nil)).Elem()
}

func (i TmpRecordingRuleMap) ToTmpRecordingRuleMapOutput() TmpRecordingRuleMapOutput {
	return i.ToTmpRecordingRuleMapOutputWithContext(context.Background())
}

func (i TmpRecordingRuleMap) ToTmpRecordingRuleMapOutputWithContext(ctx context.Context) TmpRecordingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpRecordingRuleMapOutput)
}

type TmpRecordingRuleOutput struct{ *pulumi.OutputState }

func (TmpRecordingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpRecordingRule)(nil)).Elem()
}

func (o TmpRecordingRuleOutput) ToTmpRecordingRuleOutput() TmpRecordingRuleOutput {
	return o
}

func (o TmpRecordingRuleOutput) ToTmpRecordingRuleOutputWithContext(ctx context.Context) TmpRecordingRuleOutput {
	return o
}

// Recording rule group.
func (o TmpRecordingRuleOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpRecordingRule) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Instance id.
func (o TmpRecordingRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpRecordingRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Recording rule name.
func (o TmpRecordingRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpRecordingRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Rule state.
func (o TmpRecordingRuleOutput) RuleState() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TmpRecordingRule) pulumi.IntPtrOutput { return v.RuleState }).(pulumi.IntPtrOutput)
}

type TmpRecordingRuleArrayOutput struct{ *pulumi.OutputState }

func (TmpRecordingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpRecordingRule)(nil)).Elem()
}

func (o TmpRecordingRuleArrayOutput) ToTmpRecordingRuleArrayOutput() TmpRecordingRuleArrayOutput {
	return o
}

func (o TmpRecordingRuleArrayOutput) ToTmpRecordingRuleArrayOutputWithContext(ctx context.Context) TmpRecordingRuleArrayOutput {
	return o
}

func (o TmpRecordingRuleArrayOutput) Index(i pulumi.IntInput) TmpRecordingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpRecordingRule {
		return vs[0].([]*TmpRecordingRule)[vs[1].(int)]
	}).(TmpRecordingRuleOutput)
}

type TmpRecordingRuleMapOutput struct{ *pulumi.OutputState }

func (TmpRecordingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpRecordingRule)(nil)).Elem()
}

func (o TmpRecordingRuleMapOutput) ToTmpRecordingRuleMapOutput() TmpRecordingRuleMapOutput {
	return o
}

func (o TmpRecordingRuleMapOutput) ToTmpRecordingRuleMapOutputWithContext(ctx context.Context) TmpRecordingRuleMapOutput {
	return o
}

func (o TmpRecordingRuleMapOutput) MapIndex(k pulumi.StringInput) TmpRecordingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpRecordingRule {
		return vs[0].(map[string]*TmpRecordingRule)[vs[1].(string)]
	}).(TmpRecordingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpRecordingRuleInput)(nil)).Elem(), &TmpRecordingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpRecordingRuleArrayInput)(nil)).Elem(), TmpRecordingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpRecordingRuleMapInput)(nil)).Elem(), TmpRecordingRuleMap{})
	pulumi.RegisterOutputType(TmpRecordingRuleOutput{})
	pulumi.RegisterOutputType(TmpRecordingRuleArrayOutput{})
	pulumi.RegisterOutputType(TmpRecordingRuleMapOutput{})
}
