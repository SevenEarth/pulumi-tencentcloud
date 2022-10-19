// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create CBS set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cbs.NewStorageSet(ctx, "storage", &Cbs.StorageSetArgs{
//				AvailabilityZone: pulumi.String("ap-guangzhou-3"),
//				DiskCount:        pulumi.Int(10),
//				Encrypt:          pulumi.Bool(false),
//				ProjectId:        pulumi.Int(0),
//				StorageName:      pulumi.String("mystorage"),
//				StorageSize:      pulumi.Int(100),
//				StorageType:      pulumi.String("CLOUD_SSD"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StorageSet struct {
	pulumi.CustomResourceState

	// Indicates whether the CBS is mounted the CVM.
	Attached pulumi.BoolOutput `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// The number of disks to be purchased. Default 1.
	DiskCount pulumi.IntPtrOutput `pulumi:"diskCount"`
	// disk id list.
	DiskIds pulumi.StringArrayOutput `pulumi:"diskIds"`
	// Indicates whether CBS is encrypted.
	Encrypt pulumi.BoolPtrOutput `pulumi:"encrypt"`
	// ID of the project to which the instance belongs.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// ID of the snapshot. If specified, created the CBS by this snapshot.
	SnapshotId pulumi.StringOutput `pulumi:"snapshotId"`
	// Name of CBS. The maximum length can not exceed 60 bytes.
	StorageName pulumi.StringOutput `pulumi:"storageName"`
	// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
	StorageSize pulumi.IntOutput `pulumi:"storageSize"`
	// Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
	StorageStatus pulumi.StringOutput `pulumi:"storageStatus"`
	// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance pulumi.IntPtrOutput `pulumi:"throughputPerformance"`
}

// NewStorageSet registers a new resource with the given unique name, arguments, and options.
func NewStorageSet(ctx *pulumi.Context,
	name string, args *StorageSetArgs, opts ...pulumi.ResourceOption) (*StorageSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.StorageName == nil {
		return nil, errors.New("invalid value for required argument 'StorageName'")
	}
	if args.StorageSize == nil {
		return nil, errors.New("invalid value for required argument 'StorageSize'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	var resource StorageSet
	err := ctx.RegisterResource("tencentcloud:Cbs/storageSet:StorageSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageSet gets an existing StorageSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSetState, opts ...pulumi.ResourceOption) (*StorageSet, error) {
	var resource StorageSet
	err := ctx.ReadResource("tencentcloud:Cbs/storageSet:StorageSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageSet resources.
type storageSetState struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached *bool `pulumi:"attached"`
	// The available zone that the CBS instance locates at.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
	ChargeType *string `pulumi:"chargeType"`
	// The number of disks to be purchased. Default 1.
	DiskCount *int `pulumi:"diskCount"`
	// disk id list.
	DiskIds []string `pulumi:"diskIds"`
	// Indicates whether CBS is encrypted.
	Encrypt *bool `pulumi:"encrypt"`
	// ID of the project to which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the snapshot. If specified, created the CBS by this snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// Name of CBS. The maximum length can not exceed 60 bytes.
	StorageName *string `pulumi:"storageName"`
	// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
	StorageSize *int `pulumi:"storageSize"`
	// Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
	StorageStatus *string `pulumi:"storageStatus"`
	// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
	StorageType *string `pulumi:"storageType"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance *int `pulumi:"throughputPerformance"`
}

type StorageSetState struct {
	// Indicates whether the CBS is mounted the CVM.
	Attached pulumi.BoolPtrInput
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringPtrInput
	// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrInput
	// The number of disks to be purchased. Default 1.
	DiskCount pulumi.IntPtrInput
	// disk id list.
	DiskIds pulumi.StringArrayInput
	// Indicates whether CBS is encrypted.
	Encrypt pulumi.BoolPtrInput
	// ID of the project to which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the snapshot. If specified, created the CBS by this snapshot.
	SnapshotId pulumi.StringPtrInput
	// Name of CBS. The maximum length can not exceed 60 bytes.
	StorageName pulumi.StringPtrInput
	// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
	StorageSize pulumi.IntPtrInput
	// Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
	StorageStatus pulumi.StringPtrInput
	// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
	StorageType pulumi.StringPtrInput
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance pulumi.IntPtrInput
}

