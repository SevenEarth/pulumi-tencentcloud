// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to set as scalingGroup status
//
// ## Example Usage
// ### Deactivate Scaling Group
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
// 		_, err = As.NewScalingGroupStatus(ctx, "scalingGroupStatus", &As.ScalingGroupStatusArgs{
// 			AutoScalingGroupId: exampleScalingGroup.ID(),
// 			Enable:             pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enable Scaling Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := As.NewScalingGroupStatus(ctx, "scalingGroupStatus", &As.ScalingGroupStatusArgs{
// 			AutoScalingGroupId: pulumi.Any(tencentcloud_as_scaling_group.Example.Id),
// 			Enable:             pulumi.Bool(true),
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
// as scaling_group_status can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:As/scalingGroupStatus:ScalingGroupStatus scaling_group_status scaling_group_id
// ```
type ScalingGroupStatus struct {
	pulumi.CustomResourceState

	// Scaling group ID.
	AutoScalingGroupId pulumi.StringOutput `pulumi:"autoScalingGroupId"`
	// If enable auto scaling group.
	Enable pulumi.BoolOutput `pulumi:"enable"`
}

// NewScalingGroupStatus registers a new resource with the given unique name, arguments, and options.
func NewScalingGroupStatus(ctx *pulumi.Context,
	name string, args *ScalingGroupStatusArgs, opts ...pulumi.ResourceOption) (*ScalingGroupStatus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingGroupId'")
	}
	if args.Enable == nil {
		return nil, errors.New("invalid value for required argument 'Enable'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScalingGroupStatus
	err := ctx.RegisterResource("tencentcloud:As/scalingGroupStatus:ScalingGroupStatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingGroupStatus gets an existing ScalingGroupStatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingGroupStatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingGroupStatusState, opts ...pulumi.ResourceOption) (*ScalingGroupStatus, error) {
	var resource ScalingGroupStatus
	err := ctx.ReadResource("tencentcloud:As/scalingGroupStatus:ScalingGroupStatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingGroupStatus resources.
type scalingGroupStatusState struct {
	// Scaling group ID.
	AutoScalingGroupId *string `pulumi:"autoScalingGroupId"`
	// If enable auto scaling group.
	Enable *bool `pulumi:"enable"`
}

type ScalingGroupStatusState struct {
	// Scaling group ID.
	AutoScalingGroupId pulumi.StringPtrInput
	// If enable auto scaling group.
	Enable pulumi.BoolPtrInput
}

func (ScalingGroupStatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupStatusState)(nil)).Elem()
}

type scalingGroupStatusArgs struct {
	// Scaling group ID.
	AutoScalingGroupId string `pulumi:"autoScalingGroupId"`
	// If enable auto scaling group.
	Enable bool `pulumi:"enable"`
}

// The set of arguments for constructing a ScalingGroupStatus resource.
type ScalingGroupStatusArgs struct {
	// Scaling group ID.
	AutoScalingGroupId pulumi.StringInput
	// If enable auto scaling group.
	Enable pulumi.BoolInput
}

func (ScalingGroupStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupStatusArgs)(nil)).Elem()
}

type ScalingGroupStatusInput interface {
	pulumi.Input

	ToScalingGroupStatusOutput() ScalingGroupStatusOutput
	ToScalingGroupStatusOutputWithContext(ctx context.Context) ScalingGroupStatusOutput
}

func (*ScalingGroupStatus) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroupStatus)(nil)).Elem()
}

func (i *ScalingGroupStatus) ToScalingGroupStatusOutput() ScalingGroupStatusOutput {
	return i.ToScalingGroupStatusOutputWithContext(context.Background())
}

func (i *ScalingGroupStatus) ToScalingGroupStatusOutputWithContext(ctx context.Context) ScalingGroupStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupStatusOutput)
}

// ScalingGroupStatusArrayInput is an input type that accepts ScalingGroupStatusArray and ScalingGroupStatusArrayOutput values.
// You can construct a concrete instance of `ScalingGroupStatusArrayInput` via:
//
//          ScalingGroupStatusArray{ ScalingGroupStatusArgs{...} }
type ScalingGroupStatusArrayInput interface {
	pulumi.Input

	ToScalingGroupStatusArrayOutput() ScalingGroupStatusArrayOutput
	ToScalingGroupStatusArrayOutputWithContext(context.Context) ScalingGroupStatusArrayOutput
}

type ScalingGroupStatusArray []ScalingGroupStatusInput

