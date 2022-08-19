// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CBS snapshots.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cbs.GetSnapshots(ctx, &cbs.GetSnapshotsArgs{
// 			ResultOutputFile: pulumi.StringRef("mytestpath"),
// 			SnapshotId:       pulumi.StringRef("snap-f3io7adt"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSnapshots(ctx *pulumi.Context, args *GetSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetSnapshotsResult, error) {
	var rv GetSnapshotsResult
	err := ctx.Invoke("tencentcloud:Cbs/getSnapshots:getSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// ID of the project within the snapshot.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the snapshot to be queried.
	SnapshotId *string `pulumi:"snapshotId"`
	// Name of the snapshot to be queried.
	SnapshotName *string `pulumi:"snapshotName"`
	// ID of the the CBS which this snapshot created from.
	StorageId *string `pulumi:"storageId"`
	// Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
	StorageUsage *string `pulumi:"storageUsage"`
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResult struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the project within the snapshot.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of snapshot. Each element contains the following attributes:
	SnapshotLists []GetSnapshotsSnapshotList `pulumi:"snapshotLists"`
	// Name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// ID of the the CBS which this snapshot created from.
	StorageId *string `pulumi:"storageId"`
	// Types of CBS which this snapshot created from.
	StorageUsage *string `pulumi:"storageUsage"`
}

func GetSnapshotsOutput(ctx *pulumi.Context, args GetSnapshotsOutputArgs, opts ...pulumi.InvokeOption) GetSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSnapshotsResult, error) {
			args := v.(GetSnapshotsArgs)
			r, err := GetSnapshots(ctx, &args, opts...)
			var s GetSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSnapshotsResultOutput)
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsOutputArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// ID of the project within the snapshot.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the snapshot to be queried.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	// Name of the snapshot to be queried.
	SnapshotName pulumi.StringPtrInput `pulumi:"snapshotName"`
	// ID of the the CBS which this snapshot created from.
	StorageId pulumi.StringPtrInput `pulumi:"storageId"`
	// Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
	StorageUsage pulumi.StringPtrInput `pulumi:"storageUsage"`
}

func (GetSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsResult)(nil)).Elem()
}

func (o GetSnapshotsResultOutput) ToGetSnapshotsResultOutput() GetSnapshotsResultOutput {
	return o
}

func (o GetSnapshotsResultOutput) ToGetSnapshotsResultOutputWithContext(ctx context.Context) GetSnapshotsResultOutput {
	return o
}

// The available zone that the CBS instance locates at.
func (o GetSnapshotsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the project within the snapshot.
func (o GetSnapshotsResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetSnapshotsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of the snapshot.
func (o GetSnapshotsResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// A list of snapshot. Each element contains the following attributes:
func (o GetSnapshotsResultOutput) SnapshotLists() GetSnapshotsSnapshotListArrayOutput {
	return o.ApplyT(func(v GetSnapshotsResult) []GetSnapshotsSnapshotList { return v.SnapshotLists }).(GetSnapshotsSnapshotListArrayOutput)
}

// Name of the snapshot.
func (o GetSnapshotsResultOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

// ID of the the CBS which this snapshot created from.
func (o GetSnapshotsResultOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

// Types of CBS which this snapshot created from.
func (o GetSnapshotsResultOutput) StorageUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.StorageUsage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotsResultOutput{})
}
