// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a cls alarm
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cls.NewAlarm(ctx, "example", &Cls.AlarmArgs{
//				AlarmLevel: pulumi.Int(0),
//				AlarmNoticeIds: pulumi.StringArray{
//					pulumi.String("notice-0850756b-245d-4bc7-bb27-2a58fffc780b"),
//				},
//				AlarmPeriod: pulumi.Int(15),
//				AlarmTargets: cls.AlarmAlarmTargetArray{
//					&cls.AlarmAlarmTargetArgs{
//						EndTimeOffset:   pulumi.Int(0),
//						LogsetId:        pulumi.String("33aaf0ae-6163-411b-a415-9f27450f68db"),
//						Number:          pulumi.Int(1),
//						Query:           pulumi.String("status:>500 | select count(*) as errorCounts"),
//						StartTimeOffset: -15,
//						SyntaxRule:      pulumi.Int(1),
//						TopicId:         pulumi.String("88735a07-bea4-4985-8763-e9deb6da4fad"),
//					},
//				},
//				Analyses: cls.AlarmAnalysisArray{
//					&cls.AlarmAnalysisArgs{
//						ConfigInfos: cls.AlarmAnalysisConfigInfoArray{
//							&cls.AlarmAnalysisConfigInfoArgs{
//								Key:   pulumi.String("QueryIndex"),
//								Value: pulumi.String("1"),
//							},
//						},
//						Content: pulumi.String("__FILENAME__"),
//						Name:    pulumi.String("terraform"),
//						Type:    pulumi.String("field"),
//					},
//				},
//				Condition:       pulumi.String("test"),
//				MessageTemplate: pulumi.String("{{.Label}}"),
//				MonitorTime: &cls.AlarmMonitorTimeArgs{
//					Time: pulumi.Int(1),
//					Type: pulumi.String("Period"),
//				},
//				Status: pulumi.Bool(true),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//				TriggerCount: pulumi.Int(1),
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
// cls alarm can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Cls/alarm:Alarm example alarm-d8529662-e10f-440c-ba80-50f3dcf215a3
// ```
type Alarm struct {
	pulumi.CustomResourceState

	// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
	AlarmLevel pulumi.IntOutput `pulumi:"alarmLevel"`
	// list of alarm notice id.
	AlarmNoticeIds pulumi.StringArrayOutput `pulumi:"alarmNoticeIds"`
	// alarm repeat cycle.
	AlarmPeriod pulumi.IntOutput `pulumi:"alarmPeriod"`
	// list of alarm target.
	AlarmTargets AlarmAlarmTargetArrayOutput `pulumi:"alarmTargets"`
	// multidimensional analysis.
	Analyses AlarmAnalysisArrayOutput `pulumi:"analyses"`
	// user define callback.
	CallBack AlarmCallBackOutput `pulumi:"callBack"`
	// triggering conditions.
	Condition pulumi.StringOutput `pulumi:"condition"`
	// user define alarm notice.
	MessageTemplate pulumi.StringPtrOutput `pulumi:"messageTemplate"`
	// monitor task execution time.
	MonitorTime AlarmMonitorTimeOutput `pulumi:"monitorTime"`
	// log alarm name.
	Name pulumi.StringOutput `pulumi:"name"`
	// whether to enable the alarm policy.
	Status pulumi.BoolOutput `pulumi:"status"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// continuous cycle.
	TriggerCount pulumi.IntOutput `pulumi:"triggerCount"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlarmNoticeIds == nil {
		return nil, errors.New("invalid value for required argument 'AlarmNoticeIds'")
	}
	if args.AlarmPeriod == nil {
		return nil, errors.New("invalid value for required argument 'AlarmPeriod'")
	}
	if args.AlarmTargets == nil {
		return nil, errors.New("invalid value for required argument 'AlarmTargets'")
	}
	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	if args.MonitorTime == nil {
		return nil, errors.New("invalid value for required argument 'MonitorTime'")
	}
	if args.TriggerCount == nil {
		return nil, errors.New("invalid value for required argument 'TriggerCount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alarm
	err := ctx.RegisterResource("tencentcloud:Cls/alarm:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("tencentcloud:Cls/alarm:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
	// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
	AlarmLevel *int `pulumi:"alarmLevel"`
	// list of alarm notice id.
	AlarmNoticeIds []string `pulumi:"alarmNoticeIds"`
	// alarm repeat cycle.
	AlarmPeriod *int `pulumi:"alarmPeriod"`
	// list of alarm target.
	AlarmTargets []AlarmAlarmTarget `pulumi:"alarmTargets"`
	// multidimensional analysis.
	Analyses []AlarmAnalysis `pulumi:"analyses"`
	// user define callback.
	CallBack *AlarmCallBack `pulumi:"callBack"`
	// triggering conditions.
	Condition *string `pulumi:"condition"`
	// user define alarm notice.
	MessageTemplate *string `pulumi:"messageTemplate"`
	// monitor task execution time.
	MonitorTime *AlarmMonitorTime `pulumi:"monitorTime"`
	// log alarm name.
	Name *string `pulumi:"name"`
	// whether to enable the alarm policy.
	Status *bool `pulumi:"status"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// continuous cycle.
	TriggerCount *int `pulumi:"triggerCount"`
}

type AlarmState struct {
	// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
	AlarmLevel pulumi.IntPtrInput
	// list of alarm notice id.
	AlarmNoticeIds pulumi.StringArrayInput
	// alarm repeat cycle.
	AlarmPeriod pulumi.IntPtrInput
	// list of alarm target.
	AlarmTargets AlarmAlarmTargetArrayInput
	// multidimensional analysis.
	Analyses AlarmAnalysisArrayInput
	// user define callback.
	CallBack AlarmCallBackPtrInput
	// triggering conditions.
	Condition pulumi.StringPtrInput
	// user define alarm notice.
	MessageTemplate pulumi.StringPtrInput
	// monitor task execution time.
	MonitorTime AlarmMonitorTimePtrInput
	// log alarm name.
	Name pulumi.StringPtrInput
	// whether to enable the alarm policy.
	Status pulumi.BoolPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// continuous cycle.
	TriggerCount pulumi.IntPtrInput
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
	AlarmLevel *int `pulumi:"alarmLevel"`
	// list of alarm notice id.
	AlarmNoticeIds []string `pulumi:"alarmNoticeIds"`
	// alarm repeat cycle.
	AlarmPeriod int `pulumi:"alarmPeriod"`
	// list of alarm target.
	AlarmTargets []AlarmAlarmTarget `pulumi:"alarmTargets"`
	// multidimensional analysis.
	Analyses []AlarmAnalysis `pulumi:"analyses"`
	// user define callback.
	CallBack *AlarmCallBack `pulumi:"callBack"`
	// triggering conditions.
	Condition string `pulumi:"condition"`
	// user define alarm notice.
	MessageTemplate *string `pulumi:"messageTemplate"`
	// monitor task execution time.
	MonitorTime AlarmMonitorTime `pulumi:"monitorTime"`
	// log alarm name.
	Name *string `pulumi:"name"`
	// whether to enable the alarm policy.
	Status *bool `pulumi:"status"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// continuous cycle.
	TriggerCount int `pulumi:"triggerCount"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
	AlarmLevel pulumi.IntPtrInput
	// list of alarm notice id.
	AlarmNoticeIds pulumi.StringArrayInput
	// alarm repeat cycle.
	AlarmPeriod pulumi.IntInput
	// list of alarm target.
	AlarmTargets AlarmAlarmTargetArrayInput
	// multidimensional analysis.
	Analyses AlarmAnalysisArrayInput
	// user define callback.
	CallBack AlarmCallBackPtrInput
	// triggering conditions.
	Condition pulumi.StringInput
	// user define alarm notice.
	MessageTemplate pulumi.StringPtrInput
	// monitor task execution time.
	MonitorTime AlarmMonitorTimeInput
	// log alarm name.
	Name pulumi.StringPtrInput
	// whether to enable the alarm policy.
	Status pulumi.BoolPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// continuous cycle.
	TriggerCount pulumi.IntInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}

type AlarmInput interface {
	pulumi.Input

	ToAlarmOutput() AlarmOutput
	ToAlarmOutputWithContext(ctx context.Context) AlarmOutput
}

func (*Alarm) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

// AlarmArrayInput is an input type that accepts AlarmArray and AlarmArrayOutput values.
// You can construct a concrete instance of `AlarmArrayInput` via:
//
//	AlarmArray{ AlarmArgs{...} }
type AlarmArrayInput interface {
	pulumi.Input

	ToAlarmArrayOutput() AlarmArrayOutput
	ToAlarmArrayOutputWithContext(context.Context) AlarmArrayOutput
}

type AlarmArray []AlarmInput

func (AlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (i AlarmArray) ToAlarmArrayOutput() AlarmArrayOutput {
	return i.ToAlarmArrayOutputWithContext(context.Background())
}

func (i AlarmArray) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmArrayOutput)
}

// AlarmMapInput is an input type that accepts AlarmMap and AlarmMapOutput values.
// You can construct a concrete instance of `AlarmMapInput` via:
//
//	AlarmMap{ "key": AlarmArgs{...} }
type AlarmMapInput interface {
	pulumi.Input

	ToAlarmMapOutput() AlarmMapOutput
	ToAlarmMapOutputWithContext(context.Context) AlarmMapOutput
}

type AlarmMap map[string]AlarmInput

func (AlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (i AlarmMap) ToAlarmMapOutput() AlarmMapOutput {
	return i.ToAlarmMapOutputWithContext(context.Background())
}

func (i AlarmMap) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmMapOutput)
}

type AlarmOutput struct{ *pulumi.OutputState }

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

// Alarm level. 0: Warning; 1: Info; 2: Critical. Default is 0.
func (o AlarmOutput) AlarmLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntOutput { return v.AlarmLevel }).(pulumi.IntOutput)
}

// list of alarm notice id.
func (o AlarmOutput) AlarmNoticeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringArrayOutput { return v.AlarmNoticeIds }).(pulumi.StringArrayOutput)
}

// alarm repeat cycle.
func (o AlarmOutput) AlarmPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntOutput { return v.AlarmPeriod }).(pulumi.IntOutput)
}

// list of alarm target.
func (o AlarmOutput) AlarmTargets() AlarmAlarmTargetArrayOutput {
	return o.ApplyT(func(v *Alarm) AlarmAlarmTargetArrayOutput { return v.AlarmTargets }).(AlarmAlarmTargetArrayOutput)
}

// multidimensional analysis.
func (o AlarmOutput) Analyses() AlarmAnalysisArrayOutput {
	return o.ApplyT(func(v *Alarm) AlarmAnalysisArrayOutput { return v.Analyses }).(AlarmAnalysisArrayOutput)
}

// user define callback.
func (o AlarmOutput) CallBack() AlarmCallBackOutput {
	return o.ApplyT(func(v *Alarm) AlarmCallBackOutput { return v.CallBack }).(AlarmCallBackOutput)
}

// triggering conditions.
func (o AlarmOutput) Condition() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Condition }).(pulumi.StringOutput)
}

// user define alarm notice.
func (o AlarmOutput) MessageTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.MessageTemplate }).(pulumi.StringPtrOutput)
}

// monitor task execution time.
func (o AlarmOutput) MonitorTime() AlarmMonitorTimeOutput {
	return o.ApplyT(func(v *Alarm) AlarmMonitorTimeOutput { return v.MonitorTime }).(AlarmMonitorTimeOutput)
}

// log alarm name.
func (o AlarmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// whether to enable the alarm policy.
func (o AlarmOutput) Status() pulumi.BoolOutput {
	return o.ApplyT(func(v *Alarm) pulumi.BoolOutput { return v.Status }).(pulumi.BoolOutput)
}

// Tag description list.
func (o AlarmOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Alarm) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// continuous cycle.
func (o AlarmOutput) TriggerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntOutput { return v.TriggerCount }).(pulumi.IntOutput)
}

type AlarmArrayOutput struct{ *pulumi.OutputState }

func (AlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (o AlarmArrayOutput) ToAlarmArrayOutput() AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) Index(i pulumi.IntInput) AlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].([]*Alarm)[vs[1].(int)]
	}).(AlarmOutput)
}

type AlarmMapOutput struct{ *pulumi.OutputState }

func (AlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (o AlarmMapOutput) ToAlarmMapOutput() AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) MapIndex(k pulumi.StringInput) AlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].(map[string]*Alarm)[vs[1].(string)]
	}).(AlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmInput)(nil)).Elem(), &Alarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmArrayInput)(nil)).Elem(), AlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmMapInput)(nil)).Elem(), AlarmMap{})
	pulumi.RegisterOutputType(AlarmOutput{})
	pulumi.RegisterOutputType(AlarmArrayOutput{})
	pulumi.RegisterOutputType(AlarmMapOutput{})
}
