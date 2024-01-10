// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a mysql instance resource to create read-only database instances.
//
// > **NOTE:** Read-only instances can be purchased only for two-node or three-node source instances on MySQL 5.6 or above with the InnoDB engine at a specification of 1 GB memory and 50 GB disk capacity or above.
// **NOTE:** The terminate operation of read only mysql does NOT take effect immediately, maybe takes for several hours. so during that time, VPCs associated with that mysql instance can't be terminated also.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "cdb",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("mysql test"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := Mysql.NewInstance(ctx, "exampleInstance", &Mysql.InstanceArgs{
//				InternetService:  pulumi.Int(1),
//				EngineVersion:    pulumi.String("5.7"),
//				ChargeType:       pulumi.String("POSTPAID"),
//				RootPassword:     pulumi.String("PassWord123"),
//				SlaveDeployMode:  pulumi.Int(0),
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				SlaveSyncMode:    pulumi.Int(1),
//				InstanceName:     pulumi.String("tf-example-mysql"),
//				MemSize:          pulumi.Int(4000),
//				VolumeSize:       pulumi.Int(200),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				IntranetPort:     pulumi.Int(3306),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"name": pulumi.Any("test"),
//				},
//				Parameters: pulumi.AnyMap{
//					"character_set_server": pulumi.Any("UTF8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewReadonlyInstance(ctx, "exampleReadonlyInstance", &Mysql.ReadonlyInstanceArgs{
//				MasterInstanceId: exampleInstance.ID(),
//				InstanceName:     pulumi.String("tf-example"),
//				MemSize:          pulumi.Int(128000),
//				VolumeSize:       pulumi.Int(255),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				IntranetPort:     pulumi.Int(3306),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"createBy": pulumi.Any("terraform"),
//				},
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
// mysql read-only database instances can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Mysql/readonlyInstance:ReadonlyInstance default cdb-dnqksd9f
//
// ```
type ReadonlyInstance struct {
	pulumi.CustomResourceState

	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// CPU cores.
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrOutput `pulumi:"fastUpgrade"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// The name of a mysql instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// instance intranet IP.
	IntranetIp pulumi.StringOutput `pulumi:"intranetIp"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrOutput `pulumi:"intranetPort"`
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked pulumi.IntOutput `pulumi:"locked"`
	// Indicates the master instance ID of recovery instances.
	MasterInstanceId pulumi.StringOutput `pulumi:"masterInstanceId"`
	// The zone information of the primary instance is required when you purchase a disaster recovery instance.
	MasterRegion pulumi.StringOutput `pulumi:"masterRegion"`
	// Memory size (in MB).
	MemSize pulumi.IntOutput `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrOutput `pulumi:"paramTemplateId"`
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
	// Security groups to use.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrOutput `pulumi:"slaveDeployMode"`
	// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
	Status pulumi.IntOutput `pulumi:"status"`
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Instance tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Indicates which kind of operations is being executed.
	TaskStatus pulumi.IntOutput `pulumi:"taskStatus"`
	// Disk size (in GB).
	VolumeSize pulumi.IntOutput `pulumi:"volumeSize"`
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
	WaitSwitch pulumi.IntPtrOutput `pulumi:"waitSwitch"`
	// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewReadonlyInstance registers a new resource with the given unique name, arguments, and options.
func NewReadonlyInstance(ctx *pulumi.Context,
	name string, args *ReadonlyInstanceArgs, opts ...pulumi.ResourceOption) (*ReadonlyInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.MasterInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'MasterInstanceId'")
	}
	if args.MemSize == nil {
		return nil, errors.New("invalid value for required argument 'MemSize'")
	}
	if args.VolumeSize == nil {
		return nil, errors.New("invalid value for required argument 'VolumeSize'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReadonlyInstance
	err := ctx.RegisterResource("tencentcloud:Mysql/readonlyInstance:ReadonlyInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadonlyInstance gets an existing ReadonlyInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadonlyInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadonlyInstanceState, opts ...pulumi.ResourceOption) (*ReadonlyInstance, error) {
	var resource ReadonlyInstance
	err := ctx.ReadResource("tencentcloud:Mysql/readonlyInstance:ReadonlyInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadonlyInstance resources.
type readonlyInstanceState struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// CPU cores.
	Cpu *int `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType *string `pulumi:"deviceType"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade *int `pulumi:"fastUpgrade"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The name of a mysql instance.
	InstanceName *string `pulumi:"instanceName"`
	// instance intranet IP.
	IntranetIp *string `pulumi:"intranetIp"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort *int `pulumi:"intranetPort"`
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked *int `pulumi:"locked"`
	// Indicates the master instance ID of recovery instances.
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	// The zone information of the primary instance is required when you purchase a disaster recovery instance.
	MasterRegion *string `pulumi:"masterRegion"`
	// Memory size (in MB).
	MemSize *int `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId *int `pulumi:"paramTemplateId"`
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
	// Security groups to use.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode *int `pulumi:"slaveDeployMode"`
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
	// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
	WaitSwitch *int `pulumi:"waitSwitch"`
	// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
	Zone *string `pulumi:"zone"`
}

