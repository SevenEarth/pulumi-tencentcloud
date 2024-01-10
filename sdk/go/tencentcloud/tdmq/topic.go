// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a TDMQ topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Tdmq.NewInstance(ctx, "exampleInstance", &Tdmq.InstanceArgs{
//				ClusterName: pulumi.String("tf_example"),
//				Remark:      pulumi.String("remark."),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNamespace, err := Tdmq.NewNamespace(ctx, "exampleNamespace", &Tdmq.NamespaceArgs{
//				EnvironName: pulumi.String("tf_example"),
//				MsgTtl:      pulumi.Int(300),
//				ClusterId:   exampleInstance.ID(),
//				RetentionPolicy: &tdmq.NamespaceRetentionPolicyArgs{
//					TimeInMinutes: pulumi.Int(60),
//					SizeInMb:      pulumi.Int(10),
//				},
//				Remark: pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tdmq.NewTopic(ctx, "exampleTopic", &Tdmq.TopicArgs{
//				EnvironId:       exampleNamespace.EnvironName,
//				ClusterId:       exampleInstance.ID(),
//				TopicName:       pulumi.String("tf-example-topic"),
//				Partitions:      pulumi.Int(6),
//				PulsarTopicType: pulumi.Int(3),
//				Remark:          pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Tdmq Topic can be imported, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tdmq/topic:Topic test topic_id
//
// ```
type Topic struct {
	pulumi.CustomResourceState

	// The Dedicated Cluster Id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation time of resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of tdmq namespace.
	EnvironId pulumi.StringOutput `pulumi:"environId"`
	// The partitions of topic.
	Partitions pulumi.IntOutput `pulumi:"partitions"`
	// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
	PulsarTopicType pulumi.IntOutput `pulumi:"pulsarTopicType"`
	// Description of the namespace.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The name of topic to be created.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
	//
	// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
	TopicType pulumi.IntOutput `pulumi:"topicType"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.EnvironId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironId'")
	}
	if args.Partitions == nil {
		return nil, errors.New("invalid value for required argument 'Partitions'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("tencentcloud:Tdmq/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("tencentcloud:Tdmq/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// The Dedicated Cluster Id.
	ClusterId *string `pulumi:"clusterId"`
	// Creation time of resource.
	CreateTime *string `pulumi:"createTime"`
	// The name of tdmq namespace.
	EnvironId *string `pulumi:"environId"`
	// The partitions of topic.
	Partitions *int `pulumi:"partitions"`
	// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
	PulsarTopicType *int `pulumi:"pulsarTopicType"`
	// Description of the namespace.
	Remark *string `pulumi:"remark"`
	// The name of topic to be created.
	TopicName *string `pulumi:"topicName"`
	// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
	//
	// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
	TopicType *int `pulumi:"topicType"`
}

type TopicState struct {
	// The Dedicated Cluster Id.
	ClusterId pulumi.StringPtrInput
	// Creation time of resource.
	CreateTime pulumi.StringPtrInput
	// The name of tdmq namespace.
	EnvironId pulumi.StringPtrInput
	// The partitions of topic.
	Partitions pulumi.IntPtrInput
	// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
	PulsarTopicType pulumi.IntPtrInput
	// Description of the namespace.
	Remark pulumi.StringPtrInput
	// The name of topic to be created.
	TopicName pulumi.StringPtrInput
	// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
	//
	// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
	TopicType pulumi.IntPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The Dedicated Cluster Id.
	ClusterId string `pulumi:"clusterId"`
	// The name of tdmq namespace.
	EnvironId string `pulumi:"environId"`
	// The partitions of topic.
	Partitions int `pulumi:"partitions"`
	// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
	PulsarTopicType *int `pulumi:"pulsarTopicType"`
	// Description of the namespace.
	Remark *string `pulumi:"remark"`
	// The name of topic to be created.
	TopicName string `pulumi:"topicName"`
	// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
	//
	// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
	TopicType *int `pulumi:"topicType"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The Dedicated Cluster Id.
	ClusterId pulumi.StringInput
	// The name of tdmq namespace.
	EnvironId pulumi.StringInput
	// The partitions of topic.
	Partitions pulumi.IntInput
	// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
	PulsarTopicType pulumi.IntPtrInput
	// Description of the namespace.
	Remark pulumi.StringPtrInput
	// The name of topic to be created.
	TopicName pulumi.StringInput
	// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
	//
	// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
	TopicType pulumi.IntPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//	TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//	TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// The Dedicated Cluster Id.
func (o TopicOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation time of resource.
func (o TopicOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The name of tdmq namespace.
func (o TopicOutput) EnvironId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.EnvironId }).(pulumi.StringOutput)
}

// The partitions of topic.
func (o TopicOutput) Partitions() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.Partitions }).(pulumi.IntOutput)
}

// Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
func (o TopicOutput) PulsarTopicType() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.PulsarTopicType }).(pulumi.IntOutput)
}

// Description of the namespace.
func (o TopicOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// The name of topic to be created.
func (o TopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

// This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
//
// Deprecated: This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
func (o TopicOutput) TopicType() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.TopicType }).(pulumi.IntOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].([]*Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].(map[string]*Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
