// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a pts cronJob
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Pts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Pts.NewCronJob(ctx, "cronJob", &Pts.CronJobArgs{
// 			CronExpression: pulumi.String("* 1 * * *"),
// 			FrequencyType:  pulumi.Int(2),
// 			JobOwner:       pulumi.String("userName"),
// 			Note:           pulumi.String("desc"),
// 			NoticeId:       pulumi.String("notice-vp6i38jt"),
// 			ProjectId:      pulumi.String("project-7qkzxhea"),
// 			ScenarioId:     pulumi.String("scenario-c22lqb1w"),
// 			ScenarioName:   pulumi.String("pts-js(2022-11-10 21:53:53)"),
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
// pts cron_job can be imported using the projectId#cronJobId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Pts/cronJob:CronJob cron_job project-7qkzxhea#scenario-c22lqb1w
// ```
type CronJob struct {
	pulumi.CustomResourceState

	// Reason for suspension.
	AbortReason pulumi.IntOutput `pulumi:"abortReason"`
	// App ID.
	AppId pulumi.IntOutput `pulumi:"appId"`
	// Creation time; type: Timestamp ISO8601.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
	CronExpression pulumi.StringOutput `pulumi:"cronExpression"`
	// End Time; type: Timestamp ISO8601.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
	FrequencyType pulumi.IntOutput `pulumi:"frequencyType"`
	// Job Owner.
	JobOwner pulumi.StringOutput `pulumi:"jobOwner"`
	// Cron Job Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Note.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// Notice ID.
	NoticeId pulumi.StringPtrOutput `pulumi:"noticeId"`
	// Project Id.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Scenario Id.
	ScenarioId pulumi.StringOutput `pulumi:"scenarioId"`
	// Scenario Name.
	ScenarioName pulumi.StringOutput `pulumi:"scenarioName"`
	// Scheduled task status.
	Status pulumi.IntOutput `pulumi:"status"`
	// Sub-user ID.
	SubAccountUin pulumi.StringOutput `pulumi:"subAccountUin"`
	// User ID.
	Uin pulumi.StringOutput `pulumi:"uin"`
	// Update time; type: Timestamp ISO8601.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCronJob registers a new resource with the given unique name, arguments, and options.
func NewCronJob(ctx *pulumi.Context,
	name string, args *CronJobArgs, opts ...pulumi.ResourceOption) (*CronJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExpression == nil {
		return nil, errors.New("invalid value for required argument 'CronExpression'")
	}
	if args.FrequencyType == nil {
		return nil, errors.New("invalid value for required argument 'FrequencyType'")
	}
	if args.JobOwner == nil {
		return nil, errors.New("invalid value for required argument 'JobOwner'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ScenarioId == nil {
		return nil, errors.New("invalid value for required argument 'ScenarioId'")
	}
	if args.ScenarioName == nil {
		return nil, errors.New("invalid value for required argument 'ScenarioName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CronJob
	err := ctx.RegisterResource("tencentcloud:Pts/cronJob:CronJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJob gets an existing CronJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobState, opts ...pulumi.ResourceOption) (*CronJob, error) {
	var resource CronJob
	err := ctx.ReadResource("tencentcloud:Pts/cronJob:CronJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJob resources.
type cronJobState struct {
	// Reason for suspension.
	AbortReason *int `pulumi:"abortReason"`
	// App ID.
	AppId *int `pulumi:"appId"`
	// Creation time; type: Timestamp ISO8601.
	CreatedAt *string `pulumi:"createdAt"`
	// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
	CronExpression *string `pulumi:"cronExpression"`
	// End Time; type: Timestamp ISO8601.
	EndTime *string `pulumi:"endTime"`
	// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
	FrequencyType *int `pulumi:"frequencyType"`
	// Job Owner.
	JobOwner *string `pulumi:"jobOwner"`
	// Cron Job Name.
	Name *string `pulumi:"name"`
	// Note.
	Note *string `pulumi:"note"`
	// Notice ID.
	NoticeId *string `pulumi:"noticeId"`
	// Project Id.
	ProjectId *string `pulumi:"projectId"`
	// Scenario Id.
	ScenarioId *string `pulumi:"scenarioId"`
	// Scenario Name.
	ScenarioName *string `pulumi:"scenarioName"`
	// Scheduled task status.
	Status *int `pulumi:"status"`
	// Sub-user ID.
	SubAccountUin *string `pulumi:"subAccountUin"`
	// User ID.
	Uin *string `pulumi:"uin"`
	// Update time; type: Timestamp ISO8601.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CronJobState struct {
	// Reason for suspension.
	AbortReason pulumi.IntPtrInput
	// App ID.
	AppId pulumi.IntPtrInput
	// Creation time; type: Timestamp ISO8601.
	CreatedAt pulumi.StringPtrInput
	// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
	CronExpression pulumi.StringPtrInput
	// End Time; type: Timestamp ISO8601.
	EndTime pulumi.StringPtrInput
	// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
	FrequencyType pulumi.IntPtrInput
	// Job Owner.
	JobOwner pulumi.StringPtrInput
	// Cron Job Name.
	Name pulumi.StringPtrInput
	// Note.
	Note pulumi.StringPtrInput
	// Notice ID.
	NoticeId pulumi.StringPtrInput
	// Project Id.
	ProjectId pulumi.StringPtrInput
	// Scenario Id.
	ScenarioId pulumi.StringPtrInput
	// Scenario Name.
	ScenarioName pulumi.StringPtrInput
	// Scheduled task status.
	Status pulumi.IntPtrInput
	// Sub-user ID.
	SubAccountUin pulumi.StringPtrInput
	// User ID.
	Uin pulumi.StringPtrInput
	// Update time; type: Timestamp ISO8601.
	UpdatedAt pulumi.StringPtrInput
}

func (CronJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobState)(nil)).Elem()
}

type cronJobArgs struct {
	// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
	CronExpression string `pulumi:"cronExpression"`
	// End Time; type: Timestamp ISO8601.
	EndTime *string `pulumi:"endTime"`
	// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
	FrequencyType int `pulumi:"frequencyType"`
	// Job Owner.
	JobOwner string `pulumi:"jobOwner"`
	// Cron Job Name.
	Name *string `pulumi:"name"`
	// Note.
	Note *string `pulumi:"note"`
	// Notice ID.
	NoticeId *string `pulumi:"noticeId"`
	// Project Id.
	ProjectId string `pulumi:"projectId"`
	// Scenario Id.
	ScenarioId string `pulumi:"scenarioId"`
	// Scenario Name.
	ScenarioName string `pulumi:"scenarioName"`
}

// The set of arguments for constructing a CronJob resource.
type CronJobArgs struct {
	// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
	CronExpression pulumi.StringInput
	// End Time; type: Timestamp ISO8601.
	EndTime pulumi.StringPtrInput
	// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
	FrequencyType pulumi.IntInput
	// Job Owner.
	JobOwner pulumi.StringInput
	// Cron Job Name.
	Name pulumi.StringPtrInput
	// Note.
	Note pulumi.StringPtrInput
	// Notice ID.
	NoticeId pulumi.StringPtrInput
	// Project Id.
	ProjectId pulumi.StringInput
	// Scenario Id.
	ScenarioId pulumi.StringInput
	// Scenario Name.
	ScenarioName pulumi.StringInput
}

func (CronJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobArgs)(nil)).Elem()
}

type CronJobInput interface {
	pulumi.Input

	ToCronJobOutput() CronJobOutput
	ToCronJobOutputWithContext(ctx context.Context) CronJobOutput
}

func (*CronJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil)).Elem()
}

func (i *CronJob) ToCronJobOutput() CronJobOutput {
	return i.ToCronJobOutputWithContext(context.Background())
}

func (i *CronJob) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobOutput)
}

// CronJobArrayInput is an input type that accepts CronJobArray and CronJobArrayOutput values.
// You can construct a concrete instance of `CronJobArrayInput` via:
//
//          CronJobArray{ CronJobArgs{...} }
type CronJobArrayInput interface {
	pulumi.Input

	ToCronJobArrayOutput() CronJobArrayOutput
	ToCronJobArrayOutputWithContext(context.Context) CronJobArrayOutput
}

type CronJobArray []CronJobInput

func (CronJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJob)(nil)).Elem()
}

