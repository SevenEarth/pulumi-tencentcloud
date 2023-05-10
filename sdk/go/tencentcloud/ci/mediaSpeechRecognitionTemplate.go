// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ci mediaSpeechRecognitionTemplate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ci"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ci.NewMediaSpeechRecognitionTemplate(ctx, "mediaSpeechRecognitionTemplate", &Ci.MediaSpeechRecognitionTemplateArgs{
// 			Bucket: pulumi.String("terraform-ci-1308919341"),
// 			SpeechRecognition: &ci.MediaSpeechRecognitionTemplateSpeechRecognitionArgs{
// 				ChannelNum:         pulumi.String("1"),
// 				ConvertNumMode:     pulumi.String("0"),
// 				EngineModelType:    pulumi.String("16k_zh"),
// 				FilterDirty:        pulumi.String("0"),
// 				FilterModal:        pulumi.String("1"),
// 				FilterPunc:         pulumi.String("0"),
// 				OutputFileType:     pulumi.String("txt"),
// 				ResTextFormat:      pulumi.String("1"),
// 				SpeakerDiarization: pulumi.String("1"),
// 				SpeakerNumber:      pulumi.String("0"),
// 			},
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
// ci media_speech_recognition_template can be imported using the bucket#templateId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate media_speech_recognition_template terraform-ci-xxxxxx#t1d794430f2f1f4350b11e905ce2c6167e
// ```
type MediaSpeechRecognitionTemplate struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// audio configuration.
	SpeechRecognition MediaSpeechRecognitionTemplateSpeechRecognitionOutput `pulumi:"speechRecognition"`
}

// NewMediaSpeechRecognitionTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaSpeechRecognitionTemplate(ctx *pulumi.Context,
	name string, args *MediaSpeechRecognitionTemplateArgs, opts ...pulumi.ResourceOption) (*MediaSpeechRecognitionTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.SpeechRecognition == nil {
		return nil, errors.New("invalid value for required argument 'SpeechRecognition'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MediaSpeechRecognitionTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaSpeechRecognitionTemplate gets an existing MediaSpeechRecognitionTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaSpeechRecognitionTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaSpeechRecognitionTemplateState, opts ...pulumi.ResourceOption) (*MediaSpeechRecognitionTemplate, error) {
	var resource MediaSpeechRecognitionTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaSpeechRecognitionTemplate resources.
type mediaSpeechRecognitionTemplateState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// audio configuration.
	SpeechRecognition *MediaSpeechRecognitionTemplateSpeechRecognition `pulumi:"speechRecognition"`
}

type MediaSpeechRecognitionTemplateState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// audio configuration.
	SpeechRecognition MediaSpeechRecognitionTemplateSpeechRecognitionPtrInput
}

func (MediaSpeechRecognitionTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSpeechRecognitionTemplateState)(nil)).Elem()
}

type mediaSpeechRecognitionTemplateArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// audio configuration.
	SpeechRecognition MediaSpeechRecognitionTemplateSpeechRecognition `pulumi:"speechRecognition"`
}

// The set of arguments for constructing a MediaSpeechRecognitionTemplate resource.
type MediaSpeechRecognitionTemplateArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// audio configuration.
	SpeechRecognition MediaSpeechRecognitionTemplateSpeechRecognitionInput
}

func (MediaSpeechRecognitionTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSpeechRecognitionTemplateArgs)(nil)).Elem()
}

type MediaSpeechRecognitionTemplateInput interface {
	pulumi.Input

	ToMediaSpeechRecognitionTemplateOutput() MediaSpeechRecognitionTemplateOutput
	ToMediaSpeechRecognitionTemplateOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateOutput
}

func (*MediaSpeechRecognitionTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (i *MediaSpeechRecognitionTemplate) ToMediaSpeechRecognitionTemplateOutput() MediaSpeechRecognitionTemplateOutput {
	return i.ToMediaSpeechRecognitionTemplateOutputWithContext(context.Background())
}

func (i *MediaSpeechRecognitionTemplate) ToMediaSpeechRecognitionTemplateOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSpeechRecognitionTemplateOutput)
}

// MediaSpeechRecognitionTemplateArrayInput is an input type that accepts MediaSpeechRecognitionTemplateArray and MediaSpeechRecognitionTemplateArrayOutput values.
// You can construct a concrete instance of `MediaSpeechRecognitionTemplateArrayInput` via:
//
//          MediaSpeechRecognitionTemplateArray{ MediaSpeechRecognitionTemplateArgs{...} }
type MediaSpeechRecognitionTemplateArrayInput interface {
	pulumi.Input

	ToMediaSpeechRecognitionTemplateArrayOutput() MediaSpeechRecognitionTemplateArrayOutput
	ToMediaSpeechRecognitionTemplateArrayOutputWithContext(context.Context) MediaSpeechRecognitionTemplateArrayOutput
}

type MediaSpeechRecognitionTemplateArray []MediaSpeechRecognitionTemplateInput

func (MediaSpeechRecognitionTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (i MediaSpeechRecognitionTemplateArray) ToMediaSpeechRecognitionTemplateArrayOutput() MediaSpeechRecognitionTemplateArrayOutput {
	return i.ToMediaSpeechRecognitionTemplateArrayOutputWithContext(context.Background())
}

func (i MediaSpeechRecognitionTemplateArray) ToMediaSpeechRecognitionTemplateArrayOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSpeechRecognitionTemplateArrayOutput)
}

// MediaSpeechRecognitionTemplateMapInput is an input type that accepts MediaSpeechRecognitionTemplateMap and MediaSpeechRecognitionTemplateMapOutput values.
// You can construct a concrete instance of `MediaSpeechRecognitionTemplateMapInput` via:
//
//          MediaSpeechRecognitionTemplateMap{ "key": MediaSpeechRecognitionTemplateArgs{...} }
type MediaSpeechRecognitionTemplateMapInput interface {
	pulumi.Input

	ToMediaSpeechRecognitionTemplateMapOutput() MediaSpeechRecognitionTemplateMapOutput
	ToMediaSpeechRecognitionTemplateMapOutputWithContext(context.Context) MediaSpeechRecognitionTemplateMapOutput
}

type MediaSpeechRecognitionTemplateMap map[string]MediaSpeechRecognitionTemplateInput

func (MediaSpeechRecognitionTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (i MediaSpeechRecognitionTemplateMap) ToMediaSpeechRecognitionTemplateMapOutput() MediaSpeechRecognitionTemplateMapOutput {
	return i.ToMediaSpeechRecognitionTemplateMapOutputWithContext(context.Background())
}

func (i MediaSpeechRecognitionTemplateMap) ToMediaSpeechRecognitionTemplateMapOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSpeechRecognitionTemplateMapOutput)
}

type MediaSpeechRecognitionTemplateOutput struct{ *pulumi.OutputState }

func (MediaSpeechRecognitionTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (o MediaSpeechRecognitionTemplateOutput) ToMediaSpeechRecognitionTemplateOutput() MediaSpeechRecognitionTemplateOutput {
	return o
}

func (o MediaSpeechRecognitionTemplateOutput) ToMediaSpeechRecognitionTemplateOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateOutput {
	return o
}

// bucket name.
func (o MediaSpeechRecognitionTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSpeechRecognitionTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaSpeechRecognitionTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSpeechRecognitionTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// audio configuration.
func (o MediaSpeechRecognitionTemplateOutput) SpeechRecognition() MediaSpeechRecognitionTemplateSpeechRecognitionOutput {
	return o.ApplyT(func(v *MediaSpeechRecognitionTemplate) MediaSpeechRecognitionTemplateSpeechRecognitionOutput {
		return v.SpeechRecognition
	}).(MediaSpeechRecognitionTemplateSpeechRecognitionOutput)
}

type MediaSpeechRecognitionTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaSpeechRecognitionTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (o MediaSpeechRecognitionTemplateArrayOutput) ToMediaSpeechRecognitionTemplateArrayOutput() MediaSpeechRecognitionTemplateArrayOutput {
	return o
}

func (o MediaSpeechRecognitionTemplateArrayOutput) ToMediaSpeechRecognitionTemplateArrayOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateArrayOutput {
	return o
}

func (o MediaSpeechRecognitionTemplateArrayOutput) Index(i pulumi.IntInput) MediaSpeechRecognitionTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaSpeechRecognitionTemplate {
		return vs[0].([]*MediaSpeechRecognitionTemplate)[vs[1].(int)]
	}).(MediaSpeechRecognitionTemplateOutput)
}

type MediaSpeechRecognitionTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaSpeechRecognitionTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSpeechRecognitionTemplate)(nil)).Elem()
}

func (o MediaSpeechRecognitionTemplateMapOutput) ToMediaSpeechRecognitionTemplateMapOutput() MediaSpeechRecognitionTemplateMapOutput {
	return o
}

func (o MediaSpeechRecognitionTemplateMapOutput) ToMediaSpeechRecognitionTemplateMapOutputWithContext(ctx context.Context) MediaSpeechRecognitionTemplateMapOutput {
	return o
}

func (o MediaSpeechRecognitionTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaSpeechRecognitionTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaSpeechRecognitionTemplate {
		return vs[0].(map[string]*MediaSpeechRecognitionTemplate)[vs[1].(string)]
	}).(MediaSpeechRecognitionTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSpeechRecognitionTemplateInput)(nil)).Elem(), &MediaSpeechRecognitionTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSpeechRecognitionTemplateArrayInput)(nil)).Elem(), MediaSpeechRecognitionTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSpeechRecognitionTemplateMapInput)(nil)).Elem(), MediaSpeechRecognitionTemplateMap{})
	pulumi.RegisterOutputType(MediaSpeechRecognitionTemplateOutput{})
	pulumi.RegisterOutputType(MediaSpeechRecognitionTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaSpeechRecognitionTemplateMapOutput{})
}
