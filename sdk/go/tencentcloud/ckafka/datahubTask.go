// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a ckafka datahubTask
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ckafka.NewDatahubTask(ctx, "datahubTask", &Ckafka.DatahubTaskArgs{
//				SourceResource: &ckafka.DatahubTaskSourceResourceArgs{
//					PostgreSqlParam: &ckafka.DatahubTaskSourceResourcePostgreSqlParamArgs{
//						Database:         pulumi.String("postgres"),
//						IsTableRegular:   pulumi.Bool(false),
//						KeyColumns:       pulumi.String(""),
//						PluginName:       pulumi.String("decoderbufs"),
//						RecordWithSchema: pulumi.Bool(false),
//						Resource:         pulumi.String("resource-y9nxnw46"),
//						SnapshotMode:     pulumi.String("never"),
//						Table:            pulumi.String("*"),
//					},
//					Type: pulumi.String("POSTGRESQL"),
//				},
//				TargetResource: &ckafka.DatahubTaskTargetResourceArgs{
//					TopicParam: &ckafka.DatahubTaskTargetResourceTopicParamArgs{
//						CompressionType:    pulumi.String("none"),
//						Resource:           pulumi.String("1308726196-keep-topic"),
//						UseAutoCreateTopic: pulumi.Bool(false),
//					},
//					Type: pulumi.String("TOPIC"),
//				},
//				TaskName: pulumi.String("test-task123321"),
//				TaskType: pulumi.String("SOURCE"),
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
// ckafka datahub_task can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Ckafka/datahubTask:DatahubTask datahub_task datahub_task_id
// ```
type DatahubTask struct {
	pulumi.CustomResourceState

	// SchemaId.
	SchemaId pulumi.StringPtrOutput `pulumi:"schemaId"`
	// data resource.
	SourceResource DatahubTaskSourceResourcePtrOutput `pulumi:"sourceResource"`
	// Target Resource.
	TargetResource DatahubTaskTargetResourcePtrOutput `pulumi:"targetResource"`
	// name of the task.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
	// type of the task, SOURCE(data input), SINK(data output).
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// Data Processing Rules.
	TransformParam DatahubTaskTransformParamPtrOutput `pulumi:"transformParam"`
	// Data processing rules.
	TransformsParam DatahubTaskTransformsParamPtrOutput `pulumi:"transformsParam"`
}

// NewDatahubTask registers a new resource with the given unique name, arguments, and options.
func NewDatahubTask(ctx *pulumi.Context,
	name string, args *DatahubTaskArgs, opts ...pulumi.ResourceOption) (*DatahubTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatahubTask
	err := ctx.RegisterResource("tencentcloud:Ckafka/datahubTask:DatahubTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatahubTask gets an existing DatahubTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatahubTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatahubTaskState, opts ...pulumi.ResourceOption) (*DatahubTask, error) {
	var resource DatahubTask
	err := ctx.ReadResource("tencentcloud:Ckafka/datahubTask:DatahubTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatahubTask resources.
type datahubTaskState struct {
	// SchemaId.
	SchemaId *string `pulumi:"schemaId"`
	// data resource.
	SourceResource *DatahubTaskSourceResource `pulumi:"sourceResource"`
	// Target Resource.
	TargetResource *DatahubTaskTargetResource `pulumi:"targetResource"`
	// name of the task.
	TaskName *string `pulumi:"taskName"`
	// type of the task, SOURCE(data input), SINK(data output).
	TaskType *string `pulumi:"taskType"`
	// Data Processing Rules.
	TransformParam *DatahubTaskTransformParam `pulumi:"transformParam"`
	// Data processing rules.
	TransformsParam *DatahubTaskTransformsParam `pulumi:"transformsParam"`
}

type DatahubTaskState struct {
	// SchemaId.
	SchemaId pulumi.StringPtrInput
	// data resource.
	SourceResource DatahubTaskSourceResourcePtrInput
	// Target Resource.
	TargetResource DatahubTaskTargetResourcePtrInput
	// name of the task.
	TaskName pulumi.StringPtrInput
	// type of the task, SOURCE(data input), SINK(data output).
	TaskType pulumi.StringPtrInput
	// Data Processing Rules.
	TransformParam DatahubTaskTransformParamPtrInput
	// Data processing rules.
	TransformsParam DatahubTaskTransformsParamPtrInput
}

func (DatahubTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*datahubTaskState)(nil)).Elem()
}

type datahubTaskArgs struct {
	// SchemaId.
	SchemaId *string `pulumi:"schemaId"`
	// data resource.
	SourceResource *DatahubTaskSourceResource `pulumi:"sourceResource"`
	// Target Resource.
	TargetResource *DatahubTaskTargetResource `pulumi:"targetResource"`
	// name of the task.
	TaskName string `pulumi:"taskName"`
	// type of the task, SOURCE(data input), SINK(data output).
	TaskType string `pulumi:"taskType"`
	// Data Processing Rules.
	TransformParam *DatahubTaskTransformParam `pulumi:"transformParam"`
	// Data processing rules.
	TransformsParam *DatahubTaskTransformsParam `pulumi:"transformsParam"`
}

// The set of arguments for constructing a DatahubTask resource.
type DatahubTaskArgs struct {
	// SchemaId.
	SchemaId pulumi.StringPtrInput
	// data resource.
	SourceResource DatahubTaskSourceResourcePtrInput
	// Target Resource.
	TargetResource DatahubTaskTargetResourcePtrInput
	// name of the task.
	TaskName pulumi.StringInput
	// type of the task, SOURCE(data input), SINK(data output).
	TaskType pulumi.StringInput
	// Data Processing Rules.
	TransformParam DatahubTaskTransformParamPtrInput
	// Data processing rules.
	TransformsParam DatahubTaskTransformsParamPtrInput
}

func (DatahubTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datahubTaskArgs)(nil)).Elem()
}

type DatahubTaskInput interface {
	pulumi.Input

	ToDatahubTaskOutput() DatahubTaskOutput
	ToDatahubTaskOutputWithContext(ctx context.Context) DatahubTaskOutput
}

func (*DatahubTask) ElementType() reflect.Type {
	return reflect.TypeOf((**DatahubTask)(nil)).Elem()
}

func (i *DatahubTask) ToDatahubTaskOutput() DatahubTaskOutput {
	return i.ToDatahubTaskOutputWithContext(context.Background())
}

func (i *DatahubTask) ToDatahubTaskOutputWithContext(ctx context.Context) DatahubTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatahubTaskOutput)
}

