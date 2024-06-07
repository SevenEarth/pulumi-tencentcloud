// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dasb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a dasb bindDeviceAccountPassword
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dasb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDevice, err := Dasb.NewDevice(ctx, "exampleDevice", &Dasb.DeviceArgs{
//				OsName: pulumi.String("Linux"),
//				Ip:     pulumi.String("192.168.0.1"),
//				Port:   pulumi.Int(80),
//			})
//			if err != nil {
//				return err
//			}
//			exampleDeviceAccount, err := Dasb.NewDeviceAccount(ctx, "exampleDeviceAccount", &Dasb.DeviceAccountArgs{
//				DeviceId: exampleDevice.ID(),
//				Account:  pulumi.String("root"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Dasb.NewBindDeviceAccountPassword(ctx, "exampleBindDeviceAccountPassword", &Dasb.BindDeviceAccountPasswordArgs{
//				DeviceAccountId: exampleDeviceAccount.ID(),
//				Password:        pulumi.String("TerraformPassword"),
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
type BindDeviceAccountPassword struct {
	pulumi.CustomResourceState

	// Host account ID.
	DeviceAccountId pulumi.IntOutput `pulumi:"deviceAccountId"`
	// Host account password.
	Password pulumi.StringOutput `pulumi:"password"`
}

// NewBindDeviceAccountPassword registers a new resource with the given unique name, arguments, and options.
func NewBindDeviceAccountPassword(ctx *pulumi.Context,
	name string, args *BindDeviceAccountPasswordArgs, opts ...pulumi.ResourceOption) (*BindDeviceAccountPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceAccountId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BindDeviceAccountPassword
	err := ctx.RegisterResource("tencentcloud:Dasb/bindDeviceAccountPassword:BindDeviceAccountPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBindDeviceAccountPassword gets an existing BindDeviceAccountPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBindDeviceAccountPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindDeviceAccountPasswordState, opts ...pulumi.ResourceOption) (*BindDeviceAccountPassword, error) {
	var resource BindDeviceAccountPassword
	err := ctx.ReadResource("tencentcloud:Dasb/bindDeviceAccountPassword:BindDeviceAccountPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BindDeviceAccountPassword resources.
type bindDeviceAccountPasswordState struct {
	// Host account ID.
	DeviceAccountId *int `pulumi:"deviceAccountId"`
	// Host account password.
	Password *string `pulumi:"password"`
}

type BindDeviceAccountPasswordState struct {
	// Host account ID.
	DeviceAccountId pulumi.IntPtrInput
	// Host account password.
	Password pulumi.StringPtrInput
}

func (BindDeviceAccountPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindDeviceAccountPasswordState)(nil)).Elem()
}

type bindDeviceAccountPasswordArgs struct {
	// Host account ID.
	DeviceAccountId int `pulumi:"deviceAccountId"`
	// Host account password.
	Password string `pulumi:"password"`
}

// The set of arguments for constructing a BindDeviceAccountPassword resource.
type BindDeviceAccountPasswordArgs struct {
	// Host account ID.
	DeviceAccountId pulumi.IntInput
	// Host account password.
	Password pulumi.StringInput
}

func (BindDeviceAccountPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindDeviceAccountPasswordArgs)(nil)).Elem()
}

type BindDeviceAccountPasswordInput interface {
	pulumi.Input

	ToBindDeviceAccountPasswordOutput() BindDeviceAccountPasswordOutput
	ToBindDeviceAccountPasswordOutputWithContext(ctx context.Context) BindDeviceAccountPasswordOutput
}

func (*BindDeviceAccountPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**BindDeviceAccountPassword)(nil)).Elem()
}

func (i *BindDeviceAccountPassword) ToBindDeviceAccountPasswordOutput() BindDeviceAccountPasswordOutput {
	return i.ToBindDeviceAccountPasswordOutputWithContext(context.Background())
}

func (i *BindDeviceAccountPassword) ToBindDeviceAccountPasswordOutputWithContext(ctx context.Context) BindDeviceAccountPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindDeviceAccountPasswordOutput)
}

// BindDeviceAccountPasswordArrayInput is an input type that accepts BindDeviceAccountPasswordArray and BindDeviceAccountPasswordArrayOutput values.
// You can construct a concrete instance of `BindDeviceAccountPasswordArrayInput` via:
//
//	BindDeviceAccountPasswordArray{ BindDeviceAccountPasswordArgs{...} }
type BindDeviceAccountPasswordArrayInput interface {
	pulumi.Input

	ToBindDeviceAccountPasswordArrayOutput() BindDeviceAccountPasswordArrayOutput
	ToBindDeviceAccountPasswordArrayOutputWithContext(context.Context) BindDeviceAccountPasswordArrayOutput
}

type BindDeviceAccountPasswordArray []BindDeviceAccountPasswordInput

func (BindDeviceAccountPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindDeviceAccountPassword)(nil)).Elem()
}

