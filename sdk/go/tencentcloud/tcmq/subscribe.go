// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcmq subscribe
//
// ## Import
//
// tcmq subscribe can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tcmq/subscribe:Subscribe subscribe subscribe_id
// ```
type Subscribe struct {
	pulumi.CustomResourceState

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKeys pulumi.StringArrayOutput `pulumi:"bindingKeys"`
	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags pulumi.StringArrayOutput `pulumi:"filterTags"`
	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat pulumi.StringPtrOutput `pulumi:"notifyContentFormat"`
	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
	NotifyStrategy pulumi.StringPtrOutput `pulumi:"notifyStrategy"`
	// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName pulumi.StringOutput `pulumi:"subscriptionName"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewSubscribe registers a new resource with the given unique name, arguments, and options.
func NewSubscribe(ctx *pulumi.Context,
	name string, args *SubscribeArgs, opts ...pulumi.ResourceOption) (*Subscribe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Subscribe
	err := ctx.RegisterResource("tencentcloud:Tcmq/subscribe:Subscribe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscribe gets an existing Subscribe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscribe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscribeState, opts ...pulumi.ResourceOption) (*Subscribe, error) {
	var resource Subscribe
	err := ctx.ReadResource("tencentcloud:Tcmq/subscribe:Subscribe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscribe resources.
type subscribeState struct {
	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKeys []string `pulumi:"bindingKeys"`
	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint *string `pulumi:"endpoint"`
	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []string `pulumi:"filterTags"`
	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `pulumi:"notifyContentFormat"`
	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
	NotifyStrategy *string `pulumi:"notifyStrategy"`
	// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol *string `pulumi:"protocol"`
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `pulumi:"subscriptionName"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `pulumi:"topicName"`
}

type SubscribeState struct {
	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKeys pulumi.StringArrayInput
	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint pulumi.StringPtrInput
	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags pulumi.StringArrayInput
	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat pulumi.StringPtrInput
	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
	NotifyStrategy pulumi.StringPtrInput
	// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol pulumi.StringPtrInput
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName pulumi.StringPtrInput
}

func (SubscribeState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribeState)(nil)).Elem()
}

type subscribeArgs struct {
	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKeys []string `pulumi:"bindingKeys"`
	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint string `pulumi:"endpoint"`
	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []string `pulumi:"filterTags"`
	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `pulumi:"notifyContentFormat"`
	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
	NotifyStrategy *string `pulumi:"notifyStrategy"`
	// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol string `pulumi:"protocol"`
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName string `pulumi:"subscriptionName"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Subscribe resource.
type SubscribeArgs struct {
	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKeys pulumi.StringArrayInput
	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint pulumi.StringInput
	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags pulumi.StringArrayInput
	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat pulumi.StringPtrInput
	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
	NotifyStrategy pulumi.StringPtrInput
	// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol pulumi.StringInput
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName pulumi.StringInput
}

func (SubscribeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribeArgs)(nil)).Elem()
}

type SubscribeInput interface {
	pulumi.Input

	ToSubscribeOutput() SubscribeOutput
	ToSubscribeOutputWithContext(ctx context.Context) SubscribeOutput
}

func (*Subscribe) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscribe)(nil)).Elem()
}

func (i *Subscribe) ToSubscribeOutput() SubscribeOutput {
	return i.ToSubscribeOutputWithContext(context.Background())
}

func (i *Subscribe) ToSubscribeOutputWithContext(ctx context.Context) SubscribeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribeOutput)
}

// SubscribeArrayInput is an input type that accepts SubscribeArray and SubscribeArrayOutput values.
// You can construct a concrete instance of `SubscribeArrayInput` via:
//
//          SubscribeArray{ SubscribeArgs{...} }
type SubscribeArrayInput interface {
	pulumi.Input

	ToSubscribeArrayOutput() SubscribeArrayOutput
	ToSubscribeArrayOutputWithContext(context.Context) SubscribeArrayOutput
}

type SubscribeArray []SubscribeInput

func (SubscribeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscribe)(nil)).Elem()
}

func (i SubscribeArray) ToSubscribeArrayOutput() SubscribeArrayOutput {
	return i.ToSubscribeArrayOutputWithContext(context.Background())
}