func (StorageSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSetState)(nil)).Elem()
}

type storageSetArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
	ChargeType *string `pulumi:"chargeType"`
	// The number of disks to be purchased. Default 1.
	DiskCount *int `pulumi:"diskCount"`
	// Indicates whether CBS is encrypted.
	Encrypt *bool `pulumi:"encrypt"`
	// ID of the project to which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the snapshot. If specified, created the CBS by this snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// Name of CBS. The maximum length can not exceed 60 bytes.
	StorageName string `pulumi:"storageName"`
	// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
	StorageSize int `pulumi:"storageSize"`
	// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
	StorageType string `pulumi:"storageType"`
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance *int `pulumi:"throughputPerformance"`
}

// The set of arguments for constructing a StorageSet resource.
type StorageSetArgs struct {
	// The available zone that the CBS instance locates at.
	AvailabilityZone pulumi.StringInput
	// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrInput
	// The number of disks to be purchased. Default 1.
	DiskCount pulumi.IntPtrInput
	// Indicates whether CBS is encrypted.
	Encrypt pulumi.BoolPtrInput
	// ID of the project to which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the snapshot. If specified, created the CBS by this snapshot.
	SnapshotId pulumi.StringPtrInput
	// Name of CBS. The maximum length can not exceed 60 bytes.
	StorageName pulumi.StringInput
	// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
	StorageSize pulumi.IntInput
	// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
	StorageType pulumi.StringInput
	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	ThroughputPerformance pulumi.IntPtrInput
}

func (StorageSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSetArgs)(nil)).Elem()
}

type StorageSetInput interface {
	pulumi.Input

	ToStorageSetOutput() StorageSetOutput
	ToStorageSetOutputWithContext(ctx context.Context) StorageSetOutput
}

func (*StorageSet) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSet)(nil)).Elem()
}

func (i *StorageSet) ToStorageSetOutput() StorageSetOutput {
	return i.ToStorageSetOutputWithContext(context.Background())
}

func (i *StorageSet) ToStorageSetOutputWithContext(ctx context.Context) StorageSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSetOutput)
}

// StorageSetArrayInput is an input type that accepts StorageSetArray and StorageSetArrayOutput values.
// You can construct a concrete instance of `StorageSetArrayInput` via:
//
//	StorageSetArray{ StorageSetArgs{...} }
type StorageSetArrayInput interface {
	pulumi.Input

	ToStorageSetArrayOutput() StorageSetArrayOutput
	ToStorageSetArrayOutputWithContext(context.Context) StorageSetArrayOutput
}

type StorageSetArray []StorageSetInput

func (StorageSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageSet)(nil)).Elem()
}

func (i StorageSetArray) ToStorageSetArrayOutput() StorageSetArrayOutput {
	return i.ToStorageSetArrayOutputWithContext(context.Background())
}

func (i StorageSetArray) ToStorageSetArrayOutputWithContext(ctx context.Context) StorageSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSetArrayOutput)
}

// StorageSetMapInput is an input type that accepts StorageSetMap and StorageSetMapOutput values.
// You can construct a concrete instance of `StorageSetMapInput` via:
//
//	StorageSetMap{ "key": StorageSetArgs{...} }
type StorageSetMapInput interface {
	pulumi.Input

	ToStorageSetMapOutput() StorageSetMapOutput
	ToStorageSetMapOutputWithContext(context.Context) StorageSetMapOutput
}

type StorageSetMap map[string]StorageSetInput

func (StorageSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageSet)(nil)).Elem()
}

