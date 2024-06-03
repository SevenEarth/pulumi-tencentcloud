// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

type ImportImage struct {
	pulumi.CustomResourceState

	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// Boot mode.
	BootMode pulumi.StringPtrOutput `pulumi:"bootMode"`
	// Dry run to check the parameters without performing the operation.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Whether to force import the image.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Image description.
	ImageDescription pulumi.StringPtrOutput `pulumi:"imageDescription"`
	// Image name.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// Address on COS where the image to be imported is stored.
	ImageUrl pulumi.StringOutput `pulumi:"imageUrl"`
	// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
	// license BYOL: Bring Your Own License.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecifications ImportImageTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
}

// NewImportImage registers a new resource with the given unique name, arguments, and options.
func NewImportImage(ctx *pulumi.Context,
	name string, args *ImportImageArgs, opts ...pulumi.ResourceOption) (*ImportImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Architecture == nil {
		return nil, errors.New("invalid value for required argument 'Architecture'")
	}
	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.ImageUrl == nil {
		return nil, errors.New("invalid value for required argument 'ImageUrl'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.OsVersion == nil {
		return nil, errors.New("invalid value for required argument 'OsVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImportImage
	err := ctx.RegisterResource("tencentcloud:Cvm/importImage:ImportImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImportImage gets an existing ImportImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImportImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImportImageState, opts ...pulumi.ResourceOption) (*ImportImage, error) {
	var resource ImportImage
	err := ctx.ReadResource("tencentcloud:Cvm/importImage:ImportImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImportImage resources.
type importImageState struct {
	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture *string `pulumi:"architecture"`
	// Boot mode.
	BootMode *string `pulumi:"bootMode"`
	// Dry run to check the parameters without performing the operation.
	DryRun *bool `pulumi:"dryRun"`
	// Whether to force import the image.
	Force *bool `pulumi:"force"`
	// Image description.
	ImageDescription *string `pulumi:"imageDescription"`
	// Image name.
	ImageName *string `pulumi:"imageName"`
	// Address on COS where the image to be imported is stored.
	ImageUrl *string `pulumi:"imageUrl"`
	// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
	// license BYOL: Bring Your Own License.
	LicenseType *string `pulumi:"licenseType"`
	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsType *string `pulumi:"osType"`
	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsVersion *string `pulumi:"osVersion"`
	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecifications []ImportImageTagSpecification `pulumi:"tagSpecifications"`
}

type ImportImageState struct {
	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture pulumi.StringPtrInput
	// Boot mode.
	BootMode pulumi.StringPtrInput
	// Dry run to check the parameters without performing the operation.
	DryRun pulumi.BoolPtrInput
	// Whether to force import the image.
	Force pulumi.BoolPtrInput
	// Image description.
	ImageDescription pulumi.StringPtrInput
	// Image name.
	ImageName pulumi.StringPtrInput
	// Address on COS where the image to be imported is stored.
	ImageUrl pulumi.StringPtrInput
	// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
	// license BYOL: Bring Your Own License.
	LicenseType pulumi.StringPtrInput
	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsType pulumi.StringPtrInput
	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsVersion pulumi.StringPtrInput
	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecifications ImportImageTagSpecificationArrayInput
}

func (ImportImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*importImageState)(nil)).Elem()
}

type importImageArgs struct {
	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture string `pulumi:"architecture"`
	// Boot mode.
	BootMode *string `pulumi:"bootMode"`
	// Dry run to check the parameters without performing the operation.
	DryRun *bool `pulumi:"dryRun"`
	// Whether to force import the image.
	Force *bool `pulumi:"force"`
	// Image description.
	ImageDescription *string `pulumi:"imageDescription"`
	// Image name.
	ImageName string `pulumi:"imageName"`
	// Address on COS where the image to be imported is stored.
	ImageUrl string `pulumi:"imageUrl"`
	// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
	// license BYOL: Bring Your Own License.
	LicenseType *string `pulumi:"licenseType"`
	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsType string `pulumi:"osType"`
	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsVersion string `pulumi:"osVersion"`
	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecifications []ImportImageTagSpecification `pulumi:"tagSpecifications"`
}

// The set of arguments for constructing a ImportImage resource.
type ImportImageArgs struct {
	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture pulumi.StringInput
	// Boot mode.
	BootMode pulumi.StringPtrInput
	// Dry run to check the parameters without performing the operation.
	DryRun pulumi.BoolPtrInput
	// Whether to force import the image.
	Force pulumi.BoolPtrInput
	// Image description.
	ImageDescription pulumi.StringPtrInput
	// Image name.
	ImageName pulumi.StringInput
	// Address on COS where the image to be imported is stored.
	ImageUrl pulumi.StringInput
	// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
	// license BYOL: Bring Your Own License.
	LicenseType pulumi.StringPtrInput
	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsType pulumi.StringInput
	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
	// systems.
	OsVersion pulumi.StringInput
	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecifications ImportImageTagSpecificationArrayInput
}

func (ImportImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*importImageArgs)(nil)).Elem()
}

type ImportImageInput interface {
	pulumi.Input

	ToImportImageOutput() ImportImageOutput
	ToImportImageOutputWithContext(ctx context.Context) ImportImageOutput
}

func (*ImportImage) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportImage)(nil)).Elem()
}

