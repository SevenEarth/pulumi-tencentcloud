// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a sqlserver configInstanceParam
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
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
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewConfigInstanceParam(ctx, "exampleConfigInstanceParam", &Sqlserver.ConfigInstanceParamArgs{
//				InstanceId: exampleBasicInstance.ID(),
//				ParamLists: sqlserver.ConfigInstanceParamParamListArray{
//					&sqlserver.ConfigInstanceParamParamListArgs{
//						Name:         pulumi.String("fill factor(%)"),
//						CurrentValue: pulumi.String("90"),
//					},
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// sqlserver config_instance_param can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Sqlserver/configInstanceParam:ConfigInstanceParam example config_instance_param
// ```
type ConfigInstanceParam struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
	ParamLists ConfigInstanceParamParamListArrayOutput `pulumi:"paramLists"`
}

// NewConfigInstanceParam registers a new resource with the given unique name, arguments, and options.
func NewConfigInstanceParam(ctx *pulumi.Context,
	name string, args *ConfigInstanceParamArgs, opts ...pulumi.ResourceOption) (*ConfigInstanceParam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ParamLists == nil {
		return nil, errors.New("invalid value for required argument 'ParamLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigInstanceParam
	err := ctx.RegisterResource("tencentcloud:Sqlserver/configInstanceParam:ConfigInstanceParam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigInstanceParam gets an existing ConfigInstanceParam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigInstanceParam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigInstanceParamState, opts ...pulumi.ResourceOption) (*ConfigInstanceParam, error) {
	var resource ConfigInstanceParam
	err := ctx.ReadResource("tencentcloud:Sqlserver/configInstanceParam:ConfigInstanceParam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigInstanceParam resources.
type configInstanceParamState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
	ParamLists []ConfigInstanceParamParamList `pulumi:"paramLists"`
}

type ConfigInstanceParamState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
	ParamLists ConfigInstanceParamParamListArrayInput
}

func (ConfigInstanceParamState) ElementType() reflect.Type {
	return reflect.TypeOf((*configInstanceParamState)(nil)).Elem()
}

type configInstanceParamArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
	ParamLists []ConfigInstanceParamParamList `pulumi:"paramLists"`
}

// The set of arguments for constructing a ConfigInstanceParam resource.
type ConfigInstanceParamArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
	// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
	ParamLists ConfigInstanceParamParamListArrayInput
}

func (ConfigInstanceParamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configInstanceParamArgs)(nil)).Elem()
}

type ConfigInstanceParamInput interface {
	pulumi.Input

	ToConfigInstanceParamOutput() ConfigInstanceParamOutput
	ToConfigInstanceParamOutputWithContext(ctx context.Context) ConfigInstanceParamOutput
}

func (*ConfigInstanceParam) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigInstanceParam)(nil)).Elem()
}

func (i *ConfigInstanceParam) ToConfigInstanceParamOutput() ConfigInstanceParamOutput {
	return i.ToConfigInstanceParamOutputWithContext(context.Background())
}

func (i *ConfigInstanceParam) ToConfigInstanceParamOutputWithContext(ctx context.Context) ConfigInstanceParamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigInstanceParamOutput)
}

// ConfigInstanceParamArrayInput is an input type that accepts ConfigInstanceParamArray and ConfigInstanceParamArrayOutput values.
// You can construct a concrete instance of `ConfigInstanceParamArrayInput` via:
//
//	ConfigInstanceParamArray{ ConfigInstanceParamArgs{...} }
type ConfigInstanceParamArrayInput interface {
	pulumi.Input

	ToConfigInstanceParamArrayOutput() ConfigInstanceParamArrayOutput
	ToConfigInstanceParamArrayOutputWithContext(context.Context) ConfigInstanceParamArrayOutput
}

type ConfigInstanceParamArray []ConfigInstanceParamInput

func (ConfigInstanceParamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigInstanceParam)(nil)).Elem()
}

func (i ConfigInstanceParamArray) ToConfigInstanceParamArrayOutput() ConfigInstanceParamArrayOutput {
	return i.ToConfigInstanceParamArrayOutputWithContext(context.Background())
}

func (i ConfigInstanceParamArray) ToConfigInstanceParamArrayOutputWithContext(ctx context.Context) ConfigInstanceParamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigInstanceParamArrayOutput)
}

