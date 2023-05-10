// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to start a dts migrateJob
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.NewMigrateJobStartOperation(ctx, "start", &Dts.MigrateJobStartOperationArgs{
// 			JobId: pulumi.Any(tencentcloud_dts_migrate_job.Job.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MigrateJobStartOperation struct {
	pulumi.CustomResourceState

	// Job Id from `Dts.MigrateJob`.
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewMigrateJobStartOperation registers a new resource with the given unique name, arguments, and options.
func NewMigrateJobStartOperation(ctx *pulumi.Context,
	name string, args *MigrateJobStartOperationArgs, opts ...pulumi.ResourceOption) (*MigrateJobStartOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MigrateJobStartOperation
	err := ctx.RegisterResource("tencentcloud:Dts/migrateJobStartOperation:MigrateJobStartOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigrateJobStartOperation gets an existing MigrateJobStartOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigrateJobStartOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateJobStartOperationState, opts ...pulumi.ResourceOption) (*MigrateJobStartOperation, error) {
	var resource MigrateJobStartOperation
	err := ctx.ReadResource("tencentcloud:Dts/migrateJobStartOperation:MigrateJobStartOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MigrateJobStartOperation resources.
type migrateJobStartOperationState struct {
	// Job Id from `Dts.MigrateJob`.
	JobId *string `pulumi:"jobId"`
}

type MigrateJobStartOperationState struct {
	// Job Id from `Dts.MigrateJob`.
	JobId pulumi.StringPtrInput
}

func (MigrateJobStartOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateJobStartOperationState)(nil)).Elem()
}

type migrateJobStartOperationArgs struct {
	// Job Id from `Dts.MigrateJob`.
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a MigrateJobStartOperation resource.
type MigrateJobStartOperationArgs struct {
	// Job Id from `Dts.MigrateJob`.
	JobId pulumi.StringInput
}

func (MigrateJobStartOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateJobStartOperationArgs)(nil)).Elem()
}

type MigrateJobStartOperationInput interface {
	pulumi.Input

	ToMigrateJobStartOperationOutput() MigrateJobStartOperationOutput
	ToMigrateJobStartOperationOutputWithContext(ctx context.Context) MigrateJobStartOperationOutput
}

func (*MigrateJobStartOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateJobStartOperation)(nil)).Elem()
}

func (i *MigrateJobStartOperation) ToMigrateJobStartOperationOutput() MigrateJobStartOperationOutput {
	return i.ToMigrateJobStartOperationOutputWithContext(context.Background())
}

func (i *MigrateJobStartOperation) ToMigrateJobStartOperationOutputWithContext(ctx context.Context) MigrateJobStartOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobStartOperationOutput)
}

// MigrateJobStartOperationArrayInput is an input type that accepts MigrateJobStartOperationArray and MigrateJobStartOperationArrayOutput values.
// You can construct a concrete instance of `MigrateJobStartOperationArrayInput` via:
//
//          MigrateJobStartOperationArray{ MigrateJobStartOperationArgs{...} }
type MigrateJobStartOperationArrayInput interface {
	pulumi.Input

	ToMigrateJobStartOperationArrayOutput() MigrateJobStartOperationArrayOutput
	ToMigrateJobStartOperationArrayOutputWithContext(context.Context) MigrateJobStartOperationArrayOutput
}

type MigrateJobStartOperationArray []MigrateJobStartOperationInput

func (MigrateJobStartOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MigrateJobStartOperation)(nil)).Elem()
}

func (i MigrateJobStartOperationArray) ToMigrateJobStartOperationArrayOutput() MigrateJobStartOperationArrayOutput {
	return i.ToMigrateJobStartOperationArrayOutputWithContext(context.Background())
}

func (i MigrateJobStartOperationArray) ToMigrateJobStartOperationArrayOutputWithContext(ctx context.Context) MigrateJobStartOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobStartOperationArrayOutput)
}

