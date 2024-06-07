// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tdmq subscription
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Tdmq.NewInstance(ctx, "exampleInstance", &Tdmq.InstanceArgs{
//				ClusterName: pulumi.String("tf_example"),
//				Remark:      pulumi.String("remark."),
//				Tags: pulumi.Map{
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
//			exampleTopic, err := Tdmq.NewTopic(ctx, "exampleTopic", &Tdmq.TopicArgs{
//				ClusterId:       exampleInstance.ID(),
//				EnvironId:       exampleNamespace.EnvironName,
//				TopicName:       pulumi.String("tf-example-topic"),
//				Partitions:      pulumi.Int(1),
//				PulsarTopicType: pulumi.Int(3),
//				Remark:          pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tdmq.NewSubscription(ctx, "exampleSubscription", &Tdmq.SubscriptionArgs{
//				ClusterId:             exampleInstance.ID(),
//				EnvironmentId:         exampleNamespace.EnvironName,
//				TopicName:             exampleTopic.TopicName,
//				SubscriptionName:      pulumi.String("tf-example-subscription"),
//				Remark:                pulumi.String("remark."),
//				AutoCreatePolicyTopic: pulumi.Bool(true),
//				AutoDeletePolicyTopic: pulumi.Bool(true),
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
// tdmq subscription can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tdmq/subscription:Subscription example pulsar-q4k5898krpqj#tf_example#tf-example-topic#tf-example-subscription#true
// ```
type Subscription struct {
	pulumi.CustomResourceState

	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
	AutoCreatePolicyTopic pulumi.BoolPtrOutput `pulumi:"autoCreatePolicyTopic"`
	// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
	AutoDeletePolicyTopic pulumi.BoolPtrOutput `pulumi:"autoDeletePolicyTopic"`
	// Pulsar cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName pulumi.StringOutput `pulumi:"subscriptionName"`
	// Topic name.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Subscription
	err := ctx.RegisterResource("tencentcloud:Tdmq/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("tencentcloud:Tdmq/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
	AutoCreatePolicyTopic *bool `pulumi:"autoCreatePolicyTopic"`
	// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
	AutoDeletePolicyTopic *bool `pulumi:"autoDeletePolicyTopic"`
	// Pulsar cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentId *string `pulumi:"environmentId"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName *string `pulumi:"subscriptionName"`
	// Topic name.
	TopicName *string `pulumi:"topicName"`
}

type SubscriptionState struct {
	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
	AutoCreatePolicyTopic pulumi.BoolPtrInput
	// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
	AutoDeletePolicyTopic pulumi.BoolPtrInput
	// Pulsar cluster ID.
	ClusterId pulumi.StringPtrInput
	// Environment (namespace) name.
	EnvironmentId pulumi.StringPtrInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName pulumi.StringPtrInput
	// Topic name.
	TopicName pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
	AutoCreatePolicyTopic *bool `pulumi:"autoCreatePolicyTopic"`
	// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
	AutoDeletePolicyTopic *bool `pulumi:"autoDeletePolicyTopic"`
	// Pulsar cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentId string `pulumi:"environmentId"`
	// Remarks (up to 128 characters).
	Remark *string `pulumi:"remark"`
	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName string `pulumi:"subscriptionName"`
	// Topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
	AutoCreatePolicyTopic pulumi.BoolPtrInput
	// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
	AutoDeletePolicyTopic pulumi.BoolPtrInput
	// Pulsar cluster ID.
	ClusterId pulumi.StringInput
	// Environment (namespace) name.
	EnvironmentId pulumi.StringInput
	// Remarks (up to 128 characters).
	Remark pulumi.StringPtrInput
	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName pulumi.StringInput
	// Topic name.
	TopicName pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

// SubscriptionArrayInput is an input type that accepts SubscriptionArray and SubscriptionArrayOutput values.
// You can construct a concrete instance of `SubscriptionArrayInput` via:
//
//	SubscriptionArray{ SubscriptionArgs{...} }
type SubscriptionArrayInput interface {
	pulumi.Input

	ToSubscriptionArrayOutput() SubscriptionArrayOutput
	ToSubscriptionArrayOutputWithContext(context.Context) SubscriptionArrayOutput
}

type SubscriptionArray []SubscriptionInput

func (SubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscription)(nil)).Elem()
}

func (i SubscriptionArray) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return i.ToSubscriptionArrayOutputWithContext(context.Background())
}

func (i SubscriptionArray) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionArrayOutput)
}

// SubscriptionMapInput is an input type that accepts SubscriptionMap and SubscriptionMapOutput values.
// You can construct a concrete instance of `SubscriptionMapInput` via:
//
//	SubscriptionMap{ "key": SubscriptionArgs{...} }
type SubscriptionMapInput interface {
	pulumi.Input

	ToSubscriptionMapOutput() SubscriptionMapOutput
	ToSubscriptionMapOutputWithContext(context.Context) SubscriptionMapOutput
}

type SubscriptionMap map[string]SubscriptionInput

func (SubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscription)(nil)).Elem()
}

func (i SubscriptionMap) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return i.ToSubscriptionMapOutputWithContext(context.Background())
}

func (i SubscriptionMap) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionMapOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

// Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
func (o SubscriptionOutput) AutoCreatePolicyTopic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.AutoCreatePolicyTopic }).(pulumi.BoolPtrOutput)
}

// Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
func (o SubscriptionOutput) AutoDeletePolicyTopic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.AutoDeletePolicyTopic }).(pulumi.BoolPtrOutput)
}

// Pulsar cluster ID.
func (o SubscriptionOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Environment (namespace) name.
func (o SubscriptionOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// Remarks (up to 128 characters).
func (o SubscriptionOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Subscriber name, which can contain up to 128 characters.
func (o SubscriptionOutput) SubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.SubscriptionName }).(pulumi.StringOutput)
}

// Topic name.
func (o SubscriptionOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type SubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscription)(nil)).Elem()
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) Index(i pulumi.IntInput) SubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Subscription {
		return vs[0].([]*Subscription)[vs[1].(int)]
	}).(SubscriptionOutput)
}

type SubscriptionMapOutput struct{ *pulumi.OutputState }

func (SubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscription)(nil)).Elem()
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) MapIndex(k pulumi.StringInput) SubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Subscription {
		return vs[0].(map[string]*Subscription)[vs[1].(string)]
	}).(SubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionInput)(nil)).Elem(), &Subscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionArrayInput)(nil)).Elem(), SubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionMapInput)(nil)).Elem(), SubscriptionMap{})
	pulumi.RegisterOutputType(SubscriptionOutput{})
	pulumi.RegisterOutputType(SubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionMapOutput{})
}
