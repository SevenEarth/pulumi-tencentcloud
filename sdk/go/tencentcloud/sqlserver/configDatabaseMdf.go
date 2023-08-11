// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver configDatabaseMdf
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
// 			Product: "sqlserver",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
// 			AvailabilityZone: pulumi.String(zones.Zones[4].Name),
// 			VpcId:            vpc.ID(),
// 			CidrBlock:        pulumi.String("10.0.0.0/16"),
// 			IsMulticast:      pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
// 			Description: pulumi.String("desc."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBasicInstance, err := Sqlserver.NewBasicInstance(ctx, "exampleBasicInstance", &Sqlserver.BasicInstanceArgs{
// 			AvailabilityZone: pulumi.String(zones.Zones[4].Name),
// 			ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
// 			VpcId:            vpc.ID(),
// 			SubnetId:         subnet.ID(),
// 			ProjectId:        pulumi.Int(0),
// 			Memory:           pulumi.Int(4),
// 			Storage:          pulumi.Int(100),
// 			Cpu:              pulumi.Int(2),
// 			MachineType:      pulumi.String("CLOUD_PREMIUM"),
// 			MaintenanceWeekSets: pulumi.IntArray{
// 				pulumi.Int(1),
// 				pulumi.Int(2),
// 				pulumi.Int(3),
// 			},
// 			MaintenanceStartTime: pulumi.String("09:00"),
// 			MaintenanceTimeSpan:  pulumi.Int(3),
// 			SecurityGroups: pulumi.StringArray{
// 				securityGroup.ID(),
// 			},
// 			Tags: pulumi.AnyMap{
// 				"test": pulumi.Any("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDb, err := Sqlserver.NewDb(ctx, "exampleDb", &Sqlserver.DbArgs{
// 			InstanceId: exampleBasicInstance.ID(),
// 			Charset:    pulumi.String("Chinese_PRC_BIN"),
// 			Remark:     pulumi.String("test-remark"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Sqlserver.NewConfigDatabaseMdf(ctx, "exampleConfigDatabaseMdf", &Sqlserver.ConfigDatabaseMdfArgs{
// 			DbName:     exampleDb.Name,
// 			InstanceId: exampleBasicInstance.ID(),
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
// sqlserver config_database_mdf can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf example mssql-i9ma6oy7#tf_example_db
// ```
type ConfigDatabaseMdf struct {
	pulumi.CustomResourceState

	// Array of database names.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewConfigDatabaseMdf registers a new resource with the given unique name, arguments, and options.
func NewConfigDatabaseMdf(ctx *pulumi.Context,
	name string, args *ConfigDatabaseMdfArgs, opts ...pulumi.ResourceOption) (*ConfigDatabaseMdf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ConfigDatabaseMdf
	err := ctx.RegisterResource("tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigDatabaseMdf gets an existing ConfigDatabaseMdf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigDatabaseMdf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigDatabaseMdfState, opts ...pulumi.ResourceOption) (*ConfigDatabaseMdf, error) {
	var resource ConfigDatabaseMdf
	err := ctx.ReadResource("tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigDatabaseMdf resources.
type configDatabaseMdfState struct {
	// Array of database names.
	DbName *string `pulumi:"dbName"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type ConfigDatabaseMdfState struct {
	// Array of database names.
	DbName pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (ConfigDatabaseMdfState) ElementType() reflect.Type {
	return reflect.TypeOf((*configDatabaseMdfState)(nil)).Elem()
}

type configDatabaseMdfArgs struct {
	// Array of database names.
	DbName string `pulumi:"dbName"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a ConfigDatabaseMdf resource.
type ConfigDatabaseMdfArgs struct {
	// Array of database names.
	DbName pulumi.StringInput
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (ConfigDatabaseMdfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configDatabaseMdfArgs)(nil)).Elem()
}

type ConfigDatabaseMdfInput interface {
	pulumi.Input

	ToConfigDatabaseMdfOutput() ConfigDatabaseMdfOutput
	ToConfigDatabaseMdfOutputWithContext(ctx context.Context) ConfigDatabaseMdfOutput
}

func (*ConfigDatabaseMdf) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigDatabaseMdf)(nil)).Elem()
}

func (i *ConfigDatabaseMdf) ToConfigDatabaseMdfOutput() ConfigDatabaseMdfOutput {
	return i.ToConfigDatabaseMdfOutputWithContext(context.Background())
}

func (i *ConfigDatabaseMdf) ToConfigDatabaseMdfOutputWithContext(ctx context.Context) ConfigDatabaseMdfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseMdfOutput)
}

// ConfigDatabaseMdfArrayInput is an input type that accepts ConfigDatabaseMdfArray and ConfigDatabaseMdfArrayOutput values.
// You can construct a concrete instance of `ConfigDatabaseMdfArrayInput` via:
//
//          ConfigDatabaseMdfArray{ ConfigDatabaseMdfArgs{...} }
type ConfigDatabaseMdfArrayInput interface {
	pulumi.Input

	ToConfigDatabaseMdfArrayOutput() ConfigDatabaseMdfArrayOutput
	ToConfigDatabaseMdfArrayOutputWithContext(context.Context) ConfigDatabaseMdfArrayOutput
}

