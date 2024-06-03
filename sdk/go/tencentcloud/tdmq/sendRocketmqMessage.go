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

// Provides a resource to create a tdmq sendRocketmqMessage
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
//			_, err := Tdmq.NewSendRocketmqMessage(ctx, "sendRocketmqMessage", &Tdmq.SendRocketmqMessageArgs{
//				ClusterId:   pulumi.String("rocketmq-7k45z9dkpnne"),
//				MsgBody:     pulumi.String("msg key"),
//				MsgKey:      pulumi.String("msg tag"),
//				MsgTag:      pulumi.String("msg value"),
//				NamespaceId: pulumi.String("test_ns"),
//				TopicName:   pulumi.String("test_topic"),
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
type SendRocketmqMessage struct {
	pulumi.CustomResourceState

	// Cluster id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Information.
	MsgBody pulumi.StringOutput `pulumi:"msgBody"`
	// Message key information.
	MsgKey pulumi.StringPtrOutput `pulumi:"msgKey"`
	// Message tag information.
	MsgTag pulumi.StringPtrOutput `pulumi:"msgTag"`
	// Namespaces.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// topic name.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewSendRocketmqMessage registers a new resource with the given unique name, arguments, and options.
func NewSendRocketmqMessage(ctx *pulumi.Context,
	name string, args *SendRocketmqMessageArgs, opts ...pulumi.ResourceOption) (*SendRocketmqMessage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.MsgBody == nil {
		return nil, errors.New("invalid value for required argument 'MsgBody'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SendRocketmqMessage
	err := ctx.RegisterResource("tencentcloud:Tdmq/sendRocketmqMessage:SendRocketmqMessage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSendRocketmqMessage gets an existing SendRocketmqMessage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSendRocketmqMessage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SendRocketmqMessageState, opts ...pulumi.ResourceOption) (*SendRocketmqMessage, error) {
	var resource SendRocketmqMessage
	err := ctx.ReadResource("tencentcloud:Tdmq/sendRocketmqMessage:SendRocketmqMessage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SendRocketmqMessage resources.
type sendRocketmqMessageState struct {
	// Cluster id.
	ClusterId *string `pulumi:"clusterId"`
	// Information.
	MsgBody *string `pulumi:"msgBody"`
	// Message key information.
	MsgKey *string `pulumi:"msgKey"`
	// Message tag information.
	MsgTag *string `pulumi:"msgTag"`
	// Namespaces.
	NamespaceId *string `pulumi:"namespaceId"`
	// topic name.
	TopicName *string `pulumi:"topicName"`
}

type SendRocketmqMessageState struct {
	// Cluster id.
	ClusterId pulumi.StringPtrInput
	// Information.
	MsgBody pulumi.StringPtrInput
	// Message key information.
	MsgKey pulumi.StringPtrInput
	// Message tag information.
	MsgTag pulumi.StringPtrInput
	// Namespaces.
	NamespaceId pulumi.StringPtrInput
	// topic name.
	TopicName pulumi.StringPtrInput
}

func (SendRocketmqMessageState) ElementType() reflect.Type {
	return reflect.TypeOf((*sendRocketmqMessageState)(nil)).Elem()
}

type sendRocketmqMessageArgs struct {
	// Cluster id.
	ClusterId string `pulumi:"clusterId"`
	// Information.
	MsgBody string `pulumi:"msgBody"`
	// Message key information.
	MsgKey *string `pulumi:"msgKey"`
	// Message tag information.
	MsgTag *string `pulumi:"msgTag"`
	// Namespaces.
	NamespaceId string `pulumi:"namespaceId"`
	// topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a SendRocketmqMessage resource.
type SendRocketmqMessageArgs struct {
	// Cluster id.
	ClusterId pulumi.StringInput
	// Information.
	MsgBody pulumi.StringInput
	// Message key information.
	MsgKey pulumi.StringPtrInput
	// Message tag information.
	MsgTag pulumi.StringPtrInput
	// Namespaces.
	NamespaceId pulumi.StringInput
	// topic name.
	TopicName pulumi.StringInput
}

func (SendRocketmqMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sendRocketmqMessageArgs)(nil)).Elem()
}

type SendRocketmqMessageInput interface {
	pulumi.Input

	ToSendRocketmqMessageOutput() SendRocketmqMessageOutput
	ToSendRocketmqMessageOutputWithContext(ctx context.Context) SendRocketmqMessageOutput
}

func (*SendRocketmqMessage) ElementType() reflect.Type {
	return reflect.TypeOf((**SendRocketmqMessage)(nil)).Elem()
}

func (i *SendRocketmqMessage) ToSendRocketmqMessageOutput() SendRocketmqMessageOutput {
	return i.ToSendRocketmqMessageOutputWithContext(context.Background())
}

func (i *SendRocketmqMessage) ToSendRocketmqMessageOutputWithContext(ctx context.Context) SendRocketmqMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SendRocketmqMessageOutput)
}

// SendRocketmqMessageArrayInput is an input type that accepts SendRocketmqMessageArray and SendRocketmqMessageArrayOutput values.
// You can construct a concrete instance of `SendRocketmqMessageArrayInput` via:
//
//	SendRocketmqMessageArray{ SendRocketmqMessageArgs{...} }
type SendRocketmqMessageArrayInput interface {
	pulumi.Input

	ToSendRocketmqMessageArrayOutput() SendRocketmqMessageArrayOutput
	ToSendRocketmqMessageArrayOutputWithContext(context.Context) SendRocketmqMessageArrayOutput
}

type SendRocketmqMessageArray []SendRocketmqMessageInput

func (SendRocketmqMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SendRocketmqMessage)(nil)).Elem()
}

func (i SendRocketmqMessageArray) ToSendRocketmqMessageArrayOutput() SendRocketmqMessageArrayOutput {
	return i.ToSendRocketmqMessageArrayOutputWithContext(context.Background())
}

func (i SendRocketmqMessageArray) ToSendRocketmqMessageArrayOutputWithContext(ctx context.Context) SendRocketmqMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SendRocketmqMessageArrayOutput)
}

