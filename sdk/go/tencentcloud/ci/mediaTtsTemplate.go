// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ci mediaTtsTemplate
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
// 		_, err := Ci.NewMediaTtsTemplate(ctx, "mediaTtsTemplate", &Ci.MediaTtsTemplateArgs{
// 			Bucket:    pulumi.String("terraform-ci-xxxxxx"),
// 			Codec:     pulumi.String("pcm"),
// 			Mode:      pulumi.String("Asyc"),
// 			Speed:     pulumi.String("100"),
// 			VoiceType: pulumi.String("ruxue"),
// 			Volume:    pulumi.String("0"),
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
// ci media_tts_template can be imported using the bucket#templateId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate media_tts_template terraform-ci-xxxxxx#t1ed421df8bd2140b6b73474f70f99b0f8
// ```
type MediaTtsTemplate struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
	Codec pulumi.StringPtrOutput `pulumi:"codec"`
	// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Speech rate, the default value is 100, [50,200].
	Speed pulumi.StringPtrOutput `pulumi:"speed"`
	// Timbre, the default value is ruxue.
	VoiceType pulumi.StringPtrOutput `pulumi:"voiceType"`
	// Volume, default value 0, [-10,10].
	Volume pulumi.StringPtrOutput `pulumi:"volume"`
}

// NewMediaTtsTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaTtsTemplate(ctx *pulumi.Context,
	name string, args *MediaTtsTemplateArgs, opts ...pulumi.ResourceOption) (*MediaTtsTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MediaTtsTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaTtsTemplate gets an existing MediaTtsTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaTtsTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaTtsTemplateState, opts ...pulumi.ResourceOption) (*MediaTtsTemplate, error) {
	var resource MediaTtsTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaTtsTemplate:MediaTtsTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaTtsTemplate resources.
type mediaTtsTemplateState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
	Codec *string `pulumi:"codec"`
	// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
	Mode *string `pulumi:"mode"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// Speech rate, the default value is 100, [50,200].
	Speed *string `pulumi:"speed"`
	// Timbre, the default value is ruxue.
	VoiceType *string `pulumi:"voiceType"`
	// Volume, default value 0, [-10,10].
	Volume *string `pulumi:"volume"`
}

type MediaTtsTemplateState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
	Codec pulumi.StringPtrInput
	// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
	Mode pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// Speech rate, the default value is 100, [50,200].
	Speed pulumi.StringPtrInput
	// Timbre, the default value is ruxue.
	VoiceType pulumi.StringPtrInput
	// Volume, default value 0, [-10,10].
	Volume pulumi.StringPtrInput
}

func (MediaTtsTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaTtsTemplateState)(nil)).Elem()
}

type mediaTtsTemplateArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
	Codec *string `pulumi:"codec"`
	// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
	Mode *string `pulumi:"mode"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// Speech rate, the default value is 100, [50,200].
	Speed *string `pulumi:"speed"`
	// Timbre, the default value is ruxue.
	VoiceType *string `pulumi:"voiceType"`
	// Volume, default value 0, [-10,10].
	Volume *string `pulumi:"volume"`
}

// The set of arguments for constructing a MediaTtsTemplate resource.
type MediaTtsTemplateArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
	Codec pulumi.StringPtrInput
	// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
	Mode pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// Speech rate, the default value is 100, [50,200].
	Speed pulumi.StringPtrInput
	// Timbre, the default value is ruxue.
	VoiceType pulumi.StringPtrInput
	// Volume, default value 0, [-10,10].
	Volume pulumi.StringPtrInput
}

func (MediaTtsTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaTtsTemplateArgs)(nil)).Elem()
}

type MediaTtsTemplateInput interface {
	pulumi.Input

	ToMediaTtsTemplateOutput() MediaTtsTemplateOutput
	ToMediaTtsTemplateOutputWithContext(ctx context.Context) MediaTtsTemplateOutput
}

func (*MediaTtsTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaTtsTemplate)(nil)).Elem()
}

func (i *MediaTtsTemplate) ToMediaTtsTemplateOutput() MediaTtsTemplateOutput {
	return i.ToMediaTtsTemplateOutputWithContext(context.Background())
}

func (i *MediaTtsTemplate) ToMediaTtsTemplateOutputWithContext(ctx context.Context) MediaTtsTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaTtsTemplateOutput)
}

// MediaTtsTemplateArrayInput is an input type that accepts MediaTtsTemplateArray and MediaTtsTemplateArrayOutput values.
// You can construct a concrete instance of `MediaTtsTemplateArrayInput` via:
//
//          MediaTtsTemplateArray{ MediaTtsTemplateArgs{...} }
type MediaTtsTemplateArrayInput interface {
	pulumi.Input

	ToMediaTtsTemplateArrayOutput() MediaTtsTemplateArrayOutput
	ToMediaTtsTemplateArrayOutputWithContext(context.Context) MediaTtsTemplateArrayOutput
}

