// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ckafka consumerGroupModifyOffset
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ckafka.NewConsumerGroupModifyOffset(ctx, "consumerGroupModifyOffset", &Ckafka.ConsumerGroupModifyOffsetArgs{
// 			Group:      pulumi.String("xxxxxx"),
// 			InstanceId: pulumi.String("ckafka-xxxxxx"),
// 			Offset:     pulumi.Int(0),
// 			Strategy:   pulumi.Int(2),
// 			Topics: pulumi.StringArray{
// 				pulumi.String("xxxxxx"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ConsumerGroupModifyOffset struct {
	pulumi.CustomResourceState

	// kafka group.
	Group pulumi.StringOutput `pulumi:"group"`
	// Kafka instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The offset location that needs to be reset. When strategy is 2, this field must be included.
	Offset pulumi.IntPtrOutput `pulumi:"offset"`
	// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
	Partitions pulumi.IntArrayOutput `pulumi:"partitions"`
	// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
	Shift pulumi.IntPtrOutput `pulumi:"shift"`
	// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
	ShiftTimestamp pulumi.IntPtrOutput `pulumi:"shiftTimestamp"`
	// Reset the policy of offset.
	// `0`: Move the offset forward or backward shift bar;
	// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
	// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
	Strategy pulumi.IntOutput `pulumi:"strategy"`
	// Indicates the topics that needs to be reset. Leave it empty means all.
	Topics pulumi.StringArrayOutput `pulumi:"topics"`
}

// NewConsumerGroupModifyOffset registers a new resource with the given unique name, arguments, and options.
func NewConsumerGroupModifyOffset(ctx *pulumi.Context,
	name string, args *ConsumerGroupModifyOffsetArgs, opts ...pulumi.ResourceOption) (*ConsumerGroupModifyOffset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ConsumerGroupModifyOffset
	err := ctx.RegisterResource("tencentcloud:Ckafka/consumerGroupModifyOffset:ConsumerGroupModifyOffset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerGroupModifyOffset gets an existing ConsumerGroupModifyOffset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerGroupModifyOffset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerGroupModifyOffsetState, opts ...pulumi.ResourceOption) (*ConsumerGroupModifyOffset, error) {
	var resource ConsumerGroupModifyOffset
	err := ctx.ReadResource("tencentcloud:Ckafka/consumerGroupModifyOffset:ConsumerGroupModifyOffset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerGroupModifyOffset resources.
type consumerGroupModifyOffsetState struct {
	// kafka group.
	Group *string `pulumi:"group"`
	// Kafka instance id.
	InstanceId *string `pulumi:"instanceId"`
	// The offset location that needs to be reset. When strategy is 2, this field must be included.
	Offset *int `pulumi:"offset"`
	// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
	Partitions []int `pulumi:"partitions"`
	// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
	Shift *int `pulumi:"shift"`
	// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
	ShiftTimestamp *int `pulumi:"shiftTimestamp"`
	// Reset the policy of offset.
	// `0`: Move the offset forward or backward shift bar;
	// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
	// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
	Strategy *int `pulumi:"strategy"`
	// Indicates the topics that needs to be reset. Leave it empty means all.
	Topics []string `pulumi:"topics"`
}

type ConsumerGroupModifyOffsetState struct {
	// kafka group.
	Group pulumi.StringPtrInput
	// Kafka instance id.
	InstanceId pulumi.StringPtrInput
	// The offset location that needs to be reset. When strategy is 2, this field must be included.
	Offset pulumi.IntPtrInput
	// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
	Partitions pulumi.IntArrayInput
	// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
	Shift pulumi.IntPtrInput
	// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
	ShiftTimestamp pulumi.IntPtrInput
	// Reset the policy of offset.
	// `0`: Move the offset forward or backward shift bar;
	// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
	// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
	Strategy pulumi.IntPtrInput
	// Indicates the topics that needs to be reset. Leave it empty means all.
	Topics pulumi.StringArrayInput
}

func (ConsumerGroupModifyOffsetState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupModifyOffsetState)(nil)).Elem()
}

type consumerGroupModifyOffsetArgs struct {
	// kafka group.
	Group string `pulumi:"group"`
	// Kafka instance id.
	InstanceId string `pulumi:"instanceId"`
	// The offset location that needs to be reset. When strategy is 2, this field must be included.
	Offset *int `pulumi:"offset"`
	// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
	Partitions []int `pulumi:"partitions"`
	// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
	Shift *int `pulumi:"shift"`
	// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
	ShiftTimestamp *int `pulumi:"shiftTimestamp"`
	// Reset the policy of offset.
	// `0`: Move the offset forward or backward shift bar;
	// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
	// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
	Strategy int `pulumi:"strategy"`
	// Indicates the topics that needs to be reset. Leave it empty means all.
	Topics []string `pulumi:"topics"`
}

// The set of arguments for constructing a ConsumerGroupModifyOffset resource.
type ConsumerGroupModifyOffsetArgs struct {
	// kafka group.
	Group pulumi.StringInput
	// Kafka instance id.
	InstanceId pulumi.StringInput
	// The offset location that needs to be reset. When strategy is 2, this field must be included.
	Offset pulumi.IntPtrInput
	// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
	Partitions pulumi.IntArrayInput
	// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
	Shift pulumi.IntPtrInput
	// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
	ShiftTimestamp pulumi.IntPtrInput
	// Reset the policy of offset.
	// `0`: Move the offset forward or backward shift bar;
	// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
	// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
	Strategy pulumi.IntInput
	// Indicates the topics that needs to be reset. Leave it empty means all.
	Topics pulumi.StringArrayInput
}

func (ConsumerGroupModifyOffsetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupModifyOffsetArgs)(nil)).Elem()
}

