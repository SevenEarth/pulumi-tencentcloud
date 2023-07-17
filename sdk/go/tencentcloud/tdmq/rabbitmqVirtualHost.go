// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tdmq rabbitmqVirtualHost
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tdmq.NewRabbitmqVirtualHost(ctx, "rabbitmqVirtualHost", &Tdmq.RabbitmqVirtualHostArgs{
// 			Description: pulumi.String("desc"),
// 			InstanceId:  pulumi.String("amqp-kzbe8p3n"),
// 			TraceFlag:   pulumi.Bool(false),
// 			VirtualHost: pulumi.String("vh-test-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RabbitmqVirtualHost struct {
	pulumi.CustomResourceState

	// describe.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Cluster instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Message track switch, true is on, false is off, default is off.
	TraceFlag pulumi.BoolPtrOutput `pulumi:"traceFlag"`
	// vhost name.
	VirtualHost pulumi.StringOutput `pulumi:"virtualHost"`
}

// NewRabbitmqVirtualHost registers a new resource with the given unique name, arguments, and options.
func NewRabbitmqVirtualHost(ctx *pulumi.Context,
	name string, args *RabbitmqVirtualHostArgs, opts ...pulumi.ResourceOption) (*RabbitmqVirtualHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VirtualHost == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHost'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RabbitmqVirtualHost
	err := ctx.RegisterResource("tencentcloud:Tdmq/rabbitmqVirtualHost:RabbitmqVirtualHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRabbitmqVirtualHost gets an existing RabbitmqVirtualHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRabbitmqVirtualHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RabbitmqVirtualHostState, opts ...pulumi.ResourceOption) (*RabbitmqVirtualHost, error) {
	var resource RabbitmqVirtualHost
	err := ctx.ReadResource("tencentcloud:Tdmq/rabbitmqVirtualHost:RabbitmqVirtualHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RabbitmqVirtualHost resources.
type rabbitmqVirtualHostState struct {
	// describe.
	Description *string `pulumi:"description"`
	// Cluster instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Message track switch, true is on, false is off, default is off.
	TraceFlag *bool `pulumi:"traceFlag"`
	// vhost name.
	VirtualHost *string `pulumi:"virtualHost"`
}

type RabbitmqVirtualHostState struct {
	// describe.
	Description pulumi.StringPtrInput
	// Cluster instance ID.
	InstanceId pulumi.StringPtrInput
	// Message track switch, true is on, false is off, default is off.
	TraceFlag pulumi.BoolPtrInput
	// vhost name.
	VirtualHost pulumi.StringPtrInput
}

func (RabbitmqVirtualHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitmqVirtualHostState)(nil)).Elem()
}

type rabbitmqVirtualHostArgs struct {
	// describe.
	Description *string `pulumi:"description"`
	// Cluster instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Message track switch, true is on, false is off, default is off.
	TraceFlag *bool `pulumi:"traceFlag"`
	// vhost name.
	VirtualHost string `pulumi:"virtualHost"`
}

// The set of arguments for constructing a RabbitmqVirtualHost resource.
type RabbitmqVirtualHostArgs struct {
	// describe.
	Description pulumi.StringPtrInput
	// Cluster instance ID.
	InstanceId pulumi.StringInput
	// Message track switch, true is on, false is off, default is off.
	TraceFlag pulumi.BoolPtrInput
	// vhost name.
	VirtualHost pulumi.StringInput
}

func (RabbitmqVirtualHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rabbitmqVirtualHostArgs)(nil)).Elem()
}

type RabbitmqVirtualHostInput interface {
	pulumi.Input

	ToRabbitmqVirtualHostOutput() RabbitmqVirtualHostOutput
	ToRabbitmqVirtualHostOutputWithContext(ctx context.Context) RabbitmqVirtualHostOutput
}

func (*RabbitmqVirtualHost) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitmqVirtualHost)(nil)).Elem()
}

func (i *RabbitmqVirtualHost) ToRabbitmqVirtualHostOutput() RabbitmqVirtualHostOutput {
	return i.ToRabbitmqVirtualHostOutputWithContext(context.Background())
}

func (i *RabbitmqVirtualHost) ToRabbitmqVirtualHostOutputWithContext(ctx context.Context) RabbitmqVirtualHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVirtualHostOutput)
}

// RabbitmqVirtualHostArrayInput is an input type that accepts RabbitmqVirtualHostArray and RabbitmqVirtualHostArrayOutput values.
// You can construct a concrete instance of `RabbitmqVirtualHostArrayInput` via:
//
//          RabbitmqVirtualHostArray{ RabbitmqVirtualHostArgs{...} }
type RabbitmqVirtualHostArrayInput interface {
	pulumi.Input

	ToRabbitmqVirtualHostArrayOutput() RabbitmqVirtualHostArrayOutput
	ToRabbitmqVirtualHostArrayOutputWithContext(context.Context) RabbitmqVirtualHostArrayOutput
}

type RabbitmqVirtualHostArray []RabbitmqVirtualHostInput

func (RabbitmqVirtualHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitmqVirtualHost)(nil)).Elem()
}

