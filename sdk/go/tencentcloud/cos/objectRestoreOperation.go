// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to restore object
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewObjectRestoreOperation(ctx, "objectRestore", &Cos.ObjectRestoreOperationArgs{
//				Bucket: pulumi.String("keep-test-1308919341"),
//				Days:   pulumi.Int(2),
//				Key:    pulumi.String("test-restore.txt"),
//				Tier:   pulumi.String("Expedited"),
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
type ObjectRestoreOperation struct {
	pulumi.CustomResourceState

	// Bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies the valid duration of the restored temporary copy in days.
	Days pulumi.IntOutput `pulumi:"days"`
	// Object key.
	Key pulumi.StringOutput `pulumi:"key"`
	// when restoring, Tier can be specified as the supported recovery model.
	// There are three recovery models for recovering archived storage type data, which are:
	// - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
	// - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
	// - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
	//   For deep recovery archive storage type data, there are two recovery models, which are:
	// - Standard: standard retrieval mode, recovery time is 12-24 hours.
	// - Bulk: batch retrieval mode, recovery time is 24-48 hours.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewObjectRestoreOperation registers a new resource with the given unique name, arguments, and options.
func NewObjectRestoreOperation(ctx *pulumi.Context,
	name string, args *ObjectRestoreOperationArgs, opts ...pulumi.ResourceOption) (*ObjectRestoreOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Days == nil {
		return nil, errors.New("invalid value for required argument 'Days'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectRestoreOperation
	err := ctx.RegisterResource("tencentcloud:Cos/objectRestoreOperation:ObjectRestoreOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectRestoreOperation gets an existing ObjectRestoreOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectRestoreOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectRestoreOperationState, opts ...pulumi.ResourceOption) (*ObjectRestoreOperation, error) {
	var resource ObjectRestoreOperation
	err := ctx.ReadResource("tencentcloud:Cos/objectRestoreOperation:ObjectRestoreOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectRestoreOperation resources.
type objectRestoreOperationState struct {
	// Bucket.
	Bucket *string `pulumi:"bucket"`
	// Specifies the valid duration of the restored temporary copy in days.
	Days *int `pulumi:"days"`
	// Object key.
	Key *string `pulumi:"key"`
	// when restoring, Tier can be specified as the supported recovery model.
	// There are three recovery models for recovering archived storage type data, which are:
	// - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
	// - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
	// - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
	//   For deep recovery archive storage type data, there are two recovery models, which are:
	// - Standard: standard retrieval mode, recovery time is 12-24 hours.
	// - Bulk: batch retrieval mode, recovery time is 24-48 hours.
	Tier *string `pulumi:"tier"`
}

type ObjectRestoreOperationState struct {
	// Bucket.
	Bucket pulumi.StringPtrInput
	// Specifies the valid duration of the restored temporary copy in days.
	Days pulumi.IntPtrInput
	// Object key.
	Key pulumi.StringPtrInput
	// when restoring, Tier can be specified as the supported recovery model.
	// There are three recovery models for recovering archived storage type data, which are:
	// - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
	// - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
	// - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
	//   For deep recovery archive storage type data, there are two recovery models, which are:
	// - Standard: standard retrieval mode, recovery time is 12-24 hours.
	// - Bulk: batch retrieval mode, recovery time is 24-48 hours.
	Tier pulumi.StringPtrInput
}

func (ObjectRestoreOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectRestoreOperationState)(nil)).Elem()
}

type objectRestoreOperationArgs struct {
	// Bucket.
	Bucket string `pulumi:"bucket"`
	// Specifies the valid duration of the restored temporary copy in days.
	Days int `pulumi:"days"`
	// Object key.
	Key string `pulumi:"key"`
	// when restoring, Tier can be specified as the supported recovery model.
	// There are three recovery models for recovering archived storage type data, which are:
	// - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
	// - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
	// - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
	//   For deep recovery archive storage type data, there are two recovery models, which are:
	// - Standard: standard retrieval mode, recovery time is 12-24 hours.
	// - Bulk: batch retrieval mode, recovery time is 24-48 hours.
	Tier string `pulumi:"tier"`
}

// The set of arguments for constructing a ObjectRestoreOperation resource.
type ObjectRestoreOperationArgs struct {
	// Bucket.
	Bucket pulumi.StringInput
	// Specifies the valid duration of the restored temporary copy in days.
	Days pulumi.IntInput
	// Object key.
	Key pulumi.StringInput
	// when restoring, Tier can be specified as the supported recovery model.
	// There are three recovery models for recovering archived storage type data, which are:
	// - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
	// - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
	// - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
	//   For deep recovery archive storage type data, there are two recovery models, which are:
	// - Standard: standard retrieval mode, recovery time is 12-24 hours.
	// - Bulk: batch retrieval mode, recovery time is 24-48 hours.
	Tier pulumi.StringInput
}

func (ObjectRestoreOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectRestoreOperationArgs)(nil)).Elem()
}

type ObjectRestoreOperationInput interface {
	pulumi.Input

	ToObjectRestoreOperationOutput() ObjectRestoreOperationOutput
	ToObjectRestoreOperationOutputWithContext(ctx context.Context) ObjectRestoreOperationOutput
}

func (*ObjectRestoreOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectRestoreOperation)(nil)).Elem()
}

func (i *ObjectRestoreOperation) ToObjectRestoreOperationOutput() ObjectRestoreOperationOutput {
	return i.ToObjectRestoreOperationOutputWithContext(context.Background())
}

func (i *ObjectRestoreOperation) ToObjectRestoreOperationOutputWithContext(ctx context.Context) ObjectRestoreOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectRestoreOperationOutput)
}

