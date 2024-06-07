// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a group of AS (Auto scaling) instances.
//
// ## Example Usage
//
// ### Create a basic Scaling Group
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "as",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			image, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
//				ImageTypes: []string{
//					"PUBLIC_IMAGE",
//				},
//				OsName: pulumi.StringRef("TencentOS Server 3.2 (Final)"),
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
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//			})
//			if err != nil {
//				return err
//			}
//			exampleScalingConfig, err := As.NewScalingConfig(ctx, "exampleScalingConfig", &As.ScalingConfigArgs{
//				ConfigurationName: pulumi.String("tf-example"),
//				ImageId:           pulumi.String(image.Images[0].ImageId),
//				InstanceTypes: pulumi.StringArray{
//					pulumi.String("SA1.SMALL1"),
//					pulumi.String("SA2.SMALL1"),
//					pulumi.String("SA2.SMALL2"),
//					pulumi.String("SA2.SMALL4"),
//				},
//				InstanceNameSettings: &as.ScalingConfigInstanceNameSettingsArgs{
//					InstanceName: pulumi.String("test-ins-name"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = As.NewScalingGroup(ctx, "exampleScalingGroup", &As.ScalingGroupArgs{
//				ScalingGroupName: pulumi.String("tf-example"),
//				ConfigurationId:  exampleScalingConfig.ID(),
//				MaxSize:          pulumi.Int(1),
//				MinSize:          pulumi.Int(0),
//				VpcId:            vpc.ID(),
//				SubnetIds: pulumi.StringArray{
//					subnet.ID(),
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
// AutoScaling Groups can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:As/scalingGroup:ScalingGroup scaling_group asg-n32ymck2
// ```
type ScalingGroup struct {
	pulumi.CustomResourceState

	// An available ID for a launch configuration.
	ConfigurationId pulumi.StringOutput `pulumi:"configurationId"`
	// The time when the AS group was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown pulumi.IntPtrOutput `pulumi:"defaultCooldown"`
	// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// List of application load balancers, which can't be specified with `loadBalancerIds` together.
	ForwardBalancerIds ScalingGroupForwardBalancerIdArrayOutput `pulumi:"forwardBalancerIds"`
	// Instance number of a scaling group.
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// ID list of traditional load balancers.
	LoadBalancerIds pulumi.StringArrayOutput `pulumi:"loadBalancerIds"`
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy pulumi.StringPtrOutput `pulumi:"multiZoneSubnetPolicy"`
	// Specifies to which project the scaling group belongs.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy pulumi.BoolPtrOutput `pulumi:"replaceLoadBalancerUnhealthy"`
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy pulumi.BoolPtrOutput `pulumi:"replaceMonitorUnhealthy"`
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy pulumi.StringPtrOutput `pulumi:"retryPolicy"`
	// Name of a scaling group.
	ScalingGroupName pulumi.StringOutput `pulumi:"scalingGroupName"`
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode pulumi.StringPtrOutput `pulumi:"scalingMode"`
	// Current status of a scaling group.
	Status pulumi.StringOutput `pulumi:"status"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Tags of a scaling group.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies pulumi.StringOutput `pulumi:"terminationPolicies"`
	// ID of VPC network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// List of available zones, for Basic network it is required.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewScalingGroup registers a new resource with the given unique name, arguments, and options.
func NewScalingGroup(ctx *pulumi.Context,
	name string, args *ScalingGroupArgs, opts ...pulumi.ResourceOption) (*ScalingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationId'")
	}
	if args.MaxSize == nil {
		return nil, errors.New("invalid value for required argument 'MaxSize'")
	}
	if args.MinSize == nil {
		return nil, errors.New("invalid value for required argument 'MinSize'")
	}
	if args.ScalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScalingGroup
	err := ctx.RegisterResource("tencentcloud:As/scalingGroup:ScalingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingGroup gets an existing ScalingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingGroupState, opts ...pulumi.ResourceOption) (*ScalingGroup, error) {
	var resource ScalingGroup
	err := ctx.ReadResource("tencentcloud:As/scalingGroup:ScalingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingGroup resources.
type scalingGroupState struct {
	// An available ID for a launch configuration.
	ConfigurationId *string `pulumi:"configurationId"`
	// The time when the AS group was created.
	CreateTime *string `pulumi:"createTime"`
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// List of application load balancers, which can't be specified with `loadBalancerIds` together.
	ForwardBalancerIds []ScalingGroupForwardBalancerId `pulumi:"forwardBalancerIds"`
	// Instance number of a scaling group.
	InstanceCount *int `pulumi:"instanceCount"`
	// ID list of traditional load balancers.
	LoadBalancerIds []string `pulumi:"loadBalancerIds"`
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize *int `pulumi:"maxSize"`
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize *int `pulumi:"minSize"`
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy *string `pulumi:"multiZoneSubnetPolicy"`
	// Specifies to which project the scaling group belongs.
	ProjectId *int `pulumi:"projectId"`
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy *bool `pulumi:"replaceLoadBalancerUnhealthy"`
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy *bool `pulumi:"replaceMonitorUnhealthy"`
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy *string `pulumi:"retryPolicy"`
	// Name of a scaling group.
	ScalingGroupName *string `pulumi:"scalingGroupName"`
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode *string `pulumi:"scalingMode"`
	// Current status of a scaling group.
	Status *string `pulumi:"status"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tags of a scaling group.
	Tags map[string]interface{} `pulumi:"tags"`
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies *string `pulumi:"terminationPolicies"`
	// ID of VPC network.
	VpcId *string `pulumi:"vpcId"`
	// List of available zones, for Basic network it is required.
	Zones []string `pulumi:"zones"`
}