type ConsumerGroupModifyOffsetInput interface {
	pulumi.Input

	ToConsumerGroupModifyOffsetOutput() ConsumerGroupModifyOffsetOutput
	ToConsumerGroupModifyOffsetOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetOutput
}

func (*ConsumerGroupModifyOffset) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerGroupModifyOffset)(nil)).Elem()
}

func (i *ConsumerGroupModifyOffset) ToConsumerGroupModifyOffsetOutput() ConsumerGroupModifyOffsetOutput {
	return i.ToConsumerGroupModifyOffsetOutputWithContext(context.Background())
}

func (i *ConsumerGroupModifyOffset) ToConsumerGroupModifyOffsetOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerGroupModifyOffsetOutput)
}

// ConsumerGroupModifyOffsetArrayInput is an input type that accepts ConsumerGroupModifyOffsetArray and ConsumerGroupModifyOffsetArrayOutput values.
// You can construct a concrete instance of `ConsumerGroupModifyOffsetArrayInput` via:
//
//          ConsumerGroupModifyOffsetArray{ ConsumerGroupModifyOffsetArgs{...} }
type ConsumerGroupModifyOffsetArrayInput interface {
	pulumi.Input

	ToConsumerGroupModifyOffsetArrayOutput() ConsumerGroupModifyOffsetArrayOutput
	ToConsumerGroupModifyOffsetArrayOutputWithContext(context.Context) ConsumerGroupModifyOffsetArrayOutput
}

type ConsumerGroupModifyOffsetArray []ConsumerGroupModifyOffsetInput

func (ConsumerGroupModifyOffsetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerGroupModifyOffset)(nil)).Elem()
}

func (i ConsumerGroupModifyOffsetArray) ToConsumerGroupModifyOffsetArrayOutput() ConsumerGroupModifyOffsetArrayOutput {
	return i.ToConsumerGroupModifyOffsetArrayOutputWithContext(context.Background())
}

func (i ConsumerGroupModifyOffsetArray) ToConsumerGroupModifyOffsetArrayOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerGroupModifyOffsetArrayOutput)
}

// ConsumerGroupModifyOffsetMapInput is an input type that accepts ConsumerGroupModifyOffsetMap and ConsumerGroupModifyOffsetMapOutput values.
// You can construct a concrete instance of `ConsumerGroupModifyOffsetMapInput` via:
//
//          ConsumerGroupModifyOffsetMap{ "key": ConsumerGroupModifyOffsetArgs{...} }
type ConsumerGroupModifyOffsetMapInput interface {
	pulumi.Input

	ToConsumerGroupModifyOffsetMapOutput() ConsumerGroupModifyOffsetMapOutput
	ToConsumerGroupModifyOffsetMapOutputWithContext(context.Context) ConsumerGroupModifyOffsetMapOutput
}

type ConsumerGroupModifyOffsetMap map[string]ConsumerGroupModifyOffsetInput

func (ConsumerGroupModifyOffsetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerGroupModifyOffset)(nil)).Elem()
}

func (i ConsumerGroupModifyOffsetMap) ToConsumerGroupModifyOffsetMapOutput() ConsumerGroupModifyOffsetMapOutput {
	return i.ToConsumerGroupModifyOffsetMapOutputWithContext(context.Background())
}

