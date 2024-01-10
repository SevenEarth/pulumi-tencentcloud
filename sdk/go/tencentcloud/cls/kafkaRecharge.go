// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cls kafkaRecharge
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			logset, err := Cls.NewLogset(ctx, "logset", &Cls.LogsetArgs{
//				LogsetName: pulumi.String("tf-example-logset"),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			topic, err := Cls.NewTopic(ctx, "topic", &Cls.TopicArgs{
//				TopicName:          pulumi.String("tf-example-topic"),
//				LogsetId:           logset.ID(),
//				AutoSplit:          pulumi.Bool(false),
//				MaxSplitPartitions: pulumi.Int(20),
//				PartitionCount:     pulumi.Int(1),
//				Period:             pulumi.Int(10),
//				StorageType:        pulumi.String("hot"),
//				Tags: pulumi.AnyMap{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Cls.NewKafkaRecharge(ctx, "kafkaRecharge", &Cls.KafkaRechargeArgs{
//				TopicId:          topic.ID(),
//				KafkaType:        pulumi.Int(0),
//				Offset:           -2,
//				IsEncryptionAddr: pulumi.Bool(true),
//				UserKafkaTopics:  pulumi.String("recharge"),
//				KafkaInstance:    pulumi.String("ckafka-qzoeaqx8"),
//				LogRechargeRule: &cls.KafkaRechargeLogRechargeRuleArgs{
//					RechargeType:      pulumi.String("json_log"),
//					EncodingFormat:    pulumi.Int(0),
//					DefaultTimeSwitch: pulumi.Bool(true),
//				},
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
// cls kafka_recharge can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cls/kafkaRecharge:KafkaRecharge kafka_recharge kafka_recharge_id
//
// ```
type KafkaRecharge struct {
	pulumi.CustomResourceState

	// user consumer group name.
	ConsumerGroupName pulumi.StringPtrOutput `pulumi:"consumerGroupName"`
	// ServerAddr is encryption addr.
	IsEncryptionAddr pulumi.BoolOutput `pulumi:"isEncryptionAddr"`
	// CKafka Instance id.
	KafkaInstance pulumi.StringPtrOutput `pulumi:"kafkaInstance"`
	// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
	KafkaType pulumi.IntOutput `pulumi:"kafkaType"`
	// log recharge rule.
	LogRechargeRule KafkaRechargeLogRechargeRuleOutput `pulumi:"logRechargeRule"`
	// kafka recharge name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The translation is: -2: Earliest (default) -1: Latest.
	Offset pulumi.IntOutput `pulumi:"offset"`
	// encryption protocol.
	Protocol KafkaRechargeProtocolOutput `pulumi:"protocol"`
	// Server addr.
	ServerAddr pulumi.StringPtrOutput `pulumi:"serverAddr"`
	// recharge for cls TopicId.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
	// user need recharge kafka topic list.
	UserKafkaTopics pulumi.StringOutput `pulumi:"userKafkaTopics"`
}

// NewKafkaRecharge registers a new resource with the given unique name, arguments, and options.
func NewKafkaRecharge(ctx *pulumi.Context,
	name string, args *KafkaRechargeArgs, opts ...pulumi.ResourceOption) (*KafkaRecharge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KafkaType == nil {
		return nil, errors.New("invalid value for required argument 'KafkaType'")
	}
	if args.Offset == nil {
		return nil, errors.New("invalid value for required argument 'Offset'")
	}
	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	if args.UserKafkaTopics == nil {
		return nil, errors.New("invalid value for required argument 'UserKafkaTopics'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KafkaRecharge
	err := ctx.RegisterResource("tencentcloud:Cls/kafkaRecharge:KafkaRecharge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaRecharge gets an existing KafkaRecharge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaRecharge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaRechargeState, opts ...pulumi.ResourceOption) (*KafkaRecharge, error) {
	var resource KafkaRecharge
	err := ctx.ReadResource("tencentcloud:Cls/kafkaRecharge:KafkaRecharge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaRecharge resources.
type kafkaRechargeState struct {
	// user consumer group name.
	ConsumerGroupName *string `pulumi:"consumerGroupName"`
	// ServerAddr is encryption addr.
	IsEncryptionAddr *bool `pulumi:"isEncryptionAddr"`
	// CKafka Instance id.
	KafkaInstance *string `pulumi:"kafkaInstance"`
	// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
	KafkaType *int `pulumi:"kafkaType"`
	// log recharge rule.
	LogRechargeRule *KafkaRechargeLogRechargeRule `pulumi:"logRechargeRule"`
	// kafka recharge name.
	Name *string `pulumi:"name"`
	// The translation is: -2: Earliest (default) -1: Latest.
	Offset *int `pulumi:"offset"`
	// encryption protocol.
	Protocol *KafkaRechargeProtocol `pulumi:"protocol"`
	// Server addr.
	ServerAddr *string `pulumi:"serverAddr"`
	// recharge for cls TopicId.
	TopicId *string `pulumi:"topicId"`
	// user need recharge kafka topic list.
	UserKafkaTopics *string `pulumi:"userKafkaTopics"`
}

type KafkaRechargeState struct {
	// user consumer group name.
	ConsumerGroupName pulumi.StringPtrInput
	// ServerAddr is encryption addr.
	IsEncryptionAddr pulumi.BoolPtrInput
	// CKafka Instance id.
	KafkaInstance pulumi.StringPtrInput
	// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
	KafkaType pulumi.IntPtrInput
	// log recharge rule.
	LogRechargeRule KafkaRechargeLogRechargeRulePtrInput
	// kafka recharge name.
	Name pulumi.StringPtrInput
	// The translation is: -2: Earliest (default) -1: Latest.
	Offset pulumi.IntPtrInput
	// encryption protocol.
	Protocol KafkaRechargeProtocolPtrInput
	// Server addr.
	ServerAddr pulumi.StringPtrInput
	// recharge for cls TopicId.
	TopicId pulumi.StringPtrInput
	// user need recharge kafka topic list.
	UserKafkaTopics pulumi.StringPtrInput
}

func (KafkaRechargeState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaRechargeState)(nil)).Elem()
}

type kafkaRechargeArgs struct {
	// user consumer group name.
	ConsumerGroupName *string `pulumi:"consumerGroupName"`
	// ServerAddr is encryption addr.
	IsEncryptionAddr *bool `pulumi:"isEncryptionAddr"`
	// CKafka Instance id.
	KafkaInstance *string `pulumi:"kafkaInstance"`
	// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
	KafkaType int `pulumi:"kafkaType"`
	// log recharge rule.
	LogRechargeRule *KafkaRechargeLogRechargeRule `pulumi:"logRechargeRule"`
	// kafka recharge name.
	Name *string `pulumi:"name"`
	// The translation is: -2: Earliest (default) -1: Latest.
	Offset int `pulumi:"offset"`
	// encryption protocol.
	Protocol *KafkaRechargeProtocol `pulumi:"protocol"`
	// Server addr.
	ServerAddr *string `pulumi:"serverAddr"`
	// recharge for cls TopicId.
	TopicId string `pulumi:"topicId"`
	// user need recharge kafka topic list.
	UserKafkaTopics string `pulumi:"userKafkaTopics"`
}

// The set of arguments for constructing a KafkaRecharge resource.
type KafkaRechargeArgs struct {
	// user consumer group name.
	ConsumerGroupName pulumi.StringPtrInput
	// ServerAddr is encryption addr.
	IsEncryptionAddr pulumi.BoolPtrInput
	// CKafka Instance id.
	KafkaInstance pulumi.StringPtrInput
	// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
	KafkaType pulumi.IntInput
	// log recharge rule.
	LogRechargeRule KafkaRechargeLogRechargeRulePtrInput
	// kafka recharge name.
	Name pulumi.StringPtrInput
	// The translation is: -2: Earliest (default) -1: Latest.
	Offset pulumi.IntInput
	// encryption protocol.
	Protocol KafkaRechargeProtocolPtrInput
	// Server addr.
	ServerAddr pulumi.StringPtrInput
	// recharge for cls TopicId.
	TopicId pulumi.StringInput
	// user need recharge kafka topic list.
	UserKafkaTopics pulumi.StringInput
}

func (KafkaRechargeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaRechargeArgs)(nil)).Elem()
}

type KafkaRechargeInput interface {
	pulumi.Input

	ToKafkaRechargeOutput() KafkaRechargeOutput
	ToKafkaRechargeOutputWithContext(ctx context.Context) KafkaRechargeOutput
}

func (*KafkaRecharge) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaRecharge)(nil)).Elem()
}

