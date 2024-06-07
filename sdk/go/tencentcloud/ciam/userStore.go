// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ciam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a ciam user store
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ciam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ciam.NewUserStore(ctx, "userStore", &Ciam.UserStoreArgs{
//				UserPoolDesc: pulumi.String("for terraform test 123"),
//				UserPoolLogo: pulumi.String("https://ciam-prd-1302490086.cos.ap-guangzhou.myqcloud.com/temporary/92630252a2c5422d9663db5feafd619b.png"),
//				UserPoolName: pulumi.String("tf_user_store"),
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
//
// ## Import
//
// ciam user_store can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Ciam/userStore:UserStore user_store userStoreId
// ```
type UserStore struct {
	pulumi.CustomResourceState

	// User Store Description.
	UserPoolDesc pulumi.StringPtrOutput `pulumi:"userPoolDesc"`
	// User Store Logo.
	UserPoolLogo pulumi.StringPtrOutput `pulumi:"userPoolLogo"`
	// User Store Name.
	UserPoolName pulumi.StringOutput `pulumi:"userPoolName"`
}

// NewUserStore registers a new resource with the given unique name, arguments, and options.
func NewUserStore(ctx *pulumi.Context,
	name string, args *UserStoreArgs, opts ...pulumi.ResourceOption) (*UserStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserPoolName == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserStore
	err := ctx.RegisterResource("tencentcloud:Ciam/userStore:UserStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserStore gets an existing UserStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserStoreState, opts ...pulumi.ResourceOption) (*UserStore, error) {
	var resource UserStore
	err := ctx.ReadResource("tencentcloud:Ciam/userStore:UserStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserStore resources.
type userStoreState struct {
	// User Store Description.
	UserPoolDesc *string `pulumi:"userPoolDesc"`
	// User Store Logo.
	UserPoolLogo *string `pulumi:"userPoolLogo"`
	// User Store Name.
	UserPoolName *string `pulumi:"userPoolName"`
}

type UserStoreState struct {
	// User Store Description.
	UserPoolDesc pulumi.StringPtrInput
	// User Store Logo.
	UserPoolLogo pulumi.StringPtrInput
	// User Store Name.
	UserPoolName pulumi.StringPtrInput
}

func (UserStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*userStoreState)(nil)).Elem()
}

type userStoreArgs struct {
	// User Store Description.
	UserPoolDesc *string `pulumi:"userPoolDesc"`
	// User Store Logo.
	UserPoolLogo *string `pulumi:"userPoolLogo"`
	// User Store Name.
	UserPoolName string `pulumi:"userPoolName"`
}

// The set of arguments for constructing a UserStore resource.
type UserStoreArgs struct {
	// User Store Description.
	UserPoolDesc pulumi.StringPtrInput
	// User Store Logo.
	UserPoolLogo pulumi.StringPtrInput
	// User Store Name.
	UserPoolName pulumi.StringInput
}

func (UserStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userStoreArgs)(nil)).Elem()
}

type UserStoreInput interface {
	pulumi.Input

	ToUserStoreOutput() UserStoreOutput
	ToUserStoreOutputWithContext(ctx context.Context) UserStoreOutput
}

func (*UserStore) ElementType() reflect.Type {
	return reflect.TypeOf((**UserStore)(nil)).Elem()
}

func (i *UserStore) ToUserStoreOutput() UserStoreOutput {
	return i.ToUserStoreOutputWithContext(context.Background())
}

func (i *UserStore) ToUserStoreOutputWithContext(ctx context.Context) UserStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserStoreOutput)
}

// UserStoreArrayInput is an input type that accepts UserStoreArray and UserStoreArrayOutput values.
// You can construct a concrete instance of `UserStoreArrayInput` via:
//
//	UserStoreArray{ UserStoreArgs{...} }
type UserStoreArrayInput interface {
	pulumi.Input

	ToUserStoreArrayOutput() UserStoreArrayOutput
	ToUserStoreArrayOutputWithContext(context.Context) UserStoreArrayOutput
}

