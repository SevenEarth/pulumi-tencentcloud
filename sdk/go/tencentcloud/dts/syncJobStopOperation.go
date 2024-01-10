// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dts syncJobStopOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dts.NewSyncJobStopOperation(ctx, "syncJobStopOperation", &Dts.SyncJobStopOperationArgs{
//				JobId: pulumi.String("sync-werwfs23"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SyncJobStopOperation struct {
	pulumi.CustomResourceState

	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewSyncJobStopOperation registers a new resource with the given unique name, arguments, and options.
func NewSyncJobStopOperation(ctx *pulumi.Context,
	name string, args *SyncJobStopOperationArgs, opts ...pulumi.ResourceOption) (*SyncJobStopOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SyncJobStopOperation
	err := ctx.RegisterResource("tencentcloud:Dts/syncJobStopOperation:SyncJobStopOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncJobStopOperation gets an existing SyncJobStopOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncJobStopOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncJobStopOperationState, opts ...pulumi.ResourceOption) (*SyncJobStopOperation, error) {
	var resource SyncJobStopOperation
	err := ctx.ReadResource("tencentcloud:Dts/syncJobStopOperation:SyncJobStopOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncJobStopOperation resources.
type syncJobStopOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId *string `pulumi:"jobId"`
}

type SyncJobStopOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringPtrInput
}

func (SyncJobStopOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobStopOperationState)(nil)).Elem()
}

type syncJobStopOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a SyncJobStopOperation resource.
type SyncJobStopOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringInput
}

func (SyncJobStopOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobStopOperationArgs)(nil)).Elem()
}

type SyncJobStopOperationInput interface {
	pulumi.Input

	ToSyncJobStopOperationOutput() SyncJobStopOperationOutput
	ToSyncJobStopOperationOutputWithContext(ctx context.Context) SyncJobStopOperationOutput
}

func (*SyncJobStopOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobStopOperation)(nil)).Elem()
}

func (i *SyncJobStopOperation) ToSyncJobStopOperationOutput() SyncJobStopOperationOutput {
	return i.ToSyncJobStopOperationOutputWithContext(context.Background())
}

func (i *SyncJobStopOperation) ToSyncJobStopOperationOutputWithContext(ctx context.Context) SyncJobStopOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStopOperationOutput)
}

// SyncJobStopOperationArrayInput is an input type that accepts SyncJobStopOperationArray and SyncJobStopOperationArrayOutput values.
// You can construct a concrete instance of `SyncJobStopOperationArrayInput` via:
//
//	SyncJobStopOperationArray{ SyncJobStopOperationArgs{...} }
type SyncJobStopOperationArrayInput interface {
	pulumi.Input

	ToSyncJobStopOperationArrayOutput() SyncJobStopOperationArrayOutput
	ToSyncJobStopOperationArrayOutputWithContext(context.Context) SyncJobStopOperationArrayOutput
}

type SyncJobStopOperationArray []SyncJobStopOperationInput

func (SyncJobStopOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobStopOperation)(nil)).Elem()
}

func (i SyncJobStopOperationArray) ToSyncJobStopOperationArrayOutput() SyncJobStopOperationArrayOutput {
	return i.ToSyncJobStopOperationArrayOutputWithContext(context.Background())
}

func (i SyncJobStopOperationArray) ToSyncJobStopOperationArrayOutputWithContext(ctx context.Context) SyncJobStopOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStopOperationArrayOutput)
}

// SyncJobStopOperationMapInput is an input type that accepts SyncJobStopOperationMap and SyncJobStopOperationMapOutput values.
// You can construct a concrete instance of `SyncJobStopOperationMapInput` via:
//
//	SyncJobStopOperationMap{ "key": SyncJobStopOperationArgs{...} }
type SyncJobStopOperationMapInput interface {
	pulumi.Input

	ToSyncJobStopOperationMapOutput() SyncJobStopOperationMapOutput
	ToSyncJobStopOperationMapOutputWithContext(context.Context) SyncJobStopOperationMapOutput
}

type SyncJobStopOperationMap map[string]SyncJobStopOperationInput

func (SyncJobStopOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobStopOperation)(nil)).Elem()
}

func (i SyncJobStopOperationMap) ToSyncJobStopOperationMapOutput() SyncJobStopOperationMapOutput {
	return i.ToSyncJobStopOperationMapOutputWithContext(context.Background())
}

func (i SyncJobStopOperationMap) ToSyncJobStopOperationMapOutputWithContext(ctx context.Context) SyncJobStopOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStopOperationMapOutput)
}

type SyncJobStopOperationOutput struct{ *pulumi.OutputState }

func (SyncJobStopOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobStopOperation)(nil)).Elem()
}

func (o SyncJobStopOperationOutput) ToSyncJobStopOperationOutput() SyncJobStopOperationOutput {
	return o
}

func (o SyncJobStopOperationOutput) ToSyncJobStopOperationOutputWithContext(ctx context.Context) SyncJobStopOperationOutput {
	return o
}

// Synchronization instance id (i.e. identifies a synchronization job).
func (o SyncJobStopOperationOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncJobStopOperation) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type SyncJobStopOperationArrayOutput struct{ *pulumi.OutputState }

func (SyncJobStopOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobStopOperation)(nil)).Elem()
}

func (o SyncJobStopOperationArrayOutput) ToSyncJobStopOperationArrayOutput() SyncJobStopOperationArrayOutput {
	return o
}

func (o SyncJobStopOperationArrayOutput) ToSyncJobStopOperationArrayOutputWithContext(ctx context.Context) SyncJobStopOperationArrayOutput {
	return o
}

func (o SyncJobStopOperationArrayOutput) Index(i pulumi.IntInput) SyncJobStopOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncJobStopOperation {
		return vs[0].([]*SyncJobStopOperation)[vs[1].(int)]
	}).(SyncJobStopOperationOutput)
}

type SyncJobStopOperationMapOutput struct{ *pulumi.OutputState }

func (SyncJobStopOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobStopOperation)(nil)).Elem()
}

func (o SyncJobStopOperationMapOutput) ToSyncJobStopOperationMapOutput() SyncJobStopOperationMapOutput {
	return o
}

func (o SyncJobStopOperationMapOutput) ToSyncJobStopOperationMapOutputWithContext(ctx context.Context) SyncJobStopOperationMapOutput {
	return o
}

func (o SyncJobStopOperationMapOutput) MapIndex(k pulumi.StringInput) SyncJobStopOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncJobStopOperation {
		return vs[0].(map[string]*SyncJobStopOperation)[vs[1].(string)]
	}).(SyncJobStopOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStopOperationInput)(nil)).Elem(), &SyncJobStopOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStopOperationArrayInput)(nil)).Elem(), SyncJobStopOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStopOperationMapInput)(nil)).Elem(), SyncJobStopOperationMap{})
	pulumi.RegisterOutputType(SyncJobStopOperationOutput{})
	pulumi.RegisterOutputType(SyncJobStopOperationArrayOutput{})
	pulumi.RegisterOutputType(SyncJobStopOperationMapOutput{})
}
