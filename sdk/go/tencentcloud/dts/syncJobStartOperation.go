// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dts syncJobStartOperation
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
//			_, err := Dts.NewSyncJobStartOperation(ctx, "syncJobStartOperation", &Dts.SyncJobStartOperationArgs{
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
type SyncJobStartOperation struct {
	pulumi.CustomResourceState

	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringOutput `pulumi:"jobId"`
}

// NewSyncJobStartOperation registers a new resource with the given unique name, arguments, and options.
func NewSyncJobStartOperation(ctx *pulumi.Context,
	name string, args *SyncJobStartOperationArgs, opts ...pulumi.ResourceOption) (*SyncJobStartOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SyncJobStartOperation
	err := ctx.RegisterResource("tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncJobStartOperation gets an existing SyncJobStartOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncJobStartOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncJobStartOperationState, opts ...pulumi.ResourceOption) (*SyncJobStartOperation, error) {
	var resource SyncJobStartOperation
	err := ctx.ReadResource("tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncJobStartOperation resources.
type syncJobStartOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId *string `pulumi:"jobId"`
}

type SyncJobStartOperationState struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringPtrInput
}

func (SyncJobStartOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobStartOperationState)(nil)).Elem()
}

type syncJobStartOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId string `pulumi:"jobId"`
}

// The set of arguments for constructing a SyncJobStartOperation resource.
type SyncJobStartOperationArgs struct {
	// Synchronization instance id (i.e. identifies a synchronization job).
	JobId pulumi.StringInput
}

func (SyncJobStartOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncJobStartOperationArgs)(nil)).Elem()
}

type SyncJobStartOperationInput interface {
	pulumi.Input

	ToSyncJobStartOperationOutput() SyncJobStartOperationOutput
	ToSyncJobStartOperationOutputWithContext(ctx context.Context) SyncJobStartOperationOutput
}

func (*SyncJobStartOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobStartOperation)(nil)).Elem()
}

func (i *SyncJobStartOperation) ToSyncJobStartOperationOutput() SyncJobStartOperationOutput {
	return i.ToSyncJobStartOperationOutputWithContext(context.Background())
}

func (i *SyncJobStartOperation) ToSyncJobStartOperationOutputWithContext(ctx context.Context) SyncJobStartOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStartOperationOutput)
}

// SyncJobStartOperationArrayInput is an input type that accepts SyncJobStartOperationArray and SyncJobStartOperationArrayOutput values.
// You can construct a concrete instance of `SyncJobStartOperationArrayInput` via:
//
//	SyncJobStartOperationArray{ SyncJobStartOperationArgs{...} }
type SyncJobStartOperationArrayInput interface {
	pulumi.Input

	ToSyncJobStartOperationArrayOutput() SyncJobStartOperationArrayOutput
	ToSyncJobStartOperationArrayOutputWithContext(context.Context) SyncJobStartOperationArrayOutput
}

type SyncJobStartOperationArray []SyncJobStartOperationInput

func (SyncJobStartOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobStartOperation)(nil)).Elem()
}

func (i SyncJobStartOperationArray) ToSyncJobStartOperationArrayOutput() SyncJobStartOperationArrayOutput {
	return i.ToSyncJobStartOperationArrayOutputWithContext(context.Background())
}

func (i SyncJobStartOperationArray) ToSyncJobStartOperationArrayOutputWithContext(ctx context.Context) SyncJobStartOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStartOperationArrayOutput)
}

// SyncJobStartOperationMapInput is an input type that accepts SyncJobStartOperationMap and SyncJobStartOperationMapOutput values.
// You can construct a concrete instance of `SyncJobStartOperationMapInput` via:
//
//	SyncJobStartOperationMap{ "key": SyncJobStartOperationArgs{...} }
type SyncJobStartOperationMapInput interface {
	pulumi.Input

	ToSyncJobStartOperationMapOutput() SyncJobStartOperationMapOutput
	ToSyncJobStartOperationMapOutputWithContext(context.Context) SyncJobStartOperationMapOutput
}

type SyncJobStartOperationMap map[string]SyncJobStartOperationInput

func (SyncJobStartOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobStartOperation)(nil)).Elem()
}

func (i SyncJobStartOperationMap) ToSyncJobStartOperationMapOutput() SyncJobStartOperationMapOutput {
	return i.ToSyncJobStartOperationMapOutputWithContext(context.Background())
}

func (i SyncJobStartOperationMap) ToSyncJobStartOperationMapOutputWithContext(ctx context.Context) SyncJobStartOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncJobStartOperationMapOutput)
}

type SyncJobStartOperationOutput struct{ *pulumi.OutputState }

func (SyncJobStartOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncJobStartOperation)(nil)).Elem()
}

func (o SyncJobStartOperationOutput) ToSyncJobStartOperationOutput() SyncJobStartOperationOutput {
	return o
}

func (o SyncJobStartOperationOutput) ToSyncJobStartOperationOutputWithContext(ctx context.Context) SyncJobStartOperationOutput {
	return o
}

// Synchronization instance id (i.e. identifies a synchronization job).
func (o SyncJobStartOperationOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncJobStartOperation) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

type SyncJobStartOperationArrayOutput struct{ *pulumi.OutputState }

func (SyncJobStartOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncJobStartOperation)(nil)).Elem()
}

func (o SyncJobStartOperationArrayOutput) ToSyncJobStartOperationArrayOutput() SyncJobStartOperationArrayOutput {
	return o
}

func (o SyncJobStartOperationArrayOutput) ToSyncJobStartOperationArrayOutputWithContext(ctx context.Context) SyncJobStartOperationArrayOutput {
	return o
}

func (o SyncJobStartOperationArrayOutput) Index(i pulumi.IntInput) SyncJobStartOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncJobStartOperation {
		return vs[0].([]*SyncJobStartOperation)[vs[1].(int)]
	}).(SyncJobStartOperationOutput)
}

type SyncJobStartOperationMapOutput struct{ *pulumi.OutputState }

func (SyncJobStartOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncJobStartOperation)(nil)).Elem()
}

func (o SyncJobStartOperationMapOutput) ToSyncJobStartOperationMapOutput() SyncJobStartOperationMapOutput {
	return o
}

func (o SyncJobStartOperationMapOutput) ToSyncJobStartOperationMapOutputWithContext(ctx context.Context) SyncJobStartOperationMapOutput {
	return o
}

func (o SyncJobStartOperationMapOutput) MapIndex(k pulumi.StringInput) SyncJobStartOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncJobStartOperation {
		return vs[0].(map[string]*SyncJobStartOperation)[vs[1].(string)]
	}).(SyncJobStartOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStartOperationInput)(nil)).Elem(), &SyncJobStartOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStartOperationArrayInput)(nil)).Elem(), SyncJobStartOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncJobStartOperationMapInput)(nil)).Elem(), SyncJobStartOperationMap{})
	pulumi.RegisterOutputType(SyncJobStartOperationOutput{})
	pulumi.RegisterOutputType(SyncJobStartOperationArrayOutput{})
	pulumi.RegisterOutputType(SyncJobStartOperationMapOutput{})
}
