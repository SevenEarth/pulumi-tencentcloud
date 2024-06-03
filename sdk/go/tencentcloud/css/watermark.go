// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a css watermark
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Css.NewWatermark(ctx, "watermark", &Css.WatermarkArgs{
//				Height:        pulumi.Int(0),
//				PictureUrl:    pulumi.String("picture_url"),
//				WatermarkName: pulumi.String("watermark_name"),
//				Width:         pulumi.Int(0),
//				XPosition:     pulumi.Int(0),
//				YPosition:     pulumi.Int(0),
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
//
// ## Import
//
// css watermark can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Css/watermark:Watermark watermark watermark_id
// ```
type Watermark struct {
	pulumi.CustomResourceState

	// height of the picture.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// watermark url.
	PictureUrl pulumi.StringOutput `pulumi:"pictureUrl"`
	// status. 0: not used, 1: used.
	Status pulumi.IntOutput `pulumi:"status"`
	// watermark name.
	WatermarkName pulumi.StringOutput `pulumi:"watermarkName"`
	// width of the picture.
	Width pulumi.IntPtrOutput `pulumi:"width"`
	// x position of the picture.
	XPosition pulumi.IntPtrOutput `pulumi:"xPosition"`
	// y position of the picture.
	YPosition pulumi.IntPtrOutput `pulumi:"yPosition"`
}

// NewWatermark registers a new resource with the given unique name, arguments, and options.
func NewWatermark(ctx *pulumi.Context,
	name string, args *WatermarkArgs, opts ...pulumi.ResourceOption) (*Watermark, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PictureUrl == nil {
		return nil, errors.New("invalid value for required argument 'PictureUrl'")
	}
	if args.WatermarkName == nil {
		return nil, errors.New("invalid value for required argument 'WatermarkName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Watermark
	err := ctx.RegisterResource("tencentcloud:Css/watermark:Watermark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatermark gets an existing Watermark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatermark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatermarkState, opts ...pulumi.ResourceOption) (*Watermark, error) {
	var resource Watermark
	err := ctx.ReadResource("tencentcloud:Css/watermark:Watermark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Watermark resources.
type watermarkState struct {
	// height of the picture.
	Height *int `pulumi:"height"`
	// watermark url.
	PictureUrl *string `pulumi:"pictureUrl"`
	// status. 0: not used, 1: used.
	Status *int `pulumi:"status"`
	// watermark name.
	WatermarkName *string `pulumi:"watermarkName"`
	// width of the picture.
	Width *int `pulumi:"width"`
	// x position of the picture.
	XPosition *int `pulumi:"xPosition"`
	// y position of the picture.
	YPosition *int `pulumi:"yPosition"`
}

type WatermarkState struct {
	// height of the picture.
	Height pulumi.IntPtrInput
	// watermark url.
	PictureUrl pulumi.StringPtrInput
	// status. 0: not used, 1: used.
	Status pulumi.IntPtrInput
	// watermark name.
	WatermarkName pulumi.StringPtrInput
	// width of the picture.
	Width pulumi.IntPtrInput
	// x position of the picture.
	XPosition pulumi.IntPtrInput
	// y position of the picture.
	YPosition pulumi.IntPtrInput
}

func (WatermarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*watermarkState)(nil)).Elem()
}

type watermarkArgs struct {
	// height of the picture.
	Height *int `pulumi:"height"`
	// watermark url.
	PictureUrl string `pulumi:"pictureUrl"`
	// watermark name.
	WatermarkName string `pulumi:"watermarkName"`
	// width of the picture.
	Width *int `pulumi:"width"`
	// x position of the picture.
	XPosition *int `pulumi:"xPosition"`
	// y position of the picture.
	YPosition *int `pulumi:"yPosition"`
}

// The set of arguments for constructing a Watermark resource.
type WatermarkArgs struct {
	// height of the picture.
	Height pulumi.IntPtrInput
	// watermark url.
	PictureUrl pulumi.StringInput
	// watermark name.
	WatermarkName pulumi.StringInput
	// width of the picture.
	Width pulumi.IntPtrInput
	// x position of the picture.
	XPosition pulumi.IntPtrInput
	// y position of the picture.
	YPosition pulumi.IntPtrInput
}

func (WatermarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watermarkArgs)(nil)).Elem()
}

type WatermarkInput interface {
	pulumi.Input

	ToWatermarkOutput() WatermarkOutput
	ToWatermarkOutputWithContext(ctx context.Context) WatermarkOutput
}

func (*Watermark) ElementType() reflect.Type {
	return reflect.TypeOf((**Watermark)(nil)).Elem()
}

func (i *Watermark) ToWatermarkOutput() WatermarkOutput {
	return i.ToWatermarkOutputWithContext(context.Background())
}

func (i *Watermark) ToWatermarkOutputWithContext(ctx context.Context) WatermarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkOutput)
}

// WatermarkArrayInput is an input type that accepts WatermarkArray and WatermarkArrayOutput values.
// You can construct a concrete instance of `WatermarkArrayInput` via:
//
//	WatermarkArray{ WatermarkArgs{...} }
type WatermarkArrayInput interface {
	pulumi.Input

	ToWatermarkArrayOutput() WatermarkArrayOutput
	ToWatermarkArrayOutputWithContext(context.Context) WatermarkArrayOutput
}

type WatermarkArray []WatermarkInput

func (WatermarkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Watermark)(nil)).Elem()
}

