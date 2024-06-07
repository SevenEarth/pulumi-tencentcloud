// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a mps aiAnalysisTemplate
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mps.NewAiAnalysisTemplate(ctx, "aiAnalysisTemplate", &Mps.AiAnalysisTemplateArgs{
//				ClassificationConfigure: &mps.AiAnalysisTemplateClassificationConfigureArgs{
//					Switch: pulumi.String("OFF"),
//				},
//				CoverConfigure: &mps.AiAnalysisTemplateCoverConfigureArgs{
//					Switch: pulumi.String("ON"),
//				},
//				FrameTagConfigure: &mps.AiAnalysisTemplateFrameTagConfigureArgs{
//					Switch: pulumi.String("ON"),
//				},
//				TagConfigure: &mps.AiAnalysisTemplateTagConfigureArgs{
//					Switch: pulumi.String("ON"),
//				},
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
// mps ai_analysis_template can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Mps/aiAnalysisTemplate:AiAnalysisTemplate ai_analysis_template ai_analysis_template_id
// ```
type AiAnalysisTemplate struct {
	pulumi.CustomResourceState

	// Ai classification task control parameters.
	ClassificationConfigure AiAnalysisTemplateClassificationConfigurePtrOutput `pulumi:"classificationConfigure"`
	// Ai analysis template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Ai cover task control parameters.
	CoverConfigure AiAnalysisTemplateCoverConfigurePtrOutput `pulumi:"coverConfigure"`
	// Ai frame tag task control parameters.
	FrameTagConfigure AiAnalysisTemplateFrameTagConfigurePtrOutput `pulumi:"frameTagConfigure"`
	// Ai analysis template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ai tag task control parameters.
	TagConfigure AiAnalysisTemplateTagConfigurePtrOutput `pulumi:"tagConfigure"`
}

// NewAiAnalysisTemplate registers a new resource with the given unique name, arguments, and options.
func NewAiAnalysisTemplate(ctx *pulumi.Context,
	name string, args *AiAnalysisTemplateArgs, opts ...pulumi.ResourceOption) (*AiAnalysisTemplate, error) {
	if args == nil {
		args = &AiAnalysisTemplateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiAnalysisTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/aiAnalysisTemplate:AiAnalysisTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiAnalysisTemplate gets an existing AiAnalysisTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiAnalysisTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiAnalysisTemplateState, opts ...pulumi.ResourceOption) (*AiAnalysisTemplate, error) {
	var resource AiAnalysisTemplate
	err := ctx.ReadResource("tencentcloud:Mps/aiAnalysisTemplate:AiAnalysisTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiAnalysisTemplate resources.
type aiAnalysisTemplateState struct {
	// Ai classification task control parameters.
	ClassificationConfigure *AiAnalysisTemplateClassificationConfigure `pulumi:"classificationConfigure"`
	// Ai analysis template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Ai cover task control parameters.
	CoverConfigure *AiAnalysisTemplateCoverConfigure `pulumi:"coverConfigure"`
	// Ai frame tag task control parameters.
	FrameTagConfigure *AiAnalysisTemplateFrameTagConfigure `pulumi:"frameTagConfigure"`
	// Ai analysis template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Ai tag task control parameters.
	TagConfigure *AiAnalysisTemplateTagConfigure `pulumi:"tagConfigure"`
}

type AiAnalysisTemplateState struct {
	// Ai classification task control parameters.
	ClassificationConfigure AiAnalysisTemplateClassificationConfigurePtrInput
	// Ai analysis template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Ai cover task control parameters.
	CoverConfigure AiAnalysisTemplateCoverConfigurePtrInput
	// Ai frame tag task control parameters.
	FrameTagConfigure AiAnalysisTemplateFrameTagConfigurePtrInput
	// Ai analysis template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Ai tag task control parameters.
	TagConfigure AiAnalysisTemplateTagConfigurePtrInput
}

func (AiAnalysisTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiAnalysisTemplateState)(nil)).Elem()
}

type aiAnalysisTemplateArgs struct {
	// Ai classification task control parameters.
	ClassificationConfigure *AiAnalysisTemplateClassificationConfigure `pulumi:"classificationConfigure"`
	// Ai analysis template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Ai cover task control parameters.
	CoverConfigure *AiAnalysisTemplateCoverConfigure `pulumi:"coverConfigure"`
	// Ai frame tag task control parameters.
	FrameTagConfigure *AiAnalysisTemplateFrameTagConfigure `pulumi:"frameTagConfigure"`
	// Ai analysis template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Ai tag task control parameters.
	TagConfigure *AiAnalysisTemplateTagConfigure `pulumi:"tagConfigure"`
}

// The set of arguments for constructing a AiAnalysisTemplate resource.
type AiAnalysisTemplateArgs struct {
	// Ai classification task control parameters.
	ClassificationConfigure AiAnalysisTemplateClassificationConfigurePtrInput
	// Ai analysis template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Ai cover task control parameters.
	CoverConfigure AiAnalysisTemplateCoverConfigurePtrInput
	// Ai frame tag task control parameters.
	FrameTagConfigure AiAnalysisTemplateFrameTagConfigurePtrInput
	// Ai analysis template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Ai tag task control parameters.
	TagConfigure AiAnalysisTemplateTagConfigurePtrInput
}

func (AiAnalysisTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiAnalysisTemplateArgs)(nil)).Elem()
}

type AiAnalysisTemplateInput interface {
	pulumi.Input

	ToAiAnalysisTemplateOutput() AiAnalysisTemplateOutput
	ToAiAnalysisTemplateOutputWithContext(ctx context.Context) AiAnalysisTemplateOutput
}

func (*AiAnalysisTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**AiAnalysisTemplate)(nil)).Elem()
}

