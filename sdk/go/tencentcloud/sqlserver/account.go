// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create SQL Server account
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sqlserver.NewAccount(ctx, "foo", &Sqlserver.AccountArgs{
// 			InstanceId: pulumi.Any(tencentcloud_sqlserver_instance.Example.Id),
// 			Password:   pulumi.String("test1233"),
// 			Remark:     pulumi.String("testt"),
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
// SQL Server account can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/account:Account foo mssql-3cdq7kx5#tf_sqlserver_account
// ```
type Account struct {
	pulumi.CustomResourceState

	// Create time of the SQL Server account.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Instance ID that the account belongs to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicate that the account is root account or not.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// Name of the SQL Server account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password of the SQL Server account.
	Password pulumi.StringOutput `pulumi:"password"`
	// Remark of the SQL Server account.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
	Status pulumi.IntOutput `pulumi:"status"`
	// Last updated time of the SQL Server account.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("tencentcloud:Sqlserver/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Sqlserver/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Create time of the SQL Server account.
	CreateTime *string `pulumi:"createTime"`
	// Instance ID that the account belongs to.
	InstanceId *string `pulumi:"instanceId"`
	// Indicate that the account is root account or not.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Name of the SQL Server account.
	Name *string `pulumi:"name"`
	// Password of the SQL Server account.
	Password *string `pulumi:"password"`
	// Remark of the SQL Server account.
	Remark *string `pulumi:"remark"`
	// Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
	Status *int `pulumi:"status"`
	// Last updated time of the SQL Server account.
	UpdateTime *string `pulumi:"updateTime"`
}

type AccountState struct {
	// Create time of the SQL Server account.
	CreateTime pulumi.StringPtrInput
	// Instance ID that the account belongs to.
	InstanceId pulumi.StringPtrInput
	// Indicate that the account is root account or not.
	IsAdmin pulumi.BoolPtrInput
	// Name of the SQL Server account.
	Name pulumi.StringPtrInput
	// Password of the SQL Server account.
	Password pulumi.StringPtrInput
	// Remark of the SQL Server account.
	Remark pulumi.StringPtrInput
	// Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
	Status pulumi.IntPtrInput
	// Last updated time of the SQL Server account.
	UpdateTime pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Instance ID that the account belongs to.
	InstanceId string `pulumi:"instanceId"`
	// Indicate that the account is root account or not.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Name of the SQL Server account.
	Name *string `pulumi:"name"`
	// Password of the SQL Server account.
	Password string `pulumi:"password"`
	// Remark of the SQL Server account.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Instance ID that the account belongs to.
	InstanceId pulumi.StringInput
	// Indicate that the account is root account or not.
	IsAdmin pulumi.BoolPtrInput
	// Name of the SQL Server account.
	Name pulumi.StringPtrInput
	// Password of the SQL Server account.
	Password pulumi.StringInput
	// Remark of the SQL Server account.
	Remark pulumi.StringPtrInput
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

// Create time of the SQL Server account.
func (o AccountOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Instance ID that the account belongs to.
func (o AccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicate that the account is root account or not.
func (o AccountOutput) IsAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.BoolPtrOutput { return v.IsAdmin }).(pulumi.BoolPtrOutput)
}

// Name of the SQL Server account.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password of the SQL Server account.
func (o AccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Remark of the SQL Server account.
func (o AccountOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
func (o AccountOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Account) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Last updated time of the SQL Server account.
func (o AccountOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
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