func (i *KafkaRecharge) ToKafkaRechargeOutput() KafkaRechargeOutput {
	return i.ToKafkaRechargeOutputWithContext(context.Background())
}

func (i *KafkaRecharge) ToKafkaRechargeOutputWithContext(ctx context.Context) KafkaRechargeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRechargeOutput)
}

// KafkaRechargeArrayInput is an input type that accepts KafkaRechargeArray and KafkaRechargeArrayOutput values.
// You can construct a concrete instance of `KafkaRechargeArrayInput` via:
//
//	KafkaRechargeArray{ KafkaRechargeArgs{...} }
type KafkaRechargeArrayInput interface {
	pulumi.Input

	ToKafkaRechargeArrayOutput() KafkaRechargeArrayOutput
	ToKafkaRechargeArrayOutputWithContext(context.Context) KafkaRechargeArrayOutput
}

type KafkaRechargeArray []KafkaRechargeInput

func (KafkaRechargeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaRecharge)(nil)).Elem()
}

func (i KafkaRechargeArray) ToKafkaRechargeArrayOutput() KafkaRechargeArrayOutput {
	return i.ToKafkaRechargeArrayOutputWithContext(context.Background())
}

func (i KafkaRechargeArray) ToKafkaRechargeArrayOutputWithContext(ctx context.Context) KafkaRechargeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRechargeArrayOutput)
}