// SendRocketmqMessageMapInput is an input type that accepts SendRocketmqMessageMap and SendRocketmqMessageMapOutput values.
// You can construct a concrete instance of `SendRocketmqMessageMapInput` via:
//
//	SendRocketmqMessageMap{ "key": SendRocketmqMessageArgs{...} }
type SendRocketmqMessageMapInput interface {
	pulumi.Input

	ToSendRocketmqMessageMapOutput() SendRocketmqMessageMapOutput
	ToSendRocketmqMessageMapOutputWithContext(context.Context) SendRocketmqMessageMapOutput
}

type SendRocketmqMessageMap map[string]SendRocketmqMessageInput

func (SendRocketmqMessageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SendRocketmqMessage)(nil)).Elem()
}

func (i SendRocketmqMessageMap) ToSendRocketmqMessageMapOutput() SendRocketmqMessageMapOutput {
	return i.ToSendRocketmqMessageMapOutputWithContext(context.Background())
}

func (i SendRocketmqMessageMap) ToSendRocketmqMessageMapOutputWithContext(ctx context.Context) SendRocketmqMessageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SendRocketmqMessageMapOutput)
}

type SendRocketmqMessageOutput struct{ *pulumi.OutputState }

func (SendRocketmqMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SendRocketmqMessage)(nil)).Elem()
}

func (o SendRocketmqMessageOutput) ToSendRocketmqMessageOutput() SendRocketmqMessageOutput {
	return o
}

func (o SendRocketmqMessageOutput) ToSendRocketmqMessageOutputWithContext(ctx context.Context) SendRocketmqMessageOutput {
	return o
}

// Cluster id.
func (o SendRocketmqMessageOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Information.
func (o SendRocketmqMessageOutput) MsgBody() pulumi.StringOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringOutput { return v.MsgBody }).(pulumi.StringOutput)
}

// Message key information.
func (o SendRocketmqMessageOutput) MsgKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringPtrOutput { return v.MsgKey }).(pulumi.StringPtrOutput)
}

// Message tag information.
func (o SendRocketmqMessageOutput) MsgTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringPtrOutput { return v.MsgTag }).(pulumi.StringPtrOutput)
}

// Namespaces.
func (o SendRocketmqMessageOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// topic name.
func (o SendRocketmqMessageOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *SendRocketmqMessage) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type SendRocketmqMessageArrayOutput struct{ *pulumi.OutputState }

func (SendRocketmqMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SendRocketmqMessage)(nil)).Elem()
}

func (o SendRocketmqMessageArrayOutput) ToSendRocketmqMessageArrayOutput() SendRocketmqMessageArrayOutput {
	return o
}

func (o SendRocketmqMessageArrayOutput) ToSendRocketmqMessageArrayOutputWithContext(ctx context.Context) SendRocketmqMessageArrayOutput {
	return o
}

func (o SendRocketmqMessageArrayOutput) Index(i pulumi.IntInput) SendRocketmqMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SendRocketmqMessage {
		return vs[0].([]*SendRocketmqMessage)[vs[1].(int)]
	}).(SendRocketmqMessageOutput)
}

type SendRocketmqMessageMapOutput struct{ *pulumi.OutputState }

func (SendRocketmqMessageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SendRocketmqMessage)(nil)).Elem()
}

func (o SendRocketmqMessageMapOutput) ToSendRocketmqMessageMapOutput() SendRocketmqMessageMapOutput {
	return o
}

func (o SendRocketmqMessageMapOutput) ToSendRocketmqMessageMapOutputWithContext(ctx context.Context) SendRocketmqMessageMapOutput {
	return o
}

func (o SendRocketmqMessageMapOutput) MapIndex(k pulumi.StringInput) SendRocketmqMessageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SendRocketmqMessage {
		return vs[0].(map[string]*SendRocketmqMessage)[vs[1].(string)]
	}).(SendRocketmqMessageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SendRocketmqMessageInput)(nil)).Elem(), &SendRocketmqMessage{})
	pulumi.RegisterInputType(reflect.TypeOf((*SendRocketmqMessageArrayInput)(nil)).Elem(), SendRocketmqMessageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SendRocketmqMessageMapInput)(nil)).Elem(), SendRocketmqMessageMap{})
	pulumi.RegisterOutputType(SendRocketmqMessageOutput{})
	pulumi.RegisterOutputType(SendRocketmqMessageArrayOutput{})
	pulumi.RegisterOutputType(SendRocketmqMessageMapOutput{})
}
