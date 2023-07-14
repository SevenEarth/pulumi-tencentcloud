// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mariadb restartInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mariadb.NewRestartInstance(ctx, "restartInstance", &Mariadb.RestartInstanceArgs{
// 			InstanceId: pulumi.String("tdsql-9vqvls95"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RestartInstance struct {
	pulumi.CustomResourceState

	// instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// expected restart time.
	RestartTime pulumi.StringPtrOutput `pulumi:"restartTime"`
}

// NewRestartInstance registers a new resource with the given unique name, arguments, and options.
func NewRestartInstance(ctx *pulumi.Context,
	name string, args *RestartInstanceArgs, opts ...pulumi.ResourceOption) (*RestartInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RestartInstance
	err := ctx.RegisterResource("tencentcloud:Mariadb/restartInstance:RestartInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestartInstance gets an existing RestartInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestartInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestartInstanceState, opts ...pulumi.ResourceOption) (*RestartInstance, error) {
	var resource RestartInstance
	err := ctx.ReadResource("tencentcloud:Mariadb/restartInstance:RestartInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestartInstance resources.
type restartInstanceState struct {
	// instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// expected restart time.
	RestartTime *string `pulumi:"restartTime"`
}

type RestartInstanceState struct {
	// instance ID.
	InstanceId pulumi.StringPtrInput
	// expected restart time.
	RestartTime pulumi.StringPtrInput
}

func (RestartInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*restartInstanceState)(nil)).Elem()
}

type restartInstanceArgs struct {
	// instance ID.
	InstanceId string `pulumi:"instanceId"`
	// expected restart time.
	RestartTime *string `pulumi:"restartTime"`
}

// The set of arguments for constructing a RestartInstance resource.
type RestartInstanceArgs struct {
	// instance ID.
	InstanceId pulumi.StringInput
	// expected restart time.
	RestartTime pulumi.StringPtrInput
}

func (RestartInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restartInstanceArgs)(nil)).Elem()
}

type RestartInstanceInput interface {
	pulumi.Input

	ToRestartInstanceOutput() RestartInstanceOutput
	ToRestartInstanceOutputWithContext(ctx context.Context) RestartInstanceOutput
}

func (*RestartInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartInstance)(nil)).Elem()
}

func (i *RestartInstance) ToRestartInstanceOutput() RestartInstanceOutput {
	return i.ToRestartInstanceOutputWithContext(context.Background())
}

func (i *RestartInstance) ToRestartInstanceOutputWithContext(ctx context.Context) RestartInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartInstanceOutput)
}

// RestartInstanceArrayInput is an input type that accepts RestartInstanceArray and RestartInstanceArrayOutput values.
// You can construct a concrete instance of `RestartInstanceArrayInput` via:
//
//          RestartInstanceArray{ RestartInstanceArgs{...} }
type RestartInstanceArrayInput interface {
	pulumi.Input

	ToRestartInstanceArrayOutput() RestartInstanceArrayOutput
	ToRestartInstanceArrayOutputWithContext(context.Context) RestartInstanceArrayOutput
}

type RestartInstanceArray []RestartInstanceInput

func (RestartInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartInstance)(nil)).Elem()
}

func (i RestartInstanceArray) ToRestartInstanceArrayOutput() RestartInstanceArrayOutput {
	return i.ToRestartInstanceArrayOutputWithContext(context.Background())
}

func (i RestartInstanceArray) ToRestartInstanceArrayOutputWithContext(ctx context.Context) RestartInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartInstanceArrayOutput)
}

// RestartInstanceMapInput is an input type that accepts RestartInstanceMap and RestartInstanceMapOutput values.
// You can construct a concrete instance of `RestartInstanceMapInput` via:
//
//          RestartInstanceMap{ "key": RestartInstanceArgs{...} }
type RestartInstanceMapInput interface {
	pulumi.Input

	ToRestartInstanceMapOutput() RestartInstanceMapOutput
	ToRestartInstanceMapOutputWithContext(context.Context) RestartInstanceMapOutput
}

type RestartInstanceMap map[string]RestartInstanceInput

func (RestartInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartInstance)(nil)).Elem()
}

func (i RestartInstanceMap) ToRestartInstanceMapOutput() RestartInstanceMapOutput {
	return i.ToRestartInstanceMapOutputWithContext(context.Background())
}

func (i RestartInstanceMap) ToRestartInstanceMapOutputWithContext(ctx context.Context) RestartInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartInstanceMapOutput)
}

type RestartInstanceOutput struct{ *pulumi.OutputState }

func (RestartInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartInstance)(nil)).Elem()
}

func (o RestartInstanceOutput) ToRestartInstanceOutput() RestartInstanceOutput {
	return o
}

func (o RestartInstanceOutput) ToRestartInstanceOutputWithContext(ctx context.Context) RestartInstanceOutput {
	return o
}

// instance ID.
func (o RestartInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestartInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// expected restart time.
func (o RestartInstanceOutput) RestartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestartInstance) pulumi.StringPtrOutput { return v.RestartTime }).(pulumi.StringPtrOutput)
}

type RestartInstanceArrayOutput struct{ *pulumi.OutputState }

func (RestartInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartInstance)(nil)).Elem()
}

func (o RestartInstanceArrayOutput) ToRestartInstanceArrayOutput() RestartInstanceArrayOutput {
	return o
}

func (o RestartInstanceArrayOutput) ToRestartInstanceArrayOutputWithContext(ctx context.Context) RestartInstanceArrayOutput {
	return o
}

func (o RestartInstanceArrayOutput) Index(i pulumi.IntInput) RestartInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RestartInstance {
		return vs[0].([]*RestartInstance)[vs[1].(int)]
	}).(RestartInstanceOutput)
}

type RestartInstanceMapOutput struct{ *pulumi.OutputState }

func (RestartInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartInstance)(nil)).Elem()
}

func (o RestartInstanceMapOutput) ToRestartInstanceMapOutput() RestartInstanceMapOutput {
	return o
}

func (o RestartInstanceMapOutput) ToRestartInstanceMapOutputWithContext(ctx context.Context) RestartInstanceMapOutput {
	return o
}

func (o RestartInstanceMapOutput) MapIndex(k pulumi.StringInput) RestartInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RestartInstance {
		return vs[0].(map[string]*RestartInstance)[vs[1].(string)]
	}).(RestartInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestartInstanceInput)(nil)).Elem(), &RestartInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartInstanceArrayInput)(nil)).Elem(), RestartInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartInstanceMapInput)(nil)).Elem(), RestartInstanceMap{})
	pulumi.RegisterOutputType(RestartInstanceOutput{})
	pulumi.RegisterOutputType(RestartInstanceArrayOutput{})
	pulumi.RegisterOutputType(RestartInstanceMapOutput{})
}
