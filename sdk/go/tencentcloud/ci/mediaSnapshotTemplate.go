// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ci mediaSnapshotTemplate
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
// 		_, err := Ci.NewMediaSnapshotTemplate(ctx, "mediaSnapshotTemplate", &Ci.MediaSnapshotTemplateArgs{
// 			Bucket: pulumi.String("terraform-ci-xxxxxx"),
// 			Snapshot: &ci.MediaSnapshotTemplateSnapshotArgs{
// 				Count:           pulumi.String("10"),
// 				SnapshotOutMode: pulumi.String("SnapshotAndSprite"),
// 				SpriteSnapshotConfig: &ci.MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs{
// 					Color:   pulumi.String("White"),
// 					Columns: pulumi.String("10"),
// 					Lines:   pulumi.String("10"),
// 					Margin:  pulumi.String("10"),
// 					Padding: pulumi.String("10"),
// 				},
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
// ci media_snapshot_template can be imported using the bucket#templateId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate media_snapshot_template terraform-ci-xxxxxx#t18210645f96564eaf80e86b1f58c20152
// ```
type MediaSnapshotTemplate struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// screenshot.
	Snapshot MediaSnapshotTemplateSnapshotOutput `pulumi:"snapshot"`
	// Template ID.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewMediaSnapshotTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaSnapshotTemplate(ctx *pulumi.Context,
	name string, args *MediaSnapshotTemplateArgs, opts ...pulumi.ResourceOption) (*MediaSnapshotTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Snapshot == nil {
		return nil, errors.New("invalid value for required argument 'Snapshot'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MediaSnapshotTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaSnapshotTemplate gets an existing MediaSnapshotTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaSnapshotTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaSnapshotTemplateState, opts ...pulumi.ResourceOption) (*MediaSnapshotTemplate, error) {
	var resource MediaSnapshotTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaSnapshotTemplate:MediaSnapshotTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaSnapshotTemplate resources.
type mediaSnapshotTemplateState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// creation time.
	CreateTime *string `pulumi:"createTime"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// screenshot.
	Snapshot *MediaSnapshotTemplateSnapshot `pulumi:"snapshot"`
	// Template ID.
	TemplateId *string `pulumi:"templateId"`
	// update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type MediaSnapshotTemplateState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// creation time.
	CreateTime pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// screenshot.
	Snapshot MediaSnapshotTemplateSnapshotPtrInput
	// Template ID.
	TemplateId pulumi.StringPtrInput
	// update time.
	UpdateTime pulumi.StringPtrInput
}

func (MediaSnapshotTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSnapshotTemplateState)(nil)).Elem()
}

type mediaSnapshotTemplateArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// screenshot.
	Snapshot MediaSnapshotTemplateSnapshot `pulumi:"snapshot"`
}

// The set of arguments for constructing a MediaSnapshotTemplate resource.
type MediaSnapshotTemplateArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// screenshot.
	Snapshot MediaSnapshotTemplateSnapshotInput
}

func (MediaSnapshotTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaSnapshotTemplateArgs)(nil)).Elem()
}

type MediaSnapshotTemplateInput interface {
	pulumi.Input

	ToMediaSnapshotTemplateOutput() MediaSnapshotTemplateOutput
	ToMediaSnapshotTemplateOutputWithContext(ctx context.Context) MediaSnapshotTemplateOutput
}

func (*MediaSnapshotTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSnapshotTemplate)(nil)).Elem()
}

func (i *MediaSnapshotTemplate) ToMediaSnapshotTemplateOutput() MediaSnapshotTemplateOutput {
	return i.ToMediaSnapshotTemplateOutputWithContext(context.Background())
}

func (i *MediaSnapshotTemplate) ToMediaSnapshotTemplateOutputWithContext(ctx context.Context) MediaSnapshotTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSnapshotTemplateOutput)
}

// MediaSnapshotTemplateArrayInput is an input type that accepts MediaSnapshotTemplateArray and MediaSnapshotTemplateArrayOutput values.
// You can construct a concrete instance of `MediaSnapshotTemplateArrayInput` via:
//
//          MediaSnapshotTemplateArray{ MediaSnapshotTemplateArgs{...} }
type MediaSnapshotTemplateArrayInput interface {
	pulumi.Input

	ToMediaSnapshotTemplateArrayOutput() MediaSnapshotTemplateArrayOutput
	ToMediaSnapshotTemplateArrayOutputWithContext(context.Context) MediaSnapshotTemplateArrayOutput
}

type MediaSnapshotTemplateArray []MediaSnapshotTemplateInput

func (MediaSnapshotTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSnapshotTemplate)(nil)).Elem()
}

func (i MediaSnapshotTemplateArray) ToMediaSnapshotTemplateArrayOutput() MediaSnapshotTemplateArrayOutput {
	return i.ToMediaSnapshotTemplateArrayOutputWithContext(context.Background())
}

func (i MediaSnapshotTemplateArray) ToMediaSnapshotTemplateArrayOutputWithContext(ctx context.Context) MediaSnapshotTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSnapshotTemplateArrayOutput)
}

// MediaSnapshotTemplateMapInput is an input type that accepts MediaSnapshotTemplateMap and MediaSnapshotTemplateMapOutput values.
// You can construct a concrete instance of `MediaSnapshotTemplateMapInput` via:
//
//          MediaSnapshotTemplateMap{ "key": MediaSnapshotTemplateArgs{...} }
type MediaSnapshotTemplateMapInput interface {
	pulumi.Input

	ToMediaSnapshotTemplateMapOutput() MediaSnapshotTemplateMapOutput
	ToMediaSnapshotTemplateMapOutputWithContext(context.Context) MediaSnapshotTemplateMapOutput
}

type MediaSnapshotTemplateMap map[string]MediaSnapshotTemplateInput

func (MediaSnapshotTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSnapshotTemplate)(nil)).Elem()
}

func (i MediaSnapshotTemplateMap) ToMediaSnapshotTemplateMapOutput() MediaSnapshotTemplateMapOutput {
	return i.ToMediaSnapshotTemplateMapOutputWithContext(context.Background())
}

func (i MediaSnapshotTemplateMap) ToMediaSnapshotTemplateMapOutputWithContext(ctx context.Context) MediaSnapshotTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaSnapshotTemplateMapOutput)
}

type MediaSnapshotTemplateOutput struct{ *pulumi.OutputState }

func (MediaSnapshotTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaSnapshotTemplate)(nil)).Elem()
}

func (o MediaSnapshotTemplateOutput) ToMediaSnapshotTemplateOutput() MediaSnapshotTemplateOutput {
	return o
}

func (o MediaSnapshotTemplateOutput) ToMediaSnapshotTemplateOutputWithContext(ctx context.Context) MediaSnapshotTemplateOutput {
	return o
}

// bucket name.
func (o MediaSnapshotTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// creation time.
func (o MediaSnapshotTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaSnapshotTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// screenshot.
func (o MediaSnapshotTemplateOutput) Snapshot() MediaSnapshotTemplateSnapshotOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) MediaSnapshotTemplateSnapshotOutput { return v.Snapshot }).(MediaSnapshotTemplateSnapshotOutput)
}

// Template ID.
func (o MediaSnapshotTemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// update time.
func (o MediaSnapshotTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaSnapshotTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type MediaSnapshotTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaSnapshotTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaSnapshotTemplate)(nil)).Elem()
}

func (o MediaSnapshotTemplateArrayOutput) ToMediaSnapshotTemplateArrayOutput() MediaSnapshotTemplateArrayOutput {
	return o
}

func (o MediaSnapshotTemplateArrayOutput) ToMediaSnapshotTemplateArrayOutputWithContext(ctx context.Context) MediaSnapshotTemplateArrayOutput {
	return o
}

func (o MediaSnapshotTemplateArrayOutput) Index(i pulumi.IntInput) MediaSnapshotTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaSnapshotTemplate {
		return vs[0].([]*MediaSnapshotTemplate)[vs[1].(int)]
	}).(MediaSnapshotTemplateOutput)
}

type MediaSnapshotTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaSnapshotTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaSnapshotTemplate)(nil)).Elem()
}

func (o MediaSnapshotTemplateMapOutput) ToMediaSnapshotTemplateMapOutput() MediaSnapshotTemplateMapOutput {
	return o
}

func (o MediaSnapshotTemplateMapOutput) ToMediaSnapshotTemplateMapOutputWithContext(ctx context.Context) MediaSnapshotTemplateMapOutput {
	return o
}

func (o MediaSnapshotTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaSnapshotTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaSnapshotTemplate {
		return vs[0].(map[string]*MediaSnapshotTemplate)[vs[1].(string)]
	}).(MediaSnapshotTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSnapshotTemplateInput)(nil)).Elem(), &MediaSnapshotTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSnapshotTemplateArrayInput)(nil)).Elem(), MediaSnapshotTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaSnapshotTemplateMapInput)(nil)).Elem(), MediaSnapshotTemplateMap{})
	pulumi.RegisterOutputType(MediaSnapshotTemplateOutput{})
	pulumi.RegisterOutputType(MediaSnapshotTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaSnapshotTemplateMapOutput{})
}
