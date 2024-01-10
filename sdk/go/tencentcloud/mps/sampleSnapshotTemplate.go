// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mps sampleSnapshotTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mps.NewSampleSnapshotTemplate(ctx, "sampleSnapshotTemplate", &Mps.SampleSnapshotTemplateArgs{
//				FillType:           pulumi.String("stretch"),
//				Format:             pulumi.String("jpg"),
//				Height:             pulumi.Int(128),
//				ResolutionAdaptive: pulumi.String("open"),
//				SampleInterval:     pulumi.Int(10),
//				SampleType:         pulumi.String("Percent"),
//				Width:              pulumi.Int(140),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// mps sample_snapshot_template can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Mps/sampleSnapshotTemplate:SampleSnapshotTemplate sample_snapshot_template sample_snapshot_template_id
//
// ```
type SampleSnapshotTemplate struct {
	pulumi.CustomResourceState

	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrOutput `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Sample snapshot template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrOutput `pulumi:"resolutionAdaptive"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntOutput `pulumi:"sampleInterval"`
	// Sampling snapshot type, optional value:Percent/Time.
	SampleType pulumi.StringOutput `pulumi:"sampleType"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewSampleSnapshotTemplate registers a new resource with the given unique name, arguments, and options.
func NewSampleSnapshotTemplate(ctx *pulumi.Context,
	name string, args *SampleSnapshotTemplateArgs, opts ...pulumi.ResourceOption) (*SampleSnapshotTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SampleInterval == nil {
		return nil, errors.New("invalid value for required argument 'SampleInterval'")
	}
	if args.SampleType == nil {
		return nil, errors.New("invalid value for required argument 'SampleType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SampleSnapshotTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/sampleSnapshotTemplate:SampleSnapshotTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSampleSnapshotTemplate gets an existing SampleSnapshotTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSampleSnapshotTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SampleSnapshotTemplateState, opts ...pulumi.ResourceOption) (*SampleSnapshotTemplate, error) {
	var resource SampleSnapshotTemplate
	err := ctx.ReadResource("tencentcloud:Mps/sampleSnapshotTemplate:SampleSnapshotTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SampleSnapshotTemplate resources.
type sampleSnapshotTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Sample snapshot template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval *int `pulumi:"sampleInterval"`
	// Sampling snapshot type, optional value:Percent/Time.
	SampleType *string `pulumi:"sampleType"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

type SampleSnapshotTemplateState struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Sample snapshot template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntPtrInput
	// Sampling snapshot type, optional value:Percent/Time.
	SampleType pulumi.StringPtrInput
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SampleSnapshotTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sampleSnapshotTemplateState)(nil)).Elem()
}

type sampleSnapshotTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format *string `pulumi:"format"`
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height *int `pulumi:"height"`
	// Sample snapshot template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval int `pulumi:"sampleInterval"`
	// Sampling snapshot type, optional value:Percent/Time.
	SampleType string `pulumi:"sampleType"`
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a SampleSnapshotTemplate resource.
type SampleSnapshotTemplateArgs struct {
	// Template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format, the value can be jpg, png, webp. Default is jpg.
	Format pulumi.StringPtrInput
	// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Height pulumi.IntPtrInput
	// Sample snapshot template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
	SampleInterval pulumi.IntInput
	// Sampling snapshot type, optional value:Percent/Time.
	SampleType pulumi.StringInput
	// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SampleSnapshotTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sampleSnapshotTemplateArgs)(nil)).Elem()
}

type SampleSnapshotTemplateInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput
	ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput
}

func (*SampleSnapshotTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**SampleSnapshotTemplate)(nil)).Elem()
}

func (i *SampleSnapshotTemplate) ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput {
	return i.ToSampleSnapshotTemplateOutputWithContext(context.Background())
}

func (i *SampleSnapshotTemplate) ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateOutput)
}

// SampleSnapshotTemplateArrayInput is an input type that accepts SampleSnapshotTemplateArray and SampleSnapshotTemplateArrayOutput values.
// You can construct a concrete instance of `SampleSnapshotTemplateArrayInput` via:
//
//	SampleSnapshotTemplateArray{ SampleSnapshotTemplateArgs{...} }
type SampleSnapshotTemplateArrayInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput
	ToSampleSnapshotTemplateArrayOutputWithContext(context.Context) SampleSnapshotTemplateArrayOutput
}

