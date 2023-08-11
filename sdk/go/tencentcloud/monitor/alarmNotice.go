// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a alarm notice resource for monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Monitor.NewAlarmNotice(ctx, "foo", &Monitor.AlarmNoticeArgs{
// 			NoticeLanguage: pulumi.String("zh-CN"),
// 			NoticeType:     pulumi.String("ALL"),
// 			UrlNotices: monitor.AlarmNoticeUrlNoticeArray{
// 				&monitor.AlarmNoticeUrlNoticeArgs{
// 					EndTime:   pulumi.Int(0),
// 					StartTime: pulumi.Int(1),
// 					Url:       pulumi.String("https://www.mytest.com/validate"),
// 					Weekdays: pulumi.IntArray{
// 						pulumi.Int(1),
// 						pulumi.Int(2),
// 						pulumi.Int(3),
// 						pulumi.Int(4),
// 						pulumi.Int(5),
// 						pulumi.Int(6),
// 						pulumi.Int(7),
// 					},
// 				},
// 			},
// 			UserNotices: monitor.AlarmNoticeUserNoticeArray{
// 				&monitor.AlarmNoticeUserNoticeArgs{
// 					EndTime:               pulumi.Int(1),
// 					GroupIds:              pulumi.IntArray{},
// 					NeedPhoneArriveNotice: pulumi.Int(1),
// 					NoticeWays: pulumi.StringArray{
// 						pulumi.String("SMS"),
// 						pulumi.String("EMAIL"),
// 					},
// 					PhoneCallType:       pulumi.String("CIRCLE"),
// 					PhoneCircleInterval: pulumi.Int(50),
// 					PhoneCircleTimes:    pulumi.Int(2),
// 					PhoneInnerInterval:  pulumi.Int(60),
// 					PhoneOrders: pulumi.IntArray{
// 						pulumi.Int(10001),
// 					},
// 					ReceiverType: pulumi.String("USER"),
// 					StartTime:    pulumi.Int(0),
// 					UserIds: pulumi.IntArray{
// 						pulumi.Int(10001),
// 					},
// 					Weekdays: pulumi.IntArray{
// 						pulumi.Int(1),
// 						pulumi.Int(2),
// 						pulumi.Int(3),
// 						pulumi.Int(4),
// 						pulumi.Int(5),
// 						pulumi.Int(6),
// 						pulumi.Int(7),
// 					},
// 				},
// 			},
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
// Monitor Alarm Notice can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Monitor/alarmNotice:AlarmNotice import-test noticeId
// ```
type AlarmNotice struct {
	pulumi.CustomResourceState

	// Amp consumer ID.
	AmpConsumerId pulumi.StringOutput `pulumi:"ampConsumerId"`
	// A maximum of one alarm notification can be pushed to the CLS service.
	ClsNotices AlarmNoticeClsNoticeArrayOutput `pulumi:"clsNotices"`
	// Whether it is the system default notification template 0=No 1=Yes.
	IsPreset pulumi.IntOutput `pulumi:"isPreset"`
	// Notification template name within 60.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification language zh-CN=Chinese en-US=English.
	NoticeLanguage pulumi.StringOutput `pulumi:"noticeLanguage"`
	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	NoticeType pulumi.StringOutput `pulumi:"noticeType"`
	// List of alarm policy IDs bound to the alarm notification template.
	PolicyIds pulumi.StringArrayOutput `pulumi:"policyIds"`
	// Last modified time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Last Modified By.
	UpdatedBy pulumi.StringOutput `pulumi:"updatedBy"`
	// The maximum number of callback notifications is 3.
	UrlNotices AlarmNoticeUrlNoticeArrayOutput `pulumi:"urlNotices"`
	// Alarm notification template list.(At most five).
	UserNotices AlarmNoticeUserNoticeArrayOutput `pulumi:"userNotices"`
}

