// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create ckafka topic.
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
// 		_, err := Ckafka.NewTopic(ctx, "foo", &Ckafka.TopicArgs{
// 			CleanUpPolicy:   pulumi.String("delete"),
// 			EnableWhiteList: pulumi.Bool(true),
// 			InstanceId:      pulumi.String("ckafka-f9ife4zz"),
// 			IpWhiteLists: pulumi.StringArray{
// 				pulumi.String("ip1"),
// 				pulumi.String("ip2"),
// 			},
// 			MaxMessageBytes:             pulumi.Int(0),
// 			Note:                        pulumi.String("topic note"),
// 			PartitionNum:                pulumi.Int(1),
// 			ReplicaNum:                  pulumi.Int(2),
// 			Retention:                   pulumi.Int(60000),
// 			Segment:                     pulumi.Int(3600000),
// 			SyncReplicaMinNum:           pulumi.Int(1),
// 			TopicName:                   pulumi.String("example"),
// 			UncleanLeaderElectionEnable: pulumi.Bool(false),
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
// ckafka topic can be imported using the instance_id#topic_name, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ckafka/topic:Topic foo ckafka-f9ife4zz#example
// ```
type Topic struct {
	pulumi.CustomResourceState

	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy pulumi.StringPtrOutput `pulumi:"cleanUpPolicy"`
	// Create time of the CKafka topic.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList pulumi.BoolPtrOutput `pulumi:"enableWhiteList"`
	// Data backup cos bucket: the bucket address that is dumped to cos.
	ForwardCosBucket pulumi.StringOutput `pulumi:"forwardCosBucket"`
	// Periodic frequency of data backup to cos.
	ForwardInterval pulumi.IntOutput `pulumi:"forwardInterval"`
	// Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
	ForwardStatus pulumi.IntOutput `pulumi:"forwardStatus"`
	// Ckafka instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IpWhiteLists pulumi.StringArrayOutput `pulumi:"ipWhiteLists"`
	// Max message bytes.
	MaxMessageBytes pulumi.IntPtrOutput `pulumi:"maxMessageBytes"`
	// Message storage location.
	MessageStorageLocation pulumi.StringOutput `pulumi:"messageStorageLocation"`
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// The number of partition.
	PartitionNum pulumi.IntOutput `pulumi:"partitionNum"`
	// The number of replica.
	ReplicaNum pulumi.IntOutput `pulumi:"replicaNum"`
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention pulumi.IntPtrOutput `pulumi:"retention"`
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment pulumi.IntPtrOutput `pulumi:"segment"`
	// Number of bytes rolled by shard.
	SegmentBytes pulumi.IntOutput `pulumi:"segmentBytes"`
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum pulumi.IntPtrOutput `pulumi:"syncReplicaMinNum"`
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable pulumi.BoolPtrOutput `pulumi:"uncleanLeaderElectionEnable"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PartitionNum == nil {
		return nil, errors.New("invalid value for required argument 'PartitionNum'")
	}
	if args.ReplicaNum == nil {
		return nil, errors.New("invalid value for required argument 'ReplicaNum'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource Topic
	err := ctx.RegisterResource("tencentcloud:Ckafka/topic:Topic", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Ckafka/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy *string `pulumi:"cleanUpPolicy"`
	// Create time of the CKafka topic.
	CreateTime *string `pulumi:"createTime"`
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList *bool `pulumi:"enableWhiteList"`
	// Data backup cos bucket: the bucket address that is dumped to cos.
	ForwardCosBucket *string `pulumi:"forwardCosBucket"`
	// Periodic frequency of data backup to cos.
	ForwardInterval *int `pulumi:"forwardInterval"`
	// Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
	ForwardStatus *int `pulumi:"forwardStatus"`
	// Ckafka instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IpWhiteLists []string `pulumi:"ipWhiteLists"`
	// Max message bytes.
	MaxMessageBytes *int `pulumi:"maxMessageBytes"`
	// Message storage location.
	MessageStorageLocation *string `pulumi:"messageStorageLocation"`
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note *string `pulumi:"note"`
	// The number of partition.
	PartitionNum *int `pulumi:"partitionNum"`
	// The number of replica.
	ReplicaNum *int `pulumi:"replicaNum"`
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention *int `pulumi:"retention"`
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment *int `pulumi:"segment"`
	// Number of bytes rolled by shard.
	SegmentBytes *int `pulumi:"segmentBytes"`
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum *int `pulumi:"syncReplicaMinNum"`
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName *string `pulumi:"topicName"`
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable *bool `pulumi:"uncleanLeaderElectionEnable"`
}

type TopicState struct {
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy pulumi.StringPtrInput
	// Create time of the CKafka topic.
	CreateTime pulumi.StringPtrInput
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList pulumi.BoolPtrInput
	// Data backup cos bucket: the bucket address that is dumped to cos.
	ForwardCosBucket pulumi.StringPtrInput
	// Periodic frequency of data backup to cos.
	ForwardInterval pulumi.IntPtrInput
	// Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
	ForwardStatus pulumi.IntPtrInput
	// Ckafka instance ID.
	InstanceId pulumi.StringPtrInput
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IpWhiteLists pulumi.StringArrayInput
	// Max message bytes.
	MaxMessageBytes pulumi.IntPtrInput
	// Message storage location.
	MessageStorageLocation pulumi.StringPtrInput
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note pulumi.StringPtrInput
	// The number of partition.
	PartitionNum pulumi.IntPtrInput
	// The number of replica.
	ReplicaNum pulumi.IntPtrInput
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention pulumi.IntPtrInput
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment pulumi.IntPtrInput
	// Number of bytes rolled by shard.
	SegmentBytes pulumi.IntPtrInput
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum pulumi.IntPtrInput
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName pulumi.StringPtrInput
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable pulumi.BoolPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy *string `pulumi:"cleanUpPolicy"`
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList *bool `pulumi:"enableWhiteList"`
	// Ckafka instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IpWhiteLists []string `pulumi:"ipWhiteLists"`
	// Max message bytes.
	MaxMessageBytes *int `pulumi:"maxMessageBytes"`
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note *string `pulumi:"note"`
	// The number of partition.
	PartitionNum int `pulumi:"partitionNum"`
	// The number of replica.
	ReplicaNum int `pulumi:"replicaNum"`
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention *int `pulumi:"retention"`
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment *int `pulumi:"segment"`
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum *int `pulumi:"syncReplicaMinNum"`
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName string `pulumi:"topicName"`
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable *bool `pulumi:"uncleanLeaderElectionEnable"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	CleanUpPolicy pulumi.StringPtrInput
	// Whether to open the ip whitelist, `true`: open, `false`: close.
	EnableWhiteList pulumi.BoolPtrInput
	// Ckafka instance ID.
	InstanceId pulumi.StringInput
	// Ip whitelist, quota limit, required when enableWhileList=true.
	IpWhiteLists pulumi.StringArrayInput
	// Max message bytes.
	MaxMessageBytes pulumi.IntPtrInput
	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	Note pulumi.StringPtrInput
	// The number of partition.
	PartitionNum pulumi.IntInput
	// The number of replica.
	ReplicaNum pulumi.IntInput
	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	Retention pulumi.IntPtrInput
	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	Segment pulumi.IntPtrInput
	// Min number of sync replicas, Default is `1`.
	SyncReplicaMinNum pulumi.IntPtrInput
	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	TopicName pulumi.StringInput
	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
	UncleanLeaderElectionEnable pulumi.BoolPtrInput
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
//          TopicArray{ TopicArgs{...} }
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
//          TopicMap{ "key": TopicArgs{...} }
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

// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
func (o TopicOutput) CleanUpPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.CleanUpPolicy }).(pulumi.StringPtrOutput)
}

