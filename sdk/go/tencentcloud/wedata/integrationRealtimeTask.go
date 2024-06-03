// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wedata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a wedata integrationRealtimeTask
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Wedata"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Wedata.NewIntegrationRealtimeTask(ctx, "example", &Wedata.IntegrationRealtimeTaskArgs{
//				Description: pulumi.String("description."),
//				ProjectId:   pulumi.String("1612982498218618880"),
//				SyncType:    pulumi.Int(1),
//				TaskInfo: &wedata.IntegrationRealtimeTaskTaskInfoArgs{
//					Configs: wedata.IntegrationRealtimeTaskTaskInfoConfigArray{
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("concurrency"),
//							Value: pulumi.String("1"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("TaskManager"),
//							Value: pulumi.String("1"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("JobManager"),
//							Value: pulumi.String("1"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("TolerateDirtyData"),
//							Value: pulumi.String("0"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("CheckpointingInterval"),
//							Value: pulumi.String("1"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("CheckpointingIntervalUnit"),
//							Value: pulumi.String("min"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("RestartStrategyFixedDelayAttempts"),
//							Value: pulumi.String("-1"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("ResourceAllocationType"),
//							Value: pulumi.String("0"),
//						},
//						&wedata.IntegrationRealtimeTaskTaskInfoConfigArgs{
//							Name:  pulumi.String("TaskAlarmRegularList"),
//							Value: pulumi.String("35"),
//						},
//					},
//					ExecutorId: pulumi.String("20230313175748567418"),
//					Incharge:   pulumi.String("100028439226"),
//				},
//				TaskMode: pulumi.String("1"),
//				TaskName: pulumi.String("tf_example"),
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
// wedata integration_realtime_task can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Wedata/integrationRealtimeTask:IntegrationRealtimeTask example 1776563389209296896#h9d39630a-ae45-4460-90b2-0b093cbfef5d
// ```
type IntegrationRealtimeTask struct {
	pulumi.CustomResourceState

	// Description information.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
	SyncType pulumi.IntOutput `pulumi:"syncType"`
	// The task id to which the node belongs.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// Task Information.
	TaskInfo IntegrationRealtimeTaskTaskInfoOutput `pulumi:"taskInfo"`
	// Task display mode, 0: canvas mode, 1: form mode.
	TaskMode pulumi.StringOutput `pulumi:"taskMode"`
	// Task name.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
}

// NewIntegrationRealtimeTask registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRealtimeTask(ctx *pulumi.Context,
	name string, args *IntegrationRealtimeTaskArgs, opts ...pulumi.ResourceOption) (*IntegrationRealtimeTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SyncType == nil {
		return nil, errors.New("invalid value for required argument 'SyncType'")
	}
	if args.TaskInfo == nil {
		return nil, errors.New("invalid value for required argument 'TaskInfo'")
	}
	if args.TaskMode == nil {
		return nil, errors.New("invalid value for required argument 'TaskMode'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationRealtimeTask
	err := ctx.RegisterResource("tencentcloud:Wedata/integrationRealtimeTask:IntegrationRealtimeTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRealtimeTask gets an existing IntegrationRealtimeTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRealtimeTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRealtimeTaskState, opts ...pulumi.ResourceOption) (*IntegrationRealtimeTask, error) {
	var resource IntegrationRealtimeTask
	err := ctx.ReadResource("tencentcloud:Wedata/integrationRealtimeTask:IntegrationRealtimeTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRealtimeTask resources.
type integrationRealtimeTaskState struct {
	// Description information.
	Description *string `pulumi:"description"`
	// Project ID.
	ProjectId *string `pulumi:"projectId"`
	// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
	SyncType *int `pulumi:"syncType"`
	// The task id to which the node belongs.
	TaskId *string `pulumi:"taskId"`
	// Task Information.
	TaskInfo *IntegrationRealtimeTaskTaskInfo `pulumi:"taskInfo"`
	// Task display mode, 0: canvas mode, 1: form mode.
	TaskMode *string `pulumi:"taskMode"`
	// Task name.
	TaskName *string `pulumi:"taskName"`
}

type IntegrationRealtimeTaskState struct {
	// Description information.
	Description pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.StringPtrInput
	// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
	SyncType pulumi.IntPtrInput
	// The task id to which the node belongs.
	TaskId pulumi.StringPtrInput
	// Task Information.
	TaskInfo IntegrationRealtimeTaskTaskInfoPtrInput
	// Task display mode, 0: canvas mode, 1: form mode.
	TaskMode pulumi.StringPtrInput
	// Task name.
	TaskName pulumi.StringPtrInput
}

func (IntegrationRealtimeTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRealtimeTaskState)(nil)).Elem()
}

type integrationRealtimeTaskArgs struct {
	// Description information.
	Description *string `pulumi:"description"`
	// Project ID.
	ProjectId string `pulumi:"projectId"`
	// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
	SyncType int `pulumi:"syncType"`
	// Task Information.
	TaskInfo IntegrationRealtimeTaskTaskInfo `pulumi:"taskInfo"`
	// Task display mode, 0: canvas mode, 1: form mode.
	TaskMode string `pulumi:"taskMode"`
	// Task name.
	TaskName string `pulumi:"taskName"`
}

// The set of arguments for constructing a IntegrationRealtimeTask resource.
type IntegrationRealtimeTaskArgs struct {
	// Description information.
	Description pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.StringInput
	// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
	SyncType pulumi.IntInput
	// Task Information.
	TaskInfo IntegrationRealtimeTaskTaskInfoInput
	// Task display mode, 0: canvas mode, 1: form mode.
	TaskMode pulumi.StringInput
	// Task name.
	TaskName pulumi.StringInput
}

func (IntegrationRealtimeTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRealtimeTaskArgs)(nil)).Elem()
}

type IntegrationRealtimeTaskInput interface {
	pulumi.Input

	ToIntegrationRealtimeTaskOutput() IntegrationRealtimeTaskOutput
	ToIntegrationRealtimeTaskOutputWithContext(ctx context.Context) IntegrationRealtimeTaskOutput
}

func (*IntegrationRealtimeTask) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRealtimeTask)(nil)).Elem()
}