// NewAlarmNotice registers a new resource with the given unique name, arguments, and options.
func NewAlarmNotice(ctx *pulumi.Context,
	name string, args *AlarmNoticeArgs, opts ...pulumi.ResourceOption) (*AlarmNotice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NoticeLanguage == nil {
		return nil, errors.New("invalid value for required argument 'NoticeLanguage'")
	}
	if args.NoticeType == nil {
		return nil, errors.New("invalid value for required argument 'NoticeType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AlarmNotice
	err := ctx.RegisterResource("tencentcloud:Monitor/alarmNotice:AlarmNotice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarmNotice gets an existing AlarmNotice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarmNotice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmNoticeState, opts ...pulumi.ResourceOption) (*AlarmNotice, error) {
	var resource AlarmNotice
	err := ctx.ReadResource("tencentcloud:Monitor/alarmNotice:AlarmNotice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlarmNotice resources.
type alarmNoticeState struct {
	// Amp consumer ID.
	AmpConsumerId *string `pulumi:"ampConsumerId"`
	// A maximum of one alarm notification can be pushed to the CLS service.
	ClsNotices []AlarmNoticeClsNotice `pulumi:"clsNotices"`
	// Whether it is the system default notification template 0=No 1=Yes.
	IsPreset *int `pulumi:"isPreset"`
	// Notification template name within 60.
	Name *string `pulumi:"name"`
	// Notification language zh-CN=Chinese en-US=English.
	NoticeLanguage *string `pulumi:"noticeLanguage"`
	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	NoticeType *string `pulumi:"noticeType"`
	// List of alarm policy IDs bound to the alarm notification template.
	PolicyIds []string `pulumi:"policyIds"`
	// Last modified time.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Last Modified By.
	UpdatedBy *string `pulumi:"updatedBy"`
	// The maximum number of callback notifications is 3.
	UrlNotices []AlarmNoticeUrlNotice `pulumi:"urlNotices"`
	// Alarm notification template list.(At most five).
	UserNotices []AlarmNoticeUserNotice `pulumi:"userNotices"`
}

type AlarmNoticeState struct {
	// Amp consumer ID.
	AmpConsumerId pulumi.StringPtrInput
	// A maximum of one alarm notification can be pushed to the CLS service.
	ClsNotices AlarmNoticeClsNoticeArrayInput
	// Whether it is the system default notification template 0=No 1=Yes.
	IsPreset pulumi.IntPtrInput
	// Notification template name within 60.
	Name pulumi.StringPtrInput
	// Notification language zh-CN=Chinese en-US=English.
	NoticeLanguage pulumi.StringPtrInput
	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	NoticeType pulumi.StringPtrInput
	// List of alarm policy IDs bound to the alarm notification template.
	PolicyIds pulumi.StringArrayInput
	// Last modified time.
	UpdatedAt pulumi.StringPtrInput
	// Last Modified By.
	UpdatedBy pulumi.StringPtrInput
	// The maximum number of callback notifications is 3.
	UrlNotices AlarmNoticeUrlNoticeArrayInput
	// Alarm notification template list.(At most five).
	UserNotices AlarmNoticeUserNoticeArrayInput
}

func (AlarmNoticeState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmNoticeState)(nil)).Elem()
}

type alarmNoticeArgs struct {
	// A maximum of one alarm notification can be pushed to the CLS service.
	ClsNotices []AlarmNoticeClsNotice `pulumi:"clsNotices"`
	// Notification template name within 60.
	Name *string `pulumi:"name"`
	// Notification language zh-CN=Chinese en-US=English.
	NoticeLanguage string `pulumi:"noticeLanguage"`
	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	NoticeType string `pulumi:"noticeType"`
	// The maximum number of callback notifications is 3.
	UrlNotices []AlarmNoticeUrlNotice `pulumi:"urlNotices"`
	// Alarm notification template list.(At most five).
	UserNotices []AlarmNoticeUserNotice `pulumi:"userNotices"`
}

// The set of arguments for constructing a AlarmNotice resource.
type AlarmNoticeArgs struct {
	// A maximum of one alarm notification can be pushed to the CLS service.
	ClsNotices AlarmNoticeClsNoticeArrayInput
	// Notification template name within 60.
	Name pulumi.StringPtrInput
	// Notification language zh-CN=Chinese en-US=English.
	NoticeLanguage pulumi.StringInput
	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	NoticeType pulumi.StringInput
	// The maximum number of callback notifications is 3.
	UrlNotices AlarmNoticeUrlNoticeArrayInput
	// Alarm notification template list.(At most five).
	UserNotices AlarmNoticeUserNoticeArrayInput
}

func (AlarmNoticeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmNoticeArgs)(nil)).Elem()
}

type AlarmNoticeInput interface {
	pulumi.Input

	ToAlarmNoticeOutput() AlarmNoticeOutput
	ToAlarmNoticeOutputWithContext(ctx context.Context) AlarmNoticeOutput
}

func (*AlarmNotice) ElementType() reflect.Type {
	return reflect.TypeOf((**AlarmNotice)(nil)).Elem()
}

func (i *AlarmNotice) ToAlarmNoticeOutput() AlarmNoticeOutput {
	return i.ToAlarmNoticeOutputWithContext(context.Background())
}

func (i *AlarmNotice) ToAlarmNoticeOutputWithContext(ctx context.Context) AlarmNoticeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmNoticeOutput)
}