func (i CronJobArray) ToCronJobArrayOutput() CronJobArrayOutput {
	return i.ToCronJobArrayOutputWithContext(context.Background())
}

func (i CronJobArray) ToCronJobArrayOutputWithContext(ctx context.Context) CronJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobArrayOutput)
}

// CronJobMapInput is an input type that accepts CronJobMap and CronJobMapOutput values.
// You can construct a concrete instance of `CronJobMapInput` via:
//
//          CronJobMap{ "key": CronJobArgs{...} }
type CronJobMapInput interface {
	pulumi.Input

	ToCronJobMapOutput() CronJobMapOutput
	ToCronJobMapOutputWithContext(context.Context) CronJobMapOutput
}

type CronJobMap map[string]CronJobInput

func (CronJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJob)(nil)).Elem()
}

func (i CronJobMap) ToCronJobMapOutput() CronJobMapOutput {
	return i.ToCronJobMapOutputWithContext(context.Background())
}

func (i CronJobMap) ToCronJobMapOutputWithContext(ctx context.Context) CronJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobMapOutput)
}

type CronJobOutput struct{ *pulumi.OutputState }

func (CronJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil)).Elem()
}

func (o CronJobOutput) ToCronJobOutput() CronJobOutput {
	return o
}

func (o CronJobOutput) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return o
}

