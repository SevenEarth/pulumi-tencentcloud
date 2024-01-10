// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver configTerminateDbInstance
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "sqlserver",
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
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("desc."),
//			})
//			if err != nil {
//				return err
//			}
//			exampleBasicInstance, err := Sqlserver.NewBasicInstance(ctx, "exampleBasicInstance", &Sqlserver.BasicInstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				ProjectId:        pulumi.Int(0),
//				Memory:           pulumi.Int(4),
//				Storage:          pulumi.Int(100),
//				Cpu:              pulumi.Int(2),
//				MachineType:      pulumi.String("CLOUD_PREMIUM"),
//				MaintenanceWeekSets: pulumi.IntArray{
//					pulumi.Int(1),
//					pulumi.Int(2),
//					pulumi.Int(3),
//				},
//				MaintenanceStartTime: pulumi.String("09:00"),
//				MaintenanceTimeSpan:  pulumi.Int(3),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewConfigTerminateDbInstance(ctx, "exampleConfigTerminateDbInstance", &Sqlserver.ConfigTerminateDbInstanceArgs{
//				InstanceId: exampleBasicInstance.ID(),
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
// sqlserver config_terminate_db_instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Sqlserver/configTerminateDbInstance:ConfigTerminateDbInstance example mssql-i9ma6oy7
//
// ```
type ConfigTerminateDbInstance struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewConfigTerminateDbInstance registers a new resource with the given unique name, arguments, and options.
func NewConfigTerminateDbInstance(ctx *pulumi.Context,
	name string, args *ConfigTerminateDbInstanceArgs, opts ...pulumi.ResourceOption) (*ConfigTerminateDbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ConfigTerminateDbInstance
	err := ctx.RegisterResource("tencentcloud:Sqlserver/configTerminateDbInstance:ConfigTerminateDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigTerminateDbInstance gets an existing ConfigTerminateDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigTerminateDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigTerminateDbInstanceState, opts ...pulumi.ResourceOption) (*ConfigTerminateDbInstance, error) {
	var resource ConfigTerminateDbInstance
	err := ctx.ReadResource("tencentcloud:Sqlserver/configTerminateDbInstance:ConfigTerminateDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigTerminateDbInstance resources.
type configTerminateDbInstanceState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type ConfigTerminateDbInstanceState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (ConfigTerminateDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configTerminateDbInstanceState)(nil)).Elem()
}

type configTerminateDbInstanceArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a ConfigTerminateDbInstance resource.
type ConfigTerminateDbInstanceArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (ConfigTerminateDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configTerminateDbInstanceArgs)(nil)).Elem()
}

type ConfigTerminateDbInstanceInput interface {
	pulumi.Input

	ToConfigTerminateDbInstanceOutput() ConfigTerminateDbInstanceOutput
	ToConfigTerminateDbInstanceOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceOutput
}

func (*ConfigTerminateDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigTerminateDbInstance)(nil)).Elem()
}

func (i *ConfigTerminateDbInstance) ToConfigTerminateDbInstanceOutput() ConfigTerminateDbInstanceOutput {
	return i.ToConfigTerminateDbInstanceOutputWithContext(context.Background())
}

func (i *ConfigTerminateDbInstance) ToConfigTerminateDbInstanceOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigTerminateDbInstanceOutput)
}

// ConfigTerminateDbInstanceArrayInput is an input type that accepts ConfigTerminateDbInstanceArray and ConfigTerminateDbInstanceArrayOutput values.
// You can construct a concrete instance of `ConfigTerminateDbInstanceArrayInput` via:
//
//	ConfigTerminateDbInstanceArray{ ConfigTerminateDbInstanceArgs{...} }
type ConfigTerminateDbInstanceArrayInput interface {
	pulumi.Input

	ToConfigTerminateDbInstanceArrayOutput() ConfigTerminateDbInstanceArrayOutput
	ToConfigTerminateDbInstanceArrayOutputWithContext(context.Context) ConfigTerminateDbInstanceArrayOutput
}

type ConfigTerminateDbInstanceArray []ConfigTerminateDbInstanceInput

func (ConfigTerminateDbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigTerminateDbInstance)(nil)).Elem()
}

func (i ConfigTerminateDbInstanceArray) ToConfigTerminateDbInstanceArrayOutput() ConfigTerminateDbInstanceArrayOutput {
	return i.ToConfigTerminateDbInstanceArrayOutputWithContext(context.Background())
}

