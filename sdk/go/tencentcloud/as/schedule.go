// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for an AS (Auto scaling) schedule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
// 			Product: "as",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		image, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
// 			ImageTypes: []string{
// 				"PUBLIC_IMAGE",
// 			},
// 			OsName: pulumi.StringRef("TencentOS Server 3.2 (Final)"),
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
// 			VpcId:            vpc.ID(),
// 			CidrBlock:        pulumi.String("10.0.0.0/16"),
// 			AvailabilityZone: pulumi.String(zones.Zones[0].Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleScalingConfig, err := As.NewScalingConfig(ctx, "exampleScalingConfig", &As.ScalingConfigArgs{
// 			ConfigurationName: pulumi.String("tf-example"),
// 			ImageId:           pulumi.String(image.Images[0].ImageId),
// 			InstanceTypes: pulumi.StringArray{
// 				pulumi.String("SA1.SMALL1"),
// 				pulumi.String("SA2.SMALL1"),
// 				pulumi.String("SA2.SMALL2"),
// 				pulumi.String("SA2.SMALL4"),
// 			},
// 			InstanceNameSettings: &as.ScalingConfigInstanceNameSettingsArgs{
// 				InstanceName: pulumi.String("test-ins-name"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleScalingGroup, err := As.NewScalingGroup(ctx, "exampleScalingGroup", &As.ScalingGroupArgs{
// 			ScalingGroupName: pulumi.String("tf-example"),
// 			ConfigurationId:  exampleScalingConfig.ID(),
// 			MaxSize:          pulumi.Int(1),
// 			MinSize:          pulumi.Int(0),
// 			VpcId:            vpc.ID(),
// 			SubnetIds: pulumi.StringArray{
// 				subnet.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = As.NewSchedule(ctx, "exampleSchedule", &As.ScheduleArgs{
// 			ScalingGroupId:     exampleScalingGroup.ID(),
// 			ScheduleActionName: pulumi.String("tf-as-schedule"),
// 			MaxSize:            pulumi.Int(10),
// 			MinSize:            pulumi.Int(0),
// 			DesiredCapacity:    pulumi.Int(0),
// 			StartTime:          pulumi.String("2019-01-01T00:00:00+08:00"),
// 			EndTime:            pulumi.String("2019-12-01T00:00:00+08:00"),
// 			Recurrence:         pulumi.String("0 0 * * *"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Schedule struct {
	pulumi.CustomResourceState

	// The desired number of CVM instances that should be running in the group.
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group.
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
	Recurrence pulumi.StringPtrOutput `pulumi:"recurrence"`
	// ID of a scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// The name of this scaling action.
	ScheduleActionName pulumi.StringOutput `pulumi:"scheduleActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DesiredCapacity == nil {
		return nil, errors.New("invalid value for required argument 'DesiredCapacity'")
	}
	if args.MaxSize == nil {
		return nil, errors.New("invalid value for required argument 'MaxSize'")
	}
	if args.MinSize == nil {
		return nil, errors.New("invalid value for required argument 'MinSize'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	if args.ScheduleActionName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleActionName'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Schedule
	err := ctx.RegisterResource("tencentcloud:As/schedule:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("tencentcloud:As/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	// The desired number of CVM instances that should be running in the group.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	EndTime *string `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group.
	MinSize *int `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
	Recurrence *string `pulumi:"recurrence"`
	// ID of a scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The name of this scaling action.
	ScheduleActionName *string `pulumi:"scheduleActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	StartTime *string `pulumi:"startTime"`
}

