// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to download object
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cos.NewObjectDownloadOperation(ctx, "objectDownload", &Cos.ObjectDownloadOperationArgs{
// 			Bucket:       pulumi.String("xxxxxxx"),
// 			DownloadPath: pulumi.String("/tmp/test.txt"),
// 			Key:          pulumi.String("test.txt"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ObjectDownloadOperation struct {
	pulumi.CustomResourceState

	// Bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Download path.
	DownloadPath pulumi.StringOutput `pulumi:"downloadPath"`
	// Object key.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewObjectDownloadOperation registers a new resource with the given unique name, arguments, and options.
func NewObjectDownloadOperation(ctx *pulumi.Context,
	name string, args *ObjectDownloadOperationArgs, opts ...pulumi.ResourceOption) (*ObjectDownloadOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.DownloadPath == nil {
		return nil, errors.New("invalid value for required argument 'DownloadPath'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ObjectDownloadOperation
	err := ctx.RegisterResource("tencentcloud:Cos/objectDownloadOperation:ObjectDownloadOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectDownloadOperation gets an existing ObjectDownloadOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectDownloadOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectDownloadOperationState, opts ...pulumi.ResourceOption) (*ObjectDownloadOperation, error) {
	var resource ObjectDownloadOperation
	err := ctx.ReadResource("tencentcloud:Cos/objectDownloadOperation:ObjectDownloadOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectDownloadOperation resources.
type objectDownloadOperationState struct {
	// Bucket.
	Bucket *string `pulumi:"bucket"`
	// Download path.
	DownloadPath *string `pulumi:"downloadPath"`
	// Object key.
	Key *string `pulumi:"key"`
}

type ObjectDownloadOperationState struct {
	// Bucket.
	Bucket pulumi.StringPtrInput
	// Download path.
	DownloadPath pulumi.StringPtrInput
	// Object key.
	Key pulumi.StringPtrInput
}

func (ObjectDownloadOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectDownloadOperationState)(nil)).Elem()
}

type objectDownloadOperationArgs struct {
	// Bucket.
	Bucket string `pulumi:"bucket"`
	// Download path.
	DownloadPath string `pulumi:"downloadPath"`
	// Object key.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ObjectDownloadOperation resource.
type ObjectDownloadOperationArgs struct {
	// Bucket.
	Bucket pulumi.StringInput
	// Download path.
	DownloadPath pulumi.StringInput
	// Object key.
	Key pulumi.StringInput
}

func (ObjectDownloadOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectDownloadOperationArgs)(nil)).Elem()
}

type ObjectDownloadOperationInput interface {
	pulumi.Input

	ToObjectDownloadOperationOutput() ObjectDownloadOperationOutput
	ToObjectDownloadOperationOutputWithContext(ctx context.Context) ObjectDownloadOperationOutput
}

func (*ObjectDownloadOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectDownloadOperation)(nil)).Elem()
}

func (i *ObjectDownloadOperation) ToObjectDownloadOperationOutput() ObjectDownloadOperationOutput {
	return i.ToObjectDownloadOperationOutputWithContext(context.Background())
}

func (i *ObjectDownloadOperation) ToObjectDownloadOperationOutputWithContext(ctx context.Context) ObjectDownloadOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectDownloadOperationOutput)
}

// ObjectDownloadOperationArrayInput is an input type that accepts ObjectDownloadOperationArray and ObjectDownloadOperationArrayOutput values.
// You can construct a concrete instance of `ObjectDownloadOperationArrayInput` via:
//
//          ObjectDownloadOperationArray{ ObjectDownloadOperationArgs{...} }
type ObjectDownloadOperationArrayInput interface {
	pulumi.Input

	ToObjectDownloadOperationArrayOutput() ObjectDownloadOperationArrayOutput
	ToObjectDownloadOperationArrayOutputWithContext(context.Context) ObjectDownloadOperationArrayOutput
}

type ObjectDownloadOperationArray []ObjectDownloadOperationInput

func (ObjectDownloadOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectDownloadOperation)(nil)).Elem()
}

func (i ObjectDownloadOperationArray) ToObjectDownloadOperationArrayOutput() ObjectDownloadOperationArrayOutput {
	return i.ToObjectDownloadOperationArrayOutputWithContext(context.Background())
}

