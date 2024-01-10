// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfw

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cfw syncAsset
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfw"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfw.NewSyncAsset(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SyncAsset struct {
	pulumi.CustomResourceState
}

// NewSyncAsset registers a new resource with the given unique name, arguments, and options.
func NewSyncAsset(ctx *pulumi.Context,
	name string, args *SyncAssetArgs, opts ...pulumi.ResourceOption) (*SyncAsset, error) {
	if args == nil {
		args = &SyncAssetArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SyncAsset
	err := ctx.RegisterResource("tencentcloud:Cfw/syncAsset:SyncAsset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAsset gets an existing SyncAsset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAssetState, opts ...pulumi.ResourceOption) (*SyncAsset, error) {
	var resource SyncAsset
	err := ctx.ReadResource("tencentcloud:Cfw/syncAsset:SyncAsset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAsset resources.
type syncAssetState struct {
}

type SyncAssetState struct {
}

func (SyncAssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAssetState)(nil)).Elem()
}

type syncAssetArgs struct {
}

// The set of arguments for constructing a SyncAsset resource.
type SyncAssetArgs struct {
}

func (SyncAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAssetArgs)(nil)).Elem()
}

type SyncAssetInput interface {
	pulumi.Input

	ToSyncAssetOutput() SyncAssetOutput
	ToSyncAssetOutputWithContext(ctx context.Context) SyncAssetOutput
}

func (*SyncAsset) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAsset)(nil)).Elem()
}

func (i *SyncAsset) ToSyncAssetOutput() SyncAssetOutput {
	return i.ToSyncAssetOutputWithContext(context.Background())
}

func (i *SyncAsset) ToSyncAssetOutputWithContext(ctx context.Context) SyncAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssetOutput)
}

// SyncAssetArrayInput is an input type that accepts SyncAssetArray and SyncAssetArrayOutput values.
// You can construct a concrete instance of `SyncAssetArrayInput` via:
//
//	SyncAssetArray{ SyncAssetArgs{...} }
type SyncAssetArrayInput interface {
	pulumi.Input

	ToSyncAssetArrayOutput() SyncAssetArrayOutput
	ToSyncAssetArrayOutputWithContext(context.Context) SyncAssetArrayOutput
}

type SyncAssetArray []SyncAssetInput

func (SyncAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAsset)(nil)).Elem()
}

func (i SyncAssetArray) ToSyncAssetArrayOutput() SyncAssetArrayOutput {
	return i.ToSyncAssetArrayOutputWithContext(context.Background())
}

func (i SyncAssetArray) ToSyncAssetArrayOutputWithContext(ctx context.Context) SyncAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssetArrayOutput)
}

// SyncAssetMapInput is an input type that accepts SyncAssetMap and SyncAssetMapOutput values.
// You can construct a concrete instance of `SyncAssetMapInput` via:
//
//	SyncAssetMap{ "key": SyncAssetArgs{...} }
type SyncAssetMapInput interface {
	pulumi.Input

	ToSyncAssetMapOutput() SyncAssetMapOutput
	ToSyncAssetMapOutputWithContext(context.Context) SyncAssetMapOutput
}

type SyncAssetMap map[string]SyncAssetInput

func (SyncAssetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAsset)(nil)).Elem()
}

func (i SyncAssetMap) ToSyncAssetMapOutput() SyncAssetMapOutput {
	return i.ToSyncAssetMapOutputWithContext(context.Background())
}

func (i SyncAssetMap) ToSyncAssetMapOutputWithContext(ctx context.Context) SyncAssetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAssetMapOutput)
}

type SyncAssetOutput struct{ *pulumi.OutputState }

func (SyncAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAsset)(nil)).Elem()
}

func (o SyncAssetOutput) ToSyncAssetOutput() SyncAssetOutput {
	return o
}

func (o SyncAssetOutput) ToSyncAssetOutputWithContext(ctx context.Context) SyncAssetOutput {
	return o
}

type SyncAssetArrayOutput struct{ *pulumi.OutputState }

func (SyncAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncAsset)(nil)).Elem()
}

func (o SyncAssetArrayOutput) ToSyncAssetArrayOutput() SyncAssetArrayOutput {
	return o
}

func (o SyncAssetArrayOutput) ToSyncAssetArrayOutputWithContext(ctx context.Context) SyncAssetArrayOutput {
	return o
}

func (o SyncAssetArrayOutput) Index(i pulumi.IntInput) SyncAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncAsset {
		return vs[0].([]*SyncAsset)[vs[1].(int)]
	}).(SyncAssetOutput)
}

type SyncAssetMapOutput struct{ *pulumi.OutputState }

func (SyncAssetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncAsset)(nil)).Elem()
}

func (o SyncAssetMapOutput) ToSyncAssetMapOutput() SyncAssetMapOutput {
	return o
}

func (o SyncAssetMapOutput) ToSyncAssetMapOutputWithContext(ctx context.Context) SyncAssetMapOutput {
	return o
}

func (o SyncAssetMapOutput) MapIndex(k pulumi.StringInput) SyncAssetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncAsset {
		return vs[0].(map[string]*SyncAsset)[vs[1].(string)]
	}).(SyncAssetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssetInput)(nil)).Elem(), &SyncAsset{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssetArrayInput)(nil)).Elem(), SyncAssetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncAssetMapInput)(nil)).Elem(), SyncAssetMap{})
	pulumi.RegisterOutputType(SyncAssetOutput{})
	pulumi.RegisterOutputType(SyncAssetArrayOutput{})
	pulumi.RegisterOutputType(SyncAssetMapOutput{})
}
