// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tsf task
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tsf.NewTask(ctx, "task", &Tsf.TaskArgs{
// 			AdvanceSettings: &tsf.TaskAdvanceSettingsArgs{
// 				SubTaskConcurrency: pulumi.Int(2),
// 			},
// 			ExecuteType:     pulumi.String("unicast"),
// 			GroupId:         pulumi.String("group-y8pnmoga"),
// 			RetryCount:      pulumi.Int(0),
// 			RetryInterval:   pulumi.Int(0),
// 			SuccessOperator: pulumi.String("GTE"),
// 			SuccessRatio:    pulumi.String("100"),
// 			TaskArgument:    pulumi.String("a=c"),
// 			TaskContent:     pulumi.String("/test"),
// 			TaskName:        pulumi.String("terraform-test"),
// 			TaskRule: &tsf.TaskTaskRuleArgs{
// 				Expression: pulumi.String("0 * 1 * * ? "),
// 				RuleType:   pulumi.String("Cron"),
// 			},
// 			TaskType: pulumi.String("java"),
// 			TimeOut:  pulumi.Int(60000),
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
// tsf task can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tsf/task:Task task task-y37eqq95
// ```
type Task struct {
	pulumi.CustomResourceState

	// advanced settings.
	AdvanceSettings TaskAdvanceSettingsOutput `pulumi:"advanceSettings"`
	// ID of the workflow to which it belongs.
	BelongFlowIds pulumi.StringArrayOutput `pulumi:"belongFlowIds"`
	// execution type, unicast/broadcast.
	ExecuteType pulumi.StringOutput `pulumi:"executeType"`
	// deployment group ID.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Program id list.
	ProgramIdLists pulumi.StringArrayOutput `pulumi:"programIdLists"`
	// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
	RetryCount pulumi.IntOutput `pulumi:"retryCount"`
	// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
	RetryInterval pulumi.IntOutput `pulumi:"retryInterval"`
	// Fragmentation parameters.
	ShardArguments TaskShardArgumentArrayOutput `pulumi:"shardArguments"`
	// number of shards.
	ShardCount pulumi.IntOutput `pulumi:"shardCount"`
	// the operator to judge the success of the task.
	SuccessOperator pulumi.StringOutput `pulumi:"successOperator"`
	// The threshold for judging the success rate of the task, such as 100.
	SuccessRatio pulumi.StringOutput `pulumi:"successRatio"`
	// task parameters, the length limit is 10000 characters.
	TaskArgument pulumi.StringOutput `pulumi:"taskArgument"`
	// task content, length limit 65536 bytes.
	TaskContent pulumi.StringOutput `pulumi:"taskContent"`
	// task ID.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// task history ID.
	TaskLogId pulumi.StringOutput `pulumi:"taskLogId"`
	// task name, task length 64 characters.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
	// trigger rule.
	TaskRule TaskTaskRuleOutput `pulumi:"taskRule"`
	// Whether to enable the task, ENABLED/DISABLED.
	TaskState pulumi.StringOutput `pulumi:"taskState"`
	// task type, java.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// task timeout, time unit ms.
	TimeOut pulumi.IntOutput `pulumi:"timeOut"`
	// trigger type.
	TriggerType pulumi.StringOutput `pulumi:"triggerType"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecuteType == nil {
		return nil, errors.New("invalid value for required argument 'ExecuteType'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.TaskContent == nil {
		return nil, errors.New("invalid value for required argument 'TaskContent'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	if args.TimeOut == nil {
		return nil, errors.New("invalid value for required argument 'TimeOut'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Task
	err := ctx.RegisterResource("tencentcloud:Tsf/task:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("tencentcloud:Tsf/task:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
	// advanced settings.
	AdvanceSettings *TaskAdvanceSettings `pulumi:"advanceSettings"`
	// ID of the workflow to which it belongs.
	BelongFlowIds []string `pulumi:"belongFlowIds"`
	// execution type, unicast/broadcast.
	ExecuteType *string `pulumi:"executeType"`
	// deployment group ID.
	GroupId *string `pulumi:"groupId"`
	// Program id list.
	ProgramIdLists []string `pulumi:"programIdLists"`
	// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
	RetryCount *int `pulumi:"retryCount"`
	// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
	RetryInterval *int `pulumi:"retryInterval"`
	// Fragmentation parameters.
	ShardArguments []TaskShardArgument `pulumi:"shardArguments"`
	// number of shards.
	ShardCount *int `pulumi:"shardCount"`
	// the operator to judge the success of the task.
	SuccessOperator *string `pulumi:"successOperator"`
	// The threshold for judging the success rate of the task, such as 100.
	SuccessRatio *string `pulumi:"successRatio"`
	// task parameters, the length limit is 10000 characters.
	TaskArgument *string `pulumi:"taskArgument"`
	// task content, length limit 65536 bytes.
	TaskContent *string `pulumi:"taskContent"`
	// task ID.
	TaskId *string `pulumi:"taskId"`
	// task history ID.
	TaskLogId *string `pulumi:"taskLogId"`
	// task name, task length 64 characters.
	TaskName *string `pulumi:"taskName"`
	// trigger rule.
	TaskRule *TaskTaskRule `pulumi:"taskRule"`
	// Whether to enable the task, ENABLED/DISABLED.
	TaskState *string `pulumi:"taskState"`
	// task type, java.
	TaskType *string `pulumi:"taskType"`
	// task timeout, time unit ms.
	TimeOut *int `pulumi:"timeOut"`
	// trigger type.
	TriggerType *string `pulumi:"triggerType"`
}

type TaskState struct {
	// advanced settings.
	AdvanceSettings TaskAdvanceSettingsPtrInput
	// ID of the workflow to which it belongs.
	BelongFlowIds pulumi.StringArrayInput
	// execution type, unicast/broadcast.
	ExecuteType pulumi.StringPtrInput
	// deployment group ID.
	GroupId pulumi.StringPtrInput
	// Program id list.
	ProgramIdLists pulumi.StringArrayInput
	// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
	RetryCount pulumi.IntPtrInput
	// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
	RetryInterval pulumi.IntPtrInput
	// Fragmentation parameters.
	ShardArguments TaskShardArgumentArrayInput
	// number of shards.
	ShardCount pulumi.IntPtrInput
	// the operator to judge the success of the task.
	SuccessOperator pulumi.StringPtrInput
	// The threshold for judging the success rate of the task, such as 100.
	SuccessRatio pulumi.StringPtrInput
	// task parameters, the length limit is 10000 characters.
	TaskArgument pulumi.StringPtrInput
	// task content, length limit 65536 bytes.
	TaskContent pulumi.StringPtrInput
	// task ID.
	TaskId pulumi.StringPtrInput
	// task history ID.
	TaskLogId pulumi.StringPtrInput
	// task name, task length 64 characters.
	TaskName pulumi.StringPtrInput
	// trigger rule.
	TaskRule TaskTaskRulePtrInput
	// Whether to enable the task, ENABLED/DISABLED.
	TaskState pulumi.StringPtrInput
	// task type, java.
	TaskType pulumi.StringPtrInput
	// task timeout, time unit ms.
	TimeOut pulumi.IntPtrInput
	// trigger type.
	TriggerType pulumi.StringPtrInput
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// advanced settings.
	AdvanceSettings *TaskAdvanceSettings `pulumi:"advanceSettings"`
	// execution type, unicast/broadcast.
	ExecuteType string `pulumi:"executeType"`
	// deployment group ID.
	GroupId string `pulumi:"groupId"`
	// Program id list.
	ProgramIdLists []string `pulumi:"programIdLists"`
	// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
	RetryCount *int `pulumi:"retryCount"`
	// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
	RetryInterval *int `pulumi:"retryInterval"`
	// Fragmentation parameters.
	ShardArguments []TaskShardArgument `pulumi:"shardArguments"`
	// number of shards.
	ShardCount *int `pulumi:"shardCount"`
	// the operator to judge the success of the task.
	SuccessOperator *string `pulumi:"successOperator"`
	// The threshold for judging the success rate of the task, such as 100.
	SuccessRatio *string `pulumi:"successRatio"`
	// task parameters, the length limit is 10000 characters.
	TaskArgument *string `pulumi:"taskArgument"`
	// task content, length limit 65536 bytes.
	TaskContent string `pulumi:"taskContent"`
	// task name, task length 64 characters.
	TaskName string `pulumi:"taskName"`
	// trigger rule.
	TaskRule *TaskTaskRule `pulumi:"taskRule"`
	// task type, java.
	TaskType string `pulumi:"taskType"`
	// task timeout, time unit ms.
	TimeOut int `pulumi:"timeOut"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// advanced settings.
	AdvanceSettings TaskAdvanceSettingsPtrInput
	// execution type, unicast/broadcast.
	ExecuteType pulumi.StringInput
	// deployment group ID.
	GroupId pulumi.StringInput
	// Program id list.
	ProgramIdLists pulumi.StringArrayInput
	// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
	RetryCount pulumi.IntPtrInput
	// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
	RetryInterval pulumi.IntPtrInput
	// Fragmentation parameters.
	ShardArguments TaskShardArgumentArrayInput
	// number of shards.
	ShardCount pulumi.IntPtrInput
	// the operator to judge the success of the task.
	SuccessOperator pulumi.StringPtrInput
	// The threshold for judging the success rate of the task, such as 100.
	SuccessRatio pulumi.StringPtrInput
	// task parameters, the length limit is 10000 characters.
	TaskArgument pulumi.StringPtrInput
	// task content, length limit 65536 bytes.
	TaskContent pulumi.StringInput
	// task name, task length 64 characters.
	TaskName pulumi.StringInput
	// trigger rule.
	TaskRule TaskTaskRulePtrInput
	// task type, java.
	TaskType pulumi.StringInput
	// task timeout, time unit ms.
	TimeOut pulumi.IntInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}

type TaskInput interface {
	pulumi.Input

	ToTaskOutput() TaskOutput
	ToTaskOutputWithContext(ctx context.Context) TaskOutput
}

func (*Task) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

// TaskArrayInput is an input type that accepts TaskArray and TaskArrayOutput values.
// You can construct a concrete instance of `TaskArrayInput` via:
//
//          TaskArray{ TaskArgs{...} }
type TaskArrayInput interface {
	pulumi.Input

	ToTaskArrayOutput() TaskArrayOutput
	ToTaskArrayOutputWithContext(context.Context) TaskArrayOutput
}

type TaskArray []TaskInput

func (TaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (i TaskArray) ToTaskArrayOutput() TaskArrayOutput {
	return i.ToTaskArrayOutputWithContext(context.Background())
}

func (i TaskArray) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskArrayOutput)
}

// TaskMapInput is an input type that accepts TaskMap and TaskMapOutput values.
// You can construct a concrete instance of `TaskMapInput` via:
//
//          TaskMap{ "key": TaskArgs{...} }
type TaskMapInput interface {
	pulumi.Input

	ToTaskMapOutput() TaskMapOutput
	ToTaskMapOutputWithContext(context.Context) TaskMapOutput
}

type TaskMap map[string]TaskInput

func (TaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (i TaskMap) ToTaskMapOutput() TaskMapOutput {
	return i.ToTaskMapOutputWithContext(context.Background())
}

func (i TaskMap) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskMapOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

// advanced settings.
func (o TaskOutput) AdvanceSettings() TaskAdvanceSettingsOutput {
	return o.ApplyT(func(v *Task) TaskAdvanceSettingsOutput { return v.AdvanceSettings }).(TaskAdvanceSettingsOutput)
}

// ID of the workflow to which it belongs.
func (o TaskOutput) BelongFlowIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Task) pulumi.StringArrayOutput { return v.BelongFlowIds }).(pulumi.StringArrayOutput)
}

