// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cls scheduledSql
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		logset, err := Cls.NewLogset(ctx, "logset", &Cls.LogsetArgs{
// 			LogsetName: pulumi.String("tf-example-logset"),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		topic, err := Cls.NewTopic(ctx, "topic", &Cls.TopicArgs{
// 			TopicName:          pulumi.String("tf-example-topic"),
// 			LogsetId:           logset.ID(),
// 			AutoSplit:          pulumi.Bool(false),
// 			MaxSplitPartitions: pulumi.Int(20),
// 			PartitionCount:     pulumi.Int(1),
// 			Period:             pulumi.Int(10),
// 			StorageType:        pulumi.String("hot"),
// 			Tags: pulumi.AnyMap{
// 				"test": pulumi.Any("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Cls.NewScheduledSql(ctx, "scheduledSql", &Cls.ScheduledSqlArgs{
// 			SrcTopicId: topic.ID(),
// 			EnableFlag: pulumi.Int(1),
// 			DstResource: &cls.ScheduledSqlDstResourceArgs{
// 				TopicId:    topic.ID(),
// 				Region:     pulumi.String("ap-guangzhou"),
// 				BizType:    pulumi.Int(0),
// 				MetricName: pulumi.String("test"),
// 			},
// 			ScheduledSqlContent: pulumi.String("xxx"),
// 			ProcessStartTime:    pulumi.Int(1690515360000),
// 			ProcessType:         pulumi.Int(1),
// 			ProcessPeriod:       pulumi.Int(10),
// 			ProcessTimeWindow:   pulumi.String("@m-15m,@m"),
// 			ProcessDelay:        pulumi.Int(5),
// 			SrcTopicRegion:      pulumi.String("ap-guangzhou"),
// 			ProcessEndTime:      pulumi.Int(1690515360000),
// 			SyntaxRule:          pulumi.Int(0),
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
// cls scheduled_sql can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cls/scheduledSql:ScheduledSql scheduled_sql scheduled_sql_id
// ```
type ScheduledSql struct {
	pulumi.CustomResourceState

	// scheduled slq dst resource.
	DstResource ScheduledSqlDstResourceOutput `pulumi:"dstResource"`
	// task enable flag.
	EnableFlag pulumi.IntOutput `pulumi:"enableFlag"`
	// task name.
	Name pulumi.StringOutput `pulumi:"name"`
	// process delay.
	ProcessDelay pulumi.IntOutput `pulumi:"processDelay"`
	// process end timestamp.
	ProcessEndTime pulumi.IntOutput `pulumi:"processEndTime"`
	// process period.
	ProcessPeriod pulumi.IntOutput `pulumi:"processPeriod"`
	// process start timestamp.
	ProcessStartTime pulumi.IntOutput `pulumi:"processStartTime"`
	// process time window.
	ProcessTimeWindow pulumi.StringOutput `pulumi:"processTimeWindow"`
	// process type.
	ProcessType pulumi.IntOutput `pulumi:"processType"`
	// scheduled sql content.
	ScheduledSqlContent pulumi.StringOutput `pulumi:"scheduledSqlContent"`
	// src topic id.
	SrcTopicId pulumi.StringOutput `pulumi:"srcTopicId"`
	// src topic region.
	SrcTopicRegion pulumi.StringOutput `pulumi:"srcTopicRegion"`
	// syntax rule.
	SyntaxRule pulumi.IntPtrOutput `pulumi:"syntaxRule"`
}

// NewScheduledSql registers a new resource with the given unique name, arguments, and options.
func NewScheduledSql(ctx *pulumi.Context,
	name string, args *ScheduledSqlArgs, opts ...pulumi.ResourceOption) (*ScheduledSql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DstResource == nil {
		return nil, errors.New("invalid value for required argument 'DstResource'")
	}
	if args.EnableFlag == nil {
		return nil, errors.New("invalid value for required argument 'EnableFlag'")
	}
	if args.ProcessDelay == nil {
		return nil, errors.New("invalid value for required argument 'ProcessDelay'")
	}
	if args.ProcessPeriod == nil {
		return nil, errors.New("invalid value for required argument 'ProcessPeriod'")
	}
	if args.ProcessStartTime == nil {
		return nil, errors.New("invalid value for required argument 'ProcessStartTime'")
	}
	if args.ProcessTimeWindow == nil {
		return nil, errors.New("invalid value for required argument 'ProcessTimeWindow'")
	}
	if args.ProcessType == nil {
		return nil, errors.New("invalid value for required argument 'ProcessType'")
	}
	if args.ScheduledSqlContent == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledSqlContent'")
	}
	if args.SrcTopicId == nil {
		return nil, errors.New("invalid value for required argument 'SrcTopicId'")
	}
	if args.SrcTopicRegion == nil {
		return nil, errors.New("invalid value for required argument 'SrcTopicRegion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScheduledSql
	err := ctx.RegisterResource("tencentcloud:Cls/scheduledSql:ScheduledSql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledSql gets an existing ScheduledSql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledSql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledSqlState, opts ...pulumi.ResourceOption) (*ScheduledSql, error) {
	var resource ScheduledSql
	err := ctx.ReadResource("tencentcloud:Cls/scheduledSql:ScheduledSql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledSql resources.
type scheduledSqlState struct {
	// scheduled slq dst resource.
	DstResource *ScheduledSqlDstResource `pulumi:"dstResource"`
	// task enable flag.
	EnableFlag *int `pulumi:"enableFlag"`
	// task name.
	Name *string `pulumi:"name"`
	// process delay.
	ProcessDelay *int `pulumi:"processDelay"`
	// process end timestamp.
	ProcessEndTime *int `pulumi:"processEndTime"`
	// process period.
	ProcessPeriod *int `pulumi:"processPeriod"`
	// process start timestamp.
	ProcessStartTime *int `pulumi:"processStartTime"`
	// process time window.
	ProcessTimeWindow *string `pulumi:"processTimeWindow"`
	// process type.
	ProcessType *int `pulumi:"processType"`
	// scheduled sql content.
	ScheduledSqlContent *string `pulumi:"scheduledSqlContent"`
	// src topic id.
	SrcTopicId *string `pulumi:"srcTopicId"`
	// src topic region.
	SrcTopicRegion *string `pulumi:"srcTopicRegion"`
	// syntax rule.
	SyntaxRule *int `pulumi:"syntaxRule"`
}

type ScheduledSqlState struct {
	// scheduled slq dst resource.
	DstResource ScheduledSqlDstResourcePtrInput
	// task enable flag.
	EnableFlag pulumi.IntPtrInput
	// task name.
	Name pulumi.StringPtrInput
	// process delay.
	ProcessDelay pulumi.IntPtrInput
	// process end timestamp.
	ProcessEndTime pulumi.IntPtrInput
	// process period.
	ProcessPeriod pulumi.IntPtrInput
	// process start timestamp.
	ProcessStartTime pulumi.IntPtrInput
	// process time window.
	ProcessTimeWindow pulumi.StringPtrInput
	// process type.
	ProcessType pulumi.IntPtrInput
	// scheduled sql content.
	ScheduledSqlContent pulumi.StringPtrInput
	// src topic id.
	SrcTopicId pulumi.StringPtrInput
	// src topic region.
	SrcTopicRegion pulumi.StringPtrInput
	// syntax rule.
	SyntaxRule pulumi.IntPtrInput
}

func (ScheduledSqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSqlState)(nil)).Elem()
}

type scheduledSqlArgs struct {
	// scheduled slq dst resource.
	DstResource ScheduledSqlDstResource `pulumi:"dstResource"`
	// task enable flag.
	EnableFlag int `pulumi:"enableFlag"`
	// task name.
	Name *string `pulumi:"name"`
	// process delay.
	ProcessDelay int `pulumi:"processDelay"`
	// process end timestamp.
	ProcessEndTime *int `pulumi:"processEndTime"`
	// process period.
	ProcessPeriod int `pulumi:"processPeriod"`
	// process start timestamp.
	ProcessStartTime int `pulumi:"processStartTime"`
	// process time window.
	ProcessTimeWindow string `pulumi:"processTimeWindow"`
	// process type.
	ProcessType int `pulumi:"processType"`
	// scheduled sql content.
	ScheduledSqlContent string `pulumi:"scheduledSqlContent"`
	// src topic id.
	SrcTopicId string `pulumi:"srcTopicId"`
	// src topic region.
	SrcTopicRegion string `pulumi:"srcTopicRegion"`
	// syntax rule.
	SyntaxRule *int `pulumi:"syntaxRule"`
}

// The set of arguments for constructing a ScheduledSql resource.
type ScheduledSqlArgs struct {
	// scheduled slq dst resource.
	DstResource ScheduledSqlDstResourceInput
	// task enable flag.
	EnableFlag pulumi.IntInput
	// task name.
	Name pulumi.StringPtrInput
	// process delay.
	ProcessDelay pulumi.IntInput
	// process end timestamp.
	ProcessEndTime pulumi.IntPtrInput
	// process period.
	ProcessPeriod pulumi.IntInput
	// process start timestamp.
	ProcessStartTime pulumi.IntInput
	// process time window.
	ProcessTimeWindow pulumi.StringInput
	// process type.
	ProcessType pulumi.IntInput
	// scheduled sql content.
	ScheduledSqlContent pulumi.StringInput
	// src topic id.
	SrcTopicId pulumi.StringInput
	// src topic region.
	SrcTopicRegion pulumi.StringInput
	// syntax rule.
	SyntaxRule pulumi.IntPtrInput
}

func (ScheduledSqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSqlArgs)(nil)).Elem()
}

type ScheduledSqlInput interface {
	pulumi.Input

	ToScheduledSqlOutput() ScheduledSqlOutput
	ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput
}

func (*ScheduledSql) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledSql)(nil)).Elem()
}

