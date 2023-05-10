// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SQL Server instance resource to create basic database instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Sqlserver.NewBasicInstance(ctx, "foo", &Sqlserver.BasicInstanceArgs{
// 			AvailabilityZone: pulumi.Any(_var.Availability_zone),
// 			ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
// 			VpcId:            pulumi.String("vpc-26w7r56z"),
// 			SubnetId:         pulumi.String("subnet-lvlr6eeu"),
// 			ProjectId:        pulumi.Int(0),
// 			Memory:           pulumi.Int(2),
// 			Storage:          pulumi.Int(20),
// 			Cpu:              pulumi.Int(1),
// 			MachineType:      pulumi.String("CLOUD_PREMIUM"),
// 			MaintenanceWeekSets: pulumi.IntArray{
// 				pulumi.Int(1),
// 				pulumi.Int(2),
// 				pulumi.Int(3),
// 			},
// 			MaintenanceStartTime: pulumi.String("09:00"),
// 			MaintenanceTimeSpan:  pulumi.Int(3),
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("sg-nltpbqg1"),
// 			},
// 			Tags: pulumi.AnyMap{
// 				"test": pulumi.Any("test"),
// 			},
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
// SQL Server basic instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/basicInstance:BasicInstance foo mssql-3cdq7kx5
// ```
type BasicInstance struct {
	pulumi.CustomResourceState

	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
	AutoRenew pulumi.IntPtrOutput `pulumi:"autoRenew"`
	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	AutoVoucher pulumi.IntPtrOutput `pulumi:"autoVoucher"`
	// Availability zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// The CPU number of the SQL Server basic instance.
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// Create time of the SQL Server basic instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Start time of the maintenance in one day, format like `HH:mm`.
	MaintenanceStartTime pulumi.StringOutput `pulumi:"maintenanceStartTime"`
	// The timespan of maintenance in one day, unit is hour.
	MaintenanceTimeSpan pulumi.IntOutput `pulumi:"maintenanceTimeSpan"`
	// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
	MaintenanceWeekSets pulumi.IntArrayOutput `pulumi:"maintenanceWeekSets"`
	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Name of the SQL Server basic instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Project ID, default value is 0.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Security group bound to the instance.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
	Status pulumi.IntOutput `pulumi:"status"`
	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
	Storage pulumi.IntOutput `pulumi:"storage"`
	// ID of subnet.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// The tags of the SQL Server basic instance.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// IP for private access.
	Vip pulumi.StringOutput `pulumi:"vip"`
	// An array of voucher IDs, currently only one can be used for a single order.
	VoucherIds pulumi.StringArrayOutput `pulumi:"voucherIds"`
	// ID of VPC.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
	// Port for private access.
	Vport pulumi.IntOutput `pulumi:"vport"`
}

