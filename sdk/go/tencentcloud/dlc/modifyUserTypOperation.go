// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dlc modifyUserTypOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dlc.NewModifyUserTypOperation(ctx, "modifyUserTypOperation", &Dlc.ModifyUserTypOperationArgs{
//				UserId:   pulumi.String("127382378"),
//				UserType: pulumi.String("ADMIN"),
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
// dlc modify_user_typ_operation can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation modify_user_typ_operation modify_user_typ_operation_id
//
// ```
type ModifyUserTypOperation struct {
	pulumi.CustomResourceState

	// User id (uin), if left blank, it defaults to the caller's sub-uin.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
	UserType pulumi.StringOutput `pulumi:"userType"`
}

// NewModifyUserTypOperation registers a new resource with the given unique name, arguments, and options.
func NewModifyUserTypOperation(ctx *pulumi.Context,
	name string, args *ModifyUserTypOperationArgs, opts ...pulumi.ResourceOption) (*ModifyUserTypOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.UserType == nil {
		return nil, errors.New("invalid value for required argument 'UserType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ModifyUserTypOperation
	err := ctx.RegisterResource("tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModifyUserTypOperation gets an existing ModifyUserTypOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModifyUserTypOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModifyUserTypOperationState, opts ...pulumi.ResourceOption) (*ModifyUserTypOperation, error) {
	var resource ModifyUserTypOperation
	err := ctx.ReadResource("tencentcloud:Dlc/modifyUserTypOperation:ModifyUserTypOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModifyUserTypOperation resources.
type modifyUserTypOperationState struct {
	// User id (uin), if left blank, it defaults to the caller's sub-uin.
	UserId *string `pulumi:"userId"`
	// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
	UserType *string `pulumi:"userType"`
}

type ModifyUserTypOperationState struct {
	// User id (uin), if left blank, it defaults to the caller's sub-uin.
	UserId pulumi.StringPtrInput
	// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
	UserType pulumi.StringPtrInput
}

func (ModifyUserTypOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyUserTypOperationState)(nil)).Elem()
}

type modifyUserTypOperationArgs struct {
	// User id (uin), if left blank, it defaults to the caller's sub-uin.
	UserId string `pulumi:"userId"`
	// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
	UserType string `pulumi:"userType"`
}

// The set of arguments for constructing a ModifyUserTypOperation resource.
type ModifyUserTypOperationArgs struct {
	// User id (uin), if left blank, it defaults to the caller's sub-uin.
	UserId pulumi.StringInput
	// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
	UserType pulumi.StringInput
}

func (ModifyUserTypOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modifyUserTypOperationArgs)(nil)).Elem()
}

type ModifyUserTypOperationInput interface {
	pulumi.Input

	ToModifyUserTypOperationOutput() ModifyUserTypOperationOutput
	ToModifyUserTypOperationOutputWithContext(ctx context.Context) ModifyUserTypOperationOutput
}

func (*ModifyUserTypOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyUserTypOperation)(nil)).Elem()
}

func (i *ModifyUserTypOperation) ToModifyUserTypOperationOutput() ModifyUserTypOperationOutput {
	return i.ToModifyUserTypOperationOutputWithContext(context.Background())
}

func (i *ModifyUserTypOperation) ToModifyUserTypOperationOutputWithContext(ctx context.Context) ModifyUserTypOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyUserTypOperationOutput)
}

// ModifyUserTypOperationArrayInput is an input type that accepts ModifyUserTypOperationArray and ModifyUserTypOperationArrayOutput values.
// You can construct a concrete instance of `ModifyUserTypOperationArrayInput` via:
//
//	ModifyUserTypOperationArray{ ModifyUserTypOperationArgs{...} }
type ModifyUserTypOperationArrayInput interface {
	pulumi.Input

	ToModifyUserTypOperationArrayOutput() ModifyUserTypOperationArrayOutput
	ToModifyUserTypOperationArrayOutputWithContext(context.Context) ModifyUserTypOperationArrayOutput
}

type ModifyUserTypOperationArray []ModifyUserTypOperationInput

func (ModifyUserTypOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyUserTypOperation)(nil)).Elem()
}

func (i ModifyUserTypOperationArray) ToModifyUserTypOperationArrayOutput() ModifyUserTypOperationArrayOutput {
	return i.ToModifyUserTypOperationArrayOutputWithContext(context.Background())
}

