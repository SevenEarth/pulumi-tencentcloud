// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cloud file system(CFS).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfs.NewFileSystem(ctx, "foo", &Cfs.FileSystemArgs{
//				AccessGroupId:    pulumi.String("pgroup-7nx89k7l"),
//				AvailabilityZone: pulumi.String("ap-guangzhou-3"),
//				Protocol:         pulumi.String("NFS"),
//				SubnetId:         pulumi.String("subnet-9mu2t9iw"),
//				VpcId:            pulumi.String("vpc-ah9fbkap"),
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
// Cloud file system can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cfs/fileSystem:FileSystem foo cfs-6hgquxmj
//
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// ID of a access group.
	AccessGroupId pulumi.StringOutput `pulumi:"accessGroupId"`
	// The available zone that the file system locates at.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Create time of the file system.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// IP of mount point.
	MountIp pulumi.StringOutput `pulumi:"mountIp"`
	// Name of a file system.
	Name pulumi.StringOutput `pulumi:"name"`
	// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// ID of a subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Instance tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// ID of a VPC network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupId'")
	}
	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FileSystem
	err := ctx.RegisterResource("tencentcloud:Cfs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("tencentcloud:Cfs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// ID of a access group.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// The available zone that the file system locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Create time of the file system.
	CreateTime *string `pulumi:"createTime"`
	// IP of mount point.
	MountIp *string `pulumi:"mountIp"`
	// Name of a file system.
	Name *string `pulumi:"name"`
	// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
	Protocol *string `pulumi:"protocol"`
	// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
	StorageType *string `pulumi:"storageType"`
	// ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of a VPC network.
	VpcId *string `pulumi:"vpcId"`
}

type FileSystemState struct {
	// ID of a access group.
	AccessGroupId pulumi.StringPtrInput
	// The available zone that the file system locates at.
	AvailabilityZone pulumi.StringPtrInput
	// Create time of the file system.
	CreateTime pulumi.StringPtrInput
	// IP of mount point.
	MountIp pulumi.StringPtrInput
	// Name of a file system.
	Name pulumi.StringPtrInput
	// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
	Protocol pulumi.StringPtrInput
	// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
	StorageType pulumi.StringPtrInput
	// ID of a subnet.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// ID of a VPC network.
	VpcId pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// ID of a access group.
	AccessGroupId string `pulumi:"accessGroupId"`
	// The available zone that the file system locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// IP of mount point.
	MountIp *string `pulumi:"mountIp"`
	// Name of a file system.
	Name *string `pulumi:"name"`
	// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
	Protocol *string `pulumi:"protocol"`
	// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
	StorageType *string `pulumi:"storageType"`
	// ID of a subnet.
	SubnetId string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of a VPC network.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// ID of a access group.
	AccessGroupId pulumi.StringInput
	// The available zone that the file system locates at.
	AvailabilityZone pulumi.StringInput
	// IP of mount point.
	MountIp pulumi.StringPtrInput
	// Name of a file system.
	Name pulumi.StringPtrInput
	// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
	Protocol pulumi.StringPtrInput
	// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
	StorageType pulumi.StringPtrInput
	// ID of a subnet.
	SubnetId pulumi.StringInput
	// Instance tags.
	Tags pulumi.MapInput
	// ID of a VPC network.
	VpcId pulumi.StringInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//	FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//	FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct{ *pulumi.OutputState }

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

// ID of a access group.
func (o FileSystemOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.AccessGroupId }).(pulumi.StringOutput)
}

// The available zone that the file system locates at.
func (o FileSystemOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Create time of the file system.
func (o FileSystemOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// IP of mount point.
func (o FileSystemOutput) MountIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.MountIp }).(pulumi.StringOutput)
}

// Name of a file system.
func (o FileSystemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
func (o FileSystemOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
func (o FileSystemOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.StorageType }).(pulumi.StringPtrOutput)
}

// ID of a subnet.
func (o FileSystemOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Instance tags.
func (o FileSystemOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// ID of a VPC network.
func (o FileSystemOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].([]*FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].(map[string]*FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemArrayInput)(nil)).Elem(), FileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemMapInput)(nil)).Elem(), FileSystemMap{})
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
