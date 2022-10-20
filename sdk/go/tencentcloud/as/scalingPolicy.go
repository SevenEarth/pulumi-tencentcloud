// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for an AS (Auto scaling) policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := As.NewScalingPolicy(ctx, "scalingPolicy", &As.ScalingPolicyArgs{
//				AdjustmentType:     pulumi.String("EXACT_CAPACITY"),
//				AdjustmentValue:    pulumi.Int(0),
//				ComparisonOperator: pulumi.String("GREATER_THAN"),
//				ContinuousTime:     pulumi.Int(10),
//				Cooldown:           pulumi.Int(360),
//				MetricName:         pulumi.String("CPU_UTILIZATION"),
//				Period:             pulumi.Int(300),
//				PolicyName:         pulumi.String("tf-as-scaling-policy"),
//				ScalingGroupId:     pulumi.String("asg-n32ymck2"),
//				Statistic:          pulumi.String("AVERAGE"),
//				Threshold:          pulumi.Int(80),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ScalingPolicy struct {
	pulumi.CustomResourceState

	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	AdjustmentType pulumi.StringOutput `pulumi:"adjustmentType"`
	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	AdjustmentValue pulumi.IntOutput `pulumi:"adjustmentValue"`
	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	ComparisonOperator pulumi.StringOutput `pulumi:"comparisonOperator"`
	// Retry times. Valid value ranges: (1~10).
	ContinuousTime pulumi.IntOutput `pulumi:"continuousTime"`
	// Cooldwon time in second. Default is `30`0.
	Cooldown pulumi.IntPtrOutput `pulumi:"cooldown"`
	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// An ID group of users to be notified when an alarm is triggered.
	NotificationUserGroupIds pulumi.StringArrayOutput `pulumi:"notificationUserGroupIds"`
	// Time period in second. Valid values: `60` and `300`.
	Period pulumi.IntOutput `pulumi:"period"`
	// Name of a policy used to define a reaction when an alarm is triggered.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// ID of a scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	Statistic pulumi.StringPtrOutput `pulumi:"statistic"`
	// Alarm threshold.
	Threshold pulumi.IntOutput `pulumi:"threshold"`
}

// NewScalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewScalingPolicy(ctx *pulumi.Context,
	name string, args *ScalingPolicyArgs, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdjustmentType == nil {
		return nil, errors.New("invalid value for required argument 'AdjustmentType'")
	}
	if args.AdjustmentValue == nil {
		return nil, errors.New("invalid value for required argument 'AdjustmentValue'")
	}
	if args.ComparisonOperator == nil {
		return nil, errors.New("invalid value for required argument 'ComparisonOperator'")
	}
	if args.ContinuousTime == nil {
		return nil, errors.New("invalid value for required argument 'ContinuousTime'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	if args.Threshold == nil {
		return nil, errors.New("invalid value for required argument 'Threshold'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScalingPolicy
	err := ctx.RegisterResource("tencentcloud:As/scalingPolicy:ScalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPolicy gets an existing ScalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPolicyState, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	var resource ScalingPolicy
	err := ctx.ReadResource("tencentcloud:As/scalingPolicy:ScalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPolicy resources.
type scalingPolicyState struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	AdjustmentValue *int `pulumi:"adjustmentValue"`
	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	ComparisonOperator *string `pulumi:"comparisonOperator"`
	// Retry times. Valid value ranges: (1~10).
	ContinuousTime *int `pulumi:"continuousTime"`
	// Cooldwon time in second. Default is `30`0.
	Cooldown *int `pulumi:"cooldown"`
	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	MetricName *string `pulumi:"metricName"`
	// An ID group of users to be notified when an alarm is triggered.
	NotificationUserGroupIds []string `pulumi:"notificationUserGroupIds"`
	// Time period in second. Valid values: `60` and `300`.
	Period *int `pulumi:"period"`
	// Name of a policy used to define a reaction when an alarm is triggered.
	PolicyName *string `pulumi:"policyName"`
	// ID of a scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	Statistic *string `pulumi:"statistic"`
	// Alarm threshold.
	Threshold *int `pulumi:"threshold"`
}

type ScalingPolicyState struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	AdjustmentType pulumi.StringPtrInput
	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	AdjustmentValue pulumi.IntPtrInput
	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	ComparisonOperator pulumi.StringPtrInput
	// Retry times. Valid value ranges: (1~10).
	ContinuousTime pulumi.IntPtrInput
	// Cooldwon time in second. Default is `30`0.
	Cooldown pulumi.IntPtrInput
	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	MetricName pulumi.StringPtrInput
	// An ID group of users to be notified when an alarm is triggered.
	NotificationUserGroupIds pulumi.StringArrayInput
	// Time period in second. Valid values: `60` and `300`.
	Period pulumi.IntPtrInput
	// Name of a policy used to define a reaction when an alarm is triggered.
	PolicyName pulumi.StringPtrInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringPtrInput
	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	Statistic pulumi.StringPtrInput
	// Alarm threshold.
	Threshold pulumi.IntPtrInput
}

func (ScalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyState)(nil)).Elem()
}

type scalingPolicyArgs struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	AdjustmentType string `pulumi:"adjustmentType"`
	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	AdjustmentValue int `pulumi:"adjustmentValue"`
	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	ComparisonOperator string `pulumi:"comparisonOperator"`
	// Retry times. Valid value ranges: (1~10).
	ContinuousTime int `pulumi:"continuousTime"`
	// Cooldwon time in second. Default is `30`0.
	Cooldown *int `pulumi:"cooldown"`
	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	MetricName string `pulumi:"metricName"`
	// An ID group of users to be notified when an alarm is triggered.
	NotificationUserGroupIds []string `pulumi:"notificationUserGroupIds"`
	// Time period in second. Valid values: `60` and `300`.
	Period int `pulumi:"period"`
	// Name of a policy used to define a reaction when an alarm is triggered.
	PolicyName string `pulumi:"policyName"`
	// ID of a scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	Statistic *string `pulumi:"statistic"`
	// Alarm threshold.
	Threshold int `pulumi:"threshold"`
}