func (i StorageSetMap) ToStorageSetMapOutput() StorageSetMapOutput {
	return i.ToStorageSetMapOutputWithContext(context.Background())
}

func (i StorageSetMap) ToStorageSetMapOutputWithContext(ctx context.Context) StorageSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSetMapOutput)
}

type StorageSetOutput struct{ *pulumi.OutputState }

func (StorageSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSet)(nil)).Elem()
}

func (o StorageSetOutput) ToStorageSetOutput() StorageSetOutput {
	return o
}

func (o StorageSetOutput) ToStorageSetOutputWithContext(ctx context.Context) StorageSetOutput {
	return o
}

// Indicates whether the CBS is mounted the CVM.
func (o StorageSetOutput) Attached() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.BoolOutput { return v.Attached }).(pulumi.BoolOutput)
}

// The available zone that the CBS instance locates at.
func (o StorageSetOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
func (o StorageSetOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// The number of disks to be purchased. Default 1.
func (o StorageSetOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.IntPtrOutput { return v.DiskCount }).(pulumi.IntPtrOutput)
}

// disk id list.
func (o StorageSetOutput) DiskIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringArrayOutput { return v.DiskIds }).(pulumi.StringArrayOutput)
}

// Indicates whether CBS is encrypted.
func (o StorageSetOutput) Encrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.BoolPtrOutput { return v.Encrypt }).(pulumi.BoolPtrOutput)
}

// ID of the project to which the instance belongs.
func (o StorageSetOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// ID of the snapshot. If specified, created the CBS by this snapshot.
func (o StorageSetOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringOutput { return v.SnapshotId }).(pulumi.StringOutput)
}

// Name of CBS. The maximum length can not exceed 60 bytes.
func (o StorageSetOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringOutput { return v.StorageName }).(pulumi.StringOutput)
}

// Volume of CBS, and unit is GB. If storage type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
func (o StorageSetOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.IntOutput { return v.StorageSize }).(pulumi.IntOutput)
}

// Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
func (o StorageSetOutput) StorageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringOutput { return v.StorageStatus }).(pulumi.StringOutput)
}

// Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.
func (o StorageSetOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
func (o StorageSetOutput) ThroughputPerformance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageSet) pulumi.IntPtrOutput { return v.ThroughputPerformance }).(pulumi.IntPtrOutput)
}

type StorageSetArrayOutput struct{ *pulumi.OutputState }

func (StorageSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageSet)(nil)).Elem()
}

func (o StorageSetArrayOutput) ToStorageSetArrayOutput() StorageSetArrayOutput {
	return o
}

func (o StorageSetArrayOutput) ToStorageSetArrayOutputWithContext(ctx context.Context) StorageSetArrayOutput {
	return o
}

func (o StorageSetArrayOutput) Index(i pulumi.IntInput) StorageSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageSet {
		return vs[0].([]*StorageSet)[vs[1].(int)]
	}).(StorageSetOutput)
}

type StorageSetMapOutput struct{ *pulumi.OutputState }

func (StorageSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageSet)(nil)).Elem()
}

func (o StorageSetMapOutput) ToStorageSetMapOutput() StorageSetMapOutput {
	return o
}

func (o StorageSetMapOutput) ToStorageSetMapOutputWithContext(ctx context.Context) StorageSetMapOutput {
	return o
}

func (o StorageSetMapOutput) MapIndex(k pulumi.StringInput) StorageSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageSet {
		return vs[0].(map[string]*StorageSet)[vs[1].(string)]
	}).(StorageSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageSetInput)(nil)).Elem(), &StorageSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageSetArrayInput)(nil)).Elem(), StorageSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageSetMapInput)(nil)).Elem(), StorageSetMap{})
	pulumi.RegisterOutputType(StorageSetOutput{})
	pulumi.RegisterOutputType(StorageSetArrayOutput{})
	pulumi.RegisterOutputType(StorageSetMapOutput{})
}
