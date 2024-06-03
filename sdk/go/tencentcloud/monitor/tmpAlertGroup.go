// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a monitor tmpAlertGroup
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.NewTmpAlertGroup(ctx, "tmpAlertGroup", &Monitor.TmpAlertGroupArgs{
//				AmpReceivers: pulumi.StringArray{
//					pulumi.String("notice-om017kc2"),
//				},
//				CustomReceiver: &monitor.TmpAlertGroupCustomReceiverArgs{
//					Type: pulumi.String("amp"),
//				},
//				GroupName:      pulumi.String("tf-test"),
//				InstanceId:     pulumi.String("prom-ip429jis"),
//				RepeatInterval: pulumi.String("5m"),
//				Rules: monitor.TmpAlertGroupRuleArray{
//					&monitor.TmpAlertGroupRuleArgs{
//						Annotations: pulumi.Map{
//							"description": pulumi.Any("Agent {{$labels.instance}} is deactivated, please pay attention!"),
//							"summary":     pulumi.Any("Agent health check"),
//						},
//						Duration: pulumi.String("1m"),
//						Expr:     pulumi.String("up{job=\"prometheus-agent\"} != 1"),
//						Labels: pulumi.Map{
//							"severity": pulumi.Any("critical"),
//						},
//						RuleName: pulumi.String("Agent health check"),
//						State:    pulumi.Int(2),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// monitor tmp_alert_group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Monitor/tmpAlertGroup:TmpAlertGroup tmp_alert_group instance_id#group_id
// ```
type TmpAlertGroup struct {
	pulumi.CustomResourceState

	// Tencent cloud notification template id list.
	AmpReceivers pulumi.StringArrayOutput `pulumi:"ampReceivers"`
	// User custom notification template, such as webhook, alertmanager.
	CustomReceiver TmpAlertGroupCustomReceiverPtrOutput `pulumi:"customReceiver"`
	// Alarm group id.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Unique alert group name.
	GroupName pulumi.StringPtrOutput `pulumi:"groupName"`
	// Instance id.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// Alert message send interval, default 1 hour.
	RepeatInterval pulumi.StringPtrOutput `pulumi:"repeatInterval"`
	// A list of alert rules.
	Rules TmpAlertGroupRuleArrayOutput `pulumi:"rules"`
}

// NewTmpAlertGroup registers a new resource with the given unique name, arguments, and options.
func NewTmpAlertGroup(ctx *pulumi.Context,
	name string, args *TmpAlertGroupArgs, opts ...pulumi.ResourceOption) (*TmpAlertGroup, error) {
	if args == nil {
		args = &TmpAlertGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TmpAlertGroup
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpAlertGroup:TmpAlertGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpAlertGroup gets an existing TmpAlertGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpAlertGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpAlertGroupState, opts ...pulumi.ResourceOption) (*TmpAlertGroup, error) {
	var resource TmpAlertGroup
	err := ctx.ReadResource("tencentcloud:Monitor/tmpAlertGroup:TmpAlertGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpAlertGroup resources.
type tmpAlertGroupState struct {
	// Tencent cloud notification template id list.
	AmpReceivers []string `pulumi:"ampReceivers"`
	// User custom notification template, such as webhook, alertmanager.
	CustomReceiver *TmpAlertGroupCustomReceiver `pulumi:"customReceiver"`
	// Alarm group id.
	GroupId *string `pulumi:"groupId"`
	// Unique alert group name.
	GroupName *string `pulumi:"groupName"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Alert message send interval, default 1 hour.
	RepeatInterval *string `pulumi:"repeatInterval"`
	// A list of alert rules.
	Rules []TmpAlertGroupRule `pulumi:"rules"`
}

type TmpAlertGroupState struct {
	// Tencent cloud notification template id list.
	AmpReceivers pulumi.StringArrayInput
	// User custom notification template, such as webhook, alertmanager.
	CustomReceiver TmpAlertGroupCustomReceiverPtrInput
	// Alarm group id.
	GroupId pulumi.StringPtrInput
	// Unique alert group name.
	GroupName pulumi.StringPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Alert message send interval, default 1 hour.
	RepeatInterval pulumi.StringPtrInput
	// A list of alert rules.
	Rules TmpAlertGroupRuleArrayInput
}

func (TmpAlertGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpAlertGroupState)(nil)).Elem()
}

type tmpAlertGroupArgs struct {
	// Tencent cloud notification template id list.
	AmpReceivers []string `pulumi:"ampReceivers"`
	// User custom notification template, such as webhook, alertmanager.
	CustomReceiver *TmpAlertGroupCustomReceiver `pulumi:"customReceiver"`
	// Unique alert group name.
	GroupName *string `pulumi:"groupName"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Alert message send interval, default 1 hour.
	RepeatInterval *string `pulumi:"repeatInterval"`
	// A list of alert rules.
	Rules []TmpAlertGroupRule `pulumi:"rules"`
}

// The set of arguments for constructing a TmpAlertGroup resource.
type TmpAlertGroupArgs struct {
	// Tencent cloud notification template id list.
	AmpReceivers pulumi.StringArrayInput
	// User custom notification template, such as webhook, alertmanager.
	CustomReceiver TmpAlertGroupCustomReceiverPtrInput
	// Unique alert group name.
	GroupName pulumi.StringPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Alert message send interval, default 1 hour.
	RepeatInterval pulumi.StringPtrInput
	// A list of alert rules.
	Rules TmpAlertGroupRuleArrayInput
}

func (TmpAlertGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpAlertGroupArgs)(nil)).Elem()
}

type TmpAlertGroupInput interface {
	pulumi.Input

	ToTmpAlertGroupOutput() TmpAlertGroupOutput
	ToTmpAlertGroupOutputWithContext(ctx context.Context) TmpAlertGroupOutput
}

func (*TmpAlertGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpAlertGroup)(nil)).Elem()
}

func (i *TmpAlertGroup) ToTmpAlertGroupOutput() TmpAlertGroupOutput {
	return i.ToTmpAlertGroupOutputWithContext(context.Background())
}

func (i *TmpAlertGroup) ToTmpAlertGroupOutputWithContext(ctx context.Context) TmpAlertGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertGroupOutput)
}