type MediaTtsTemplateArray []MediaTtsTemplateInput

func (MediaTtsTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaTtsTemplate)(nil)).Elem()
}

func (i MediaTtsTemplateArray) ToMediaTtsTemplateArrayOutput() MediaTtsTemplateArrayOutput {
	return i.ToMediaTtsTemplateArrayOutputWithContext(context.Background())
}

func (i MediaTtsTemplateArray) ToMediaTtsTemplateArrayOutputWithContext(ctx context.Context) MediaTtsTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaTtsTemplateArrayOutput)
}

// MediaTtsTemplateMapInput is an input type that accepts MediaTtsTemplateMap and MediaTtsTemplateMapOutput values.
// You can construct a concrete instance of `MediaTtsTemplateMapInput` via:
//
//          MediaTtsTemplateMap{ "key": MediaTtsTemplateArgs{...} }
type MediaTtsTemplateMapInput interface {
	pulumi.Input

	ToMediaTtsTemplateMapOutput() MediaTtsTemplateMapOutput
	ToMediaTtsTemplateMapOutputWithContext(context.Context) MediaTtsTemplateMapOutput
}

type MediaTtsTemplateMap map[string]MediaTtsTemplateInput

func (MediaTtsTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaTtsTemplate)(nil)).Elem()
}

func (i MediaTtsTemplateMap) ToMediaTtsTemplateMapOutput() MediaTtsTemplateMapOutput {
	return i.ToMediaTtsTemplateMapOutputWithContext(context.Background())
}

func (i MediaTtsTemplateMap) ToMediaTtsTemplateMapOutputWithContext(ctx context.Context) MediaTtsTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaTtsTemplateMapOutput)
}

type MediaTtsTemplateOutput struct{ *pulumi.OutputState }

func (MediaTtsTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaTtsTemplate)(nil)).Elem()
}

func (o MediaTtsTemplateOutput) ToMediaTtsTemplateOutput() MediaTtsTemplateOutput {
	return o
}

func (o MediaTtsTemplateOutput) ToMediaTtsTemplateOutputWithContext(ctx context.Context) MediaTtsTemplateOutput {
	return o
}

// bucket name.
func (o MediaTtsTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Audio format, default wav (synchronous)/pcm (asynchronous, wav, mp3, pcm.
func (o MediaTtsTemplateOutput) Codec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringPtrOutput { return v.Codec }).(pulumi.StringPtrOutput)
}

// Processing mode, default value Asyc, Asyc (asynchronous composition), Sync (synchronous composition), When Asyc is selected, the codec only supports pcm.
func (o MediaTtsTemplateOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaTtsTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Speech rate, the default value is 100, [50,200].
func (o MediaTtsTemplateOutput) Speed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringPtrOutput { return v.Speed }).(pulumi.StringPtrOutput)
}

// Timbre, the default value is ruxue.
func (o MediaTtsTemplateOutput) VoiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringPtrOutput { return v.VoiceType }).(pulumi.StringPtrOutput)
}

// Volume, default value 0, [-10,10].
func (o MediaTtsTemplateOutput) Volume() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaTtsTemplate) pulumi.StringPtrOutput { return v.Volume }).(pulumi.StringPtrOutput)
}

type MediaTtsTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaTtsTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaTtsTemplate)(nil)).Elem()
}

func (o MediaTtsTemplateArrayOutput) ToMediaTtsTemplateArrayOutput() MediaTtsTemplateArrayOutput {
	return o
}

func (o MediaTtsTemplateArrayOutput) ToMediaTtsTemplateArrayOutputWithContext(ctx context.Context) MediaTtsTemplateArrayOutput {
	return o
}

func (o MediaTtsTemplateArrayOutput) Index(i pulumi.IntInput) MediaTtsTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaTtsTemplate {
		return vs[0].([]*MediaTtsTemplate)[vs[1].(int)]
	}).(MediaTtsTemplateOutput)
}

type MediaTtsTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaTtsTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaTtsTemplate)(nil)).Elem()
}

func (o MediaTtsTemplateMapOutput) ToMediaTtsTemplateMapOutput() MediaTtsTemplateMapOutput {
	return o
}

func (o MediaTtsTemplateMapOutput) ToMediaTtsTemplateMapOutputWithContext(ctx context.Context) MediaTtsTemplateMapOutput {
	return o
}

func (o MediaTtsTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaTtsTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaTtsTemplate {
		return vs[0].(map[string]*MediaTtsTemplate)[vs[1].(string)]
	}).(MediaTtsTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaTtsTemplateInput)(nil)).Elem(), &MediaTtsTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaTtsTemplateArrayInput)(nil)).Elem(), MediaTtsTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaTtsTemplateMapInput)(nil)).Elem(), MediaTtsTemplateMap{})
	pulumi.RegisterOutputType(MediaTtsTemplateOutput{})
	pulumi.RegisterOutputType(MediaTtsTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaTtsTemplateMapOutput{})
}
