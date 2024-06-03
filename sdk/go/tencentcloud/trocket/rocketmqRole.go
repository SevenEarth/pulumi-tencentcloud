// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trocket

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a trocket rocketmqRole
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Trocket"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rocketmqInstance, err := Trocket.NewRocketmqInstance(ctx, "rocketmqInstance", &Trocket.RocketmqInstanceArgs{
//				InstanceType: pulumi.String("EXPERIMENT"),
//				SkuCode:      pulumi.String("experiment_500"),
//				Remark:       pulumi.String("test"),
//				VpcId:        pulumi.String("vpc-xxxxx"),
//				SubnetId:     pulumi.String("subnet-xxxxx"),
//				Tags: pulumi.Map{
//					"tag_key":   pulumi.Any("rocketmq"),
//					"tag_value": pulumi.Any("5.x"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			rocketmqRole, err := Trocket.NewRocketmqRole(ctx, "rocketmqRole", &Trocket.RocketmqRoleArgs{
//				InstanceId: rocketmqInstance.ID(),
//				Role:       pulumi.String("test_role"),
//				Remark:     pulumi.String("test for terraform"),
//				PermWrite:  pulumi.Bool(false),
//				PermRead:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("accessKey", rocketmqRole.AccessKey)
//			ctx.Export("secretKey", rocketmqRole.SecretKey)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// trocket rocketmq_role can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Trocket/rocketmqRole:RocketmqRole rocketmq_role instanceId#role
// ```
type RocketmqRole struct {
	pulumi.CustomResourceState

	// Access key.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// Created time.
	CreatedTime pulumi.IntOutput `pulumi:"createdTime"`
	// ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Modified time.
	ModifiedTime pulumi.IntOutput `pulumi:"modifiedTime"`
	// Whether to enable consumption permission.
	PermRead pulumi.BoolOutput `pulumi:"permRead"`
	// Whether to enable production permission.
	PermWrite pulumi.BoolOutput `pulumi:"permWrite"`
	// remark.
	Remark pulumi.StringOutput `pulumi:"remark"`
	// Name of role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Secret key.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
}

// NewRocketmqRole registers a new resource with the given unique name, arguments, and options.
func NewRocketmqRole(ctx *pulumi.Context,
	name string, args *RocketmqRoleArgs, opts ...pulumi.ResourceOption) (*RocketmqRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PermRead == nil {
		return nil, errors.New("invalid value for required argument 'PermRead'")
	}
	if args.PermWrite == nil {
		return nil, errors.New("invalid value for required argument 'PermWrite'")
	}
	if args.Remark == nil {
		return nil, errors.New("invalid value for required argument 'Remark'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RocketmqRole
	err := ctx.RegisterResource("tencentcloud:Trocket/rocketmqRole:RocketmqRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRocketmqRole gets an existing RocketmqRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRocketmqRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RocketmqRoleState, opts ...pulumi.ResourceOption) (*RocketmqRole, error) {
	var resource RocketmqRole
	err := ctx.ReadResource("tencentcloud:Trocket/rocketmqRole:RocketmqRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RocketmqRole resources.
type rocketmqRoleState struct {
	// Access key.
	AccessKey *string `pulumi:"accessKey"`
	// Created time.
	CreatedTime *int `pulumi:"createdTime"`
	// ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Modified time.
	ModifiedTime *int `pulumi:"modifiedTime"`
	// Whether to enable consumption permission.
	PermRead *bool `pulumi:"permRead"`
	// Whether to enable production permission.
	PermWrite *bool `pulumi:"permWrite"`
	// remark.
	Remark *string `pulumi:"remark"`
	// Name of role.
	Role *string `pulumi:"role"`
	// Secret key.
	SecretKey *string `pulumi:"secretKey"`
}

type RocketmqRoleState struct {
	// Access key.
	AccessKey pulumi.StringPtrInput
	// Created time.
	CreatedTime pulumi.IntPtrInput
	// ID of instance.
	InstanceId pulumi.StringPtrInput
	// Modified time.
	ModifiedTime pulumi.IntPtrInput
	// Whether to enable consumption permission.
	PermRead pulumi.BoolPtrInput
	// Whether to enable production permission.
	PermWrite pulumi.BoolPtrInput
	// remark.
	Remark pulumi.StringPtrInput
	// Name of role.
	Role pulumi.StringPtrInput
	// Secret key.
	SecretKey pulumi.StringPtrInput
}

func (RocketmqRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqRoleState)(nil)).Elem()
}

type rocketmqRoleArgs struct {
	// ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Whether to enable consumption permission.
	PermRead bool `pulumi:"permRead"`
	// Whether to enable production permission.
	PermWrite bool `pulumi:"permWrite"`
	// remark.
	Remark string `pulumi:"remark"`
	// Name of role.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RocketmqRole resource.
type RocketmqRoleArgs struct {
	// ID of instance.
	InstanceId pulumi.StringInput
	// Whether to enable consumption permission.
	PermRead pulumi.BoolInput
	// Whether to enable production permission.
	PermWrite pulumi.BoolInput
	// remark.
	Remark pulumi.StringInput
	// Name of role.
	Role pulumi.StringInput
}

func (RocketmqRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqRoleArgs)(nil)).Elem()
}

type RocketmqRoleInput interface {
	pulumi.Input

	ToRocketmqRoleOutput() RocketmqRoleOutput
	ToRocketmqRoleOutputWithContext(ctx context.Context) RocketmqRoleOutput
}

func (*RocketmqRole) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqRole)(nil)).Elem()
}