// MigrateJobStartOperationMapInput is an input type that accepts MigrateJobStartOperationMap and MigrateJobStartOperationMapOutput values.
// You can construct a concrete instance of `MigrateJobStartOperationMapInput` via:
//
//          MigrateJobStartOperationMap{ "key": MigrateJobStartOperationArgs{...} }
type MigrateJobStartOperationMapInput interface {
	pulumi.Input

	ToMigrateJobStartOperationMapOutput() MigrateJobStartOperationMapOutput
	ToMigrateJobStartOperationMapOutputWithContext(context.Context) MigrateJobStartOperationMapOutput
}

type MigrateJobStartOperationMap map[string]MigrateJobStartOperationInput

func (MigrateJobStartOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MigrateJobStartOperation)(nil)).Elem()
}

func (i MigrateJobStartOperationMap) ToMigrateJobStartOperationMapOutput() MigrateJobStartOperationMapOutput {
	return i.ToMigrateJobStartOperationMapOutputWithContext(context.Background())
}

func (i MigrateJobStartOperationMap) ToMigrateJobStartOperationMapOutputWithContext(ctx context.Context) MigrateJobStartOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateJobStartOperationMapOutput)
}

type MigrateJobStartOperationOutput struct{ *pulumi.OutputState }

func (MigrateJobStartOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateJobStartOperation)(nil)).Elem()
}

func (o MigrateJobStartOperationOutput) ToMigrateJobStartOperationOutput() MigrateJobStartOperationOutput {
	return o
}

func (o MigrateJobStartOperationOutput) ToMigrateJobStartOperationOutputWithContext(ctx context.Context) MigrateJobStartOperationOutput {
	return o
}

// Job Id from `Dts.MigrateJob`.
func (o MigrateJobStartOperationOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateJobStartOperation) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type MigrateJobStartOperationArrayOutput struct{ *pulumi.OutputState }

func (MigrateJobStartOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MigrateJobStartOperation)(nil)).Elem()
}

func (o MigrateJobStartOperationArrayOutput) ToMigrateJobStartOperationArrayOutput() MigrateJobStartOperationArrayOutput {
	return o
}

func (o MigrateJobStartOperationArrayOutput) ToMigrateJobStartOperationArrayOutputWithContext(ctx context.Context) MigrateJobStartOperationArrayOutput {
	return o
}

func (o MigrateJobStartOperationArrayOutput) Index(i pulumi.IntInput) MigrateJobStartOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MigrateJobStartOperation {
		return vs[0].([]*MigrateJobStartOperation)[vs[1].(int)]
	}).(MigrateJobStartOperationOutput)
}

type MigrateJobStartOperationMapOutput struct{ *pulumi.OutputState }

func (MigrateJobStartOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MigrateJobStartOperation)(nil)).Elem()
}

func (o MigrateJobStartOperationMapOutput) ToMigrateJobStartOperationMapOutput() MigrateJobStartOperationMapOutput {
	return o
}

func (o MigrateJobStartOperationMapOutput) ToMigrateJobStartOperationMapOutputWithContext(ctx context.Context) MigrateJobStartOperationMapOutput {
	return o
}

func (o MigrateJobStartOperationMapOutput) MapIndex(k pulumi.StringInput) MigrateJobStartOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MigrateJobStartOperation {
		return vs[0].(map[string]*MigrateJobStartOperation)[vs[1].(string)]
	}).(MigrateJobStartOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobStartOperationInput)(nil)).Elem(), &MigrateJobStartOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobStartOperationArrayInput)(nil)).Elem(), MigrateJobStartOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateJobStartOperationMapInput)(nil)).Elem(), MigrateJobStartOperationMap{})
	pulumi.RegisterOutputType(MigrateJobStartOperationOutput{})
	pulumi.RegisterOutputType(MigrateJobStartOperationArrayOutput{})
	pulumi.RegisterOutputType(MigrateJobStartOperationMapOutput{})
}
