// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a VOD snapshot by time offset template.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vod"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vod.NewSnapshotByTimeOffsetTemplate(ctx, "foo", &Vod.SnapshotByTimeOffsetTemplateArgs{
//				Comment:            pulumi.String("test"),
//				FillType:           pulumi.String("white"),
//				Format:             pulumi.String("png"),
//				Height:             pulumi.Int(128),
//				ResolutionAdaptive: pulumi.Bool(false),
//				Width:              pulumi.Int(130),
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
// VOD snapshot by time offset template can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vod/snapshotByTimeOffsetTemplate:SnapshotByTimeOffsetTemplate foo 46906
//
// ```
type SnapshotByTimeOffsetTemplate struct {
	pulumi.CustomResourceState

	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Creation time of template in ISO date format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType pulumi.StringPtrOutput `pulumi:"fillType"`
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive pulumi.BoolPtrOutput `pulumi:"resolutionAdaptive"`
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId pulumi.IntPtrOutput `pulumi:"subAppId"`
	// Last modified time of template in ISO date format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewSnapshotByTimeOffsetTemplate registers a new resource with the given unique name, arguments, and options.
func NewSnapshotByTimeOffsetTemplate(ctx *pulumi.Context,
	name string, args *SnapshotByTimeOffsetTemplateArgs, opts ...pulumi.ResourceOption) (*SnapshotByTimeOffsetTemplate, error) {
	if args == nil {
		args = &SnapshotByTimeOffsetTemplateArgs{}
	}

	var resource SnapshotByTimeOffsetTemplate
	err := ctx.RegisterResource("tencentcloud:Vod/snapshotByTimeOffsetTemplate:SnapshotByTimeOffsetTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotByTimeOffsetTemplate gets an existing SnapshotByTimeOffsetTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotByTimeOffsetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotByTimeOffsetTemplateState, opts ...pulumi.ResourceOption) (*SnapshotByTimeOffsetTemplate, error) {
	var resource SnapshotByTimeOffsetTemplate
	err := ctx.ReadResource("tencentcloud:Vod/snapshotByTimeOffsetTemplate:SnapshotByTimeOffsetTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotByTimeOffsetTemplate resources.
type snapshotByTimeOffsetTemplateState struct {
	// Template description. Length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Creation time of template in ISO date format.
	CreateTime *string `pulumi:"createTime"`
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType *string `pulumi:"fillType"`
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format *string `pulumi:"format"`
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height *int `pulumi:"height"`
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive *bool `pulumi:"resolutionAdaptive"`
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *int `pulumi:"subAppId"`
	// Last modified time of template in ISO date format.
	UpdateTime *string `pulumi:"updateTime"`
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width *int `pulumi:"width"`
}

type SnapshotByTimeOffsetTemplateState struct {
	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Creation time of template in ISO date format.
	CreateTime pulumi.StringPtrInput
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType pulumi.StringPtrInput
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format pulumi.StringPtrInput
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height pulumi.IntPtrInput
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive pulumi.BoolPtrInput
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId pulumi.IntPtrInput
	// Last modified time of template in ISO date format.
	UpdateTime pulumi.StringPtrInput
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width pulumi.IntPtrInput
}

func (SnapshotByTimeOffsetTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotByTimeOffsetTemplateState)(nil)).Elem()
}

type snapshotByTimeOffsetTemplateArgs struct {
	// Template description. Length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType *string `pulumi:"fillType"`
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format *string `pulumi:"format"`
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height *int `pulumi:"height"`
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive *bool `pulumi:"resolutionAdaptive"`
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *int `pulumi:"subAppId"`
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a SnapshotByTimeOffsetTemplate resource.
type SnapshotByTimeOffsetTemplateArgs struct {
	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType pulumi.StringPtrInput
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format pulumi.StringPtrInput
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height pulumi.IntPtrInput
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive pulumi.BoolPtrInput
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId pulumi.IntPtrInput
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width pulumi.IntPtrInput
}

func (SnapshotByTimeOffsetTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotByTimeOffsetTemplateArgs)(nil)).Elem()
}

type SnapshotByTimeOffsetTemplateInput interface {
	pulumi.Input

	ToSnapshotByTimeOffsetTemplateOutput() SnapshotByTimeOffsetTemplateOutput
	ToSnapshotByTimeOffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateOutput
}

func (*SnapshotByTimeOffsetTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (i *SnapshotByTimeOffsetTemplate) ToSnapshotByTimeOffsetTemplateOutput() SnapshotByTimeOffsetTemplateOutput {
	return i.ToSnapshotByTimeOffsetTemplateOutputWithContext(context.Background())
}

func (i *SnapshotByTimeOffsetTemplate) ToSnapshotByTimeOffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeOffsetTemplateOutput)
}

// SnapshotByTimeOffsetTemplateArrayInput is an input type that accepts SnapshotByTimeOffsetTemplateArray and SnapshotByTimeOffsetTemplateArrayOutput values.
// You can construct a concrete instance of `SnapshotByTimeOffsetTemplateArrayInput` via:
//
//	SnapshotByTimeOffsetTemplateArray{ SnapshotByTimeOffsetTemplateArgs{...} }
type SnapshotByTimeOffsetTemplateArrayInput interface {
	pulumi.Input

	ToSnapshotByTimeOffsetTemplateArrayOutput() SnapshotByTimeOffsetTemplateArrayOutput
	ToSnapshotByTimeOffsetTemplateArrayOutputWithContext(context.Context) SnapshotByTimeOffsetTemplateArrayOutput
}

type SnapshotByTimeOffsetTemplateArray []SnapshotByTimeOffsetTemplateInput

func (SnapshotByTimeOffsetTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (i SnapshotByTimeOffsetTemplateArray) ToSnapshotByTimeOffsetTemplateArrayOutput() SnapshotByTimeOffsetTemplateArrayOutput {
	return i.ToSnapshotByTimeOffsetTemplateArrayOutputWithContext(context.Background())
}

func (i SnapshotByTimeOffsetTemplateArray) ToSnapshotByTimeOffsetTemplateArrayOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeOffsetTemplateArrayOutput)
}

// SnapshotByTimeOffsetTemplateMapInput is an input type that accepts SnapshotByTimeOffsetTemplateMap and SnapshotByTimeOffsetTemplateMapOutput values.
// You can construct a concrete instance of `SnapshotByTimeOffsetTemplateMapInput` via:
//
//	SnapshotByTimeOffsetTemplateMap{ "key": SnapshotByTimeOffsetTemplateArgs{...} }
type SnapshotByTimeOffsetTemplateMapInput interface {
	pulumi.Input

	ToSnapshotByTimeOffsetTemplateMapOutput() SnapshotByTimeOffsetTemplateMapOutput
	ToSnapshotByTimeOffsetTemplateMapOutputWithContext(context.Context) SnapshotByTimeOffsetTemplateMapOutput
}

type SnapshotByTimeOffsetTemplateMap map[string]SnapshotByTimeOffsetTemplateInput

func (SnapshotByTimeOffsetTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (i SnapshotByTimeOffsetTemplateMap) ToSnapshotByTimeOffsetTemplateMapOutput() SnapshotByTimeOffsetTemplateMapOutput {
	return i.ToSnapshotByTimeOffsetTemplateMapOutputWithContext(context.Background())
}

func (i SnapshotByTimeOffsetTemplateMap) ToSnapshotByTimeOffsetTemplateMapOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotByTimeOffsetTemplateMapOutput)
}

type SnapshotByTimeOffsetTemplateOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeOffsetTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeOffsetTemplateOutput) ToSnapshotByTimeOffsetTemplateOutput() SnapshotByTimeOffsetTemplateOutput {
	return o
}

func (o SnapshotByTimeOffsetTemplateOutput) ToSnapshotByTimeOffsetTemplateOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateOutput {
	return o
}

// Template description. Length limit: 256 characters.
func (o SnapshotByTimeOffsetTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Creation time of template in ISO date format.
func (o SnapshotByTimeOffsetTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
func (o SnapshotByTimeOffsetTemplateOutput) FillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringPtrOutput { return v.FillType }).(pulumi.StringPtrOutput)
}

// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
func (o SnapshotByTimeOffsetTemplateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
func (o SnapshotByTimeOffsetTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Name of a time point screen capturing template. Length limit: 64 characters.
func (o SnapshotByTimeOffsetTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
func (o SnapshotByTimeOffsetTemplateOutput) ResolutionAdaptive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.BoolPtrOutput { return v.ResolutionAdaptive }).(pulumi.BoolPtrOutput)
}

// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
func (o SnapshotByTimeOffsetTemplateOutput) SubAppId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.IntPtrOutput { return v.SubAppId }).(pulumi.IntPtrOutput)
}

// Last modified time of template in ISO date format.
func (o SnapshotByTimeOffsetTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
func (o SnapshotByTimeOffsetTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotByTimeOffsetTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type SnapshotByTimeOffsetTemplateArrayOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeOffsetTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeOffsetTemplateArrayOutput) ToSnapshotByTimeOffsetTemplateArrayOutput() SnapshotByTimeOffsetTemplateArrayOutput {
	return o
}

func (o SnapshotByTimeOffsetTemplateArrayOutput) ToSnapshotByTimeOffsetTemplateArrayOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateArrayOutput {
	return o
}

func (o SnapshotByTimeOffsetTemplateArrayOutput) Index(i pulumi.IntInput) SnapshotByTimeOffsetTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotByTimeOffsetTemplate {
		return vs[0].([]*SnapshotByTimeOffsetTemplate)[vs[1].(int)]
	}).(SnapshotByTimeOffsetTemplateOutput)
}

type SnapshotByTimeOffsetTemplateMapOutput struct{ *pulumi.OutputState }

func (SnapshotByTimeOffsetTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotByTimeOffsetTemplate)(nil)).Elem()
}

func (o SnapshotByTimeOffsetTemplateMapOutput) ToSnapshotByTimeOffsetTemplateMapOutput() SnapshotByTimeOffsetTemplateMapOutput {
	return o
}

func (o SnapshotByTimeOffsetTemplateMapOutput) ToSnapshotByTimeOffsetTemplateMapOutputWithContext(ctx context.Context) SnapshotByTimeOffsetTemplateMapOutput {
	return o
}

func (o SnapshotByTimeOffsetTemplateMapOutput) MapIndex(k pulumi.StringInput) SnapshotByTimeOffsetTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotByTimeOffsetTemplate {
		return vs[0].(map[string]*SnapshotByTimeOffsetTemplate)[vs[1].(string)]
	}).(SnapshotByTimeOffsetTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeOffsetTemplateInput)(nil)).Elem(), &SnapshotByTimeOffsetTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeOffsetTemplateArrayInput)(nil)).Elem(), SnapshotByTimeOffsetTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotByTimeOffsetTemplateMapInput)(nil)).Elem(), SnapshotByTimeOffsetTemplateMap{})
	pulumi.RegisterOutputType(SnapshotByTimeOffsetTemplateOutput{})
	pulumi.RegisterOutputType(SnapshotByTimeOffsetTemplateArrayOutput{})
	pulumi.RegisterOutputType(SnapshotByTimeOffsetTemplateMapOutput{})
}
