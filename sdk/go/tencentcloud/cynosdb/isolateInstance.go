// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cynosdb isolateInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.NewAccount(ctx, "account", &Cynosdb.AccountArgs{
// 			AccountName:        pulumi.String("terraform_test"),
// 			AccountPassword:    pulumi.String("Password@1234"),
// 			ClusterId:          pulumi.String("cynosdbmysql-bws8h88b"),
// 			Description:        pulumi.String("testx"),
// 			Host:               pulumi.String("%"),
// 			MaxUserConnections: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IsolateInstance struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// isolate, activate.
	Operate pulumi.StringOutput `pulumi:"operate"`
}

// NewIsolateInstance registers a new resource with the given unique name, arguments, and options.
func NewIsolateInstance(ctx *pulumi.Context,
	name string, args *IsolateInstanceArgs, opts ...pulumi.ResourceOption) (*IsolateInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Operate == nil {
		return nil, errors.New("invalid value for required argument 'Operate'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IsolateInstance
	err := ctx.RegisterResource("tencentcloud:Cynosdb/isolateInstance:IsolateInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIsolateInstance gets an existing IsolateInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIsolateInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IsolateInstanceState, opts ...pulumi.ResourceOption) (*IsolateInstance, error) {
	var resource IsolateInstance
	err := ctx.ReadResource("tencentcloud:Cynosdb/isolateInstance:IsolateInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IsolateInstance resources.
type isolateInstanceState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// isolate, activate.
	Operate *string `pulumi:"operate"`
}

type IsolateInstanceState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// isolate, activate.
	Operate pulumi.StringPtrInput
}

func (IsolateInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*isolateInstanceState)(nil)).Elem()
}

type isolateInstanceArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// isolate, activate.
	Operate string `pulumi:"operate"`
}

// The set of arguments for constructing a IsolateInstance resource.
type IsolateInstanceArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// isolate, activate.
	Operate pulumi.StringInput
}

func (IsolateInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*isolateInstanceArgs)(nil)).Elem()
}

type IsolateInstanceInput interface {
	pulumi.Input

	ToIsolateInstanceOutput() IsolateInstanceOutput
	ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput
}

func (*IsolateInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**IsolateInstance)(nil)).Elem()
}

func (i *IsolateInstance) ToIsolateInstanceOutput() IsolateInstanceOutput {
	return i.ToIsolateInstanceOutputWithContext(context.Background())
}

func (i *IsolateInstance) ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceOutput)
}

// IsolateInstanceArrayInput is an input type that accepts IsolateInstanceArray and IsolateInstanceArrayOutput values.
// You can construct a concrete instance of `IsolateInstanceArrayInput` via:
//
//          IsolateInstanceArray{ IsolateInstanceArgs{...} }
type IsolateInstanceArrayInput interface {
	pulumi.Input

	ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput
	ToIsolateInstanceArrayOutputWithContext(context.Context) IsolateInstanceArrayOutput
}

type IsolateInstanceArray []IsolateInstanceInput

func (IsolateInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsolateInstance)(nil)).Elem()
}

func (i IsolateInstanceArray) ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput {
	return i.ToIsolateInstanceArrayOutputWithContext(context.Background())
}

func (i IsolateInstanceArray) ToIsolateInstanceArrayOutputWithContext(ctx context.Context) IsolateInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceArrayOutput)
}

// IsolateInstanceMapInput is an input type that accepts IsolateInstanceMap and IsolateInstanceMapOutput values.
// You can construct a concrete instance of `IsolateInstanceMapInput` via:
//
//          IsolateInstanceMap{ "key": IsolateInstanceArgs{...} }
type IsolateInstanceMapInput interface {
	pulumi.Input

	ToIsolateInstanceMapOutput() IsolateInstanceMapOutput
	ToIsolateInstanceMapOutputWithContext(context.Context) IsolateInstanceMapOutput
}

type IsolateInstanceMap map[string]IsolateInstanceInput

func (IsolateInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsolateInstance)(nil)).Elem()
}

func (i IsolateInstanceMap) ToIsolateInstanceMapOutput() IsolateInstanceMapOutput {
	return i.ToIsolateInstanceMapOutputWithContext(context.Background())
}

func (i IsolateInstanceMap) ToIsolateInstanceMapOutputWithContext(ctx context.Context) IsolateInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsolateInstanceMapOutput)
}

type IsolateInstanceOutput struct{ *pulumi.OutputState }

func (IsolateInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceOutput) ToIsolateInstanceOutput() IsolateInstanceOutput {
	return o
}

func (o IsolateInstanceOutput) ToIsolateInstanceOutputWithContext(ctx context.Context) IsolateInstanceOutput {
	return o
}

// Cluster ID.
func (o IsolateInstanceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *IsolateInstance) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Instance ID.
func (o IsolateInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IsolateInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// isolate, activate.
func (o IsolateInstanceOutput) Operate() pulumi.StringOutput {
	return o.ApplyT(func(v *IsolateInstance) pulumi.StringOutput { return v.Operate }).(pulumi.StringOutput)
}

type IsolateInstanceArrayOutput struct{ *pulumi.OutputState }

func (IsolateInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceArrayOutput) ToIsolateInstanceArrayOutput() IsolateInstanceArrayOutput {
	return o
}

func (o IsolateInstanceArrayOutput) ToIsolateInstanceArrayOutputWithContext(ctx context.Context) IsolateInstanceArrayOutput {
	return o
}

func (o IsolateInstanceArrayOutput) Index(i pulumi.IntInput) IsolateInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IsolateInstance {
		return vs[0].([]*IsolateInstance)[vs[1].(int)]
	}).(IsolateInstanceOutput)
}

type IsolateInstanceMapOutput struct{ *pulumi.OutputState }

func (IsolateInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsolateInstance)(nil)).Elem()
}

func (o IsolateInstanceMapOutput) ToIsolateInstanceMapOutput() IsolateInstanceMapOutput {
	return o
}

func (o IsolateInstanceMapOutput) ToIsolateInstanceMapOutputWithContext(ctx context.Context) IsolateInstanceMapOutput {
	return o
}

func (o IsolateInstanceMapOutput) MapIndex(k pulumi.StringInput) IsolateInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IsolateInstance {
		return vs[0].(map[string]*IsolateInstance)[vs[1].(string)]
	}).(IsolateInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceInput)(nil)).Elem(), &IsolateInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceArrayInput)(nil)).Elem(), IsolateInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsolateInstanceMapInput)(nil)).Elem(), IsolateInstanceMap{})
	pulumi.RegisterOutputType(IsolateInstanceOutput{})
	pulumi.RegisterOutputType(IsolateInstanceArrayOutput{})
	pulumi.RegisterOutputType(IsolateInstanceMapOutput{})
}