// NewBasicInstance registers a new resource with the given unique name, arguments, and options.
func NewBasicInstance(ctx *pulumi.Context,
	name string, args *BasicInstanceArgs, opts ...pulumi.ResourceOption) (*BasicInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cpu == nil {
		return nil, errors.New("invalid value for required argument 'Cpu'")
	}
	if args.MachineType == nil {
		return nil, errors.New("invalid value for required argument 'MachineType'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BasicInstance
	err := ctx.RegisterResource("tencentcloud:Sqlserver/basicInstance:BasicInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicInstance gets an existing BasicInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicInstanceState, opts ...pulumi.ResourceOption) (*BasicInstance, error) {
	var resource BasicInstance
	err := ctx.ReadResource("tencentcloud:Sqlserver/basicInstance:BasicInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicInstance resources.
type basicInstanceState struct {
	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
	AutoRenew *int `pulumi:"autoRenew"`
	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	AutoVoucher *int `pulumi:"autoVoucher"`
	// Availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
	ChargeType *string `pulumi:"chargeType"`
	// The CPU number of the SQL Server basic instance.
	Cpu *int `pulumi:"cpu"`
	// Create time of the SQL Server basic instance.
	CreateTime *string `pulumi:"createTime"`
	// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	EngineVersion *string `pulumi:"engineVersion"`
	// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
	MachineType *string `pulumi:"machineType"`
	// Start time of the maintenance in one day, format like `HH:mm`.
	MaintenanceStartTime *string `pulumi:"maintenanceStartTime"`
	// The timespan of maintenance in one day, unit is hour.
	MaintenanceTimeSpan *int `pulumi:"maintenanceTimeSpan"`
	// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
	MaintenanceWeekSets []int `pulumi:"maintenanceWeekSets"`
	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
	Memory *int `pulumi:"memory"`
	// Name of the SQL Server basic instance.
	Name *string `pulumi:"name"`
	// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
	Period *int `pulumi:"period"`
	// Project ID, default value is 0.
	ProjectId *int `pulumi:"projectId"`
	// Security group bound to the instance.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
	Status *int `pulumi:"status"`
	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
	Storage *int `pulumi:"storage"`
	// ID of subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The tags of the SQL Server basic instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// IP for private access.
	Vip *string `pulumi:"vip"`
	// An array of voucher IDs, currently only one can be used for a single order.
	VoucherIds []string `pulumi:"voucherIds"`
	// ID of VPC.
	VpcId *string `pulumi:"vpcId"`
	// Port for private access.
	Vport *int `pulumi:"vport"`
}

type BasicInstanceState struct {
	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
	AutoRenew pulumi.IntPtrInput
	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	AutoVoucher pulumi.IntPtrInput
	// Availability zone.
	AvailabilityZone pulumi.StringPtrInput
	// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
	ChargeType pulumi.StringPtrInput
	// The CPU number of the SQL Server basic instance.
	Cpu pulumi.IntPtrInput
	// Create time of the SQL Server basic instance.
	CreateTime pulumi.StringPtrInput
	// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	EngineVersion pulumi.StringPtrInput
	// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
	MachineType pulumi.StringPtrInput
	// Start time of the maintenance in one day, format like `HH:mm`.
	MaintenanceStartTime pulumi.StringPtrInput
	// The timespan of maintenance in one day, unit is hour.
	MaintenanceTimeSpan pulumi.IntPtrInput
	// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
	MaintenanceWeekSets pulumi.IntArrayInput
	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
	Memory pulumi.IntPtrInput
	// Name of the SQL Server basic instance.
	Name pulumi.StringPtrInput
	// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
	Period pulumi.IntPtrInput
	// Project ID, default value is 0.
	ProjectId pulumi.IntPtrInput
	// Security group bound to the instance.
	SecurityGroups pulumi.StringArrayInput
	// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
	Status pulumi.IntPtrInput
	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
	Storage pulumi.IntPtrInput
	// ID of subnet.
	SubnetId pulumi.StringPtrInput
	// The tags of the SQL Server basic instance.
	Tags pulumi.MapInput
	// IP for private access.
	Vip pulumi.StringPtrInput
	// An array of voucher IDs, currently only one can be used for a single order.
	VoucherIds pulumi.StringArrayInput
	// ID of VPC.
	VpcId pulumi.StringPtrInput
	// Port for private access.
	Vport pulumi.IntPtrInput
}

func (BasicInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicInstanceState)(nil)).Elem()
}

type basicInstanceArgs struct {
	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
	AutoRenew *int `pulumi:"autoRenew"`
	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	AutoVoucher *int `pulumi:"autoVoucher"`
	// Availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
	ChargeType *string `pulumi:"chargeType"`
	// The CPU number of the SQL Server basic instance.
	Cpu int `pulumi:"cpu"`
	// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	EngineVersion *string `pulumi:"engineVersion"`
	// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
	MachineType string `pulumi:"machineType"`
	// Start time of the maintenance in one day, format like `HH:mm`.
	MaintenanceStartTime *string `pulumi:"maintenanceStartTime"`
	// The timespan of maintenance in one day, unit is hour.
	MaintenanceTimeSpan *int `pulumi:"maintenanceTimeSpan"`
	// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
	MaintenanceWeekSets []int `pulumi:"maintenanceWeekSets"`
	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
	Memory int `pulumi:"memory"`
	// Name of the SQL Server basic instance.
	Name *string `pulumi:"name"`
	// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
	Period *int `pulumi:"period"`
	// Project ID, default value is 0.
	ProjectId *int `pulumi:"projectId"`
	// Security group bound to the instance.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
	Storage int `pulumi:"storage"`
	// ID of subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The tags of the SQL Server basic instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// An array of voucher IDs, currently only one can be used for a single order.
	VoucherIds []string `pulumi:"voucherIds"`
	// ID of VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a BasicInstance resource.
type BasicInstanceArgs struct {
	// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
	AutoRenew pulumi.IntPtrInput
	// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
	AutoVoucher pulumi.IntPtrInput
	// Availability zone.
	AvailabilityZone pulumi.StringPtrInput
	// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
	ChargeType pulumi.StringPtrInput
	// The CPU number of the SQL Server basic instance.
	Cpu pulumi.IntInput
	// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
	EngineVersion pulumi.StringPtrInput
	// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
	MachineType pulumi.StringInput
	// Start time of the maintenance in one day, format like `HH:mm`.
	MaintenanceStartTime pulumi.StringPtrInput
	// The timespan of maintenance in one day, unit is hour.
	MaintenanceTimeSpan pulumi.IntPtrInput
	// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
	MaintenanceWeekSets pulumi.IntArrayInput
	// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
	Memory pulumi.IntInput
	// Name of the SQL Server basic instance.
	Name pulumi.StringPtrInput
	// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
	Period pulumi.IntPtrInput
	// Project ID, default value is 0.
	ProjectId pulumi.IntPtrInput
	// Security group bound to the instance.
	SecurityGroups pulumi.StringArrayInput
	// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
	Storage pulumi.IntInput
	// ID of subnet.
	SubnetId pulumi.StringPtrInput
	// The tags of the SQL Server basic instance.
	Tags pulumi.MapInput
	// An array of voucher IDs, currently only one can be used for a single order.
	VoucherIds pulumi.StringArrayInput
	// ID of VPC.
	VpcId pulumi.StringPtrInput
}

func (BasicInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicInstanceArgs)(nil)).Elem()
}

type BasicInstanceInput interface {
	pulumi.Input

	ToBasicInstanceOutput() BasicInstanceOutput
	ToBasicInstanceOutputWithContext(ctx context.Context) BasicInstanceOutput
}

func (*BasicInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicInstance)(nil)).Elem()
}

