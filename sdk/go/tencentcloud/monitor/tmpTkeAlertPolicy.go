// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tke tmpAlertPolicy
type TmpTkeAlertPolicy struct {
	pulumi.CustomResourceState

	// Alarm notification channels.
	AlertRule TmpTkeAlertPolicyAlertRuleOutput `pulumi:"alertRule"`
	// Instance Id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewTmpTkeAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewTmpTkeAlertPolicy(ctx *pulumi.Context,
	name string, args *TmpTkeAlertPolicyArgs, opts ...pulumi.ResourceOption) (*TmpTkeAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertRule == nil {
		return nil, errors.New("invalid value for required argument 'AlertRule'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource TmpTkeAlertPolicy
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpTkeAlertPolicy gets an existing TmpTkeAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpTkeAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpTkeAlertPolicyState, opts ...pulumi.ResourceOption) (*TmpTkeAlertPolicy, error) {
	var resource TmpTkeAlertPolicy
	err := ctx.ReadResource("tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpTkeAlertPolicy resources.
type tmpTkeAlertPolicyState struct {
	// Alarm notification channels.
	AlertRule *TmpTkeAlertPolicyAlertRule `pulumi:"alertRule"`
	// Instance Id.
	InstanceId *string `pulumi:"instanceId"`
}

type TmpTkeAlertPolicyState struct {
	// Alarm notification channels.
	AlertRule TmpTkeAlertPolicyAlertRulePtrInput
	// Instance Id.
	InstanceId pulumi.StringPtrInput
}

func (TmpTkeAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeAlertPolicyState)(nil)).Elem()
}

type tmpTkeAlertPolicyArgs struct {
	// Alarm notification channels.
	AlertRule TmpTkeAlertPolicyAlertRule `pulumi:"alertRule"`
	// Instance Id.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a TmpTkeAlertPolicy resource.
type TmpTkeAlertPolicyArgs struct {
	// Alarm notification channels.
	AlertRule TmpTkeAlertPolicyAlertRuleInput
	// Instance Id.
	InstanceId pulumi.StringInput
}

func (TmpTkeAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeAlertPolicyArgs)(nil)).Elem()
}

type TmpTkeAlertPolicyInput interface {
	pulumi.Input

	ToTmpTkeAlertPolicyOutput() TmpTkeAlertPolicyOutput
	ToTmpTkeAlertPolicyOutputWithContext(ctx context.Context) TmpTkeAlertPolicyOutput
}

func (*TmpTkeAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeAlertPolicy)(nil)).Elem()
}

func (i *TmpTkeAlertPolicy) ToTmpTkeAlertPolicyOutput() TmpTkeAlertPolicyOutput {
	return i.ToTmpTkeAlertPolicyOutputWithContext(context.Background())
}

func (i *TmpTkeAlertPolicy) ToTmpTkeAlertPolicyOutputWithContext(ctx context.Context) TmpTkeAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeAlertPolicyOutput)
}

// TmpTkeAlertPolicyArrayInput is an input type that accepts TmpTkeAlertPolicyArray and TmpTkeAlertPolicyArrayOutput values.
// You can construct a concrete instance of `TmpTkeAlertPolicyArrayInput` via:
//
//	TmpTkeAlertPolicyArray{ TmpTkeAlertPolicyArgs{...} }
type TmpTkeAlertPolicyArrayInput interface {
	pulumi.Input

	ToTmpTkeAlertPolicyArrayOutput() TmpTkeAlertPolicyArrayOutput
	ToTmpTkeAlertPolicyArrayOutputWithContext(context.Context) TmpTkeAlertPolicyArrayOutput
}

type TmpTkeAlertPolicyArray []TmpTkeAlertPolicyInput

func (TmpTkeAlertPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeAlertPolicy)(nil)).Elem()
}

func (i TmpTkeAlertPolicyArray) ToTmpTkeAlertPolicyArrayOutput() TmpTkeAlertPolicyArrayOutput {
	return i.ToTmpTkeAlertPolicyArrayOutputWithContext(context.Background())
}

func (i TmpTkeAlertPolicyArray) ToTmpTkeAlertPolicyArrayOutputWithContext(ctx context.Context) TmpTkeAlertPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeAlertPolicyArrayOutput)
}