func (i ObjectDownloadOperationArray) ToObjectDownloadOperationArrayOutputWithContext(ctx context.Context) ObjectDownloadOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectDownloadOperationArrayOutput)
}

// ObjectDownloadOperationMapInput is an input type that accepts ObjectDownloadOperationMap and ObjectDownloadOperationMapOutput values.
// You can construct a concrete instance of `ObjectDownloadOperationMapInput` via:
//
//          ObjectDownloadOperationMap{ "key": ObjectDownloadOperationArgs{...} }
type ObjectDownloadOperationMapInput interface {
	pulumi.Input

	ToObjectDownloadOperationMapOutput() ObjectDownloadOperationMapOutput
	ToObjectDownloadOperationMapOutputWithContext(context.Context) ObjectDownloadOperationMapOutput
}

type ObjectDownloadOperationMap map[string]ObjectDownloadOperationInput

func (ObjectDownloadOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectDownloadOperation)(nil)).Elem()
}

func (i ObjectDownloadOperationMap) ToObjectDownloadOperationMapOutput() ObjectDownloadOperationMapOutput {
	return i.ToObjectDownloadOperationMapOutputWithContext(context.Background())
}

func (i ObjectDownloadOperationMap) ToObjectDownloadOperationMapOutputWithContext(ctx context.Context) ObjectDownloadOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectDownloadOperationMapOutput)
}

type ObjectDownloadOperationOutput struct{ *pulumi.OutputState }

func (ObjectDownloadOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectDownloadOperation)(nil)).Elem()
}

func (o ObjectDownloadOperationOutput) ToObjectDownloadOperationOutput() ObjectDownloadOperationOutput {
	return o
}

func (o ObjectDownloadOperationOutput) ToObjectDownloadOperationOutputWithContext(ctx context.Context) ObjectDownloadOperationOutput {
	return o
}

// Bucket.
func (o ObjectDownloadOperationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectDownloadOperation) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Download path.
func (o ObjectDownloadOperationOutput) DownloadPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectDownloadOperation) pulumi.StringOutput { return v.DownloadPath }).(pulumi.StringOutput)
}

// Object key.
func (o ObjectDownloadOperationOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectDownloadOperation) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ObjectDownloadOperationArrayOutput struct{ *pulumi.OutputState }

func (ObjectDownloadOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectDownloadOperation)(nil)).Elem()
}

func (o ObjectDownloadOperationArrayOutput) ToObjectDownloadOperationArrayOutput() ObjectDownloadOperationArrayOutput {
	return o
}

func (o ObjectDownloadOperationArrayOutput) ToObjectDownloadOperationArrayOutputWithContext(ctx context.Context) ObjectDownloadOperationArrayOutput {
	return o
}

func (o ObjectDownloadOperationArrayOutput) Index(i pulumi.IntInput) ObjectDownloadOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectDownloadOperation {
		return vs[0].([]*ObjectDownloadOperation)[vs[1].(int)]
	}).(ObjectDownloadOperationOutput)
}

type ObjectDownloadOperationMapOutput struct{ *pulumi.OutputState }

func (ObjectDownloadOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectDownloadOperation)(nil)).Elem()
}

func (o ObjectDownloadOperationMapOutput) ToObjectDownloadOperationMapOutput() ObjectDownloadOperationMapOutput {
	return o
}

func (o ObjectDownloadOperationMapOutput) ToObjectDownloadOperationMapOutputWithContext(ctx context.Context) ObjectDownloadOperationMapOutput {
	return o
}

func (o ObjectDownloadOperationMapOutput) MapIndex(k pulumi.StringInput) ObjectDownloadOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectDownloadOperation {
		return vs[0].(map[string]*ObjectDownloadOperation)[vs[1].(string)]
	}).(ObjectDownloadOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectDownloadOperationInput)(nil)).Elem(), &ObjectDownloadOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectDownloadOperationArrayInput)(nil)).Elem(), ObjectDownloadOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectDownloadOperationMapInput)(nil)).Elem(), ObjectDownloadOperationMap{})
	pulumi.RegisterOutputType(ObjectDownloadOperationOutput{})
	pulumi.RegisterOutputType(ObjectDownloadOperationArrayOutput{})
	pulumi.RegisterOutputType(ObjectDownloadOperationMapOutput{})
}