// Reason for suspension.
func (o CronJobOutput) AbortReason() pulumi.IntOutput {
	return o.ApplyT(func(v *CronJob) pulumi.IntOutput { return v.AbortReason }).(pulumi.IntOutput)
}

// App ID.
func (o CronJobOutput) AppId() pulumi.IntOutput {
	return o.ApplyT(func(v *CronJob) pulumi.IntOutput { return v.AppId }).(pulumi.IntOutput)
}

// Creation time; type: Timestamp ISO8601.
func (o CronJobOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
func (o CronJobOutput) CronExpression() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.CronExpression }).(pulumi.StringOutput)
}

// End Time; type: Timestamp ISO8601.
func (o CronJobOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
func (o CronJobOutput) FrequencyType() pulumi.IntOutput {
	return o.ApplyT(func(v *CronJob) pulumi.IntOutput { return v.FrequencyType }).(pulumi.IntOutput)
}

// Job Owner.
func (o CronJobOutput) JobOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.JobOwner }).(pulumi.StringOutput)
}

// Cron Job Name.
func (o CronJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Note.
func (o CronJobOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// Notice ID.
func (o CronJobOutput) NoticeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.NoticeId }).(pulumi.StringPtrOutput)
}

// Project Id.
func (o CronJobOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Scenario Id.
func (o CronJobOutput) ScenarioId() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.ScenarioId }).(pulumi.StringOutput)
}

// Scenario Name.
func (o CronJobOutput) ScenarioName() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.ScenarioName }).(pulumi.StringOutput)
}

// Scheduled task status.
func (o CronJobOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *CronJob) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Sub-user ID.
func (o CronJobOutput) SubAccountUin() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.SubAccountUin }).(pulumi.StringOutput)
}

// User ID.
func (o CronJobOutput) Uin() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.Uin }).(pulumi.StringOutput)
}

// Update time; type: Timestamp ISO8601.
func (o CronJobOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type CronJobArrayOutput struct{ *pulumi.OutputState }

func (CronJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJob)(nil)).Elem()
}

func (o CronJobArrayOutput) ToCronJobArrayOutput() CronJobArrayOutput {
	return o
}

func (o CronJobArrayOutput) ToCronJobArrayOutputWithContext(ctx context.Context) CronJobArrayOutput {
	return o
}

func (o CronJobArrayOutput) Index(i pulumi.IntInput) CronJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CronJob {
		return vs[0].([]*CronJob)[vs[1].(int)]
	}).(CronJobOutput)
}

type CronJobMapOutput struct{ *pulumi.OutputState }

func (CronJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJob)(nil)).Elem()
}

func (o CronJobMapOutput) ToCronJobMapOutput() CronJobMapOutput {
	return o
}

func (o CronJobMapOutput) ToCronJobMapOutputWithContext(ctx context.Context) CronJobMapOutput {
	return o
}

func (o CronJobMapOutput) MapIndex(k pulumi.StringInput) CronJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CronJob {
		return vs[0].(map[string]*CronJob)[vs[1].(string)]
	}).(CronJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobInput)(nil)).Elem(), &CronJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobArrayInput)(nil)).Elem(), CronJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobMapInput)(nil)).Elem(), CronJobMap{})
	pulumi.RegisterOutputType(CronJobOutput{})
	pulumi.RegisterOutputType(CronJobArrayOutput{})
	pulumi.RegisterOutputType(CronJobMapOutput{})
}