// DatahubTaskArrayInput is an input type that accepts DatahubTaskArray and DatahubTaskArrayOutput values.
// You can construct a concrete instance of `DatahubTaskArrayInput` via:
//
//	DatahubTaskArray{ DatahubTaskArgs{...} }
type DatahubTaskArrayInput interface {
	pulumi.Input

	ToDatahubTaskArrayOutput() DatahubTaskArrayOutput
	ToDatahubTaskArrayOutputWithContext(context.Context) DatahubTaskArrayOutput
}

type DatahubTaskArray []DatahubTaskInput

func (DatahubTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatahubTask)(nil)).Elem()
}

func (i DatahubTaskArray) ToDatahubTaskArrayOutput() DatahubTaskArrayOutput {
	return i.ToDatahubTaskArrayOutputWithContext(context.Background())
}

func (i DatahubTaskArray) ToDatahubTaskArrayOutputWithContext(ctx context.Context) DatahubTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatahubTaskArrayOutput)
}

// DatahubTaskMapInput is an input type that accepts DatahubTaskMap and DatahubTaskMapOutput values.
// You can construct a concrete instance of `DatahubTaskMapInput` via:
//
//	DatahubTaskMap{ "key": DatahubTaskArgs{...} }
type DatahubTaskMapInput interface {
	pulumi.Input

	ToDatahubTaskMapOutput() DatahubTaskMapOutput
	ToDatahubTaskMapOutputWithContext(context.Context) DatahubTaskMapOutput
}

type DatahubTaskMap map[string]DatahubTaskInput

func (DatahubTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatahubTask)(nil)).Elem()
}