type ScheduleState struct {
	// The desired number of CVM instances that should be running in the group.
	DesiredCapacity pulumi.IntPtrInput
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	EndTime pulumi.StringPtrInput
	// The maximum size for the Auto Scaling group.
	MaxSize pulumi.IntPtrInput
	// The minimum size for the Auto Scaling group.
	MinSize pulumi.IntPtrInput
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
	Recurrence pulumi.StringPtrInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringPtrInput
	// The name of this scaling action.
	ScheduleActionName pulumi.StringPtrInput
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	StartTime pulumi.StringPtrInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// The desired number of CVM instances that should be running in the group.
	DesiredCapacity int `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	EndTime *string `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group.
	MaxSize int `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group.
	MinSize int `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
	Recurrence *string `pulumi:"recurrence"`
	// ID of a scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The name of this scaling action.
	ScheduleActionName string `pulumi:"scheduleActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	StartTime string `pulumi:"startTime"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The desired number of CVM instances that should be running in the group.
	DesiredCapacity pulumi.IntInput
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	EndTime pulumi.StringPtrInput
	// The maximum size for the Auto Scaling group.
	MaxSize pulumi.IntInput
	// The minimum size for the Auto Scaling group.
	MinSize pulumi.IntInput
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
	Recurrence pulumi.StringPtrInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringInput
	// The name of this scaling action.
	ScheduleActionName pulumi.StringInput
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
	StartTime pulumi.StringInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (*Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

// ScheduleArrayInput is an input type that accepts ScheduleArray and ScheduleArrayOutput values.
// You can construct a concrete instance of `ScheduleArrayInput` via:
//
//          ScheduleArray{ ScheduleArgs{...} }
type ScheduleArrayInput interface {
	pulumi.Input

	ToScheduleArrayOutput() ScheduleArrayOutput
	ToScheduleArrayOutputWithContext(context.Context) ScheduleArrayOutput
}

type ScheduleArray []ScheduleInput

func (ScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schedule)(nil)).Elem()
}

func (i ScheduleArray) ToScheduleArrayOutput() ScheduleArrayOutput {
	return i.ToScheduleArrayOutputWithContext(context.Background())
}

func (i ScheduleArray) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleArrayOutput)
}

// ScheduleMapInput is an input type that accepts ScheduleMap and ScheduleMapOutput values.
// You can construct a concrete instance of `ScheduleMapInput` via:
//
//          ScheduleMap{ "key": ScheduleArgs{...} }
type ScheduleMapInput interface {
	pulumi.Input

	ToScheduleMapOutput() ScheduleMapOutput
	ToScheduleMapOutputWithContext(context.Context) ScheduleMapOutput
}

type ScheduleMap map[string]ScheduleInput

func (ScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schedule)(nil)).Elem()
}

func (i ScheduleMap) ToScheduleMapOutput() ScheduleMapOutput {
	return i.ToScheduleMapOutputWithContext(context.Background())
}

func (i ScheduleMap) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMapOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

// The desired number of CVM instances that should be running in the group.
func (o ScheduleOutput) DesiredCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.DesiredCapacity }).(pulumi.IntOutput)
}

// The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
func (o ScheduleOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The maximum size for the Auto Scaling group.
func (o ScheduleOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// The minimum size for the Auto Scaling group.
func (o ScheduleOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.MinSize }).(pulumi.IntOutput)
}

// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with endTime together.
func (o ScheduleOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Recurrence }).(pulumi.StringPtrOutput)
}

// ID of a scaling group.
func (o ScheduleOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The name of this scaling action.
func (o ScheduleOutput) ScheduleActionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ScheduleActionName }).(pulumi.StringOutput)
}

// The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
func (o ScheduleOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

type ScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schedule)(nil)).Elem()
}

func (o ScheduleArrayOutput) ToScheduleArrayOutput() ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) Index(i pulumi.IntInput) ScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Schedule {
		return vs[0].([]*Schedule)[vs[1].(int)]
	}).(ScheduleOutput)
}

type ScheduleMapOutput struct{ *pulumi.OutputState }

func (ScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schedule)(nil)).Elem()
}

func (o ScheduleMapOutput) ToScheduleMapOutput() ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) MapIndex(k pulumi.StringInput) ScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Schedule {
		return vs[0].(map[string]*Schedule)[vs[1].(string)]
	}).(ScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleInput)(nil)).Elem(), &Schedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleArrayInput)(nil)).Elem(), ScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleMapInput)(nil)).Elem(), ScheduleMap{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(ScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScheduleMapOutput{})
}