// ObjectRestoreOperationArrayInput is an input type that accepts ObjectRestoreOperationArray and ObjectRestoreOperationArrayOutput values.
// You can construct a concrete instance of `ObjectRestoreOperationArrayInput` via:
//
//	ObjectRestoreOperationArray{ ObjectRestoreOperationArgs{...} }
type ObjectRestoreOperationArrayInput interface {
	pulumi.Input

	ToObjectRestoreOperationArrayOutput() ObjectRestoreOperationArrayOutput
	ToObjectRestoreOperationArrayOutputWithContext(context.Context) ObjectRestoreOperationArrayOutput
}

type ObjectRestoreOperationArray []ObjectRestoreOperationInput

func (ObjectRestoreOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectRestoreOperation)(nil)).Elem()
}

func (i ObjectRestoreOperationArray) ToObjectRestoreOperationArrayOutput() ObjectRestoreOperationArrayOutput {
	return i.ToObjectRestoreOperationArrayOutputWithContext(context.Background())
}

func (i ObjectRestoreOperationArray) ToObjectRestoreOperationArrayOutputWithContext(ctx context.Context) ObjectRestoreOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectRestoreOperationArrayOutput)
}

// ObjectRestoreOperationMapInput is an input type that accepts ObjectRestoreOperationMap and ObjectRestoreOperationMapOutput values.
// You can construct a concrete instance of `ObjectRestoreOperationMapInput` via:
//
//	ObjectRestoreOperationMap{ "key": ObjectRestoreOperationArgs{...} }
type ObjectRestoreOperationMapInput interface {
	pulumi.Input

	ToObjectRestoreOperationMapOutput() ObjectRestoreOperationMapOutput
	ToObjectRestoreOperationMapOutputWithContext(context.Context) ObjectRestoreOperationMapOutput
}

type ObjectRestoreOperationMap map[string]ObjectRestoreOperationInput

func (ObjectRestoreOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectRestoreOperation)(nil)).Elem()
}

func (i ObjectRestoreOperationMap) ToObjectRestoreOperationMapOutput() ObjectRestoreOperationMapOutput {
	return i.ToObjectRestoreOperationMapOutputWithContext(context.Background())
}