func (i *RocketmqRole) ToRocketmqRoleOutput() RocketmqRoleOutput {
	return i.ToRocketmqRoleOutputWithContext(context.Background())
}

func (i *RocketmqRole) ToRocketmqRoleOutputWithContext(ctx context.Context) RocketmqRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqRoleOutput)
}

// RocketmqRoleArrayInput is an input type that accepts RocketmqRoleArray and RocketmqRoleArrayOutput values.
// You can construct a concrete instance of `RocketmqRoleArrayInput` via:
//
//	RocketmqRoleArray{ RocketmqRoleArgs{...} }
type RocketmqRoleArrayInput interface {
	pulumi.Input

	ToRocketmqRoleArrayOutput() RocketmqRoleArrayOutput
	ToRocketmqRoleArrayOutputWithContext(context.Context) RocketmqRoleArrayOutput
}

type RocketmqRoleArray []RocketmqRoleInput

func (RocketmqRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqRole)(nil)).Elem()
}

func (i RocketmqRoleArray) ToRocketmqRoleArrayOutput() RocketmqRoleArrayOutput {
	return i.ToRocketmqRoleArrayOutputWithContext(context.Background())
}

func (i RocketmqRoleArray) ToRocketmqRoleArrayOutputWithContext(ctx context.Context) RocketmqRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqRoleArrayOutput)
}

// RocketmqRoleMapInput is an input type that accepts RocketmqRoleMap and RocketmqRoleMapOutput values.
// You can construct a concrete instance of `RocketmqRoleMapInput` via:
//
//	RocketmqRoleMap{ "key": RocketmqRoleArgs{...} }
type RocketmqRoleMapInput interface {
	pulumi.Input

	ToRocketmqRoleMapOutput() RocketmqRoleMapOutput
	ToRocketmqRoleMapOutputWithContext(context.Context) RocketmqRoleMapOutput
}

type RocketmqRoleMap map[string]RocketmqRoleInput

func (RocketmqRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqRole)(nil)).Elem()
}

func (i RocketmqRoleMap) ToRocketmqRoleMapOutput() RocketmqRoleMapOutput {
	return i.ToRocketmqRoleMapOutputWithContext(context.Background())
}

func (i RocketmqRoleMap) ToRocketmqRoleMapOutputWithContext(ctx context.Context) RocketmqRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqRoleMapOutput)
}

type RocketmqRoleOutput struct{ *pulumi.OutputState }

func (RocketmqRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqRole)(nil)).Elem()
}

func (o RocketmqRoleOutput) ToRocketmqRoleOutput() RocketmqRoleOutput {
	return o
}

func (o RocketmqRoleOutput) ToRocketmqRoleOutputWithContext(ctx context.Context) RocketmqRoleOutput {
	return o
}

// Access key.
func (o RocketmqRoleOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// Created time.
func (o RocketmqRoleOutput) CreatedTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.IntOutput { return v.CreatedTime }).(pulumi.IntOutput)
}

// ID of instance.
func (o RocketmqRoleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Modified time.
func (o RocketmqRoleOutput) ModifiedTime() pulumi.IntOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.IntOutput { return v.ModifiedTime }).(pulumi.IntOutput)
}

// Whether to enable consumption permission.
func (o RocketmqRoleOutput) PermRead() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.BoolOutput { return v.PermRead }).(pulumi.BoolOutput)
}

// Whether to enable production permission.
func (o RocketmqRoleOutput) PermWrite() pulumi.BoolOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.BoolOutput { return v.PermWrite }).(pulumi.BoolOutput)
}

// remark.
func (o RocketmqRoleOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.StringOutput { return v.Remark }).(pulumi.StringOutput)
}

// Name of role.
func (o RocketmqRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Secret key.
func (o RocketmqRoleOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqRole) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

type RocketmqRoleArrayOutput struct{ *pulumi.OutputState }

func (RocketmqRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqRole)(nil)).Elem()
}

func (o RocketmqRoleArrayOutput) ToRocketmqRoleArrayOutput() RocketmqRoleArrayOutput {
	return o
}

func (o RocketmqRoleArrayOutput) ToRocketmqRoleArrayOutputWithContext(ctx context.Context) RocketmqRoleArrayOutput {
	return o
}

func (o RocketmqRoleArrayOutput) Index(i pulumi.IntInput) RocketmqRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RocketmqRole {
		return vs[0].([]*RocketmqRole)[vs[1].(int)]
	}).(RocketmqRoleOutput)
}

type RocketmqRoleMapOutput struct{ *pulumi.OutputState }

func (RocketmqRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqRole)(nil)).Elem()
}

func (o RocketmqRoleMapOutput) ToRocketmqRoleMapOutput() RocketmqRoleMapOutput {
	return o
}

func (o RocketmqRoleMapOutput) ToRocketmqRoleMapOutputWithContext(ctx context.Context) RocketmqRoleMapOutput {
	return o
}

func (o RocketmqRoleMapOutput) MapIndex(k pulumi.StringInput) RocketmqRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RocketmqRole {
		return vs[0].(map[string]*RocketmqRole)[vs[1].(string)]
	}).(RocketmqRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqRoleInput)(nil)).Elem(), &RocketmqRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqRoleArrayInput)(nil)).Elem(), RocketmqRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqRoleMapInput)(nil)).Elem(), RocketmqRoleMap{})
	pulumi.RegisterOutputType(RocketmqRoleOutput{})
	pulumi.RegisterOutputType(RocketmqRoleArrayOutput{})
	pulumi.RegisterOutputType(RocketmqRoleMapOutput{})
}