// The set of arguments for constructing a ScalingPolicy resource.
type ScalingPolicyArgs struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	AdjustmentType pulumi.StringInput
	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	AdjustmentValue pulumi.IntInput
	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	ComparisonOperator pulumi.StringInput
	// Retry times. Valid value ranges: (1~10).
	ContinuousTime pulumi.IntInput
	// Cooldwon time in second. Default is `30`0.
	Cooldown pulumi.IntPtrInput
	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	MetricName pulumi.StringInput
	// An ID group of users to be notified when an alarm is triggered.
	NotificationUserGroupIds pulumi.StringArrayInput
	// Time period in second. Valid values: `60` and `300`.
	Period pulumi.IntInput
	// Name of a policy used to define a reaction when an alarm is triggered.
	PolicyName pulumi.StringInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringInput
	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	Statistic pulumi.StringPtrInput
	// Alarm threshold.
	Threshold pulumi.IntInput
}

func (ScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyArgs)(nil)).Elem()
}

type ScalingPolicyInput interface {
	pulumi.Input

	ToScalingPolicyOutput() ScalingPolicyOutput
	ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput
}

func (*ScalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (i *ScalingPolicy) ToScalingPolicyOutput() ScalingPolicyOutput {
	return i.ToScalingPolicyOutputWithContext(context.Background())
}

func (i *ScalingPolicy) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyOutput)
}

// ScalingPolicyArrayInput is an input type that accepts ScalingPolicyArray and ScalingPolicyArrayOutput values.
// You can construct a concrete instance of `ScalingPolicyArrayInput` via:
//
//	ScalingPolicyArray{ ScalingPolicyArgs{...} }
type ScalingPolicyArrayInput interface {
	pulumi.Input

	ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput
	ToScalingPolicyArrayOutputWithContext(context.Context) ScalingPolicyArrayOutput
}

type ScalingPolicyArray []ScalingPolicyInput

func (ScalingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingPolicy)(nil)).Elem()
}

func (i ScalingPolicyArray) ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput {
	return i.ToScalingPolicyArrayOutputWithContext(context.Background())
}

func (i ScalingPolicyArray) ToScalingPolicyArrayOutputWithContext(ctx context.Context) ScalingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyArrayOutput)
}

