// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor tmpCvmAgent
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.NewTmpCvmAgent(ctx, "tmpCvmAgent", &Monitor.TmpCvmAgentArgs{
//				InstanceId: pulumi.String("prom-dko9d0nu"),
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
// monitor tmpCvmAgent can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent tmpCvmAgent tmpCvmAgent_id
//
// ```
type TmpCvmAgent struct {
	pulumi.CustomResourceState

	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Agent name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTmpCvmAgent registers a new resource with the given unique name, arguments, and options.
func NewTmpCvmAgent(ctx *pulumi.Context,
	name string, args *TmpCvmAgentArgs, opts ...pulumi.ResourceOption) (*TmpCvmAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource TmpCvmAgent
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpCvmAgent gets an existing TmpCvmAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpCvmAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpCvmAgentState, opts ...pulumi.ResourceOption) (*TmpCvmAgent, error) {
	var resource TmpCvmAgent
	err := ctx.ReadResource("tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpCvmAgent resources.
type tmpCvmAgentState struct {
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Agent name.
	Name *string `pulumi:"name"`
}

type TmpCvmAgentState struct {
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Agent name.
	Name pulumi.StringPtrInput
}

func (TmpCvmAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpCvmAgentState)(nil)).Elem()
}

type tmpCvmAgentArgs struct {
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Agent name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TmpCvmAgent resource.
type TmpCvmAgentArgs struct {
	// Instance id.
	InstanceId pulumi.StringInput
	// Agent name.
	Name pulumi.StringPtrInput
}

func (TmpCvmAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpCvmAgentArgs)(nil)).Elem()
}

type TmpCvmAgentInput interface {
	pulumi.Input

	ToTmpCvmAgentOutput() TmpCvmAgentOutput
	ToTmpCvmAgentOutputWithContext(ctx context.Context) TmpCvmAgentOutput
}

func (*TmpCvmAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpCvmAgent)(nil)).Elem()
}

func (i *TmpCvmAgent) ToTmpCvmAgentOutput() TmpCvmAgentOutput {
	return i.ToTmpCvmAgentOutputWithContext(context.Background())
}

func (i *TmpCvmAgent) ToTmpCvmAgentOutputWithContext(ctx context.Context) TmpCvmAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpCvmAgentOutput)
}

// TmpCvmAgentArrayInput is an input type that accepts TmpCvmAgentArray and TmpCvmAgentArrayOutput values.
// You can construct a concrete instance of `TmpCvmAgentArrayInput` via:
//
//	TmpCvmAgentArray{ TmpCvmAgentArgs{...} }
type TmpCvmAgentArrayInput interface {
	pulumi.Input

	ToTmpCvmAgentArrayOutput() TmpCvmAgentArrayOutput
	ToTmpCvmAgentArrayOutputWithContext(context.Context) TmpCvmAgentArrayOutput
}

type TmpCvmAgentArray []TmpCvmAgentInput

func (TmpCvmAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpCvmAgent)(nil)).Elem()
}

func (i TmpCvmAgentArray) ToTmpCvmAgentArrayOutput() TmpCvmAgentArrayOutput {
	return i.ToTmpCvmAgentArrayOutputWithContext(context.Background())
}

func (i TmpCvmAgentArray) ToTmpCvmAgentArrayOutputWithContext(ctx context.Context) TmpCvmAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpCvmAgentArrayOutput)
}

// TmpCvmAgentMapInput is an input type that accepts TmpCvmAgentMap and TmpCvmAgentMapOutput values.
// You can construct a concrete instance of `TmpCvmAgentMapInput` via:
//
//	TmpCvmAgentMap{ "key": TmpCvmAgentArgs{...} }
type TmpCvmAgentMapInput interface {
	pulumi.Input

	ToTmpCvmAgentMapOutput() TmpCvmAgentMapOutput
	ToTmpCvmAgentMapOutputWithContext(context.Context) TmpCvmAgentMapOutput
}

type TmpCvmAgentMap map[string]TmpCvmAgentInput

func (TmpCvmAgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpCvmAgent)(nil)).Elem()
}

func (i TmpCvmAgentMap) ToTmpCvmAgentMapOutput() TmpCvmAgentMapOutput {
	return i.ToTmpCvmAgentMapOutputWithContext(context.Background())
}

func (i TmpCvmAgentMap) ToTmpCvmAgentMapOutputWithContext(ctx context.Context) TmpCvmAgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpCvmAgentMapOutput)
}

type TmpCvmAgentOutput struct{ *pulumi.OutputState }

func (TmpCvmAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpCvmAgent)(nil)).Elem()
}

func (o TmpCvmAgentOutput) ToTmpCvmAgentOutput() TmpCvmAgentOutput {
	return o
}

func (o TmpCvmAgentOutput) ToTmpCvmAgentOutputWithContext(ctx context.Context) TmpCvmAgentOutput {
	return o
}

// Instance id.
func (o TmpCvmAgentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpCvmAgent) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Agent name.
func (o TmpCvmAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpCvmAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TmpCvmAgentArrayOutput struct{ *pulumi.OutputState }

func (TmpCvmAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpCvmAgent)(nil)).Elem()
}

func (o TmpCvmAgentArrayOutput) ToTmpCvmAgentArrayOutput() TmpCvmAgentArrayOutput {
	return o
}

func (o TmpCvmAgentArrayOutput) ToTmpCvmAgentArrayOutputWithContext(ctx context.Context) TmpCvmAgentArrayOutput {
	return o
}

func (o TmpCvmAgentArrayOutput) Index(i pulumi.IntInput) TmpCvmAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpCvmAgent {
		return vs[0].([]*TmpCvmAgent)[vs[1].(int)]
	}).(TmpCvmAgentOutput)
}

type TmpCvmAgentMapOutput struct{ *pulumi.OutputState }

func (TmpCvmAgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpCvmAgent)(nil)).Elem()
}

func (o TmpCvmAgentMapOutput) ToTmpCvmAgentMapOutput() TmpCvmAgentMapOutput {
	return o
}

func (o TmpCvmAgentMapOutput) ToTmpCvmAgentMapOutputWithContext(ctx context.Context) TmpCvmAgentMapOutput {
	return o
}

func (o TmpCvmAgentMapOutput) MapIndex(k pulumi.StringInput) TmpCvmAgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpCvmAgent {
		return vs[0].(map[string]*TmpCvmAgent)[vs[1].(string)]
	}).(TmpCvmAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpCvmAgentInput)(nil)).Elem(), &TmpCvmAgent{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpCvmAgentArrayInput)(nil)).Elem(), TmpCvmAgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpCvmAgentMapInput)(nil)).Elem(), TmpCvmAgentMap{})
	pulumi.RegisterOutputType(TmpCvmAgentOutput{})
	pulumi.RegisterOutputType(TmpCvmAgentArrayOutput{})
	pulumi.RegisterOutputType(TmpCvmAgentMapOutput{})
}