type UserStoreArray []UserStoreInput

func (UserStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserStore)(nil)).Elem()
}

func (i UserStoreArray) ToUserStoreArrayOutput() UserStoreArrayOutput {
	return i.ToUserStoreArrayOutputWithContext(context.Background())
}

func (i UserStoreArray) ToUserStoreArrayOutputWithContext(ctx context.Context) UserStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserStoreArrayOutput)
}

// UserStoreMapInput is an input type that accepts UserStoreMap and UserStoreMapOutput values.
// You can construct a concrete instance of `UserStoreMapInput` via:
//
//	UserStoreMap{ "key": UserStoreArgs{...} }
type UserStoreMapInput interface {
	pulumi.Input

	ToUserStoreMapOutput() UserStoreMapOutput
	ToUserStoreMapOutputWithContext(context.Context) UserStoreMapOutput
}

type UserStoreMap map[string]UserStoreInput

func (UserStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserStore)(nil)).Elem()
}

func (i UserStoreMap) ToUserStoreMapOutput() UserStoreMapOutput {
	return i.ToUserStoreMapOutputWithContext(context.Background())
}

func (i UserStoreMap) ToUserStoreMapOutputWithContext(ctx context.Context) UserStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserStoreMapOutput)
}

type UserStoreOutput struct{ *pulumi.OutputState }

func (UserStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserStore)(nil)).Elem()
}

func (o UserStoreOutput) ToUserStoreOutput() UserStoreOutput {
	return o
}

func (o UserStoreOutput) ToUserStoreOutputWithContext(ctx context.Context) UserStoreOutput {
	return o
}

// User Store Description.
func (o UserStoreOutput) UserPoolDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserStore) pulumi.StringPtrOutput { return v.UserPoolDesc }).(pulumi.StringPtrOutput)
}

// User Store Logo.
func (o UserStoreOutput) UserPoolLogo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserStore) pulumi.StringPtrOutput { return v.UserPoolLogo }).(pulumi.StringPtrOutput)
}

// User Store Name.
func (o UserStoreOutput) UserPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserStore) pulumi.StringOutput { return v.UserPoolName }).(pulumi.StringOutput)
}

type UserStoreArrayOutput struct{ *pulumi.OutputState }

func (UserStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserStore)(nil)).Elem()
}

func (o UserStoreArrayOutput) ToUserStoreArrayOutput() UserStoreArrayOutput {
	return o
}

func (o UserStoreArrayOutput) ToUserStoreArrayOutputWithContext(ctx context.Context) UserStoreArrayOutput {
	return o
}

func (o UserStoreArrayOutput) Index(i pulumi.IntInput) UserStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserStore {
		return vs[0].([]*UserStore)[vs[1].(int)]
	}).(UserStoreOutput)
}

type UserStoreMapOutput struct{ *pulumi.OutputState }

func (UserStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserStore)(nil)).Elem()
}

func (o UserStoreMapOutput) ToUserStoreMapOutput() UserStoreMapOutput {
	return o
}

func (o UserStoreMapOutput) ToUserStoreMapOutputWithContext(ctx context.Context) UserStoreMapOutput {
	return o
}

func (o UserStoreMapOutput) MapIndex(k pulumi.StringInput) UserStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserStore {
		return vs[0].(map[string]*UserStore)[vs[1].(string)]
	}).(UserStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserStoreInput)(nil)).Elem(), &UserStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserStoreArrayInput)(nil)).Elem(), UserStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserStoreMapInput)(nil)).Elem(), UserStoreMap{})
	pulumi.RegisterOutputType(UserStoreOutput{})
	pulumi.RegisterOutputType(UserStoreArrayOutput{})
	pulumi.RegisterOutputType(UserStoreMapOutput{})
}
