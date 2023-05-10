// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ci bucketPicStyle
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ci"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ci.NewBucketPicStyle(ctx, "bucketPicStyle", &Ci.BucketPicStyleArgs{
// 			Bucket:    pulumi.String("terraform-ci-xxxxxx"),
// 			StyleBody: pulumi.String("imageMogr2/thumbnail/20x/crop/20x20/gravity/center/interlace/0/quality/100"),
// 			StyleName: pulumi.String("rayscale_2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ci bucket_pic_style can be imported using the bucket#styleName, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ci/bucketPicStyle:BucketPicStyle bucket_pic_style terraform-ci-xxxxxx#rayscale_2
// ```
type BucketPicStyle struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// style details, example: mageMogr2/grayscale/1.
	StyleBody pulumi.StringOutput `pulumi:"styleBody"`
	// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
	StyleName pulumi.StringOutput `pulumi:"styleName"`
}

// NewBucketPicStyle registers a new resource with the given unique name, arguments, and options.
func NewBucketPicStyle(ctx *pulumi.Context,
	name string, args *BucketPicStyleArgs, opts ...pulumi.ResourceOption) (*BucketPicStyle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.StyleBody == nil {
		return nil, errors.New("invalid value for required argument 'StyleBody'")
	}
	if args.StyleName == nil {
		return nil, errors.New("invalid value for required argument 'StyleName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BucketPicStyle
	err := ctx.RegisterResource("tencentcloud:Ci/bucketPicStyle:BucketPicStyle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPicStyle gets an existing BucketPicStyle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPicStyle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPicStyleState, opts ...pulumi.ResourceOption) (*BucketPicStyle, error) {
	var resource BucketPicStyle
	err := ctx.ReadResource("tencentcloud:Ci/bucketPicStyle:BucketPicStyle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPicStyle resources.
type bucketPicStyleState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// style details, example: mageMogr2/grayscale/1.
	StyleBody *string `pulumi:"styleBody"`
	// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
	StyleName *string `pulumi:"styleName"`
}

type BucketPicStyleState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// style details, example: mageMogr2/grayscale/1.
	StyleBody pulumi.StringPtrInput
	// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
	StyleName pulumi.StringPtrInput
}

func (BucketPicStyleState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPicStyleState)(nil)).Elem()
}

type bucketPicStyleArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// style details, example: mageMogr2/grayscale/1.
	StyleBody string `pulumi:"styleBody"`
	// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
	StyleName string `pulumi:"styleName"`
}

// The set of arguments for constructing a BucketPicStyle resource.
type BucketPicStyleArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// style details, example: mageMogr2/grayscale/1.
	StyleBody pulumi.StringInput
	// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
	StyleName pulumi.StringInput
}

func (BucketPicStyleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPicStyleArgs)(nil)).Elem()
}

type BucketPicStyleInput interface {
	pulumi.Input

	ToBucketPicStyleOutput() BucketPicStyleOutput
	ToBucketPicStyleOutputWithContext(ctx context.Context) BucketPicStyleOutput
}

func (*BucketPicStyle) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPicStyle)(nil)).Elem()
}

func (i *BucketPicStyle) ToBucketPicStyleOutput() BucketPicStyleOutput {
	return i.ToBucketPicStyleOutputWithContext(context.Background())
}

func (i *BucketPicStyle) ToBucketPicStyleOutputWithContext(ctx context.Context) BucketPicStyleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPicStyleOutput)
}

// BucketPicStyleArrayInput is an input type that accepts BucketPicStyleArray and BucketPicStyleArrayOutput values.
// You can construct a concrete instance of `BucketPicStyleArrayInput` via:
//
//          BucketPicStyleArray{ BucketPicStyleArgs{...} }
type BucketPicStyleArrayInput interface {
	pulumi.Input

	ToBucketPicStyleArrayOutput() BucketPicStyleArrayOutput
	ToBucketPicStyleArrayOutputWithContext(context.Context) BucketPicStyleArrayOutput
}

type BucketPicStyleArray []BucketPicStyleInput

func (BucketPicStyleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPicStyle)(nil)).Elem()
}

