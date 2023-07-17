// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a as completeLifecycle
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := As.NewCompleteLifecycle(ctx, "completeLifecycle", &As.CompleteLifecycleArgs{
// 			InstanceId:            pulumi.String("ins-xxxxxxxx"),
// 			LifecycleActionResult: pulumi.String("CONTINUE"),
// 			LifecycleHookId:       pulumi.String("ash-xxxxxxxx"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CompleteLifecycle struct {
	pulumi.CustomResourceState

	// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
	LifecycleActionResult pulumi.StringOutput `pulumi:"lifecycleActionResult"`
	// Either InstanceId or LifecycleActionToken must be specified.
	LifecycleActionToken pulumi.StringPtrOutput `pulumi:"lifecycleActionToken"`
	// Lifecycle hook ID.
	LifecycleHookId pulumi.StringOutput `pulumi:"lifecycleHookId"`
}

// NewCompleteLifecycle registers a new resource with the given unique name, arguments, and options.
func NewCompleteLifecycle(ctx *pulumi.Context,
	name string, args *CompleteLifecycleArgs, opts ...pulumi.ResourceOption) (*CompleteLifecycle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LifecycleActionResult == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleActionResult'")
	}
	if args.LifecycleHookId == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleHookId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CompleteLifecycle
	err := ctx.RegisterResource("tencentcloud:As/completeLifecycle:CompleteLifecycle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompleteLifecycle gets an existing CompleteLifecycle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompleteLifecycle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompleteLifecycleState, opts ...pulumi.ResourceOption) (*CompleteLifecycle, error) {
	var resource CompleteLifecycle
	err := ctx.ReadResource("tencentcloud:As/completeLifecycle:CompleteLifecycle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CompleteLifecycle resources.
type completeLifecycleState struct {
	// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
	InstanceId *string `pulumi:"instanceId"`
	// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
	LifecycleActionResult *string `pulumi:"lifecycleActionResult"`
	// Either InstanceId or LifecycleActionToken must be specified.
	LifecycleActionToken *string `pulumi:"lifecycleActionToken"`
	// Lifecycle hook ID.
	LifecycleHookId *string `pulumi:"lifecycleHookId"`
}

type CompleteLifecycleState struct {
	// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
	InstanceId pulumi.StringPtrInput
	// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
	LifecycleActionResult pulumi.StringPtrInput
	// Either InstanceId or LifecycleActionToken must be specified.
	LifecycleActionToken pulumi.StringPtrInput
	// Lifecycle hook ID.
	LifecycleHookId pulumi.StringPtrInput
}

func (CompleteLifecycleState) ElementType() reflect.Type {
	return reflect.TypeOf((*completeLifecycleState)(nil)).Elem()
}

type completeLifecycleArgs struct {
	// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
	InstanceId *string `pulumi:"instanceId"`
	// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
	LifecycleActionResult string `pulumi:"lifecycleActionResult"`
	// Either InstanceId or LifecycleActionToken must be specified.
	LifecycleActionToken *string `pulumi:"lifecycleActionToken"`
	// Lifecycle hook ID.
	LifecycleHookId string `pulumi:"lifecycleHookId"`
}

// The set of arguments for constructing a CompleteLifecycle resource.
type CompleteLifecycleArgs struct {
	// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
	InstanceId pulumi.StringPtrInput
	// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
	LifecycleActionResult pulumi.StringInput
	// Either InstanceId or LifecycleActionToken must be specified.
	LifecycleActionToken pulumi.StringPtrInput
	// Lifecycle hook ID.
	LifecycleHookId pulumi.StringInput
}

func (CompleteLifecycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*completeLifecycleArgs)(nil)).Elem()
}

type CompleteLifecycleInput interface {
	pulumi.Input

	ToCompleteLifecycleOutput() CompleteLifecycleOutput
	ToCompleteLifecycleOutputWithContext(ctx context.Context) CompleteLifecycleOutput
}

func (*CompleteLifecycle) ElementType() reflect.Type {
	return reflect.TypeOf((**CompleteLifecycle)(nil)).Elem()
}

func (i *CompleteLifecycle) ToCompleteLifecycleOutput() CompleteLifecycleOutput {
	return i.ToCompleteLifecycleOutputWithContext(context.Background())
}

func (i *CompleteLifecycle) ToCompleteLifecycleOutputWithContext(ctx context.Context) CompleteLifecycleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteLifecycleOutput)
}

// CompleteLifecycleArrayInput is an input type that accepts CompleteLifecycleArray and CompleteLifecycleArrayOutput values.
// You can construct a concrete instance of `CompleteLifecycleArrayInput` via:
//
//          CompleteLifecycleArray{ CompleteLifecycleArgs{...} }
type CompleteLifecycleArrayInput interface {
	pulumi.Input

	ToCompleteLifecycleArrayOutput() CompleteLifecycleArrayOutput
	ToCompleteLifecycleArrayOutputWithContext(context.Context) CompleteLifecycleArrayOutput
}

type CompleteLifecycleArray []CompleteLifecycleInput

