// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cat

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cat taskSet
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cat"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cat"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"ipType":            0,
// 			"grabBag":           0,
// 			"filterIp":          0,
// 			"netIcmpOn":         1,
// 			"netIcmpActivex":    0,
// 			"netIcmpTimeout":    20,
// 			"netIcmpInterval":   0.5,
// 			"netIcmpNum":        20,
// 			"netIcmpSize":       32,
// 			"netIcmpDataCut":    1,
// 			"netDnsOn":          1,
// 			"netDnsTimeout":     5,
// 			"netDnsQuerymethod": 1,
// 			"netDnsNs":          "",
// 			"netDigOn":          1,
// 			"netDnsServer":      2,
// 			"netTracertOn":      1,
// 			"netTracertTimeout": 60,
// 			"netTracertNum":     30,
// 			"whiteList":         "",
// 			"blackList":         "",
// 			"netIcmpActivexStr": "",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := Cat.NewTaskSet(ctx, "taskSet", &Cat.TaskSetArgs{
// 			BatchTasks: &cat.TaskSetBatchTasksArgs{
// 				Name:          pulumi.String("demo"),
// 				TargetAddress: pulumi.String("http://www.baidu.com"),
// 			},
// 			TaskType: pulumi.Int(5),
// 			Nodes: pulumi.StringArray{
// 				pulumi.String("12136"),
// 				pulumi.String("12137"),
// 				pulumi.String("12138"),
// 				pulumi.String("12141"),
// 				pulumi.String("12144"),
// 			},
// 			Interval:     pulumi.Int(6),
// 			Parameters:   pulumi.String(json0),
// 			TaskCategory: pulumi.Int(1),
// 			Cron:         pulumi.String("* 0-6 * * *"),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
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
// cat task_set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cat/taskSet:TaskSet task_set taskSet_id
// ```
type TaskSet struct {
	pulumi.CustomResourceState

	// Batch task name address.
	BatchTasks TaskSetBatchTasksOutput `pulumi:"batchTasks"`
	// Timer task cron expression.
	Cron pulumi.StringPtrOutput `pulumi:"cron"`
	// Task interval minutes in (1,5,10,15,30,60,120,240).
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Task Nodes.
	Nodes pulumi.StringArrayOutput `pulumi:"nodes"`
	// tasks parameters.
	Parameters pulumi.StringOutput `pulumi:"parameters"`
	// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
	Status pulumi.IntOutput `pulumi:"status"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Task category,1:PC,2:Mobile.
	TaskCategory pulumi.IntOutput `pulumi:"taskCategory"`
	// Task Id.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
	TaskType pulumi.IntOutput `pulumi:"taskType"`
}

// NewTaskSet registers a new resource with the given unique name, arguments, and options.
func NewTaskSet(ctx *pulumi.Context,
	name string, args *TaskSetArgs, opts ...pulumi.ResourceOption) (*TaskSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BatchTasks == nil {
		return nil, errors.New("invalid value for required argument 'BatchTasks'")
	}
	if args.Interval == nil {
		return nil, errors.New("invalid value for required argument 'Interval'")
	}
	if args.Nodes == nil {
		return nil, errors.New("invalid value for required argument 'Nodes'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.TaskCategory == nil {
		return nil, errors.New("invalid value for required argument 'TaskCategory'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TaskSet
	err := ctx.RegisterResource("tencentcloud:Cat/taskSet:TaskSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskSet gets an existing TaskSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskSetState, opts ...pulumi.ResourceOption) (*TaskSet, error) {
	var resource TaskSet
	err := ctx.ReadResource("tencentcloud:Cat/taskSet:TaskSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskSet resources.
type taskSetState struct {
	// Batch task name address.
	BatchTasks *TaskSetBatchTasks `pulumi:"batchTasks"`
	// Timer task cron expression.
	Cron *string `pulumi:"cron"`
	// Task interval minutes in (1,5,10,15,30,60,120,240).
	Interval *int `pulumi:"interval"`
	// Task Nodes.
	Nodes []string `pulumi:"nodes"`
	// tasks parameters.
	Parameters *string `pulumi:"parameters"`
	// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
	Status *int `pulumi:"status"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Task category,1:PC,2:Mobile.
	TaskCategory *int `pulumi:"taskCategory"`
	// Task Id.
	TaskId *string `pulumi:"taskId"`
	// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
	TaskType *int `pulumi:"taskType"`
}

