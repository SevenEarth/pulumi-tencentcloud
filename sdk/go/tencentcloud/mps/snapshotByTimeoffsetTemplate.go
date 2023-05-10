// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mps snapshotByTimeoffsetTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mps.NewSnapshotByTimeoffsetTemplate(ctx, "snapshotByTimeoffsetTemplate", &Mps.SnapshotByTimeoffsetTemplateArgs{
// 			FillType:           pulumi.String("stretch"),
// 			Format:             pulumi.String("jpg"),
// 			Height:             pulumi.Int(128),
// 			ResolutionAdaptive: pulumi.String("open"),
// 			Width:              pulumi.Int(140),
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
// mps snapshot_by_timeoffset_template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate snapshot_by_timeoffset_template snapshot_by_timeoffset_template_id
// ```
type SnapshotByTimeoffsetTemplate struct {
	pulumi.CustomResourceState

	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrOutput `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Snapshot by timeoffset template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrOutput `pulumi:"resolutionAdaptive"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewSnapshotByTimeoffsetTemplate registers a new resource with the given unique name, arguments, and options.
func NewSnapshotByTimeoffsetTemplate(ctx *pulumi.Context,
	name string, args *SnapshotByTimeoffsetTemplateArgs, opts ...pulumi.ResourceOption) (*SnapshotByTimeoffsetTemplate, error) {
	if args == nil {
		args = &SnapshotByTimeoffsetTemplateArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SnapshotByTimeoffsetTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotByTimeoffsetTemplate gets an existing SnapshotByTimeoffsetTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotByTimeoffsetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotByTimeoffsetTemplateState, opts ...pulumi.ResourceOption) (*SnapshotByTimeoffsetTemplate, error) {
	var resource SnapshotByTimeoffsetTemplate
	err := ctx.ReadResource("tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotByTimeoffsetTemplate resources.
type snapshotByTimeoffsetTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Snapshot by timeoffset template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

type SnapshotByTimeoffsetTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Snapshot by timeoffset template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SnapshotByTimeoffsetTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotByTimeoffsetTemplateState)(nil)).Elem()
}

type snapshotByTimeoffsetTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Snapshot by timeoffset template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a SnapshotByTimeoffsetTemplate resource.
type SnapshotByTimeoffsetTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Snapshot by timeoffset template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SnapshotByTimeoffsetTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotByTimeoffsetTemplateArgs)(nil)).Elem()
}

type SnapshotByTimeoffsetTemplateInput interface {
	pulumi.Input

	ToSnapshotByTimeoffsetTemplateOutput() SnapshotByTimeoffsetTemplateOutput
	ToSnapshotByTimeoffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateOutput
}

func (*SnapshotByTimeoffsetTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (i *SnapshotByTimeoffsetTemplate) ToSnapshotByTimeoffsetTemplateOutput() SnapshotByTimeoffsetTemplateOutput {
	return i.ToSnapshotByTimeoffsetTemplateOutputWithContext(context.Background())
}

func (i *SnapshotByTimeoffsetTemplate) ToSnapshotByTimeoffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeoffsetTemplateOutput)
}

// SnapshotByTimeoffsetTemplateArrayInput is an input type that accepts SnapshotByTimeoffsetTemplateArray and SnapshotByTimeoffsetTemplateArrayOutput values.
// You can construct a concrete instance of `SnapshotByTimeoffsetTemplateArrayInput` via:
//
//          SnapshotByTimeoffsetTemplateArray{ SnapshotByTimeoffsetTemplateArgs{...} }
type SnapshotByTimeoffsetTemplateArrayInput interface {
	pulumi.Input

	ToSnapshotByTimeoffsetTemplateArrayOutput() SnapshotByTimeoffsetTemplateArrayOutput
	ToSnapshotByTimeoffsetTemplateArrayOutputWithContext(context.Context) SnapshotByTimeoffsetTemplateArrayOutput
}

type SnapshotByTimeoffsetTemplateArray []SnapshotByTimeoffsetTemplateInput

func (SnapshotByTimeoffsetTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (i SnapshotByTimeoffsetTemplateArray) ToSnapshotByTimeoffsetTemplateArrayOutput() SnapshotByTimeoffsetTemplateArrayOutput {
	return i.ToSnapshotByTimeoffsetTemplateArrayOutputWithContext(context.Background())
}

func (i SnapshotByTimeoffsetTemplateArray) ToSnapshotByTimeoffsetTemplateArrayOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeoffsetTemplateArrayOutput)
}

