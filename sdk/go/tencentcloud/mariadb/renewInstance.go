// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mariadb renewInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.NewRenewInstance(ctx, "renewInstance", &Mariadb.RenewInstanceArgs{
//				InstanceId: pulumi.String("tdsql-9vqvls95"),
//				Period:     pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RenewInstance struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Renewal duration, unit: month.
	Period pulumi.IntOutput `pulumi:"period"`
}

// NewRenewInstance registers a new resource with the given unique name, arguments, and options.
func NewRenewInstance(ctx *pulumi.Context,
	name string, args *RenewInstanceArgs, opts ...pulumi.ResourceOption) (*RenewInstance, error) {
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
	var resource RenewInstance
	err := ctx.RegisterResource("tencentcloud:Mariadb/renewInstance:RenewInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRenewInstance gets an existing RenewInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRenewInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenewInstanceState, opts ...pulumi.ResourceOption) (*RenewInstance, error) {
	var resource RenewInstance
	err := ctx.ReadResource("tencentcloud:Mariadb/renewInstance:RenewInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RenewInstance resources.
type renewInstanceState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Renewal duration, unit: month.
	Period *int `pulumi:"period"`
}

type RenewInstanceState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Renewal duration, unit: month.
	Period pulumi.IntPtrInput
}

func (RenewInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*renewInstanceState)(nil)).Elem()
}

type renewInstanceArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Renewal duration, unit: month.
	Period int `pulumi:"period"`
}

// The set of arguments for constructing a RenewInstance resource.
type RenewInstanceArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
	// Renewal duration, unit: month.
	Period pulumi.IntInput
}

func (RenewInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renewInstanceArgs)(nil)).Elem()
}

type RenewInstanceInput interface {
	pulumi.Input

	ToRenewInstanceOutput() RenewInstanceOutput
	ToRenewInstanceOutputWithContext(ctx context.Context) RenewInstanceOutput
}

func (*RenewInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewInstance)(nil)).Elem()
}

func (i *RenewInstance) ToRenewInstanceOutput() RenewInstanceOutput {
	return i.ToRenewInstanceOutputWithContext(context.Background())
}

func (i *RenewInstance) ToRenewInstanceOutputWithContext(ctx context.Context) RenewInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceOutput)
}

// RenewInstanceArrayInput is an input type that accepts RenewInstanceArray and RenewInstanceArrayOutput values.
// You can construct a concrete instance of `RenewInstanceArrayInput` via:
//
//	RenewInstanceArray{ RenewInstanceArgs{...} }
type RenewInstanceArrayInput interface {
	pulumi.Input

	ToRenewInstanceArrayOutput() RenewInstanceArrayOutput
	ToRenewInstanceArrayOutputWithContext(context.Context) RenewInstanceArrayOutput
}

type RenewInstanceArray []RenewInstanceInput

func (RenewInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewInstance)(nil)).Elem()
}

func (i RenewInstanceArray) ToRenewInstanceArrayOutput() RenewInstanceArrayOutput {
	return i.ToRenewInstanceArrayOutputWithContext(context.Background())
}

func (i RenewInstanceArray) ToRenewInstanceArrayOutputWithContext(ctx context.Context) RenewInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceArrayOutput)
}

// RenewInstanceMapInput is an input type that accepts RenewInstanceMap and RenewInstanceMapOutput values.
// You can construct a concrete instance of `RenewInstanceMapInput` via:
//
//	RenewInstanceMap{ "key": RenewInstanceArgs{...} }
type RenewInstanceMapInput interface {
	pulumi.Input

	ToRenewInstanceMapOutput() RenewInstanceMapOutput
	ToRenewInstanceMapOutputWithContext(context.Context) RenewInstanceMapOutput
}

type RenewInstanceMap map[string]RenewInstanceInput

func (RenewInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewInstance)(nil)).Elem()
}

func (i RenewInstanceMap) ToRenewInstanceMapOutput() RenewInstanceMapOutput {
	return i.ToRenewInstanceMapOutputWithContext(context.Background())
}

func (i RenewInstanceMap) ToRenewInstanceMapOutputWithContext(ctx context.Context) RenewInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewInstanceMapOutput)
}

type RenewInstanceOutput struct{ *pulumi.OutputState }

func (RenewInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewInstance)(nil)).Elem()
}

func (o RenewInstanceOutput) ToRenewInstanceOutput() RenewInstanceOutput {
	return o
}

func (o RenewInstanceOutput) ToRenewInstanceOutputWithContext(ctx context.Context) RenewInstanceOutput {
	return o
}

// Instance ID.
func (o RenewInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RenewInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Renewal duration, unit: month.
func (o RenewInstanceOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *RenewInstance) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

type RenewInstanceArrayOutput struct{ *pulumi.OutputState }

func (RenewInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewInstance)(nil)).Elem()
}

func (o RenewInstanceArrayOutput) ToRenewInstanceArrayOutput() RenewInstanceArrayOutput {
	return o
}

func (o RenewInstanceArrayOutput) ToRenewInstanceArrayOutputWithContext(ctx context.Context) RenewInstanceArrayOutput {
	return o
}

func (o RenewInstanceArrayOutput) Index(i pulumi.IntInput) RenewInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RenewInstance {
		return vs[0].([]*RenewInstance)[vs[1].(int)]
	}).(RenewInstanceOutput)
}

type RenewInstanceMapOutput struct{ *pulumi.OutputState }

func (RenewInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewInstance)(nil)).Elem()
}

func (o RenewInstanceMapOutput) ToRenewInstanceMapOutput() RenewInstanceMapOutput {
	return o
}

func (o RenewInstanceMapOutput) ToRenewInstanceMapOutputWithContext(ctx context.Context) RenewInstanceMapOutput {
	return o
}

func (o RenewInstanceMapOutput) MapIndex(k pulumi.StringInput) RenewInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RenewInstance {
		return vs[0].(map[string]*RenewInstance)[vs[1].(string)]
	}).(RenewInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceInput)(nil)).Elem(), &RenewInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceArrayInput)(nil)).Elem(), RenewInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewInstanceMapInput)(nil)).Elem(), RenewInstanceMap{})
	pulumi.RegisterOutputType(RenewInstanceOutput{})
	pulumi.RegisterOutputType(RenewInstanceArrayOutput{})
	pulumi.RegisterOutputType(RenewInstanceMapOutput{})
}