func (i ConfigTerminateDbInstanceArray) ToConfigTerminateDbInstanceArrayOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigTerminateDbInstanceArrayOutput)
}

// ConfigTerminateDbInstanceMapInput is an input type that accepts ConfigTerminateDbInstanceMap and ConfigTerminateDbInstanceMapOutput values.
// You can construct a concrete instance of `ConfigTerminateDbInstanceMapInput` via:
//
//	ConfigTerminateDbInstanceMap{ "key": ConfigTerminateDbInstanceArgs{...} }
type ConfigTerminateDbInstanceMapInput interface {
	pulumi.Input

	ToConfigTerminateDbInstanceMapOutput() ConfigTerminateDbInstanceMapOutput
	ToConfigTerminateDbInstanceMapOutputWithContext(context.Context) ConfigTerminateDbInstanceMapOutput
}

type ConfigTerminateDbInstanceMap map[string]ConfigTerminateDbInstanceInput

func (ConfigTerminateDbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigTerminateDbInstance)(nil)).Elem()
}

func (i ConfigTerminateDbInstanceMap) ToConfigTerminateDbInstanceMapOutput() ConfigTerminateDbInstanceMapOutput {
	return i.ToConfigTerminateDbInstanceMapOutputWithContext(context.Background())
}

func (i ConfigTerminateDbInstanceMap) ToConfigTerminateDbInstanceMapOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigTerminateDbInstanceMapOutput)
}

type ConfigTerminateDbInstanceOutput struct{ *pulumi.OutputState }

func (ConfigTerminateDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigTerminateDbInstance)(nil)).Elem()
}

func (o ConfigTerminateDbInstanceOutput) ToConfigTerminateDbInstanceOutput() ConfigTerminateDbInstanceOutput {
	return o
}

func (o ConfigTerminateDbInstanceOutput) ToConfigTerminateDbInstanceOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceOutput {
	return o
}

// Instance ID.
func (o ConfigTerminateDbInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigTerminateDbInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type ConfigTerminateDbInstanceArrayOutput struct{ *pulumi.OutputState }

func (ConfigTerminateDbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigTerminateDbInstance)(nil)).Elem()
}

func (o ConfigTerminateDbInstanceArrayOutput) ToConfigTerminateDbInstanceArrayOutput() ConfigTerminateDbInstanceArrayOutput {
	return o
}

func (o ConfigTerminateDbInstanceArrayOutput) ToConfigTerminateDbInstanceArrayOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceArrayOutput {
	return o
}

func (o ConfigTerminateDbInstanceArrayOutput) Index(i pulumi.IntInput) ConfigTerminateDbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigTerminateDbInstance {
		return vs[0].([]*ConfigTerminateDbInstance)[vs[1].(int)]
	}).(ConfigTerminateDbInstanceOutput)
}

type ConfigTerminateDbInstanceMapOutput struct{ *pulumi.OutputState }

func (ConfigTerminateDbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigTerminateDbInstance)(nil)).Elem()
}

func (o ConfigTerminateDbInstanceMapOutput) ToConfigTerminateDbInstanceMapOutput() ConfigTerminateDbInstanceMapOutput {
	return o
}

func (o ConfigTerminateDbInstanceMapOutput) ToConfigTerminateDbInstanceMapOutputWithContext(ctx context.Context) ConfigTerminateDbInstanceMapOutput {
	return o
}

func (o ConfigTerminateDbInstanceMapOutput) MapIndex(k pulumi.StringInput) ConfigTerminateDbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigTerminateDbInstance {
		return vs[0].(map[string]*ConfigTerminateDbInstance)[vs[1].(string)]
	}).(ConfigTerminateDbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigTerminateDbInstanceInput)(nil)).Elem(), &ConfigTerminateDbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigTerminateDbInstanceArrayInput)(nil)).Elem(), ConfigTerminateDbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigTerminateDbInstanceMapInput)(nil)).Elem(), ConfigTerminateDbInstanceMap{})
	pulumi.RegisterOutputType(ConfigTerminateDbInstanceOutput{})
	pulumi.RegisterOutputType(ConfigTerminateDbInstanceArrayOutput{})
	pulumi.RegisterOutputType(ConfigTerminateDbInstanceMapOutput{})
}
