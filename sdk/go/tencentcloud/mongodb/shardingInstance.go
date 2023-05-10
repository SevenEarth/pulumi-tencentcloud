// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a Mongodb sharding instance.
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
// 		_, err := Mongodb.NewShardingInstance(ctx, "mongodb", &Mongodb.ShardingInstanceArgs{
// 			AvailableZone: pulumi.String("ap-guangzhou-3"),
// 			EngineVersion: pulumi.String("MONGO_36_WT"),
// 			InstanceName:  pulumi.String("mongodb"),
// 			MachineType:   pulumi.String("HIO10G"),
// 			Memory:        pulumi.Int(4),
// 			MongosCpu:     pulumi.Int(1),
// 			MongosMemory:  pulumi.Int(2),
// 			MongosNodeNum: pulumi.Int(3),
// 			NodesPerShard: pulumi.Int(3),
// 			Password:      pulumi.String("password1234"),
// 			ProjectId:     pulumi.Int(0),
// 			ShardQuantity: pulumi.Int(2),
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
// Mongodb sharding instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mongodb/shardingInstance:ShardingInstance mongodb cmgo-41s6jwy4
// ```
type ShardingInstance struct {
	pulumi.CustomResourceState

	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	AvailabilityZoneLists pulumi.StringArrayOutput `pulumi:"availabilityZoneLists"`
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringOutput `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// Creation time of the Mongodb instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	HiddenZone pulumi.StringOutput `pulumi:"hiddenZone"`
	// Name of the Mongodb instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Number of mongos cpu.
	MongosCpu pulumi.IntOutput `pulumi:"mongosCpu"`
	// Mongos memory size in GB.
	MongosMemory pulumi.IntOutput `pulumi:"mongosMemory"`
	// Number of mongos.
	MongosNodeNum pulumi.IntOutput `pulumi:"mongosNodeNum"`
	// Number of nodes per shard, at least 3(one master and two slaves).
	NodesPerShard pulumi.IntOutput `pulumi:"nodesPerShard"`
	// Password of this Mongodb account.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrOutput `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Number of sharding.
	ShardQuantity pulumi.IntOutput `pulumi:"shardQuantity"`
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

// NewShardingInstance registers a new resource with the given unique name, arguments, and options.
func NewShardingInstance(ctx *pulumi.Context,
	name string, args *ShardingInstanceArgs, opts ...pulumi.ResourceOption) (*ShardingInstance, error) {
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
	if args.NodesPerShard == nil {
		return nil, errors.New("invalid value for required argument 'NodesPerShard'")
	}
	if args.ShardQuantity == nil {
		return nil, errors.New("invalid value for required argument 'ShardQuantity'")
	}
	if args.Volume == nil {
		return nil, errors.New("invalid value for required argument 'Volume'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ShardingInstance
	err := ctx.RegisterResource("tencentcloud:Mongodb/shardingInstance:ShardingInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShardingInstance gets an existing ShardingInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShardingInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShardingInstanceState, opts ...pulumi.ResourceOption) (*ShardingInstance, error) {
	var resource ShardingInstance
	err := ctx.ReadResource("tencentcloud:Mongodb/shardingInstance:ShardingInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShardingInstance resources.
type shardingInstanceState struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	AvailabilityZoneLists []string `pulumi:"availabilityZoneLists"`
	// The available zone of the Mongodb.
	AvailableZone *string `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// Creation time of the Mongodb instance.
	CreateTime *string `pulumi:"createTime"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion *string `pulumi:"engineVersion"`
	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	HiddenZone *string `pulumi:"hiddenZone"`
	// Name of the Mongodb instance.
	InstanceName *string `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType *string `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory *int `pulumi:"memory"`
	// Number of mongos cpu.
	MongosCpu *int `pulumi:"mongosCpu"`
	// Mongos memory size in GB.
	MongosMemory *int `pulumi:"mongosMemory"`
	// Number of mongos.
	MongosNodeNum *int `pulumi:"mongosNodeNum"`
	// Number of nodes per shard, at least 3(one master and two slaves).
	NodesPerShard *int `pulumi:"nodesPerShard"`
	// Password of this Mongodb account.
	Password *string `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Number of sharding.
	ShardQuantity *int `pulumi:"shardQuantity"`
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

type ShardingInstanceState struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrInput
	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	AvailabilityZoneLists pulumi.StringArrayInput
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringPtrInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrInput
	// Creation time of the Mongodb instance.
	CreateTime pulumi.StringPtrInput
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringPtrInput
	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	HiddenZone pulumi.StringPtrInput
	// Name of the Mongodb instance.
	InstanceName pulumi.StringPtrInput
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringPtrInput
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntPtrInput
	// Number of mongos cpu.
	MongosCpu pulumi.IntPtrInput
	// Mongos memory size in GB.
	MongosMemory pulumi.IntPtrInput
	// Number of mongos.
	MongosNodeNum pulumi.IntPtrInput
	// Number of nodes per shard, at least 3(one master and two slaves).
	NodesPerShard pulumi.IntPtrInput
	// Password of this Mongodb account.
	Password pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrInput
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayInput
	// Number of sharding.
	ShardQuantity pulumi.IntPtrInput
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

func (ShardingInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingInstanceState)(nil)).Elem()
}

type shardingInstanceArgs struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	AvailabilityZoneLists []string `pulumi:"availabilityZoneLists"`
	// The available zone of the Mongodb.
	AvailableZone string `pulumi:"availableZone"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion string `pulumi:"engineVersion"`
	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	HiddenZone *string `pulumi:"hiddenZone"`
	// Name of the Mongodb instance.
	InstanceName string `pulumi:"instanceName"`
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType string `pulumi:"machineType"`
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory int `pulumi:"memory"`
	// Number of mongos cpu.
	MongosCpu *int `pulumi:"mongosCpu"`
	// Mongos memory size in GB.
	MongosMemory *int `pulumi:"mongosMemory"`
	// Number of mongos.
	MongosNodeNum *int `pulumi:"mongosNodeNum"`
	// Number of nodes per shard, at least 3(one master and two slaves).
	NodesPerShard int `pulumi:"nodesPerShard"`
	// Password of this Mongodb account.
	Password *string `pulumi:"password"`
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// ID of the project which the instance belongs.
	ProjectId *int `pulumi:"projectId"`
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Number of sharding.
	ShardQuantity int `pulumi:"shardQuantity"`
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId *string `pulumi:"subnetId"`
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags map[string]interface{} `pulumi:"tags"`
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume int `pulumi:"volume"`
	// ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ShardingInstance resource.
type ShardingInstanceArgs struct {
	// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
	AutoRenewFlag pulumi.IntPtrInput
	// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
	// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
	// - Version 4.2 and above are not supported.
	// - Read-only disaster recovery instances are not supported.
	// - Basic network cannot be selected.
	AvailabilityZoneLists pulumi.StringArrayInput
	// The available zone of the Mongodb.
	AvailableZone pulumi.StringInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
	ChargeType pulumi.StringPtrInput
	// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
	EngineVersion pulumi.StringInput
	// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
	HiddenZone pulumi.StringPtrInput
	// Name of the Mongodb instance.
	InstanceName pulumi.StringInput
	// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
	MachineType pulumi.StringInput
	// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Memory pulumi.IntInput
	// Number of mongos cpu.
	MongosCpu pulumi.IntPtrInput
	// Mongos memory size in GB.
	MongosMemory pulumi.IntPtrInput
	// Number of mongos.
	MongosNodeNum pulumi.IntPtrInput
	// Number of nodes per shard, at least 3(one master and two slaves).
	NodesPerShard pulumi.IntInput
	// Password of this Mongodb account.
	Password pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
	PrepaidPeriod pulumi.IntPtrInput
	// ID of the project which the instance belongs.
	ProjectId pulumi.IntPtrInput
	// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
	SecurityGroups pulumi.StringArrayInput
	// Number of sharding.
	ShardQuantity pulumi.IntInput
	// ID of the subnet within this VPC. The value is required if `vpcId` is set.
	SubnetId pulumi.StringPtrInput
	// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
	Tags pulumi.MapInput
	// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
	Volume pulumi.IntInput
	// ID of the VPC.
	VpcId pulumi.StringPtrInput
}

func (ShardingInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingInstanceArgs)(nil)).Elem()
}

type ShardingInstanceInput interface {
	pulumi.Input

	ToShardingInstanceOutput() ShardingInstanceOutput
	ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput
}

func (*ShardingInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingInstance)(nil)).Elem()
}

func (i *ShardingInstance) ToShardingInstanceOutput() ShardingInstanceOutput {
	return i.ToShardingInstanceOutputWithContext(context.Background())
}

func (i *ShardingInstance) ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceOutput)
}

