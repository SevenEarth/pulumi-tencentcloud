// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tem scaleRule
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tem"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tem.NewScaleRule(ctx, "scaleRule", &Tem.ScaleRuleArgs{
//				EnvironmentId: pulumi.String("en-o5edaepv"),
//				ApplicationId: pulumi.String("app-3j29aa2p"),
//				WorkloadId:    pulumi.Any(resource.Tencentcloud_tem_workload.Workload.Id),
//				Autoscaler: &tem.ScaleRuleAutoscalerArgs{
//					AutoscalerName: pulumi.String("test3123"),
//					Description:    pulumi.String("test"),
//					Enabled:        pulumi.Bool(true),
//					MinReplicas:    pulumi.Int(1),
//					MaxReplicas:    pulumi.Int(4),
//					CronHorizontalAutoscalers: tem.ScaleRuleAutoscalerCronHorizontalAutoscalerArray{
//						&tem.ScaleRuleAutoscalerCronHorizontalAutoscalerArgs{
//							Name:     pulumi.String("test"),
//							Period:   pulumi.String("* * *"),
//							Priority: pulumi.Int(1),
//							Enabled:  pulumi.Bool(true),
//							Schedules: tem.ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleArray{
//								&tem.ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleArgs{
//									StartAt:        pulumi.String("03:00"),
//									TargetReplicas: pulumi.Int(1),
//								},
//							},
//						},
//						&tem.ScaleRuleAutoscalerCronHorizontalAutoscalerArgs{
//							Name:     pulumi.String("test123123"),
//							Period:   pulumi.String("* * *"),
//							Priority: pulumi.Int(0),
//							Enabled:  pulumi.Bool(true),
//							Schedules: tem.ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleArray{
//								&tem.ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleArgs{
//									StartAt:        pulumi.String("04:13"),
//									TargetReplicas: pulumi.Int(1),
//								},
//							},
//						},
//					},
//					HorizontalAutoscalers: tem.ScaleRuleAutoscalerHorizontalAutoscalerArray{
//						&tem.ScaleRuleAutoscalerHorizontalAutoscalerArgs{
//							Metrics:     pulumi.String("CPU"),
//							Enabled:     pulumi.Bool(true),
//							MaxReplicas: pulumi.Int(4),
//							MinReplicas: pulumi.Int(1),
//							Threshold:   pulumi.Int(60),
//						},
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
// tem scaleRule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tem/scaleRule:ScaleRule scaleRule environmentId#applicationId#scaleRuleId
// ```
type ScaleRule struct {
	pulumi.CustomResourceState

	// application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// .
	Autoscaler ScaleRuleAutoscalerOutput `pulumi:"autoscaler"`
	// environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringOutput `pulumi:"workloadId"`
}

// NewScaleRule registers a new resource with the given unique name, arguments, and options.
func NewScaleRule(ctx *pulumi.Context,
	name string, args *ScaleRuleArgs, opts ...pulumi.ResourceOption) (*ScaleRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Autoscaler == nil {
		return nil, errors.New("invalid value for required argument 'Autoscaler'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.WorkloadId == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScaleRule
	err := ctx.RegisterResource("tencentcloud:Tem/scaleRule:ScaleRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScaleRule gets an existing ScaleRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScaleRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScaleRuleState, opts ...pulumi.ResourceOption) (*ScaleRule, error) {
	var resource ScaleRule
	err := ctx.ReadResource("tencentcloud:Tem/scaleRule:ScaleRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScaleRule resources.
type scaleRuleState struct {
	// application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// .
	Autoscaler *ScaleRuleAutoscaler `pulumi:"autoscaler"`
	// environment ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId *string `pulumi:"workloadId"`
}

type ScaleRuleState struct {
	// application ID.
	ApplicationId pulumi.StringPtrInput
	// .
	Autoscaler ScaleRuleAutoscalerPtrInput
	// environment ID.
	EnvironmentId pulumi.StringPtrInput
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringPtrInput
}

func (ScaleRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleRuleState)(nil)).Elem()
}

type scaleRuleArgs struct {
	// application ID.
	ApplicationId string `pulumi:"applicationId"`
	// .
	Autoscaler ScaleRuleAutoscaler `pulumi:"autoscaler"`
	// environment ID.
	EnvironmentId string `pulumi:"environmentId"`
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId string `pulumi:"workloadId"`
}

// The set of arguments for constructing a ScaleRule resource.
type ScaleRuleArgs struct {
	// application ID.
	ApplicationId pulumi.StringInput
	// .
	Autoscaler ScaleRuleAutoscalerInput
	// environment ID.
	EnvironmentId pulumi.StringInput
	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	WorkloadId pulumi.StringInput
}

func (ScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleRuleArgs)(nil)).Elem()
}

type ScaleRuleInput interface {
	pulumi.Input

	ToScaleRuleOutput() ScaleRuleOutput
	ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput
}

func (*ScaleRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleRule)(nil)).Elem()
}