func (i WatermarkArray) ToWatermarkArrayOutput() WatermarkArrayOutput {
	return i.ToWatermarkArrayOutputWithContext(context.Background())
}

func (i WatermarkArray) ToWatermarkArrayOutputWithContext(ctx context.Context) WatermarkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkArrayOutput)
}

// WatermarkMapInput is an input type that accepts WatermarkMap and WatermarkMapOutput values.
// You can construct a concrete instance of `WatermarkMapInput` via:
//
//	WatermarkMap{ "key": WatermarkArgs{...} }
type WatermarkMapInput interface {
	pulumi.Input

	ToWatermarkMapOutput() WatermarkMapOutput
	ToWatermarkMapOutputWithContext(context.Context) WatermarkMapOutput
}

type WatermarkMap map[string]WatermarkInput

func (WatermarkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Watermark)(nil)).Elem()
}

func (i WatermarkMap) ToWatermarkMapOutput() WatermarkMapOutput {
	return i.ToWatermarkMapOutputWithContext(context.Background())
}

func (i WatermarkMap) ToWatermarkMapOutputWithContext(ctx context.Context) WatermarkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatermarkMapOutput)
}

type WatermarkOutput struct{ *pulumi.OutputState }

func (WatermarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watermark)(nil)).Elem()
}

func (o WatermarkOutput) ToWatermarkOutput() WatermarkOutput {
	return o
}

func (o WatermarkOutput) ToWatermarkOutputWithContext(ctx context.Context) WatermarkOutput {
	return o
}

// height of the picture.
func (o WatermarkOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watermark) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// watermark url.
func (o WatermarkOutput) PictureUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Watermark) pulumi.StringOutput { return v.PictureUrl }).(pulumi.StringOutput)
}

// status. 0: not used, 1: used.
func (o WatermarkOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Watermark) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// watermark name.
func (o WatermarkOutput) WatermarkName() pulumi.StringOutput {
	return o.ApplyT(func(v *Watermark) pulumi.StringOutput { return v.WatermarkName }).(pulumi.StringOutput)
}

// width of the picture.
func (o WatermarkOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watermark) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

// x position of the picture.
func (o WatermarkOutput) XPosition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watermark) pulumi.IntPtrOutput { return v.XPosition }).(pulumi.IntPtrOutput)
}

// y position of the picture.
func (o WatermarkOutput) YPosition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watermark) pulumi.IntPtrOutput { return v.YPosition }).(pulumi.IntPtrOutput)
}

type WatermarkArrayOutput struct{ *pulumi.OutputState }

func (WatermarkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Watermark)(nil)).Elem()
}

func (o WatermarkArrayOutput) ToWatermarkArrayOutput() WatermarkArrayOutput {
	return o
}

func (o WatermarkArrayOutput) ToWatermarkArrayOutputWithContext(ctx context.Context) WatermarkArrayOutput {
	return o
}

func (o WatermarkArrayOutput) Index(i pulumi.IntInput) WatermarkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Watermark {
		return vs[0].([]*Watermark)[vs[1].(int)]
	}).(WatermarkOutput)
}

type WatermarkMapOutput struct{ *pulumi.OutputState }

func (WatermarkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Watermark)(nil)).Elem()
}

func (o WatermarkMapOutput) ToWatermarkMapOutput() WatermarkMapOutput {
	return o
}

func (o WatermarkMapOutput) ToWatermarkMapOutputWithContext(ctx context.Context) WatermarkMapOutput {
	return o
}

func (o WatermarkMapOutput) MapIndex(k pulumi.StringInput) WatermarkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Watermark {
		return vs[0].(map[string]*Watermark)[vs[1].(string)]
	}).(WatermarkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkInput)(nil)).Elem(), &Watermark{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkArrayInput)(nil)).Elem(), WatermarkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatermarkMapInput)(nil)).Elem(), WatermarkMap{})
	pulumi.RegisterOutputType(WatermarkOutput{})
	pulumi.RegisterOutputType(WatermarkArrayOutput{})
	pulumi.RegisterOutputType(WatermarkMapOutput{})
}
