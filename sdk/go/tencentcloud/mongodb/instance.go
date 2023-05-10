// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a Mongodb instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mongodb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mongodb.NewInstance(ctx, "mongodb", &Mongodb.InstanceArgs{
// 			AvailableZone: pulumi.String("ap-guangzhou-2"),
// 			EngineVersion: pulumi.String("MONGO_3_WT"),
// 			InstanceName:  pulumi.String("mongodb"),
// 			MachineType:   pulumi.String("GIO"),
// 			Memory:        pulumi.Int(4),
// 			Password:      pulumi.String("password1234"),
// 			ProjectId:     pulumi.Int(0),
// 			SubnetId:      pulumi.String("subnet-lk0svi3p"),
// 			Volume:        pulumi.Int(100),
// 			VpcId:         pulumi.String("vpc-mz3efvbw"),
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
// Mongodb instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mongodb/instance:Instance mongodb cmgo-41s6jwy4
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringOutput `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// Creation time of the Mongodb instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Name of the Mongodb instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Password of this Mongodb account.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrOutput `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// List of standby instances' info.
	StandbyInstanceLists InstanceStandbyInstanceListArrayOutput `pulumi:"standbyInstanceLists"`
	// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
	Status pulumi.IntOutput `pulumi:"status"`
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// IP of the Mongodb instance.
	Vip pulumi.StringOutput `pulumi:"vip"`
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume pulumi.IntOutput `pulumi:"volume"`
	// ID of the VPC.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
	// IP port of the Mongodb instance.
	Vport pulumi.IntOutput `pulumi:"vport"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailableZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailableZone'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.MachineType == nil {
		return nil, errors.New("invalid value for required argument 'MachineType'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.Volume == nil {
		return nil, errors.New("invalid value for required argument 'Volume'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Mongodb/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Mongodb/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// The available zone of the Mongodb.
	AvailableZone *string `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// Creation time of the Mongodb instance.
	CreateTime *string `pulumi:"createTime"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion *string `pulumi:"engineVersion"`
	// Name of the Mongodb instance.
	InstanceName *string `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType *string `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory *int `pulumi:"memory"`
	// Password of this Mongodb account.
	Password *string `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups []string `pulumi:"securityGroups"`
	// List of standby instances' info.
	StandbyInstanceLists []InstanceStandbyInstanceList `pulumi:"standbyInstanceLists"`
	// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
	Status *int `pulumi:"status"`
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId *string `pulumi:"subnetId"`
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags map[string]interface{} `pulumi:"tags"`
	// IP of the Mongodb instance.
	Vip *string `pulumi:"vip"`
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume *int `pulumi:"volume"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// IP port of the Mongodb instance.
	Vport *int `pulumi:"vport"`
}

type InstanceState struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrInput
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringPtrInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrInput
	// Creation time of the Mongodb instance.
	CreateTime pulumi.StringPtrInput
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringPtrInput
	// Name of the Mongodb instance.
	InstanceName pulumi.StringPtrInput
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringPtrInput
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntPtrInput
	// Password of this Mongodb account.
	Password pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrInput
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayInput
	// List of standby instances' info.
	StandbyInstanceLists InstanceStandbyInstanceListArrayInput
	// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
	Status pulumi.IntPtrInput
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId pulumi.StringPtrInput
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags pulumi.MapInput
	// IP of the Mongodb instance.
	Vip pulumi.StringPtrInput
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume pulumi.IntPtrInput
	// ID of the VPC.
	VpcId pulumi.StringPtrInput
	// IP port of the Mongodb instance.
	Vport pulumi.IntPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// The available zone of the Mongodb.
	AvailableZone string `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion string `pulumi:"engineVersion"`
	// Name of the Mongodb instance.
	InstanceName string `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType string `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory int `pulumi:"memory"`
	// Password of this Mongodb account.
	Password *string `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups []string `pulumi:"securityGroups"`
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId *string `pulumi:"subnetId"`
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags map[string]interface{} `pulumi:"tags"`
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume int `pulumi:"volume"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrInput
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrInput
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringInput
	// Name of the Mongodb instance.
	InstanceName pulumi.StringInput
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringInput
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntInput
	// Password of this Mongodb account.
	Password pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrInput
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayInput
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId pulumi.StringPtrInput
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags pulumi.MapInput
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume pulumi.IntInput
	// ID of the VPC.
	VpcId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
func (o InstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// The available zone of the Mongodb.
func (o InstanceOutput) AvailableZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailableZone }).(pulumi.StringOutput)
}

// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
func (o InstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// Creation time of the Mongodb instance.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
func (o InstanceOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// Name of the Mongodb instance.
func (o InstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
func (o InstanceOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.MachineType }).(pulumi.StringOutput)
}

// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
func (o InstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// Password of this Mongodb account.
func (o InstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
func (o InstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// ID of the project which the instance belongs.
func (o InstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// List of standby instances' info.
func (o InstanceOutput) StandbyInstanceLists() InstanceStandbyInstanceListArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceStandbyInstanceListArrayOutput { return v.StandbyInstanceLists }).(InstanceStandbyInstanceListArrayOutput)
}

// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
func (o InstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// ID of the subnet within this VPC. The value is required if `vpcId` is set.
func (o InstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// IP of the Mongodb instance.
func (o InstanceOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
func (o InstanceOutput) Volume() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Volume }).(pulumi.IntOutput)
}

// ID of the VPC.
func (o InstanceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

// IP port of the Mongodb instance.
func (o InstanceOutput) Vport() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Vport }).(pulumi.IntOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