func (i *ScaleRule) ToScaleRuleOutput() ScaleRuleOutput {
	return i.ToScaleRuleOutputWithContext(context.Background())
}

func (i *ScaleRule) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleOutput)
}

// ScaleRuleArrayInput is an input type that accepts ScaleRuleArray and ScaleRuleArrayOutput values.
// You can construct a concrete instance of `ScaleRuleArrayInput` via:
//
//	ScaleRuleArray{ ScaleRuleArgs{...} }
type ScaleRuleArrayInput interface {
	pulumi.Input

	ToScaleRuleArrayOutput() ScaleRuleArrayOutput
	ToScaleRuleArrayOutputWithContext(context.Context) ScaleRuleArrayOutput
}

type ScaleRuleArray []ScaleRuleInput

func (ScaleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArray) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return i.ToScaleRuleArrayOutputWithContext(context.Background())
}

func (i ScaleRuleArray) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleArrayOutput)
}

// ScaleRuleMapInput is an input type that accepts ScaleRuleMap and ScaleRuleMapOutput values.
// You can construct a concrete instance of `ScaleRuleMapInput` via:
//
//	ScaleRuleMap{ "key": ScaleRuleArgs{...} }
type ScaleRuleMapInput interface {
	pulumi.Input

	ToScaleRuleMapOutput() ScaleRuleMapOutput
	ToScaleRuleMapOutputWithContext(context.Context) ScaleRuleMapOutput
}

type ScaleRuleMap map[string]ScaleRuleInput

func (ScaleRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScaleRule)(nil)).Elem()
}

func (i ScaleRuleMap) ToScaleRuleMapOutput() ScaleRuleMapOutput {
	return i.ToScaleRuleMapOutputWithContext(context.Background())
}

func (i ScaleRuleMap) ToScaleRuleMapOutputWithContext(ctx context.Context) ScaleRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleMapOutput)
}

type ScaleRuleOutput struct{ *pulumi.OutputState }

func (ScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleRule)(nil)).Elem()
}

func (o ScaleRuleOutput) ToScaleRuleOutput() ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return o
}

// application ID.
func (o ScaleRuleOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScaleRule) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// .
func (o ScaleRuleOutput) Autoscaler() ScaleRuleAutoscalerOutput {
	return o.ApplyT(func(v *ScaleRule) ScaleRuleAutoscalerOutput { return v.Autoscaler }).(ScaleRuleAutoscalerOutput)
}

// environment ID.
func (o ScaleRuleOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScaleRule) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
func (o ScaleRuleOutput) WorkloadId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScaleRule) pulumi.StringOutput { return v.WorkloadId }).(pulumi.StringOutput)
}

type ScaleRuleArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScaleRule)(nil)).Elem()
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) Index(i pulumi.IntInput) ScaleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScaleRule {
		return vs[0].([]*ScaleRule)[vs[1].(int)]
	}).(ScaleRuleOutput)
}

type ScaleRuleMapOutput struct{ *pulumi.OutputState }

func (ScaleRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScaleRule)(nil)).Elem()
}

func (o ScaleRuleMapOutput) ToScaleRuleMapOutput() ScaleRuleMapOutput {
	return o
}

func (o ScaleRuleMapOutput) ToScaleRuleMapOutputWithContext(ctx context.Context) ScaleRuleMapOutput {
	return o
}

func (o ScaleRuleMapOutput) MapIndex(k pulumi.StringInput) ScaleRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScaleRule {
		return vs[0].(map[string]*ScaleRule)[vs[1].(string)]
	}).(ScaleRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleRuleInput)(nil)).Elem(), &ScaleRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleRuleArrayInput)(nil)).Elem(), ScaleRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleRuleMapInput)(nil)).Elem(), ScaleRuleMap{})
	pulumi.RegisterOutputType(ScaleRuleOutput{})
	pulumi.RegisterOutputType(ScaleRuleArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleMapOutput{})
}
