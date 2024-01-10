// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dcdb flushBinlogOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dcdb.NewFlushBinlogOperation(ctx, "flushOperation", &Dcdb.FlushBinlogOperationArgs{
//				InstanceId: pulumi.Any(local.Dcdb_id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FlushBinlogOperation struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewFlushBinlogOperation registers a new resource with the given unique name, arguments, and options.
func NewFlushBinlogOperation(ctx *pulumi.Context,
	name string, args *FlushBinlogOperationArgs, opts ...pulumi.ResourceOption) (*FlushBinlogOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FlushBinlogOperation
	err := ctx.RegisterResource("tencentcloud:Dcdb/flushBinlogOperation:FlushBinlogOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlushBinlogOperation gets an existing FlushBinlogOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlushBinlogOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlushBinlogOperationState, opts ...pulumi.ResourceOption) (*FlushBinlogOperation, error) {
	var resource FlushBinlogOperation
	err := ctx.ReadResource("tencentcloud:Dcdb/flushBinlogOperation:FlushBinlogOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlushBinlogOperation resources.
type flushBinlogOperationState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type FlushBinlogOperationState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (FlushBinlogOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*flushBinlogOperationState)(nil)).Elem()
}

type flushBinlogOperationArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a FlushBinlogOperation resource.
type FlushBinlogOperationArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (FlushBinlogOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flushBinlogOperationArgs)(nil)).Elem()
}

type FlushBinlogOperationInput interface {
	pulumi.Input

	ToFlushBinlogOperationOutput() FlushBinlogOperationOutput
	ToFlushBinlogOperationOutputWithContext(ctx context.Context) FlushBinlogOperationOutput
}

func (*FlushBinlogOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**FlushBinlogOperation)(nil)).Elem()
}

func (i *FlushBinlogOperation) ToFlushBinlogOperationOutput() FlushBinlogOperationOutput {
	return i.ToFlushBinlogOperationOutputWithContext(context.Background())
}

func (i *FlushBinlogOperation) ToFlushBinlogOperationOutputWithContext(ctx context.Context) FlushBinlogOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlushBinlogOperationOutput)
}

// FlushBinlogOperationArrayInput is an input type that accepts FlushBinlogOperationArray and FlushBinlogOperationArrayOutput values.
// You can construct a concrete instance of `FlushBinlogOperationArrayInput` via:
//
//	FlushBinlogOperationArray{ FlushBinlogOperationArgs{...} }
type FlushBinlogOperationArrayInput interface {
	pulumi.Input

	ToFlushBinlogOperationArrayOutput() FlushBinlogOperationArrayOutput
	ToFlushBinlogOperationArrayOutputWithContext(context.Context) FlushBinlogOperationArrayOutput
}

type FlushBinlogOperationArray []FlushBinlogOperationInput

func (FlushBinlogOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlushBinlogOperation)(nil)).Elem()
}

func (i FlushBinlogOperationArray) ToFlushBinlogOperationArrayOutput() FlushBinlogOperationArrayOutput {
	return i.ToFlushBinlogOperationArrayOutputWithContext(context.Background())
}

func (i FlushBinlogOperationArray) ToFlushBinlogOperationArrayOutputWithContext(ctx context.Context) FlushBinlogOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlushBinlogOperationArrayOutput)
}

// FlushBinlogOperationMapInput is an input type that accepts FlushBinlogOperationMap and FlushBinlogOperationMapOutput values.
// You can construct a concrete instance of `FlushBinlogOperationMapInput` via:
//
//	FlushBinlogOperationMap{ "key": FlushBinlogOperationArgs{...} }
type FlushBinlogOperationMapInput interface {
	pulumi.Input

	ToFlushBinlogOperationMapOutput() FlushBinlogOperationMapOutput
	ToFlushBinlogOperationMapOutputWithContext(context.Context) FlushBinlogOperationMapOutput
}

type FlushBinlogOperationMap map[string]FlushBinlogOperationInput

func (FlushBinlogOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlushBinlogOperation)(nil)).Elem()
}

func (i FlushBinlogOperationMap) ToFlushBinlogOperationMapOutput() FlushBinlogOperationMapOutput {
	return i.ToFlushBinlogOperationMapOutputWithContext(context.Background())
}

func (i FlushBinlogOperationMap) ToFlushBinlogOperationMapOutputWithContext(ctx context.Context) FlushBinlogOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlushBinlogOperationMapOutput)
}

type FlushBinlogOperationOutput struct{ *pulumi.OutputState }

func (FlushBinlogOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlushBinlogOperation)(nil)).Elem()
}

func (o FlushBinlogOperationOutput) ToFlushBinlogOperationOutput() FlushBinlogOperationOutput {
	return o
}

func (o FlushBinlogOperationOutput) ToFlushBinlogOperationOutputWithContext(ctx context.Context) FlushBinlogOperationOutput {
	return o
}

// Instance ID.
func (o FlushBinlogOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlushBinlogOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type FlushBinlogOperationArrayOutput struct{ *pulumi.OutputState }

func (FlushBinlogOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlushBinlogOperation)(nil)).Elem()
}

func (o FlushBinlogOperationArrayOutput) ToFlushBinlogOperationArrayOutput() FlushBinlogOperationArrayOutput {
	return o
}

func (o FlushBinlogOperationArrayOutput) ToFlushBinlogOperationArrayOutputWithContext(ctx context.Context) FlushBinlogOperationArrayOutput {
	return o
}

func (o FlushBinlogOperationArrayOutput) Index(i pulumi.IntInput) FlushBinlogOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlushBinlogOperation {
		return vs[0].([]*FlushBinlogOperation)[vs[1].(int)]
	}).(FlushBinlogOperationOutput)
}

type FlushBinlogOperationMapOutput struct{ *pulumi.OutputState }

func (FlushBinlogOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlushBinlogOperation)(nil)).Elem()
}

func (o FlushBinlogOperationMapOutput) ToFlushBinlogOperationMapOutput() FlushBinlogOperationMapOutput {
	return o
}

func (o FlushBinlogOperationMapOutput) ToFlushBinlogOperationMapOutputWithContext(ctx context.Context) FlushBinlogOperationMapOutput {
	return o
}

func (o FlushBinlogOperationMapOutput) MapIndex(k pulumi.StringInput) FlushBinlogOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlushBinlogOperation {
		return vs[0].(map[string]*FlushBinlogOperation)[vs[1].(string)]
	}).(FlushBinlogOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlushBinlogOperationInput)(nil)).Elem(), &FlushBinlogOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlushBinlogOperationArrayInput)(nil)).Elem(), FlushBinlogOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlushBinlogOperationMapInput)(nil)).Elem(), FlushBinlogOperationMap{})
	pulumi.RegisterOutputType(FlushBinlogOperationOutput{})
	pulumi.RegisterOutputType(FlushBinlogOperationArrayOutput{})
	pulumi.RegisterOutputType(FlushBinlogOperationMapOutput{})
}
