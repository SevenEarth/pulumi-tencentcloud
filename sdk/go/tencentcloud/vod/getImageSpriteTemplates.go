// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of VOD image sprite templates.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vod"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vod"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooImageSpriteTemplate, err := Vod.NewImageSpriteTemplate(ctx, "fooImageSpriteTemplate", &Vod.ImageSpriteTemplateArgs{
// 			SampleType:         pulumi.String("Percent"),
// 			SampleInterval:     pulumi.Int(10),
// 			RowCount:           pulumi.Int(3),
// 			ColumnCount:        pulumi.Int(3),
// 			Comment:            pulumi.String("test"),
// 			FillType:           pulumi.String("stretch"),
// 			Width:              pulumi.Int(128),
// 			Height:             pulumi.Int(128),
// 			ResolutionAdaptive: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = Vod.GetImageSpriteTemplatesOutput(ctx, vod.GetImageSpriteTemplatesOutputArgs{
// 			Type:       pulumi.String("Custom"),
// 			Definition: fooImageSpriteTemplate.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func GetImageSpriteTemplates(ctx *pulumi.Context, args *GetImageSpriteTemplatesArgs, opts ...pulumi.InvokeOption) (*GetImageSpriteTemplatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetImageSpriteTemplatesResult
	err := ctx.Invoke("tencentcloud:Vod/getImageSpriteTemplates:getImageSpriteTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageSpriteTemplates.
type GetImageSpriteTemplatesArgs struct {
	// Unique ID filter of image sprite template.
	Definition *string `pulumi:"definition"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId *int `pulumi:"subAppId"`
	// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getImageSpriteTemplates.
type GetImageSpriteTemplatesResult struct {
	// Unique ID of image sprite template.
	Definition *string `pulumi:"definition"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SubAppId         *int    `pulumi:"subAppId"`
	// A list of image sprite templates. Each element contains the following attributes:
	TemplateLists []GetImageSpriteTemplatesTemplateList `pulumi:"templateLists"`
	// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
	Type *string `pulumi:"type"`
}

func GetImageSpriteTemplatesOutput(ctx *pulumi.Context, args GetImageSpriteTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetImageSpriteTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageSpriteTemplatesResult, error) {
			args := v.(GetImageSpriteTemplatesArgs)
			r, err := GetImageSpriteTemplates(ctx, &args, opts...)
			var s GetImageSpriteTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageSpriteTemplatesResultOutput)
}

// A collection of arguments for invoking getImageSpriteTemplates.
type GetImageSpriteTemplatesOutputArgs struct {
	// Unique ID filter of image sprite template.
	Definition pulumi.StringPtrInput `pulumi:"definition"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	SubAppId pulumi.IntPtrInput `pulumi:"subAppId"`
	// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetImageSpriteTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageSpriteTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getImageSpriteTemplates.
type GetImageSpriteTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetImageSpriteTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageSpriteTemplatesResult)(nil)).Elem()
}

func (o GetImageSpriteTemplatesResultOutput) ToGetImageSpriteTemplatesResultOutput() GetImageSpriteTemplatesResultOutput {
	return o
}

func (o GetImageSpriteTemplatesResultOutput) ToGetImageSpriteTemplatesResultOutputWithContext(ctx context.Context) GetImageSpriteTemplatesResultOutput {
	return o
}

// Unique ID of image sprite template.
func (o GetImageSpriteTemplatesResultOutput) Definition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.Definition }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetImageSpriteTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetImageSpriteTemplatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetImageSpriteTemplatesResultOutput) SubAppId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *int { return v.SubAppId }).(pulumi.IntPtrOutput)
}

// A list of image sprite templates. Each element contains the following attributes:
func (o GetImageSpriteTemplatesResultOutput) TemplateLists() GetImageSpriteTemplatesTemplateListArrayOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) []GetImageSpriteTemplatesTemplateList { return v.TemplateLists }).(GetImageSpriteTemplatesTemplateListArrayOutput)
}

// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
func (o GetImageSpriteTemplatesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageSpriteTemplatesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageSpriteTemplatesResultOutput{})
}