func (i *IntegrationRealtimeTask) ToIntegrationRealtimeTaskOutput() IntegrationRealtimeTaskOutput {
	return i.ToIntegrationRealtimeTaskOutputWithContext(context.Background())
}

func (i *IntegrationRealtimeTask) ToIntegrationRealtimeTaskOutputWithContext(ctx context.Context) IntegrationRealtimeTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRealtimeTaskOutput)
}

// IntegrationRealtimeTaskArrayInput is an input type that accepts IntegrationRealtimeTaskArray and IntegrationRealtimeTaskArrayOutput values.
// You can construct a concrete instance of `IntegrationRealtimeTaskArrayInput` via:
//
//	IntegrationRealtimeTaskArray{ IntegrationRealtimeTaskArgs{...} }
type IntegrationRealtimeTaskArrayInput interface {
	pulumi.Input

	ToIntegrationRealtimeTaskArrayOutput() IntegrationRealtimeTaskArrayOutput
	ToIntegrationRealtimeTaskArrayOutputWithContext(context.Context) IntegrationRealtimeTaskArrayOutput
}

type IntegrationRealtimeTaskArray []IntegrationRealtimeTaskInput

func (IntegrationRealtimeTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRealtimeTask)(nil)).Elem()
}

func (i IntegrationRealtimeTaskArray) ToIntegrationRealtimeTaskArrayOutput() IntegrationRealtimeTaskArrayOutput {
	return i.ToIntegrationRealtimeTaskArrayOutputWithContext(context.Background())
}

func (i IntegrationRealtimeTaskArray) ToIntegrationRealtimeTaskArrayOutputWithContext(ctx context.Context) IntegrationRealtimeTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRealtimeTaskArrayOutput)
}

// IntegrationRealtimeTaskMapInput is an input type that accepts IntegrationRealtimeTaskMap and IntegrationRealtimeTaskMapOutput values.
// You can construct a concrete instance of `IntegrationRealtimeTaskMapInput` via:
//
//	IntegrationRealtimeTaskMap{ "key": IntegrationRealtimeTaskArgs{...} }
type IntegrationRealtimeTaskMapInput interface {
	pulumi.Input

	ToIntegrationRealtimeTaskMapOutput() IntegrationRealtimeTaskMapOutput
	ToIntegrationRealtimeTaskMapOutputWithContext(context.Context) IntegrationRealtimeTaskMapOutput
}