// KafkaRechargeMapInput is an input type that accepts KafkaRechargeMap and KafkaRechargeMapOutput values.
// You can construct a concrete instance of `KafkaRechargeMapInput` via:
//
//	KafkaRechargeMap{ "key": KafkaRechargeArgs{...} }
type KafkaRechargeMapInput interface {
	pulumi.Input

	ToKafkaRechargeMapOutput() KafkaRechargeMapOutput
	ToKafkaRechargeMapOutputWithContext(context.Context) KafkaRechargeMapOutput
}

type KafkaRechargeMap map[string]KafkaRechargeInput

func (KafkaRechargeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaRecharge)(nil)).Elem()
}

func (i KafkaRechargeMap) ToKafkaRechargeMapOutput() KafkaRechargeMapOutput {
	return i.ToKafkaRechargeMapOutputWithContext(context.Background())
}

func (i KafkaRechargeMap) ToKafkaRechargeMapOutputWithContext(ctx context.Context) KafkaRechargeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRechargeMapOutput)
}

type KafkaRechargeOutput struct{ *pulumi.OutputState }

func (KafkaRechargeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaRecharge)(nil)).Elem()
}

func (o KafkaRechargeOutput) ToKafkaRechargeOutput() KafkaRechargeOutput {
	return o
}

func (o KafkaRechargeOutput) ToKafkaRechargeOutputWithContext(ctx context.Context) KafkaRechargeOutput {
	return o
}

// user consumer group name.
func (o KafkaRechargeOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringPtrOutput { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

// ServerAddr is encryption addr.
func (o KafkaRechargeOutput) IsEncryptionAddr() pulumi.BoolOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.BoolOutput { return v.IsEncryptionAddr }).(pulumi.BoolOutput)
}

// CKafka Instance id.
func (o KafkaRechargeOutput) KafkaInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringPtrOutput { return v.KafkaInstance }).(pulumi.StringPtrOutput)
}

// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
func (o KafkaRechargeOutput) KafkaType() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.IntOutput { return v.KafkaType }).(pulumi.IntOutput)
}

// log recharge rule.
func (o KafkaRechargeOutput) LogRechargeRule() KafkaRechargeLogRechargeRuleOutput {
	return o.ApplyT(func(v *KafkaRecharge) KafkaRechargeLogRechargeRuleOutput { return v.LogRechargeRule }).(KafkaRechargeLogRechargeRuleOutput)
}

// kafka recharge name.
func (o KafkaRechargeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The translation is: -2: Earliest (default) -1: Latest.
func (o KafkaRechargeOutput) Offset() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.IntOutput { return v.Offset }).(pulumi.IntOutput)
}

// encryption protocol.
func (o KafkaRechargeOutput) Protocol() KafkaRechargeProtocolOutput {
	return o.ApplyT(func(v *KafkaRecharge) KafkaRechargeProtocolOutput { return v.Protocol }).(KafkaRechargeProtocolOutput)
}

// Server addr.
func (o KafkaRechargeOutput) ServerAddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringPtrOutput { return v.ServerAddr }).(pulumi.StringPtrOutput)
}

// recharge for cls TopicId.
func (o KafkaRechargeOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

// user need recharge kafka topic list.
func (o KafkaRechargeOutput) UserKafkaTopics() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaRecharge) pulumi.StringOutput { return v.UserKafkaTopics }).(pulumi.StringOutput)
}

type KafkaRechargeArrayOutput struct{ *pulumi.OutputState }

func (KafkaRechargeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaRecharge)(nil)).Elem()
}

func (o KafkaRechargeArrayOutput) ToKafkaRechargeArrayOutput() KafkaRechargeArrayOutput {
	return o
}

func (o KafkaRechargeArrayOutput) ToKafkaRechargeArrayOutputWithContext(ctx context.Context) KafkaRechargeArrayOutput {
	return o
}

func (o KafkaRechargeArrayOutput) Index(i pulumi.IntInput) KafkaRechargeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaRecharge {
		return vs[0].([]*KafkaRecharge)[vs[1].(int)]
	}).(KafkaRechargeOutput)
}

type KafkaRechargeMapOutput struct{ *pulumi.OutputState }

func (KafkaRechargeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaRecharge)(nil)).Elem()
}

func (o KafkaRechargeMapOutput) ToKafkaRechargeMapOutput() KafkaRechargeMapOutput {
	return o
}

func (o KafkaRechargeMapOutput) ToKafkaRechargeMapOutputWithContext(ctx context.Context) KafkaRechargeMapOutput {
	return o
}

func (o KafkaRechargeMapOutput) MapIndex(k pulumi.StringInput) KafkaRechargeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaRecharge {
		return vs[0].(map[string]*KafkaRecharge)[vs[1].(string)]
	}).(KafkaRechargeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaRechargeInput)(nil)).Elem(), &KafkaRecharge{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaRechargeArrayInput)(nil)).Elem(), KafkaRechargeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaRechargeMapInput)(nil)).Elem(), KafkaRechargeMap{})
	pulumi.RegisterOutputType(KafkaRechargeOutput{})
	pulumi.RegisterOutputType(KafkaRechargeArrayOutput{})
	pulumi.RegisterOutputType(KafkaRechargeMapOutput{})
}