func (i *AiAnalysisTemplate) ToAiAnalysisTemplateOutput() AiAnalysisTemplateOutput {
	return i.ToAiAnalysisTemplateOutputWithContext(context.Background())
}

func (i *AiAnalysisTemplate) ToAiAnalysisTemplateOutputWithContext(ctx context.Context) AiAnalysisTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiAnalysisTemplateOutput)
}

// AiAnalysisTemplateArrayInput is an input type that accepts AiAnalysisTemplateArray and AiAnalysisTemplateArrayOutput values.
// You can construct a concrete instance of `AiAnalysisTemplateArrayInput` via:
//
//	AiAnalysisTemplateArray{ AiAnalysisTemplateArgs{...} }
type AiAnalysisTemplateArrayInput interface {
	pulumi.Input

	ToAiAnalysisTemplateArrayOutput() AiAnalysisTemplateArrayOutput
	ToAiAnalysisTemplateArrayOutputWithContext(context.Context) AiAnalysisTemplateArrayOutput
}

type AiAnalysisTemplateArray []AiAnalysisTemplateInput

func (AiAnalysisTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiAnalysisTemplate)(nil)).Elem()
}

func (i AiAnalysisTemplateArray) ToAiAnalysisTemplateArrayOutput() AiAnalysisTemplateArrayOutput {
	return i.ToAiAnalysisTemplateArrayOutputWithContext(context.Background())
}

func (i AiAnalysisTemplateArray) ToAiAnalysisTemplateArrayOutputWithContext(ctx context.Context) AiAnalysisTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiAnalysisTemplateArrayOutput)
}

// AiAnalysisTemplateMapInput is an input type that accepts AiAnalysisTemplateMap and AiAnalysisTemplateMapOutput values.
// You can construct a concrete instance of `AiAnalysisTemplateMapInput` via:
//
//	AiAnalysisTemplateMap{ "key": AiAnalysisTemplateArgs{...} }
type AiAnalysisTemplateMapInput interface {
	pulumi.Input

	ToAiAnalysisTemplateMapOutput() AiAnalysisTemplateMapOutput
	ToAiAnalysisTemplateMapOutputWithContext(context.Context) AiAnalysisTemplateMapOutput
}

type AiAnalysisTemplateMap map[string]AiAnalysisTemplateInput

func (AiAnalysisTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiAnalysisTemplate)(nil)).Elem()
}

func (i AiAnalysisTemplateMap) ToAiAnalysisTemplateMapOutput() AiAnalysisTemplateMapOutput {
	return i.ToAiAnalysisTemplateMapOutputWithContext(context.Background())
}

func (i AiAnalysisTemplateMap) ToAiAnalysisTemplateMapOutputWithContext(ctx context.Context) AiAnalysisTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiAnalysisTemplateMapOutput)
}