// AlarmNoticeArrayInput is an input type that accepts AlarmNoticeArray and AlarmNoticeArrayOutput values.
// You can construct a concrete instance of `AlarmNoticeArrayInput` via:
//
//          AlarmNoticeArray{ AlarmNoticeArgs{...} }
type AlarmNoticeArrayInput interface {
	pulumi.Input

	ToAlarmNoticeArrayOutput() AlarmNoticeArrayOutput
	ToAlarmNoticeArrayOutputWithContext(context.Context) AlarmNoticeArrayOutput
}

type AlarmNoticeArray []AlarmNoticeInput

func (AlarmNoticeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlarmNotice)(nil)).Elem()
}

func (i AlarmNoticeArray) ToAlarmNoticeArrayOutput() AlarmNoticeArrayOutput {
	return i.ToAlarmNoticeArrayOutputWithContext(context.Background())
}

func (i AlarmNoticeArray) ToAlarmNoticeArrayOutputWithContext(ctx context.Context) AlarmNoticeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmNoticeArrayOutput)
}

// AlarmNoticeMapInput is an input type that accepts AlarmNoticeMap and AlarmNoticeMapOutput values.
// You can construct a concrete instance of `AlarmNoticeMapInput` via:
//
//          AlarmNoticeMap{ "key": AlarmNoticeArgs{...} }
type AlarmNoticeMapInput interface {
	pulumi.Input

	ToAlarmNoticeMapOutput() AlarmNoticeMapOutput
	ToAlarmNoticeMapOutputWithContext(context.Context) AlarmNoticeMapOutput
}

type AlarmNoticeMap map[string]AlarmNoticeInput

func (AlarmNoticeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlarmNotice)(nil)).Elem()
}

func (i AlarmNoticeMap) ToAlarmNoticeMapOutput() AlarmNoticeMapOutput {
	return i.ToAlarmNoticeMapOutputWithContext(context.Background())
}

func (i AlarmNoticeMap) ToAlarmNoticeMapOutputWithContext(ctx context.Context) AlarmNoticeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmNoticeMapOutput)
}

type AlarmNoticeOutput struct{ *pulumi.OutputState }

func (AlarmNoticeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlarmNotice)(nil)).Elem()
}

func (o AlarmNoticeOutput) ToAlarmNoticeOutput() AlarmNoticeOutput {
	return o
}

func (o AlarmNoticeOutput) ToAlarmNoticeOutputWithContext(ctx context.Context) AlarmNoticeOutput {
	return o
}

// Amp consumer ID.
func (o AlarmNoticeOutput) AmpConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.AmpConsumerId }).(pulumi.StringOutput)
}