// ShardingInstanceArrayInput is an input type that accepts ShardingInstanceArray and ShardingInstanceArrayOutput values.
// You can construct a concrete instance of `ShardingInstanceArrayInput` via:
//
//          ShardingInstanceArray{ ShardingInstanceArgs{...} }
type ShardingInstanceArrayInput interface {
	pulumi.Input

	ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput
	ToShardingInstanceArrayOutputWithContext(context.Context) ShardingInstanceArrayOutput
}

type ShardingInstanceArray []ShardingInstanceInput

func (ShardingInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingInstance)(nil)).Elem()
}

func (i ShardingInstanceArray) ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput {
	return i.ToShardingInstanceArrayOutputWithContext(context.Background())
}

func (i ShardingInstanceArray) ToShardingInstanceArrayOutputWithContext(ctx context.Context) ShardingInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceArrayOutput)
}

// ShardingInstanceMapInput is an input type that accepts ShardingInstanceMap and ShardingInstanceMapOutput values.
// You can construct a concrete instance of `ShardingInstanceMapInput` via:
//
//          ShardingInstanceMap{ "key": ShardingInstanceArgs{...} }
type ShardingInstanceMapInput interface {
	pulumi.Input

	ToShardingInstanceMapOutput() ShardingInstanceMapOutput
	ToShardingInstanceMapOutputWithContext(context.Context) ShardingInstanceMapOutput
}

type ShardingInstanceMap map[string]ShardingInstanceInput