type ReadonlyInstanceState struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrInput
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrInput
	// CPU cores.
	Cpu pulumi.IntPtrInput
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringPtrInput
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrInput
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrInput
	// The name of a mysql instance.
	InstanceName pulumi.StringPtrInput
	// instance intranet IP.
	IntranetIp pulumi.StringPtrInput
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrInput
	// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
	Locked pulumi.IntPtrInput
	// Indicates the master instance ID of recovery instances.
	MasterInstanceId pulumi.StringPtrInput
	// The zone information of the primary instance is required when you purchase a disaster recovery instance.
	MasterRegion pulumi.StringPtrInput
	// Memory size (in MB).
	MemSize pulumi.IntPtrInput
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrInput
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
	// Security groups to use.
	SecurityGroups pulumi.StringArrayInput
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrInput
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
	// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
	WaitSwitch pulumi.IntPtrInput
	// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
	Zone pulumi.StringPtrInput
}

func (ReadonlyInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyInstanceState)(nil)).Elem()
}

type readonlyInstanceArgs struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// CPU cores.
	Cpu *int `pulumi:"cpu"`
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType *string `pulumi:"deviceType"`
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade *int `pulumi:"fastUpgrade"`
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The name of a mysql instance.
	InstanceName string `pulumi:"instanceName"`
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort *int `pulumi:"intranetPort"`
	// Indicates the master instance ID of recovery instances.
	MasterInstanceId string `pulumi:"masterInstanceId"`
	// The zone information of the primary instance is required when you purchase a disaster recovery instance.
	MasterRegion *string `pulumi:"masterRegion"`
	// Memory size (in MB).
	MemSize int `pulumi:"memSize"`
	// Specify parameter template id.
	ParamTemplateId *int `pulumi:"paramTemplateId"`
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
	// Security groups to use.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode *int `pulumi:"slaveDeployMode"`
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Disk size (in GB).
	VolumeSize int `pulumi:"volumeSize"`
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId *string `pulumi:"vpcId"`
	// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
	WaitSwitch *int `pulumi:"waitSwitch"`
	// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ReadonlyInstance resource.
type ReadonlyInstanceArgs struct {
	// Auto renew flag. NOTES: Only supported prepaid instance.
	AutoRenewFlag pulumi.IntPtrInput
	// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
	ChargeType pulumi.StringPtrInput
	// CPU cores.
	Cpu pulumi.IntPtrInput
	// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
	DeviceType pulumi.StringPtrInput
	// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
	FastUpgrade pulumi.IntPtrInput
	// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
	ForceDelete pulumi.BoolPtrInput
	// The name of a mysql instance.
	InstanceName pulumi.StringInput
	// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
	IntranetPort pulumi.IntPtrInput
	// Indicates the master instance ID of recovery instances.
	MasterInstanceId pulumi.StringInput
	// The zone information of the primary instance is required when you purchase a disaster recovery instance.
	MasterRegion pulumi.StringPtrInput
	// Memory size (in MB).
	MemSize pulumi.IntInput
	// Specify parameter template id.
	ParamTemplateId pulumi.IntPtrInput
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
	// Security groups to use.
	SecurityGroups pulumi.StringArrayInput
	// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
	SlaveDeployMode pulumi.IntPtrInput
	// Private network ID. If `vpcId` is set, this value is required.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// Disk size (in GB).
	VolumeSize pulumi.IntInput
	// ID of VPC, which can be modified once every 24 hours and can't be removed.
	VpcId pulumi.StringPtrInput
	// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
	WaitSwitch pulumi.IntPtrInput
	// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
	Zone pulumi.StringPtrInput
}