type IntegrationRealtimeTaskMap map[string]IntegrationRealtimeTaskInput

func (IntegrationRealtimeTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRealtimeTask)(nil)).Elem()
}

func (i IntegrationRealtimeTaskMap) ToIntegrationRealtimeTaskMapOutput() IntegrationRealtimeTaskMapOutput {
	return i.ToIntegrationRealtimeTaskMapOutputWithContext(context.Background())
}

func (i IntegrationRealtimeTaskMap) ToIntegrationRealtimeTaskMapOutputWithContext(ctx context.Context) IntegrationRealtimeTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRealtimeTaskMapOutput)
}

type IntegrationRealtimeTaskOutput struct{ *pulumi.OutputState }

func (IntegrationRealtimeTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRealtimeTask)(nil)).Elem()
}

func (o IntegrationRealtimeTaskOutput) ToIntegrationRealtimeTaskOutput() IntegrationRealtimeTaskOutput {
	return o
}

func (o IntegrationRealtimeTaskOutput) ToIntegrationRealtimeTaskOutputWithContext(ctx context.Context) IntegrationRealtimeTaskOutput {
	return o
}

// Description information.
func (o IntegrationRealtimeTaskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Project ID.
func (o IntegrationRealtimeTaskOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
func (o IntegrationRealtimeTaskOutput) SyncType() pulumi.IntOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.IntOutput { return v.SyncType }).(pulumi.IntOutput)
}

// The task id to which the node belongs.
func (o IntegrationRealtimeTaskOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// Task Information.
func (o IntegrationRealtimeTaskOutput) TaskInfo() IntegrationRealtimeTaskTaskInfoOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) IntegrationRealtimeTaskTaskInfoOutput { return v.TaskInfo }).(IntegrationRealtimeTaskTaskInfoOutput)
}

// Task display mode, 0: canvas mode, 1: form mode.
func (o IntegrationRealtimeTaskOutput) TaskMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.StringOutput { return v.TaskMode }).(pulumi.StringOutput)
}

// Task name.
func (o IntegrationRealtimeTaskOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationRealtimeTask) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

type IntegrationRealtimeTaskArrayOutput struct{ *pulumi.OutputState }

func (IntegrationRealtimeTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRealtimeTask)(nil)).Elem()
}

func (o IntegrationRealtimeTaskArrayOutput) ToIntegrationRealtimeTaskArrayOutput() IntegrationRealtimeTaskArrayOutput {
	return o
}

func (o IntegrationRealtimeTaskArrayOutput) ToIntegrationRealtimeTaskArrayOutputWithContext(ctx context.Context) IntegrationRealtimeTaskArrayOutput {
	return o
}

func (o IntegrationRealtimeTaskArrayOutput) Index(i pulumi.IntInput) IntegrationRealtimeTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationRealtimeTask {
		return vs[0].([]*IntegrationRealtimeTask)[vs[1].(int)]
	}).(IntegrationRealtimeTaskOutput)
}

type IntegrationRealtimeTaskMapOutput struct{ *pulumi.OutputState }

func (IntegrationRealtimeTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRealtimeTask)(nil)).Elem()
}

func (o IntegrationRealtimeTaskMapOutput) ToIntegrationRealtimeTaskMapOutput() IntegrationRealtimeTaskMapOutput {
	return o
}

func (o IntegrationRealtimeTaskMapOutput) ToIntegrationRealtimeTaskMapOutputWithContext(ctx context.Context) IntegrationRealtimeTaskMapOutput {
	return o
}

func (o IntegrationRealtimeTaskMapOutput) MapIndex(k pulumi.StringInput) IntegrationRealtimeTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationRealtimeTask {
		return vs[0].(map[string]*IntegrationRealtimeTask)[vs[1].(string)]
	}).(IntegrationRealtimeTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRealtimeTaskInput)(nil)).Elem(), &IntegrationRealtimeTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRealtimeTaskArrayInput)(nil)).Elem(), IntegrationRealtimeTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRealtimeTaskMapInput)(nil)).Elem(), IntegrationRealtimeTaskMap{})
	pulumi.RegisterOutputType(IntegrationRealtimeTaskOutput{})
	pulumi.RegisterOutputType(IntegrationRealtimeTaskArrayOutput{})
	pulumi.RegisterOutputType(IntegrationRealtimeTaskMapOutput{})
}