func (ShardingInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingInstance)(nil)).Elem()
}

func (i ShardingInstanceMap) ToShardingInstanceMapOutput() ShardingInstanceMapOutput {
	return i.ToShardingInstanceMapOutputWithContext(context.Background())
}

func (i ShardingInstanceMap) ToShardingInstanceMapOutputWithContext(ctx context.Context) ShardingInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceMapOutput)
}

type ShardingInstanceOutput struct{ *pulumi.OutputState }

func (ShardingInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingInstance)(nil)).Elem()
}

func (o ShardingInstanceOutput) ToShardingInstanceOutput() ShardingInstanceOutput {
	return o
}

func (o ShardingInstanceOutput) ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput {
	return o
}

// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
func (o ShardingInstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
// - Version 4.2 and above are not supported.
// - Read-only disaster recovery instances are not supported.
// - Basic network cannot be selected.
func (o ShardingInstanceOutput) AvailabilityZoneLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringArrayOutput { return v.AvailabilityZoneLists }).(pulumi.StringArrayOutput)
}

// The available zone of the Mongodb.
func (o ShardingInstanceOutput) AvailableZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.AvailableZone }).(pulumi.StringOutput)
}

// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
func (o ShardingInstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// Creation time of the Mongodb instance.
func (o ShardingInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
func (o ShardingInstanceOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
func (o ShardingInstanceOutput) HiddenZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.HiddenZone }).(pulumi.StringOutput)
}

// Name of the Mongodb instance.
func (o ShardingInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
func (o ShardingInstanceOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.MachineType }).(pulumi.StringOutput)
}

// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
func (o ShardingInstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// Number of mongos cpu.
func (o ShardingInstanceOutput) MongosCpu() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.MongosCpu }).(pulumi.IntOutput)
}

// Mongos memory size in GB.
func (o ShardingInstanceOutput) MongosMemory() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.MongosMemory }).(pulumi.IntOutput)
}

// Number of mongos.
func (o ShardingInstanceOutput) MongosNodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.MongosNodeNum }).(pulumi.IntOutput)
}

// Number of nodes per shard, at least 3(one master and two slaves).
func (o ShardingInstanceOutput) NodesPerShard() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.NodesPerShard }).(pulumi.IntOutput)
}

// Password of this Mongodb account.
func (o ShardingInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when chargeType is set to `PREPAID`.
func (o ShardingInstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// ID of the project which the instance belongs.
func (o ShardingInstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// ID of the security group. NOTE: for instance which `engineVersion` is `MONGO_40_WT`, `securityGroups` is not supported.
func (o ShardingInstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Number of sharding.
func (o ShardingInstanceOutput) ShardQuantity() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.ShardQuantity }).(pulumi.IntOutput)
}

// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
func (o ShardingInstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// ID of the subnet within this VPC. The value is required if `vpcId` is set.
func (o ShardingInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
func (o ShardingInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// IP of the Mongodb instance.
func (o ShardingInstanceOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
func (o ShardingInstanceOutput) Volume() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.Volume }).(pulumi.IntOutput)
}

// ID of the VPC.
func (o ShardingInstanceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

// IP port of the Mongodb instance.
func (o ShardingInstanceOutput) Vport() pulumi.IntOutput {
	return o.ApplyT(func(v *ShardingInstance) pulumi.IntOutput { return v.Vport }).(pulumi.IntOutput)
}

type ShardingInstanceArrayOutput struct{ *pulumi.OutputState }

func (ShardingInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingInstance)(nil)).Elem()
}

func (o ShardingInstanceArrayOutput) ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput {
	return o
}

func (o ShardingInstanceArrayOutput) ToShardingInstanceArrayOutputWithContext(ctx context.Context) ShardingInstanceArrayOutput {
	return o
}

func (o ShardingInstanceArrayOutput) Index(i pulumi.IntInput) ShardingInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShardingInstance {
		return vs[0].([]*ShardingInstance)[vs[1].(int)]
	}).(ShardingInstanceOutput)
}

type ShardingInstanceMapOutput struct{ *pulumi.OutputState }

func (ShardingInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingInstance)(nil)).Elem()
}

func (o ShardingInstanceMapOutput) ToShardingInstanceMapOutput() ShardingInstanceMapOutput {
	return o
}

func (o ShardingInstanceMapOutput) ToShardingInstanceMapOutputWithContext(ctx context.Context) ShardingInstanceMapOutput {
	return o
}

func (o ShardingInstanceMapOutput) MapIndex(k pulumi.StringInput) ShardingInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShardingInstance {
		return vs[0].(map[string]*ShardingInstance)[vs[1].(string)]
	}).(ShardingInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingInstanceInput)(nil)).Elem(), &ShardingInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingInstanceArrayInput)(nil)).Elem(), ShardingInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingInstanceMapInput)(nil)).Elem(), ShardingInstanceMap{})
	pulumi.RegisterOutputType(ShardingInstanceOutput{})
	pulumi.RegisterOutputType(ShardingInstanceArrayOutput{})
	pulumi.RegisterOutputType(ShardingInstanceMapOutput{})
}
