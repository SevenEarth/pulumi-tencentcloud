// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cbs disk_backup.
//
// > **NOTE:** Backup quota must greater than 1.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cbs.NewDiskBackup(ctx, "diskBackup", &Cbs.DiskBackupArgs{
// 			DiskBackupName: pulumi.String("xxx"),
// 			DiskId:         pulumi.String("disk-xxx"),
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
// cbs disk_backup can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cbs/diskBackup:DiskBackup disk_backup disk_backup_id
// ```
type DiskBackup struct {
	pulumi.CustomResourceState

	// Backup point name.
	DiskBackupName pulumi.StringPtrOutput `pulumi:"diskBackupName"`
	// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
}

// NewDiskBackup registers a new resource with the given unique name, arguments, and options.
func NewDiskBackup(ctx *pulumi.Context,
	name string, args *DiskBackupArgs, opts ...pulumi.ResourceOption) (*DiskBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DiskBackup
	err := ctx.RegisterResource("tencentcloud:Cbs/diskBackup:DiskBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskBackup gets an existing DiskBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskBackupState, opts ...pulumi.ResourceOption) (*DiskBackup, error) {
	var resource DiskBackup
	err := ctx.ReadResource("tencentcloud:Cbs/diskBackup:DiskBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskBackup resources.
type diskBackupState struct {
	// Backup point name.
	DiskBackupName *string `pulumi:"diskBackupName"`
	// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
	DiskId *string `pulumi:"diskId"`
}

type DiskBackupState struct {
	// Backup point name.
	DiskBackupName pulumi.StringPtrInput
	// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
	DiskId pulumi.StringPtrInput
}

func (DiskBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskBackupState)(nil)).Elem()
}

type diskBackupArgs struct {
	// Backup point name.
	DiskBackupName *string `pulumi:"diskBackupName"`
	// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
	DiskId string `pulumi:"diskId"`
}

// The set of arguments for constructing a DiskBackup resource.
type DiskBackupArgs struct {
	// Backup point name.
	DiskBackupName pulumi.StringPtrInput
	// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
	DiskId pulumi.StringInput
}

func (DiskBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskBackupArgs)(nil)).Elem()
}

type DiskBackupInput interface {
	pulumi.Input

	ToDiskBackupOutput() DiskBackupOutput
	ToDiskBackupOutputWithContext(ctx context.Context) DiskBackupOutput
}

func (*DiskBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskBackup)(nil)).Elem()
}

func (i *DiskBackup) ToDiskBackupOutput() DiskBackupOutput {
	return i.ToDiskBackupOutputWithContext(context.Background())
}

func (i *DiskBackup) ToDiskBackupOutputWithContext(ctx context.Context) DiskBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupOutput)
}

// DiskBackupArrayInput is an input type that accepts DiskBackupArray and DiskBackupArrayOutput values.
// You can construct a concrete instance of `DiskBackupArrayInput` via:
//
//          DiskBackupArray{ DiskBackupArgs{...} }
type DiskBackupArrayInput interface {
	pulumi.Input

	ToDiskBackupArrayOutput() DiskBackupArrayOutput
	ToDiskBackupArrayOutputWithContext(context.Context) DiskBackupArrayOutput
}

type DiskBackupArray []DiskBackupInput

func (DiskBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskBackup)(nil)).Elem()
}

func (i DiskBackupArray) ToDiskBackupArrayOutput() DiskBackupArrayOutput {
	return i.ToDiskBackupArrayOutputWithContext(context.Background())
}

func (i DiskBackupArray) ToDiskBackupArrayOutputWithContext(ctx context.Context) DiskBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupArrayOutput)
}

// DiskBackupMapInput is an input type that accepts DiskBackupMap and DiskBackupMapOutput values.
// You can construct a concrete instance of `DiskBackupMapInput` via:
//
//          DiskBackupMap{ "key": DiskBackupArgs{...} }
type DiskBackupMapInput interface {
	pulumi.Input

	ToDiskBackupMapOutput() DiskBackupMapOutput
	ToDiskBackupMapOutputWithContext(context.Context) DiskBackupMapOutput
}

type DiskBackupMap map[string]DiskBackupInput

func (DiskBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskBackup)(nil)).Elem()
}

func (i DiskBackupMap) ToDiskBackupMapOutput() DiskBackupMapOutput {
	return i.ToDiskBackupMapOutputWithContext(context.Background())
}

func (i DiskBackupMap) ToDiskBackupMapOutputWithContext(ctx context.Context) DiskBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupMapOutput)
}

type DiskBackupOutput struct{ *pulumi.OutputState }

func (DiskBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskBackup)(nil)).Elem()
}

func (o DiskBackupOutput) ToDiskBackupOutput() DiskBackupOutput {
	return o
}

func (o DiskBackupOutput) ToDiskBackupOutputWithContext(ctx context.Context) DiskBackupOutput {
	return o
}

// Backup point name.
func (o DiskBackupOutput) DiskBackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskBackup) pulumi.StringPtrOutput { return v.DiskBackupName }).(pulumi.StringPtrOutput)
}

// ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
func (o DiskBackupOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskBackup) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

type DiskBackupArrayOutput struct{ *pulumi.OutputState }

func (DiskBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskBackup)(nil)).Elem()
}

func (o DiskBackupArrayOutput) ToDiskBackupArrayOutput() DiskBackupArrayOutput {
	return o
}

func (o DiskBackupArrayOutput) ToDiskBackupArrayOutputWithContext(ctx context.Context) DiskBackupArrayOutput {
	return o
}

func (o DiskBackupArrayOutput) Index(i pulumi.IntInput) DiskBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DiskBackup {
		return vs[0].([]*DiskBackup)[vs[1].(int)]
	}).(DiskBackupOutput)
}

type DiskBackupMapOutput struct{ *pulumi.OutputState }

func (DiskBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskBackup)(nil)).Elem()
}

func (o DiskBackupMapOutput) ToDiskBackupMapOutput() DiskBackupMapOutput {
	return o
}

func (o DiskBackupMapOutput) ToDiskBackupMapOutputWithContext(ctx context.Context) DiskBackupMapOutput {
	return o
}

func (o DiskBackupMapOutput) MapIndex(k pulumi.StringInput) DiskBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DiskBackup {
		return vs[0].(map[string]*DiskBackup)[vs[1].(string)]
	}).(DiskBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupInput)(nil)).Elem(), &DiskBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupArrayInput)(nil)).Elem(), DiskBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupMapInput)(nil)).Elem(), DiskBackupMap{})
	pulumi.RegisterOutputType(DiskBackupOutput{})
	pulumi.RegisterOutputType(DiskBackupArrayOutput{})
	pulumi.RegisterOutputType(DiskBackupMapOutput{})
}