func (i BucketPicStyleArray) ToBucketPicStyleArrayOutput() BucketPicStyleArrayOutput {
	return i.ToBucketPicStyleArrayOutputWithContext(context.Background())
}

func (i BucketPicStyleArray) ToBucketPicStyleArrayOutputWithContext(ctx context.Context) BucketPicStyleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPicStyleArrayOutput)
}

// BucketPicStyleMapInput is an input type that accepts BucketPicStyleMap and BucketPicStyleMapOutput values.
// You can construct a concrete instance of `BucketPicStyleMapInput` via:
//
//          BucketPicStyleMap{ "key": BucketPicStyleArgs{...} }
type BucketPicStyleMapInput interface {
	pulumi.Input

	ToBucketPicStyleMapOutput() BucketPicStyleMapOutput
	ToBucketPicStyleMapOutputWithContext(context.Context) BucketPicStyleMapOutput
}

type BucketPicStyleMap map[string]BucketPicStyleInput

func (BucketPicStyleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPicStyle)(nil)).Elem()
}

func (i BucketPicStyleMap) ToBucketPicStyleMapOutput() BucketPicStyleMapOutput {
	return i.ToBucketPicStyleMapOutputWithContext(context.Background())
}

func (i BucketPicStyleMap) ToBucketPicStyleMapOutputWithContext(ctx context.Context) BucketPicStyleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPicStyleMapOutput)
}

type BucketPicStyleOutput struct{ *pulumi.OutputState }

func (BucketPicStyleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPicStyle)(nil)).Elem()
}

func (o BucketPicStyleOutput) ToBucketPicStyleOutput() BucketPicStyleOutput {
	return o
}

func (o BucketPicStyleOutput) ToBucketPicStyleOutputWithContext(ctx context.Context) BucketPicStyleOutput {
	return o
}

// bucket name.
func (o BucketPicStyleOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPicStyle) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// style details, example: mageMogr2/grayscale/1.
func (o BucketPicStyleOutput) StyleBody() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPicStyle) pulumi.StringOutput { return v.StyleBody }).(pulumi.StringOutput)
}

// style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
func (o BucketPicStyleOutput) StyleName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPicStyle) pulumi.StringOutput { return v.StyleName }).(pulumi.StringOutput)
}

type BucketPicStyleArrayOutput struct{ *pulumi.OutputState }

func (BucketPicStyleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPicStyle)(nil)).Elem()
}

func (o BucketPicStyleArrayOutput) ToBucketPicStyleArrayOutput() BucketPicStyleArrayOutput {
	return o
}

func (o BucketPicStyleArrayOutput) ToBucketPicStyleArrayOutputWithContext(ctx context.Context) BucketPicStyleArrayOutput {
	return o
}

func (o BucketPicStyleArrayOutput) Index(i pulumi.IntInput) BucketPicStyleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketPicStyle {
		return vs[0].([]*BucketPicStyle)[vs[1].(int)]
	}).(BucketPicStyleOutput)
}

type BucketPicStyleMapOutput struct{ *pulumi.OutputState }

func (BucketPicStyleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPicStyle)(nil)).Elem()
}

func (o BucketPicStyleMapOutput) ToBucketPicStyleMapOutput() BucketPicStyleMapOutput {
	return o
}

func (o BucketPicStyleMapOutput) ToBucketPicStyleMapOutputWithContext(ctx context.Context) BucketPicStyleMapOutput {
	return o
}

func (o BucketPicStyleMapOutput) MapIndex(k pulumi.StringInput) BucketPicStyleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketPicStyle {
		return vs[0].(map[string]*BucketPicStyle)[vs[1].(string)]
	}).(BucketPicStyleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPicStyleInput)(nil)).Elem(), &BucketPicStyle{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPicStyleArrayInput)(nil)).Elem(), BucketPicStyleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPicStyleMapInput)(nil)).Elem(), BucketPicStyleMap{})
	pulumi.RegisterOutputType(BucketPicStyleOutput{})
	pulumi.RegisterOutputType(BucketPicStyleArrayOutput{})
	pulumi.RegisterOutputType(BucketPicStyleMapOutput{})
}
