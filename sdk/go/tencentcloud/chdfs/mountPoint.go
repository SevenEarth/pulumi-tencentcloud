// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chdfs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a chdfs mountPoint
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Chdfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Chdfs.NewMountPoint(ctx, "mountPoint", &Chdfs.MountPointArgs{
//				FileSystemId:     pulumi.String("f14mpfy5lh4e"),
//				MountPointName:   pulumi.String("terraform-test"),
//				MountPointStatus: pulumi.Int(1),
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
// chdfs mount_point can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Chdfs/mountPoint:MountPoint mount_point mount_point_id
// ```
type MountPoint struct {
	pulumi.CustomResourceState

	// file system id you want to mount.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// mount point name.
	MountPointName pulumi.StringOutput `pulumi:"mountPointName"`
	// mount status 1:open, 2:close.
	MountPointStatus pulumi.IntOutput `pulumi:"mountPointStatus"`
}

// NewMountPoint registers a new resource with the given unique name, arguments, and options.
func NewMountPoint(ctx *pulumi.Context,
	name string, args *MountPointArgs, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.MountPointName == nil {
		return nil, errors.New("invalid value for required argument 'MountPointName'")
	}
	if args.MountPointStatus == nil {
		return nil, errors.New("invalid value for required argument 'MountPointStatus'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MountPoint
	err := ctx.RegisterResource("tencentcloud:Chdfs/mountPoint:MountPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMountPoint gets an existing MountPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMountPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountPointState, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	var resource MountPoint
	err := ctx.ReadResource("tencentcloud:Chdfs/mountPoint:MountPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MountPoint resources.
type mountPointState struct {
	// file system id you want to mount.
	FileSystemId *string `pulumi:"fileSystemId"`
	// mount point name.
	MountPointName *string `pulumi:"mountPointName"`
	// mount status 1:open, 2:close.
	MountPointStatus *int `pulumi:"mountPointStatus"`
}

type MountPointState struct {
	// file system id you want to mount.
	FileSystemId pulumi.StringPtrInput
	// mount point name.
	MountPointName pulumi.StringPtrInput
	// mount status 1:open, 2:close.
	MountPointStatus pulumi.IntPtrInput
}

func (MountPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointState)(nil)).Elem()
}

type mountPointArgs struct {
	// file system id you want to mount.
	FileSystemId string `pulumi:"fileSystemId"`
	// mount point name.
	MountPointName string `pulumi:"mountPointName"`
	// mount status 1:open, 2:close.
	MountPointStatus int `pulumi:"mountPointStatus"`
}

// The set of arguments for constructing a MountPoint resource.
type MountPointArgs struct {
	// file system id you want to mount.
	FileSystemId pulumi.StringInput
	// mount point name.
	MountPointName pulumi.StringInput
	// mount status 1:open, 2:close.
	MountPointStatus pulumi.IntInput
}

func (MountPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointArgs)(nil)).Elem()
}

type MountPointInput interface {
	pulumi.Input

	ToMountPointOutput() MountPointOutput
	ToMountPointOutputWithContext(ctx context.Context) MountPointOutput
}

func (*MountPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (i *MountPoint) ToMountPointOutput() MountPointOutput {
	return i.ToMountPointOutputWithContext(context.Background())
}

func (i *MountPoint) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointOutput)
}

// MountPointArrayInput is an input type that accepts MountPointArray and MountPointArrayOutput values.
// You can construct a concrete instance of `MountPointArrayInput` via:
//
//	MountPointArray{ MountPointArgs{...} }
type MountPointArrayInput interface {
	pulumi.Input

	ToMountPointArrayOutput() MountPointArrayOutput
	ToMountPointArrayOutputWithContext(context.Context) MountPointArrayOutput
}

type MountPointArray []MountPointInput

func (MountPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (i MountPointArray) ToMountPointArrayOutput() MountPointArrayOutput {
	return i.ToMountPointArrayOutputWithContext(context.Background())
}

func (i MountPointArray) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointArrayOutput)
}

// MountPointMapInput is an input type that accepts MountPointMap and MountPointMapOutput values.
// You can construct a concrete instance of `MountPointMapInput` via:
//
//	MountPointMap{ "key": MountPointArgs{...} }
type MountPointMapInput interface {
	pulumi.Input

	ToMountPointMapOutput() MountPointMapOutput
	ToMountPointMapOutputWithContext(context.Context) MountPointMapOutput
}

type MountPointMap map[string]MountPointInput

func (MountPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (i MountPointMap) ToMountPointMapOutput() MountPointMapOutput {
	return i.ToMountPointMapOutputWithContext(context.Background())
}

func (i MountPointMap) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapOutput)
}

type MountPointOutput struct{ *pulumi.OutputState }

func (MountPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (o MountPointOutput) ToMountPointOutput() MountPointOutput {
	return o
}

func (o MountPointOutput) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return o
}

// file system id you want to mount.
func (o MountPointOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// mount point name.
func (o MountPointOutput) MountPointName() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.MountPointName }).(pulumi.StringOutput)
}

// mount status 1:open, 2:close.
func (o MountPointOutput) MountPointStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.IntOutput { return v.MountPointStatus }).(pulumi.IntOutput)
}

type MountPointArrayOutput struct{ *pulumi.OutputState }

func (MountPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (o MountPointArrayOutput) ToMountPointArrayOutput() MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) Index(i pulumi.IntInput) MountPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].([]*MountPoint)[vs[1].(int)]
	}).(MountPointOutput)
}

type MountPointMapOutput struct{ *pulumi.OutputState }

func (MountPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (o MountPointMapOutput) ToMountPointMapOutput() MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) MapIndex(k pulumi.StringInput) MountPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].(map[string]*MountPoint)[vs[1].(string)]
	}).(MountPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointInput)(nil)).Elem(), &MountPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointArrayInput)(nil)).Elem(), MountPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointMapInput)(nil)).Elem(), MountPointMap{})
	pulumi.RegisterOutputType(MountPointOutput{})
	pulumi.RegisterOutputType(MountPointArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapOutput{})
}