func (i *ScheduledSql) ToScheduledSqlOutput() ScheduledSqlOutput {
	return i.ToScheduledSqlOutputWithContext(context.Background())
}

func (i *ScheduledSql) ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlOutput)
}

// ScheduledSqlArrayInput is an input type that accepts ScheduledSqlArray and ScheduledSqlArrayOutput values.
// You can construct a concrete instance of `ScheduledSqlArrayInput` via:
//
//          ScheduledSqlArray{ ScheduledSqlArgs{...} }
type ScheduledSqlArrayInput interface {
	pulumi.Input

	ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput
	ToScheduledSqlArrayOutputWithContext(context.Context) ScheduledSqlArrayOutput
}

type ScheduledSqlArray []ScheduledSqlInput

func (ScheduledSqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledSql)(nil)).Elem()
}

func (i ScheduledSqlArray) ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput {
	return i.ToScheduledSqlArrayOutputWithContext(context.Background())
}

func (i ScheduledSqlArray) ToScheduledSqlArrayOutputWithContext(ctx context.Context) ScheduledSqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlArrayOutput)
}

// ScheduledSqlMapInput is an input type that accepts ScheduledSqlMap and ScheduledSqlMapOutput values.
// You can construct a concrete instance of `ScheduledSqlMapInput` via:
//
//          ScheduledSqlMap{ "key": ScheduledSqlArgs{...} }
type ScheduledSqlMapInput interface {
	pulumi.Input

	ToScheduledSqlMapOutput() ScheduledSqlMapOutput
	ToScheduledSqlMapOutputWithContext(context.Context) ScheduledSqlMapOutput
}