type SampleSnapshotTemplateArray []SampleSnapshotTemplateInput

func (SampleSnapshotTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SampleSnapshotTemplate)(nil)).Elem()
}

func (i SampleSnapshotTemplateArray) ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput {
	return i.ToSampleSnapshotTemplateArrayOutputWithContext(context.Background())
}

func (i SampleSnapshotTemplateArray) ToSampleSnapshotTemplateArrayOutputWithContext(ctx context.Context) SampleSnapshotTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateArrayOutput)
}

// SampleSnapshotTemplateMapInput is an input type that accepts SampleSnapshotTemplateMap and SampleSnapshotTemplateMapOutput values.
// You can construct a concrete instance of `SampleSnapshotTemplateMapInput` via:
//
//	SampleSnapshotTemplateMap{ "key": SampleSnapshotTemplateArgs{...} }
type SampleSnapshotTemplateMapInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput
	ToSampleSnapshotTemplateMapOutputWithContext(context.Context) SampleSnapshotTemplateMapOutput
}

type SampleSnapshotTemplateMap map[string]SampleSnapshotTemplateInput

func (SampleSnapshotTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SampleSnapshotTemplate)(nil)).Elem()
}

func (i SampleSnapshotTemplateMap) ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput {
	return i.ToSampleSnapshotTemplateMapOutputWithContext(context.Background())
}

func (i SampleSnapshotTemplateMap) ToSampleSnapshotTemplateMapOutputWithContext(ctx context.Context) SampleSnapshotTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateMapOutput)
}

type SampleSnapshotTemplateOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateOutput) ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput {
	return o
}

func (o SampleSnapshotTemplateOutput) ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput {
	return o
}

// Template description information, length limit: 256 characters.
func (o SampleSnapshotTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
func (o SampleSnapshotTemplateOutput) FillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.FillType }).(pulumi.StringPtrOutput)
}

// Image format, the value can be jpg, png, webp. Default is jpg.
func (o SampleSnapshotTemplateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o SampleSnapshotTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Sample snapshot template name, length limit: 64 characters.
func (o SampleSnapshotTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
func (o SampleSnapshotTemplateOutput) ResolutionAdaptive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.ResolutionAdaptive }).(pulumi.StringPtrOutput)
}

// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
func (o SampleSnapshotTemplateOutput) SampleInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntOutput { return v.SampleInterval }).(pulumi.IntOutput)
}

// Sampling snapshot type, optional value:Percent/Time.
func (o SampleSnapshotTemplateOutput) SampleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringOutput { return v.SampleType }).(pulumi.StringOutput)
}

// The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
func (o SampleSnapshotTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type SampleSnapshotTemplateArrayOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateArrayOutput) ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput {
	return o
}

func (o SampleSnapshotTemplateArrayOutput) ToSampleSnapshotTemplateArrayOutputWithContext(ctx context.Context) SampleSnapshotTemplateArrayOutput {
	return o
}

func (o SampleSnapshotTemplateArrayOutput) Index(i pulumi.IntInput) SampleSnapshotTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SampleSnapshotTemplate {
		return vs[0].([]*SampleSnapshotTemplate)[vs[1].(int)]
	}).(SampleSnapshotTemplateOutput)
}

type SampleSnapshotTemplateMapOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateMapOutput) ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput {
	return o
}

func (o SampleSnapshotTemplateMapOutput) ToSampleSnapshotTemplateMapOutputWithContext(ctx context.Context) SampleSnapshotTemplateMapOutput {
	return o
}

func (o SampleSnapshotTemplateMapOutput) MapIndex(k pulumi.StringInput) SampleSnapshotTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SampleSnapshotTemplate {
		return vs[0].(map[string]*SampleSnapshotTemplate)[vs[1].(string)]
	}).(SampleSnapshotTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateInput)(nil)).Elem(), &SampleSnapshotTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateArrayInput)(nil)).Elem(), SampleSnapshotTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateMapInput)(nil)).Elem(), SampleSnapshotTemplateMap{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateOutput{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateArrayOutput{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateMapOutput{})
}