type ScalingGroupState struct {
	// An available ID for a launch configuration.
	ConfigurationId pulumi.StringPtrInput
	// The time when the AS group was created.
	CreateTime pulumi.StringPtrInput
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown pulumi.IntPtrInput
	// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
	DesiredCapacity pulumi.IntPtrInput
	// List of application load balancers, which can't be specified with `loadBalancerIds` together.
	ForwardBalancerIds ScalingGroupForwardBalancerIdArrayInput
	// Instance number of a scaling group.
	InstanceCount pulumi.IntPtrInput
	// ID list of traditional load balancers.
	LoadBalancerIds pulumi.StringArrayInput
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize pulumi.IntPtrInput
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize pulumi.IntPtrInput
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy pulumi.StringPtrInput
	// Specifies to which project the scaling group belongs.
	ProjectId pulumi.IntPtrInput
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy pulumi.BoolPtrInput
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy pulumi.BoolPtrInput
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy pulumi.StringPtrInput
	// Name of a scaling group.
	ScalingGroupName pulumi.StringPtrInput
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode pulumi.StringPtrInput
	// Current status of a scaling group.
	Status pulumi.StringPtrInput
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayInput
	// Tags of a scaling group.
	Tags pulumi.MapInput
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies pulumi.StringPtrInput
	// ID of VPC network.
	VpcId pulumi.StringPtrInput
	// List of available zones, for Basic network it is required.
	Zones pulumi.StringArrayInput
}

func (ScalingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupState)(nil)).Elem()
}