type ScheduledSqlMap map[string]ScheduledSqlInput

func (ScheduledSqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledSql)(nil)).Elem()
}

func (i ScheduledSqlMap) ToScheduledSqlMapOutput() ScheduledSqlMapOutput {
	return i.ToScheduledSqlMapOutputWithContext(context.Background())
}

func (i ScheduledSqlMap) ToScheduledSqlMapOutputWithContext(ctx context.Context) ScheduledSqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlMapOutput)
}

type ScheduledSqlOutput struct{ *pulumi.OutputState }

func (ScheduledSqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlOutput) ToScheduledSqlOutput() ScheduledSqlOutput {
	return o
}

func (o ScheduledSqlOutput) ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput {
	return o
}

// scheduled slq dst resource.
func (o ScheduledSqlOutput) DstResource() ScheduledSqlDstResourceOutput {
	return o.ApplyT(func(v *ScheduledSql) ScheduledSqlDstResourceOutput { return v.DstResource }).(ScheduledSqlDstResourceOutput)
}

// task enable flag.
func (o ScheduledSqlOutput) EnableFlag() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.EnableFlag }).(pulumi.IntOutput)
}

// task name.
func (o ScheduledSqlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// process delay.
func (o ScheduledSqlOutput) ProcessDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.ProcessDelay }).(pulumi.IntOutput)
}