// ScalingPolicyMapInput is an input type that accepts ScalingPolicyMap and ScalingPolicyMapOutput values.
// You can construct a concrete instance of `ScalingPolicyMapInput` via:
//
//	ScalingPolicyMap{ "key": ScalingPolicyArgs{...} }
type ScalingPolicyMapInput interface {
	pulumi.Input

	ToScalingPolicyMapOutput() ScalingPolicyMapOutput
	ToScalingPolicyMapOutputWithContext(context.Context) ScalingPolicyMapOutput
}

type ScalingPolicyMap map[string]ScalingPolicyInput

func (ScalingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingPolicy)(nil)).Elem()
}

func (i ScalingPolicyMap) ToScalingPolicyMapOutput() ScalingPolicyMapOutput {
	return i.ToScalingPolicyMapOutputWithContext(context.Background())
}

func (i ScalingPolicyMap) ToScalingPolicyMapOutputWithContext(ctx context.Context) ScalingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyMapOutput)
}

type ScalingPolicyOutput struct{ *pulumi.OutputState }

func (ScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyOutput) ToScalingPolicyOutput() ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return o
}

// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
func (o ScalingPolicyOutput) AdjustmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.AdjustmentType }).(pulumi.StringOutput)
}

// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
func (o ScalingPolicyOutput) AdjustmentValue() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.IntOutput { return v.AdjustmentValue }).(pulumi.IntOutput)
}

// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
func (o ScalingPolicyOutput) ComparisonOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.ComparisonOperator }).(pulumi.StringOutput)
}

// Retry times. Valid value ranges: (1~10).
func (o ScalingPolicyOutput) ContinuousTime() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.IntOutput { return v.ContinuousTime }).(pulumi.IntOutput)
}

// Cooldwon time in second. Default is `30`0.
func (o ScalingPolicyOutput) Cooldown() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.IntPtrOutput { return v.Cooldown }).(pulumi.IntPtrOutput)
}

// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
func (o ScalingPolicyOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

// An ID group of users to be notified when an alarm is triggered.
func (o ScalingPolicyOutput) NotificationUserGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringArrayOutput { return v.NotificationUserGroupIds }).(pulumi.StringArrayOutput)
}

// Time period in second. Valid values: `60` and `300`.
func (o ScalingPolicyOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// Name of a policy used to define a reaction when an alarm is triggered.
func (o ScalingPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// ID of a scaling group.
func (o ScalingPolicyOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
func (o ScalingPolicyOutput) Statistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.Statistic }).(pulumi.StringPtrOutput)
}

// Alarm threshold.
func (o ScalingPolicyOutput) Threshold() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.IntOutput { return v.Threshold }).(pulumi.IntOutput)
}

type ScalingPolicyArrayOutput struct{ *pulumi.OutputState }

func (ScalingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyArrayOutput) ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput {
	return o
}

func (o ScalingPolicyArrayOutput) ToScalingPolicyArrayOutputWithContext(ctx context.Context) ScalingPolicyArrayOutput {
	return o
}

func (o ScalingPolicyArrayOutput) Index(i pulumi.IntInput) ScalingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingPolicy {
		return vs[0].([]*ScalingPolicy)[vs[1].(int)]
	}).(ScalingPolicyOutput)
}

type ScalingPolicyMapOutput struct{ *pulumi.OutputState }

func (ScalingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyMapOutput) ToScalingPolicyMapOutput() ScalingPolicyMapOutput {
	return o
}

func (o ScalingPolicyMapOutput) ToScalingPolicyMapOutputWithContext(ctx context.Context) ScalingPolicyMapOutput {
	return o
}

func (o ScalingPolicyMapOutput) MapIndex(k pulumi.StringInput) ScalingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingPolicy {
		return vs[0].(map[string]*ScalingPolicy)[vs[1].(string)]
	}).(ScalingPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyInput)(nil)).Elem(), &ScalingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyArrayInput)(nil)).Elem(), ScalingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyMapInput)(nil)).Elem(), ScalingPolicyMap{})
	pulumi.RegisterOutputType(ScalingPolicyOutput{})
	pulumi.RegisterOutputType(ScalingPolicyArrayOutput{})
	pulumi.RegisterOutputType(ScalingPolicyMapOutput{})
}