// execution type, unicast/broadcast.
func (o TaskOutput) ExecuteType() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.ExecuteType }).(pulumi.StringOutput)
}

// deployment group ID.
func (o TaskOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Program id list.
func (o TaskOutput) ProgramIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Task) pulumi.StringArrayOutput { return v.ProgramIdLists }).(pulumi.StringArrayOutput)
}

// number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
func (o TaskOutput) RetryCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.RetryCount }).(pulumi.IntOutput)
}

// retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
func (o TaskOutput) RetryInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.RetryInterval }).(pulumi.IntOutput)
}

// Fragmentation parameters.
func (o TaskOutput) ShardArguments() TaskShardArgumentArrayOutput {
	return o.ApplyT(func(v *Task) TaskShardArgumentArrayOutput { return v.ShardArguments }).(TaskShardArgumentArrayOutput)
}

// number of shards.
func (o TaskOutput) ShardCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.ShardCount }).(pulumi.IntOutput)
}

// the operator to judge the success of the task.
func (o TaskOutput) SuccessOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.SuccessOperator }).(pulumi.StringOutput)
}

// The threshold for judging the success rate of the task, such as 100.
func (o TaskOutput) SuccessRatio() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.SuccessRatio }).(pulumi.StringOutput)
}

// task parameters, the length limit is 10000 characters.
func (o TaskOutput) TaskArgument() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskArgument }).(pulumi.StringOutput)
}

