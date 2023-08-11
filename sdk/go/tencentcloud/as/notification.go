// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for an AS (Auto scaling) notification.
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
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
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
// 		exampleGroup, err := Cam.NewGroup(ctx, "exampleGroup", &Cam.GroupArgs{
// 			Remark: pulumi.String("desc."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = As.NewNotification(ctx, "asNotification", &As.NotificationArgs{
// 			ScalingGroupId: exampleScalingGroup.ID(),
// 			NotificationTypes: pulumi.StringArray{
// 				pulumi.String("SCALE_OUT_SUCCESSFUL"),
// 				pulumi.String("SCALE_OUT_FAILED"),
// 				pulumi.String("SCALE_IN_FAILED"),
// 				pulumi.String("REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL"),
// 				pulumi.String("REPLACE_UNHEALTHY_INSTANCE_FAILED"),
// 			},
// 			NotificationUserGroupIds: pulumi.StringArray{
// 				exampleGroup.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Notification struct {
	pulumi.CustomResourceState

	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes pulumi.StringArrayOutput `pulumi:"notificationTypes"`
	// A group of user IDs to be notified.
	NotificationUserGroupIds pulumi.StringArrayOutput `pulumi:"notificationUserGroupIds"`
	// ID of a scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NotificationTypes == nil {
		return nil, errors.New("invalid value for required argument 'NotificationTypes'")
	}
	if args.NotificationUserGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'NotificationUserGroupIds'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Notification
	err := ctx.RegisterResource("tencentcloud:As/notification:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("tencentcloud:As/notification:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes []string `pulumi:"notificationTypes"`
	// A group of user IDs to be notified.
	NotificationUserGroupIds []string `pulumi:"notificationUserGroupIds"`
	// ID of a scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

type NotificationState struct {
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes pulumi.StringArrayInput
	// A group of user IDs to be notified.
	NotificationUserGroupIds pulumi.StringArrayInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringPtrInput
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes []string `pulumi:"notificationTypes"`
	// A group of user IDs to be notified.
	NotificationUserGroupIds []string `pulumi:"notificationUserGroupIds"`
	// ID of a scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
	NotificationTypes pulumi.StringArrayInput
	// A group of user IDs to be notified.
	NotificationUserGroupIds pulumi.StringArrayInput
	// ID of a scaling group.
	ScalingGroupId pulumi.StringInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}

type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(ctx context.Context) NotificationOutput
}

func (*Notification) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (i *Notification) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i *Notification) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}

// NotificationArrayInput is an input type that accepts NotificationArray and NotificationArrayOutput values.
// You can construct a concrete instance of `NotificationArrayInput` via:
//
//          NotificationArray{ NotificationArgs{...} }
type NotificationArrayInput interface {
	pulumi.Input

	ToNotificationArrayOutput() NotificationArrayOutput
	ToNotificationArrayOutputWithContext(context.Context) NotificationArrayOutput
}

type NotificationArray []NotificationInput

func (NotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (i NotificationArray) ToNotificationArrayOutput() NotificationArrayOutput {
	return i.ToNotificationArrayOutputWithContext(context.Background())
}

func (i NotificationArray) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationArrayOutput)
}

// NotificationMapInput is an input type that accepts NotificationMap and NotificationMapOutput values.
// You can construct a concrete instance of `NotificationMapInput` via:
//
//          NotificationMap{ "key": NotificationArgs{...} }
type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

// A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
func (o NotificationOutput) NotificationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.NotificationTypes }).(pulumi.StringArrayOutput)
}

// A group of user IDs to be notified.
func (o NotificationOutput) NotificationUserGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.NotificationUserGroupIds }).(pulumi.StringArrayOutput)
}

// ID of a scaling group.
func (o NotificationOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

type NotificationArrayOutput struct{ *pulumi.OutputState }

func (NotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (o NotificationArrayOutput) ToNotificationArrayOutput() NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) Index(i pulumi.IntInput) NotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].([]*Notification)[vs[1].(int)]
	}).(NotificationOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].(map[string]*Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationInput)(nil)).Elem(), &Notification{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationArrayInput)(nil)).Elem(), NotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationMapInput)(nil)).Elem(), NotificationMap{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationArrayOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
}