// TmpTkeAlertPolicyMapInput is an input type that accepts TmpTkeAlertPolicyMap and TmpTkeAlertPolicyMapOutput values.
// You can construct a concrete instance of `TmpTkeAlertPolicyMapInput` via:
//
//	TmpTkeAlertPolicyMap{ "key": TmpTkeAlertPolicyArgs{...} }
type TmpTkeAlertPolicyMapInput interface {
	pulumi.Input

	ToTmpTkeAlertPolicyMapOutput() TmpTkeAlertPolicyMapOutput
	ToTmpTkeAlertPolicyMapOutputWithContext(context.Context) TmpTkeAlertPolicyMapOutput
}

type TmpTkeAlertPolicyMap map[string]TmpTkeAlertPolicyInput

func (TmpTkeAlertPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeAlertPolicy)(nil)).Elem()
}

func (i TmpTkeAlertPolicyMap) ToTmpTkeAlertPolicyMapOutput() TmpTkeAlertPolicyMapOutput {
	return i.ToTmpTkeAlertPolicyMapOutputWithContext(context.Background())
}

func (i TmpTkeAlertPolicyMap) ToTmpTkeAlertPolicyMapOutputWithContext(ctx context.Context) TmpTkeAlertPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeAlertPolicyMapOutput)
}

type TmpTkeAlertPolicyOutput struct{ *pulumi.OutputState }

func (TmpTkeAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeAlertPolicy)(nil)).Elem()
}

func (o TmpTkeAlertPolicyOutput) ToTmpTkeAlertPolicyOutput() TmpTkeAlertPolicyOutput {
	return o
}

func (o TmpTkeAlertPolicyOutput) ToTmpTkeAlertPolicyOutputWithContext(ctx context.Context) TmpTkeAlertPolicyOutput {
	return o
}

// Alarm notification channels.
func (o TmpTkeAlertPolicyOutput) AlertRule() TmpTkeAlertPolicyAlertRuleOutput {
	return o.ApplyT(func(v *TmpTkeAlertPolicy) TmpTkeAlertPolicyAlertRuleOutput { return v.AlertRule }).(TmpTkeAlertPolicyAlertRuleOutput)
}

// Instance Id.
func (o TmpTkeAlertPolicyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeAlertPolicy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type TmpTkeAlertPolicyArrayOutput struct{ *pulumi.OutputState }

func (TmpTkeAlertPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeAlertPolicy)(nil)).Elem()
}

func (o TmpTkeAlertPolicyArrayOutput) ToTmpTkeAlertPolicyArrayOutput() TmpTkeAlertPolicyArrayOutput {
	return o
}

func (o TmpTkeAlertPolicyArrayOutput) ToTmpTkeAlertPolicyArrayOutputWithContext(ctx context.Context) TmpTkeAlertPolicyArrayOutput {
	return o
}

func (o TmpTkeAlertPolicyArrayOutput) Index(i pulumi.IntInput) TmpTkeAlertPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpTkeAlertPolicy {
		return vs[0].([]*TmpTkeAlertPolicy)[vs[1].(int)]
	}).(TmpTkeAlertPolicyOutput)
}

type TmpTkeAlertPolicyMapOutput struct{ *pulumi.OutputState }

func (TmpTkeAlertPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeAlertPolicy)(nil)).Elem()
}

func (o TmpTkeAlertPolicyMapOutput) ToTmpTkeAlertPolicyMapOutput() TmpTkeAlertPolicyMapOutput {
	return o
}

func (o TmpTkeAlertPolicyMapOutput) ToTmpTkeAlertPolicyMapOutputWithContext(ctx context.Context) TmpTkeAlertPolicyMapOutput {
	return o
}

func (o TmpTkeAlertPolicyMapOutput) MapIndex(k pulumi.StringInput) TmpTkeAlertPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpTkeAlertPolicy {
		return vs[0].(map[string]*TmpTkeAlertPolicy)[vs[1].(string)]
	}).(TmpTkeAlertPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeAlertPolicyInput)(nil)).Elem(), &TmpTkeAlertPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeAlertPolicyArrayInput)(nil)).Elem(), TmpTkeAlertPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeAlertPolicyMapInput)(nil)).Elem(), TmpTkeAlertPolicyMap{})
	pulumi.RegisterOutputType(TmpTkeAlertPolicyOutput{})
	pulumi.RegisterOutputType(TmpTkeAlertPolicyArrayOutput{})
	pulumi.RegisterOutputType(TmpTkeAlertPolicyMapOutput{})
}