func (i *BasicInstance) ToBasicInstanceOutput() BasicInstanceOutput {
	return i.ToBasicInstanceOutputWithContext(context.Background())
}

func (i *BasicInstance) ToBasicInstanceOutputWithContext(ctx context.Context) BasicInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicInstanceOutput)
}

// BasicInstanceArrayInput is an input type that accepts BasicInstanceArray and BasicInstanceArrayOutput values.
// You can construct a concrete instance of `BasicInstanceArrayInput` via:
//
//          BasicInstanceArray{ BasicInstanceArgs{...} }
type BasicInstanceArrayInput interface {
	pulumi.Input

	ToBasicInstanceArrayOutput() BasicInstanceArrayOutput
	ToBasicInstanceArrayOutputWithContext(context.Context) BasicInstanceArrayOutput
}

type BasicInstanceArray []BasicInstanceInput

func (BasicInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicInstance)(nil)).Elem()
}

func (i BasicInstanceArray) ToBasicInstanceArrayOutput() BasicInstanceArrayOutput {
	return i.ToBasicInstanceArrayOutputWithContext(context.Background())
}

func (i BasicInstanceArray) ToBasicInstanceArrayOutputWithContext(ctx context.Context) BasicInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicInstanceArrayOutput)
}

// BasicInstanceMapInput is an input type that accepts BasicInstanceMap and BasicInstanceMapOutput values.
// You can construct a concrete instance of `BasicInstanceMapInput` via:
//
//          BasicInstanceMap{ "key": BasicInstanceArgs{...} }
type BasicInstanceMapInput interface {
	pulumi.Input

	ToBasicInstanceMapOutput() BasicInstanceMapOutput
	ToBasicInstanceMapOutputWithContext(context.Context) BasicInstanceMapOutput
}

type BasicInstanceMap map[string]BasicInstanceInput

func (BasicInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicInstance)(nil)).Elem()
}

func (i BasicInstanceMap) ToBasicInstanceMapOutput() BasicInstanceMapOutput {
	return i.ToBasicInstanceMapOutputWithContext(context.Background())
}

func (i BasicInstanceMap) ToBasicInstanceMapOutputWithContext(ctx context.Context) BasicInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicInstanceMapOutput)
}

type BasicInstanceOutput struct{ *pulumi.OutputState }

func (BasicInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicInstance)(nil)).Elem()
}

func (o BasicInstanceOutput) ToBasicInstanceOutput() BasicInstanceOutput {
	return o
}