func (ReadonlyInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyInstanceArgs)(nil)).Elem()
}

type ReadonlyInstanceInput interface {
	pulumi.Input

	ToReadonlyInstanceOutput() ReadonlyInstanceOutput
	ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput
}

func (*ReadonlyInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyInstance)(nil)).Elem()
}

func (i *ReadonlyInstance) ToReadonlyInstanceOutput() ReadonlyInstanceOutput {
	return i.ToReadonlyInstanceOutputWithContext(context.Background())
}

func (i *ReadonlyInstance) ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceOutput)
}

// ReadonlyInstanceArrayInput is an input type that accepts ReadonlyInstanceArray and ReadonlyInstanceArrayOutput values.
// You can construct a concrete instance of `ReadonlyInstanceArrayInput` via:
//
//	ReadonlyInstanceArray{ ReadonlyInstanceArgs{...} }
type ReadonlyInstanceArrayInput interface {
	pulumi.Input

	ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput
	ToReadonlyInstanceArrayOutputWithContext(context.Context) ReadonlyInstanceArrayOutput
}

type ReadonlyInstanceArray []ReadonlyInstanceInput

func (ReadonlyInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyInstance)(nil)).Elem()
}

func (i ReadonlyInstanceArray) ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput {
	return i.ToReadonlyInstanceArrayOutputWithContext(context.Background())
}

func (i ReadonlyInstanceArray) ToReadonlyInstanceArrayOutputWithContext(ctx context.Context) ReadonlyInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceArrayOutput)
}

// ReadonlyInstanceMapInput is an input type that accepts ReadonlyInstanceMap and ReadonlyInstanceMapOutput values.
// You can construct a concrete instance of `ReadonlyInstanceMapInput` via:
//
//	ReadonlyInstanceMap{ "key": ReadonlyInstanceArgs{...} }
type ReadonlyInstanceMapInput interface {
	pulumi.Input

	ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput
	ToReadonlyInstanceMapOutputWithContext(context.Context) ReadonlyInstanceMapOutput
}

type ReadonlyInstanceMap map[string]ReadonlyInstanceInput

func (ReadonlyInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyInstance)(nil)).Elem()
}

func (i ReadonlyInstanceMap) ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput {
	return i.ToReadonlyInstanceMapOutputWithContext(context.Background())
}

func (i ReadonlyInstanceMap) ToReadonlyInstanceMapOutputWithContext(ctx context.Context) ReadonlyInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceMapOutput)
}

type ReadonlyInstanceOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceOutput) ToReadonlyInstanceOutput() ReadonlyInstanceOutput {
	return o
}

func (o ReadonlyInstanceOutput) ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput {
	return o
}

// Auto renew flag. NOTES: Only supported prepaid instance.
func (o ReadonlyInstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
func (o ReadonlyInstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// CPU cores.
func (o ReadonlyInstanceOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.Cpu }).(pulumi.IntOutput)
}

// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
func (o ReadonlyInstanceOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
func (o ReadonlyInstanceOutput) FastUpgrade() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.FastUpgrade }).(pulumi.IntPtrOutput)
}

// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
func (o ReadonlyInstanceOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// The name of a mysql instance.
func (o ReadonlyInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// instance intranet IP.
func (o ReadonlyInstanceOutput) IntranetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.IntranetIp }).(pulumi.StringOutput)
}

// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
func (o ReadonlyInstanceOutput) IntranetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.IntranetPort }).(pulumi.IntPtrOutput)
}

// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
func (o ReadonlyInstanceOutput) Locked() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.Locked }).(pulumi.IntOutput)
}

// Indicates the master instance ID of recovery instances.
func (o ReadonlyInstanceOutput) MasterInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.MasterInstanceId }).(pulumi.StringOutput)
}

// The zone information of the primary instance is required when you purchase a disaster recovery instance.
func (o ReadonlyInstanceOutput) MasterRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.MasterRegion }).(pulumi.StringOutput)
}

// Memory size (in MB).
func (o ReadonlyInstanceOutput) MemSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.MemSize }).(pulumi.IntOutput)
}

// Specify parameter template id.
func (o ReadonlyInstanceOutput) ParamTemplateId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.ParamTemplateId }).(pulumi.IntPtrOutput)
}

// It has been deprecated from version 1.36.0. Please use `chargeType` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
//
// Deprecated: It has been deprecated from version 1.36.0. Please use `charge_type` instead.
func (o ReadonlyInstanceOutput) PayType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.PayType }).(pulumi.IntPtrOutput)
}

// It has been deprecated from version 1.36.0. Please use `prepaidPeriod` instead. Period of instance. NOTES: Only supported prepaid instance.
//
// Deprecated: It has been deprecated from version 1.36.0. Please use `prepaid_period` instead.
func (o ReadonlyInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Period of instance. NOTES: Only supported prepaid instance.
func (o ReadonlyInstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// Security groups to use.
func (o ReadonlyInstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
func (o ReadonlyInstanceOutput) SlaveDeployMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.SlaveDeployMode }).(pulumi.IntPtrOutput)
}

// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
func (o ReadonlyInstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Private network ID. If `vpcId` is set, this value is required.
func (o ReadonlyInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Instance tags.
func (o ReadonlyInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Indicates which kind of operations is being executed.
func (o ReadonlyInstanceOutput) TaskStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.TaskStatus }).(pulumi.IntOutput)
}

// Disk size (in GB).
func (o ReadonlyInstanceOutput) VolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.VolumeSize }).(pulumi.IntOutput)
}

// ID of VPC, which can be modified once every 24 hours and can't be removed.
func (o ReadonlyInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Switch the method of accessing new instances, default is `0`. Supported values include: `0` - switch immediately, `1` - switch in time window.
func (o ReadonlyInstanceOutput) WaitSwitch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.WaitSwitch }).(pulumi.IntPtrOutput)
}

// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
func (o ReadonlyInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ReadonlyInstanceArrayOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceArrayOutput) ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput {
	return o
}

func (o ReadonlyInstanceArrayOutput) ToReadonlyInstanceArrayOutputWithContext(ctx context.Context) ReadonlyInstanceArrayOutput {
	return o
}

func (o ReadonlyInstanceArrayOutput) Index(i pulumi.IntInput) ReadonlyInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReadonlyInstance {
		return vs[0].([]*ReadonlyInstance)[vs[1].(int)]
	}).(ReadonlyInstanceOutput)
}

type ReadonlyInstanceMapOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceMapOutput) ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput {
	return o
}

func (o ReadonlyInstanceMapOutput) ToReadonlyInstanceMapOutputWithContext(ctx context.Context) ReadonlyInstanceMapOutput {
	return o
}

func (o ReadonlyInstanceMapOutput) MapIndex(k pulumi.StringInput) ReadonlyInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReadonlyInstance {
		return vs[0].(map[string]*ReadonlyInstance)[vs[1].(string)]
	}).(ReadonlyInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceInput)(nil)).Elem(), &ReadonlyInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceArrayInput)(nil)).Elem(), ReadonlyInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceMapInput)(nil)).Elem(), ReadonlyInstanceMap{})
	pulumi.RegisterOutputType(ReadonlyInstanceOutput{})
	pulumi.RegisterOutputType(ReadonlyInstanceArrayOutput{})
	pulumi.RegisterOutputType(ReadonlyInstanceMapOutput{})
}