func (i ModifyUserTypOperationArray) ToModifyUserTypOperationArrayOutputWithContext(ctx context.Context) ModifyUserTypOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyUserTypOperationArrayOutput)
}

// ModifyUserTypOperationMapInput is an input type that accepts ModifyUserTypOperationMap and ModifyUserTypOperationMapOutput values.
// You can construct a concrete instance of `ModifyUserTypOperationMapInput` via:
//
//	ModifyUserTypOperationMap{ "key": ModifyUserTypOperationArgs{...} }
type ModifyUserTypOperationMapInput interface {
	pulumi.Input

	ToModifyUserTypOperationMapOutput() ModifyUserTypOperationMapOutput
	ToModifyUserTypOperationMapOutputWithContext(context.Context) ModifyUserTypOperationMapOutput
}

type ModifyUserTypOperationMap map[string]ModifyUserTypOperationInput

func (ModifyUserTypOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyUserTypOperation)(nil)).Elem()
}

func (i ModifyUserTypOperationMap) ToModifyUserTypOperationMapOutput() ModifyUserTypOperationMapOutput {
	return i.ToModifyUserTypOperationMapOutputWithContext(context.Background())
}

func (i ModifyUserTypOperationMap) ToModifyUserTypOperationMapOutputWithContext(ctx context.Context) ModifyUserTypOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModifyUserTypOperationMapOutput)
}

type ModifyUserTypOperationOutput struct{ *pulumi.OutputState }

func (ModifyUserTypOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModifyUserTypOperation)(nil)).Elem()
}

func (o ModifyUserTypOperationOutput) ToModifyUserTypOperationOutput() ModifyUserTypOperationOutput {
	return o
}

func (o ModifyUserTypOperationOutput) ToModifyUserTypOperationOutputWithContext(ctx context.Context) ModifyUserTypOperationOutput {
	return o
}

// User id (uin), if left blank, it defaults to the caller's sub-uin.
func (o ModifyUserTypOperationOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyUserTypOperation) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
func (o ModifyUserTypOperationOutput) UserType() pulumi.StringOutput {
	return o.ApplyT(func(v *ModifyUserTypOperation) pulumi.StringOutput { return v.UserType }).(pulumi.StringOutput)
}

type ModifyUserTypOperationArrayOutput struct{ *pulumi.OutputState }

func (ModifyUserTypOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ModifyUserTypOperation)(nil)).Elem()
}

func (o ModifyUserTypOperationArrayOutput) ToModifyUserTypOperationArrayOutput() ModifyUserTypOperationArrayOutput {
	return o
}

func (o ModifyUserTypOperationArrayOutput) ToModifyUserTypOperationArrayOutputWithContext(ctx context.Context) ModifyUserTypOperationArrayOutput {
	return o
}

func (o ModifyUserTypOperationArrayOutput) Index(i pulumi.IntInput) ModifyUserTypOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ModifyUserTypOperation {
		return vs[0].([]*ModifyUserTypOperation)[vs[1].(int)]
	}).(ModifyUserTypOperationOutput)
}

type ModifyUserTypOperationMapOutput struct{ *pulumi.OutputState }

func (ModifyUserTypOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ModifyUserTypOperation)(nil)).Elem()
}

func (o ModifyUserTypOperationMapOutput) ToModifyUserTypOperationMapOutput() ModifyUserTypOperationMapOutput {
	return o
}

func (o ModifyUserTypOperationMapOutput) ToModifyUserTypOperationMapOutputWithContext(ctx context.Context) ModifyUserTypOperationMapOutput {
	return o
}

func (o ModifyUserTypOperationMapOutput) MapIndex(k pulumi.StringInput) ModifyUserTypOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ModifyUserTypOperation {
		return vs[0].(map[string]*ModifyUserTypOperation)[vs[1].(string)]
	}).(ModifyUserTypOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyUserTypOperationInput)(nil)).Elem(), &ModifyUserTypOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyUserTypOperationArrayInput)(nil)).Elem(), ModifyUserTypOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModifyUserTypOperationMapInput)(nil)).Elem(), ModifyUserTypOperationMap{})
	pulumi.RegisterOutputType(ModifyUserTypOperationOutput{})
	pulumi.RegisterOutputType(ModifyUserTypOperationArrayOutput{})
	pulumi.RegisterOutputType(ModifyUserTypOperationMapOutput{})
}