// task content, length limit 65536 bytes.
func (o TaskOutput) TaskContent() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskContent }).(pulumi.StringOutput)
}

// task ID.
func (o TaskOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// task history ID.
func (o TaskOutput) TaskLogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskLogId }).(pulumi.StringOutput)
}

// task name, task length 64 characters.
func (o TaskOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

// trigger rule.
func (o TaskOutput) TaskRule() TaskTaskRuleOutput {
	return o.ApplyT(func(v *Task) TaskTaskRuleOutput { return v.TaskRule }).(TaskTaskRuleOutput)
}

// Whether to enable the task, ENABLED/DISABLED.
func (o TaskOutput) TaskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskState }).(pulumi.StringOutput)
}

// task type, java.
func (o TaskOutput) TaskType() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TaskType }).(pulumi.StringOutput)
}

// task timeout, time unit ms.
func (o TaskOutput) TimeOut() pulumi.IntOutput {
	return o.ApplyT(func(v *Task) pulumi.IntOutput { return v.TimeOut }).(pulumi.IntOutput)
}

// trigger type.
func (o TaskOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.TriggerType }).(pulumi.StringOutput)
}

type TaskArrayOutput struct{ *pulumi.OutputState }

func (TaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (o TaskArrayOutput) ToTaskArrayOutput() TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) Index(i pulumi.IntInput) TaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Task {
		return vs[0].([]*Task)[vs[1].(int)]
	}).(TaskOutput)
}

type TaskMapOutput struct{ *pulumi.OutputState }

func (TaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (o TaskMapOutput) ToTaskMapOutput() TaskMapOutput {
	return o
}

func (o TaskMapOutput) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return o
}

func (o TaskMapOutput) MapIndex(k pulumi.StringInput) TaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Task {
		return vs[0].(map[string]*Task)[vs[1].(string)]
	}).(TaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskInput)(nil)).Elem(), &Task{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskArrayInput)(nil)).Elem(), TaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskMapInput)(nil)).Elem(), TaskMap{})
	pulumi.RegisterOutputType(TaskOutput{})
	pulumi.RegisterOutputType(TaskArrayOutput{})
	pulumi.RegisterOutputType(TaskMapOutput{})
}