// process end timestamp.
func (o ScheduledSqlOutput) ProcessEndTime() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.ProcessEndTime }).(pulumi.IntOutput)
}

// process period.
func (o ScheduledSqlOutput) ProcessPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.ProcessPeriod }).(pulumi.IntOutput)
}

// process start timestamp.
func (o ScheduledSqlOutput) ProcessStartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.ProcessStartTime }).(pulumi.IntOutput)
}

// process time window.
func (o ScheduledSqlOutput) ProcessTimeWindow() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.ProcessTimeWindow }).(pulumi.StringOutput)
}

// process type.
func (o ScheduledSqlOutput) ProcessType() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntOutput { return v.ProcessType }).(pulumi.IntOutput)
}

// scheduled sql content.
func (o ScheduledSqlOutput) ScheduledSqlContent() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.ScheduledSqlContent }).(pulumi.StringOutput)
}

// src topic id.
func (o ScheduledSqlOutput) SrcTopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.SrcTopicId }).(pulumi.StringOutput)
}

// src topic region.
func (o ScheduledSqlOutput) SrcTopicRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.SrcTopicRegion }).(pulumi.StringOutput)
}

// syntax rule.
func (o ScheduledSqlOutput) SyntaxRule() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.IntPtrOutput { return v.SyntaxRule }).(pulumi.IntPtrOutput)
}

type ScheduledSqlArrayOutput struct{ *pulumi.OutputState }

func (ScheduledSqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlArrayOutput) ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput {
	return o
}

func (o ScheduledSqlArrayOutput) ToScheduledSqlArrayOutputWithContext(ctx context.Context) ScheduledSqlArrayOutput {
	return o
}

func (o ScheduledSqlArrayOutput) Index(i pulumi.IntInput) ScheduledSqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScheduledSql {
		return vs[0].([]*ScheduledSql)[vs[1].(int)]
	}).(ScheduledSqlOutput)
}

type ScheduledSqlMapOutput struct{ *pulumi.OutputState }

func (ScheduledSqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlMapOutput) ToScheduledSqlMapOutput() ScheduledSqlMapOutput {
	return o
}

func (o ScheduledSqlMapOutput) ToScheduledSqlMapOutputWithContext(ctx context.Context) ScheduledSqlMapOutput {
	return o
}

func (o ScheduledSqlMapOutput) MapIndex(k pulumi.StringInput) ScheduledSqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScheduledSql {
		return vs[0].(map[string]*ScheduledSql)[vs[1].(string)]
	}).(ScheduledSqlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlInput)(nil)).Elem(), &ScheduledSql{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlArrayInput)(nil)).Elem(), ScheduledSqlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlMapInput)(nil)).Elem(), ScheduledSqlMap{})
	pulumi.RegisterOutputType(ScheduledSqlOutput{})
	pulumi.RegisterOutputType(ScheduledSqlArrayOutput{})
	pulumi.RegisterOutputType(ScheduledSqlMapOutput{})
}
