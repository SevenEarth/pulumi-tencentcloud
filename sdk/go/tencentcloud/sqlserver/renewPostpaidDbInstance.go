// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver renewPostpaidDbInstance
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
// 		_, err := Sqlserver.NewRenewPostpaidDbInstance(ctx, "renewPostpaidDbInstance", &Sqlserver.RenewPostpaidDbInstanceArgs{
// 			InstanceId: pulumi.String("mssql-i1z41iwd"),
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
// sqlserver renew_postpaid_db_instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/renewPostpaidDbInstance:RenewPostpaidDbInstance renew_postpaid_db_instance renew_postpaid_db_instance_id
// ```
type RenewPostpaidDbInstance struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewRenewPostpaidDbInstance registers a new resource with the given unique name, arguments, and options.
func NewRenewPostpaidDbInstance(ctx *pulumi.Context,
	name string, args *RenewPostpaidDbInstanceArgs, opts ...pulumi.ResourceOption) (*RenewPostpaidDbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RenewPostpaidDbInstance
	err := ctx.RegisterResource("tencentcloud:Sqlserver/renewPostpaidDbInstance:RenewPostpaidDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRenewPostpaidDbInstance gets an existing RenewPostpaidDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRenewPostpaidDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RenewPostpaidDbInstanceState, opts ...pulumi.ResourceOption) (*RenewPostpaidDbInstance, error) {
	var resource RenewPostpaidDbInstance
	err := ctx.ReadResource("tencentcloud:Sqlserver/renewPostpaidDbInstance:RenewPostpaidDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RenewPostpaidDbInstance resources.
type renewPostpaidDbInstanceState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type RenewPostpaidDbInstanceState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (RenewPostpaidDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*renewPostpaidDbInstanceState)(nil)).Elem()
}

type renewPostpaidDbInstanceArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a RenewPostpaidDbInstance resource.
type RenewPostpaidDbInstanceArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (RenewPostpaidDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*renewPostpaidDbInstanceArgs)(nil)).Elem()
}

type RenewPostpaidDbInstanceInput interface {
	pulumi.Input

	ToRenewPostpaidDbInstanceOutput() RenewPostpaidDbInstanceOutput
	ToRenewPostpaidDbInstanceOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceOutput
}

func (*RenewPostpaidDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewPostpaidDbInstance)(nil)).Elem()
}

func (i *RenewPostpaidDbInstance) ToRenewPostpaidDbInstanceOutput() RenewPostpaidDbInstanceOutput {
	return i.ToRenewPostpaidDbInstanceOutputWithContext(context.Background())
}

func (i *RenewPostpaidDbInstance) ToRenewPostpaidDbInstanceOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewPostpaidDbInstanceOutput)
}

// RenewPostpaidDbInstanceArrayInput is an input type that accepts RenewPostpaidDbInstanceArray and RenewPostpaidDbInstanceArrayOutput values.
// You can construct a concrete instance of `RenewPostpaidDbInstanceArrayInput` via:
//
//          RenewPostpaidDbInstanceArray{ RenewPostpaidDbInstanceArgs{...} }
type RenewPostpaidDbInstanceArrayInput interface {
	pulumi.Input

	ToRenewPostpaidDbInstanceArrayOutput() RenewPostpaidDbInstanceArrayOutput
	ToRenewPostpaidDbInstanceArrayOutputWithContext(context.Context) RenewPostpaidDbInstanceArrayOutput
}

type RenewPostpaidDbInstanceArray []RenewPostpaidDbInstanceInput

func (RenewPostpaidDbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewPostpaidDbInstance)(nil)).Elem()
}

func (i RenewPostpaidDbInstanceArray) ToRenewPostpaidDbInstanceArrayOutput() RenewPostpaidDbInstanceArrayOutput {
	return i.ToRenewPostpaidDbInstanceArrayOutputWithContext(context.Background())
}

func (i RenewPostpaidDbInstanceArray) ToRenewPostpaidDbInstanceArrayOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewPostpaidDbInstanceArrayOutput)
}