func (i SubscribeArray) ToSubscribeArrayOutputWithContext(ctx context.Context) SubscribeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribeArrayOutput)
}

// SubscribeMapInput is an input type that accepts SubscribeMap and SubscribeMapOutput values.
// You can construct a concrete instance of `SubscribeMapInput` via:
//
//          SubscribeMap{ "key": SubscribeArgs{...} }
type SubscribeMapInput interface {
	pulumi.Input

	ToSubscribeMapOutput() SubscribeMapOutput
	ToSubscribeMapOutputWithContext(context.Context) SubscribeMapOutput
}

type SubscribeMap map[string]SubscribeInput

func (SubscribeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscribe)(nil)).Elem()
}

func (i SubscribeMap) ToSubscribeMapOutput() SubscribeMapOutput {
	return i.ToSubscribeMapOutputWithContext(context.Background())
}

func (i SubscribeMap) ToSubscribeMapOutputWithContext(ctx context.Context) SubscribeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribeMapOutput)
}

type SubscribeOutput struct{ *pulumi.OutputState }

func (SubscribeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscribe)(nil)).Elem()
}

func (o SubscribeOutput) ToSubscribeOutput() SubscribeOutput {
	return o
}

func (o SubscribeOutput) ToSubscribeOutputWithContext(ctx context.Context) SubscribeOutput {
	return o
}

// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
func (o SubscribeOutput) BindingKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringArrayOutput { return v.BindingKeys }).(pulumi.StringArrayOutput)
}

// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
func (o SubscribeOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
func (o SubscribeOutput) FilterTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringArrayOutput { return v.FilterTags }).(pulumi.StringArrayOutput)
}

// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
func (o SubscribeOutput) NotifyContentFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringPtrOutput { return v.NotifyContentFormat }).(pulumi.StringPtrOutput)
}

// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. `BACKOFF_RETRY`: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. `EXPONENTIAL_DECAY_RETRY`: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: `EXPONENTIAL_DECAY_RETRY`.
func (o SubscribeOutput) NotifyStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringPtrOutput { return v.NotifyStrategy }).(pulumi.StringPtrOutput)
}

// ubscription protocol. Currently, two protocols are supported: `http` and `queue`. To use the `http` protocol, you need to build your own web server to receive messages. With the `queue` protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
func (o SubscribeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
func (o SubscribeOutput) SubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringOutput { return v.SubscriptionName }).(pulumi.StringOutput)
}

// Tag description list.
func (o SubscribeOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
func (o SubscribeOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscribe) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type SubscribeArrayOutput struct{ *pulumi.OutputState }

func (SubscribeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscribe)(nil)).Elem()
}

func (o SubscribeArrayOutput) ToSubscribeArrayOutput() SubscribeArrayOutput {
	return o
}

func (o SubscribeArrayOutput) ToSubscribeArrayOutputWithContext(ctx context.Context) SubscribeArrayOutput {
	return o
}

func (o SubscribeArrayOutput) Index(i pulumi.IntInput) SubscribeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Subscribe {
		return vs[0].([]*Subscribe)[vs[1].(int)]
	}).(SubscribeOutput)
}

type SubscribeMapOutput struct{ *pulumi.OutputState }

func (SubscribeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscribe)(nil)).Elem()
}

func (o SubscribeMapOutput) ToSubscribeMapOutput() SubscribeMapOutput {
	return o
}

func (o SubscribeMapOutput) ToSubscribeMapOutputWithContext(ctx context.Context) SubscribeMapOutput {
	return o
}

func (o SubscribeMapOutput) MapIndex(k pulumi.StringInput) SubscribeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Subscribe {
		return vs[0].(map[string]*Subscribe)[vs[1].(string)]
	}).(SubscribeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribeInput)(nil)).Elem(), &Subscribe{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribeArrayInput)(nil)).Elem(), SubscribeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribeMapInput)(nil)).Elem(), SubscribeMap{})
	pulumi.RegisterOutputType(SubscribeOutput{})
	pulumi.RegisterOutputType(SubscribeArrayOutput{})
	pulumi.RegisterOutputType(SubscribeMapOutput{})
}
