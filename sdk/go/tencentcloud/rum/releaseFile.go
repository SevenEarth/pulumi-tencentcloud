// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReleaseFile struct {
	pulumi.CustomResourceState

	// Release file hash.
	FileHash pulumi.StringOutput `pulumi:"fileHash"`
	// Release file unique key.
	FileKey pulumi.StringOutput `pulumi:"fileKey"`
	// Release file name.
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// Project ID.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Release file id.
	ReleaseFileId pulumi.IntOutput `pulumi:"releaseFileId"`
	// Release File version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewReleaseFile registers a new resource with the given unique name, arguments, and options.
func NewReleaseFile(ctx *pulumi.Context,
	name string, args *ReleaseFileArgs, opts ...pulumi.ResourceOption) (*ReleaseFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileHash == nil {
		return nil, errors.New("invalid value for required argument 'FileHash'")
	}
	if args.FileKey == nil {
		return nil, errors.New("invalid value for required argument 'FileKey'")
	}
	if args.FileName == nil {
		return nil, errors.New("invalid value for required argument 'FileName'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ReleaseFileId == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseFileId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReleaseFile
	err := ctx.RegisterResource("tencentcloud:Rum/releaseFile:ReleaseFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseFile gets an existing ReleaseFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseFileState, opts ...pulumi.ResourceOption) (*ReleaseFile, error) {
	var resource ReleaseFile
	err := ctx.ReadResource("tencentcloud:Rum/releaseFile:ReleaseFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseFile resources.
type releaseFileState struct {
	// Release file hash.
	FileHash *string `pulumi:"fileHash"`
	// Release file unique key.
	FileKey *string `pulumi:"fileKey"`
	// Release file name.
	FileName *string `pulumi:"fileName"`
	// Project ID.
	ProjectId *int `pulumi:"projectId"`
	// Release file id.
	ReleaseFileId *int `pulumi:"releaseFileId"`
	// Release File version.
	Version *string `pulumi:"version"`
}

type ReleaseFileState struct {
	// Release file hash.
	FileHash pulumi.StringPtrInput
	// Release file unique key.
	FileKey pulumi.StringPtrInput
	// Release file name.
	FileName pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.IntPtrInput
	// Release file id.
	ReleaseFileId pulumi.IntPtrInput
	// Release File version.
	Version pulumi.StringPtrInput
}

func (ReleaseFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseFileState)(nil)).Elem()
}

type releaseFileArgs struct {
	// Release file hash.
	FileHash string `pulumi:"fileHash"`
	// Release file unique key.
	FileKey string `pulumi:"fileKey"`
	// Release file name.
	FileName string `pulumi:"fileName"`
	// Project ID.
	ProjectId int `pulumi:"projectId"`
	// Release file id.
	ReleaseFileId int `pulumi:"releaseFileId"`
	// Release File version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a ReleaseFile resource.
type ReleaseFileArgs struct {
	// Release file hash.
	FileHash pulumi.StringInput
	// Release file unique key.
	FileKey pulumi.StringInput
	// Release file name.
	FileName pulumi.StringInput
	// Project ID.
	ProjectId pulumi.IntInput
	// Release file id.
	ReleaseFileId pulumi.IntInput
	// Release File version.
	Version pulumi.StringInput
}

func (ReleaseFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseFileArgs)(nil)).Elem()
}

type ReleaseFileInput interface {
	pulumi.Input

	ToReleaseFileOutput() ReleaseFileOutput
	ToReleaseFileOutputWithContext(ctx context.Context) ReleaseFileOutput
}

func (*ReleaseFile) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseFile)(nil)).Elem()
}

func (i *ReleaseFile) ToReleaseFileOutput() ReleaseFileOutput {
	return i.ToReleaseFileOutputWithContext(context.Background())
}

func (i *ReleaseFile) ToReleaseFileOutputWithContext(ctx context.Context) ReleaseFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseFileOutput)
}

// ReleaseFileArrayInput is an input type that accepts ReleaseFileArray and ReleaseFileArrayOutput values.
// You can construct a concrete instance of `ReleaseFileArrayInput` via:
//
//	ReleaseFileArray{ ReleaseFileArgs{...} }
type ReleaseFileArrayInput interface {
	pulumi.Input

	ToReleaseFileArrayOutput() ReleaseFileArrayOutput
	ToReleaseFileArrayOutputWithContext(context.Context) ReleaseFileArrayOutput
}