type scalingGroupArgs struct {
	// An available ID for a launch configuration.
	ConfigurationId string `pulumi:"configurationId"`
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// List of application load balancers, which can't be specified with `loadBalancerIds` together.
	ForwardBalancerIds []ScalingGroupForwardBalancerId `pulumi:"forwardBalancerIds"`
	// ID list of traditional load balancers.
	LoadBalancerIds []string `pulumi:"loadBalancerIds"`
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize int `pulumi:"maxSize"`
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize int `pulumi:"minSize"`
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy *string `pulumi:"multiZoneSubnetPolicy"`
	// Specifies to which project the scaling group belongs.
	ProjectId *int `pulumi:"projectId"`
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy *bool `pulumi:"replaceLoadBalancerUnhealthy"`
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy *bool `pulumi:"replaceMonitorUnhealthy"`
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy *string `pulumi:"retryPolicy"`
	// Name of a scaling group.
	ScalingGroupName string `pulumi:"scalingGroupName"`
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode *string `pulumi:"scalingMode"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tags of a scaling group.
	Tags map[string]interface{} `pulumi:"tags"`
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies *string `pulumi:"terminationPolicies"`
	// ID of VPC network.
	VpcId string `pulumi:"vpcId"`
	// List of available zones, for Basic network it is required.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ScalingGroup resource.
type ScalingGroupArgs struct {
	// An available ID for a launch configuration.
	ConfigurationId pulumi.StringInput
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown pulumi.IntPtrInput
	// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
	DesiredCapacity pulumi.IntPtrInput
	// List of application load balancers, which can't be specified with `loadBalancerIds` together.
	ForwardBalancerIds ScalingGroupForwardBalancerIdArrayInput
	// ID list of traditional load balancers.
	LoadBalancerIds pulumi.StringArrayInput
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize pulumi.IntInput
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize pulumi.IntInput
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy pulumi.StringPtrInput
	// Specifies to which project the scaling group belongs.
	ProjectId pulumi.IntPtrInput
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy pulumi.BoolPtrInput
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy pulumi.BoolPtrInput
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy pulumi.StringPtrInput
	// Name of a scaling group.
	ScalingGroupName pulumi.StringInput
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode pulumi.StringPtrInput
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayInput
	// Tags of a scaling group.
	Tags pulumi.MapInput
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies pulumi.StringPtrInput
	// ID of VPC network.
	VpcId pulumi.StringInput
	// List of available zones, for Basic network it is required.
	Zones pulumi.StringArrayInput
}

func (ScalingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupArgs)(nil)).Elem()
}

type ScalingGroupInput interface {
	pulumi.Input

	ToScalingGroupOutput() ScalingGroupOutput
	ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput
}

func (*ScalingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroup)(nil)).Elem()
}

func (i *ScalingGroup) ToScalingGroupOutput() ScalingGroupOutput {
	return i.ToScalingGroupOutputWithContext(context.Background())
}

func (i *ScalingGroup) ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupOutput)
}

// ScalingGroupArrayInput is an input type that accepts ScalingGroupArray and ScalingGroupArrayOutput values.
// You can construct a concrete instance of `ScalingGroupArrayInput` via:
//
//	ScalingGroupArray{ ScalingGroupArgs{...} }
type ScalingGroupArrayInput interface {
	pulumi.Input

	ToScalingGroupArrayOutput() ScalingGroupArrayOutput
	ToScalingGroupArrayOutputWithContext(context.Context) ScalingGroupArrayOutput
}

type ScalingGroupArray []ScalingGroupInput

func (ScalingGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroup)(nil)).Elem()
}

func (i ScalingGroupArray) ToScalingGroupArrayOutput() ScalingGroupArrayOutput {
	return i.ToScalingGroupArrayOutputWithContext(context.Background())
}

func (i ScalingGroupArray) ToScalingGroupArrayOutputWithContext(ctx context.Context) ScalingGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupArrayOutput)
}

// ScalingGroupMapInput is an input type that accepts ScalingGroupMap and ScalingGroupMapOutput values.
// You can construct a concrete instance of `ScalingGroupMapInput` via:
//
//	ScalingGroupMap{ "key": ScalingGroupArgs{...} }
type ScalingGroupMapInput interface {
	pulumi.Input

	ToScalingGroupMapOutput() ScalingGroupMapOutput
	ToScalingGroupMapOutputWithContext(context.Context) ScalingGroupMapOutput
}

type ScalingGroupMap map[string]ScalingGroupInput

func (ScalingGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroup)(nil)).Elem()
}

func (i ScalingGroupMap) ToScalingGroupMapOutput() ScalingGroupMapOutput {
	return i.ToScalingGroupMapOutputWithContext(context.Background())
}

func (i ScalingGroupMap) ToScalingGroupMapOutputWithContext(ctx context.Context) ScalingGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupMapOutput)
}

type ScalingGroupOutput struct{ *pulumi.OutputState }

func (ScalingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupOutput) ToScalingGroupOutput() ScalingGroupOutput {
	return o
}

func (o ScalingGroupOutput) ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput {
	return o
}

// An available ID for a launch configuration.
func (o ScalingGroupOutput) ConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.ConfigurationId }).(pulumi.StringOutput)
}

// The time when the AS group was created.
func (o ScalingGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Default cooldown time in second, and default value is `300`.
func (o ScalingGroupOutput) DefaultCooldown() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntPtrOutput { return v.DefaultCooldown }).(pulumi.IntPtrOutput)
}

// Desired volume of CVM instances, which is between `maxSize` and `minSize`.
func (o ScalingGroupOutput) DesiredCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.DesiredCapacity }).(pulumi.IntOutput)
}

// List of application load balancers, which can't be specified with `loadBalancerIds` together.
func (o ScalingGroupOutput) ForwardBalancerIds() ScalingGroupForwardBalancerIdArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) ScalingGroupForwardBalancerIdArrayOutput { return v.ForwardBalancerIds }).(ScalingGroupForwardBalancerIdArrayOutput)
}

// Instance number of a scaling group.
func (o ScalingGroupOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.InstanceCount }).(pulumi.IntOutput)
}

// ID list of traditional load balancers.
func (o ScalingGroupOutput) LoadBalancerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringArrayOutput { return v.LoadBalancerIds }).(pulumi.StringArrayOutput)
}

// Maximum number of CVM instances. Valid value ranges: (0~2000).
func (o ScalingGroupOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// Minimum number of CVM instances. Valid value ranges: (0~2000).
func (o ScalingGroupOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.MinSize }).(pulumi.IntOutput)
}

// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
func (o ScalingGroupOutput) MultiZoneSubnetPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringPtrOutput { return v.MultiZoneSubnetPolicy }).(pulumi.StringPtrOutput)
}

// Specifies to which project the scaling group belongs.
func (o ScalingGroupOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
func (o ScalingGroupOutput) ReplaceLoadBalancerUnhealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.BoolPtrOutput { return v.ReplaceLoadBalancerUnhealthy }).(pulumi.BoolPtrOutput)
}

// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
func (o ScalingGroupOutput) ReplaceMonitorUnhealthy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.BoolPtrOutput { return v.ReplaceMonitorUnhealthy }).(pulumi.BoolPtrOutput)
}

// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
func (o ScalingGroupOutput) RetryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringPtrOutput { return v.RetryPolicy }).(pulumi.StringPtrOutput)
}

// Name of a scaling group.
func (o ScalingGroupOutput) ScalingGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.ScalingGroupName }).(pulumi.StringOutput)
}

// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
func (o ScalingGroupOutput) ScalingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringPtrOutput { return v.ScalingMode }).(pulumi.StringPtrOutput)
}

// Current status of a scaling group.
func (o ScalingGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// ID list of subnet, and for VPC it is required.
func (o ScalingGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Tags of a scaling group.
func (o ScalingGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
func (o ScalingGroupOutput) TerminationPolicies() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.TerminationPolicies }).(pulumi.StringOutput)
}

// ID of VPC network.
func (o ScalingGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// List of available zones, for Basic network it is required.
func (o ScalingGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type ScalingGroupArrayOutput struct{ *pulumi.OutputState }

func (ScalingGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupArrayOutput) ToScalingGroupArrayOutput() ScalingGroupArrayOutput {
	return o
}

func (o ScalingGroupArrayOutput) ToScalingGroupArrayOutputWithContext(ctx context.Context) ScalingGroupArrayOutput {
	return o
}

func (o ScalingGroupArrayOutput) Index(i pulumi.IntInput) ScalingGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingGroup {
		return vs[0].([]*ScalingGroup)[vs[1].(int)]
	}).(ScalingGroupOutput)
}

type ScalingGroupMapOutput struct{ *pulumi.OutputState }

func (ScalingGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupMapOutput) ToScalingGroupMapOutput() ScalingGroupMapOutput {
	return o
}

func (o ScalingGroupMapOutput) ToScalingGroupMapOutputWithContext(ctx context.Context) ScalingGroupMapOutput {
	return o
}

func (o ScalingGroupMapOutput) MapIndex(k pulumi.StringInput) ScalingGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingGroup {
		return vs[0].(map[string]*ScalingGroup)[vs[1].(string)]
	}).(ScalingGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupInput)(nil)).Elem(), &ScalingGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupArrayInput)(nil)).Elem(), ScalingGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupMapInput)(nil)).Elem(), ScalingGroupMap{})
	pulumi.RegisterOutputType(ScalingGroupOutput{})
	pulumi.RegisterOutputType(ScalingGroupArrayOutput{})
	pulumi.RegisterOutputType(ScalingGroupMapOutput{})
}