func (CompleteLifecycleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompleteLifecycle)(nil)).Elem()
}

func (i CompleteLifecycleArray) ToCompleteLifecycleArrayOutput() CompleteLifecycleArrayOutput {
	return i.ToCompleteLifecycleArrayOutputWithContext(context.Background())
}

func (i CompleteLifecycleArray) ToCompleteLifecycleArrayOutputWithContext(ctx context.Context) CompleteLifecycleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteLifecycleArrayOutput)
}

// CompleteLifecycleMapInput is an input type that accepts CompleteLifecycleMap and CompleteLifecycleMapOutput values.
// You can construct a concrete instance of `CompleteLifecycleMapInput` via:
//
//          CompleteLifecycleMap{ "key": CompleteLifecycleArgs{...} }
type CompleteLifecycleMapInput interface {
	pulumi.Input

	ToCompleteLifecycleMapOutput() CompleteLifecycleMapOutput
	ToCompleteLifecycleMapOutputWithContext(context.Context) CompleteLifecycleMapOutput
}

type CompleteLifecycleMap map[string]CompleteLifecycleInput

func (CompleteLifecycleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompleteLifecycle)(nil)).Elem()
}

func (i CompleteLifecycleMap) ToCompleteLifecycleMapOutput() CompleteLifecycleMapOutput {
	return i.ToCompleteLifecycleMapOutputWithContext(context.Background())
}

func (i CompleteLifecycleMap) ToCompleteLifecycleMapOutputWithContext(ctx context.Context) CompleteLifecycleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteLifecycleMapOutput)
}

type CompleteLifecycleOutput struct{ *pulumi.OutputState }

func (CompleteLifecycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompleteLifecycle)(nil)).Elem()
}

func (o CompleteLifecycleOutput) ToCompleteLifecycleOutput() CompleteLifecycleOutput {
	return o
}

func (o CompleteLifecycleOutput) ToCompleteLifecycleOutputWithContext(ctx context.Context) CompleteLifecycleOutput {
	return o
}

// Instance ID. Either InstanceId or LifecycleActionToken must be specified.
func (o CompleteLifecycleOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompleteLifecycle) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// Result of the lifecycle action. Value range: `CONTINUE`, `ABANDON`.
func (o CompleteLifecycleOutput) LifecycleActionResult() pulumi.StringOutput {
	return o.ApplyT(func(v *CompleteLifecycle) pulumi.StringOutput { return v.LifecycleActionResult }).(pulumi.StringOutput)
}

// Either InstanceId or LifecycleActionToken must be specified.
func (o CompleteLifecycleOutput) LifecycleActionToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompleteLifecycle) pulumi.StringPtrOutput { return v.LifecycleActionToken }).(pulumi.StringPtrOutput)
}

// Lifecycle hook ID.
func (o CompleteLifecycleOutput) LifecycleHookId() pulumi.StringOutput {
	return o.ApplyT(func(v *CompleteLifecycle) pulumi.StringOutput { return v.LifecycleHookId }).(pulumi.StringOutput)
}

type CompleteLifecycleArrayOutput struct{ *pulumi.OutputState }

func (CompleteLifecycleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompleteLifecycle)(nil)).Elem()
}

func (o CompleteLifecycleArrayOutput) ToCompleteLifecycleArrayOutput() CompleteLifecycleArrayOutput {
	return o
}

func (o CompleteLifecycleArrayOutput) ToCompleteLifecycleArrayOutputWithContext(ctx context.Context) CompleteLifecycleArrayOutput {
	return o
}

func (o CompleteLifecycleArrayOutput) Index(i pulumi.IntInput) CompleteLifecycleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CompleteLifecycle {
		return vs[0].([]*CompleteLifecycle)[vs[1].(int)]
	}).(CompleteLifecycleOutput)
}

type CompleteLifecycleMapOutput struct{ *pulumi.OutputState }

func (CompleteLifecycleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompleteLifecycle)(nil)).Elem()
}

func (o CompleteLifecycleMapOutput) ToCompleteLifecycleMapOutput() CompleteLifecycleMapOutput {
	return o
}

func (o CompleteLifecycleMapOutput) ToCompleteLifecycleMapOutputWithContext(ctx context.Context) CompleteLifecycleMapOutput {
	return o
}

func (o CompleteLifecycleMapOutput) MapIndex(k pulumi.StringInput) CompleteLifecycleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CompleteLifecycle {
		return vs[0].(map[string]*CompleteLifecycle)[vs[1].(string)]
	}).(CompleteLifecycleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteLifecycleInput)(nil)).Elem(), &CompleteLifecycle{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteLifecycleArrayInput)(nil)).Elem(), CompleteLifecycleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteLifecycleMapInput)(nil)).Elem(), CompleteLifecycleMap{})
	pulumi.RegisterOutputType(CompleteLifecycleOutput{})
	pulumi.RegisterOutputType(CompleteLifecycleArrayOutput{})
	pulumi.RegisterOutputType(CompleteLifecycleMapOutput{})
}
