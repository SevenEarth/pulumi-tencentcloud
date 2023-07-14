// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis renewInstanceOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Redis.NewRenewInstanceOperation(ctx, "renewInstanceOperation", &Redis.RenewInstanceOperationArgs{
// 			InstanceId:    pulumi.String("crs-c1nl9rpv"),
// 			ModifyPayMode: pulumi.String("prepaid"),
// 			Period:        pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RenewInstanceOperation struct {
	pulumi.CustomResourceState

	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
	ModifyPayMode pulumi.StringPtrOutput `pulumi:"modifyPayMode"`
	// Purchase duration, in months.
	Period pulumi.IntOutput `pulumi:"period"`
}

// NewRenewInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewRenewInstanceOperation(ctx *pulumi.Context,
	name string, args *RenewInstanceOperationArgs, opts ...pulumi.ResourceOption) (*RenewInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RenewInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Redis/renewInstanceOperation:RenewInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRenewInstanceOperation gets an existing RenewInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRenewInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenewInstanceOperationState, opts ...pulumi.ResourceOption) (*RenewInstanceOperation, error) {
	var resource RenewInstanceOperation
	err := ctx.ReadResource("tencentcloud:Redis/renewInstanceOperation:RenewInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RenewInstanceOperation resources.
type renewInstanceOperationState struct {
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
	ModifyPayMode *string `pulumi:"modifyPayMode"`
	// Purchase duration, in months.
	Period *int `pulumi:"period"`
}

type RenewInstanceOperationState struct {
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
	// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
	ModifyPayMode pulumi.StringPtrInput
	// Purchase duration, in months.
	Period pulumi.IntPtrInput
}

func (RenewInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*renewInstanceOperationState)(nil)).Elem()
}

type renewInstanceOperationArgs struct {
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
	ModifyPayMode *string `pulumi:"modifyPayMode"`
	// Purchase duration, in months.
	Period int `pulumi:"period"`
}

// The set of arguments for constructing a RenewInstanceOperation resource.
type RenewInstanceOperationArgs struct {
	// The ID of instance.
	InstanceId pulumi.StringInput
	// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
	ModifyPayMode pulumi.StringPtrInput
	// Purchase duration, in months.
	Period pulumi.IntInput
}

func (RenewInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renewInstanceOperationArgs)(nil)).Elem()
}

type RenewInstanceOperationInput interface {
	pulumi.Input

	ToRenewInstanceOperationOutput() RenewInstanceOperationOutput
	ToRenewInstanceOperationOutputWithContext(ctx context.Context) RenewInstanceOperationOutput
}

func (*RenewInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewInstanceOperation)(nil)).Elem()
}

func (i *RenewInstanceOperation) ToRenewInstanceOperationOutput() RenewInstanceOperationOutput {
	return i.ToRenewInstanceOperationOutputWithContext(context.Background())
}

func (i *RenewInstanceOperation) ToRenewInstanceOperationOutputWithContext(ctx context.Context) RenewInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceOperationOutput)
}

// RenewInstanceOperationArrayInput is an input type that accepts RenewInstanceOperationArray and RenewInstanceOperationArrayOutput values.
// You can construct a concrete instance of `RenewInstanceOperationArrayInput` via:
//
//          RenewInstanceOperationArray{ RenewInstanceOperationArgs{...} }
type RenewInstanceOperationArrayInput interface {
	pulumi.Input

	ToRenewInstanceOperationArrayOutput() RenewInstanceOperationArrayOutput
	ToRenewInstanceOperationArrayOutputWithContext(context.Context) RenewInstanceOperationArrayOutput
}

type RenewInstanceOperationArray []RenewInstanceOperationInput

func (RenewInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewInstanceOperation)(nil)).Elem()
}

func (i RenewInstanceOperationArray) ToRenewInstanceOperationArrayOutput() RenewInstanceOperationArrayOutput {
	return i.ToRenewInstanceOperationArrayOutputWithContext(context.Background())
}