// ConfigInstanceParamMapInput is an input type that accepts ConfigInstanceParamMap and ConfigInstanceParamMapOutput values.
// You can construct a concrete instance of `ConfigInstanceParamMapInput` via:
//
//	ConfigInstanceParamMap{ "key": ConfigInstanceParamArgs{...} }
type ConfigInstanceParamMapInput interface {
	pulumi.Input

	ToConfigInstanceParamMapOutput() ConfigInstanceParamMapOutput
	ToConfigInstanceParamMapOutputWithContext(context.Context) ConfigInstanceParamMapOutput
}

type ConfigInstanceParamMap map[string]ConfigInstanceParamInput

func (ConfigInstanceParamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigInstanceParam)(nil)).Elem()
}

func (i ConfigInstanceParamMap) ToConfigInstanceParamMapOutput() ConfigInstanceParamMapOutput {
	return i.ToConfigInstanceParamMapOutputWithContext(context.Background())
}

func (i ConfigInstanceParamMap) ToConfigInstanceParamMapOutputWithContext(ctx context.Context) ConfigInstanceParamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigInstanceParamMapOutput)
}

type ConfigInstanceParamOutput struct{ *pulumi.OutputState }

func (ConfigInstanceParamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigInstanceParam)(nil)).Elem()
}

func (o ConfigInstanceParamOutput) ToConfigInstanceParamOutput() ConfigInstanceParamOutput {
	return o
}

func (o ConfigInstanceParamOutput) ToConfigInstanceParamOutputWithContext(ctx context.Context) ConfigInstanceParamOutput {
	return o
}

// Instance ID.
func (o ConfigInstanceParamOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigInstanceParam) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// List of modified parameters. Each list element has two fields: Name and CurrentValue. Set Name to the parameter name and CurrentValue to the new value after modification. Note: if the instance needs to be restarted for the modified parameter to take effect, it will be restarted immediately or during the maintenance time. Before you modify a parameter, you can use the DescribeInstanceParams API to query whether the instance needs to be restarted.
func (o ConfigInstanceParamOutput) ParamLists() ConfigInstanceParamParamListArrayOutput {
	return o.ApplyT(func(v *ConfigInstanceParam) ConfigInstanceParamParamListArrayOutput { return v.ParamLists }).(ConfigInstanceParamParamListArrayOutput)
}

type ConfigInstanceParamArrayOutput struct{ *pulumi.OutputState }

func (ConfigInstanceParamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigInstanceParam)(nil)).Elem()
}

func (o ConfigInstanceParamArrayOutput) ToConfigInstanceParamArrayOutput() ConfigInstanceParamArrayOutput {
	return o
}

func (o ConfigInstanceParamArrayOutput) ToConfigInstanceParamArrayOutputWithContext(ctx context.Context) ConfigInstanceParamArrayOutput {
	return o
}

func (o ConfigInstanceParamArrayOutput) Index(i pulumi.IntInput) ConfigInstanceParamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigInstanceParam {
		return vs[0].([]*ConfigInstanceParam)[vs[1].(int)]
	}).(ConfigInstanceParamOutput)
}

type ConfigInstanceParamMapOutput struct{ *pulumi.OutputState }

func (ConfigInstanceParamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigInstanceParam)(nil)).Elem()
}

func (o ConfigInstanceParamMapOutput) ToConfigInstanceParamMapOutput() ConfigInstanceParamMapOutput {
	return o
}

func (o ConfigInstanceParamMapOutput) ToConfigInstanceParamMapOutputWithContext(ctx context.Context) ConfigInstanceParamMapOutput {
	return o
}

func (o ConfigInstanceParamMapOutput) MapIndex(k pulumi.StringInput) ConfigInstanceParamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigInstanceParam {
		return vs[0].(map[string]*ConfigInstanceParam)[vs[1].(string)]
	}).(ConfigInstanceParamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInstanceParamInput)(nil)).Elem(), &ConfigInstanceParam{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInstanceParamArrayInput)(nil)).Elem(), ConfigInstanceParamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInstanceParamMapInput)(nil)).Elem(), ConfigInstanceParamMap{})
	pulumi.RegisterOutputType(ConfigInstanceParamOutput{})
	pulumi.RegisterOutputType(ConfigInstanceParamArrayOutput{})
	pulumi.RegisterOutputType(ConfigInstanceParamMapOutput{})
}
