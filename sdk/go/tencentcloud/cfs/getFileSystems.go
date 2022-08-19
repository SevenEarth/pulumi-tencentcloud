// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the detail information of cloud file systems(CFS).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cfs.GetFileSystems(ctx, &cfs.GetFileSystemsArgs{
// 			AvailabilityZone: pulumi.StringRef("ap-guangzhou-3"),
// 			FileSystemId:     pulumi.StringRef("cfs-6hgquxmj"),
// 			Name:             pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetFileSystems(ctx *pulumi.Context, args *GetFileSystemsArgs, opts ...pulumi.InvokeOption) (*GetFileSystemsResult, error) {
	var rv GetFileSystemsResult
	err := ctx.Invoke("tencentcloud:Cfs/getFileSystems:getFileSystems", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFileSystems.
type GetFileSystemsArgs struct {
	// The available zone that the file system locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// A specified file system ID used to query.
	FileSystemId *string `pulumi:"fileSystemId"`
	// A file system name used to query.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of a vpc subnet.
	SubnetId *string `pulumi:"subnetId"`
	// ID of the vpc to be queried.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getFileSystems.
type GetFileSystemsResult struct {
	// The available zone that the file system locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// ID of the file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// An information list of cloud file system. Each element contains the following attributes:
	FileSystemLists []GetFileSystemsFileSystemList `pulumi:"fileSystemLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the file system.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SubnetId         *string `pulumi:"subnetId"`
	VpcId            *string `pulumi:"vpcId"`
}

func GetFileSystemsOutput(ctx *pulumi.Context, args GetFileSystemsOutputArgs, opts ...pulumi.InvokeOption) GetFileSystemsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFileSystemsResult, error) {
			args := v.(GetFileSystemsArgs)
			r, err := GetFileSystems(ctx, &args, opts...)
			var s GetFileSystemsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFileSystemsResultOutput)
}

// A collection of arguments for invoking getFileSystems.
type GetFileSystemsOutputArgs struct {
	// The available zone that the file system locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// A specified file system ID used to query.
	FileSystemId pulumi.StringPtrInput `pulumi:"fileSystemId"`
	// A file system name used to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of a vpc subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// ID of the vpc to be queried.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetFileSystemsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileSystemsArgs)(nil)).Elem()
}

// A collection of values returned by getFileSystems.
type GetFileSystemsResultOutput struct{ *pulumi.OutputState }

func (GetFileSystemsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileSystemsResult)(nil)).Elem()
}

func (o GetFileSystemsResultOutput) ToGetFileSystemsResultOutput() GetFileSystemsResultOutput {
	return o
}

func (o GetFileSystemsResultOutput) ToGetFileSystemsResultOutputWithContext(ctx context.Context) GetFileSystemsResultOutput {
	return o
}

// The available zone that the file system locates at.
func (o GetFileSystemsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// ID of the file system.
func (o GetFileSystemsResultOutput) FileSystemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.FileSystemId }).(pulumi.StringPtrOutput)
}

// An information list of cloud file system. Each element contains the following attributes:
func (o GetFileSystemsResultOutput) FileSystemLists() GetFileSystemsFileSystemListArrayOutput {
	return o.ApplyT(func(v GetFileSystemsResult) []GetFileSystemsFileSystemList { return v.FileSystemLists }).(GetFileSystemsFileSystemListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFileSystemsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileSystemsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the file system.
func (o GetFileSystemsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetFileSystemsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetFileSystemsResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o GetFileSystemsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFileSystemsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFileSystemsResultOutput{})
}
