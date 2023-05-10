// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dcdb account
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dcdb.NewAccount(ctx, "account", &Dcdb.AccountArgs{
// 			Description:        pulumi.String("this is a test account"),
// 			Host:               pulumi.String("127.0.0.1"),
// 			InstanceId:         pulumi.String("tdsqlshard-kkpoxvnv"),
// 			MaxUserConnections: pulumi.Int(10),
// 			Password:           pulumi.String("===password==="),
// 			ReadOnly:           pulumi.Int(0),
// 			UserName:           pulumi.String("mysql"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// dcdb account can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Dcdb/account:Account account account_id
// ```
type Account struct {
	pulumi.CustomResourceState

	// description for account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// db host.
	Host pulumi.StringOutput `pulumi:"host"`
	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// max user connections.
	MaxUserConnections pulumi.IntPtrOutput `pulumi:"maxUserConnections"`
	// password.
	Password pulumi.StringOutput `pulumi:"password"`
	// whether the account is readonly. 0 means not a readonly account.
	ReadOnly pulumi.IntPtrOutput `pulumi:"readOnly"`
	// account name.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("tencentcloud:Dcdb/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("tencentcloud:Dcdb/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// description for account.
	Description *string `pulumi:"description"`
	// db host.
	Host *string `pulumi:"host"`
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// max user connections.
	MaxUserConnections *int `pulumi:"maxUserConnections"`
	// password.
	Password *string `pulumi:"password"`
	// whether the account is readonly. 0 means not a readonly account.
	ReadOnly *int `pulumi:"readOnly"`
	// account name.
	UserName *string `pulumi:"userName"`
}

type AccountState struct {
	// description for account.
	Description pulumi.StringPtrInput
	// db host.
	Host pulumi.StringPtrInput
	// instance id.
	InstanceId pulumi.StringPtrInput
	// max user connections.
	MaxUserConnections pulumi.IntPtrInput
	// password.
	Password pulumi.StringPtrInput
	// whether the account is readonly. 0 means not a readonly account.
	ReadOnly pulumi.IntPtrInput
	// account name.
	UserName pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// description for account.
	Description *string `pulumi:"description"`
	// db host.
	Host string `pulumi:"host"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// max user connections.
	MaxUserConnections *int `pulumi:"maxUserConnections"`
	// password.
	Password string `pulumi:"password"`
	// whether the account is readonly. 0 means not a readonly account.
	ReadOnly *int `pulumi:"readOnly"`
	// account name.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// description for account.
	Description pulumi.StringPtrInput
	// db host.
	Host pulumi.StringInput
	// instance id.
	InstanceId pulumi.StringInput
	// max user connections.
	MaxUserConnections pulumi.IntPtrInput
	// password.
	Password pulumi.StringInput
	// whether the account is readonly. 0 means not a readonly account.
	ReadOnly pulumi.IntPtrInput
	// account name.
	UserName pulumi.StringInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//          AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//          AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// description for account.
func (o AccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// db host.
func (o AccountOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// instance id.
func (o AccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// max user connections.
func (o AccountOutput) MaxUserConnections() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.MaxUserConnections }).(pulumi.IntPtrOutput)
}

// password.
func (o AccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// whether the account is readonly. 0 means not a readonly account.
func (o AccountOutput) ReadOnly() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.ReadOnly }).(pulumi.IntPtrOutput)
}

// account name.
func (o AccountOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