func (ScalingGroupStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroupStatus)(nil)).Elem()
}

func (i ScalingGroupStatusArray) ToScalingGroupStatusArrayOutput() ScalingGroupStatusArrayOutput {
	return i.ToScalingGroupStatusArrayOutputWithContext(context.Background())
}

func (i ScalingGroupStatusArray) ToScalingGroupStatusArrayOutputWithContext(ctx context.Context) ScalingGroupStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupStatusArrayOutput)
}

// ScalingGroupStatusMapInput is an input type that accepts ScalingGroupStatusMap and ScalingGroupStatusMapOutput values.
// You can construct a concrete instance of `ScalingGroupStatusMapInput` via:
//
//          ScalingGroupStatusMap{ "key": ScalingGroupStatusArgs{...} }
type ScalingGroupStatusMapInput interface {
	pulumi.Input

	ToScalingGroupStatusMapOutput() ScalingGroupStatusMapOutput
	ToScalingGroupStatusMapOutputWithContext(context.Context) ScalingGroupStatusMapOutput
}

type ScalingGroupStatusMap map[string]ScalingGroupStatusInput

func (ScalingGroupStatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroupStatus)(nil)).Elem()
}

func (i ScalingGroupStatusMap) ToScalingGroupStatusMapOutput() ScalingGroupStatusMapOutput {
	return i.ToScalingGroupStatusMapOutputWithContext(context.Background())
}

func (i ScalingGroupStatusMap) ToScalingGroupStatusMapOutputWithContext(ctx context.Context) ScalingGroupStatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupStatusMapOutput)
}

type ScalingGroupStatusOutput struct{ *pulumi.OutputState }

func (ScalingGroupStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroupStatus)(nil)).Elem()
}

func (o ScalingGroupStatusOutput) ToScalingGroupStatusOutput() ScalingGroupStatusOutput {
	return o
}

func (o ScalingGroupStatusOutput) ToScalingGroupStatusOutputWithContext(ctx context.Context) ScalingGroupStatusOutput {
	return o
}

// Scaling group ID.
func (o ScalingGroupStatusOutput) AutoScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroupStatus) pulumi.StringOutput { return v.AutoScalingGroupId }).(pulumi.StringOutput)
}

// If enable auto scaling group.
func (o ScalingGroupStatusOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScalingGroupStatus) pulumi.BoolOutput { return v.Enable }).(pulumi.BoolOutput)
}

type ScalingGroupStatusArrayOutput struct{ *pulumi.OutputState }

func (ScalingGroupStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroupStatus)(nil)).Elem()
}

func (o ScalingGroupStatusArrayOutput) ToScalingGroupStatusArrayOutput() ScalingGroupStatusArrayOutput {
	return o
}

func (o ScalingGroupStatusArrayOutput) ToScalingGroupStatusArrayOutputWithContext(ctx context.Context) ScalingGroupStatusArrayOutput {
	return o
}

func (o ScalingGroupStatusArrayOutput) Index(i pulumi.IntInput) ScalingGroupStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingGroupStatus {
		return vs[0].([]*ScalingGroupStatus)[vs[1].(int)]
	}).(ScalingGroupStatusOutput)
}

type ScalingGroupStatusMapOutput struct{ *pulumi.OutputState }

func (ScalingGroupStatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroupStatus)(nil)).Elem()
}

func (o ScalingGroupStatusMapOutput) ToScalingGroupStatusMapOutput() ScalingGroupStatusMapOutput {
	return o
}

func (o ScalingGroupStatusMapOutput) ToScalingGroupStatusMapOutputWithContext(ctx context.Context) ScalingGroupStatusMapOutput {
	return o
}

func (o ScalingGroupStatusMapOutput) MapIndex(k pulumi.StringInput) ScalingGroupStatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingGroupStatus {
		return vs[0].(map[string]*ScalingGroupStatus)[vs[1].(string)]
	}).(ScalingGroupStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupStatusInput)(nil)).Elem(), &ScalingGroupStatus{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupStatusArrayInput)(nil)).Elem(), ScalingGroupStatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupStatusMapInput)(nil)).Elem(), ScalingGroupStatusMap{})
	pulumi.RegisterOutputType(ScalingGroupStatusOutput{})
	pulumi.RegisterOutputType(ScalingGroupStatusArrayOutput{})
	pulumi.RegisterOutputType(ScalingGroupStatusMapOutput{})
}