func (i RenewInstanceOperationArray) ToRenewInstanceOperationArrayOutputWithContext(ctx context.Context) RenewInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceOperationArrayOutput)
}

// RenewInstanceOperationMapInput is an input type that accepts RenewInstanceOperationMap and RenewInstanceOperationMapOutput values.
// You can construct a concrete instance of `RenewInstanceOperationMapInput` via:
//
//          RenewInstanceOperationMap{ "key": RenewInstanceOperationArgs{...} }
type RenewInstanceOperationMapInput interface {
	pulumi.Input

	ToRenewInstanceOperationMapOutput() RenewInstanceOperationMapOutput
	ToRenewInstanceOperationMapOutputWithContext(context.Context) RenewInstanceOperationMapOutput
}

type RenewInstanceOperationMap map[string]RenewInstanceOperationInput

func (RenewInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewInstanceOperation)(nil)).Elem()
}

func (i RenewInstanceOperationMap) ToRenewInstanceOperationMapOutput() RenewInstanceOperationMapOutput {
	return i.ToRenewInstanceOperationMapOutputWithContext(context.Background())
}

func (i RenewInstanceOperationMap) ToRenewInstanceOperationMapOutputWithContext(ctx context.Context) RenewInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceOperationMapOutput)
}

type RenewInstanceOperationOutput struct{ *pulumi.OutputState }

func (RenewInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewInstanceOperation)(nil)).Elem()
}

func (o RenewInstanceOperationOutput) ToRenewInstanceOperationOutput() RenewInstanceOperationOutput {
	return o
}

func (o RenewInstanceOperationOutput) ToRenewInstanceOperationOutputWithContext(ctx context.Context) RenewInstanceOperationOutput {
	return o
}

// The ID of instance.
func (o RenewInstanceOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RenewInstanceOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Identifies whether the billing model is modified:The current instance billing mode is pay-as-you-go, which is prepaid and renewed.The billing mode of the current instance is subscription and you can not set this parameter.
func (o RenewInstanceOperationOutput) ModifyPayMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RenewInstanceOperation) pulumi.StringPtrOutput { return v.ModifyPayMode }).(pulumi.StringPtrOutput)
}

// Purchase duration, in months.
func (o RenewInstanceOperationOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *RenewInstanceOperation) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

type RenewInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (RenewInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewInstanceOperation)(nil)).Elem()
}

func (o RenewInstanceOperationArrayOutput) ToRenewInstanceOperationArrayOutput() RenewInstanceOperationArrayOutput {
	return o
}

func (o RenewInstanceOperationArrayOutput) ToRenewInstanceOperationArrayOutputWithContext(ctx context.Context) RenewInstanceOperationArrayOutput {
	return o
}

func (o RenewInstanceOperationArrayOutput) Index(i pulumi.IntInput) RenewInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RenewInstanceOperation {
		return vs[0].([]*RenewInstanceOperation)[vs[1].(int)]
	}).(RenewInstanceOperationOutput)
}

type RenewInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (RenewInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewInstanceOperation)(nil)).Elem()
}

func (o RenewInstanceOperationMapOutput) ToRenewInstanceOperationMapOutput() RenewInstanceOperationMapOutput {
	return o
}

func (o RenewInstanceOperationMapOutput) ToRenewInstanceOperationMapOutputWithContext(ctx context.Context) RenewInstanceOperationMapOutput {
	return o
}

func (o RenewInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) RenewInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RenewInstanceOperation {
		return vs[0].(map[string]*RenewInstanceOperation)[vs[1].(string)]
	}).(RenewInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceOperationInput)(nil)).Elem(), &RenewInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceOperationArrayInput)(nil)).Elem(), RenewInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceOperationMapInput)(nil)).Elem(), RenewInstanceOperationMap{})
	pulumi.RegisterOutputType(RenewInstanceOperationOutput{})
	pulumi.RegisterOutputType(RenewInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(RenewInstanceOperationMapOutput{})
}