type ReleaseFileArray []ReleaseFileInput

func (ReleaseFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseFile)(nil)).Elem()
}

func (i ReleaseFileArray) ToReleaseFileArrayOutput() ReleaseFileArrayOutput {
	return i.ToReleaseFileArrayOutputWithContext(context.Background())
}

func (i ReleaseFileArray) ToReleaseFileArrayOutputWithContext(ctx context.Context) ReleaseFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseFileArrayOutput)
}

// ReleaseFileMapInput is an input type that accepts ReleaseFileMap and ReleaseFileMapOutput values.
// You can construct a concrete instance of `ReleaseFileMapInput` via:
//
//	ReleaseFileMap{ "key": ReleaseFileArgs{...} }
type ReleaseFileMapInput interface {
	pulumi.Input

	ToReleaseFileMapOutput() ReleaseFileMapOutput
	ToReleaseFileMapOutputWithContext(context.Context) ReleaseFileMapOutput
}

type ReleaseFileMap map[string]ReleaseFileInput

func (ReleaseFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseFile)(nil)).Elem()
}

func (i ReleaseFileMap) ToReleaseFileMapOutput() ReleaseFileMapOutput {
	return i.ToReleaseFileMapOutputWithContext(context.Background())
}

func (i ReleaseFileMap) ToReleaseFileMapOutputWithContext(ctx context.Context) ReleaseFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseFileMapOutput)
}

type ReleaseFileOutput struct{ *pulumi.OutputState }

func (ReleaseFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseFile)(nil)).Elem()
}

func (o ReleaseFileOutput) ToReleaseFileOutput() ReleaseFileOutput {
	return o
}

func (o ReleaseFileOutput) ToReleaseFileOutputWithContext(ctx context.Context) ReleaseFileOutput {
	return o
}

// Release file hash.
func (o ReleaseFileOutput) FileHash() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.StringOutput { return v.FileHash }).(pulumi.StringOutput)
}

// Release file unique key.
func (o ReleaseFileOutput) FileKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.StringOutput { return v.FileKey }).(pulumi.StringOutput)
}

// Release file name.
func (o ReleaseFileOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// Project ID.
func (o ReleaseFileOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Release file id.
func (o ReleaseFileOutput) ReleaseFileId() pulumi.IntOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.IntOutput { return v.ReleaseFileId }).(pulumi.IntOutput)
}

// Release File version.
func (o ReleaseFileOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseFile) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ReleaseFileArrayOutput struct{ *pulumi.OutputState }

func (ReleaseFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseFile)(nil)).Elem()
}

func (o ReleaseFileArrayOutput) ToReleaseFileArrayOutput() ReleaseFileArrayOutput {
	return o
}

func (o ReleaseFileArrayOutput) ToReleaseFileArrayOutputWithContext(ctx context.Context) ReleaseFileArrayOutput {
	return o
}

func (o ReleaseFileArrayOutput) Index(i pulumi.IntInput) ReleaseFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseFile {
		return vs[0].([]*ReleaseFile)[vs[1].(int)]
	}).(ReleaseFileOutput)
}

type ReleaseFileMapOutput struct{ *pulumi.OutputState }

func (ReleaseFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseFile)(nil)).Elem()
}

func (o ReleaseFileMapOutput) ToReleaseFileMapOutput() ReleaseFileMapOutput {
	return o
}

func (o ReleaseFileMapOutput) ToReleaseFileMapOutputWithContext(ctx context.Context) ReleaseFileMapOutput {
	return o
}

func (o ReleaseFileMapOutput) MapIndex(k pulumi.StringInput) ReleaseFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseFile {
		return vs[0].(map[string]*ReleaseFile)[vs[1].(string)]
	}).(ReleaseFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseFileInput)(nil)).Elem(), &ReleaseFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseFileArrayInput)(nil)).Elem(), ReleaseFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseFileMapInput)(nil)).Elem(), ReleaseFileMap{})
	pulumi.RegisterOutputType(ReleaseFileOutput{})
	pulumi.RegisterOutputType(ReleaseFileArrayOutput{})
	pulumi.RegisterOutputType(ReleaseFileMapOutput{})
}