type ConfigDatabaseMdfArray []ConfigDatabaseMdfInput

func (ConfigDatabaseMdfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigDatabaseMdf)(nil)).Elem()
}

func (i ConfigDatabaseMdfArray) ToConfigDatabaseMdfArrayOutput() ConfigDatabaseMdfArrayOutput {
	return i.ToConfigDatabaseMdfArrayOutputWithContext(context.Background())
}

func (i ConfigDatabaseMdfArray) ToConfigDatabaseMdfArrayOutputWithContext(ctx context.Context) ConfigDatabaseMdfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseMdfArrayOutput)
}

// ConfigDatabaseMdfMapInput is an input type that accepts ConfigDatabaseMdfMap and ConfigDatabaseMdfMapOutput values.
// You can construct a concrete instance of `ConfigDatabaseMdfMapInput` via:
//
//          ConfigDatabaseMdfMap{ "key": ConfigDatabaseMdfArgs{...} }
type ConfigDatabaseMdfMapInput interface {
	pulumi.Input

	ToConfigDatabaseMdfMapOutput() ConfigDatabaseMdfMapOutput
	ToConfigDatabaseMdfMapOutputWithContext(context.Context) ConfigDatabaseMdfMapOutput
}

type ConfigDatabaseMdfMap map[string]ConfigDatabaseMdfInput

func (ConfigDatabaseMdfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigDatabaseMdf)(nil)).Elem()
}

func (i ConfigDatabaseMdfMap) ToConfigDatabaseMdfMapOutput() ConfigDatabaseMdfMapOutput {
	return i.ToConfigDatabaseMdfMapOutputWithContext(context.Background())
}

func (i ConfigDatabaseMdfMap) ToConfigDatabaseMdfMapOutputWithContext(ctx context.Context) ConfigDatabaseMdfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigDatabaseMdfMapOutput)
}

type ConfigDatabaseMdfOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseMdfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigDatabaseMdf)(nil)).Elem()
}

func (o ConfigDatabaseMdfOutput) ToConfigDatabaseMdfOutput() ConfigDatabaseMdfOutput {
	return o
}

func (o ConfigDatabaseMdfOutput) ToConfigDatabaseMdfOutputWithContext(ctx context.Context) ConfigDatabaseMdfOutput {
	return o
}

// Array of database names.
func (o ConfigDatabaseMdfOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigDatabaseMdf) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// Instance ID.
func (o ConfigDatabaseMdfOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigDatabaseMdf) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type ConfigDatabaseMdfArrayOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseMdfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigDatabaseMdf)(nil)).Elem()
}

func (o ConfigDatabaseMdfArrayOutput) ToConfigDatabaseMdfArrayOutput() ConfigDatabaseMdfArrayOutput {
	return o
}

func (o ConfigDatabaseMdfArrayOutput) ToConfigDatabaseMdfArrayOutputWithContext(ctx context.Context) ConfigDatabaseMdfArrayOutput {
	return o
}

func (o ConfigDatabaseMdfArrayOutput) Index(i pulumi.IntInput) ConfigDatabaseMdfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigDatabaseMdf {
		return vs[0].([]*ConfigDatabaseMdf)[vs[1].(int)]
	}).(ConfigDatabaseMdfOutput)
}

type ConfigDatabaseMdfMapOutput struct{ *pulumi.OutputState }

func (ConfigDatabaseMdfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigDatabaseMdf)(nil)).Elem()
}

func (o ConfigDatabaseMdfMapOutput) ToConfigDatabaseMdfMapOutput() ConfigDatabaseMdfMapOutput {
	return o
}

func (o ConfigDatabaseMdfMapOutput) ToConfigDatabaseMdfMapOutputWithContext(ctx context.Context) ConfigDatabaseMdfMapOutput {
	return o
}

func (o ConfigDatabaseMdfMapOutput) MapIndex(k pulumi.StringInput) ConfigDatabaseMdfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigDatabaseMdf {
		return vs[0].(map[string]*ConfigDatabaseMdf)[vs[1].(string)]
	}).(ConfigDatabaseMdfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseMdfInput)(nil)).Elem(), &ConfigDatabaseMdf{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseMdfArrayInput)(nil)).Elem(), ConfigDatabaseMdfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigDatabaseMdfMapInput)(nil)).Elem(), ConfigDatabaseMdfMap{})
	pulumi.RegisterOutputType(ConfigDatabaseMdfOutput{})
	pulumi.RegisterOutputType(ConfigDatabaseMdfArrayOutput{})
	pulumi.RegisterOutputType(ConfigDatabaseMdfMapOutput{})
}