type AiAnalysisTemplateOutput struct{ *pulumi.OutputState }

func (AiAnalysisTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiAnalysisTemplate)(nil)).Elem()
}

func (o AiAnalysisTemplateOutput) ToAiAnalysisTemplateOutput() AiAnalysisTemplateOutput {
	return o
}

func (o AiAnalysisTemplateOutput) ToAiAnalysisTemplateOutputWithContext(ctx context.Context) AiAnalysisTemplateOutput {
	return o
}

// Ai classification task control parameters.
func (o AiAnalysisTemplateOutput) ClassificationConfigure() AiAnalysisTemplateClassificationConfigurePtrOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) AiAnalysisTemplateClassificationConfigurePtrOutput {
		return v.ClassificationConfigure
	}).(AiAnalysisTemplateClassificationConfigurePtrOutput)
}

// Ai analysis template description information, length limit: 256 characters.
func (o AiAnalysisTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Ai cover task control parameters.
func (o AiAnalysisTemplateOutput) CoverConfigure() AiAnalysisTemplateCoverConfigurePtrOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) AiAnalysisTemplateCoverConfigurePtrOutput { return v.CoverConfigure }).(AiAnalysisTemplateCoverConfigurePtrOutput)
}

// Ai frame tag task control parameters.
func (o AiAnalysisTemplateOutput) FrameTagConfigure() AiAnalysisTemplateFrameTagConfigurePtrOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) AiAnalysisTemplateFrameTagConfigurePtrOutput { return v.FrameTagConfigure }).(AiAnalysisTemplateFrameTagConfigurePtrOutput)
}

// Ai analysis template name, length limit: 64 characters.
func (o AiAnalysisTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Ai tag task control parameters.
func (o AiAnalysisTemplateOutput) TagConfigure() AiAnalysisTemplateTagConfigurePtrOutput {
	return o.ApplyT(func(v *AiAnalysisTemplate) AiAnalysisTemplateTagConfigurePtrOutput { return v.TagConfigure }).(AiAnalysisTemplateTagConfigurePtrOutput)
}

type AiAnalysisTemplateArrayOutput struct{ *pulumi.OutputState }

func (AiAnalysisTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiAnalysisTemplate)(nil)).Elem()
}

func (o AiAnalysisTemplateArrayOutput) ToAiAnalysisTemplateArrayOutput() AiAnalysisTemplateArrayOutput {
	return o
}

func (o AiAnalysisTemplateArrayOutput) ToAiAnalysisTemplateArrayOutputWithContext(ctx context.Context) AiAnalysisTemplateArrayOutput {
	return o
}

func (o AiAnalysisTemplateArrayOutput) Index(i pulumi.IntInput) AiAnalysisTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiAnalysisTemplate {
		return vs[0].([]*AiAnalysisTemplate)[vs[1].(int)]
	}).(AiAnalysisTemplateOutput)
}

type AiAnalysisTemplateMapOutput struct{ *pulumi.OutputState }

func (AiAnalysisTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiAnalysisTemplate)(nil)).Elem()
}

func (o AiAnalysisTemplateMapOutput) ToAiAnalysisTemplateMapOutput() AiAnalysisTemplateMapOutput {
	return o
}

func (o AiAnalysisTemplateMapOutput) ToAiAnalysisTemplateMapOutputWithContext(ctx context.Context) AiAnalysisTemplateMapOutput {
	return o
}

func (o AiAnalysisTemplateMapOutput) MapIndex(k pulumi.StringInput) AiAnalysisTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiAnalysisTemplate {
		return vs[0].(map[string]*AiAnalysisTemplate)[vs[1].(string)]
	}).(AiAnalysisTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiAnalysisTemplateInput)(nil)).Elem(), &AiAnalysisTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiAnalysisTemplateArrayInput)(nil)).Elem(), AiAnalysisTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiAnalysisTemplateMapInput)(nil)).Elem(), AiAnalysisTemplateMap{})
	pulumi.RegisterOutputType(AiAnalysisTemplateOutput{})
	pulumi.RegisterOutputType(AiAnalysisTemplateArrayOutput{})
	pulumi.RegisterOutputType(AiAnalysisTemplateMapOutput{})
}
