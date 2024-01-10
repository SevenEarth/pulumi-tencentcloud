// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql localBinlogConfig
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
//					"character_set_server": pulumi.Any("utf8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewLocalBinlogConfig(ctx, "exampleLocalBinlogConfig", &Mysql.LocalBinlogConfigArgs{
//				InstanceId: exampleInstance.ID(),
//				SaveHours:  pulumi.Int(140),
//				MaxUsage:   pulumi.Int(50),
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
// mysql local_binlog_config can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig local_binlog_config instance_id
//
// ```
type LocalBinlogConfig struct {
	pulumi.CustomResourceState

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage pulumi.IntOutput `pulumi:"maxUsage"`
	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours pulumi.IntOutput `pulumi:"saveHours"`
}

// NewLocalBinlogConfig registers a new resource with the given unique name, arguments, and options.
func NewLocalBinlogConfig(ctx *pulumi.Context,
	name string, args *LocalBinlogConfigArgs, opts ...pulumi.ResourceOption) (*LocalBinlogConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MaxUsage == nil {
		return nil, errors.New("invalid value for required argument 'MaxUsage'")
	}
	if args.SaveHours == nil {
		return nil, errors.New("invalid value for required argument 'SaveHours'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LocalBinlogConfig
	err := ctx.RegisterResource("tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalBinlogConfig gets an existing LocalBinlogConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalBinlogConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalBinlogConfigState, opts ...pulumi.ResourceOption) (*LocalBinlogConfig, error) {
	var resource LocalBinlogConfig
	err := ctx.ReadResource("tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalBinlogConfig resources.
type localBinlogConfigState struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `pulumi:"instanceId"`
	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage *int `pulumi:"maxUsage"`
	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours *int `pulumi:"saveHours"`
}

type LocalBinlogConfigState struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId pulumi.StringPtrInput
	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage pulumi.IntPtrInput
	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours pulumi.IntPtrInput
}

func (LocalBinlogConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*localBinlogConfigState)(nil)).Elem()
}

type localBinlogConfigArgs struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId string `pulumi:"instanceId"`
	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage int `pulumi:"maxUsage"`
	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours int `pulumi:"saveHours"`
}

// The set of arguments for constructing a LocalBinlogConfig resource.
type LocalBinlogConfigArgs struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId pulumi.StringInput
	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage pulumi.IntInput
	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours pulumi.IntInput
}

func (LocalBinlogConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localBinlogConfigArgs)(nil)).Elem()
}

type LocalBinlogConfigInput interface {
	pulumi.Input

	ToLocalBinlogConfigOutput() LocalBinlogConfigOutput
	ToLocalBinlogConfigOutputWithContext(ctx context.Context) LocalBinlogConfigOutput
}

func (*LocalBinlogConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalBinlogConfig)(nil)).Elem()
}

func (i *LocalBinlogConfig) ToLocalBinlogConfigOutput() LocalBinlogConfigOutput {
	return i.ToLocalBinlogConfigOutputWithContext(context.Background())
}

func (i *LocalBinlogConfig) ToLocalBinlogConfigOutputWithContext(ctx context.Context) LocalBinlogConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBinlogConfigOutput)
}

// LocalBinlogConfigArrayInput is an input type that accepts LocalBinlogConfigArray and LocalBinlogConfigArrayOutput values.
// You can construct a concrete instance of `LocalBinlogConfigArrayInput` via:
//
//	LocalBinlogConfigArray{ LocalBinlogConfigArgs{...} }
type LocalBinlogConfigArrayInput interface {
	pulumi.Input

	ToLocalBinlogConfigArrayOutput() LocalBinlogConfigArrayOutput
	ToLocalBinlogConfigArrayOutputWithContext(context.Context) LocalBinlogConfigArrayOutput
}

type LocalBinlogConfigArray []LocalBinlogConfigInput

func (LocalBinlogConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalBinlogConfig)(nil)).Elem()
}

func (i LocalBinlogConfigArray) ToLocalBinlogConfigArrayOutput() LocalBinlogConfigArrayOutput {
	return i.ToLocalBinlogConfigArrayOutputWithContext(context.Background())
}

func (i LocalBinlogConfigArray) ToLocalBinlogConfigArrayOutputWithContext(ctx context.Context) LocalBinlogConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBinlogConfigArrayOutput)
}

