// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a postgresql restartDbInstanceOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Postgresql.NewRestartDbInstanceOperation(ctx, "restartDbInstanceOperation", &Postgresql.RestartDbInstanceOperationArgs{
// 			DbInstanceId: pulumi.Any(local.Pgsql_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RestartDbInstanceOperation struct {
	pulumi.CustomResourceState

	// dbInstance ID.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
}

// NewRestartDbInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewRestartDbInstanceOperation(ctx *pulumi.Context,
	name string, args *RestartDbInstanceOperationArgs, opts ...pulumi.ResourceOption) (*RestartDbInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RestartDbInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestartDbInstanceOperation gets an existing RestartDbInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestartDbInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestartDbInstanceOperationState, opts ...pulumi.ResourceOption) (*RestartDbInstanceOperation, error) {
	var resource RestartDbInstanceOperation
	err := ctx.ReadResource("tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestartDbInstanceOperation resources.
type restartDbInstanceOperationState struct {
	// dbInstance ID.
	DbInstanceId *string `pulumi:"dbInstanceId"`
}

type RestartDbInstanceOperationState struct {
	// dbInstance ID.
	DbInstanceId pulumi.StringPtrInput
}

func (RestartDbInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*restartDbInstanceOperationState)(nil)).Elem()
}

type restartDbInstanceOperationArgs struct {
	// dbInstance ID.
	DbInstanceId string `pulumi:"dbInstanceId"`
}

// The set of arguments for constructing a RestartDbInstanceOperation resource.
type RestartDbInstanceOperationArgs struct {
	// dbInstance ID.
	DbInstanceId pulumi.StringInput
}

func (RestartDbInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restartDbInstanceOperationArgs)(nil)).Elem()
}

type RestartDbInstanceOperationInput interface {
	pulumi.Input

	ToRestartDbInstanceOperationOutput() RestartDbInstanceOperationOutput
	ToRestartDbInstanceOperationOutputWithContext(ctx context.Context) RestartDbInstanceOperationOutput
}

func (*RestartDbInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartDbInstanceOperation)(nil)).Elem()
}

func (i *RestartDbInstanceOperation) ToRestartDbInstanceOperationOutput() RestartDbInstanceOperationOutput {
	return i.ToRestartDbInstanceOperationOutputWithContext(context.Background())
}

func (i *RestartDbInstanceOperation) ToRestartDbInstanceOperationOutputWithContext(ctx context.Context) RestartDbInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstanceOperationOutput)
}

// RestartDbInstanceOperationArrayInput is an input type that accepts RestartDbInstanceOperationArray and RestartDbInstanceOperationArrayOutput values.
// You can construct a concrete instance of `RestartDbInstanceOperationArrayInput` via:
//
//          RestartDbInstanceOperationArray{ RestartDbInstanceOperationArgs{...} }
type RestartDbInstanceOperationArrayInput interface {
	pulumi.Input

	ToRestartDbInstanceOperationArrayOutput() RestartDbInstanceOperationArrayOutput
	ToRestartDbInstanceOperationArrayOutputWithContext(context.Context) RestartDbInstanceOperationArrayOutput
}

type RestartDbInstanceOperationArray []RestartDbInstanceOperationInput

func (RestartDbInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartDbInstanceOperation)(nil)).Elem()
}

func (i RestartDbInstanceOperationArray) ToRestartDbInstanceOperationArrayOutput() RestartDbInstanceOperationArrayOutput {
	return i.ToRestartDbInstanceOperationArrayOutputWithContext(context.Background())
}

func (i RestartDbInstanceOperationArray) ToRestartDbInstanceOperationArrayOutputWithContext(ctx context.Context) RestartDbInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstanceOperationArrayOutput)
}

// RestartDbInstanceOperationMapInput is an input type that accepts RestartDbInstanceOperationMap and RestartDbInstanceOperationMapOutput values.
// You can construct a concrete instance of `RestartDbInstanceOperationMapInput` via:
//
//          RestartDbInstanceOperationMap{ "key": RestartDbInstanceOperationArgs{...} }
type RestartDbInstanceOperationMapInput interface {
	pulumi.Input

	ToRestartDbInstanceOperationMapOutput() RestartDbInstanceOperationMapOutput
	ToRestartDbInstanceOperationMapOutputWithContext(context.Context) RestartDbInstanceOperationMapOutput
}

type RestartDbInstanceOperationMap map[string]RestartDbInstanceOperationInput

func (RestartDbInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartDbInstanceOperation)(nil)).Elem()
}

func (i RestartDbInstanceOperationMap) ToRestartDbInstanceOperationMapOutput() RestartDbInstanceOperationMapOutput {
	return i.ToRestartDbInstanceOperationMapOutputWithContext(context.Background())
}

func (i RestartDbInstanceOperationMap) ToRestartDbInstanceOperationMapOutputWithContext(ctx context.Context) RestartDbInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstanceOperationMapOutput)
}

type RestartDbInstanceOperationOutput struct{ *pulumi.OutputState }

func (RestartDbInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartDbInstanceOperation)(nil)).Elem()
}

func (o RestartDbInstanceOperationOutput) ToRestartDbInstanceOperationOutput() RestartDbInstanceOperationOutput {
	return o
}

func (o RestartDbInstanceOperationOutput) ToRestartDbInstanceOperationOutputWithContext(ctx context.Context) RestartDbInstanceOperationOutput {
	return o
}

// dbInstance ID.
func (o RestartDbInstanceOperationOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestartDbInstanceOperation) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

type RestartDbInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (RestartDbInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartDbInstanceOperation)(nil)).Elem()
}

func (o RestartDbInstanceOperationArrayOutput) ToRestartDbInstanceOperationArrayOutput() RestartDbInstanceOperationArrayOutput {
	return o
}

func (o RestartDbInstanceOperationArrayOutput) ToRestartDbInstanceOperationArrayOutputWithContext(ctx context.Context) RestartDbInstanceOperationArrayOutput {
	return o
}

func (o RestartDbInstanceOperationArrayOutput) Index(i pulumi.IntInput) RestartDbInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RestartDbInstanceOperation {
		return vs[0].([]*RestartDbInstanceOperation)[vs[1].(int)]
	}).(RestartDbInstanceOperationOutput)
}

type RestartDbInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (RestartDbInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartDbInstanceOperation)(nil)).Elem()
}

func (o RestartDbInstanceOperationMapOutput) ToRestartDbInstanceOperationMapOutput() RestartDbInstanceOperationMapOutput {
	return o
}

func (o RestartDbInstanceOperationMapOutput) ToRestartDbInstanceOperationMapOutputWithContext(ctx context.Context) RestartDbInstanceOperationMapOutput {
	return o
}

func (o RestartDbInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) RestartDbInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RestartDbInstanceOperation {
		return vs[0].(map[string]*RestartDbInstanceOperation)[vs[1].(string)]
	}).(RestartDbInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstanceOperationInput)(nil)).Elem(), &RestartDbInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstanceOperationArrayInput)(nil)).Elem(), RestartDbInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstanceOperationMapInput)(nil)).Elem(), RestartDbInstanceOperationMap{})
	pulumi.RegisterOutputType(RestartDbInstanceOperationOutput{})
	pulumi.RegisterOutputType(RestartDbInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(RestartDbInstanceOperationMapOutput{})
}