// Create time of the CKafka topic.
func (o TopicOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether to open the ip whitelist, `true`: open, `false`: close.
func (o TopicOutput) EnableWhiteList() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnableWhiteList }).(pulumi.BoolPtrOutput)
}

// Data backup cos bucket: the bucket address that is dumped to cos.
func (o TopicOutput) ForwardCosBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ForwardCosBucket }).(pulumi.StringOutput)
}

// Periodic frequency of data backup to cos.
func (o TopicOutput) ForwardInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.ForwardInterval }).(pulumi.IntOutput)
}

// Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
func (o TopicOutput) ForwardStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.ForwardStatus }).(pulumi.IntOutput)
}

// Ckafka instance ID.
func (o TopicOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Ip whitelist, quota limit, required when enableWhileList=true.
func (o TopicOutput) IpWhiteLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringArrayOutput { return v.IpWhiteLists }).(pulumi.StringArrayOutput)
}

// Max message bytes.
func (o TopicOutput) MaxMessageBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.MaxMessageBytes }).(pulumi.IntPtrOutput)
}

// Message storage location.
func (o TopicOutput) MessageStorageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.MessageStorageLocation }).(pulumi.StringOutput)
}

// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
func (o TopicOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// The number of partition.
func (o TopicOutput) PartitionNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.PartitionNum }).(pulumi.IntOutput)
}

// The number of replica.
func (o TopicOutput) ReplicaNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.ReplicaNum }).(pulumi.IntOutput)
}

// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
func (o TopicOutput) Retention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.Retention }).(pulumi.IntPtrOutput)
}

// Segment scrolling time, in ms, the current minimum is 3600000ms.
func (o TopicOutput) Segment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.Segment }).(pulumi.IntPtrOutput)
}

// Number of bytes rolled by shard.
func (o TopicOutput) SegmentBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.SegmentBytes }).(pulumi.IntOutput)
}

// Min number of sync replicas, Default is `1`.
func (o TopicOutput) SyncReplicaMinNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.SyncReplicaMinNum }).(pulumi.IntPtrOutput)
}

// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
func (o TopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, ` true:  `allowed, `false`: not allowed.
func (o TopicOutput) UncleanLeaderElectionEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.UncleanLeaderElectionEnable }).(pulumi.BoolPtrOutput)
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