type TaskSetState struct {
	// Batch task name address.
	BatchTasks TaskSetBatchTasksPtrInput
	// Timer task cron expression.
	Cron pulumi.StringPtrInput
	// Task interval minutes in (1,5,10,15,30,60,120,240).
	Interval pulumi.IntPtrInput
	// Task Nodes.
	Nodes pulumi.StringArrayInput
	// tasks parameters.
	Parameters pulumi.StringPtrInput
	// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
	Status pulumi.IntPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// Task category,1:PC,2:Mobile.
	TaskCategory pulumi.IntPtrInput
	// Task Id.
	TaskId pulumi.StringPtrInput
	// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
	TaskType pulumi.IntPtrInput
}

func (TaskSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskSetState)(nil)).Elem()
}

type taskSetArgs struct {
	// Batch task name address.
	BatchTasks TaskSetBatchTasks `pulumi:"batchTasks"`
	// Timer task cron expression.
	Cron *string `pulumi:"cron"`
	// Task interval minutes in (1,5,10,15,30,60,120,240).
	Interval int `pulumi:"interval"`
	// Task Nodes.
	Nodes []string `pulumi:"nodes"`
	// tasks parameters.
	Parameters string `pulumi:"parameters"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Task category,1:PC,2:Mobile.
	TaskCategory int `pulumi:"taskCategory"`
	// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
	TaskType int `pulumi:"taskType"`
}

// The set of arguments for constructing a TaskSet resource.
type TaskSetArgs struct {
	// Batch task name address.
	BatchTasks TaskSetBatchTasksInput
	// Timer task cron expression.
	Cron pulumi.StringPtrInput
	// Task interval minutes in (1,5,10,15,30,60,120,240).
	Interval pulumi.IntInput
	// Task Nodes.
	Nodes pulumi.StringArrayInput
	// tasks parameters.
	Parameters pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
	// Task category,1:PC,2:Mobile.
	TaskCategory pulumi.IntInput
	// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
	TaskType pulumi.IntInput
}

func (TaskSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskSetArgs)(nil)).Elem()
}

type TaskSetInput interface {
	pulumi.Input

	ToTaskSetOutput() TaskSetOutput
	ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput
}

func (*TaskSet) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSet)(nil)).Elem()
}

func (i *TaskSet) ToTaskSetOutput() TaskSetOutput {
	return i.ToTaskSetOutputWithContext(context.Background())
}

func (i *TaskSet) ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetOutput)
}

// TaskSetArrayInput is an input type that accepts TaskSetArray and TaskSetArrayOutput values.
// You can construct a concrete instance of `TaskSetArrayInput` via:
//
//          TaskSetArray{ TaskSetArgs{...} }
type TaskSetArrayInput interface {
	pulumi.Input

	ToTaskSetArrayOutput() TaskSetArrayOutput
	ToTaskSetArrayOutputWithContext(context.Context) TaskSetArrayOutput
}

type TaskSetArray []TaskSetInput

func (TaskSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskSet)(nil)).Elem()
}

func (i TaskSetArray) ToTaskSetArrayOutput() TaskSetArrayOutput {
	return i.ToTaskSetArrayOutputWithContext(context.Background())
}

func (i TaskSetArray) ToTaskSetArrayOutputWithContext(ctx context.Context) TaskSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetArrayOutput)
}

// TaskSetMapInput is an input type that accepts TaskSetMap and TaskSetMapOutput values.
// You can construct a concrete instance of `TaskSetMapInput` via:
//
//          TaskSetMap{ "key": TaskSetArgs{...} }
type TaskSetMapInput interface {
	pulumi.Input

	ToTaskSetMapOutput() TaskSetMapOutput
	ToTaskSetMapOutputWithContext(context.Context) TaskSetMapOutput
}

type TaskSetMap map[string]TaskSetInput

func (TaskSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskSet)(nil)).Elem()
}

func (i TaskSetMap) ToTaskSetMapOutput() TaskSetMapOutput {
	return i.ToTaskSetMapOutputWithContext(context.Background())
}

func (i TaskSetMap) ToTaskSetMapOutputWithContext(ctx context.Context) TaskSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetMapOutput)
}

type TaskSetOutput struct{ *pulumi.OutputState }

func (TaskSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSet)(nil)).Elem()
}

func (o TaskSetOutput) ToTaskSetOutput() TaskSetOutput {
	return o
}

func (o TaskSetOutput) ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput {
	return o
}

// Batch task name address.
func (o TaskSetOutput) BatchTasks() TaskSetBatchTasksOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetBatchTasksOutput { return v.BatchTasks }).(TaskSetBatchTasksOutput)
}

// Timer task cron expression.
func (o TaskSetOutput) Cron() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringPtrOutput { return v.Cron }).(pulumi.StringPtrOutput)
}

// Task interval minutes in (1,5,10,15,30,60,120,240).
func (o TaskSetOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

// Task Nodes.
func (o TaskSetOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringArrayOutput { return v.Nodes }).(pulumi.StringArrayOutput)
}

// tasks parameters.
func (o TaskSetOutput) Parameters() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.Parameters }).(pulumi.StringOutput)
}

// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
func (o TaskSetOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Tag description list.
func (o TaskSetOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Task category,1:PC,2:Mobile.
func (o TaskSetOutput) TaskCategory() pulumi.IntOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.IntOutput { return v.TaskCategory }).(pulumi.IntOutput)
}

// Task Id.
func (o TaskSetOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
func (o TaskSetOutput) TaskType() pulumi.IntOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.IntOutput { return v.TaskType }).(pulumi.IntOutput)
}

type TaskSetArrayOutput struct{ *pulumi.OutputState }

func (TaskSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskSet)(nil)).Elem()
}

func (o TaskSetArrayOutput) ToTaskSetArrayOutput() TaskSetArrayOutput {
	return o
}

func (o TaskSetArrayOutput) ToTaskSetArrayOutputWithContext(ctx context.Context) TaskSetArrayOutput {
	return o
}

func (o TaskSetArrayOutput) Index(i pulumi.IntInput) TaskSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaskSet {
		return vs[0].([]*TaskSet)[vs[1].(int)]
	}).(TaskSetOutput)
}

type TaskSetMapOutput struct{ *pulumi.OutputState }

func (TaskSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskSet)(nil)).Elem()
}

func (o TaskSetMapOutput) ToTaskSetMapOutput() TaskSetMapOutput {
	return o
}

func (o TaskSetMapOutput) ToTaskSetMapOutputWithContext(ctx context.Context) TaskSetMapOutput {
	return o
}

func (o TaskSetMapOutput) MapIndex(k pulumi.StringInput) TaskSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaskSet {
		return vs[0].(map[string]*TaskSet)[vs[1].(string)]
	}).(TaskSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetInput)(nil)).Elem(), &TaskSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetArrayInput)(nil)).Elem(), TaskSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetMapInput)(nil)).Elem(), TaskSetMap{})
	pulumi.RegisterOutputType(TaskSetOutput{})
	pulumi.RegisterOutputType(TaskSetArrayOutput{})
	pulumi.RegisterOutputType(TaskSetMapOutput{})
}