func (i ConsumerGroupModifyOffsetMap) ToConsumerGroupModifyOffsetMapOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerGroupModifyOffsetMapOutput)
}

type ConsumerGroupModifyOffsetOutput struct{ *pulumi.OutputState }

func (ConsumerGroupModifyOffsetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerGroupModifyOffset)(nil)).Elem()
}

func (o ConsumerGroupModifyOffsetOutput) ToConsumerGroupModifyOffsetOutput() ConsumerGroupModifyOffsetOutput {
	return o
}

func (o ConsumerGroupModifyOffsetOutput) ToConsumerGroupModifyOffsetOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetOutput {
	return o
}

// kafka group.
func (o ConsumerGroupModifyOffsetOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Kafka instance id.
func (o ConsumerGroupModifyOffsetOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The offset location that needs to be reset. When strategy is 2, this field must be included.
func (o ConsumerGroupModifyOffsetOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.IntPtrOutput { return v.Offset }).(pulumi.IntPtrOutput)
}

// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
func (o ConsumerGroupModifyOffsetOutput) Partitions() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.IntArrayOutput { return v.Partitions }).(pulumi.IntArrayOutput)
}

// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
func (o ConsumerGroupModifyOffsetOutput) Shift() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.IntPtrOutput { return v.Shift }).(pulumi.IntPtrOutput)
}

// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
func (o ConsumerGroupModifyOffsetOutput) ShiftTimestamp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.IntPtrOutput { return v.ShiftTimestamp }).(pulumi.IntPtrOutput)
}

// Reset the policy of offset.
// `0`: Move the offset forward or backward shift bar;
// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
func (o ConsumerGroupModifyOffsetOutput) Strategy() pulumi.IntOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.IntOutput { return v.Strategy }).(pulumi.IntOutput)
}

// Indicates the topics that needs to be reset. Leave it empty means all.
func (o ConsumerGroupModifyOffsetOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConsumerGroupModifyOffset) pulumi.StringArrayOutput { return v.Topics }).(pulumi.StringArrayOutput)
}

type ConsumerGroupModifyOffsetArrayOutput struct{ *pulumi.OutputState }

func (ConsumerGroupModifyOffsetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerGroupModifyOffset)(nil)).Elem()
}

func (o ConsumerGroupModifyOffsetArrayOutput) ToConsumerGroupModifyOffsetArrayOutput() ConsumerGroupModifyOffsetArrayOutput {
	return o
}

func (o ConsumerGroupModifyOffsetArrayOutput) ToConsumerGroupModifyOffsetArrayOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetArrayOutput {
	return o
}

func (o ConsumerGroupModifyOffsetArrayOutput) Index(i pulumi.IntInput) ConsumerGroupModifyOffsetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumerGroupModifyOffset {
		return vs[0].([]*ConsumerGroupModifyOffset)[vs[1].(int)]
	}).(ConsumerGroupModifyOffsetOutput)
}

type ConsumerGroupModifyOffsetMapOutput struct{ *pulumi.OutputState }

func (ConsumerGroupModifyOffsetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerGroupModifyOffset)(nil)).Elem()
}

func (o ConsumerGroupModifyOffsetMapOutput) ToConsumerGroupModifyOffsetMapOutput() ConsumerGroupModifyOffsetMapOutput {
	return o
}

func (o ConsumerGroupModifyOffsetMapOutput) ToConsumerGroupModifyOffsetMapOutputWithContext(ctx context.Context) ConsumerGroupModifyOffsetMapOutput {
	return o
}

func (o ConsumerGroupModifyOffsetMapOutput) MapIndex(k pulumi.StringInput) ConsumerGroupModifyOffsetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumerGroupModifyOffset {
		return vs[0].(map[string]*ConsumerGroupModifyOffset)[vs[1].(string)]
	}).(ConsumerGroupModifyOffsetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerGroupModifyOffsetInput)(nil)).Elem(), &ConsumerGroupModifyOffset{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerGroupModifyOffsetArrayInput)(nil)).Elem(), ConsumerGroupModifyOffsetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerGroupModifyOffsetMapInput)(nil)).Elem(), ConsumerGroupModifyOffsetMap{})
	pulumi.RegisterOutputType(ConsumerGroupModifyOffsetOutput{})
	pulumi.RegisterOutputType(ConsumerGroupModifyOffsetArrayOutput{})
	pulumi.RegisterOutputType(ConsumerGroupModifyOffsetMapOutput{})
}
