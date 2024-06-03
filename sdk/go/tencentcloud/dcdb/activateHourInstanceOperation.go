// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a dcdb activateHourInstanceOperation
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dcdb.NewActivateHourInstanceOperation(ctx, "activateHourInstanceOperation", &Dcdb.ActivateHourInstanceOperationArgs{
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
// <!--End PulumiCodeChooser -->
type ActivateHourInstanceOperation struct {
	pulumi.CustomResourceState

	// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewActivateHourInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewActivateHourInstanceOperation(ctx *pulumi.Context,
	name string, args *ActivateHourInstanceOperationArgs, opts ...pulumi.ResourceOption) (*ActivateHourInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActivateHourInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Dcdb/activateHourInstanceOperation:ActivateHourInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivateHourInstanceOperation gets an existing ActivateHourInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivateHourInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivateHourInstanceOperationState, opts ...pulumi.ResourceOption) (*ActivateHourInstanceOperation, error) {
	var resource ActivateHourInstanceOperation
	err := ctx.ReadResource("tencentcloud:Dcdb/activateHourInstanceOperation:ActivateHourInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActivateHourInstanceOperation resources.
type activateHourInstanceOperationState struct {
	// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `pulumi:"instanceId"`
}

type ActivateHourInstanceOperationState struct {
	// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId pulumi.StringPtrInput
}

func (ActivateHourInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*activateHourInstanceOperationState)(nil)).Elem()
}

type activateHourInstanceOperationArgs struct {
	// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a ActivateHourInstanceOperation resource.
type ActivateHourInstanceOperationArgs struct {
	// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId pulumi.StringInput
}

func (ActivateHourInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activateHourInstanceOperationArgs)(nil)).Elem()
}

type ActivateHourInstanceOperationInput interface {
	pulumi.Input

	ToActivateHourInstanceOperationOutput() ActivateHourInstanceOperationOutput
	ToActivateHourInstanceOperationOutputWithContext(ctx context.Context) ActivateHourInstanceOperationOutput
}

func (*ActivateHourInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivateHourInstanceOperation)(nil)).Elem()
}

func (i *ActivateHourInstanceOperation) ToActivateHourInstanceOperationOutput() ActivateHourInstanceOperationOutput {
	return i.ToActivateHourInstanceOperationOutputWithContext(context.Background())
}

func (i *ActivateHourInstanceOperation) ToActivateHourInstanceOperationOutputWithContext(ctx context.Context) ActivateHourInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivateHourInstanceOperationOutput)
}

// ActivateHourInstanceOperationArrayInput is an input type that accepts ActivateHourInstanceOperationArray and ActivateHourInstanceOperationArrayOutput values.
// You can construct a concrete instance of `ActivateHourInstanceOperationArrayInput` via:
//
//	ActivateHourInstanceOperationArray{ ActivateHourInstanceOperationArgs{...} }
type ActivateHourInstanceOperationArrayInput interface {
	pulumi.Input

	ToActivateHourInstanceOperationArrayOutput() ActivateHourInstanceOperationArrayOutput
	ToActivateHourInstanceOperationArrayOutputWithContext(context.Context) ActivateHourInstanceOperationArrayOutput
}

type ActivateHourInstanceOperationArray []ActivateHourInstanceOperationInput

func (ActivateHourInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActivateHourInstanceOperation)(nil)).Elem()
}

func (i ActivateHourInstanceOperationArray) ToActivateHourInstanceOperationArrayOutput() ActivateHourInstanceOperationArrayOutput {
	return i.ToActivateHourInstanceOperationArrayOutputWithContext(context.Background())
}

func (i ActivateHourInstanceOperationArray) ToActivateHourInstanceOperationArrayOutputWithContext(ctx context.Context) ActivateHourInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivateHourInstanceOperationArrayOutput)
}

// ActivateHourInstanceOperationMapInput is an input type that accepts ActivateHourInstanceOperationMap and ActivateHourInstanceOperationMapOutput values.
// You can construct a concrete instance of `ActivateHourInstanceOperationMapInput` via:
//
//	ActivateHourInstanceOperationMap{ "key": ActivateHourInstanceOperationArgs{...} }
type ActivateHourInstanceOperationMapInput interface {
	pulumi.Input

	ToActivateHourInstanceOperationMapOutput() ActivateHourInstanceOperationMapOutput
	ToActivateHourInstanceOperationMapOutputWithContext(context.Context) ActivateHourInstanceOperationMapOutput
}

type ActivateHourInstanceOperationMap map[string]ActivateHourInstanceOperationInput

func (ActivateHourInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActivateHourInstanceOperation)(nil)).Elem()
}

func (i ActivateHourInstanceOperationMap) ToActivateHourInstanceOperationMapOutput() ActivateHourInstanceOperationMapOutput {
	return i.ToActivateHourInstanceOperationMapOutputWithContext(context.Background())
}

func (i ActivateHourInstanceOperationMap) ToActivateHourInstanceOperationMapOutputWithContext(ctx context.Context) ActivateHourInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivateHourInstanceOperationMapOutput)
}

type ActivateHourInstanceOperationOutput struct{ *pulumi.OutputState }

func (ActivateHourInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivateHourInstanceOperation)(nil)).Elem()
}

func (o ActivateHourInstanceOperationOutput) ToActivateHourInstanceOperationOutput() ActivateHourInstanceOperationOutput {
	return o
}

func (o ActivateHourInstanceOperationOutput) ToActivateHourInstanceOperationOutputWithContext(ctx context.Context) ActivateHourInstanceOperationOutput {
	return o
}

// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
func (o ActivateHourInstanceOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivateHourInstanceOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type ActivateHourInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (ActivateHourInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActivateHourInstanceOperation)(nil)).Elem()
}

func (o ActivateHourInstanceOperationArrayOutput) ToActivateHourInstanceOperationArrayOutput() ActivateHourInstanceOperationArrayOutput {
	return o
}

func (o ActivateHourInstanceOperationArrayOutput) ToActivateHourInstanceOperationArrayOutputWithContext(ctx context.Context) ActivateHourInstanceOperationArrayOutput {
	return o
}

func (o ActivateHourInstanceOperationArrayOutput) Index(i pulumi.IntInput) ActivateHourInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActivateHourInstanceOperation {
		return vs[0].([]*ActivateHourInstanceOperation)[vs[1].(int)]
	}).(ActivateHourInstanceOperationOutput)
}

type ActivateHourInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (ActivateHourInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActivateHourInstanceOperation)(nil)).Elem()
}

func (o ActivateHourInstanceOperationMapOutput) ToActivateHourInstanceOperationMapOutput() ActivateHourInstanceOperationMapOutput {
	return o
}

func (o ActivateHourInstanceOperationMapOutput) ToActivateHourInstanceOperationMapOutputWithContext(ctx context.Context) ActivateHourInstanceOperationMapOutput {
	return o
}

func (o ActivateHourInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) ActivateHourInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActivateHourInstanceOperation {
		return vs[0].(map[string]*ActivateHourInstanceOperation)[vs[1].(string)]
	}).(ActivateHourInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActivateHourInstanceOperationInput)(nil)).Elem(), &ActivateHourInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivateHourInstanceOperationArrayInput)(nil)).Elem(), ActivateHourInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivateHourInstanceOperationMapInput)(nil)).Elem(), ActivateHourInstanceOperationMap{})
	pulumi.RegisterOutputType(ActivateHourInstanceOperationOutput{})
	pulumi.RegisterOutputType(ActivateHourInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(ActivateHourInstanceOperationMapOutput{})
}
