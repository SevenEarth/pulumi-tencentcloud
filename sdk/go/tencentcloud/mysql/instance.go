// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a mysql instance resource to create master database instances.
//
// > **NOTE:** If this mysql has readonly instance, the terminate operation of the mysql does NOT take effect immediately, maybe takes for several hours. so during that time, VPCs associated with that mysql instance can't be terminated also.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.NewInstance(ctx, "default", &Mysql.InstanceArgs{
//				AvailabilityZone: pulumi.String("ap-guangzhou-4"),
//				ChargeType:       pulumi.String("POSTPAID"),
//				EngineVersion:    pulumi.String("5.7"),
//				FirstSlaveZone:   pulumi.String("ap-guangzhou-4"),
//				InstanceName:     pulumi.String("myTestMysql"),
//				InternetService:  pulumi.Int(1),
//				IntranetPort:     pulumi.Int(3306),
//				MemSize:          pulumi.Int(128000),
//				Parameters: pulumi.AnyMap{
//					"max_connections": pulumi.Any("1000"),
//				},
//				ProjectId:       pulumi.Int(201901010001),
//				RootPassword:    pulumi.String("********"),
//				SecondSlaveZone: pulumi.String("ap-guangzhou-4"),
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("sg-ot8eclwz"),
//				},
//				SlaveDeployMode: pulumi.Int(0),
//				SlaveSyncMode:   pulumi.Int(1),
//				SubnetId:        pulumi.String("subnet-9uivyb1g"),
//				Tags: pulumi.AnyMap{
//					"name": pulumi.Any("test"),
//				},
//				VolumeSize: pulumi.Int(250),
//				VpcId:      pulumi.String("vpc-12mt3l31"),
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
// MySQL instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Mysql/instance:Instance foo cdb-12345678"
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// Indicates which availability zone will be used.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// CPU cores.
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringPtrOutput `pulumi:"deviceType"`
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrOutput `pulumi:"fastUpgrade"`
	// Zone information about first slave instance.
	FirstSlaveZone pulumi.StringPtrOutput `pulumi:"firstSlaveZone"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// Indicates whether GTID is enable. `0` - Not enabled; `1` - Enabled.
	Gtid pulumi.IntOutput `pulumi:"gtid"`
	// The name of a mysql instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// host for public access.
	InternetHost pulumi.StringOutput `pulumi:"internetHost"`
	// Access port for public access.
	InternetPort pulumi.IntOutput `pulumi:"internetPort"`
	// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
	InternetService pulumi.IntPtrOutput `pulumi:"internetService"`
	// instance intranet IP.
	IntranetIp pulumi.StringOutput `pulumi:"intranetIp"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrOutput `pulumi:"intranetPort"`
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked pulumi.IntOutput `pulumi:"locked"`
	// Memory size (in MB).
	MemSize pulumi.IntOutput `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrOutput `pulumi:"paramTemplateId"`
	// List of parameters to use.
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
	PayType pulumi.IntPtrOutput `pulumi:"payType"`
	// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Period of instance. NOTES: Only supported prepaid instance.
	PrepaidPeriod pulumi.IntPtrOutput `pulumi:"prepaidPeriod"`
	// Project ID, default value is 0.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	RootPassword pulumi.StringPtrOutput `pulumi:"rootPassword"`
	// Zone information about second slave instance.
	SecondSlaveZone pulumi.StringPtrOutput `pulumi:"secondSlaveZone"`
	// Security groups to use.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrOutput `pulumi:"slaveDeployMode"`
	// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
	SlaveSyncMode pulumi.IntPtrOutput `pulumi:"slaveSyncMode"`
	// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
	Status pulumi.IntOutput `pulumi:"status"`
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Instance tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Indicates which kind of operations is being executed.
	TaskStatus pulumi.IntOutput `pulumi:"taskStatus"`
	// Disk size (in GB).
	VolumeSize pulumi.IntOutput `pulumi:"volumeSize"`
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.MemSize == nil {
		return nil, errors.New("invalid value for required argument 'MemSize'")
	}
	if args.VolumeSize == nil {
		return nil, errors.New("invalid value for required argument 'VolumeSize'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Mysql/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Mysql/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// Indicates which availability zone will be used.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// CPU cores.
	Cpu *int `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType *string `pulumi:"deviceType"`
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
	EngineVersion *string `pulumi:"engineVersion"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade *int `pulumi:"fastUpgrade"`
	// Zone information about first slave instance.
	FirstSlaveZone *string `pulumi:"firstSlaveZone"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete *bool `pulumi:"forceDelete"`
	// Indicates whether GTID is enable. `0` - Not enabled; `1` - Enabled.
	Gtid *int `pulumi:"gtid"`
	// The name of a mysql instance.
	InstanceName *string `pulumi:"instanceName"`
	// host for public access.
	InternetHost *string `pulumi:"internetHost"`
	// Access port for public access.
	InternetPort *int `pulumi:"internetPort"`
	// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
	InternetService *int `pulumi:"internetService"`
	// instance intranet IP.
	IntranetIp *string `pulumi:"intranetIp"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort *int `pulumi:"intranetPort"`
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked *int `pulumi:"locked"`
	// Memory size (in MB).
	MemSize *int `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId *int `pulumi:"paramTemplateId"`
	// List of parameters to use.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
	PayType *int `pulumi:"payType"`
	// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
	Period *int `pulumi:"period"`
	// Period of instance. NOTES: Only supported prepaid instance.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Project ID, default value is 0.
	ProjectId *int `pulumi:"projectId"`
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	RootPassword *string `pulumi:"rootPassword"`
	// Zone information about second slave instance.
	SecondSlaveZone *string `pulumi:"secondSlaveZone"`
	// Security groups to use.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode *int `pulumi:"slaveDeployMode"`
	// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
	SlaveSyncMode *int `pulumi:"slaveSyncMode"`
	// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
	Status *int `pulumi:"status"`
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Indicates which kind of operations is being executed.
	TaskStatus *int `pulumi:"taskStatus"`
	// Disk size (in GB).
	VolumeSize *int `pulumi:"volumeSize"`
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrInput
	// Indicates which availability zone will be used.
	AvailabilityZone pulumi.StringPtrInput
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrInput
	// CPU cores.
	Cpu pulumi.IntPtrInput
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringPtrInput
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
	EngineVersion pulumi.StringPtrInput
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrInput
	// Zone information about first slave instance.
	FirstSlaveZone pulumi.StringPtrInput
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrInput
	// Indicates whether GTID is enable. `0` - Not enabled; `1` - Enabled.
	Gtid pulumi.IntPtrInput
	// The name of a mysql instance.
	InstanceName pulumi.StringPtrInput
	// host for public access.
	InternetHost pulumi.StringPtrInput
	// Access port for public access.
	InternetPort pulumi.IntPtrInput
	// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
	InternetService pulumi.IntPtrInput
	// instance intranet IP.
	IntranetIp pulumi.StringPtrInput
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrInput
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked pulumi.IntPtrInput
	// Memory size (in MB).
	MemSize pulumi.IntPtrInput
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrInput
	// List of parameters to use.
	Parameters pulumi.MapInput
	// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
	PayType pulumi.IntPtrInput
	// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
	Period pulumi.IntPtrInput
	// Period of instance. NOTES: Only supported prepaid instance.
	PrepaidPeriod pulumi.IntPtrInput
	// Project ID, default value is 0.
	ProjectId pulumi.IntPtrInput
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	RootPassword pulumi.StringPtrInput
	// Zone information about second slave instance.
	SecondSlaveZone pulumi.StringPtrInput
	// Security groups to use.
	SecurityGroups pulumi.StringArrayInput
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrInput
	// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
	SlaveSyncMode pulumi.IntPtrInput
	// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
	Status pulumi.IntPtrInput
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// Indicates which kind of operations is being executed.
	TaskStatus pulumi.IntPtrInput
	// Disk size (in GB).
	VolumeSize pulumi.IntPtrInput
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// Indicates which availability zone will be used.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// CPU cores.
	Cpu *int `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType *string `pulumi:"deviceType"`
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
	EngineVersion *string `pulumi:"engineVersion"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade *int `pulumi:"fastUpgrade"`
	// Zone information about first slave instance.
	FirstSlaveZone *string `pulumi:"firstSlaveZone"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The name of a mysql instance.
	InstanceName string `pulumi:"instanceName"`
	// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
	InternetService *int `pulumi:"internetService"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort *int `pulumi:"intranetPort"`
	// Memory size (in MB).
	MemSize int `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId *int `pulumi:"paramTemplateId"`
	// List of parameters to use.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
	PayType *int `pulumi:"payType"`
	// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
	Period *int `pulumi:"period"`
	// Period of instance. NOTES: Only supported prepaid instance.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Project ID, default value is 0.
	ProjectId *int `pulumi:"projectId"`
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	RootPassword *string `pulumi:"rootPassword"`
	// Zone information about second slave instance.
	SecondSlaveZone *string `pulumi:"secondSlaveZone"`
	// Security groups to use.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode *int `pulumi:"slaveDeployMode"`
	// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
	SlaveSyncMode *int `pulumi:"slaveSyncMode"`
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Disk size (in GB).
	VolumeSize int `pulumi:"volumeSize"`
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrInput
	// Indicates which availability zone will be used.
	AvailabilityZone pulumi.StringPtrInput
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrInput
	// CPU cores.
	Cpu pulumi.IntPtrInput
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringPtrInput
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
	EngineVersion pulumi.StringPtrInput
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrInput
	// Zone information about first slave instance.
	FirstSlaveZone pulumi.StringPtrInput
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrInput
	// The name of a mysql instance.
	InstanceName pulumi.StringInput
	// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
	InternetService pulumi.IntPtrInput
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrInput
	// Memory size (in MB).
	MemSize pulumi.IntInput
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrInput
	// List of parameters to use.
	Parameters pulumi.MapInput
	// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
	PayType pulumi.IntPtrInput
	// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
	//
	// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
	Period pulumi.IntPtrInput
	// Period of instance. NOTES: Only supported prepaid instance.
	PrepaidPeriod pulumi.IntPtrInput
	// Project ID, default value is 0.
	ProjectId pulumi.IntPtrInput
	// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
	RootPassword pulumi.StringPtrInput
	// Zone information about second slave instance.
	SecondSlaveZone pulumi.StringPtrInput
	// Security groups to use.
	SecurityGroups pulumi.StringArrayInput
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrInput
	// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
	SlaveSyncMode pulumi.IntPtrInput
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// Disk size (in GB).
	VolumeSize pulumi.IntInput
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
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
//	InstanceArray{ InstanceArgs{...} }
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
//	InstanceMap{ "key": InstanceArgs{...} }
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

// Auto renew flag. NOTES: Only supported prepaid instance.
func (o InstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// Indicates which availability zone will be used.
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
func (o InstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// CPU cores.
func (o InstanceOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Cpu }).(pulumi.IntOutput)
}

// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
func (o InstanceOutput) DeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DeviceType }).(pulumi.StringPtrOutput)
}

// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.
func (o InstanceOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
func (o InstanceOutput) FastUpgrade() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.FastUpgrade }).(pulumi.IntPtrOutput)
}

// Zone information about first slave instance.
func (o InstanceOutput) FirstSlaveZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.FirstSlaveZone }).(pulumi.StringPtrOutput)
}

// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
func (o InstanceOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// Indicates whether GTID is enable. `0` - Not enabled; `1` - Enabled.
func (o InstanceOutput) Gtid() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Gtid }).(pulumi.IntOutput)
}

// The name of a mysql instance.
func (o InstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// host for public access.
func (o InstanceOutput) InternetHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InternetHost }).(pulumi.StringOutput)
}

// Access port for public access.
func (o InstanceOutput) InternetPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.InternetPort }).(pulumi.IntOutput)
}

// Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.
func (o InstanceOutput) InternetService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.InternetService }).(pulumi.IntPtrOutput)
}

// instance intranet IP.
func (o InstanceOutput) IntranetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.IntranetIp }).(pulumi.StringOutput)
}

// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
func (o InstanceOutput) IntranetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.IntranetPort }).(pulumi.IntPtrOutput)
}

// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
func (o InstanceOutput) Locked() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Locked }).(pulumi.IntOutput)
}

// Memory size (in MB).
func (o InstanceOutput) MemSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.MemSize }).(pulumi.IntOutput)
}

// Specify parameter template id.
func (o InstanceOutput) ParamTemplateId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ParamTemplateId }).(pulumi.IntPtrOutput)
}

// List of parameters to use.
func (o InstanceOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Parameters }).(pulumi.MapOutput)
}

// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
//
// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
func (o InstanceOutput) PayType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PayType }).(pulumi.IntPtrOutput)
}

// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
//
// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Period of instance. NOTES: Only supported prepaid instance.
func (o InstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// Project ID, default value is 0.
func (o InstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
func (o InstanceOutput) RootPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RootPassword }).(pulumi.StringPtrOutput)
}

// Zone information about second slave instance.
func (o InstanceOutput) SecondSlaveZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SecondSlaveZone }).(pulumi.StringPtrOutput)
}

// Security groups to use.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
func (o InstanceOutput) SlaveDeployMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.SlaveDeployMode }).(pulumi.IntPtrOutput)
}

// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
func (o InstanceOutput) SlaveSyncMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.SlaveSyncMode }).(pulumi.IntPtrOutput)
}

// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
func (o InstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Private network ID. If `vpcId` is set, this value is required.
func (o InstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Instance tags.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Indicates which kind of operations is being executed.
func (o InstanceOutput) TaskStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TaskStatus }).(pulumi.IntOutput)
}

// Disk size (in GB).
func (o InstanceOutput) VolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.VolumeSize }).(pulumi.IntOutput)
}

// ID of VPC, which can be modified once every 24 hours and can't be removed.
func (o InstanceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
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