func (i BindDeviceAccountPasswordArray) ToBindDeviceAccountPasswordArrayOutput() BindDeviceAccountPasswordArrayOutput {
	return i.ToBindDeviceAccountPasswordArrayOutputWithContext(context.Background())
}

func (i BindDeviceAccountPasswordArray) ToBindDeviceAccountPasswordArrayOutputWithContext(ctx context.Context) BindDeviceAccountPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindDeviceAccountPasswordArrayOutput)
}

// BindDeviceAccountPasswordMapInput is an input type that accepts BindDeviceAccountPasswordMap and BindDeviceAccountPasswordMapOutput values.
// You can construct a concrete instance of `BindDeviceAccountPasswordMapInput` via:
//
//	BindDeviceAccountPasswordMap{ "key": BindDeviceAccountPasswordArgs{...} }
type BindDeviceAccountPasswordMapInput interface {
	pulumi.Input

	ToBindDeviceAccountPasswordMapOutput() BindDeviceAccountPasswordMapOutput
	ToBindDeviceAccountPasswordMapOutputWithContext(context.Context) BindDeviceAccountPasswordMapOutput
}

type BindDeviceAccountPasswordMap map[string]BindDeviceAccountPasswordInput

func (BindDeviceAccountPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindDeviceAccountPassword)(nil)).Elem()
}

func (i BindDeviceAccountPasswordMap) ToBindDeviceAccountPasswordMapOutput() BindDeviceAccountPasswordMapOutput {
	return i.ToBindDeviceAccountPasswordMapOutputWithContext(context.Background())
}

func (i BindDeviceAccountPasswordMap) ToBindDeviceAccountPasswordMapOutputWithContext(ctx context.Context) BindDeviceAccountPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindDeviceAccountPasswordMapOutput)
}

type BindDeviceAccountPasswordOutput struct{ *pulumi.OutputState }

func (BindDeviceAccountPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindDeviceAccountPassword)(nil)).Elem()
}

func (o BindDeviceAccountPasswordOutput) ToBindDeviceAccountPasswordOutput() BindDeviceAccountPasswordOutput {
	return o
}

func (o BindDeviceAccountPasswordOutput) ToBindDeviceAccountPasswordOutputWithContext(ctx context.Context) BindDeviceAccountPasswordOutput {
	return o
}

// Host account ID.
func (o BindDeviceAccountPasswordOutput) DeviceAccountId() pulumi.IntOutput {
	return o.ApplyT(func(v *BindDeviceAccountPassword) pulumi.IntOutput { return v.DeviceAccountId }).(pulumi.IntOutput)
}

// Host account password.
func (o BindDeviceAccountPasswordOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *BindDeviceAccountPassword) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

type BindDeviceAccountPasswordArrayOutput struct{ *pulumi.OutputState }

func (BindDeviceAccountPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindDeviceAccountPassword)(nil)).Elem()
}

func (o BindDeviceAccountPasswordArrayOutput) ToBindDeviceAccountPasswordArrayOutput() BindDeviceAccountPasswordArrayOutput {
	return o
}

func (o BindDeviceAccountPasswordArrayOutput) ToBindDeviceAccountPasswordArrayOutputWithContext(ctx context.Context) BindDeviceAccountPasswordArrayOutput {
	return o
}

func (o BindDeviceAccountPasswordArrayOutput) Index(i pulumi.IntInput) BindDeviceAccountPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BindDeviceAccountPassword {
		return vs[0].([]*BindDeviceAccountPassword)[vs[1].(int)]
	}).(BindDeviceAccountPasswordOutput)
}

type BindDeviceAccountPasswordMapOutput struct{ *pulumi.OutputState }

func (BindDeviceAccountPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindDeviceAccountPassword)(nil)).Elem()
}

func (o BindDeviceAccountPasswordMapOutput) ToBindDeviceAccountPasswordMapOutput() BindDeviceAccountPasswordMapOutput {
	return o
}

func (o BindDeviceAccountPasswordMapOutput) ToBindDeviceAccountPasswordMapOutputWithContext(ctx context.Context) BindDeviceAccountPasswordMapOutput {
	return o
}

func (o BindDeviceAccountPasswordMapOutput) MapIndex(k pulumi.StringInput) BindDeviceAccountPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BindDeviceAccountPassword {
		return vs[0].(map[string]*BindDeviceAccountPassword)[vs[1].(string)]
	}).(BindDeviceAccountPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindDeviceAccountPasswordInput)(nil)).Elem(), &BindDeviceAccountPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindDeviceAccountPasswordArrayInput)(nil)).Elem(), BindDeviceAccountPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindDeviceAccountPasswordMapInput)(nil)).Elem(), BindDeviceAccountPasswordMap{})
	pulumi.RegisterOutputType(BindDeviceAccountPasswordOutput{})
	pulumi.RegisterOutputType(BindDeviceAccountPasswordArrayOutput{})
	pulumi.RegisterOutputType(BindDeviceAccountPasswordMapOutput{})
}