// RenewPostpaidDbInstanceMapInput is an input type that accepts RenewPostpaidDbInstanceMap and RenewPostpaidDbInstanceMapOutput values.
// You can construct a concrete instance of `RenewPostpaidDbInstanceMapInput` via:
//
//          RenewPostpaidDbInstanceMap{ "key": RenewPostpaidDbInstanceArgs{...} }
type RenewPostpaidDbInstanceMapInput interface {
	pulumi.Input

	ToRenewPostpaidDbInstanceMapOutput() RenewPostpaidDbInstanceMapOutput
	ToRenewPostpaidDbInstanceMapOutputWithContext(context.Context) RenewPostpaidDbInstanceMapOutput
}

type RenewPostpaidDbInstanceMap map[string]RenewPostpaidDbInstanceInput

func (RenewPostpaidDbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewPostpaidDbInstance)(nil)).Elem()
}

func (i RenewPostpaidDbInstanceMap) ToRenewPostpaidDbInstanceMapOutput() RenewPostpaidDbInstanceMapOutput {
	return i.ToRenewPostpaidDbInstanceMapOutputWithContext(context.Background())
}

func (i RenewPostpaidDbInstanceMap) ToRenewPostpaidDbInstanceMapOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RenewPostpaidDbInstanceMapOutput)
}

type RenewPostpaidDbInstanceOutput struct{ *pulumi.OutputState }

func (RenewPostpaidDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RenewPostpaidDbInstance)(nil)).Elem()
}

func (o RenewPostpaidDbInstanceOutput) ToRenewPostpaidDbInstanceOutput() RenewPostpaidDbInstanceOutput {
	return o
}

func (o RenewPostpaidDbInstanceOutput) ToRenewPostpaidDbInstanceOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceOutput {
	return o
}

// Instance ID.
func (o RenewPostpaidDbInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RenewPostpaidDbInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type RenewPostpaidDbInstanceArrayOutput struct{ *pulumi.OutputState }

func (RenewPostpaidDbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RenewPostpaidDbInstance)(nil)).Elem()
}

func (o RenewPostpaidDbInstanceArrayOutput) ToRenewPostpaidDbInstanceArrayOutput() RenewPostpaidDbInstanceArrayOutput {
	return o
}

func (o RenewPostpaidDbInstanceArrayOutput) ToRenewPostpaidDbInstanceArrayOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceArrayOutput {
	return o
}

func (o RenewPostpaidDbInstanceArrayOutput) Index(i pulumi.IntInput) RenewPostpaidDbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RenewPostpaidDbInstance {
		return vs[0].([]*RenewPostpaidDbInstance)[vs[1].(int)]
	}).(RenewPostpaidDbInstanceOutput)
}

type RenewPostpaidDbInstanceMapOutput struct{ *pulumi.OutputState }

func (RenewPostpaidDbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RenewPostpaidDbInstance)(nil)).Elem()
}

func (o RenewPostpaidDbInstanceMapOutput) ToRenewPostpaidDbInstanceMapOutput() RenewPostpaidDbInstanceMapOutput {
	return o
}

func (o RenewPostpaidDbInstanceMapOutput) ToRenewPostpaidDbInstanceMapOutputWithContext(ctx context.Context) RenewPostpaidDbInstanceMapOutput {
	return o
}

func (o RenewPostpaidDbInstanceMapOutput) MapIndex(k pulumi.StringInput) RenewPostpaidDbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RenewPostpaidDbInstance {
		return vs[0].(map[string]*RenewPostpaidDbInstance)[vs[1].(string)]
	}).(RenewPostpaidDbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RenewPostpaidDbInstanceInput)(nil)).Elem(), &RenewPostpaidDbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewPostpaidDbInstanceArrayInput)(nil)).Elem(), RenewPostpaidDbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RenewPostpaidDbInstanceMapInput)(nil)).Elem(), RenewPostpaidDbInstanceMap{})
	pulumi.RegisterOutputType(RenewPostpaidDbInstanceOutput{})
	pulumi.RegisterOutputType(RenewPostpaidDbInstanceArrayOutput{})
	pulumi.RegisterOutputType(RenewPostpaidDbInstanceMapOutput{})
}