func (i ObjectRestoreOperationMap) ToObjectRestoreOperationMapOutputWithContext(ctx context.Context) ObjectRestoreOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectRestoreOperationMapOutput)
}

type ObjectRestoreOperationOutput struct{ *pulumi.OutputState }

func (ObjectRestoreOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectRestoreOperation)(nil)).Elem()
}

func (o ObjectRestoreOperationOutput) ToObjectRestoreOperationOutput() ObjectRestoreOperationOutput {
	return o
}

func (o ObjectRestoreOperationOutput) ToObjectRestoreOperationOutputWithContext(ctx context.Context) ObjectRestoreOperationOutput {
	return o
}

// Bucket.
func (o ObjectRestoreOperationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectRestoreOperation) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Specifies the valid duration of the restored temporary copy in days.
func (o ObjectRestoreOperationOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v *ObjectRestoreOperation) pulumi.IntOutput { return v.Days }).(pulumi.IntOutput)
}

// Object key.
func (o ObjectRestoreOperationOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectRestoreOperation) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// when restoring, Tier can be specified as the supported recovery model.
// There are three recovery models for recovering archived storage type data, which are:
//   - Expedited: quick retrieval mode, and the recovery task can be completed in 1-5 minutes.
//   - Standard: standard retrieval mode. Recovery task is completed within 3-5 hours.
//   - Bulk: batch retrieval mode, and the recovery task is completed within 5-12 hours.
//     For deep recovery archive storage type data, there are two recovery models, which are:
//   - Standard: standard retrieval mode, recovery time is 12-24 hours.
//   - Bulk: batch retrieval mode, recovery time is 24-48 hours.
func (o ObjectRestoreOperationOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectRestoreOperation) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

type ObjectRestoreOperationArrayOutput struct{ *pulumi.OutputState }

func (ObjectRestoreOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectRestoreOperation)(nil)).Elem()
}

func (o ObjectRestoreOperationArrayOutput) ToObjectRestoreOperationArrayOutput() ObjectRestoreOperationArrayOutput {
	return o
}

func (o ObjectRestoreOperationArrayOutput) ToObjectRestoreOperationArrayOutputWithContext(ctx context.Context) ObjectRestoreOperationArrayOutput {
	return o
}

func (o ObjectRestoreOperationArrayOutput) Index(i pulumi.IntInput) ObjectRestoreOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectRestoreOperation {
		return vs[0].([]*ObjectRestoreOperation)[vs[1].(int)]
	}).(ObjectRestoreOperationOutput)
}

type ObjectRestoreOperationMapOutput struct{ *pulumi.OutputState }

func (ObjectRestoreOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectRestoreOperation)(nil)).Elem()
}

func (o ObjectRestoreOperationMapOutput) ToObjectRestoreOperationMapOutput() ObjectRestoreOperationMapOutput {
	return o
}

func (o ObjectRestoreOperationMapOutput) ToObjectRestoreOperationMapOutputWithContext(ctx context.Context) ObjectRestoreOperationMapOutput {
	return o
}

func (o ObjectRestoreOperationMapOutput) MapIndex(k pulumi.StringInput) ObjectRestoreOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectRestoreOperation {
		return vs[0].(map[string]*ObjectRestoreOperation)[vs[1].(string)]
	}).(ObjectRestoreOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectRestoreOperationInput)(nil)).Elem(), &ObjectRestoreOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectRestoreOperationArrayInput)(nil)).Elem(), ObjectRestoreOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectRestoreOperationMapInput)(nil)).Elem(), ObjectRestoreOperationMap{})
	pulumi.RegisterOutputType(ObjectRestoreOperationOutput{})
	pulumi.RegisterOutputType(ObjectRestoreOperationArrayOutput{})
	pulumi.RegisterOutputType(ObjectRestoreOperationMapOutput{})
}