func (i *ImportImage) ToImportImageOutput() ImportImageOutput {
	return i.ToImportImageOutputWithContext(context.Background())
}

func (i *ImportImage) ToImportImageOutputWithContext(ctx context.Context) ImportImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportImageOutput)
}

// ImportImageArrayInput is an input type that accepts ImportImageArray and ImportImageArrayOutput values.
// You can construct a concrete instance of `ImportImageArrayInput` via:
//
//	ImportImageArray{ ImportImageArgs{...} }
type ImportImageArrayInput interface {
	pulumi.Input

	ToImportImageArrayOutput() ImportImageArrayOutput
	ToImportImageArrayOutputWithContext(context.Context) ImportImageArrayOutput
}

type ImportImageArray []ImportImageInput

func (ImportImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImportImage)(nil)).Elem()
}

func (i ImportImageArray) ToImportImageArrayOutput() ImportImageArrayOutput {
	return i.ToImportImageArrayOutputWithContext(context.Background())
}

func (i ImportImageArray) ToImportImageArrayOutputWithContext(ctx context.Context) ImportImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportImageArrayOutput)
}

// ImportImageMapInput is an input type that accepts ImportImageMap and ImportImageMapOutput values.
// You can construct a concrete instance of `ImportImageMapInput` via:
//
//	ImportImageMap{ "key": ImportImageArgs{...} }
type ImportImageMapInput interface {
	pulumi.Input

	ToImportImageMapOutput() ImportImageMapOutput
	ToImportImageMapOutputWithContext(context.Context) ImportImageMapOutput
}

type ImportImageMap map[string]ImportImageInput

func (ImportImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImportImage)(nil)).Elem()
}

func (i ImportImageMap) ToImportImageMapOutput() ImportImageMapOutput {
	return i.ToImportImageMapOutputWithContext(context.Background())
}

func (i ImportImageMap) ToImportImageMapOutputWithContext(ctx context.Context) ImportImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportImageMapOutput)
}

type ImportImageOutput struct{ *pulumi.OutputState }

func (ImportImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportImage)(nil)).Elem()
}

func (o ImportImageOutput) ToImportImageOutput() ImportImageOutput {
	return o
}

func (o ImportImageOutput) ToImportImageOutputWithContext(ctx context.Context) ImportImageOutput {
	return o
}

// OS architecture of the image to be imported, `x86_64` or `i386`.
func (o ImportImageOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringOutput { return v.Architecture }).(pulumi.StringOutput)
}

// Boot mode.
func (o ImportImageOutput) BootMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringPtrOutput { return v.BootMode }).(pulumi.StringPtrOutput)
}

// Dry run to check the parameters without performing the operation.
func (o ImportImageOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Whether to force import the image.
func (o ImportImageOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// Image description.
func (o ImportImageOutput) ImageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringPtrOutput { return v.ImageDescription }).(pulumi.StringPtrOutput)
}

// Image name.
func (o ImportImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// Address on COS where the image to be imported is stored.
func (o ImportImageOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringOutput { return v.ImageUrl }).(pulumi.StringOutput)
}

// The license type used to activate the OS after importing an image. Valid values: TencentCloud: Tencent Cloud official
// license BYOL: Bring Your Own License.
func (o ImportImageOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
// systems.
func (o ImportImageOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating
// systems.
func (o ImportImageOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportImage) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

// Tag description list. This parameter is used to bind a tag to a custom image.
func (o ImportImageOutput) TagSpecifications() ImportImageTagSpecificationArrayOutput {
	return o.ApplyT(func(v *ImportImage) ImportImageTagSpecificationArrayOutput { return v.TagSpecifications }).(ImportImageTagSpecificationArrayOutput)
}

type ImportImageArrayOutput struct{ *pulumi.OutputState }

func (ImportImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImportImage)(nil)).Elem()
}

func (o ImportImageArrayOutput) ToImportImageArrayOutput() ImportImageArrayOutput {
	return o
}

func (o ImportImageArrayOutput) ToImportImageArrayOutputWithContext(ctx context.Context) ImportImageArrayOutput {
	return o
}

func (o ImportImageArrayOutput) Index(i pulumi.IntInput) ImportImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImportImage {
		return vs[0].([]*ImportImage)[vs[1].(int)]
	}).(ImportImageOutput)
}

type ImportImageMapOutput struct{ *pulumi.OutputState }

func (ImportImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImportImage)(nil)).Elem()
}

func (o ImportImageMapOutput) ToImportImageMapOutput() ImportImageMapOutput {
	return o
}

func (o ImportImageMapOutput) ToImportImageMapOutputWithContext(ctx context.Context) ImportImageMapOutput {
	return o
}

func (o ImportImageMapOutput) MapIndex(k pulumi.StringInput) ImportImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImportImage {
		return vs[0].(map[string]*ImportImage)[vs[1].(string)]
	}).(ImportImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImportImageInput)(nil)).Elem(), &ImportImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportImageArrayInput)(nil)).Elem(), ImportImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportImageMapInput)(nil)).Elem(), ImportImageMap{})
	pulumi.RegisterOutputType(ImportImageOutput{})
	pulumi.RegisterOutputType(ImportImageArrayOutput{})
	pulumi.RegisterOutputType(ImportImageMapOutput{})
}