// TmpAlertGroupArrayInput is an input type that accepts TmpAlertGroupArray and TmpAlertGroupArrayOutput values.
// You can construct a concrete instance of `TmpAlertGroupArrayInput` via:
//
//	TmpAlertGroupArray{ TmpAlertGroupArgs{...} }
type TmpAlertGroupArrayInput interface {
	pulumi.Input

	ToTmpAlertGroupArrayOutput() TmpAlertGroupArrayOutput
	ToTmpAlertGroupArrayOutputWithContext(context.Context) TmpAlertGroupArrayOutput
}

type TmpAlertGroupArray []TmpAlertGroupInput

func (TmpAlertGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpAlertGroup)(nil)).Elem()
}

func (i TmpAlertGroupArray) ToTmpAlertGroupArrayOutput() TmpAlertGroupArrayOutput {
	return i.ToTmpAlertGroupArrayOutputWithContext(context.Background())
}

func (i TmpAlertGroupArray) ToTmpAlertGroupArrayOutputWithContext(ctx context.Context) TmpAlertGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertGroupArrayOutput)
}

// TmpAlertGroupMapInput is an input type that accepts TmpAlertGroupMap and TmpAlertGroupMapOutput values.
// You can construct a concrete instance of `TmpAlertGroupMapInput` via:
//
//	TmpAlertGroupMap{ "key": TmpAlertGroupArgs{...} }
type TmpAlertGroupMapInput interface {
	pulumi.Input

	ToTmpAlertGroupMapOutput() TmpAlertGroupMapOutput
	ToTmpAlertGroupMapOutputWithContext(context.Context) TmpAlertGroupMapOutput
}

type TmpAlertGroupMap map[string]TmpAlertGroupInput

func (TmpAlertGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpAlertGroup)(nil)).Elem()
}

func (i TmpAlertGroupMap) ToTmpAlertGroupMapOutput() TmpAlertGroupMapOutput {
	return i.ToTmpAlertGroupMapOutputWithContext(context.Background())
}

