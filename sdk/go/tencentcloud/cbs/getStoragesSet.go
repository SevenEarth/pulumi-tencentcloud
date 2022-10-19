// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CBS storages in parallel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cbs.GetStoragesSet(ctx, &cbs.GetStoragesSetArgs{
//				AvailabilityZone: pulumi.StringRef("ap-guangzhou-3"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetStoragesSet(ctx *pulumi.Context, args *GetStoragesSetArgs, opts ...pulumi.InvokeOption) (*GetStoragesSetResult, error) {
	var rv GetStoragesSetResult
	err := ctx.Invoke("tencentcloud:Cbs/getStoragesSet:getStoragesSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStoragesSet.
type GetStoragesSetArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeTypes []string `pulumi:"chargeTypes"`
	// List filter by attached instance public or private IPs.
	InstanceIps []string `pulumi:"instanceIps"`
	// List filter by attached instance name.
	InstanceNames []string `pulumi:"instanceNames"`
	// Filter by whether the disk is portable (Boolean `true` or `false`).
	Portable *bool `pulumi:"portable"`
	// ID of the project with which the CBS is associated.
	ProjectId *int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the CBS to be queried.
	StorageId *string `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName *string `pulumi:"storageName"`
	// List filter by disk state (`UNATTACHED` | `ATTACHING` | `ATTACHED` | `DETACHING` | `EXPANDING` | `ROLLBACKING` | `TORECYCLE`).
	StorageStates []string `pulumi:"storageStates"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType *string `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage *string `pulumi:"storageUsage"`
	// List filter by tag keys.
	TagKeys []string `pulumi:"tagKeys"`
	// List filter by tag values.
	TagValues []string `pulumi:"tagValues"`
}

// A collection of values returned by getStoragesSet.
type GetStoragesSetResult struct {
	// The zone of CBS.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Pay type of the CBS instance.
	ChargeTypes []string `pulumi:"chargeTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	InstanceIps   []string `pulumi:"instanceIps"`
	InstanceNames []string `pulumi:"instanceNames"`
	Portable      *bool    `pulumi:"portable"`
	// ID of the project.
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of CBS.
	StorageId *string `pulumi:"storageId"`
	// A list of storage. Each element contains the following attributes:
	StorageLists []GetStoragesSetStorageList `pulumi:"storageLists"`
	// Name of CBS.
	StorageName   *string  `pulumi:"storageName"`
	StorageStates []string `pulumi:"storageStates"`
	// Types of storage medium.
	StorageType *string `pulumi:"storageType"`
	// Types of CBS.
	StorageUsage *string  `pulumi:"storageUsage"`
	TagKeys      []string `pulumi:"tagKeys"`
	TagValues    []string `pulumi:"tagValues"`
}

func GetStoragesSetOutput(ctx *pulumi.Context, args GetStoragesSetOutputArgs, opts ...pulumi.InvokeOption) GetStoragesSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStoragesSetResult, error) {
			args := v.(GetStoragesSetArgs)
			r, err := GetStoragesSet(ctx, &args, opts...)
			var s GetStoragesSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStoragesSetResultOutput)
}

// A collection of arguments for invoking getStoragesSet.
type GetStoragesSetOutputArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
	ChargeTypes pulumi.StringArrayInput `pulumi:"chargeTypes"`
	// List filter by attached instance public or private IPs.
	InstanceIps pulumi.StringArrayInput `pulumi:"instanceIps"`
	// List filter by attached instance name.
	InstanceNames pulumi.StringArrayInput `pulumi:"instanceNames"`
	// Filter by whether the disk is portable (Boolean `true` or `false`).
	Portable pulumi.BoolPtrInput `pulumi:"portable"`
	// ID of the project with which the CBS is associated.
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the CBS to be queried.
	StorageId pulumi.StringPtrInput `pulumi:"storageId"`
	// Name of the CBS to be queried.
	StorageName pulumi.StringPtrInput `pulumi:"storageName"`
	// List filter by disk state (`UNATTACHED` | `ATTACHING` | `ATTACHED` | `DETACHING` | `EXPANDING` | `ROLLBACKING` | `TORECYCLE`).
	StorageStates pulumi.StringArrayInput `pulumi:"storageStates"`
	// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
	StorageType pulumi.StringPtrInput `pulumi:"storageType"`
	// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
	StorageUsage pulumi.StringPtrInput `pulumi:"storageUsage"`
	// List filter by tag keys.
	TagKeys pulumi.StringArrayInput `pulumi:"tagKeys"`
	// List filter by tag values.
	TagValues pulumi.StringArrayInput `pulumi:"tagValues"`
}

func (GetStoragesSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetArgs)(nil)).Elem()
}

// A collection of values returned by getStoragesSet.
type GetStoragesSetResultOutput struct{ *pulumi.OutputState }

func (GetStoragesSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetResult)(nil)).Elem()
}

func (o GetStoragesSetResultOutput) ToGetStoragesSetResultOutput() GetStoragesSetResultOutput {
	return o
}

func (o GetStoragesSetResultOutput) ToGetStoragesSetResultOutputWithContext(ctx context.Context) GetStoragesSetResultOutput {
	return o
}

// The zone of CBS.
func (o GetStoragesSetResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Pay type of the CBS instance.
func (o GetStoragesSetResultOutput) ChargeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.ChargeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStoragesSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStoragesSetResultOutput) InstanceIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.InstanceIps }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) Portable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *bool { return v.Portable }).(pulumi.BoolPtrOutput)
}

// ID of the project.
func (o GetStoragesSetResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetStoragesSetResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of CBS.
func (o GetStoragesSetResultOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

// A list of storage. Each element contains the following attributes:
func (o GetStoragesSetResultOutput) StorageLists() GetStoragesSetStorageListArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []GetStoragesSetStorageList { return v.StorageLists }).(GetStoragesSetStorageListArrayOutput)
}

// Name of CBS.
func (o GetStoragesSetResultOutput) StorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageName }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) StorageStates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.StorageStates }).(pulumi.StringArrayOutput)
}

// Types of storage medium.
func (o GetStoragesSetResultOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

// Types of CBS.
func (o GetStoragesSetResultOutput) StorageUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageUsage }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) TagKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.TagKeys }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) TagValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.TagValues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStoragesSetResultOutput{})
}