// A maximum of one alarm notification can be pushed to the CLS service.
func (o AlarmNoticeOutput) ClsNotices() AlarmNoticeClsNoticeArrayOutput {
	return o.ApplyT(func(v *AlarmNotice) AlarmNoticeClsNoticeArrayOutput { return v.ClsNotices }).(AlarmNoticeClsNoticeArrayOutput)
}

// Whether it is the system default notification template 0=No 1=Yes.
func (o AlarmNoticeOutput) IsPreset() pulumi.IntOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.IntOutput { return v.IsPreset }).(pulumi.IntOutput)
}

// Notification template name within 60.
func (o AlarmNoticeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification language zh-CN=Chinese en-US=English.
func (o AlarmNoticeOutput) NoticeLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.NoticeLanguage }).(pulumi.StringOutput)
}

// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
func (o AlarmNoticeOutput) NoticeType() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.NoticeType }).(pulumi.StringOutput)
}

// List of alarm policy IDs bound to the alarm notification template.
func (o AlarmNoticeOutput) PolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringArrayOutput { return v.PolicyIds }).(pulumi.StringArrayOutput)
}

// Last modified time.
func (o AlarmNoticeOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Last Modified By.
func (o AlarmNoticeOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *AlarmNotice) pulumi.StringOutput { return v.UpdatedBy }).(pulumi.StringOutput)
}

// The maximum number of callback notifications is 3.
func (o AlarmNoticeOutput) UrlNotices() AlarmNoticeUrlNoticeArrayOutput {
	return o.ApplyT(func(v *AlarmNotice) AlarmNoticeUrlNoticeArrayOutput { return v.UrlNotices }).(AlarmNoticeUrlNoticeArrayOutput)
}

// Alarm notification template list.(At most five).
func (o AlarmNoticeOutput) UserNotices() AlarmNoticeUserNoticeArrayOutput {
	return o.ApplyT(func(v *AlarmNotice) AlarmNoticeUserNoticeArrayOutput { return v.UserNotices }).(AlarmNoticeUserNoticeArrayOutput)
}

type AlarmNoticeArrayOutput struct{ *pulumi.OutputState }

func (AlarmNoticeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlarmNotice)(nil)).Elem()
}

func (o AlarmNoticeArrayOutput) ToAlarmNoticeArrayOutput() AlarmNoticeArrayOutput {
	return o
}

func (o AlarmNoticeArrayOutput) ToAlarmNoticeArrayOutputWithContext(ctx context.Context) AlarmNoticeArrayOutput {
	return o
}

func (o AlarmNoticeArrayOutput) Index(i pulumi.IntInput) AlarmNoticeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlarmNotice {
		return vs[0].([]*AlarmNotice)[vs[1].(int)]
	}).(AlarmNoticeOutput)
}

type AlarmNoticeMapOutput struct{ *pulumi.OutputState }

func (AlarmNoticeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlarmNotice)(nil)).Elem()
}

func (o AlarmNoticeMapOutput) ToAlarmNoticeMapOutput() AlarmNoticeMapOutput {
	return o
}

func (o AlarmNoticeMapOutput) ToAlarmNoticeMapOutputWithContext(ctx context.Context) AlarmNoticeMapOutput {
	return o
}

func (o AlarmNoticeMapOutput) MapIndex(k pulumi.StringInput) AlarmNoticeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlarmNotice {
		return vs[0].(map[string]*AlarmNotice)[vs[1].(string)]
	}).(AlarmNoticeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmNoticeInput)(nil)).Elem(), &AlarmNotice{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmNoticeArrayInput)(nil)).Elem(), AlarmNoticeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmNoticeMapInput)(nil)).Elem(), AlarmNoticeMap{})
	pulumi.RegisterOutputType(AlarmNoticeOutput{})
	pulumi.RegisterOutputType(AlarmNoticeArrayOutput{})
	pulumi.RegisterOutputType(AlarmNoticeMapOutput{})
}