// LocalBinlogConfigMapInput is an input type that accepts LocalBinlogConfigMap and LocalBinlogConfigMapOutput values.
// You can construct a concrete instance of `LocalBinlogConfigMapInput` via:
//
//	LocalBinlogConfigMap{ "key": LocalBinlogConfigArgs{...} }
type LocalBinlogConfigMapInput interface {
	pulumi.Input

	ToLocalBinlogConfigMapOutput() LocalBinlogConfigMapOutput
	ToLocalBinlogConfigMapOutputWithContext(context.Context) LocalBinlogConfigMapOutput
}

type LocalBinlogConfigMap map[string]LocalBinlogConfigInput

func (LocalBinlogConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalBinlogConfig)(nil)).Elem()
}

func (i LocalBinlogConfigMap) ToLocalBinlogConfigMapOutput() LocalBinlogConfigMapOutput {
	return i.ToLocalBinlogConfigMapOutputWithContext(context.Background())
}

func (i LocalBinlogConfigMap) ToLocalBinlogConfigMapOutputWithContext(ctx context.Context) LocalBinlogConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBinlogConfigMapOutput)
}

type LocalBinlogConfigOutput struct{ *pulumi.OutputState }

func (LocalBinlogConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalBinlogConfig)(nil)).Elem()
}

func (o LocalBinlogConfigOutput) ToLocalBinlogConfigOutput() LocalBinlogConfigOutput {
	return o
}

func (o LocalBinlogConfigOutput) ToLocalBinlogConfigOutputWithContext(ctx context.Context) LocalBinlogConfigOutput {
	return o
}

// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
func (o LocalBinlogConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalBinlogConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Space utilization of local binlog. Value range: [30,50].
func (o LocalBinlogConfigOutput) MaxUsage() pulumi.IntOutput {
	return o.ApplyT(func(v *LocalBinlogConfig) pulumi.IntOutput { return v.MaxUsage }).(pulumi.IntOutput)
}

// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
func (o LocalBinlogConfigOutput) SaveHours() pulumi.IntOutput {
	return o.ApplyT(func(v *LocalBinlogConfig) pulumi.IntOutput { return v.SaveHours }).(pulumi.IntOutput)
}

type LocalBinlogConfigArrayOutput struct{ *pulumi.OutputState }

func (LocalBinlogConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalBinlogConfig)(nil)).Elem()
}

func (o LocalBinlogConfigArrayOutput) ToLocalBinlogConfigArrayOutput() LocalBinlogConfigArrayOutput {
	return o
}

func (o LocalBinlogConfigArrayOutput) ToLocalBinlogConfigArrayOutputWithContext(ctx context.Context) LocalBinlogConfigArrayOutput {
	return o
}

func (o LocalBinlogConfigArrayOutput) Index(i pulumi.IntInput) LocalBinlogConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalBinlogConfig {
		return vs[0].([]*LocalBinlogConfig)[vs[1].(int)]
	}).(LocalBinlogConfigOutput)
}

type LocalBinlogConfigMapOutput struct{ *pulumi.OutputState }

func (LocalBinlogConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalBinlogConfig)(nil)).Elem()
}

func (o LocalBinlogConfigMapOutput) ToLocalBinlogConfigMapOutput() LocalBinlogConfigMapOutput {
	return o
}

func (o LocalBinlogConfigMapOutput) ToLocalBinlogConfigMapOutputWithContext(ctx context.Context) LocalBinlogConfigMapOutput {
	return o
}

func (o LocalBinlogConfigMapOutput) MapIndex(k pulumi.StringInput) LocalBinlogConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalBinlogConfig {
		return vs[0].(map[string]*LocalBinlogConfig)[vs[1].(string)]
	}).(LocalBinlogConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBinlogConfigInput)(nil)).Elem(), &LocalBinlogConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBinlogConfigArrayInput)(nil)).Elem(), LocalBinlogConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBinlogConfigMapInput)(nil)).Elem(), LocalBinlogConfigMap{})
	pulumi.RegisterOutputType(LocalBinlogConfigOutput{})
	pulumi.RegisterOutputType(LocalBinlogConfigArrayOutput{})
	pulumi.RegisterOutputType(LocalBinlogConfigMapOutput{})
}