// SnapshotByTimeoffsetTemplateMapInput is an input type that accepts SnapshotByTimeoffsetTemplateMap and SnapshotByTimeoffsetTemplateMapOutput values.
// You can construct a concrete instance of `SnapshotByTimeoffsetTemplateMapInput` via:
//
//          SnapshotByTimeoffsetTemplateMap{ "key": SnapshotByTimeoffsetTemplateArgs{...} }
type SnapshotByTimeoffsetTemplateMapInput interface {
	pulumi.Input

	ToSnapshotByTimeoffsetTemplateMapOutput() SnapshotByTimeoffsetTemplateMapOutput
	ToSnapshotByTimeoffsetTemplateMapOutputWithContext(context.Context) SnapshotByTimeoffsetTemplateMapOutput
}

type SnapshotByTimeoffsetTemplateMap map[string]SnapshotByTimeoffsetTemplateInput

func (SnapshotByTimeoffsetTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (i SnapshotByTimeoffsetTemplateMap) ToSnapshotByTimeoffsetTemplateMapOutput() SnapshotByTimeoffsetTemplateMapOutput {
	return i.ToSnapshotByTimeoffsetTemplateMapOutputWithContext(context.Background())
}

func (i SnapshotByTimeoffsetTemplateMap) ToSnapshotByTimeoffsetTemplateMapOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeoffsetTemplateMapOutput)
}

type SnapshotByTimeoffsetTemplateOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeoffsetTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeoffsetTemplateOutput) ToSnapshotByTimeoffsetTemplateOutput() SnapshotByTimeoffsetTemplateOutput {
	return o
}

func (o SnapshotByTimeoffsetTemplateOutput) ToSnapshotByTimeoffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateOutput {
	return o
}

// Template description information, length limit: 256 characters.
func (o SnapshotByTimeoffsetTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
func (o SnapshotByTimeoffsetTemplateOutput) FillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.StringPtrOutput { return v.FillType }).(pulumi.StringPtrOutput)
}

// Image format, the value can be jpg, png, webp. Default is jpg.
func (o SnapshotByTimeoffsetTemplateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o SnapshotByTimeoffsetTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Snapshot by timeoffset template name, length limit: 64 characters.
func (o SnapshotByTimeoffsetTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
func (o SnapshotByTimeoffsetTemplateOutput) ResolutionAdaptive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.StringPtrOutput { return v.ResolutionAdaptive }).(pulumi.StringPtrOutput)
}

// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o SnapshotByTimeoffsetTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeoffsetTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type SnapshotByTimeoffsetTemplateArrayOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeoffsetTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeoffsetTemplateArrayOutput) ToSnapshotByTimeoffsetTemplateArrayOutput() SnapshotByTimeoffsetTemplateArrayOutput {
	return o
}

func (o SnapshotByTimeoffsetTemplateArrayOutput) ToSnapshotByTimeoffsetTemplateArrayOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateArrayOutput {
	return o
}

func (o SnapshotByTimeoffsetTemplateArrayOutput) Index(i pulumi.IntInput) SnapshotByTimeoffsetTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotByTimeoffsetTemplate {
		return vs[0].([]*SnapshotByTimeoffsetTemplate)[vs[1].(int)]
	}).(SnapshotByTimeoffsetTemplateOutput)
}

type SnapshotByTimeoffsetTemplateMapOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeoffsetTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotByTimeoffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeoffsetTemplateMapOutput) ToSnapshotByTimeoffsetTemplateMapOutput() SnapshotByTimeoffsetTemplateMapOutput {
	return o
}

func (o SnapshotByTimeoffsetTemplateMapOutput) ToSnapshotByTimeoffsetTemplateMapOutputWithContext(ctx context.Context) SnapshotByTimeoffsetTemplateMapOutput {
	return o
}

func (o SnapshotByTimeoffsetTemplateMapOutput) MapIndex(k pulumi.StringInput) SnapshotByTimeoffsetTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotByTimeoffsetTemplate {
		return vs[0].(map[string]*SnapshotByTimeoffsetTemplate)[vs[1].(string)]
	}).(SnapshotByTimeoffsetTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeoffsetTemplateInput)(nil)).Elem(), &SnapshotByTimeoffsetTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeoffsetTemplateArrayInput)(nil)).Elem(), SnapshotByTimeoffsetTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeoffsetTemplateMapInput)(nil)).Elem(), SnapshotByTimeoffsetTemplateMap{})
	pulumi.RegisterOutputType(SnapshotByTimeoffsetTemplateOutput{})
	pulumi.RegisterOutputType(SnapshotByTimeoffsetTemplateArrayOutput{})
	pulumi.RegisterOutputType(SnapshotByTimeoffsetTemplateMapOutput{})
}