func (i RabbitmqVirtualHostArray) ToRabbitmqVirtualHostArrayOutput() RabbitmqVirtualHostArrayOutput {
	return i.ToRabbitmqVirtualHostArrayOutputWithContext(context.Background())
}

func (i RabbitmqVirtualHostArray) ToRabbitmqVirtualHostArrayOutputWithContext(ctx context.Context) RabbitmqVirtualHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVirtualHostArrayOutput)
}

// RabbitmqVirtualHostMapInput is an input type that accepts RabbitmqVirtualHostMap and RabbitmqVirtualHostMapOutput values.
// You can construct a concrete instance of `RabbitmqVirtualHostMapInput` via:
//
//          RabbitmqVirtualHostMap{ "key": RabbitmqVirtualHostArgs{...} }
type RabbitmqVirtualHostMapInput interface {
	pulumi.Input

	ToRabbitmqVirtualHostMapOutput() RabbitmqVirtualHostMapOutput
	ToRabbitmqVirtualHostMapOutputWithContext(context.Context) RabbitmqVirtualHostMapOutput
}

type RabbitmqVirtualHostMap map[string]RabbitmqVirtualHostInput

func (RabbitmqVirtualHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitmqVirtualHost)(nil)).Elem()
}

func (i RabbitmqVirtualHostMap) ToRabbitmqVirtualHostMapOutput() RabbitmqVirtualHostMapOutput {
	return i.ToRabbitmqVirtualHostMapOutputWithContext(context.Background())
}

func (i RabbitmqVirtualHostMap) ToRabbitmqVirtualHostMapOutputWithContext(ctx context.Context) RabbitmqVirtualHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RabbitmqVirtualHostMapOutput)
}

type RabbitmqVirtualHostOutput struct{ *pulumi.OutputState }

func (RabbitmqVirtualHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RabbitmqVirtualHost)(nil)).Elem()
}

func (o RabbitmqVirtualHostOutput) ToRabbitmqVirtualHostOutput() RabbitmqVirtualHostOutput {
	return o
}

func (o RabbitmqVirtualHostOutput) ToRabbitmqVirtualHostOutputWithContext(ctx context.Context) RabbitmqVirtualHostOutput {
	return o
}

// describe.
func (o RabbitmqVirtualHostOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RabbitmqVirtualHost) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Cluster instance ID.
func (o RabbitmqVirtualHostOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RabbitmqVirtualHost) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Message track switch, true is on, false is off, default is off.
func (o RabbitmqVirtualHostOutput) TraceFlag() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RabbitmqVirtualHost) pulumi.BoolPtrOutput { return v.TraceFlag }).(pulumi.BoolPtrOutput)
}

// vhost name.
func (o RabbitmqVirtualHostOutput) VirtualHost() pulumi.StringOutput {
	return o.ApplyT(func(v *RabbitmqVirtualHost) pulumi.StringOutput { return v.VirtualHost }).(pulumi.StringOutput)
}

type RabbitmqVirtualHostArrayOutput struct{ *pulumi.OutputState }

func (RabbitmqVirtualHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RabbitmqVirtualHost)(nil)).Elem()
}

func (o RabbitmqVirtualHostArrayOutput) ToRabbitmqVirtualHostArrayOutput() RabbitmqVirtualHostArrayOutput {
	return o
}

func (o RabbitmqVirtualHostArrayOutput) ToRabbitmqVirtualHostArrayOutputWithContext(ctx context.Context) RabbitmqVirtualHostArrayOutput {
	return o
}

func (o RabbitmqVirtualHostArrayOutput) Index(i pulumi.IntInput) RabbitmqVirtualHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RabbitmqVirtualHost {
		return vs[0].([]*RabbitmqVirtualHost)[vs[1].(int)]
	}).(RabbitmqVirtualHostOutput)
}

type RabbitmqVirtualHostMapOutput struct{ *pulumi.OutputState }

func (RabbitmqVirtualHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RabbitmqVirtualHost)(nil)).Elem()
}

func (o RabbitmqVirtualHostMapOutput) ToRabbitmqVirtualHostMapOutput() RabbitmqVirtualHostMapOutput {
	return o
}

func (o RabbitmqVirtualHostMapOutput) ToRabbitmqVirtualHostMapOutputWithContext(ctx context.Context) RabbitmqVirtualHostMapOutput {
	return o
}

func (o RabbitmqVirtualHostMapOutput) MapIndex(k pulumi.StringInput) RabbitmqVirtualHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RabbitmqVirtualHost {
		return vs[0].(map[string]*RabbitmqVirtualHost)[vs[1].(string)]
	}).(RabbitmqVirtualHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVirtualHostInput)(nil)).Elem(), &RabbitmqVirtualHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVirtualHostArrayInput)(nil)).Elem(), RabbitmqVirtualHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RabbitmqVirtualHostMapInput)(nil)).Elem(), RabbitmqVirtualHostMap{})
	pulumi.RegisterOutputType(RabbitmqVirtualHostOutput{})
	pulumi.RegisterOutputType(RabbitmqVirtualHostArrayOutput{})
	pulumi.RegisterOutputType(RabbitmqVirtualHostMapOutput{})
}