func (o BasicInstanceOutput) ToBasicInstanceOutputWithContext(ctx context.Context) BasicInstanceOutput {
	return o
}

// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
func (o BasicInstanceOutput) AutoRenew() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntPtrOutput { return v.AutoRenew }).(pulumi.IntPtrOutput)
}

// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
func (o BasicInstanceOutput) AutoVoucher() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntPtrOutput { return v.AutoVoucher }).(pulumi.IntPtrOutput)
}

// Availability zone.
func (o BasicInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
func (o BasicInstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// The CPU number of the SQL Server basic instance.
func (o BasicInstanceOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.Cpu }).(pulumi.IntOutput)
}

// Create time of the SQL Server basic instance.
func (o BasicInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
func (o BasicInstanceOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk, `CLOUD_HSSD` for virtual machine enhanced cloud disk, `CLOUD_BSSD` for virtual machine general purpose SSD cloud disk.
func (o BasicInstanceOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.MachineType }).(pulumi.StringOutput)
}

// Start time of the maintenance in one day, format like `HH:mm`.
func (o BasicInstanceOutput) MaintenanceStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.MaintenanceStartTime }).(pulumi.StringOutput)
}

// The timespan of maintenance in one day, unit is hour.
func (o BasicInstanceOutput) MaintenanceTimeSpan() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.MaintenanceTimeSpan }).(pulumi.IntOutput)
}

// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
func (o BasicInstanceOutput) MaintenanceWeekSets() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntArrayOutput { return v.MaintenanceWeekSets }).(pulumi.IntArrayOutput)
}

// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloudSqlserverSpecinfos` provides.
func (o BasicInstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// Name of the SQL Server basic instance.
func (o BasicInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
func (o BasicInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Project ID, default value is 0.
func (o BasicInstanceOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Security group bound to the instance.
func (o BasicInstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
func (o BasicInstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storageMin` and `storageMax` which data source `tencentcloudSqlserverSpecinfos` provides.
func (o BasicInstanceOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.Storage }).(pulumi.IntOutput)
}

// ID of subnet.
func (o BasicInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The tags of the SQL Server basic instance.
func (o BasicInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// IP for private access.
func (o BasicInstanceOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

// An array of voucher IDs, currently only one can be used for a single order.
func (o BasicInstanceOutput) VoucherIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringArrayOutput { return v.VoucherIds }).(pulumi.StringArrayOutput)
}

// ID of VPC.
func (o BasicInstanceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

// Port for private access.
func (o BasicInstanceOutput) Vport() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicInstance) pulumi.IntOutput { return v.Vport }).(pulumi.IntOutput)
}

type BasicInstanceArrayOutput struct{ *pulumi.OutputState }

func (BasicInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicInstance)(nil)).Elem()
}

func (o BasicInstanceArrayOutput) ToBasicInstanceArrayOutput() BasicInstanceArrayOutput {
	return o
}

func (o BasicInstanceArrayOutput) ToBasicInstanceArrayOutputWithContext(ctx context.Context) BasicInstanceArrayOutput {
	return o
}

func (o BasicInstanceArrayOutput) Index(i pulumi.IntInput) BasicInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicInstance {
		return vs[0].([]*BasicInstance)[vs[1].(int)]
	}).(BasicInstanceOutput)
}

type BasicInstanceMapOutput struct{ *pulumi.OutputState }

func (BasicInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicInstance)(nil)).Elem()
}

func (o BasicInstanceMapOutput) ToBasicInstanceMapOutput() BasicInstanceMapOutput {
	return o
}

func (o BasicInstanceMapOutput) ToBasicInstanceMapOutputWithContext(ctx context.Context) BasicInstanceMapOutput {
	return o
}

func (o BasicInstanceMapOutput) MapIndex(k pulumi.StringInput) BasicInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicInstance {
		return vs[0].(map[string]*BasicInstance)[vs[1].(string)]
	}).(BasicInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicInstanceInput)(nil)).Elem(), &BasicInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicInstanceArrayInput)(nil)).Elem(), BasicInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicInstanceMapInput)(nil)).Elem(), BasicInstanceMap{})
	pulumi.RegisterOutputType(BasicInstanceOutput{})
	pulumi.RegisterOutputType(BasicInstanceArrayOutput{})
	pulumi.RegisterOutputType(BasicInstanceMapOutput{})
}