func (i TmpAlertGroupMap) ToTmpAlertGroupMapOutputWithContext(ctx context.Context) TmpAlertGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertGroupMapOutput)
}

type TmpAlertGroupOutput struct{ *pulumi.OutputState }

func (TmpAlertGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpAlertGroup)(nil)).Elem()
}

func (o TmpAlertGroupOutput) ToTmpAlertGroupOutput() TmpAlertGroupOutput {
	return o
}

func (o TmpAlertGroupOutput) ToTmpAlertGroupOutputWithContext(ctx context.Context) TmpAlertGroupOutput {
	return o
}

// Tencent cloud notification template id list.
func (o TmpAlertGroupOutput) AmpReceivers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TmpAlertGroup) pulumi.StringArrayOutput { return v.AmpReceivers }).(pulumi.StringArrayOutput)
}

// User custom notification template, such as webhook, alertmanager.
func (o TmpAlertGroupOutput) CustomReceiver() TmpAlertGroupCustomReceiverPtrOutput {
	return o.ApplyT(func(v *TmpAlertGroup) TmpAlertGroupCustomReceiverPtrOutput { return v.CustomReceiver }).(TmpAlertGroupCustomReceiverPtrOutput)
}

// Alarm group id.
func (o TmpAlertGroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpAlertGroup) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Unique alert group name.
func (o TmpAlertGroupOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TmpAlertGroup) pulumi.StringPtrOutput { return v.GroupName }).(pulumi.StringPtrOutput)
}

// Instance id.
func (o TmpAlertGroupOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TmpAlertGroup) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// Alert message send interval, default 1 hour.
func (o TmpAlertGroupOutput) RepeatInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TmpAlertGroup) pulumi.StringPtrOutput { return v.RepeatInterval }).(pulumi.StringPtrOutput)
}

// A list of alert rules.
func (o TmpAlertGroupOutput) Rules() TmpAlertGroupRuleArrayOutput {
	return o.ApplyT(func(v *TmpAlertGroup) TmpAlertGroupRuleArrayOutput { return v.Rules }).(TmpAlertGroupRuleArrayOutput)
}

type TmpAlertGroupArrayOutput struct{ *pulumi.OutputState }

func (TmpAlertGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpAlertGroup)(nil)).Elem()
}

func (o TmpAlertGroupArrayOutput) ToTmpAlertGroupArrayOutput() TmpAlertGroupArrayOutput {
	return o
}

func (o TmpAlertGroupArrayOutput) ToTmpAlertGroupArrayOutputWithContext(ctx context.Context) TmpAlertGroupArrayOutput {
	return o
}

func (o TmpAlertGroupArrayOutput) Index(i pulumi.IntInput) TmpAlertGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpAlertGroup {
		return vs[0].([]*TmpAlertGroup)[vs[1].(int)]
	}).(TmpAlertGroupOutput)
}

type TmpAlertGroupMapOutput struct{ *pulumi.OutputState }

func (TmpAlertGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpAlertGroup)(nil)).Elem()
}

func (o TmpAlertGroupMapOutput) ToTmpAlertGroupMapOutput() TmpAlertGroupMapOutput {
	return o
}

func (o TmpAlertGroupMapOutput) ToTmpAlertGroupMapOutputWithContext(ctx context.Context) TmpAlertGroupMapOutput {
	return o
}

func (o TmpAlertGroupMapOutput) MapIndex(k pulumi.StringInput) TmpAlertGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpAlertGroup {
		return vs[0].(map[string]*TmpAlertGroup)[vs[1].(string)]
	}).(TmpAlertGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertGroupInput)(nil)).Elem(), &TmpAlertGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertGroupArrayInput)(nil)).Elem(), TmpAlertGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertGroupMapInput)(nil)).Elem(), TmpAlertGroupMap{})
	pulumi.RegisterOutputType(TmpAlertGroupOutput{})
	pulumi.RegisterOutputType(TmpAlertGroupArrayOutput{})
	pulumi.RegisterOutputType(TmpAlertGroupMapOutput{})
}