func (i DatahubTaskMap) ToDatahubTaskMapOutput() DatahubTaskMapOutput {
	return i.ToDatahubTaskMapOutputWithContext(context.Background())
}

func (i DatahubTaskMap) ToDatahubTaskMapOutputWithContext(ctx context.Context) DatahubTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatahubTaskMapOutput)
}

type DatahubTaskOutput struct{ *pulumi.OutputState }

func (DatahubTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatahubTask)(nil)).Elem()
}

func (o DatahubTaskOutput) ToDatahubTaskOutput() DatahubTaskOutput {
	return o
}

func (o DatahubTaskOutput) ToDatahubTaskOutputWithContext(ctx context.Context) DatahubTaskOutput {
	return o
}

// SchemaId.
func (o DatahubTaskOutput) SchemaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatahubTask) pulumi.StringPtrOutput { return v.SchemaId }).(pulumi.StringPtrOutput)
}

// data resource.
func (o DatahubTaskOutput) SourceResource() DatahubTaskSourceResourcePtrOutput {
	return o.ApplyT(func(v *DatahubTask) DatahubTaskSourceResourcePtrOutput { return v.SourceResource }).(DatahubTaskSourceResourcePtrOutput)
}

// Target Resource.
func (o DatahubTaskOutput) TargetResource() DatahubTaskTargetResourcePtrOutput {
	return o.ApplyT(func(v *DatahubTask) DatahubTaskTargetResourcePtrOutput { return v.TargetResource }).(DatahubTaskTargetResourcePtrOutput)
}

// name of the task.
func (o DatahubTaskOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatahubTask) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

// type of the task, SOURCE(data input), SINK(data output).
func (o DatahubTaskOutput) TaskType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatahubTask) pulumi.StringOutput { return v.TaskType }).(pulumi.StringOutput)
}

// Data Processing Rules.
func (o DatahubTaskOutput) TransformParam() DatahubTaskTransformParamPtrOutput {
	return o.ApplyT(func(v *DatahubTask) DatahubTaskTransformParamPtrOutput { return v.TransformParam }).(DatahubTaskTransformParamPtrOutput)
}

// Data processing rules.
func (o DatahubTaskOutput) TransformsParam() DatahubTaskTransformsParamPtrOutput {
	return o.ApplyT(func(v *DatahubTask) DatahubTaskTransformsParamPtrOutput { return v.TransformsParam }).(DatahubTaskTransformsParamPtrOutput)
}

type DatahubTaskArrayOutput struct{ *pulumi.OutputState }

func (DatahubTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatahubTask)(nil)).Elem()
}

func (o DatahubTaskArrayOutput) ToDatahubTaskArrayOutput() DatahubTaskArrayOutput {
	return o
}

func (o DatahubTaskArrayOutput) ToDatahubTaskArrayOutputWithContext(ctx context.Context) DatahubTaskArrayOutput {
	return o
}

func (o DatahubTaskArrayOutput) Index(i pulumi.IntInput) DatahubTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatahubTask {
		return vs[0].([]*DatahubTask)[vs[1].(int)]
	}).(DatahubTaskOutput)
}

type DatahubTaskMapOutput struct{ *pulumi.OutputState }

func (DatahubTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatahubTask)(nil)).Elem()
}

func (o DatahubTaskMapOutput) ToDatahubTaskMapOutput() DatahubTaskMapOutput {
	return o
}

func (o DatahubTaskMapOutput) ToDatahubTaskMapOutputWithContext(ctx context.Context) DatahubTaskMapOutput {
	return o
}

func (o DatahubTaskMapOutput) MapIndex(k pulumi.StringInput) DatahubTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatahubTask {
		return vs[0].(map[string]*DatahubTask)[vs[1].(string)]
	}).(DatahubTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatahubTaskInput)(nil)).Elem(), &DatahubTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatahubTaskArrayInput)(nil)).Elem(), DatahubTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatahubTaskMapInput)(nil)).Elem(), DatahubTaskMap{})
	pulumi.RegisterOutputType(DatahubTaskOutput{})
	pulumi.RegisterOutputType(DatahubTaskArrayOutput{})
	pulumi.RegisterOutputType(DatahubTaskMapOutput{})
}
